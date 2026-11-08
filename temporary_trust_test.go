package main

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

// plantTemporaryFile writes .cadenya/temporary.json by hand, the way a
// committed or planted file would arrive.
func plantTemporaryFile(t *testing.T, project string, tw map[string]any, perm os.FileMode) {
	t.Helper()
	dir := filepath.Join(project, ".cadenya")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(tw)
	path := filepath.Join(dir, "temporary.json")
	if err := os.WriteFile(path, data, perm); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(path, perm); err != nil {
		t.Fatal(err)
	}
}

func plantedWorkspace(baseURL, claimURL string) map[string]any {
	return map[string]any{
		"workspaceId": "workspace_P",
		"apiKey":      "sk_planted",
		"claimUrl":    claimURL,
		"claimToken":  "tok_P",
		"expiresAt":   time.Now().Add(time.Hour).UTC().Format(time.RFC3339),
		"baseUrl":     baseURL,
	}
}

func TestPlantedTemporaryFileIsIgnored(t *testing.T) {
	t.Run("foreign baseUrl", func(t *testing.T) {
		project := isolate(t)
		var hits atomic.Int32
		evil := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { hits.Add(1) }))
		defer evil.Close()
		plantTemporaryFile(t, project, plantedWorkspace(evil.URL, "https://app.cadenya.com/claim/x"), 0o600)
		res := run(t, "workspaces", "list")
		if res.code == 0 || hits.Load() != 0 {
			t.Fatalf("a foreign baseUrl was used (exit %d, %d hits)", res.code, hits.Load())
		}
		if !strings.Contains(res.stderr, "warning: ignoring .cadenya/temporary.json: baseUrl") {
			t.Errorf("no warning: %s", res.stderr)
		}
		if strings.Count(res.stderr, "ignoring .cadenya/temporary.json") != 1 {
			t.Errorf("warning repeated: %s", res.stderr)
		}
	})
	t.Run("untrusted claimUrl", func(t *testing.T) {
		project := isolate(t)
		plantTemporaryFile(t, project, plantedWorkspace("", "https://evil.example/claim/x"), 0o600)
		res := run(t, "workspaces", "create", "--temporary", "--new-base-url", "http://127.0.0.1:1")
		if strings.Contains(res.stderr+res.stdout, "evil.example/claim") && !strings.Contains(res.stderr, "warning: ignoring") {
			t.Errorf("printed an untrusted claim URL: %s", res.stderr)
		}
		if !strings.Contains(res.stderr, "warning: ignoring .cadenya/temporary.json: claim URL") {
			t.Errorf("no warning: %s", res.stderr)
		}
	})
	t.Run("untrusted claimUrl in REQUIRES_CLAIM", func(t *testing.T) {
		project := isolate(t)
		plantTemporaryFile(t, project, plantedWorkspace("", "http://app.cadenya.com/claim/x"), 0o600)
		var w strings.Builder
		reportTemporaryWorkspaceError(&w, statusError(403, []byte(`{"code":7,"message":"needs claim","details":[{"@type":"type.googleapis.com/google.rpc.ErrorInfo","reason":"REQUIRES_CLAIM"}]}`)))
		if strings.Contains(w.String(), "http://app.cadenya.com") {
			t.Errorf("printed a non-https claim URL: %s", w.String())
		}
	})
	t.Run("loose permissions", func(t *testing.T) {
		if runtime.GOOS == "windows" {
			t.Skip("no Unix permissions")
		}
		project := isolate(t)
		api := newFakeAPI(t)
		trustAPI(t, api.URL, "")
		plantTemporaryFile(t, project, plantedWorkspace(api.URL, "https://app.cadenya.com/claim/x"), 0o644)
		res := run(t, "workspaces", "list")
		if res.code == 0 || api.lastToken() != "" {
			t.Fatalf("a 0644 file was used (exit %d, token %q)", res.code, api.lastToken())
		}
		if !strings.Contains(res.stderr, "its mode is 0644, not 0600") {
			t.Errorf("no warning: %s", res.stderr)
		}
	})
	t.Run("symlinked .cadenya", func(t *testing.T) {
		project := isolate(t)
		elsewhere := t.TempDir()
		api := newFakeAPI(t)
		trustAPI(t, api.URL, "")
		plantTemporaryFile(t, elsewhere, plantedWorkspace(api.URL, "https://app.cadenya.com/claim/x"), 0o600)
		if err := os.Symlink(filepath.Join(elsewhere, ".cadenya"), filepath.Join(project, ".cadenya")); err != nil {
			t.Skip("symlinks unavailable:", err)
		}
		res := run(t, "workspaces", "list")
		if res.code == 0 || api.lastToken() != "" {
			t.Fatalf("a symlinked .cadenya was used (exit %d)", res.code)
		}
		if err := writeTemporaryWorkspace(&temporaryWorkspace{WorkspaceID: "w"}); err == nil {
			t.Errorf("wrote through a symlinked .cadenya")
		}
	})
	t.Run("own 0600 file on a trusted API is used", func(t *testing.T) {
		project := isolate(t)
		api := newFakeAPI(t)
		trustAPI(t, api.URL, "")
		plantTemporaryFile(t, project, plantedWorkspace(api.URL, "https://app.cadenya.dev/claim/x"), 0o600)
		res := run(t, "workspaces", "list")
		if res.code != 0 || api.lastToken() != "sk_planted" {
			t.Errorf("exit %d, token %q: %s", res.code, api.lastToken(), res.stderr)
		}
	})
}

