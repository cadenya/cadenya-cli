package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/urfave/cli/v3"

	commands "go.cadenya.com/cadenya-cli/internal/commands"
)

func aPIKeysCommand() *cli.Command {
	return &cli.Command{
		Name:                      "api-keys",
		DisableSliceFlagSeparator: true,
		Commands: []*cli.Command{
			{
				Name:                      "list",
				DisableSliceFlagSeparator: true,
				Usage:                     "List API keys",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "display", Usage: "Output mode (one of: json, yaml, table, extended)"},
					&cli.StringFlag{Name: "workspace-id", Usage: "Lists only keys that can reach this workspace (query parameter). Accepts the canonical ID or the `external_id:<value>` form."},
					&cli.Int32Flag{Name: "limit", Usage: "Maximum number of results to return."},
					&cli.StringFlag{Name: "cursor", Usage: "Pagination cursor from previous response."},
					&cli.StringFlag{Name: "prefix", Usage: "Filter by ID prefix."},
					&cli.StringFlag{Name: "query", Usage: "Free-form search query."},
					&cli.StringFlag{Name: "labels", Usage: "Filters by metadata labels. Comma-separated key=value pairs, e.g. \"env=prod,team=ai\". A resource matches only if every pair matches exactly (AND semantics)."},
					&cli.StringFlag{Name: "sort-order", Usage: "Sort order for results (asc or desc by creation time)."},
					&cli.BoolFlag{Name: "include-info", Usage: "When true, included info fields are populated. Requests with this flag count more against your rate limit."},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 0 {
						return cli.Exit(fmt.Sprintf("unexpected positional arguments: %v", cmd.Args().Slice()), 2)
					}
					_display := displayMode(cmd, "table")
					if !isOneOf(_display, []string{"json", "yaml", "table", "extended"}) {
						return cli.Exit(fmt.Sprintf("--display: invalid value %q (valid: json, yaml, table, extended)", _display), 2)
					}
					_columns := []displayColumn{{header: "ID", path: []string{"metadata", "id"}}, {header: "EXTERNAL ID", path: []string{"metadata", "externalId"}}, {header: "NAME", path: []string{"metadata", "name"}}, {header: "CREATED", path: []string{"metadata", "createdAt"}}, {header: "STATE", path: []string{"state"}}}
					var converted commands.APIKeysListConversion
					if err := commands.ConvertAPIKeysList(cmd, &converted); err != nil {
						return err
					}
					client, err := newClient(cmd)
					if err != nil {
						return err
					}
					page, err := client.APIKeys().List(ctx, &converted.Params)
					if err != nil {
						return err
					}
					return renderDisplay(_display, _columns, true, map[string]any{"items": page.Items, "nextCursor": page.NextCursor})
				},
			},
			{
				Name:                      "create",
				DisableSliceFlagSeparator: true,
				Usage:                     "Create a new API key",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "display", Usage: "Output mode (one of: json, yaml, table, extended)"},
					&cli.StringFlag{Name: "metadata", Usage: "YAML/JSON document (literal, @path, or - for stdin).", TakesFile: true},
					&cli.StringFlag{Name: "name", Usage: "Required. Human-readable name for the resource (e.g., \"Production API Key\", \"Staging Workspace\")."},
					&cli.StringFlag{Name: "external-id", Usage: "External ID for the resource (e.g., a workflow ID from an external system)."},
					&cli.StringSliceFlag{Name: "label", Usage: "Key-value pairs for categorization and filtering. Values are 0-63 alphanumeric characters with \"-\", \"_\", or \".\" allowed between; keys follow the same shape and…. KEY=VALUE (repeatable; or a document)."},
					&cli.StringFlag{Name: "spec", Usage: "YAML/JSON document (literal, @path, or - for stdin).", TakesFile: true},
					&cli.StringFlag{Name: "description", Usage: "Free-form description of what this API key is used for."},
					&cli.StringSliceFlag{Name: "permission", Usage: "Scopes granted to this key. Each entry is a colon-separated resource:verb string (e.g. \"objectives:manage\"). Resources: agents, objectives, tools, memory,…. Repeatable."},
					&cli.BoolFlag{Name: "all-workspaces", Usage: "Required. When true the key reaches every workspace in the account, including workspaces created later, and workspace_assignments is ignored on requests. Reads then list…."},
					&cli.StringSliceFlag{Name: "workspace-assignment", Usage: "Required. Complete set of workspaces the key is assigned to. Order has no meaning and duplicates are collapsed. On create, at least one assignment is required unless…. key=value,... over workspace-id (repeatable; or a document)."},
					&cli.StringFlag{Name: "file", Aliases: []string{"f"}, TakesFile: true, Usage: "Whole request body from a YAML/JSON file (or - for stdin); other flags override its values"},
					&cli.BoolFlag{Name: "dry-run", Usage: "Print the assembled request body (YAML; JSON with --display json) and exit without calling the API"},
					&cli.BoolFlag{Name: "strict", Usage: "Reject fields the request does not accept in --file and document inputs instead of dropping them with a warning"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 0 {
						return cli.Exit(fmt.Sprintf("unexpected positional arguments: %v", cmd.Args().Slice()), 2)
					}
					_display := displayMode(cmd, "table")
					if !isOneOf(_display, []string{"json", "yaml", "table", "extended"}) {
						return cli.Exit(fmt.Sprintf("--display: invalid value %q (valid: json, yaml, table, extended)", _display), 2)
					}
					_columns := []displayColumn{{header: "ID", path: []string{"metadata", "id"}}, {header: "EXTERNAL ID", path: []string{"metadata", "externalId"}}, {header: "NAME", path: []string{"metadata", "name"}}, {header: "CREATED", path: []string{"metadata", "createdAt"}}, {header: "STATE", path: []string{"state"}}}
					_stdinInputs := []string{cmd.String("file"), cmd.String("metadata"), cmd.String("name"), cmd.String("external-id"), cmd.String("spec"), cmd.String("description")}
					_stdinInputs = append(_stdinInputs, cmd.StringSlice("label")...)
					_stdinInputs = append(_stdinInputs, cmd.StringSlice("workspace-assignment")...)
					if err := stdinBudget(_stdinInputs); err != nil {
						return cli.Exit(err.Error(), 2)
					}
					var converted commands.APIKeysCreateConversion
					if err := commands.ConvertAPIKeysCreate(cmd, &converted); err != nil {
						return err
					}
					if cmd.Bool("dry-run") {
						return printDocument(_display, converted.Body)
					}
					client, err := newClient(cmd)
					if err != nil {
						return err
					}
					out, err := client.APIKeys().Create(ctx, &converted.Params)
					if err != nil {
						return err
					}
					return renderDisplay(_display, _columns, false, out)
				},
			},
			{
				Name:                      "retrieve",
				DisableSliceFlagSeparator: true,
				Usage:                     "Get an API key by ID",
				ArgsUsage:                 "<id>",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "display", Usage: "Output mode (one of: json, yaml, table, extended)"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 1 {
						return cli.Exit(fmt.Sprintf("expected exactly 1 positional argument(s) (<id>), got %d", cmd.Args().Len()), 2)
					}
					if strings.TrimSpace(cmd.Args().Get(0)) == "" {
						return cli.Exit("<id> must not be empty", 2)
					}
					_display := displayMode(cmd, "table")
					if !isOneOf(_display, []string{"json", "yaml", "table", "extended"}) {
						return cli.Exit(fmt.Sprintf("--display: invalid value %q (valid: json, yaml, table, extended)", _display), 2)
					}
					_columns := []displayColumn{{header: "ID", path: []string{"metadata", "id"}}, {header: "EXTERNAL ID", path: []string{"metadata", "externalId"}}, {header: "NAME", path: []string{"metadata", "name"}}, {header: "CREATED", path: []string{"metadata", "createdAt"}}, {header: "STATE", path: []string{"state"}}}
					pos0 := cmd.Args().Get(0) // id
					client, err := newClient(cmd)
					if err != nil {
						return err
					}
					out, err := client.APIKeys().Retrieve(ctx, pos0)
					if err != nil {
						return err
					}
					return renderDisplay(_display, _columns, false, out)
				},
			},
			{
				Name:                      "delete",
				DisableSliceFlagSeparator: true,
				Usage:                     "Delete an API key",
				ArgsUsage:                 "<id>",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "display", Usage: "Output mode (one of: json, yaml, table, extended)"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 1 {
						return cli.Exit(fmt.Sprintf("expected exactly 1 positional argument(s) (<id>), got %d", cmd.Args().Len()), 2)
					}
					if strings.TrimSpace(cmd.Args().Get(0)) == "" {
						return cli.Exit("<id> must not be empty", 2)
					}
					_display := displayMode(cmd, "json")
					if !isOneOf(_display, []string{"json", "yaml", "table", "extended"}) {
						return cli.Exit(fmt.Sprintf("--display: invalid value %q (valid: json, yaml, table, extended)", _display), 2)
					}
					if _display != "json" && _display != "yaml" {
						return cli.Exit("this command has no displayable response; use --display json", 2)
					}
					pos0 := cmd.Args().Get(0) // id
					client, err := newClient(cmd)
					if err != nil {
						return err
					}
					return client.APIKeys().Delete(ctx, pos0)
				},
			},
			{
				Name:                      "update",
				DisableSliceFlagSeparator: true,
				Usage:                     "Update an API key",
				ArgsUsage:                 "<id>",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "display", Usage: "Output mode (one of: json, yaml, table, extended)"},
					&cli.StringFlag{Name: "metadata", Usage: "YAML/JSON document (literal, @path, or - for stdin).", TakesFile: true},
					&cli.StringFlag{Name: "name", Usage: "Human-readable name for the resource (e.g., \"Production API Key\", \"Staging Workspace\")."},
					&cli.StringFlag{Name: "external-id", Usage: "External ID for the resource (e.g., a workflow ID from an external system)."},
					&cli.StringSliceFlag{Name: "label", Usage: "Key-value pairs for categorization and filtering. Values are 0-63 alphanumeric characters with \"-\", \"_\", or \".\" allowed between; keys follow the same shape and…. KEY=VALUE (repeatable; or a document)."},
					&cli.StringFlag{Name: "spec", Usage: "YAML/JSON document (literal, @path, or - for stdin).", TakesFile: true},
					&cli.StringFlag{Name: "description", Usage: "Free-form description of what this API key is used for."},
					&cli.StringSliceFlag{Name: "permission", Usage: "Scopes granted to this key. Each entry is a colon-separated resource:verb string (e.g. \"objectives:manage\"). Resources: agents, objectives, tools, memory,…. Repeatable."},
					&cli.BoolFlag{Name: "all-workspaces", Usage: "When true the key reaches every workspace in the account, including workspaces created later, and workspace_assignments is ignored on requests. Reads then list…."},
					&cli.StringSliceFlag{Name: "workspace-assignment", Usage: "Complete set of workspaces the key is assigned to. Order has no meaning and duplicates are collapsed. On create, at least one assignment is required unless…. key=value,... over workspace-id (repeatable; or a document)."},
					&cli.StringFlag{Name: "update-mask", Usage: "Fields to update. Selecting spec.workspace_assignments replaces the whole set; selecting spec.all_workspaces turns the all-workspaces grant on or off. Turning…."},
					&cli.StringFlag{Name: "file", Aliases: []string{"f"}, TakesFile: true, Usage: "Whole request body from a YAML/JSON file (or - for stdin); other flags override its values"},
					&cli.BoolFlag{Name: "dry-run", Usage: "Print the assembled request body (YAML; JSON with --display json) and exit without calling the API"},
					&cli.BoolFlag{Name: "strict", Usage: "Reject fields the request does not accept in --file and document inputs instead of dropping them with a warning"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 1 {
						return cli.Exit(fmt.Sprintf("expected exactly 1 positional argument(s) (<id>), got %d", cmd.Args().Len()), 2)
					}
					if strings.TrimSpace(cmd.Args().Get(0)) == "" {
						return cli.Exit("<id> must not be empty", 2)
					}
					_display := displayMode(cmd, "table")
					if !isOneOf(_display, []string{"json", "yaml", "table", "extended"}) {
						return cli.Exit(fmt.Sprintf("--display: invalid value %q (valid: json, yaml, table, extended)", _display), 2)
					}
					_columns := []displayColumn{{header: "ID", path: []string{"metadata", "id"}}, {header: "EXTERNAL ID", path: []string{"metadata", "externalId"}}, {header: "NAME", path: []string{"metadata", "name"}}, {header: "CREATED", path: []string{"metadata", "createdAt"}}, {header: "STATE", path: []string{"state"}}}
					_stdinInputs := []string{cmd.String("file"), cmd.String("metadata"), cmd.String("name"), cmd.String("external-id"), cmd.String("spec"), cmd.String("description"), cmd.String("update-mask")}
					_stdinInputs = append(_stdinInputs, cmd.StringSlice("label")...)
					_stdinInputs = append(_stdinInputs, cmd.StringSlice("workspace-assignment")...)
					if err := stdinBudget(_stdinInputs); err != nil {
						return cli.Exit(err.Error(), 2)
					}
					pos0 := cmd.Args().Get(0) // id
					var converted commands.APIKeysUpdateConversion
					if err := commands.ConvertAPIKeysUpdate(cmd, &converted); err != nil {
						return err
					}
					if cmd.Bool("dry-run") {
						return printDocument(_display, converted.Body)
					}
					client, err := newClient(cmd)
					if err != nil {
						return err
					}
					out, err := client.APIKeys().Update(ctx, pos0, &converted.Params)
					if err != nil {
						return err
					}
					return renderDisplay(_display, _columns, false, out)
				},
			},
			{
				Name:                      "add-workspace",
				DisableSliceFlagSeparator: true,
				Usage:                     "Assign a workspace to an API key",
				ArgsUsage:                 "<id>",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "display", Usage: "Output mode (one of: json, yaml, table, extended)"},
					&cli.StringFlag{Name: "workspace-id", Usage: "Required. The workspace to assign. Accepts the canonical ID or the `external_id:<value>` form."},
					&cli.StringFlag{Name: "file", Aliases: []string{"f"}, TakesFile: true, Usage: "Whole request body from a YAML/JSON file (or - for stdin); other flags override its values"},
					&cli.BoolFlag{Name: "dry-run", Usage: "Print the assembled request body (YAML; JSON with --display json) and exit without calling the API"},
					&cli.BoolFlag{Name: "strict", Usage: "Reject fields the request does not accept in --file and document inputs instead of dropping them with a warning"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 1 {
						return cli.Exit(fmt.Sprintf("expected exactly 1 positional argument(s) (<id>), got %d", cmd.Args().Len()), 2)
					}
					if strings.TrimSpace(cmd.Args().Get(0)) == "" {
						return cli.Exit("<id> must not be empty", 2)
					}
					_display := displayMode(cmd, "table")
					if !isOneOf(_display, []string{"json", "yaml", "table", "extended"}) {
						return cli.Exit(fmt.Sprintf("--display: invalid value %q (valid: json, yaml, table, extended)", _display), 2)
					}
					_columns := []displayColumn{{header: "ID", path: []string{"metadata", "id"}}, {header: "EXTERNAL ID", path: []string{"metadata", "externalId"}}, {header: "NAME", path: []string{"metadata", "name"}}, {header: "CREATED", path: []string{"metadata", "createdAt"}}, {header: "STATE", path: []string{"state"}}}
					_stdinInputs := []string{cmd.String("file"), cmd.String("workspace-id")}
					if err := stdinBudget(_stdinInputs); err != nil {
						return cli.Exit(err.Error(), 2)
					}
					pos0 := cmd.Args().Get(0) // id
					var converted commands.APIKeysAddWorkspaceConversion
					if err := commands.ConvertAPIKeysAddWorkspace(cmd, &converted); err != nil {
						return err
					}
					if cmd.Bool("dry-run") {
						return printDocument(_display, converted.Body)
					}
					client, err := newClient(cmd)
					if err != nil {
						return err
					}
					out, err := client.APIKeys().AddWorkspace(ctx, pos0, &converted.Params)
					if err != nil {
						return err
					}
					return renderDisplay(_display, _columns, false, out)
				},
			},
			{
				Name:                      "disable",
				DisableSliceFlagSeparator: true,
				Usage:                     "Disable an API key",
				ArgsUsage:                 "<id>",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "display", Usage: "Output mode (one of: json, yaml, table, extended)"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 1 {
						return cli.Exit(fmt.Sprintf("expected exactly 1 positional argument(s) (<id>), got %d", cmd.Args().Len()), 2)
					}
					if strings.TrimSpace(cmd.Args().Get(0)) == "" {
						return cli.Exit("<id> must not be empty", 2)
					}
					_display := displayMode(cmd, "table")
					if !isOneOf(_display, []string{"json", "yaml", "table", "extended"}) {
						return cli.Exit(fmt.Sprintf("--display: invalid value %q (valid: json, yaml, table, extended)", _display), 2)
					}
					_columns := []displayColumn{{header: "ID", path: []string{"metadata", "id"}}, {header: "EXTERNAL ID", path: []string{"metadata", "externalId"}}, {header: "NAME", path: []string{"metadata", "name"}}, {header: "CREATED", path: []string{"metadata", "createdAt"}}, {header: "STATE", path: []string{"state"}}}
					pos0 := cmd.Args().Get(0) // id
					client, err := newClient(cmd)
					if err != nil {
						return err
					}
					out, err := client.APIKeys().Disable(ctx, pos0)
					if err != nil {
						return err
					}
					return renderDisplay(_display, _columns, false, out)
				},
			},
			{
				Name:                      "enable",
				DisableSliceFlagSeparator: true,
				Usage:                     "Enable an API key",
				ArgsUsage:                 "<id>",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "display", Usage: "Output mode (one of: json, yaml, table, extended)"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 1 {
						return cli.Exit(fmt.Sprintf("expected exactly 1 positional argument(s) (<id>), got %d", cmd.Args().Len()), 2)
					}
					if strings.TrimSpace(cmd.Args().Get(0)) == "" {
						return cli.Exit("<id> must not be empty", 2)
					}
					_display := displayMode(cmd, "table")
					if !isOneOf(_display, []string{"json", "yaml", "table", "extended"}) {
						return cli.Exit(fmt.Sprintf("--display: invalid value %q (valid: json, yaml, table, extended)", _display), 2)
					}
					_columns := []displayColumn{{header: "ID", path: []string{"metadata", "id"}}, {header: "EXTERNAL ID", path: []string{"metadata", "externalId"}}, {header: "NAME", path: []string{"metadata", "name"}}, {header: "CREATED", path: []string{"metadata", "createdAt"}}, {header: "STATE", path: []string{"state"}}}
					pos0 := cmd.Args().Get(0) // id
					client, err := newClient(cmd)
					if err != nil {
						return err
					}
					out, err := client.APIKeys().Enable(ctx, pos0)
					if err != nil {
						return err
					}
					return renderDisplay(_display, _columns, false, out)
				},
			},
			{
				Name:                      "remove-workspace",
				DisableSliceFlagSeparator: true,
				Usage:                     "Remove a workspace from an API key",
				ArgsUsage:                 "<id>",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "display", Usage: "Output mode (one of: json, yaml, table, extended)"},
					&cli.StringFlag{Name: "workspace-id", Usage: "Required. The workspace to remove. Accepts the canonical ID or the `external_id:<value>` form."},
					&cli.StringFlag{Name: "file", Aliases: []string{"f"}, TakesFile: true, Usage: "Whole request body from a YAML/JSON file (or - for stdin); other flags override its values"},
					&cli.BoolFlag{Name: "dry-run", Usage: "Print the assembled request body (YAML; JSON with --display json) and exit without calling the API"},
					&cli.BoolFlag{Name: "strict", Usage: "Reject fields the request does not accept in --file and document inputs instead of dropping them with a warning"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 1 {
						return cli.Exit(fmt.Sprintf("expected exactly 1 positional argument(s) (<id>), got %d", cmd.Args().Len()), 2)
					}
					if strings.TrimSpace(cmd.Args().Get(0)) == "" {
						return cli.Exit("<id> must not be empty", 2)
					}
					_display := displayMode(cmd, "table")
					if !isOneOf(_display, []string{"json", "yaml", "table", "extended"}) {
						return cli.Exit(fmt.Sprintf("--display: invalid value %q (valid: json, yaml, table, extended)", _display), 2)
					}
					_columns := []displayColumn{{header: "ID", path: []string{"metadata", "id"}}, {header: "EXTERNAL ID", path: []string{"metadata", "externalId"}}, {header: "NAME", path: []string{"metadata", "name"}}, {header: "CREATED", path: []string{"metadata", "createdAt"}}, {header: "STATE", path: []string{"state"}}}
					_stdinInputs := []string{cmd.String("file"), cmd.String("workspace-id")}
					if err := stdinBudget(_stdinInputs); err != nil {
						return cli.Exit(err.Error(), 2)
					}
					pos0 := cmd.Args().Get(0) // id
					var converted commands.APIKeysRemoveWorkspaceConversion
					if err := commands.ConvertAPIKeysRemoveWorkspace(cmd, &converted); err != nil {
						return err
					}
					if cmd.Bool("dry-run") {
						return printDocument(_display, converted.Body)
					}
					client, err := newClient(cmd)
					if err != nil {
						return err
					}
					out, err := client.APIKeys().RemoveWorkspace(ctx, pos0, &converted.Params)
					if err != nil {
						return err
					}
					return renderDisplay(_display, _columns, false, out)
				},
			},
			{
				Name:                      "reveal",
				DisableSliceFlagSeparator: true,
				Usage:                     "Reveal an API key's token",
				ArgsUsage:                 "<id>",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "display", Usage: "Output mode (one of: json, yaml, table, extended)"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 1 {
						return cli.Exit(fmt.Sprintf("expected exactly 1 positional argument(s) (<id>), got %d", cmd.Args().Len()), 2)
					}
					if strings.TrimSpace(cmd.Args().Get(0)) == "" {
						return cli.Exit("<id> must not be empty", 2)
					}
					_display := displayMode(cmd, "table")
					if !isOneOf(_display, []string{"json", "yaml", "table", "extended"}) {
						return cli.Exit(fmt.Sprintf("--display: invalid value %q (valid: json, yaml, table, extended)", _display), 2)
					}
					_columns := []displayColumn{{header: "ID", path: []string{"metadata", "id"}}, {header: "EXTERNAL ID", path: []string{"metadata", "externalId"}}, {header: "NAME", path: []string{"metadata", "name"}}, {header: "CREATED", path: []string{"metadata", "createdAt"}}, {header: "STATE", path: []string{"state"}}}
					pos0 := cmd.Args().Get(0) // id
					client, err := newClient(cmd)
					if err != nil {
						return err
					}
					out, err := client.APIKeys().Reveal(ctx, pos0)
					if err != nil {
						return err
					}
					return renderDisplay(_display, _columns, false, out)
				},
			},
			{
				Name:                      "rotate",
				DisableSliceFlagSeparator: true,
				Usage:                     "Rotate an API key",
				ArgsUsage:                 "<id>",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "display", Usage: "Output mode (one of: json, yaml, table, extended)"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 1 {
						return cli.Exit(fmt.Sprintf("expected exactly 1 positional argument(s) (<id>), got %d", cmd.Args().Len()), 2)
					}
					if strings.TrimSpace(cmd.Args().Get(0)) == "" {
						return cli.Exit("<id> must not be empty", 2)
					}
					_display := displayMode(cmd, "table")
					if !isOneOf(_display, []string{"json", "yaml", "table", "extended"}) {
						return cli.Exit(fmt.Sprintf("--display: invalid value %q (valid: json, yaml, table, extended)", _display), 2)
					}
					_columns := []displayColumn{{header: "ID", path: []string{"metadata", "id"}}, {header: "EXTERNAL ID", path: []string{"metadata", "externalId"}}, {header: "NAME", path: []string{"metadata", "name"}}, {header: "CREATED", path: []string{"metadata", "createdAt"}}, {header: "STATE", path: []string{"state"}}}
					pos0 := cmd.Args().Get(0) // id
					client, err := newClient(cmd)
					if err != nil {
						return err
					}
					out, err := client.APIKeys().Rotate(ctx, pos0)
					if err != nil {
						return err
					}
					return renderDisplay(_display, _columns, false, out)
				},
			},
		},
	}
}
