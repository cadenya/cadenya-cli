// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/cadenya/cadenya-cli/internal/apiquery"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
	"github.com/cadenya/cadenya-go"
	"github.com/cadenya/cadenya-go/option"
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
		&requestflag.Flag[bool]{
			Name:      "include-info",
			Usage:     "When set to true you may use more of your alloted API rate-limit",
			QueryPath: "includeInfo",
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

func handleObjectivesToolCallsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workspace-id") && len(unusedArgs) > 0 {
		cmd.Set("workspace-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Objectives.ToolCalls.Get(
		ctx,
		cmd.Value("workspace-id").(string),
		cmd.Value("objective-id").(string),
		cmd.Value("tool-call-id").(string),
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
	if !cmd.IsSet("workspace-id") && len(unusedArgs) > 0 {
		cmd.Set("workspace-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := cadenya.ObjectiveToolCallListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Objectives.ToolCalls.List(
			ctx,
			cmd.Value("workspace-id").(string),
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
			cmd.Value("workspace-id").(string),
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
	if !cmd.IsSet("workspace-id") && len(unusedArgs) > 0 {
		cmd.Set("workspace-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := cadenya.ObjectiveToolCallApproveParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Objectives.ToolCalls.Approve(
		ctx,
		cmd.Value("workspace-id").(string),
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
	if !cmd.IsSet("workspace-id") && len(unusedArgs) > 0 {
		cmd.Set("workspace-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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

	params := cadenya.ObjectiveToolCallDenyParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Objectives.ToolCalls.Deny(
		ctx,
		cmd.Value("workspace-id").(string),
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
