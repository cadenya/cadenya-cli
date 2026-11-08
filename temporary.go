package main

// Temporary workspaces: a workspace created without an account, held by a
// per-project credential file until a person claims it. The create call goes
// to the provisioning host (new.cadenya.com); everything else, including
// claim status, goes to the normal API with the temporary key.

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	cli "github.com/urfave/cli/v3"

	sdk "go.cadenya.com/cadenya-go"
)

const (
	temporaryDir      = ".cadenya"
	temporaryFileName = "temporary.json"

	defaultAPIBaseURL = "https://api.cadenya.com"
	defaultNewBaseURL = "https://new.cadenya.com"
	stagingAPIHost    = "api.cadenya.dev"
	stagingNewBaseURL = "https://new.cadenya.dev"

	// exitTemporaryWorkspaceClaimed is the stable exit code for a request
	// made with the key of a temporary workspace that has been claimed.
	exitTemporaryWorkspaceClaimed = 3

	// exitTemporaryWorkspaceExpired is the exit code of claim-status when
	// the workspace expired unclaimed.
	exitTemporaryWorkspaceExpired = 4

	reasonTemporaryWorkspaceClaimed = "TEMPORARY_WORKSPACE_CLAIMED"
	reasonRequiresClaim             = "REQUIRES_CLAIM"

	temporaryCreateHint = "No Cadenya account? Run: cadenya workspaces create --temporary"
)

// temporaryInstructionsURL serves the agent guide printed by
// `workspaces create --instructions`. Fetched at run time so the guide can
// change without a CLI release.
var temporaryInstructionsURL = "https://docs.cadenya.com/llms/start-without-account.txt"

// temporaryHTTPClient carries every hand-written temporary-workspace call.
var temporaryHTTPClient = &http.Client{Timeout: 30 * time.Second}

// inCI reports a CI environment the way most CLIs do: CI set to anything
// other than an explicit false.
func inCI() bool {
	v := strings.TrimSpace(os.Getenv("CI"))
	return v != "" && v != "0" && !strings.EqualFold(v, "false")
}

// codingAgentEnv names the variables coding agents set in the shells they
// run. Detection only tags the create request; it never changes behavior.
var codingAgentEnv = []string{"CLAUDECODE", "CURSOR_AGENT", "GEMINI_CLI", "AI_AGENT", "OPENCODE", "AMP_AGENT"}

// codingAgentEnvPrefixes matches Codex sandbox variables (CODEX_SANDBOX,
// CODEX_SANDBOX_NETWORK_DISABLED) without catching unrelated CODEX_ names.
var codingAgentEnvPrefixes = []string{"CODEX_SANDBOX"}

func codingAgentDetected() bool {
	for _, name := range codingAgentEnv {
		if os.Getenv(name) != "" {
			return true
		}
	}
	for _, kv := range os.Environ() {
		for _, prefix := range codingAgentEnvPrefixes {
			if strings.HasPrefix(kv, prefix) {
				return true
			}
		}
	}
	return false
}

func temporaryWorkspaceSource() string {
	if codingAgentDetected() {
		return "TEMPORARY_WORKSPACE_SOURCE_AGENT"
	}
	return "TEMPORARY_WORKSPACE_SOURCE_CLI"
}

// newBaseURL resolves the provisioning host: --new-base-url (or
// $CADENYA_NEW_BASE_URL), then staging when --base-url points at the staging
// API, then production.
func newBaseURL(cmd *cli.Command) string {
	if v := strings.TrimSpace(cmd.String("new-base-url")); v != "" {
		return strings.TrimRight(v, "/")
	}
	if cmd.Root().IsSet("base-url") {
		if u, err := url.Parse(cmd.Root().String("base-url")); err == nil && u.Hostname() == stagingAPIHost {
			return stagingNewBaseURL
		}
	}
	return defaultNewBaseURL
}

// apiBaseURLForNew is the API base URL a workspace created on the given
// provisioning host lives behind: the explicit --base-url, the staging API
// for the staging host, and "" (the production default) otherwise.
func apiBaseURLForNew(cmd *cli.Command, newBase string) string {
	if cmd.Root().IsSet("base-url") {
		return strings.TrimRight(cmd.Root().String("base-url"), "/")
	}
	if newBase == stagingNewBaseURL {
		return "https://" + stagingAPIHost
	}
	return ""
}

// temporaryAPIBaseURL is the API base URL for calls made with the temporary
// key: --base-url (or $CADENYA_BASE_URL), then the URL recorded at creation.
func temporaryAPIBaseURL(cmd *cli.Command, tw *temporaryWorkspace) string {
	if cmd.Root().IsSet("base-url") {
		return strings.TrimRight(cmd.Root().String("base-url"), "/")
	}
	if tw != nil && tw.BaseURL != "" {
		return tw.BaseURL
	}
	return defaultAPIBaseURL
}

