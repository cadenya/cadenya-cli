package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	toml "github.com/pelletier/go-toml/v2"
	cli "github.com/urfave/cli/v3"
)

// RFC 8628 device-authorization login. The endpoints are fixed at
// generation time; the credential is handed over exactly once by the token
// endpoint after browser approval and stored in the credentials file.
const (
	authClientID       = "cadenya-cli"
	authDeviceEndpoint = "https://app.cadenya.com/api/cli/oauth/device_authorization"
	authTokenEndpoint  = "https://app.cadenya.com/api/cli/oauth/token"
)

// authEndpoint rebases a generation-time endpoint onto CADENYA_AUTH_BASE_URL
// when set — a testing knob for preview deployments. The override supplies
// scheme, host, and any path prefix; the endpoint keeps its own path.
func authEndpoint(configured string) string {
	base := os.Getenv("CADENYA_AUTH_BASE_URL")
	if base == "" {
		return configured
	}
	u, err := url.Parse(configured)
	if err != nil {
		return configured
	}
	return strings.TrimSuffix(base, "/") + u.Path
}

// credWorkspace mirrors the token response's `workspaces` extension member.
type credWorkspace struct {
	ID   string `json:"id" toml:"id"`
	Name string `json:"name" toml:"name"`
}

type credProfile struct {
	Credential string          `toml:"credential"`
	Workspaces []credWorkspace `toml:"workspaces,omitempty"`
}

func credentialsPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolving home directory: %w", err)
	}
	return filepath.Join(home, ".cadenya", "credentials"), nil
}

func loadCredentials() (map[string]credProfile, error) {
	path, err := credentialsPath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return map[string]credProfile{}, nil
	}
	if err != nil {
		return nil, err
	}
	profiles := map[string]credProfile{}
	if err := toml.Unmarshal(data, &profiles); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}
	return profiles, nil
}

func saveCredentials(profiles map[string]credProfile) error {
	path, err := credentialsPath()
	if err != nil {
		return err
	}
	if len(profiles) == 0 {
		if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return err
	}
	data, err := toml.Marshal(profiles)
	if err != nil {
		return err
	}
	// 0600 from the first byte: write a private temp file, then rename over
	// the destination so a concurrent reader never sees a torn write.
	tmp, err := os.CreateTemp(filepath.Dir(path), ".credentials-*")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name())
	if err := tmp.Chmod(0o600); err != nil {
		tmp.Close()
		return err
	}
	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmp.Name(), path)
}

func activeProfile(cmd *cli.Command) string {
	if p := cmd.Root().String("profile"); p != "" {
		return p
	}
	return "default"
}

// credentialSource names the credential ordinary commands will use, in
// precedence order: --api-key flag, then the environment, then the stored
// login, then this directory's temporary workspace. MIRRORS newClient —
// keep the two in sync, or status lies.
func credentialSource(cmd *cli.Command) string {
	if cmd.Root().IsSet("api-key") {
		return "--api-key flag"
	}
	if os.Getenv("CADENYA_API_KEY") != "" {
		if storedProfile(cmd) != nil {
			return "$CADENYA_API_KEY (environment) — overrides the stored login"
		}
		return "$CADENYA_API_KEY (environment)"
	}
	if storedProfile(cmd) != nil {
		return fmt.Sprintf("stored login (profile %q)", activeProfile(cmd))
	}
	if tw := liveTemporaryWorkspace(cmd); tw != nil {
		return fmt.Sprintf("temporary workspace %s (.cadenya/temporary.json)", tw.WorkspaceID)
	}
	return "none — set $CADENYA_API_KEY or run `cadenya auth login`"
}

// storedProfile returns the active profile's saved login, or nil. Best
// effort by design: a missing or unreadable credentials file means "not
// logged in", never a hard error in the middle of an unrelated command.
func storedProfile(cmd *cli.Command) *credProfile {
	profiles, err := loadCredentials()
	if err != nil {
		return nil
	}
	if prof, ok := profiles[activeProfile(cmd)]; ok && prof.Credential != "" {
		return &prof
	}
	return nil
}

type deviceAuthorization struct {
	DeviceCode              string `json:"device_code"`
	UserCode                string `json:"user_code"`
	VerificationURI         string `json:"verification_uri"`
	VerificationURIComplete string `json:"verification_uri_complete"`
	ExpiresIn               int    `json:"expires_in"`
	Interval                int    `json:"interval"`
}

type tokenResponse struct {
	AccessToken string          `json:"access_token"`
	TokenType   string          `json:"token_type"`
	Workspaces  []credWorkspace `json:"workspaces"`
	Error       string          `json:"error"`
}

