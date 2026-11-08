package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	cli "github.com/urfave/cli/v3"
)

// isolate gives a test its own HOME and working directory and clears every
// variable that changes credential resolution, CI detection or agent
// detection. The CLI reads the working directory, so tests in this file do
// not run in parallel.
func isolate(t *testing.T) string {
	t.Helper()
	home := t.TempDir()
	project := t.TempDir()
	t.Setenv("HOME", home)
	unset := []string{"CADENYA_API_KEY", "CADENYA_BASE_URL", "CADENYA_WORKSPACE_ID", "CADENYA_NEW_BASE_URL", "CADENYA_DEBUG", "CADENYA_AUTH_BASE_URL", "CI"}
	unset = append(unset, codingAgentEnv...)
	for _, kv := range os.Environ() {
		name, _, _ := strings.Cut(kv, "=")
		for _, prefix := range codingAgentEnvPrefixes {
			if strings.HasPrefix(name, prefix) {
				unset = append(unset, name)
			}
		}
	}
	for _, name := range unset {
		t.Setenv(name, "")
		os.Unsetenv(name)
	}
	prev, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(project); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(prev) })
	return project
}

type result struct {
	stdout, stderr string
	code           int
}

// run executes the CLI in-process and maps errors to exit codes the way
// main does.
func run(t *testing.T, args ...string) result {
	t.Helper()
	var stdout, stderr bytes.Buffer
	root := newRootCommand()
	root.Writer = &stdout
	root.ErrWriter = &stderr
	root.ExitErrHandler = func(context.Context, *cli.Command, error) {}
	err := root.Run(context.Background(), append([]string{"cadenya"}, args...))
	code := 0
	if err != nil {
		var exitErr cli.ExitCoder
		if errors.As(err, &exitErr) {
			code = exitErr.ExitCode()
			stderr.WriteString("error: " + err.Error() + "\n")
		} else {
			code = reportError(&stderr, err)
		}
	}
	return result{stdout: stdout.String(), stderr: stderr.String(), code: code}
}

// fakeProvisioning stands in for new.cadenya.com.
type fakeProvisioning struct {
	*httptest.Server
	mu       sync.Mutex
	requests []map[string]any
	created  int
}

func newFakeProvisioning(t *testing.T, expiresIn time.Duration) *fakeProvisioning {
	t.Helper()
	f := &fakeProvisioning{}
	f.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v1/temporary_workspaces" {
			http.NotFound(w, r)
			return
		}
		if r.Header.Get("Authorization") != "" {
			http.Error(w, "credentials present", http.StatusBadRequest)
			return
		}
		var body map[string]any
		_ = json.NewDecoder(r.Body).Decode(&body)
		f.mu.Lock()
		f.requests = append(f.requests, body)
		f.created++
		n := f.created
		f.mu.Unlock()
		id := "workspace_" + string(rune('A'+n-1))
		_ = json.NewEncoder(w).Encode(map[string]any{
			"workspaceId": id,
			"accountId":   "account_1",
			"apiKey":      "sk_temp_" + id,
			"claimUrl":    "https://app.cadenya.com/claim/tok_" + id,
			"claimToken":  "tok_" + id,
			"expiresAt":   time.Now().Add(expiresIn).UTC().Format(time.RFC3339),
			"nextStep":    "Open the claim link to keep " + id + ".",
		})
	}))
	t.Cleanup(f.Close)
	return f
}

func (f *fakeProvisioning) lastRequest(t *testing.T) map[string]any {
	t.Helper()
	f.mu.Lock()
	defer f.mu.Unlock()
	if len(f.requests) == 0 {
		t.Fatal("no create request was sent")
	}
	return f.requests[len(f.requests)-1]
}

func createJSON(t *testing.T, newBase string, extra ...string) (temporaryEnvelope, result) {
	t.Helper()
	args := append([]string{"workspaces", "create", "--temporary", "--new-base-url", newBase, "--display", "json"}, extra...)
	res := run(t, args...)
	if res.code != 0 {
		t.Fatalf("create exited %d: %s", res.code, res.stderr)
	}
	var env temporaryEnvelope
	if err := json.Unmarshal([]byte(res.stdout), &env); err != nil {
		t.Fatalf("stdout is not one JSON envelope: %v\n%s", err, res.stdout)
	}
	return env, res
}