type createTemporaryRequest struct {
	Name        string `json:"name,omitempty"`
	ExpireAfter string `json:"expireAfter,omitempty"`
	Source      string `json:"source"`
}

type createTemporaryResponse struct {
	WorkspaceID string    `json:"workspaceId"`
	AccountID   string    `json:"accountId"`
	APIKey      string    `json:"apiKey"`
	ClaimURL    string    `json:"claimUrl"`
	ClaimToken  string    `json:"claimToken"`
	ExpiresAt   time.Time `json:"expiresAt"`
	NextStep    string    `json:"nextStep"`
}

// rateLimitedError is the provisioning edge's 429.
type rateLimitedError struct {
	retryAfter time.Duration
	message    string
}

func (e *rateLimitedError) Error() string {
	msg := "too many temporary workspaces were created from this network"
	if e.retryAfter > 0 {
		msg += fmt.Sprintf("; try again in %s", humanDuration(e.retryAfter))
	} else {
		msg += "; try again later"
	}
	if e.message != "" {
		msg += " (" + e.message + ")"
	}
	return msg
}

var errTemporaryUnavailable = errors.New("temporary workspaces are unavailable right now; try again later, or sign up at https://app.cadenya.com and run `cadenya auth login`")

func createTemporaryWorkspace(ctx context.Context, base string, body createTemporaryRequest) (*createTemporaryResponse, error) {
	raw, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, base+"/v1/temporary_workspaces", bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "cadenya-cli/"+version)
	resp, err := temporaryHTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("creating a temporary workspace: %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, err
	}
	switch {
	case resp.StatusCode == http.StatusTooManyRequests:
		message := statusError(resp.StatusCode, data).Message
		return nil, &rateLimitedError{retryAfter: parseRetryAfter(resp.Header.Get("Retry-After")), message: message}
	case resp.StatusCode == http.StatusServiceUnavailable:
		return nil, errTemporaryUnavailable
	case resp.StatusCode < 200 || resp.StatusCode >= 300:
		return nil, statusError(resp.StatusCode, data)
	}
	var out createTemporaryResponse
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("%s answered HTTP %d with an unreadable body", base, resp.StatusCode)
	}
	if out.WorkspaceID == "" || out.APIKey == "" || out.ClaimURL == "" || out.ExpiresAt.IsZero() {
		return nil, fmt.Errorf("%s answered without a workspace, API key, claim URL or expiry", base)
	}
	if err := checkClaimURL(out.ClaimURL); err != nil {
		return nil, fmt.Errorf("%s answered with an untrusted claim URL: %w", base, err)
	}
	return &out, nil
}

// statusError decodes an error body into the SDK's error type, so reporting
// treats hand-written calls like SDK calls. It reads google.rpc.Status, and
// the edge worker's {"error":{"code","message"}} when that has no message.
func statusError(code int, data []byte) *sdk.APIError {
	apiErr := &sdk.APIError{StatusCode: code}
	var status struct {
		Code    int              `json:"code"`
		Message string           `json:"message"`
		Details []map[string]any `json:"details"`
	}
	if json.Unmarshal(data, &status) == nil {
		apiErr.Code = status.Code
		apiErr.Message = status.Message
		apiErr.Details = status.Details
	}
	if apiErr.Message == "" {
		var edge struct {
			Error struct {
				Message string `json:"message"`
			} `json:"error"`
		}
		if json.Unmarshal(data, &edge) == nil {
			apiErr.Message = edge.Error.Message
		}
	}
	return apiErr
}

func parseRetryAfter(v string) time.Duration {
	v = strings.TrimSpace(v)
	if v == "" {
		return 0
	}
	if secs, err := strconv.Atoi(v); err == nil && secs >= 0 {
		return time.Duration(secs) * time.Second
	}
	if at, err := http.ParseTime(v); err == nil {
		if d := time.Until(at); d > 0 {
			return d.Round(time.Second)
		}
	}
	return 0
}

// temporaryStatus is the API's TemporaryWorkspace message.
type temporaryStatus struct {
	WorkspaceID  string     `json:"workspaceId"`
	State        string     `json:"state"`
	ExpiresAt    *time.Time `json:"expiresAt,omitempty"`
	ClaimedAt    *time.Time `json:"claimedAt,omitempty"`
	PollInterval string     `json:"pollInterval,omitempty"`
}

const (
	stateUnclaimed = "STATE_UNCLAIMED"
	statePending   = "STATE_PENDING"
	stateClaimed   = "STATE_CLAIMED"
	stateExpired   = "STATE_EXPIRED"
)