func TestCreateRecordsOnlyTrustedBaseURL(t *testing.T) {
	isolate(t)
	prov := newFakeProvisioning(t, time.Hour)
	_, res := createJSON(t, prov.URL, "--base-url", "http://localhost:8080")
	tw, err := readTemporaryWorkspace()
	if err != nil || tw.BaseURL != "" {
		t.Errorf("recorded an untrusted base URL: %+v, %v", tw, err)
	}
	if !strings.Contains(res.stderr, "pass --base-url on every command") {
		t.Errorf("no note: %s", res.stderr)
	}
}

func TestCreateRejectsUntrustedClaimURLAndZeroExpiry(t *testing.T) {
	for name, body := range map[string]string{
		"claim url": `{"workspaceId":"w","apiKey":"k","claimUrl":"https://evil.example/c","claimToken":"t","expiresAt":"2030-01-01T00:00:00Z"}`,
		"expiry":    `{"workspaceId":"w","apiKey":"k","claimUrl":"https://app.cadenya.com/c","claimToken":"t"}`,
	} {
		t.Run(name, func(t *testing.T) {
			project := isolate(t)
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { _, _ = io.WriteString(w, body) }))
			defer srv.Close()
			res := run(t, "workspaces", "create", "--temporary", "--new-base-url", srv.URL)
			if res.code == 0 {
				t.Fatalf("accepted: %s", res.stdout)
			}
			if _, err := os.Stat(filepath.Join(project, ".cadenya", "temporary.json")); !os.IsNotExist(err) {
				t.Errorf("stored a rejected response")
			}
		})
	}
}

func TestClaimedErrorForAnotherWorkspaceKeepsFile(t *testing.T) {
	project := isolate(t)
	api := newFakeAPI(t)
	api.handler = func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = io.WriteString(w, `{"code":16,"message":"claimed","details":[{"@type":"type.googleapis.com/google.rpc.ErrorInfo","reason":"TEMPORARY_WORKSPACE_CLAIMED","metadata":{"workspace_id":"workspace_OTHER"}}]}`)
	}
	writeLiveTemporary(t, api.URL)
	// An explicit key outranks the file, so the failed request did not use it.
	res := run(t, "--api-key", "sk_other", "--base-url", api.URL, "workspaces", "list")
	if res.code != exitTemporaryWorkspaceClaimed {
		t.Errorf("exit %d", res.code)
	}
	if _, err := os.Stat(filepath.Join(project, ".cadenya", "temporary.json")); err != nil {
		t.Errorf("removed a file the request did not use: %v", err)
	}
	if strings.Contains(res.stderr, "Removed") {
		t.Errorf("claimed a removal that did not happen: %s", res.stderr)
	}
}