func TestCreateTemporaryJSONEnvelope(t *testing.T) {
	project := isolate(t)
	prov := newFakeProvisioning(t, 24*time.Hour)

	env, res := createJSON(t, prov.URL)

	var raw map[string]any
	_ = json.Unmarshal([]byte(res.stdout), &raw)
	for _, key := range []string{"status", "workspaceId", "claimUrl", "expiresAt", "apiKey", "reused", "nextStep", "next"} {
		if _, ok := raw[key]; !ok {
			t.Errorf("envelope is missing %q: %s", key, res.stdout)
		}
	}
	if len(raw) != 8 {
		t.Errorf("envelope has %d fields, want exactly 8: %s", len(raw), res.stdout)
	}
	if env.Status != "created" || env.Reused || env.WorkspaceID != "workspace_A" || env.APIKey != "sk_temp_workspace_A" {
		t.Errorf("unexpected envelope: %+v", env)
	}
	if env.NextStep != "Open the claim link to keep workspace_A." {
		t.Errorf("nextStep = %q, want the server's sentence", env.NextStep)
	}
	if len(env.Next) == 0 || env.Next[0].Command == "" || env.Next[0].When == "" {
		t.Errorf("next commands missing: %+v", env.Next)
	}
	if _, err := time.Parse(time.RFC3339, env.ExpiresAt); err != nil {
		t.Errorf("expiresAt %q is not RFC 3339", env.ExpiresAt)
	}
	if !strings.Contains(res.stderr, "This workspace expires in 24 hours") || !strings.Contains(res.stderr, env.ClaimURL) {
		t.Errorf("claim notice missing from stderr: %s", res.stderr)
	}
	if got := prov.lastRequest(t)["source"]; got != "TEMPORARY_WORKSPACE_SOURCE_CLI" {
		t.Errorf("source = %v, want CLI", got)
	}

	path := filepath.Join(project, ".cadenya", "temporary.json")
	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if info.Mode().Perm() != 0o600 {
		t.Errorf("temporary.json mode = %o, want 600", info.Mode().Perm())
	}
	var stored temporaryWorkspace
	data, _ := os.ReadFile(path)
	if err := json.Unmarshal(data, &stored); err != nil {
		t.Fatal(err)
	}
	if stored.WorkspaceID != "workspace_A" || stored.APIKey != "sk_temp_workspace_A" || stored.ClaimToken != "tok_workspace_A" || stored.ClaimURL == "" || stored.ExpiresAt.IsZero() {
		t.Errorf("stored file incomplete: %+v", stored)
	}
}

func TestCreateTemporaryHumanOutput(t *testing.T) {
	isolate(t)
	prov := newFakeProvisioning(t, 24*time.Hour)
	res := run(t, "workspaces", "new", "--temporary", "--new-base-url", prov.URL)
	if res.code != 0 {
		t.Fatalf("exit %d: %s", res.code, res.stderr)
	}
	if strings.Contains(res.stdout, "sk_temp_") {
		t.Errorf("human output printed the API key: %s", res.stdout)
	}
	if strings.Contains(res.stdout, "claim/") {
		t.Errorf("claim notice belongs on stderr, not stdout: %s", res.stdout)
	}
	if !strings.Contains(res.stderr, "Claim it to keep it (don't share this link): https://app.cadenya.com/claim/tok_workspace_A") {
		t.Errorf("stderr notice: %s", res.stderr)
	}
}

func TestCreateTemporaryAgentSourceAndFlags(t *testing.T) {
	isolate(t)
	t.Setenv("CLAUDECODE", "1")
	prov := newFakeProvisioning(t, 6*time.Hour)
	createJSON(t, prov.URL, "--name", "demo", "--expire-after", "6h")
	req := prov.lastRequest(t)
	if req["source"] != "TEMPORARY_WORKSPACE_SOURCE_AGENT" {
		t.Errorf("source = %v, want AGENT", req["source"])
	}
	if req["name"] != "demo" || req["expireAfter"] != "21600s" {
		t.Errorf("request = %v", req)
	}
}