func postAuthForm(ctx context.Context, endpoint string, form url.Values, out any) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return 0, err
	}
	if err := json.Unmarshal(body, out); err != nil {
		return resp.StatusCode, fmt.Errorf("%s answered HTTP %d with a non-JSON body", endpoint, resp.StatusCode)
	}
	return resp.StatusCode, nil
}

// openBrowser starts the platform opener. Failure is not an error: the URL
// is always printed for manual use first.
func openBrowser(target string) {
	switch runtime.GOOS {
	case "darwin":
		_ = exec.Command("open", target).Start()
	case "windows":
		_ = exec.Command("rundll32", "url.dll,FileProtocolHandler", target).Start()
	default:
		_ = exec.Command("xdg-open", target).Start()
	}
}

func authCommand() *cli.Command {
	return &cli.Command{
		Name:  "auth",
		Usage: "Log in and manage stored credentials",
		Commands: []*cli.Command{
			{
				Name:  "login",
				Usage: "Authenticate in the browser and store the credential",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "no-browser", Usage: "Print the verification URL instead of opening a browser"},
				},
				Action: authLogin,
			},
			{
				Name:   "logout",
				Usage:  "Delete the stored credential for the active profile",
				Action: authLogout,
			},
			{
				Name:   "status",
				Usage:  "Show the stored login and which credential source commands will use",
				Action: authStatus,
			},
		},
	}
}

// originOf reduces a URL to scheme://host.
func originOf(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return raw
	}
	return u.Scheme + "://" + u.Host
}

// authDefaultPollInterval is the token poll interval when the device
// authorization response names none.
var authDefaultPollInterval = 5 * time.Second

func authLogin(ctx context.Context, cmd *cli.Command) error {
	out := cmd.Root().Writer
	form := url.Values{"client_id": {authClientID}}
	// A temporary workspace in this directory is claimed by this login: the
	// claim token rides along with the device authorization, and the
	// approval page claims the workspace before issuing the credential.
	claiming := liveTemporaryWorkspace(cmd)
	if claiming != nil && claiming.ClaimToken == "" {
		claiming = nil
	}
	if claiming != nil {
		// A claim token only goes to the auth host of the API the workspace
		// lives behind: never a staging token to production, or the reverse.
		authBase := originOf(authEndpoint(authDeviceEndpoint))
		if want := claimAuthBaseURLs[claiming.effectiveBaseURL()]; want == "" || normalizeBaseURL(want) != authBase {
			fmt.Fprintf(cmd.Root().Writer, "Not claiming the temporary workspace %s with this login: it lives behind %s, and this login goes to %s. Claim it with the link in .cadenya/temporary.json, or log in against the matching environment.\n", claiming.WorkspaceID, claiming.effectiveBaseURL(), authBase)
			claiming = nil
		}
	}
	if claiming != nil {
		form.Set("claim_token", claiming.ClaimToken)
	}
	if host, err := os.Hostname(); err == nil {
		name := host
		if u := os.Getenv("USER"); u != "" {
			name = u + "@" + host
		}
		form.Set("device_name", name)
	}
	var grant deviceAuthorization
	status, err := postAuthForm(ctx, authEndpoint(authDeviceEndpoint), form, &grant)
	if err != nil {
		return err
	}
	if status != http.StatusOK || grant.DeviceCode == "" {
		return fmt.Errorf("device authorization failed (HTTP %d)", status)
	}
	verification := grant.VerificationURIComplete
	if verification == "" {
		verification = grant.VerificationURI
	}

	// The one-time code prints BEFORE the browser opens so the user can
	// visually match it against the approval page. The device_code (the
	// poll secret) is never displayed.
	fmt.Fprintf(out, "Your one-time code: %s\n", grant.UserCode)
	fmt.Fprintf(out, "Approve this login in your browser: %s\n", verification)
	if claiming != nil {
		fmt.Fprintf(out, "Approving this login also claims the temporary workspace %s in .cadenya/temporary.json.\n", claiming.WorkspaceID)
	}
	if !cmd.Bool("no-browser") {
		openBrowser(verification)
	}
	fmt.Fprintln(out, "Waiting for approval...")

	interval := time.Duration(grant.Interval) * time.Second
	if interval <= 0 {
		interval = authDefaultPollInterval
	}
	expires := time.Duration(grant.ExpiresIn) * time.Second
	if expires <= 0 {
		expires = 15 * time.Minute
	}
	deadline := time.Now().Add(expires)
	tokenForm := url.Values{
		"grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
		"device_code": {grant.DeviceCode},
		"client_id":   {authClientID},
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(interval):
		}
		if time.Now().After(deadline) {
			return errors.New("login expired before it was approved; run login again")
		}
		var token tokenResponse
		status, err := postAuthForm(ctx, authEndpoint(authTokenEndpoint), tokenForm, &token)
		if err != nil {
			if ctx.Err() != nil {
				return err
			}
			// Transient transport failure mid-poll (auth servers recycle
			// keep-alive connections on ~5s idle timers — exactly the
			// slow_down gap): the grant is still pending server-side, so
			// keep polling until the deadline instead of aborting a login
			// the user may already have approved in the browser.
			continue
		}
		switch {
		case status == http.StatusOK && token.AccessToken != "":
			profiles, err := loadCredentials()
			if err != nil {
				return err
			}
			profile := activeProfile(cmd)
			profiles[profile] = credProfile{Credential: token.AccessToken, Workspaces: token.Workspaces}
			if err := saveCredentials(profiles); err != nil {
				return err
			}
			path, _ := credentialsPath()
			fmt.Fprintf(out, "Logged in. Credential saved to %s (profile %q).\n", path, profile)
			// The freshest login is worthless while an ambient key outranks
			// it — say so NOW, not in a 403 an hour later.
			if os.Getenv("CADENYA_API_KEY") != "" {
				fmt.Fprintln(out, "Note: CADENYA_API_KEY is set in this environment and takes precedence over the stored login for API commands. Unset it to use this login.")
			}
			if claiming != nil {
				reportLoginClaim(ctx, cmd, out, claiming)
			}
			return nil
		case token.Error == "authorization_pending":
			// Keep polling.
		case token.Error == "slow_down":
			// RFC 8628 §3.5: add 5 seconds to the interval.
			interval += 5 * time.Second
		case token.Error == "access_denied":
			return errors.New("login was denied in the browser")
		case token.Error == "expired_token":
			return errors.New("login expired before it was approved; run login again")
		default:
			return fmt.Errorf("token endpoint failed (HTTP %d: %q)", status, token.Error)
		}
	}
}

