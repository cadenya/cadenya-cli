package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	cli "github.com/urfave/cli/v3"
)

// withTemporaryWorkspaceCommands adds the hand-written temporary workspace
// commands to the workspaces group, keeping cmd_workspaces.go comparable
// with a Redwood reference build.
func withTemporaryWorkspaceCommands(group *cli.Command) *cli.Command {
	group.Commands = append(group.Commands, workspacesCreateCommand(), workspacesClaimStatusCommand())
	return group
}

func workspacesCreateCommand() *cli.Command {
	return &cli.Command{
		Name:    "create",
		Aliases: []string{"new"},
		Usage:   "Create a temporary workspace without an account",
		Description: `Creates a temporary workspace and API key without signing up, and returns a
claim URL. Open the claim URL and sign in or sign up to keep the workspace; an
unclaimed workspace is permanently deleted when it expires (24 hours by
default). Claiming revokes the temporary API key: run ` + "`cadenya auth login`" + `
in the same directory to claim and log in in one step.

The credential is stored in ./.cadenya/temporary.json (mode 0600), and that file
is added to .gitignore. cadenya commands run from this directory use it when no
other credential is configured. Run again in the same directory to reuse the
live workspace; --fresh creates a new one.

Refused when a credential is already configured (--api-key, $CADENYA_API_KEY or
a stored login) and in CI ($CI set). To create a workspace in your account, use
` + "`cadenya workspace-admin create`" + `.`,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "temporary", Usage: "Create a temporary workspace that a person claims later (required)"},
			&cli.StringFlag{Name: "name", Usage: "Display name for the workspace (default: generated)"},
			&cli.DurationFlag{Name: "expire-after", Value: 24 * time.Hour, Usage: "How long the workspace lives unless claimed, from 1h to 24h"},
			&cli.BoolFlag{Name: "fresh", Usage: "Create a new temporary workspace even if this directory has a live one"},
			&cli.BoolFlag{Name: "instructions", Usage: "Print the guide for coding agents and create nothing"},
			&cli.StringFlag{Name: "new-base-url", Sources: cli.EnvVars("CADENYA_NEW_BASE_URL"), Usage: "Provisioning base URL (default: https://new.cadenya.com, or https://new.cadenya.dev when --base-url is the staging API)"},
			&cli.StringFlag{Name: "display", Usage: "Output mode (one of: json, yaml, table, extended); json prints one envelope for agents"},
		},
		Action: workspacesCreate,
	}
}

type nextCommand struct {
	Command string `json:"command"`
	When    string `json:"when"`
}

// temporaryEnvelope is the --display json output of `workspaces create
// --temporary`. Agent skills depend on these field names; keep them stable.
type temporaryEnvelope struct {
	Status      string        `json:"status"`
	WorkspaceID string        `json:"workspaceId"`
	ClaimURL    string        `json:"claimUrl"`
	ExpiresAt   string        `json:"expiresAt"`
	APIKey      string        `json:"apiKey"`
	Reused      bool          `json:"reused"`
	NextStep    string        `json:"nextStep"`
	Next        []nextCommand `json:"next"`
}

var temporaryNextCommands = []nextCommand{
	{Command: "cadenya workspaces claim-status --display json", When: "to check whether the workspace has been claimed"},
	{Command: "cadenya workspaces claim-status --wait", When: "to wait until the user claims the workspace or it expires"},
	{Command: "cadenya auth login", When: "the user wants to claim the workspace and log in in one step"},
}