func TestCreateTemporaryExpireAfterRange(t *testing.T) {
	isolate(t)
	prov := newFakeProvisioning(t, time.Hour)
	for _, v := range []string{"30m", "25h"} {
		res := run(t, "workspaces", "create", "--temporary", "--new-base-url", prov.URL, "--expire-after", v)
		if res.code != 2 {
			t.Errorf("--expire-after %s exited %d, want 2", v, res.code)
		}
	}
	if prov.created != 0 {
		t.Errorf("an out-of-range request reached the server")
	}
}

func TestCreateTemporaryReuseAndFresh(t *testing.T) {
	isolate(t)
	prov := newFakeProvisioning(t, 24*time.Hour)
	first, _ := createJSON(t, prov.URL)

	again, res := createJSON(t, prov.URL)
	if !again.Reused || again.Status != "reused" || again.WorkspaceID != first.WorkspaceID || again.APIKey != first.APIKey {
		t.Errorf("second create did not reuse: %+v", again)
	}
	if again.ExpiresAt != first.ExpiresAt {
		t.Errorf("reuse changed the expiry: %s -> %s", first.ExpiresAt, again.ExpiresAt)
	}
	if again.NextStep != first.NextStep {
		t.Errorf("reuse lost nextStep: %q", again.NextStep)
	}
	if prov.created != 1 {
		t.Errorf("reuse sent a create request (%d total)", prov.created)
	}
	if !strings.Contains(res.stderr, "Claim it to keep it") {
		t.Errorf("reuse should repeat the claim notice: %s", res.stderr)
	}

	fresh, _ := createJSON(t, prov.URL, "--fresh")
	if fresh.Reused || fresh.WorkspaceID == first.WorkspaceID {
		t.Errorf("--fresh reused: %+v", fresh)
	}
	tw, _ := readTemporaryWorkspace()
	if tw.WorkspaceID != fresh.WorkspaceID {
		t.Errorf("--fresh did not replace the file: %s", tw.WorkspaceID)
	}
}

func TestCreateTemporaryReplacesExpiredFile(t *testing.T) {
	isolate(t)
	prov := newFakeProvisioning(t, 24*time.Hour)
	if err := writeTemporaryWorkspace(&temporaryWorkspace{WorkspaceID: "workspace_old", APIKey: "sk_old", ClaimURL: "u", ExpiresAt: time.Now().Add(-time.Minute)}); err != nil {
		t.Fatal(err)
	}
	env, _ := createJSON(t, prov.URL)
	if env.Reused || env.WorkspaceID == "workspace_old" {
		t.Errorf("an expired file was reused: %+v", env)
	}
}

