// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"go.cadenya.com/cadenya-go"
	"go.cadenya.com/cadenya-go/option"

	"github.com/cadenya/cadenya-cli/internal/apiquery"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var objectivesToolCallsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a single tool call, including the content the tool returned. Media\ncontent (images, audio) is served as short-lived signed URLs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "objective-id",
			Required:  true,
			PathParam: "objectiveId",
		},
		&requestflag.Flag[string]{
			Name:      "tool-call-id",
			Required:  true,
			PathParam: "toolCallId",
		},
	},
	Action:          handleObjectivesToolCallsRetrieve,
	HideHelpCommand: true,
}

var objectivesToolCallsList = cli.Command{
	Name:    "list",
	Usage:   "Lists all tool calls for an objective",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "objective-id",
			Required:  true,
			PathParam: "objectiveId",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor from previous response",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "execution-status",
			Usage:     "Filter by tool call execution status. Useful for reverse-harness\n polling of bare tool calls waiting for externally supplied content\n (TOOL_CALL_EXECUTION_STATUS_WAITING_FOR_CONTENT).",
			QueryPath: "executionStatus",
		},
		&requestflag.Flag[bool]{
			Name:      "include-info",
			Usage:     "When set to true you may use more of your alloted API rate-limit",
			QueryPath: "includeInfo",
		},
		&requestflag.Flag[string]{
			Name:      "labels",
			Usage:     "Filters by metadata labels. Comma-separated key=value pairs,\n e.g. \"env=prod,team=ai\". A resource matches only if every pair\n matches exactly (AND semantics).",
			QueryPath: "labels",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by tool call status",
			QueryPath: "status",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleObjectivesToolCallsList,
	HideHelpCommand: true,
}

var objectivesToolCallsApprove = cli.Command{
	Name:    "approve",
	Usage:   "When an agent attempts to use a tool that requires approval, use this endpoint\nto mark it as approved.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "objective-id",
			Required:  true,
			PathParam: "objectiveId",
		},
		&requestflag.Flag[string]{
			Name:      "tool-call-id",
			Required:  true,
			PathParam: "toolCallId",
		},
	},
	Action:          handleObjectivesToolCallsApprove,
	HideHelpCommand: true,
}

var objectivesToolCallsDeny = cli.Command{
	Name:    "deny",
	Usage:   "When an agent attempts to use a tool that requires approval, use this endpoint\nto mark it as denied. Use a memo to steer the LLM to a different decision or\nusage of the tool.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "objective-id",
			Required:  true,
			PathParam: "objectiveId",
		},
		&requestflag.Flag[string]{
			Name:      "tool-call-id",
			Required:  true,
			PathParam: "toolCallId",
		},
		&requestflag.Flag[string]{
			Name:     "memo",
			Usage:    "A memo to associate to the tool call denial. Use a memo to steer the LLM to a different decision or usage of the tool.",
			BodyPath: "memo",
		},
	},
	Action:          handleObjectivesToolCallsDeny,
	HideHelpCommand: true,
}

var objectivesToolCallsSetContent = cli.Command{
	Name:    "set-content",
	Usage:   "For bare tool calls (tool sets with no execution adapter), sets the content an\nexternal API consumer supplies for the call — used for human-in-the-loop tools\nand reverse harnesses that execute tools locally and report results back.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "objective-id",
			Required:  true,
			PathParam: "objectiveId",
		},
		&requestflag.Flag[string]{
			Name:      "tool-call-id",
			Required:  true,
			PathParam: "toolCallId",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "content",
			Usage:    "The content to set on the tool call. Mirrors\n ObjectiveToolCallResult.ContentBlock but writable: media blocks carry\n raw data on input where the result-side carries a signed url on output.",
			Required: true,
			BodyPath: "content",
		},
	},
	Action:          handleObjectivesToolCallsSetContent,
	HideHelpCommand: true,
}

func handleObjectivesToolCallsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("tool-call-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-call-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := cadenya.ObjectiveToolCallGetParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Objectives.ToolCalls.Get(
		ctx,
		cmd.Value("objective-id").(string),
		cmd.Value("tool-call-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "objectives:tool-calls retrieve",
		Transform:      transform,
	})
}

func handleObjectivesToolCallsList(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := cadenya.ObjectiveToolCallListParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Objectives.ToolCalls.List(
			ctx,
			cmd.Value("objective-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "objectives:tool-calls list",
			Transform:      transform,
		})
	} else {
		iter := client.Objectives.ToolCalls.ListAutoPaging(
			ctx,
			cmd.Value("objective-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "objectives:tool-calls list",
			Transform:      transform,
		})
	}
}

func handleObjectivesToolCallsApprove(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("tool-call-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-call-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := cadenya.ObjectiveToolCallApproveParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Objectives.ToolCalls.Approve(
		ctx,
		cmd.Value("objective-id").(string),
		cmd.Value("tool-call-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "objectives:tool-calls approve",
		Transform:      transform,
	})
}

func handleObjectivesToolCallsDeny(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("tool-call-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-call-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := cadenya.ObjectiveToolCallDenyParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Objectives.ToolCalls.Deny(
		ctx,
		cmd.Value("objective-id").(string),
		cmd.Value("tool-call-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "objectives:tool-calls deny",
		Transform:      transform,
	})
}

func handleObjectivesToolCallsSetContent(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("tool-call-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-call-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := cadenya.ObjectiveToolCallSetContentParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Objectives.ToolCalls.SetContent(
		ctx,
		cmd.Value("objective-id").(string),
		cmd.Value("tool-call-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "objectives:tool-calls set-content",
		Transform:      transform,
	})
}