func workspacesCreate(ctx context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() != 0 {
		return cli.Exit(fmt.Sprintf("unexpected positional arguments: %v", cmd.Args().Slice()), 2)
	}
	stdout, stderr := cmd.Root().Writer, cmd.Root().ErrWriter
	if cmd.Bool("instructions") {
		return printTemporaryInstructions(ctx, stdout)
	}
	if !cmd.Bool("temporary") {
		return cli.Exit("workspaces create only creates temporary workspaces: pass --temporary. To create a workspace in your account, use `cadenya workspace-admin create`.", 2)
	}
	display := displayMode(cmd, "table")
	if !isOneOf(display, []string{"json", "yaml", "table", "extended"}) {
		return cli.Exit(fmt.Sprintf("--display: invalid value %q (valid: json, yaml, table, extended)", display), 2)
	}
	var expireAfter string
	if cmd.IsSet("expire-after") {
		d := cmd.Duration("expire-after")
		if d < time.Hour || d > 24*time.Hour {
			return cli.Exit(fmt.Sprintf("--expire-after: %s is out of range (1h to 24h)", d), 2)
		}
		expireAfter = fmt.Sprintf("%ds", int64(d/time.Second))
	}
	if source := explicitCredentialSource(cmd); source != "" {
		return cli.Exit(fmt.Sprintf("a Cadenya credential is already configured (%s), so --temporary is refused: temporary workspaces are for people without an account. Use your account with `cadenya auth login`, or create a workspace in it with `cadenya workspace-admin create`.", source), 2)
	}
	if inCI() {
		return cli.Exit("--temporary is refused in CI ($CI is set): a temporary workspace needs a person to claim it. Create an API key in the dashboard and set CADENYA_API_KEY in CI.", 2)
	}

	now := time.Now()
	if !cmd.Bool("fresh") {
		if tw := liveTemporaryWorkspace(); tw != nil {
			if display != "json" && display != "yaml" {
				fmt.Fprintf(stdout, "Reusing temporary workspace %s from .cadenya/temporary.json (--fresh creates a new one).\n", tw.WorkspaceID)
			}
			fmt.Fprintln(stderr, claimNotice(tw, now))
			return renderTemporary(display, stdout, tw, true)
		}
	}

	newBase := newBaseURL(cmd)
	resp, err := createTemporaryWorkspace(ctx, newBase, createTemporaryRequest{
		Name:        cmd.String("name"),
		ExpireAfter: expireAfter,
		Source:      temporaryWorkspaceSource(),
	})
	if err != nil {
		return err
	}
	tw := &temporaryWorkspace{
		WorkspaceID: resp.WorkspaceID,
		AccountID:   resp.AccountID,
		APIKey:      resp.APIKey,
		ClaimURL:    resp.ClaimURL,
		ClaimToken:  resp.ClaimToken,
		ExpiresAt:   resp.ExpiresAt,
		NextStep:    resp.NextStep,
		BaseURL:     apiBaseURLForNew(cmd, newBase),
	}
	if err := writeTemporaryWorkspace(tw); err != nil {
		return fmt.Errorf("saving .cadenya/temporary.json: %w (workspace %s was created; claim it at %s)", err, tw.WorkspaceID, tw.ClaimURL)
	}
	note, err := ensureGitignored()
	if err != nil {
		fmt.Fprintf(stderr, "warning: could not add "+gitignoreEntry+" to .gitignore: %v. Add it yourself; the file holds an API key.\n", err)
	} else if note != "" {
		fmt.Fprintln(stderr, note)
	}
	if display != "json" && display != "yaml" {
		fmt.Fprintf(stdout, "Created temporary workspace %s.\n", tw.WorkspaceID)
		fmt.Fprintln(stdout, "Its API key is saved in .cadenya/temporary.json; cadenya commands run from this directory use it.")
	}
	fmt.Fprintln(stderr, claimNotice(tw, now))
	return renderTemporary(display, stdout, tw, false)
}

func renderTemporary(display string, w io.Writer, tw *temporaryWorkspace, reused bool) error {
	if display != "json" && display != "yaml" {
		return nil
	}
	status := "created"
	if reused {
		status = "reused"
	}
	nextStep := tw.NextStep
	if nextStep == "" {
		nextStep = fmt.Sprintf("Open %s and sign in or sign up before %s to keep this workspace; it is deleted after that.", tw.ClaimURL, tw.ExpiresAt.UTC().Format(time.RFC3339))
	}
	envelope := temporaryEnvelope{
		Status:      status,
		WorkspaceID: tw.WorkspaceID,
		ClaimURL:    tw.ClaimURL,
		ExpiresAt:   tw.ExpiresAt.UTC().Format(time.RFC3339),
		APIKey:      tw.APIKey,
		Reused:      reused,
		NextStep:    nextStep,
		Next:        temporaryNextCommands,
	}
	if display == "yaml" {
		return writeYAML(w, envelope)
	}
	return writeJSON(w, envelope)
}

func printTemporaryInstructions(ctx context.Context, w io.Writer) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, temporaryInstructionsURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "cadenya-cli/"+version)
	resp, err := temporaryHTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("fetching the guide from %s: %w", temporaryInstructionsURL, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fetching the guide from %s: HTTP %d", temporaryInstructionsURL, resp.StatusCode)
	}
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return err
	}
	fmt.Fprintln(w, "Nothing has been created yet. Run `cadenya workspaces create --temporary` to create the workspace.")
	fmt.Fprintln(w)
	_, err = w.Write(body)
	return err
}