func TestCreateTemporaryGitignore(t *testing.T) {
	t.Run("created", func(t *testing.T) {
		project := isolate(t)
		prov := newFakeProvisioning(t, time.Hour)
		_, res := createJSON(t, prov.URL)
		data, err := os.ReadFile(filepath.Join(project, ".gitignore"))
		if err != nil || string(data) != ".cadenya/temporary.json\n" {
			t.Errorf(".gitignore = %q, %v", data, err)
		}
		if !strings.Contains(res.stderr, "Created .gitignore with .cadenya/temporary.json") {
			t.Errorf("did not say so: %s", res.stderr)
		}
	})
	t.Run("appended", func(t *testing.T) {
		project := isolate(t)
		if err := os.WriteFile(filepath.Join(project, ".gitignore"), []byte("node_modules"), 0o644); err != nil {
			t.Fatal(err)
		}
		prov := newFakeProvisioning(t, time.Hour)
		_, res := createJSON(t, prov.URL)
		data, _ := os.ReadFile(filepath.Join(project, ".gitignore"))
		if string(data) != "node_modules\n.cadenya/temporary.json\n" {
			t.Errorf(".gitignore = %q", data)
		}
		if !strings.Contains(res.stderr, "Added .cadenya/temporary.json to .gitignore") {
			t.Errorf("did not say so: %s", res.stderr)
		}
	})
	t.Run("already ignored", func(t *testing.T) {
		project := isolate(t)
		if err := os.WriteFile(filepath.Join(project, ".gitignore"), []byte("/.cadenya/\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		prov := newFakeProvisioning(t, time.Hour)
		_, res := createJSON(t, prov.URL)
		data, _ := os.ReadFile(filepath.Join(project, ".gitignore"))
		if string(data) != "/.cadenya/\n" {
			t.Errorf(".gitignore changed: %q", data)
		}
		if strings.Contains(res.stderr, ".gitignore") {
			t.Errorf("reported a change that did not happen: %s", res.stderr)
		}
	})
}

func TestCreateTemporaryRefusals(t *testing.T) {
	cases := map[string]func(t *testing.T) []string{
		"api key flag": func(t *testing.T) []string { return []string{"--api-key", "sk_live"} },
		"api key env": func(t *testing.T) []string {
			t.Setenv("CADENYA_API_KEY", "sk_live")
			return nil
		},
		"stored login": func(t *testing.T) []string {
			if err := saveCredentials(map[string]credProfile{"default": {Credential: "sk_login"}}); err != nil {
				t.Fatal(err)
			}
			return nil
		},
		"ci": func(t *testing.T) []string {
			t.Setenv("CI", "true")
			return nil
		},
	}
	for name, setup := range cases {
		t.Run(name, func(t *testing.T) {
			isolate(t)
			prov := newFakeProvisioning(t, time.Hour)
			global := setup(t)
			args := append(global, "workspaces", "create", "--temporary", "--new-base-url", prov.URL)
			res := run(t, args...)
			if res.code == 0 {
				t.Fatalf("not refused: %s", res.stdout)
			}
			if prov.created != 0 {
				t.Errorf("a refused create reached the server")
			}
			want := "cadenya auth login"
			if name == "ci" {
				want = "CADENYA_API_KEY"
			}
			if !strings.Contains(res.stderr, want) {
				t.Errorf("refusal should name the alternative %q: %s", want, res.stderr)
			}
		})
	}
}

func TestCreateWithoutTemporaryIsUsageError(t *testing.T) {
	isolate(t)
	res := run(t, "workspaces", "create")
	if res.code != 2 || !strings.Contains(res.stderr, "workspace-admin create") {
		t.Errorf("exit %d: %s", res.code, res.stderr)
	}
}

func TestCreateTemporaryEdgeErrors(t *testing.T) {
	t.Run("rate limited", func(t *testing.T) {
		isolate(t)
		calls := 0
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			calls++
			w.Header().Set("Retry-After", "120")
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = io.WriteString(w, `{"error":{"code":"RATE_LIMITED","message":"slow down","retryable":true}}`)
		}))
		defer srv.Close()
		res := run(t, "workspaces", "create", "--temporary", "--new-base-url", srv.URL)
		if res.code == 0 || calls != 1 {
			t.Fatalf("exit %d after %d calls", res.code, calls)
		}
		if !strings.Contains(res.stderr, "try again in 2 minutes") {
			t.Errorf("wait time missing: %s", res.stderr)
		}
	})
	t.Run("kill switch", func(t *testing.T) {
		isolate(t)
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = io.WriteString(w, `{"error":{"code":"UNAVAILABLE","message":"disabled"}}`)
		}))
		defer srv.Close()
		res := run(t, "workspaces", "create", "--temporary", "--new-base-url", srv.URL)
		if res.code == 0 || !strings.Contains(res.stderr, "temporary workspaces are unavailable right now") {
			t.Errorf("exit %d: %s", res.code, res.stderr)
		}
	})
	t.Run("status error", func(t *testing.T) {
		isolate(t)
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = io.WriteString(w, `{"code":3,"message":"name: value length must be at most 64 characters"}`)
		}))
		defer srv.Close()
		res := run(t, "workspaces", "create", "--temporary", "--new-base-url", srv.URL)
		if res.code != 1 || !strings.Contains(res.stderr, "at most 64 characters") {
			t.Errorf("exit %d: %s", res.code, res.stderr)
		}
	})
}