func TestClaimedErrorNamingTheFileRemovesIt(t *testing.T) {
	project := isolate(t)
	api := newFakeAPI(t)
	api.handler = claimedHandler // metadata.workspace_id = workspace_T
	writeLiveTemporary(t, api.URL)
	res := run(t, "--api-key", "sk_other", "--base-url", api.URL, "workspaces", "list")
	if res.code != exitTemporaryWorkspaceClaimed || !strings.Contains(res.stderr, "Removed .cadenya/temporary.json") {
		t.Errorf("exit %d: %s", res.code, res.stderr)
	}
	if _, err := os.Stat(filepath.Join(project, ".cadenya", "temporary.json")); !os.IsNotExist(err) {
		t.Errorf("file kept")
	}
}

func TestTemporaryWorkspaceIDIgnoresStoredDefault(t *testing.T) {
	isolate(t)
	api := newFakeAPI(t)
	var paths []string
	api.handler = func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		_, _ = io.WriteString(w, `{"items":[]}`)
	}
	writeLiveTemporary(t, api.URL)
	if err := saveConfigDefaults(map[string]map[string]string{"default": {"workspace-id": "workspace_STORED"}}); err != nil {
		t.Fatal(err)
	}
	res := run(t, "agents", "list", "--display", "json")
	if res.code != 0 || len(paths) == 0 || !strings.Contains(paths[0], "workspace_T") {
		t.Errorf("exit %d, paths %v: %s", res.code, paths, res.stderr)
	}
	t.Setenv("CADENYA_WORKSPACE_ID", "workspace_ENV")
	paths = nil
	res = run(t, "agents", "list", "--display", "json")
	if res.code != 0 || len(paths) == 0 || !strings.Contains(paths[0], "workspace_ENV") {
		t.Errorf("env should override: %v", paths)
	}
}

func TestConfigPlanUsesTemporaryCredential(t *testing.T) {
	project := isolate(t)
	api := newFakeAPI(t)
	var paths []string
	api.handler = func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		_, _ = io.WriteString(w, `{"items":[]}`)
	}
	writeLiveTemporary(t, api.URL)
	if err := os.WriteFile(filepath.Join(project, "cadenya.yaml"), []byte("bundleKey: demo\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := saveConfigDefaults(map[string]map[string]string{"default": {"workspace-id": "workspace_STORED"}}); err != nil {
		t.Fatal(err)
	}
	res := run(t, "config", "plan")
	if res.code != 0 {
		t.Fatalf("exit %d: %s", res.code, res.stderr)
	}
	if api.lastToken() != "sk_temp" {
		t.Errorf("token %q", api.lastToken())
	}
	if len(paths) == 0 || !strings.Contains(paths[0], "workspace_T") {
		t.Errorf("paths %v", paths)
	}
}

func TestReuseComparesBaseURL(t *testing.T) {
	isolate(t)
	prov := newFakeProvisioning(t, 24*time.Hour)
	if err := writeTemporaryWorkspace(&temporaryWorkspace{WorkspaceID: "workspace_STAGING", APIKey: "sk_s", ClaimURL: "https://app.cadenya.dev/claim/s", ClaimToken: "s", ExpiresAt: time.Now().Add(time.Hour), BaseURL: "https://api.cadenya.dev"}); err != nil {
		t.Fatal(err)
	}
	env, res := createJSON(t, prov.URL)
	if env.Reused || env.WorkspaceID == "workspace_STAGING" {
		t.Errorf("reused a staging file for a production create: %+v", env)
	}
	if !strings.Contains(res.stderr, "replacing the file") {
		t.Errorf("did not say so: %s", res.stderr)
	}
	again, _ := createJSON(t, prov.URL)
	if !again.Reused {
		t.Errorf("a matching file was not reused")
	}
}

func TestReuseReappliesGitignore(t *testing.T) {
	project := isolate(t)
	prov := newFakeProvisioning(t, 24*time.Hour)
	createJSON(t, prov.URL)
	if err := os.Remove(filepath.Join(project, ".gitignore")); err != nil {
		t.Fatal(err)
	}
	env, _ := createJSON(t, prov.URL)
	if !env.Reused {
		t.Fatal("not reused")
	}
	if data, err := os.ReadFile(filepath.Join(project, ".gitignore")); err != nil || !strings.Contains(string(data), gitignoreEntry) {
		t.Errorf(".gitignore not restored on reuse: %q %v", data, err)
	}
}