func workspacesClaimStatusCommand() *cli.Command {
	return &cli.Command{
		Name:  "claim-status",
		Usage: "Show whether this directory's temporary workspace has been claimed",
		Description: `Reads ./.cadenya/temporary.json and asks the API, with the temporary key,
whether the workspace has been claimed. state is one of STATE_UNCLAIMED,
STATE_PENDING (a claim is finishing), STATE_CLAIMED or STATE_EXPIRED.

--wait polls at the interval the server asks for until the state is
STATE_CLAIMED or STATE_EXPIRED. Once claimed, the temporary key no longer
works: .cadenya/temporary.json is removed, and ` + "`cadenya auth login`" + ` gets a new
credential.`,
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "wait", Usage: "Poll until the workspace is claimed or expires"},
			&cli.StringFlag{Name: "display", Usage: "Output mode (one of: json, yaml, table, extended)"},
		},
		Action: workspacesClaimStatus,
	}
}

func workspacesClaimStatus(ctx context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() != 0 {
		return cli.Exit(fmt.Sprintf("unexpected positional arguments: %v", cmd.Args().Slice()), 2)
	}
	display := displayMode(cmd, "table")
	if !isOneOf(display, []string{"json", "yaml", "table", "extended"}) {
		return cli.Exit(fmt.Sprintf("--display: invalid value %q (valid: json, yaml, table, extended)", display), 2)
	}
	stdout, stderr := cmd.Root().Writer, cmd.Root().ErrWriter
	tw, err := readTemporaryWorkspace()
	if err != nil {
		return err
	}
	if tw == nil {
		return cli.Exit("no temporary workspace in this directory (.cadenya/temporary.json not found). "+temporaryCreateHint, 1)
	}
	base := temporaryAPIBaseURL(cmd, tw)
	lastState := ""
	for {
		status, err := claimStatusOnce(ctx, base, tw)
		if err != nil {
			return err
		}
		if status.State == stateClaimed {
			_ = removeTemporaryWorkspace()
		}
		done := status.State == stateClaimed || status.State == stateExpired
		if !cmd.Bool("wait") || done {
			if status.State == stateClaimed {
				fmt.Fprintln(stderr, "The workspace has been claimed. Removed .cadenya/temporary.json: the temporary key no longer works. Run `cadenya auth login` in this directory, or create an API key for the workspace in the dashboard.")
			}
			return renderClaimStatus(display, stdout, status)
		}
		if status.State != lastState {
			fmt.Fprintf(stderr, "%s: waiting for the claim (claim URL: %s)...\n", status.State, tw.ClaimURL)
			lastState = status.State
		}
		interval := 5 * time.Second
		if d, err := time.ParseDuration(status.PollInterval); err == nil && d > 0 {
			interval = d
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(interval):
		}
	}
}

// claimStatusOnce fetches the claim status. The revoked key of a claimed
// workspace answers TEMPORARY_WORKSPACE_CLAIMED, which is the claimed state
// for this command; a rejected key past expiresAt is the expired state.
func claimStatusOnce(ctx context.Context, base string, tw *temporaryWorkspace) (*temporaryStatus, error) {
	status, err := fetchTemporaryStatus(ctx, base, tw)
	if err == nil {
		return status, nil
	}
	reason, apiErr := errorInfoReason(err)
	switch {
	case reason == reasonTemporaryWorkspaceClaimed:
		return &temporaryStatus{WorkspaceID: tw.WorkspaceID, State: stateClaimed}, nil
	case apiErr != nil && apiErr.StatusCode == http.StatusUnauthorized && !time.Now().Before(tw.ExpiresAt):
		expires := tw.ExpiresAt
		return &temporaryStatus{WorkspaceID: tw.WorkspaceID, State: stateExpired, ExpiresAt: &expires}, nil
	}
	return nil, err
}

func renderClaimStatus(display string, w io.Writer, status *temporaryStatus) error {
	switch display {
	case "json":
		return writeJSON(w, status)
	case "yaml":
		return writeYAML(w, status)
	}
	fmt.Fprintf(w, "Workspace: %s\nState:     %s\n", status.WorkspaceID, status.State)
	if status.ExpiresAt != nil {
		fmt.Fprintf(w, "Expires:   %s\n", status.ExpiresAt.Local().Format("Mon Jan 2 15:04 MST"))
	}
	if status.ClaimedAt != nil {
		fmt.Fprintf(w, "Claimed:   %s\n", status.ClaimedAt.Local().Format("Mon Jan 2 15:04 MST"))
	}
	return nil
}

// explicitCredentialSource names a configured account credential, or "".
// The temporary file itself does not count.
func explicitCredentialSource(cmd *cli.Command) string {
	switch {
	case cmd.Root().IsSet("api-key"):
		return "--api-key or $CADENYA_API_KEY"
	case os.Getenv("CADENYA_API_KEY") != "":
		return "$CADENYA_API_KEY"
	case storedProfile(cmd) != nil:
		return fmt.Sprintf("stored login, profile %q", activeProfile(cmd))
	}
	return ""
}