func TestNewBaseURLResolution(t *testing.T) {
	isolate(t)
	var got string
	capture := func(t *testing.T, args ...string) {
		t.Helper()
		root := newRootCommand()
		root.Writer, root.ErrWriter = io.Discard, io.Discard
		root.ExitErrHandler = func(context.Context, *cli.Command, error) {}
		for _, c := range root.Commands {
			if c.Name != "workspaces" {
				continue
			}
			for _, sub := range c.Commands {
				if sub.Name == "create" {
					sub.Action = func(_ context.Context, cmd *cli.Command) error {
						got = newBaseURL(cmd)
						return nil
					}
				}
			}
		}
		if err := root.Run(context.Background(), append([]string{"cadenya"}, args...)); err != nil {
			t.Fatal(err)
		}
	}
	capture(t, "workspaces", "create")
	if got != "https://new.cadenya.com" {
		t.Errorf("default = %s", got)
	}
	capture(t, "--base-url", "https://api.cadenya.dev", "workspaces", "create")
	if got != "https://new.cadenya.dev" {
		t.Errorf("staging inference = %s", got)
	}
	t.Setenv("CADENYA_NEW_BASE_URL", "http://localhost:9999/")
	capture(t, "--base-url", "https://api.cadenya.dev", "workspaces", "create")
	if got != "http://localhost:9999" {
		t.Errorf("env override = %s", got)
	}
}

func TestCreateTemporaryInstructions(t *testing.T) {
	project := isolate(t)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "# Start without an account\n")
	}))
	defer srv.Close()
	prev := temporaryInstructionsURL
	temporaryInstructionsURL = srv.URL
	defer func() { temporaryInstructionsURL = prev }()

	res := run(t, "workspaces", "create", "--instructions")
	if res.code != 0 {
		t.Fatalf("exit %d: %s", res.code, res.stderr)
	}
	if !strings.HasPrefix(res.stdout, "Nothing has been created yet") || !strings.Contains(res.stdout, "# Start without an account") {
		t.Errorf("stdout: %s", res.stdout)
	}
	if _, err := os.Stat(filepath.Join(project, ".cadenya")); !os.IsNotExist(err) {
		t.Errorf("--instructions created .cadenya")
	}
}

// fakeAPI stands in for api.cadenya.com: it records the bearer token of
// every request and serves the claim-status endpoint from a script.
type fakeAPI struct {
	*httptest.Server
	mu      sync.Mutex
	tokens  []string
	states  []string
	polls   int
	handler func(w http.ResponseWriter, r *http.Request)
}

func newFakeAPI(t *testing.T) *fakeAPI {
	t.Helper()
	f := &fakeAPI{}
	f.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f.mu.Lock()
		f.tokens = append(f.tokens, strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer "))
		handler := f.handler
		f.mu.Unlock()
		if handler != nil {
			handler(w, r)
			return
		}
		if strings.HasSuffix(r.URL.Path, "/temporary") {
			f.mu.Lock()
			state := f.states[min(f.polls, len(f.states)-1)]
			f.polls++
			f.mu.Unlock()
			_ = json.NewEncoder(w).Encode(map[string]any{
				"workspaceId":  strings.Split(r.URL.Path, "/")[3],
				"state":        state,
				"expiresAt":    time.Now().Add(time.Hour).UTC().Format(time.RFC3339),
				"pollInterval": "0.01s",
			})
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"items": []any{}})
	}))
	t.Cleanup(f.Close)
	return f
}

func (f *fakeAPI) lastToken() string {
	f.mu.Lock()
	defer f.mu.Unlock()
	if len(f.tokens) == 0 {
		return ""
	}
	return f.tokens[len(f.tokens)-1]
}

func writeLiveTemporary(t *testing.T, baseURL string) *temporaryWorkspace {
	t.Helper()
	tw := &temporaryWorkspace{
		WorkspaceID: "workspace_T",
		APIKey:      "sk_temp",
		ClaimURL:    "https://app.cadenya.com/claim/tok_T",
		ClaimToken:  "tok_T",
		ExpiresAt:   time.Now().Add(time.Hour),
		BaseURL:     baseURL,
	}
	if err := writeTemporaryWorkspace(tw); err != nil {
		t.Fatal(err)
	}
	return tw
}