// reportLoginClaim confirms whether a login claimed this directory's
// temporary workspace by asking the API with the temporary key, which the
// claim revokes. The login has succeeded either way.
func reportLoginClaim(ctx context.Context, cmd *cli.Command, out io.Writer, tw *temporaryWorkspace) {
	status, err := claimStatusOnce(ctx, temporaryAPIBaseURL(cmd, tw), tw)
	switch {
	case err != nil:
		fmt.Fprintf(out, "Could not confirm whether the temporary workspace %s was claimed (%v). Run `cadenya workspaces claim-status`.\n", tw.WorkspaceID, err)
	case status.State == stateClaimed:
		msg := fmt.Sprintf("Claimed temporary workspace %s. It no longer expires.", tw.WorkspaceID)
		if removed, _ := removeTemporaryWorkspace(); removed {
			msg += " Removed .cadenya/temporary.json."
		}
		fmt.Fprintln(out, msg)
	case status.State == statePending:
		fmt.Fprintf(out, "The claim of temporary workspace %s is finishing. Run `cadenya workspaces claim-status --wait`.\n", tw.WorkspaceID)
	default:
		fmt.Fprintf(out, "The temporary workspace %s was not claimed and still expires at %s. Only a person without a Cadenya account can claim it; with an existing account, create a workspace in it instead.\n",
			tw.WorkspaceID, tw.ExpiresAt.Local().Format("Mon Jan 2 15:04 MST"))
	}
}

func authLogout(_ context.Context, cmd *cli.Command) error {
	profiles, err := loadCredentials()
	if err != nil {
		return err
	}
	profile := activeProfile(cmd)
	if _, ok := profiles[profile]; !ok {
		fmt.Fprintf(cmd.Root().Writer, "No stored credential for profile %q.\n", profile)
		return nil
	}
	delete(profiles, profile)
	if err := saveCredentials(profiles); err != nil {
		return err
	}
	fmt.Fprintf(cmd.Root().Writer, "Logged out profile %q.\n", profile)
	return nil
}

func authStatus(_ context.Context, cmd *cli.Command) error {
	profile := activeProfile(cmd)
	prof := storedProfile(cmd)
	if prof == nil {
		fmt.Fprintf(cmd.Root().Writer, "Profile %q: not logged in. Run `cadenya auth login`.\n", profile)
	} else {
		fmt.Fprintf(cmd.Root().Writer, "Profile %q: logged in.\n", profile)
		for _, ws := range prof.Workspaces {
			fmt.Fprintf(cmd.Root().Writer, "  workspace: %s\n", ws.ID)
		}
	}
	// The line that answers "which key are my commands actually using?"
	fmt.Fprintf(cmd.Root().Writer, "Credential source: %s.\n", credentialSource(cmd))
	return nil
}
