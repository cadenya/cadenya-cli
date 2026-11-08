package main

// Temporary workspaces: a workspace created without an account, held by a
// per-project credential file until a person claims it. The create call goes
// to the provisioning host (new.cadenya.com); everything else, including
// claim status, goes to the normal API with the temporary key.

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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

// temporaryWorkspace is the content of .cadenya/temporary.json.
type temporaryWorkspace struct {
	WorkspaceID string    `json:"workspaceId"`
	AccountID   string    `json:"accountId,omitempty"`
	APIKey      string    `json:"apiKey"`
	ClaimURL    string    `json:"claimUrl"`
	ClaimToken  string    `json:"claimToken"`
	ExpiresAt   time.Time `json:"expiresAt"`
	NextStep    string    `json:"nextStep,omitempty"`
	// BaseURL is the API base URL the workspace lives behind, recorded when
	// it is not the production default so a staging workspace keeps working
	// without repeating --base-url.
	BaseURL string `json:"baseUrl,omitempty"`
}

func (t *temporaryWorkspace) live(now time.Time) bool {
	return t.APIKey != "" && now.Before(t.ExpiresAt)
}

func temporaryFilePath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(wd, temporaryDir, temporaryFileName), nil
}

// readTemporaryWorkspace loads this directory's temporary workspace file,
// live or not. A missing file is (nil, nil).
func readTemporaryWorkspace() (*temporaryWorkspace, error) {
	path, err := temporaryFilePath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	var tw temporaryWorkspace
	if err := json.Unmarshal(data, &tw); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", filepath.Join(temporaryDir, temporaryFileName), err)
	}
	return &tw, nil
}

// liveTemporaryWorkspace returns this directory's unexpired temporary
// workspace, or nil. Best effort, like storedProfile: an unreadable file
// means "none" in the middle of an unrelated command.
func liveTemporaryWorkspace() *temporaryWorkspace {
	tw, err := readTemporaryWorkspace()
	if err != nil || tw == nil || !tw.live(time.Now()) {
		return nil
	}
	return tw
}

func removeTemporaryWorkspace() error {
	path, err := temporaryFilePath()
	if err != nil {
		return err
	}
	if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}

// writeTemporaryWorkspace stores the file at 0600 from the first byte: a
// private temp file renamed over the destination, as for the credentials
// file.
func writeTemporaryWorkspace(tw *temporaryWorkspace) error {
	path, err := temporaryFilePath()
	if err != nil {
		return err
	}
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(tw, "", "  ")
	if err != nil {
		return err
	}
	tmp, err := os.CreateTemp(dir, ".temporary-*")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	if err := tmp.Chmod(0o600); err != nil {
		tmp.Close()
		return err
	}
	if _, err := tmp.Write(append(data, '\n')); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmp.Name(), path)
}

// gitignoreEntry is what ensureGitignored adds. Only the credential file:
// configuration bundles keep committed resources in .cadenya/.
const gitignoreEntry = ".cadenya/temporary.json"

// ensureGitignored adds .cadenya/temporary.json to the working directory's
// .gitignore, creating the file if needed. It returns a sentence describing
// the change, or "" when the file was already ignored.
func ensureGitignored() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := filepath.Join(wd, ".gitignore")
	data, err := os.ReadFile(path)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return "", err
	}
	exists := err == nil
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		switch strings.TrimSpace(scanner.Text()) {
		case ".cadenya", ".cadenya/", "/.cadenya", "/.cadenya/", ".cadenya/*", "/.cadenya/*", gitignoreEntry, "/" + gitignoreEntry:
			return "", nil
		}
	}
	var buf bytes.Buffer
	buf.Write(data)
	if len(data) > 0 && data[len(data)-1] != '\n' {
		buf.WriteByte('\n')
	}
	buf.WriteString(gitignoreEntry + "\n")
	if err := os.WriteFile(path, buf.Bytes(), 0o644); err != nil {
		return "", err
	}
	if exists {
		return "Added " + gitignoreEntry + " to .gitignore.", nil
	}
	return "Created .gitignore with " + gitignoreEntry + ".", nil
}

// inCI reports a CI environment the way most CLIs do: CI set to anything
// other than an explicit false.
func inCI() bool {
	v := strings.TrimSpace(os.Getenv("CI"))
	return v != "" && v != "0" && !strings.EqualFold(v, "false")
}

// codingAgentEnv names the variables coding agents set in the shells they
// run. Detection only tags the create request; it never changes behavior.
var codingAgentEnv = []string{"CLAUDECODE", "CURSOR_AGENT", "GEMINI_CLI", "AI_AGENT", "OPENCODE", "AMP_AGENT"}

var codingAgentEnvPrefixes = []string{"CODEX_"}

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

// apiBaseURLForNew is the API base URL to record for a workspace created on
// the given provisioning host: the explicit --base-url, the staging API for
// the staging host, and "" (the production default) otherwise.
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
		var edge struct {
			Error struct {
				Message string `json:"message"`
			} `json:"error"`
		}
		_ = json.Unmarshal(data, &edge)
		message := edge.Error.Message
		if message == "" {
			message = statusError(resp.StatusCode, data).Message
		}
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
	if out.WorkspaceID == "" || out.APIKey == "" || out.ClaimURL == "" {
		return nil, fmt.Errorf("%s answered without a workspace, API key or claim URL", base)
	}
	return &out, nil
}

// statusError decodes a google.rpc.Status body into the SDK's error type,
// so reporting treats hand-written calls like SDK calls.
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
	var apiErr *sdk.APIError
	if !errors.As(err, &apiErr) {
		return "", nil
	}
	for _, detail := range apiErr.Details {
		typ, _ := detail["@type"].(string)
		if !strings.HasSuffix(typ, "google.rpc.ErrorInfo") {
			continue
		}
		if reason, _ := detail["reason"].(string); reason != "" {
			return reason, apiErr
		}
	}
	return "", apiErr
}

const claimedGuidance = "The temporary workspace was claimed, so its API key no longer works. Removed .cadenya/temporary.json. Run `cadenya auth login` in this directory, or create an API key for the workspace in the dashboard."

// reportTemporaryWorkspaceError handles the error reasons tied to temporary
// workspaces. handled=false leaves the error to the ordinary reporter.
func reportTemporaryWorkspaceError(w io.Writer, err error) (int, bool) {
	reason, apiErr := errorInfoReason(err)
	switch reason {
	case reasonTemporaryWorkspaceClaimed:
		_ = removeTemporaryWorkspace()
		if apiErr.Message != "" {
			fmt.Fprintln(w, "error:", apiErr.Message)
		}
		fmt.Fprintln(w, claimedGuidance)
		return exitTemporaryWorkspaceClaimed, true
	case reasonRequiresClaim:
		fmt.Fprintln(w, "error:", err)
		if tw, readErr := readTemporaryWorkspace(); readErr == nil && tw != nil && tw.ClaimURL != "" {
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