func TestCredentialPrecedenceWithTemporaryFile(t *testing.T) {
	t.Run("temporary file used last", func(t *testing.T) {
		isolate(t)
		api := newFakeAPI(t)
		writeLiveTemporary(t, api.URL)
		res := run(t, "workspaces", "list", "--display", "json")
		if res.code != 0 {
			t.Fatalf("exit %d: %s", res.code, res.stderr)
		}
		if api.lastToken() != "sk_temp" {
			t.Errorf("token = %q, want the temporary key", api.lastToken())
		}
		if strings.Contains(res.stderr, "warning") {
			t.Errorf("unexpected warning: %s", res.stderr)
		}
	})
	overrides := map[string]func(t *testing.T) []string{
		"flag": func(t *testing.T) []string { return []string{"--api-key", "sk_flag"} },
		"env": func(t *testing.T) []string {
			t.Setenv("CADENYA_API_KEY", "sk_env")
			return nil
		},
		"stored login": func(t *testing.T) []string {
			if err := saveCredentials(map[string]credProfile{"default": {Credential: "sk_login"}}); err != nil {
				t.Fatal(err)
			}
			return nil
		},
	}
	want := map[string]string{"flag": "sk_flag", "env": "sk_env", "stored login": "sk_login"}
	for name, setup := range overrides {
		t.Run(name+" outranks the temporary file", func(t *testing.T) {
			isolate(t)
			api := newFakeAPI(t)
			writeLiveTemporary(t, "")
			global := append(setup(t), "--base-url", api.URL)
			res := run(t, append(global, "workspaces", "list", "--display", "json")...)
			if res.code != 0 {
				t.Fatalf("exit %d: %s", res.code, res.stderr)
			}
			if api.lastToken() != want[name] {
				t.Errorf("token = %q, want %q", api.lastToken(), want[name])
			}
			if !strings.Contains(res.stderr, "warning:") || !strings.Contains(res.stderr, ".cadenya/temporary.json") {
				t.Errorf("no override warning: %s", res.stderr)
			}
		})
	}
	t.Run("expired file ignored", func(t *testing.T) {
		isolate(t)
		api := newFakeAPI(t)
		if err := writeTemporaryWorkspace(&temporaryWorkspace{WorkspaceID: "w", APIKey: "sk_old", ExpiresAt: time.Now().Add(-time.Second), BaseURL: api.URL}); err != nil {
			t.Fatal(err)
		}
		res := run(t, "--base-url", api.URL, "workspaces", "list")
		if res.code == 0 {
			t.Fatalf("an expired temporary key was used")
		}
	})
}

func TestNoCredentialsHint(t *testing.T) {
	isolate(t)
	res := run(t, "workspaces", "list")
	if res.code == 0 || !strings.Contains(res.stderr, temporaryCreateHint) {
		t.Errorf("exit %d, stderr: %s", res.code, res.stderr)
	}
	t.Setenv("CI", "1")
	res = run(t, "workspaces", "list")
	if res.code == 0 || strings.Contains(res.stderr, temporaryCreateHint) {
		t.Errorf("hint shown in CI: %s", res.stderr)
	}
}

func claimedHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = io.WriteString(w, `{"code":16,"message":"This workspace was claimed; its temporary API key was revoked.","details":[{"@type":"type.googleapis.com/google.rpc.ErrorInfo","reason":"TEMPORARY_WORKSPACE_CLAIMED","domain":"api.cadenya.com","metadata":{"workspace_id":"workspace_T"}}]}`)
}

func TestTemporaryWorkspaceClaimedCleanup(t *testing.T) {
	project := isolate(t)
	api := newFakeAPI(t)
	api.handler = claimedHandler
	writeLiveTemporary(t, api.URL)
	res := run(t, "workspaces", "list")
	if res.code != exitTemporaryWorkspaceClaimed {
		t.Errorf("exit %d, want %d", res.code, exitTemporaryWorkspaceClaimed)
	}
	if !strings.Contains(res.stderr, "This workspace was claimed") || !strings.Contains(res.stderr, "cadenya auth login") {
		t.Errorf("stderr: %s", res.stderr)
	}
	if _, err := os.Stat(filepath.Join(project, ".cadenya", "temporary.json")); !os.IsNotExist(err) {
		t.Errorf("temporary.json was not removed")
	}
}