func TestGitignoreCRLFAndNegation(t *testing.T) {
	t.Run("crlf already ignored", func(t *testing.T) {
		project := isolate(t)
		path := filepath.Join(project, ".gitignore")
		_ = os.WriteFile(path, []byte("node_modules\r\n.cadenya/\r\n"), 0o644)
		if note, err := ensureGitignored(); err != nil || note != "" {
			t.Errorf("note %q err %v", note, err)
		}
	})
	t.Run("crlf append keeps line endings", func(t *testing.T) {
		project := isolate(t)
		path := filepath.Join(project, ".gitignore")
		_ = os.WriteFile(path, []byte("node_modules\r\n"), 0o644)
		if _, err := ensureGitignored(); err != nil {
			t.Fatal(err)
		}
		data, _ := os.ReadFile(path)
		if string(data) != "node_modules\r\n.cadenya/temporary.json\r\n" {
			t.Errorf(".gitignore = %q", data)
		}
	})
	t.Run("negation without git", func(t *testing.T) {
		project := isolate(t)
		path := filepath.Join(project, ".gitignore")
		_ = os.WriteFile(path, []byte(".cadenya/\n!.cadenya/temporary.json\n"), 0o644)
		note, err := ensureGitignored()
		if err != nil || note == "" {
			t.Fatalf("a negated entry was treated as ignored: %q %v", note, err)
		}
		data, _ := os.ReadFile(path)
		if !strings.HasSuffix(string(data), "!.cadenya/temporary.json\n.cadenya/temporary.json\n") {
			t.Errorf(".gitignore = %q", data)
		}
	})
	t.Run("negation with git", func(t *testing.T) {
		if _, err := exec.LookPath("git"); err != nil {
			t.Skip("git unavailable")
		}
		project := isolate(t)
		if out, err := exec.Command("git", "-C", project, "init", "-q").CombinedOutput(); err != nil {
			t.Skip("git init failed:", string(out))
		}
		path := filepath.Join(project, ".gitignore")
		_ = os.WriteFile(path, []byte(".cadenya/*\n!.cadenya/temporary.json\n"), 0o644)
		if _, err := ensureGitignored(); err != nil {
			t.Fatal(err)
		}
		if ignored, known := gitIgnores(project); !known || !ignored {
			t.Errorf("git does not ignore the file after ensureGitignored (known=%v)", known)
		}
	})
	t.Run("nested negation only git sees", func(t *testing.T) {
		if _, err := exec.LookPath("git"); err != nil {
			t.Skip("git unavailable")
		}
		project := isolate(t)
		if out, err := exec.Command("git", "-C", project, "init", "-q").CombinedOutput(); err != nil {
			t.Skip("git init failed:", string(out))
		}
		_ = os.WriteFile(filepath.Join(project, ".gitignore"), []byte(".cadenya/temporary.json\n"), 0o644)
		_ = os.Mkdir(filepath.Join(project, ".cadenya"), 0o700)
		_ = os.WriteFile(filepath.Join(project, ".cadenya", ".gitignore"), []byte("!temporary.json\n"), 0o644)
		if _, err := ensureGitignored(); err == nil {
			t.Errorf("reported the file as ignored although git does not ignore it")
		}
	})
	t.Run("symlinked .gitignore", func(t *testing.T) {
		project := isolate(t)
		target := filepath.Join(t.TempDir(), "elsewhere")
		_ = os.WriteFile(target, []byte("x\n"), 0o644)
		if err := os.Symlink(target, filepath.Join(project, ".gitignore")); err != nil {
			t.Skip("symlinks unavailable")
		}
		if _, err := ensureGitignored(); err == nil {
			t.Errorf("edited a symlinked .gitignore")
		}
		if data, _ := os.ReadFile(target); string(data) != "x\n" {
			t.Errorf("wrote through the symlink: %q", data)
		}
	})
}

