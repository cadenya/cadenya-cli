package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	cli "github.com/urfave/cli/v3"
)

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
	// it is a trusted non-production API so a staging workspace keeps
	// working without repeating --base-url.
	BaseURL string `json:"baseUrl,omitempty"`
}

func (t *temporaryWorkspace) live(now time.Time) bool {
	return t.APIKey != "" && now.Before(t.ExpiresAt)
}

// effectiveBaseURL is the API the workspace lives behind.
func (t *temporaryWorkspace) effectiveBaseURL() string {
	if t.BaseURL != "" {
		return t.BaseURL
	}
	return defaultAPIBaseURL
}

// The file is read from whatever directory the CLI runs in, which may be a
// checkout someone else controls. Everything in it that could send requests
// or a person somewhere is checked against these lists. Tests extend them.
var (
	trustedAPIBaseURLs = []string{defaultAPIBaseURL, "https://" + stagingAPIHost}
	trustedClaimHosts  = []string{"app.cadenya.com", "app.cadenya.dev"}
	// claimAuthBaseURLs pairs each API with the auth host whose login may
	// carry its workspaces' claim tokens.
	claimAuthBaseURLs = map[string]string{
		defaultAPIBaseURL:           "https://app.cadenya.com",
		"https://" + stagingAPIHost: "https://app.cadenya.dev",
	}
)

func normalizeBaseURL(raw string) string { return strings.TrimRight(strings.TrimSpace(raw), "/") }

func trustedAPIBaseURL(raw string) bool {
	raw = normalizeBaseURL(raw)
	for _, allowed := range trustedAPIBaseURLs {
		if raw == normalizeBaseURL(allowed) {
			return true
		}
	}
	return false
}

func checkClaimURL(raw string) error {
	u, err := url.Parse(raw)
	if err != nil {
		return fmt.Errorf("claim URL %q does not parse", raw)
	}
	if u.Scheme == "https" && u.User == nil {
		for _, host := range trustedClaimHosts {
			if u.Host == host {
				return nil
			}
		}
	}
	return fmt.Errorf("claim URL %q is not an https link on %s", raw, strings.Join(trustedClaimHosts, " or "))
}

// untrustedTemporaryFileError explains why .cadenya/temporary.json was
// ignored.
type untrustedTemporaryFileError struct{ reason string }

func (e *untrustedTemporaryFileError) Error() string {
	return "ignoring .cadenya/temporary.json: " + e.reason
}

func untrusted(format string, a ...any) error {
	return &untrustedTemporaryFileError{reason: fmt.Sprintf(format, a...)}
}

func temporaryFilePath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(wd, temporaryDir, temporaryFileName), nil
}

// checkTemporaryDir requires .cadenya to be a real directory, not a symlink.
// A missing directory is (false, nil).
func checkTemporaryDir(dir string) (bool, error) {
	info, err := os.Lstat(dir)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if !info.IsDir() {
		return false, untrusted("%s is not a directory (symlinks are refused)", temporaryDir)
	}
	return true, nil
}