func TestRequiresClaimPrintsClaimURL(t *testing.T) {
	isolate(t)
	api := newFakeAPI(t)
	api.handler = func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = io.WriteString(w, `{"code":7,"message":"Widgets need a claimed workspace.","details":[{"@type":"type.googleapis.com/google.rpc.ErrorInfo","reason":"REQUIRES_CLAIM","domain":"api.cadenya.com","metadata":{"feature":"widgets"}}]}`)
	}
	tw := writeLiveTemporary(t, api.URL)
	res := run(t, "workspaces", "list")
	if res.code != 1 || !strings.Contains(res.stderr, tw.ClaimURL) {
		t.Errorf("exit %d: %s", res.code, res.stderr)
	}
}

func TestClaimStatus(t *testing.T) {
	t.Run("json", func(t *testing.T) {
		isolate(t)
		api := newFakeAPI(t)
		api.states = []string{stateUnclaimed}
		writeLiveTemporary(t, api.URL)
		res := run(t, "workspaces", "claim-status", "--display", "json")
		if res.code != 0 {
			t.Fatalf("exit %d: %s", res.code, res.stderr)
		}
		var status map[string]any
		if err := json.Unmarshal([]byte(res.stdout), &status); err != nil {
			t.Fatal(err)
		}
		if status["state"] != stateUnclaimed || status["workspaceId"] != "workspace_T" || status["expiresAt"] == nil {
			t.Errorf("status = %v", status)
		}
		if api.lastToken() != "sk_temp" {
			t.Errorf("token = %q", api.lastToken())
		}
	})
	t.Run("uses the temporary key even when logged in", func(t *testing.T) {
		isolate(t)
		api := newFakeAPI(t)
		api.states = []string{stateUnclaimed}
		writeLiveTemporary(t, api.URL)
		if err := saveCredentials(map[string]credProfile{"default": {Credential: "sk_login"}}); err != nil {
			t.Fatal(err)
		}
		res := run(t, "workspaces", "claim-status")
		if res.code != 0 || api.lastToken() != "sk_temp" {
			t.Errorf("exit %d, token %q: %s", res.code, api.lastToken(), res.stderr)
		}
	})
	t.Run("wait until claimed", func(t *testing.T) {
		project := isolate(t)
		api := newFakeAPI(t)
		api.states = []string{stateUnclaimed, statePending, stateClaimed}
		writeLiveTemporary(t, api.URL)
		res := run(t, "workspaces", "claim-status", "--wait", "--display", "json")
		if res.code != 0 {
			t.Fatalf("exit %d: %s", res.code, res.stderr)
		}
		if api.polls != 3 {
			t.Errorf("polled %d times, want 3", api.polls)
		}
		var status temporaryStatus
		_ = json.Unmarshal([]byte(res.stdout), &status)
		if status.State != stateClaimed {
			t.Errorf("final state %q", status.State)
		}
		if _, err := os.Stat(filepath.Join(project, ".cadenya", "temporary.json")); !os.IsNotExist(err) {
			t.Errorf("temporary.json kept after the claim")
		}
	})
	t.Run("wait until expired", func(t *testing.T) {
		isolate(t)
		api := newFakeAPI(t)
		api.states = []string{stateUnclaimed, stateExpired}
		writeLiveTemporary(t, api.URL)
		res := run(t, "workspaces", "claim-status", "--wait", "--display", "json")
		if res.code != 0 || !strings.Contains(res.stdout, stateExpired) {
			t.Errorf("exit %d: %s %s", res.code, res.stdout, res.stderr)
		}
	})
	t.Run("revoked key means claimed", func(t *testing.T) {
		project := isolate(t)
		api := newFakeAPI(t)
		api.handler = claimedHandler
		writeLiveTemporary(t, api.URL)
		res := run(t, "workspaces", "claim-status", "--wait", "--display", "json")
		if res.code != 0 || !strings.Contains(res.stdout, stateClaimed) {
			t.Errorf("exit %d: %s %s", res.code, res.stdout, res.stderr)
		}
		if _, err := os.Stat(filepath.Join(project, ".cadenya", "temporary.json")); !os.IsNotExist(err) {
			t.Errorf("temporary.json kept after the claim")
		}
	})
	t.Run("no file", func(t *testing.T) {
		isolate(t)
		res := run(t, "workspaces", "claim-status")
		if res.code == 0 || !strings.Contains(res.stderr, temporaryCreateHint) {
			t.Errorf("exit %d: %s", res.code, res.stderr)
		}
	})
}