func fetchTemporaryStatus(ctx context.Context, base string, tw *temporaryWorkspace) (*temporaryStatus, error) {
	endpoint := base + "/v1/workspaces/" + url.PathEscape(tw.WorkspaceID) + "/temporary"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+tw.APIKey)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "cadenya-cli/"+version)
	resp, err := temporaryHTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("checking claim status: %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, statusError(resp.StatusCode, data)
	}
	var out temporaryStatus
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("%s answered HTTP %d with an unreadable body", base, resp.StatusCode)
	}
	return &out, nil
}

// errorInfoReason returns the google.rpc.ErrorInfo reason of an API error.
func errorInfoReason(err error) (string, *sdk.APIError) {
	reason, _, apiErr := errorInfo(err)
	return reason, apiErr
}

// errorInfo returns the reason and metadata of an API error's
// google.rpc.ErrorInfo detail.
func errorInfo(err error) (string, map[string]any, *sdk.APIError) {
	var apiErr *sdk.APIError
	if !errors.As(err, &apiErr) {
		return "", nil, nil
	}
	for _, detail := range apiErr.Details {
		typ, _ := detail["@type"].(string)
		if !strings.HasSuffix(typ, "google.rpc.ErrorInfo") {
			continue
		}
		if reason, _ := detail["reason"].(string); reason != "" {
			metadata, _ := detail["metadata"].(map[string]any)
			return reason, metadata, apiErr
		}
	}
	return "", nil, apiErr
}

const claimedGuidance = "The temporary workspace was claimed, so its API key no longer works. Run `cadenya auth login` in this directory, or create an API key for the workspace in the dashboard."

// temporaryCredentialUsed is the workspace id of the temporary file that
// newClient authenticated with in this run, or "".
var temporaryCredentialUsed string

// removeClaimedTemporaryWorkspace removes this directory's file after a
// TEMPORARY_WORKSPACE_CLAIMED answer, but only when the answer is about that
// file: the request used it, or the error names its workspace.
func removeClaimedTemporaryWorkspace(metadata map[string]any) bool {
	tw, err := readTemporaryWorkspace()
	if err != nil || tw == nil {
		return false
	}
	named, _ := metadata["workspace_id"].(string)
	if named != tw.WorkspaceID && temporaryCredentialUsed != tw.WorkspaceID {
		return false
	}
	removed, err := removeTemporaryWorkspace()
	return err == nil && removed
}

// reportTemporaryWorkspaceError handles the error reasons tied to temporary
// workspaces. handled=false leaves the error to the ordinary reporter.
func reportTemporaryWorkspaceError(w io.Writer, err error) (int, bool) {
	reason, metadata, apiErr := errorInfo(err)
	switch reason {
	case reasonTemporaryWorkspaceClaimed:
		removed := removeClaimedTemporaryWorkspace(metadata)
		if apiErr.Message != "" {
			fmt.Fprintln(w, "error:", apiErr.Message)
		}
		fmt.Fprintln(w, claimedGuidance)
		if removed {
			fmt.Fprintln(w, "Removed .cadenya/temporary.json.")
		}
		return exitTemporaryWorkspaceClaimed, true
	case reasonRequiresClaim:
		fmt.Fprintln(w, "error:", err)
		if tw, readErr := readTemporaryWorkspace(); readErr == nil && tw != nil {
			fmt.Fprintf(w, "This needs a claimed workspace. Claim it to unlock this (don't share this link): %s\n", tw.ClaimURL)
		} else {
			fmt.Fprintln(w, "This needs a claimed workspace. Open the claim link you were given when it was created.")
		}
		return 1, true
	}
	return 0, false
}

// humanDuration renders a coarse remaining time: "24 hours", "45 minutes".
func humanDuration(d time.Duration) string {
	plural := func(n int64, unit string) string {
		if n == 1 {
			return fmt.Sprintf("1 %s", unit)
		}
		return fmt.Sprintf("%d %ss", n, unit)
	}
	switch {
	case d >= time.Hour:
		return plural(int64((d+30*time.Minute)/time.Hour), "hour")
	case d >= time.Minute:
		return plural(int64((d+30*time.Second)/time.Minute), "minute")
	case d > 0:
		return plural(int64((d+time.Second-1)/time.Second), "second")
	}
	return "0 seconds"
}

func claimNotice(tw *temporaryWorkspace, now time.Time) string {
	return fmt.Sprintf("This workspace expires in %s (at %s). Claim it to keep it (don't share this link): %s",
		humanDuration(tw.ExpiresAt.Sub(now)), tw.ExpiresAt.Local().Format("Mon Jan 2 15:04 MST"), tw.ClaimURL)
}