func TestHumanModeKeepsSecretsOffStderr(t *testing.T) {
	isolate(t)
	prov := newFakeProvisioning(t, 24*time.Hour)
	res := run(t, "workspaces", "create", "--temporary", "--new-base-url", prov.URL)
	if res.code != 0 {
		t.Fatalf("exit %d: %s", res.code, res.stderr)
	}
	for _, secret := range []string{"sk_temp_workspace_A"} {
		if strings.Contains(res.stderr, secret) || strings.Contains(res.stdout, secret) {
			t.Errorf("printed the API key: %s %s", res.stdout, res.stderr)
		}
	}
	// The claim token is part of the claim URL by design; it must not
	// appear anywhere else.
	if strings.Count(res.stderr, "tok_workspace_A") != 1 || strings.Contains(res.stdout, "tok_workspace_A") {
		t.Errorf("claim token outside the claim URL: %s %s", res.stdout, res.stderr)
	}
}

func TestJSONSaveFailureStillEmitsEnvelope(t *testing.T) {
	project := isolate(t)
	prov := newFakeProvisioning(t, time.Hour)
	if err := os.WriteFile(filepath.Join(project, ".cadenya"), []byte("not a dir"), 0o644); err != nil {
		t.Fatal(err)
	}
	res := run(t, "workspaces", "create", "--temporary", "--new-base-url", prov.URL, "--display", "json")
	if res.code == 0 {
		t.Fatal("exit 0 although the file was not saved")
	}
	var env temporaryEnvelope
	if err := json.Unmarshal([]byte(res.stdout), &env); err != nil || env.APIKey == "" {
		t.Errorf("no envelope on stdout: %q", res.stdout)
	}
}

func TestEdgeErrorBodyMessage(t *testing.T) {
	isolate(t)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		_, _ = io.WriteString(w, `{"error":{"code":"NO_ORIGIN","message":"no cluster is accepting new workspaces"}}`)
	}))
	defer srv.Close()
	res := run(t, "workspaces", "create", "--temporary", "--new-base-url", srv.URL)
	if res.code == 0 || !strings.Contains(res.stderr, "no cluster is accepting new workspaces") {
		t.Errorf("exit %d: %s", res.code, res.stderr)
	}
}

func TestRetryAfterHTTPDate(t *testing.T) {
	at := time.Now().Add(2 * time.Minute).UTC().Format(http.TimeFormat)
	d := parseRetryAfter(at)
	if d < 100*time.Second || d > 120*time.Second {
		t.Errorf("parseRetryAfter(%q) = %s", at, d)
	}
	if parseRetryAfter(time.Now().Add(-time.Minute).UTC().Format(http.TimeFormat)) != 0 {
		t.Errorf("a past date should mean no wait")
	}
}

func TestPollDelayBounds(t *testing.T) {
	saved := []time.Duration{minPollInterval, maxPollInterval}
	minPollInterval, maxPollInterval = time.Second, 60*time.Second
	defer func() { minPollInterval, maxPollInterval = saved[0], saved[1] }()
	now := time.Now()
	far := now.Add(time.Hour)
	cases := []struct {
		raw     string
		expires time.Time
		want    time.Duration
	}{
		{"0s", far, time.Second},
		{"0.001s", far, time.Second},
		{"", far, 5 * time.Second},
		{"garbage", far, 5 * time.Second},
		{"86400s", far, 60 * time.Second},
		{"30s", now.Add(10 * time.Second), 10 * time.Second},
		{"30s", now.Add(-time.Minute), time.Second},
	}
	for _, c := range cases {
		if got := pollDelay(c.raw, c.expires, now); got != c.want {
			t.Errorf("pollDelay(%q, %s) = %s, want %s", c.raw, c.expires.Sub(now), got, c.want)
		}
	}
}

