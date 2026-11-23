package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/urfave/cli/v3"

	commands "go.cadenya.com/cadenya-cli/internal/commands"
)

func toolSetRulesCommand() *cli.Command {
	return &cli.Command{
		Name:                      "rules",
		DisableSliceFlagSeparator: true,
		Commands: []*cli.Command{
			{
				Name:                      "list",
				DisableSliceFlagSeparator: true,
				Usage:                     "List tool set rules",
				ArgsUsage:                 "<tool-set-id>",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "display", Usage: "Output mode (one of: json, yaml, table, extended)"},
					&cli.StringFlag{Name: "workspace-id", Usage: "The workspace that owns the tool set."},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 1 {
						return cli.Exit(fmt.Sprintf("expected exactly 1 positional argument(s) (<tool-set-id>), got %d", cmd.Args().Len()), 2)
					}
					if strings.TrimSpace(cmd.Args().Get(0)) == "" {
						return cli.Exit("<tool-set-id> must not be empty", 2)
					}
					_display := displayMode(cmd, "table")
					if !isOneOf(_display, []string{"json", "yaml", "table", "extended"}) {
						return cli.Exit(fmt.Sprintf("--display: invalid value %q (valid: json, yaml, table, extended)", _display), 2)
					}
					if _display != "json" && _display != "yaml" {
						return cli.Exit("no display columns apply to this command; use --display json or yaml", 2)
					}
					_columns := []displayColumn(nil)
					pos0 := cmd.Args().Get(0) // tool-set-id
					var converted commands.ToolSetRulesListConversion
					if err := commands.ConvertToolSetRulesList(cmd, &converted); err != nil {
						return err
					}
					client, err := newClient(cmd)
					if err != nil {
						return err
					}
					out, err := client.ToolSets().Rules().List(ctx, pos0, &converted.Params)
					if err != nil {
						return err
					}
					return renderDisplay(_display, _columns, false, out)
				},
			},
			{
				Name:                      "create",
				DisableSliceFlagSeparator: true,
				Usage:                     "Create a tool set rule",
				ArgsUsage:                 "<tool-set-id>",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "display", Usage: "Output mode (one of: json, yaml, table, extended)"},
					&cli.StringFlag{Name: "workspace-id", Usage: "The workspace that owns the tool set."},
					&cli.StringFlag{Name: "name", Usage: "Required. A short name for the rule. Unique within the tool set, ignoring case."},
					&cli.StringFlag{Name: "description", Usage: "Why the rule exists."},
					&cli.StringFlag{Name: "when", Usage: "CEL expression over `tool`. Empty matches every tool in the set."},
					&cli.StringFlag{Name: "actions", Usage: "What the rule does to the tools it applies to. YAML/JSON document (literal, @path, or - for stdin).", TakesFile: true},
					&cli.StringSliceFlag{Name: "actions-parameter-action", Usage: "Pre-call actions, applied in order. A later action on the same parameter path wins. One YAML/JSON document per occurrence (literal, @path, or -)."},
					&cli.StringSliceFlag{Name: "actions-result-action", Usage: "Post-call actions, applied in order. Result transforms chain: each one reads the previous one's output. Not supported on bare tool sets. One YAML/JSON document per occurrence (literal, @path, or -)."},
					&cli.StringFlag{Name: "actions-widget-argument-exposure", Usage: "Whether matching tools expose their final call arguments in widget events. Unset means no opinion. YAML/JSON document (literal, @path, or - for stdin).", TakesFile: true},
					&cli.BoolFlag{Name: "actions-widget-argument-exposure-enabled", Usage: ""},
					&cli.StringFlag{Name: "actions-auto-mode", Usage: "Optional automatic approval policy. Unset means no opinion. YAML/JSON document (literal, @path, or - for stdin).", TakesFile: true},
					&cli.StringSliceFlag{Name: "actions-auto-mode-question", Usage: "key=value,... over question, threshold (repeatable; or a document)."},
					&cli.StringFlag{Name: "actions-auto-mode-passing-mode", Usage: "One of: all, none, most."},
					&cli.StringFlag{Name: "actions-auto-mode-failure-message", Usage: "Optional guidance returned to the agent when the decision thresholds fail. When empty, a default denial message is used. Client errors never use this message."},
					&cli.BoolFlag{Name: "disabled", Usage: "Create the rule switched off."},
					&cli.Int32Flag{Name: "position", Usage: "Zero-based place to insert the rule at. Rules at and after it move down by one. Omit to append the rule to the end of the list."},
					&cli.StringFlag{Name: "file", Aliases: []string{"f"}, TakesFile: true, Usage: "Whole request body from a YAML/JSON file (or - for stdin); other flags override its values"},
					&cli.BoolFlag{Name: "dry-run", Usage: "Print the assembled request body (YAML; JSON with --display json) and exit without calling the API"},
					&cli.BoolFlag{Name: "strict", Usage: "Reject fields the request does not accept in --file and document inputs instead of dropping them with a warning"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 1 {
						return cli.Exit(fmt.Sprintf("expected exactly 1 positional argument(s) (<tool-set-id>), got %d", cmd.Args().Len()), 2)
					}
					if strings.TrimSpace(cmd.Args().Get(0)) == "" {
						return cli.Exit("<tool-set-id> must not be empty", 2)
					}
					_display := displayMode(cmd, "table")
					if !isOneOf(_display, []string{"json", "yaml", "table", "extended"}) {
						return cli.Exit(fmt.Sprintf("--display: invalid value %q (valid: json, yaml, table, extended)", _display), 2)
					}
					if _display != "json" && _display != "yaml" && !cmd.Bool("dry-run") {
						return cli.Exit("no display columns apply to this command; use --display json or yaml", 2)
					}
					_columns := []displayColumn(nil)
					_stdinInputs := []string{cmd.String("file"), cmd.String("name"), cmd.String("description"), cmd.String("when"), cmd.String("actions"), cmd.String("actions-widget-argument-exposure"), cmd.String("actions-auto-mode"), cmd.String("actions-auto-mode-failure-message")}
					_stdinInputs = append(_stdinInputs, cmd.StringSlice("actions-parameter-action")...)
					_stdinInputs = append(_stdinInputs, cmd.StringSlice("actions-result-action")...)
					_stdinInputs = append(_stdinInputs, cmd.StringSlice("actions-auto-mode-question")...)
					if err := stdinBudget(_stdinInputs); err != nil {
						return cli.Exit(err.Error(), 2)
					}
					pos0 := cmd.Args().Get(0) // tool-set-id
					var converted commands.ToolSetRulesCreateConversion
					if err := commands.ConvertToolSetRulesCreate(cmd, &converted); err != nil {
						return err
					}
					if cmd.Bool("dry-run") {
						return printDocument(_display, converted.Body)
					}
					client, err := newClient(cmd)
					if err != nil {
						return err
					}
					out, err := client.ToolSets().Rules().Create(ctx, pos0, &converted.Params)
					if err != nil {
						return err
					}
					return renderDisplay(_display, _columns, false, out)
				},
			},
			{
				Name:                      "retrieve",
				DisableSliceFlagSeparator: true,
				Usage:                     "Get a tool set rule by ID",
				ArgsUsage:                 "<tool-set-id> <id>",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "display", Usage: "Output mode (one of: json, yaml, table, extended)"},
					&cli.StringFlag{Name: "workspace-id", Usage: "The workspace that owns the tool set."},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 2 {
						return cli.Exit(fmt.Sprintf("expected exactly 2 positional argument(s) (<tool-set-id> <id>), got %d", cmd.Args().Len()), 2)
					}
					if strings.TrimSpace(cmd.Args().Get(0)) == "" {
						return cli.Exit("<tool-set-id> must not be empty", 2)
					}
					if strings.TrimSpace(cmd.Args().Get(1)) == "" {
						return cli.Exit("<id> must not be empty", 2)
					}
					_display := displayMode(cmd, "table")
					if !isOneOf(_display, []string{"json", "yaml", "table", "extended"}) {
						return cli.Exit(fmt.Sprintf("--display: invalid value %q (valid: json, yaml, table, extended)", _display), 2)
					}
					if _display != "json" && _display != "yaml" {
						return cli.Exit("no display columns apply to this command; use --display json or yaml", 2)
					}
					_columns := []displayColumn(nil)
					pos0 := cmd.Args().Get(0) // tool-set-id
					pos1 := cmd.Args().Get(1) // id
					var converted commands.ToolSetRulesRetrieveConversion
					if err := commands.ConvertToolSetRulesRetrieve(cmd, &converted); err != nil {
						return err
					}
					client, err := newClient(cmd)
					if err != nil {
						return err
					}
					out, err := client.ToolSets().Rules().Retrieve(ctx, pos0, pos1, &converted.Params)
					if err != nil {
						return err
					}
					return renderDisplay(_display, _columns, false, out)
				},
			},
			{
				Name:                      "delete",
				DisableSliceFlagSeparator: true,
				Usage:                     "Delete a tool set rule",
				ArgsUsage:                 "<tool-set-id> <id>",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "display", Usage: "Output mode (one of: json, yaml, table, extended)"},
					&cli.StringFlag{Name: "workspace-id", Usage: "The workspace that owns the tool set."},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 2 {
						return cli.Exit(fmt.Sprintf("expected exactly 2 positional argument(s) (<tool-set-id> <id>), got %d", cmd.Args().Len()), 2)
					}
					if strings.TrimSpace(cmd.Args().Get(0)) == "" {
						return cli.Exit("<tool-set-id> must not be empty", 2)
					}
					if strings.TrimSpace(cmd.Args().Get(1)) == "" {
						return cli.Exit("<id> must not be empty", 2)
					}
					_display := displayMode(cmd, "json")
					if !isOneOf(_display, []string{"json", "yaml", "table", "extended"}) {
						return cli.Exit(fmt.Sprintf("--display: invalid value %q (valid: json, yaml, table, extended)", _display), 2)
					}
					if _display != "json" && _display != "yaml" {
						return cli.Exit("this command has no displayable response; use --display json", 2)
					}
					pos0 := cmd.Args().Get(0) // tool-set-id
					pos1 := cmd.Args().Get(1) // id
					var converted commands.ToolSetRulesDeleteConversion
					if err := commands.ConvertToolSetRulesDelete(cmd, &converted); err != nil {
						return err
					}
					client, err := newClient(cmd)
					if err != nil {
						return err
					}
					return client.ToolSets().Rules().Delete(ctx, pos0, pos1, &converted.Params)
				},
			},
			{
				Name:                      "update",
				DisableSliceFlagSeparator: true,
				Usage:                     "Update a tool set rule",
				ArgsUsage:                 "<tool-set-id> <id>",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "display", Usage: "Output mode (one of: json, yaml, table, extended)"},
					&cli.StringFlag{Name: "workspace-id", Usage: "The workspace that owns the tool set."},
					&cli.StringFlag{Name: "name", Usage: ""},
					&cli.StringFlag{Name: "description", Usage: ""},
					&cli.StringFlag{Name: "when", Usage: ""},
					&cli.StringFlag{Name: "actions", Usage: "YAML/JSON document (literal, @path, or - for stdin).", TakesFile: true},
					&cli.StringSliceFlag{Name: "actions-parameter-action", Usage: "Pre-call actions, applied in order. A later action on the same parameter path wins. One YAML/JSON document per occurrence (literal, @path, or -)."},
					&cli.StringSliceFlag{Name: "actions-result-action", Usage: "Post-call actions, applied in order. Result transforms chain: each one reads the previous one's output. Not supported on bare tool sets. One YAML/JSON document per occurrence (literal, @path, or -)."},
					&cli.StringFlag{Name: "actions-widget-argument-exposure", Usage: "Whether matching tools expose their final call arguments in widget events. Unset means no opinion. YAML/JSON document (literal, @path, or - for stdin).", TakesFile: true},
					&cli.BoolFlag{Name: "actions-widget-argument-exposure-enabled", Usage: ""},
					&cli.StringFlag{Name: "actions-auto-mode", Usage: "Optional automatic approval policy. Unset means no opinion. YAML/JSON document (literal, @path, or - for stdin).", TakesFile: true},
					&cli.StringSliceFlag{Name: "actions-auto-mode-question", Usage: "key=value,... over question, threshold (repeatable; or a document)."},
					&cli.StringFlag{Name: "actions-auto-mode-passing-mode", Usage: "One of: all, none, most."},
					&cli.StringFlag{Name: "actions-auto-mode-failure-message", Usage: "Optional guidance returned to the agent when the decision thresholds fail. When empty, a default denial message is used. Client errors never use this message."},
					&cli.BoolFlag{Name: "disabled", Usage: ""},
					&cli.Int32Flag{Name: "position", Usage: "Moves the rule to this zero-based place. The rules in between shift by one to make room."},
					&cli.StringFlag{Name: "update-mask", Usage: "Fields to update: any of `name`, `description`, `when`, `actions`, `disabled` and `position`."},
					&cli.StringFlag{Name: "file", Aliases: []string{"f"}, TakesFile: true, Usage: "Whole request body from a YAML/JSON file (or - for stdin); other flags override its values"},
					&cli.BoolFlag{Name: "dry-run", Usage: "Print the assembled request body (YAML; JSON with --display json) and exit without calling the API"},
					&cli.BoolFlag{Name: "strict", Usage: "Reject fields the request does not accept in --file and document inputs instead of dropping them with a warning"},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					if cmd.Args().Len() != 2 {
						return cli.Exit(fmt.Sprintf("expected exactly 2 positional argument(s) (<tool-set-id> <id>), got %d", cmd.Args().Len()), 2)
					}
					if strings.TrimSpace(cmd.Args().Get(0)) == "" {
						return cli.Exit("<tool-set-id> must not be empty", 2)
					}
					if strings.TrimSpace(cmd.Args().Get(1)) == "" {
						return cli.Exit("<id> must not be empty", 2)
					}
					_display := displayMode(cmd, "table")
					if !isOneOf(_display, []string{"json", "yaml", "table", "extended"}) {
						return cli.Exit(fmt.Sprintf("--display: invalid value %q (valid: json, yaml, table, extended)", _display), 2)
					}
					if _display != "json" && _display != "yaml" && !cmd.Bool("dry-run") {
						return cli.Exit("no display columns apply to this command; use --display json or yaml", 2)
					}
					_columns := []displayColumn(nil)
					_stdinInputs := []string{cmd.String("file"), cmd.String("name"), cmd.String("description"), cmd.String("when"), cmd.String("actions"), cmd.String("actions-widget-argument-exposure"), cmd.String("actions-auto-mode"), cmd.String("actions-auto-mode-failure-message"), cmd.String("update-mask")}
					_stdinInputs = append(_stdinInputs, cmd.StringSlice("actions-parameter-action")...)
					_stdinInputs = append(_stdinInputs, cmd.StringSlice("actions-result-action")...)
					_stdinInputs = append(_stdinInputs, cmd.StringSlice("actions-auto-mode-question")...)
					if err := stdinBudget(_stdinInputs); err != nil {
						return cli.Exit(err.Error(), 2)
					}
					pos0 := cmd.Args().Get(0) // tool-set-id
					pos1 := cmd.Args().Get(1) // id
					var converted commands.ToolSetRulesUpdateConversion
					if err := commands.ConvertToolSetRulesUpdate(cmd, &converted); err != nil {
						return err
					}
					if cmd.Bool("dry-run") {
						return printDocument(_display, converted.Body)
					}
					client, err := newClient(cmd)
					if err != nil {
						return err
					}
					out, err := client.ToolSets().Rules().Update(ctx, pos0, pos1, &converted.Params)
					if err != nil {
						return err
					}
					return renderDisplay(_display, _columns, false, out)
				},
			},
		},
	}
}
