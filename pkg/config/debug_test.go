package config

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cadenya/cadenya-go"
	"github.com/cadenya/cadenya-go/option"
)

// captureLog redirects the default logger to a buffer for the duration of
// the test and returns the buffer.
func captureLog(t *testing.T) *bytes.Buffer {
	t.Helper()
	var buf bytes.Buffer
	origWriter := log.Writer()
	origFlags := log.Flags()
	log.SetOutput(&buf)
	log.SetFlags(0)
	t.Cleanup(func() {
		log.SetOutput(origWriter)
		log.SetFlags(origFlags)
	})
	return &buf
}

// runThroughMiddleware sends a single request through a client built with the
// given middleware, against an httptest server with a fixed response.
func runThroughMiddleware(t *testing.T, mw option.Middleware, reqBody string, respBody string) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, respBody)
	}))
	t.Cleanup(srv.Close)
	client := cadenya.NewClient(
		option.WithBaseURL(srv.URL+"/"),
		option.WithAPIKey("test-key"),
		option.WithMiddleware(mw),
	)
	// Any POST with a body gets the request-body dump we want to assert on.
	_, err := client.ToolSets.New(
		context.Background(),
		cadenya.ToolSetNewParams{},
		option.WithRequestBody("application/json", []byte(reqBody)),
	)
	if err != nil {
		// Don't fail on API error — mock server's body may not parse; we only
		// care about middleware behavior.
		_ = err
	}
}

func TestDebugMiddlewareRedactsSubstitutedSecrets(t *testing.T) {
	secrets := []Secret{
		{Name: "LINEAR_TOKEN", Value: "lk_secret_payload_abcdef"},
	}
	buf := captureLog(t)

	runThroughMiddleware(t,
		NewDebugMiddleware(secrets),
		`{"metadata":{"name":"x"},"spec":{"adapter":{"mcp":{"headers":{"authorization":"Bearer lk_secret_payload_abcdef"}}}}}`,
		`{"metadata":{"id":"ts_01"}}`,
	)

	out := buf.String()
	if strings.Contains(out, "lk_secret_payload_abcdef") {
		t.Errorf("debug log leaked the secret value:\n%s", out)
	}
	if !strings.Contains(out, "<$LINEAR_TOKEN>") {
		t.Errorf("expected redacted placeholder <$LINEAR_TOKEN> in log:\n%s", out)
	}
}

func TestDebugMiddlewareLeavesShortValuesAlone(t *testing.T) {
	// 3-char value: too short to safely redact (would mangle unrelated "yes"
	// occurrences). Passes through unchanged.
	secrets := []Secret{{Name: "TINY", Value: "yes"}}
	buf := captureLog(t)
	runThroughMiddleware(t,
		NewDebugMiddleware(secrets),
		`{"comment":"oh yes"}`,
		`{"ok":true}`,
	)
	out := buf.String()
	if !strings.Contains(out, "oh yes") {
		t.Errorf("short value should have passed through unredacted:\n%s", out)
	}
	if strings.Contains(out, "<$TINY>") {
		t.Errorf("short value should NOT have been replaced:\n%s", out)
	}
}

func TestDebugMiddlewareRedactsAuthorizationHeader(t *testing.T) {
	buf := captureLog(t)
	runThroughMiddleware(t, NewDebugMiddleware(nil), `{}`, `{}`)
	out := buf.String()
	// The SDK injects Authorization: Bearer <api-key> from WithAPIKey. We
	// redact the token but keep the scheme.
	if strings.Contains(out, "Bearer test-key") {
		t.Errorf("Authorization header should be redacted:\n%s", out)
	}
	if !strings.Contains(out, "Authorization: Bearer <REDACTED>") {
		t.Errorf("expected 'Authorization: Bearer <REDACTED>' in log:\n%s", out)
	}
}

func TestDebugMiddlewareRedactsSecretInResponseBody(t *testing.T) {
	secrets := []Secret{{Name: "API_TOKEN", Value: "long_enough_token_xyz"}}
	buf := captureLog(t)
	runThroughMiddleware(t,
		NewDebugMiddleware(secrets),
		`{}`,
		// Server echoes the token back (e.g. a webhook secret).
		`{"metadata":{"id":"ts_01","echoedSecret":"long_enough_token_xyz"}}`,
	)
	out := buf.String()
	if strings.Contains(out, "long_enough_token_xyz") {
		t.Errorf("response body should have redacted secret:\n%s", out)
	}
	if !strings.Contains(out, "<$API_TOKEN>") {
		t.Errorf("expected <$API_TOKEN> in response log:\n%s", out)
	}
}

func TestParseReturnsSubstitutedSecrets(t *testing.T) {
	src := `version: 1
tool_sets:
  x:
    name: "X"
    spec:
      adapter:
        mcp:
          headers:
            authorization: "Bearer $MY_TOKEN"
            other: "${OTHER}"
`
	_, secrets, err := Parse(strings.NewReader(src), map[string]string{
		"MY_TOKEN": "abcdef_long_enough",
		"OTHER":    "other_long_enough",
		"UNUSED":   "never_referenced",
	})
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}
	names := map[string]string{}
	for _, s := range secrets {
		names[s.Name] = s.Value
	}
	if names["MY_TOKEN"] != "abcdef_long_enough" {
		t.Errorf("expected MY_TOKEN captured with its value, got: %+v", secrets)
	}
	if names["OTHER"] != "other_long_enough" {
		t.Errorf("expected OTHER captured, got: %+v", secrets)
	}
	if _, unused := names["UNUSED"]; unused {
		t.Errorf("UNUSED was never referenced in YAML; should not be in secrets: %+v", secrets)
	}
}