// readTemporaryWorkspace loads this directory's temporary workspace file,
// live or not. A missing file is (nil, nil); a file that fails the trust
// checks is an *untrustedTemporaryFileError.
func readTemporaryWorkspace() (*temporaryWorkspace, error) {
	path, err := temporaryFilePath()
	if err != nil {
		return nil, err
	}
	if ok, err := checkTemporaryDir(filepath.Dir(path)); !ok || err != nil {
		return nil, err
	}
	info, err := os.Lstat(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if !info.Mode().IsRegular() {
		return nil, untrusted("it is not a regular file (symlinks are refused)")
	}
	if err := checkPrivateFile(info); err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var tw temporaryWorkspace
	if err := json.Unmarshal(data, &tw); err != nil {
		return nil, untrusted("it does not parse: %v", err)
	}
	if tw.BaseURL != "" && !trustedAPIBaseURL(tw.BaseURL) {
		return nil, untrusted("baseUrl %q is not a Cadenya API", tw.BaseURL)
	}
	tw.BaseURL = normalizeBaseURL(tw.BaseURL)
	if err := checkClaimURL(tw.ClaimURL); err != nil {
		return nil, untrusted("%v", err)
	}
	return &tw, nil
}

// warnedTemporaryFile dedupes the "ignoring" warning per command run.
var warnedTemporaryFile sync.Map

// liveTemporaryWorkspace returns this directory's unexpired temporary
// workspace, or nil. Best effort, like storedProfile: an unreadable file
// means "none" in the middle of an unrelated command, and a file failing the
// trust checks is reported once on stderr.
func liveTemporaryWorkspace(cmd *cli.Command) *temporaryWorkspace {
	tw, err := readTemporaryWorkspace()
	var bad *untrustedTemporaryFileError
	if errors.As(err, &bad) {
		root := cmd.Root()
		if _, seen := warnedTemporaryFile.LoadOrStore(root, true); !seen && root.ErrWriter != nil {
			fmt.Fprintf(root.ErrWriter, "warning: %v\n", err)
		}
	}
	if err != nil || tw == nil || !tw.live(time.Now()) {
		return nil
	}
	return tw
}

// removeTemporaryWorkspace deletes the file, reporting whether it existed.
func removeTemporaryWorkspace() (bool, error) {
	path, err := temporaryFilePath()
	if err != nil {
		return false, err
	}
	if err := os.Remove(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// writeTemporaryWorkspace stores the file at 0600 from the first byte: a
// private temp file renamed over the destination, as for the credentials
// file. It refuses a symlinked .cadenya or temporary.json.
func writeTemporaryWorkspace(tw *temporaryWorkspace) error {
	path, err := temporaryFilePath()
	if err != nil {
		return err
	}
	dir := filepath.Dir(path)
	exists, err := checkTemporaryDir(dir)
	if err != nil {
		return err
	}
	if !exists {
		if err := os.Mkdir(dir, 0o700); err != nil {
			return err
		}
	}
	if info, err := os.Lstat(path); err == nil && !info.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file (symlinks are refused)", filepath.Join(temporaryDir, temporaryFileName))
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

// gitIgnores asks git whether the credential file is ignored. known=false
// when git is missing or the directory is not a repository.
func gitIgnores(wd string) (ignored, known bool) {
	if _, err := exec.LookPath("git"); err != nil {
		return false, false
	}
	err := exec.Command("git", "-C", wd, "check-ignore", "-q", "--", gitignoreEntry).Run()
	if err == nil {
		return true, true
	}
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) && exitErr.ExitCode() == 1 {
		return false, true
	}
	return false, false
}

// scanIgnores is the fallback without git: the last .gitignore line that
// matches the credential file decides, so a later negation re-includes it.
func scanIgnores(data []byte) bool {
	matches := map[string]bool{
		".cadenya": true, ".cadenya/": true, "/.cadenya": true, "/.cadenya/": true,
		".cadenya/*": true, "/.cadenya/*": true, gitignoreEntry: true, "/" + gitignoreEntry: true,
	}
	ignored := false
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "!") {
			if matches[strings.TrimPrefix(line, "!")] {
				ignored = false
			}
			continue
		}
		if matches[line] {
			ignored = true
		}
	}
	return ignored
}

// ensureGitignored adds .cadenya/temporary.json to the working directory's
// .gitignore, creating the file if needed. It returns a sentence describing
// the change, or "" when the file was already ignored.
func ensureGitignored() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path := filepath.Join(wd, ".gitignore")
	info, err := os.Lstat(path)
	exists := err == nil
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return "", err
	}
	if exists && !info.Mode().IsRegular() {
		return "", errors.New(".gitignore is not a regular file (symlinks are refused)")
	}
	var data []byte
	if exists {
		if data, err = os.ReadFile(path); err != nil {
			return "", err
		}
	}
	ignored, known := gitIgnores(wd)
	if !known {
		ignored = scanIgnores(data)
	}
	if ignored {
		return "", nil
	}
	newline := "\n"
	if bytes.Contains(data, []byte("\r\n")) {
		newline = "\r\n"
	}
	var buf bytes.Buffer
	buf.Write(data)
	if len(data) > 0 && data[len(data)-1] != '\n' {
		buf.WriteString(newline)
	}
	buf.WriteString(gitignoreEntry + newline)
	if err := os.WriteFile(path, buf.Bytes(), 0o644); err != nil {
		return "", err
	}
	if known {
		if still, _ := gitIgnores(wd); !still {
			return "", fmt.Errorf("added %s to .gitignore, but git still does not ignore it; check for a negation rule", gitignoreEntry)
		}
	}
	if exists {
		return "Added " + gitignoreEntry + " to .gitignore.", nil
	}
	return "Created .gitignore with " + gitignoreEntry + ".", nil
}
