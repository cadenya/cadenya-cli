package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

// fakeAuthToken is fakeAuth with a caller-chosen token endpoint answer.
func fakeAuthToken(t *testing.T, status int, body map[string]any) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/cli/oauth/device_authorization":
			_ = json.NewEncoder(w).Encode(map[string]any{"device_code": "dc", "user_code": "ABCD", "verification_uri": "https://example.test/approve", "interval": 0, "expires_in": 60})
		case "/api/cli/oauth/token":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(status)
			_ = json.NewEncoder(w).Encode(body)
		}
	}))
	t.Cleanup(srv.Close)
	return srv
}

func TestAuthLoginClaimRefused(t *testing.T) {
	prevPoll := authDefaultPollInterval
	authDefaultPollInterval = time.Millisecond
	defer func() { authDefaultPollInterval = prevPoll }()

	cases := []struct {
		reason string
		want   []string
		keep   bool
	}{
		{"ACCOUNT_ALREADY_EXISTS", []string{"already have a Cadenya account", "create a workspace", "expires"}, true},
		{"TEMPORARY_WORKSPACE_EXPIRED", []string{"expired", "cannot be recovered"}, true},
		{"TEMPORARY_WORKSPACE_CLAIMED", []string{"already claimed"}, true},
		{"CLAIM_INVALID", []string{"claim token", "not valid"}, true},
		{"SOMETHING_NEW", []string{"SOMETHING_NEW"}, true},
	}
	for _, tc := range cases {
		t.Run(tc.reason, func(t *testing.T) {
			project := isolate(t)
			auth := fakeAuthToken(t, http.StatusBadRequest, map[string]any{"error": "access_denied", "claim_result": tc.reason})
			t.Setenv("CADENYA_AUTH_BASE_URL", auth.URL)
			api := newFakeAPI(t)
			api.states = []string{stateUnclaimed}
			trustAPI(t, api.URL, auth.URL)
			writeLiveTemporary(t, api.URL)
			res := run(t, "auth", "login", "--no-browser")
			if res.code == 0 {
				t.Fatalf("a refused claim must fail the login: %s", res.stdout)
			}
			if strings.Contains(res.stderr, "login was denied in the browser") && !strings.Contains(res.stderr, tc.want[0]) {
				t.Errorf("generic denial left as is: %s", res.stderr)
			}
			for _, w := range tc.want {
				if !strings.Contains(res.stderr, w) {
					t.Errorf("stderr missing %q: %s", w, res.stderr)
				}
			}
			if _, err := os.Stat(filepath.Join(project, ".cadenya", "temporary.json")); err != nil {
				t.Errorf("temporary.json must stay after a refused claim: %v", err)
			}
			if _, err := os.Stat(filepath.Join(os.Getenv("HOME"), ".cadenya", "credentials")); !os.IsNotExist(err) {
				t.Errorf("no credential may be stored after a denial")
			}
		})
	}
}

func TestAuthLoginDeniedWithoutClaimResultIsGeneric(t *testing.T) {
	prevPoll := authDefaultPollInterval
	authDefaultPollInterval = time.Millisecond
	defer func() { authDefaultPollInterval = prevPoll }()
	isolate(t)
	auth := fakeAuthToken(t, http.StatusBadRequest, map[string]any{"error": "access_denied"})
	t.Setenv("CADENYA_AUTH_BASE_URL", auth.URL)
	res := run(t, "auth", "login", "--no-browser")
	if res.code == 0 || !strings.Contains(res.stderr, "login was denied in the browser") {
		t.Errorf("exit %d: %s", res.code, res.stderr)
	}
}

func TestAuthLoginClaimResultClaimedOnSuccess(t *testing.T) {
	prevPoll := authDefaultPollInterval
	authDefaultPollInterval = time.Millisecond
	defer func() { authDefaultPollInterval = prevPoll }()
	project := isolate(t)
	auth := fakeAuthToken(t, http.StatusOK, map[string]any{"access_token": "sk_login", "token_type": "bearer", "claim_result": "claimed", "workspaces": []map[string]string{{"id": "workspace_T", "name": "Temp"}}})
	t.Setenv("CADENYA_AUTH_BASE_URL", auth.URL)
	api := newFakeAPI(t)
	// The API would say the claim is still unclaimed; the token response is
	// authoritative, so the login must not depend on it.
	api.states = []string{stateUnclaimed}
	trustAPI(t, api.URL, auth.URL)
	writeLiveTemporary(t, api.URL)
	res := run(t, "auth", "login", "--no-browser")
	if res.code != 0 {
		t.Fatalf("exit %d: %s", res.code, res.stderr)
	}
	if !strings.Contains(res.stdout, "Claimed temporary workspace workspace_T") || !strings.Contains(res.stdout, "Removed .cadenya/temporary.json") {
		t.Errorf("stdout: %s", res.stdout)
	}
	if _, err := os.Stat(filepath.Join(project, ".cadenya", "temporary.json")); !os.IsNotExist(err) {
		t.Errorf("temporary.json kept after a claim")
	}
}
