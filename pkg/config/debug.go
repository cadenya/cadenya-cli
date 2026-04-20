package config

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/cadenya/cadenya-go/option"
)

// minRedactLength is the threshold below which a substituted env-var value
// is considered too short to safely redact — replacing every occurrence of
// a 1-2 character value would mangle unrelated content in debug dumps.
// 8 is the conventional minimum for "probably-a-token"; shorter values that
// happen to be secrets are opt-outs the user accepts by using short names.
const minRedactLength = 8

// NewDebugMiddleware returns an SDK middleware that logs request/response
// pairs to log.Default() with two layers of redaction applied to the logged
// bytes (not the wire traffic):
//
//  1. Authorization header values are replaced with <REDACTED> (mirrors the
//     generated debug middleware's behavior, preserving the auth scheme
//     prefix like "Bearer " for debugging).
//  2. Any env-var value from `secrets` that is at least minRedactLength
//     characters long is replaced with "<$NAME>" wherever it appears in the
//     dumped bytes — both in request bodies (e.g. a Bearer token that came
//     from $API_TOKEN via YAML substitution) and response bodies (in case
//     the server echoes back what we sent).
//
// The wire traffic is never modified; only the log output is sanitized.
func NewDebugMiddleware(secrets []Secret) option.Middleware {
	r := newSecretRedactor(secrets)
	return func(req *http.Request, next func(*http.Request) (*http.Response, error)) (*http.Response, error) {
		if dump, err := dumpRequestRedacted(req, r); err == nil {
			log.Printf("Request Content:\n%s\n", dump)
		}

		resp, err := next(req)
		if err != nil {
			return resp, err
		}

		if dump, err := dumpResponseRedacted(resp, r); err == nil {
			log.Printf("Response Content:\n%s\n", dump)
		}
		return resp, err
	}
}

// secretRedactor applies all-at-once string replacement for substituted
// env-var values. Built on strings.NewReplacer which picks the longest
// match at each position and doesn't re-scan output (so redacted tokens
// can't re-match another replacement).
type secretRedactor struct {
	replacer *strings.Replacer
	active   bool
}

func newSecretRedactor(secrets []Secret) *secretRedactor {
	var pairs []string
	for _, s := range secrets {
		if len(s.Value) < minRedactLength {
			continue
		}
		pairs = append(pairs, s.Value, "<$"+s.Name+">")
	}
	if len(pairs) == 0 {
		return &secretRedactor{active: false}
	}
	return &secretRedactor{replacer: strings.NewReplacer(pairs...), active: true}
}

func (r *secretRedactor) Redact(b []byte) []byte {
	if !r.active {
		return b
	}
	return []byte(r.replacer.Replace(string(b)))
}

// dumpRequestRedacted clones req to mutate its Authorization header safely,
// dumps it, then applies secret redaction over the bytes.
func dumpRequestRedacted(req *http.Request, r *secretRedactor) ([]byte, error) {
	redacted, body, err := cloneRequestForLogging(req)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Body = body // restore so downstream middleware / transport see the original
	}
	dump, err := httputil.DumpRequestOut(redacted, true)
	if err != nil {
		return nil, err
	}
	return r.Redact(dump), nil
}

func dumpResponseRedacted(resp *http.Response, r *secretRedactor) ([]byte, error) {
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}
	return r.Redact(dump), nil
}

// cloneRequestForLogging produces a clone of req with its Authorization
// header replaced by "<scheme> <REDACTED>" (or "<REDACTED>" if the header
// has no scheme prefix). The original body is split so the caller can
// restore it on the unmodified request.
func cloneRequestForLogging(req *http.Request) (*http.Request, io.ReadCloser, error) {
	redacted := req.Clone(req.Context())
	redacted.Header = req.Header.Clone()

	if values := redacted.Header.Values("Authorization"); len(values) > 0 {
		redacted.Header.Del("Authorization")
		for _, v := range values {
			if scheme, _, ok := strings.Cut(v, " "); ok {
				redacted.Header.Add("Authorization", scheme+" <REDACTED>")
			} else {
				redacted.Header.Add("Authorization", "<REDACTED>")
			}
		}
	}

	if req.Body == nil || req.Body == http.NoBody {
		return redacted, nil, nil
	}
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, req.Body); err != nil {
		return nil, nil, err
	}
	if err := req.Body.Close(); err != nil {
		return nil, nil, err
	}
	redacted.Body = io.NopCloser(bytes.NewReader(buf.Bytes()))
	original := io.NopCloser(bytes.NewReader(buf.Bytes()))
	return redacted, original, nil
}