func TestClaimStatusWaitRetriesTransientErrors(t *testing.T) {
	isolate(t)
	api := newFakeAPI(t)
	var calls atomic.Int32
	api.handler = func(w http.ResponseWriter, r *http.Request) {
		switch calls.Add(1) {
		case 1:
			w.WriteHeader(http.StatusBadGateway)
			_, _ = io.WriteString(w, `{"code":14,"message":"upstream unavailable"}`)
		case 2:
			w.WriteHeader(http.StatusServiceUnavailable)
		default:
			_ = json.NewEncoder(w).Encode(map[string]any{"workspaceId": "workspace_T", "state": stateClaimed, "pollInterval": "1s"})
		}
	}
	writeLiveTemporary(t, api.URL)
	res := run(t, "workspaces", "claim-status", "--wait", "--display", "json")
	if res.code != 0 || !strings.Contains(res.stdout, stateClaimed) || calls.Load() != 3 {
		t.Errorf("exit %d after %d calls: %s %s", res.code, calls.Load(), res.stdout, res.stderr)
	}
	if !strings.Contains(res.stderr, "retrying") {
		t.Errorf("retries not reported: %s", res.stderr)
	}
}

func TestClaimStatusWaitStopsOn4xx(t *testing.T) {
	isolate(t)
	api := newFakeAPI(t)
	var calls atomic.Int32
	api.handler = func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		w.WriteHeader(http.StatusForbidden)
		_, _ = io.WriteString(w, `{"code":7,"message":"forbidden"}`)
	}
	writeLiveTemporary(t, api.URL)
	res := run(t, "workspaces", "claim-status", "--wait")
	if res.code == 0 || calls.Load() != 1 {
		t.Errorf("exit %d after %d calls", res.code, calls.Load())
	}
}

func TestClaimStatusExpiredExitCode(t *testing.T) {
	isolate(t)
	api := newFakeAPI(t)
	api.states = []string{stateExpired}
	writeLiveTemporary(t, api.URL)
	res := run(t, "workspaces", "claim-status", "--display", "json")
	if res.code != exitTemporaryWorkspaceExpired || !strings.Contains(res.stdout, stateExpired) {
		t.Errorf("exit %d: %s", res.code, res.stdout)
	}
}

func TestClaimStatusWaitCancellation(t *testing.T) {
	isolate(t)
	api := newFakeAPI(t)
	api.states = []string{stateUnclaimed}
	writeLiveTemporary(t, api.URL)
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	start := time.Now()
	_, err := runCtx(t, ctx, "workspaces", "claim-status", "--wait")
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("err = %v, want the context's error", err)
	}
	if time.Since(start) > 2*time.Second {
		t.Errorf("cancellation took %s", time.Since(start))
	}
}

func TestAuthLoginSkipsClaimAcrossEnvironments(t *testing.T) {
	prev := authDefaultPollInterval
	authDefaultPollInterval = time.Millisecond
	defer func() { authDefaultPollInterval = prev }()
	isolate(t)
	auth, sent := fakeAuth(t)
	t.Setenv("CADENYA_AUTH_BASE_URL", auth.URL)
	if err := writeTemporaryWorkspace(&temporaryWorkspace{WorkspaceID: "workspace_S", APIKey: "sk_s", ClaimURL: "https://app.cadenya.dev/claim/s", ClaimToken: "tok_S", ExpiresAt: time.Now().Add(time.Hour), BaseURL: "https://api.cadenya.dev"}); err != nil {
		t.Fatal(err)
	}
	res := run(t, "auth", "login", "--no-browser")
	if res.code != 0 {
		t.Fatalf("exit %d: %s", res.code, res.stderr)
	}
	if _, ok := (*sent)["claim_token"]; ok {
		t.Errorf("sent a staging claim token to another auth host")
	}
	if !strings.Contains(res.stdout, "Not claiming the temporary workspace workspace_S") {
		t.Errorf("did not say why: %s", res.stdout)
	}
}

func TestCodexDetectionIsNarrow(t *testing.T) {
	isolate(t)
	t.Setenv("CODEX_HOME", "/tmp/x")
	if codingAgentDetected() {
		t.Errorf("CODEX_HOME counted as a coding agent")
	}
	t.Setenv("CODEX_SANDBOX", "seatbelt")
	if !codingAgentDetected() {
		t.Errorf("CODEX_SANDBOX not detected")
	}
}
