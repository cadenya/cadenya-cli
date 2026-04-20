package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/cadenya/cadenya-cli/pkg/config"
	"github.com/cadenya/cadenya-go"
	"github.com/cadenya/cadenya-go/option"
	"github.com/urfave/cli/v3"
)

// configApply wires `cadenya config apply <path>` into urfave/cli.
//
// The apply pipeline:
//  1. Parse YAML (pkg/config.ParseFile) — resolves env vars, fails on unset.
//  2. Expand file globs (pkg/config.ExpandFromFiles).
//  3. Build plan (pkg/config.Build) — GET every resource to determine
//     create vs update vs no-change. Fails on unresolvable references.
//  4. If --plan, render and exit 0.
//  5. Otherwise render the plan and then Apply.
//
// Exit codes:
//
//	0 — plan + apply both succeeded (including all-no-change).
//	1 — apply phase failed partway. Output lists {applied, failed, skipped}.
//	2 — plan phase failed. No writes performed.
var configApply = cli.Command{
	Name:      "apply",
	Usage:     "Apply a workspace configuration from a YAML file",
	UsageText: "cadenya config apply <config.yaml> [--plan]",
	Suggest:   true,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "plan",
			Usage: "Show the plan but do not apply any changes",
		},
	},
	Action:          handleConfigApply,
	HideHelpCommand: true,
}

func handleConfigApply(ctx context.Context, cmd *cli.Command) error {
	args := cmd.Args().Slice()
	if len(args) == 0 {
		return fmt.Errorf("config apply: missing path to YAML file")
	}
	if len(args) > 1 {
		return fmt.Errorf("config apply: unexpected extra arguments: %v", args[1:])
	}
	path := args[0]

	// 1. Parse. `secrets` captures env-var substitutions so the --debug
	// middleware can redact their values from logged request/response bodies
	// before they leak into CI logs.
	cfg, baseDir, secrets, err := config.ParseFile(path)
	if err != nil {
		return exitWith(2, err)
	}

	// 2. Expand file globs (entries_from_files).
	if err := config.ExpandFromFiles(cfg, baseDir); err != nil {
		return exitWith(2, err)
	}

	// 3. Plan.
	opts := getDefaultRequestOptions(cmd)
	if cmd.Bool("debug") {
		opts = append(opts, option.WithMiddleware(config.NewDebugMiddleware(secrets)))
	}
	client := cadenya.NewClient(opts...)
	plan, err := config.Build(ctx, client, cfg)
	if err != nil {
		return exitWith(2, err)
	}
	if err := config.RenderPlan(os.Stderr, plan); err != nil {
		return err
	}
	if cmd.Bool("plan") {
		return nil
	}

	// 4. Apply.
	result, applyErr := config.Apply(ctx, client, plan, os.Stderr)
	if result != nil {
		config.RenderSummary(os.Stderr, result)
	}
	if applyErr != nil {
		return exitWith(1, applyErr)
	}
	return nil
}

// exitWith wraps an error so urfave/cli's default exit-code handling returns
// the specified code. urfave/cli/v3 honors cli.Exit for exit-code selection.
func exitWith(code int, err error) error {
	return cli.Exit(err.Error(), code)
}