func TestAuthStatusNamesTemporaryFile(t *testing.T) {
	isolate(t)
	writeLiveTemporary(t, "")
	res := run(t, "auth", "status")
	if !strings.Contains(res.stdout, ".cadenya/temporary.json") {
		t.Errorf("auth status: %s", res.stdout)
	}
}

// fakeAuth serves the device-authorization flow and records the form the
// CLI sent.
func fakeAuth(t *testing.T) (*httptest.Server, *map[string]string) {
	t.Helper()
	sent := map[string]string{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		switch r.URL.Path {
		case "/api/cli/oauth/device_authorization":
			for k := range r.PostForm {
				sent[k] = r.PostForm.Get(k)
			}
			_ = json.NewEncoder(w).Encode(map[string]any{"device_code": "dc", "user_code": "ABCD", "verification_uri": "https://example.test/approve", "interval": 0, "expires_in": 60})
		case "/api/cli/oauth/token":
			_ = json.NewEncoder(w).Encode(map[string]any{"access_token": "sk_login", "token_type": "bearer", "workspaces": []map[string]string{{"id": "workspace_T", "name": "Temp"}}})
		}
	}))
	t.Cleanup(srv.Close)
	return srv, &sent
}

func TestAuthLoginClaimsTemporaryWorkspace(t *testing.T) {
	prevPoll := authDefaultPollInterval
	authDefaultPollInterval = time.Millisecond
	defer func() { authDefaultPollInterval = prevPoll }()

	t.Run("claimed", func(t *testing.T) {
		project := isolate(t)
		auth, sent := fakeAuth(t)
		t.Setenv("CADENYA_AUTH_BASE_URL", auth.URL)
		api := newFakeAPI(t)
		api.handler = claimedHandler
		writeLiveTemporary(t, api.URL)
		res := run(t, "auth", "login", "--no-browser")
		if res.code != 0 {
			t.Fatalf("exit %d: %s", res.code, res.stderr)
		}
		if (*sent)["claim_token"] != "tok_T" {
			t.Errorf("device authorization form = %v, want claim_token", *sent)
		}
		if _, err := os.Stat(filepath.Join(project, ".cadenya", "temporary.json")); !os.IsNotExist(err) {
			t.Errorf("temporary.json kept after a claiming login")
		}
		if !strings.Contains(res.stdout, "Claimed temporary workspace workspace_T") {
			t.Errorf("stdout: %s", res.stdout)
		}
	})
	t.Run("not claimed", func(t *testing.T) {
		project := isolate(t)
		auth, _ := fakeAuth(t)
		t.Setenv("CADENYA_AUTH_BASE_URL", auth.URL)
		api := newFakeAPI(t)
		api.states = []string{stateUnclaimed}
		writeLiveTemporary(t, api.URL)
		res := run(t, "auth", "login", "--no-browser")
		if res.code != 0 {
			t.Fatalf("login must still succeed: exit %d: %s", res.code, res.stderr)
		}
		if _, err := os.Stat(filepath.Join(project, ".cadenya", "temporary.json")); err != nil {
			t.Errorf("temporary.json removed although the claim failed")
		}
		if !strings.Contains(res.stdout, "was not claimed") {
			t.Errorf("stdout: %s", res.stdout)
		}
	})
	t.Run("no temporary file sends no claim token", func(t *testing.T) {
		isolate(t)
		auth, sent := fakeAuth(t)
		t.Setenv("CADENYA_AUTH_BASE_URL", auth.URL)
		res := run(t, "auth", "login", "--no-browser")
		if res.code != 0 {
			t.Fatalf("exit %d: %s", res.code, res.stderr)
		}
		if _, ok := (*sent)["claim_token"]; ok {
			t.Errorf("claim_token sent without a temporary workspace")
		}
	})
}
