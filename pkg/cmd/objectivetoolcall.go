// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"go.cadenya.com/cadenya-go"
	"go.cadenya.com/cadenya-go/option"
	"os"

	"github.com/cadenya/cli/internal/apiquery"
	"github.com/cadenya/cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var objectivesToolCallsList = cli.Command{
	Name:    "list",
	Usage:   "Lists all tool calls for an objective",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "objective-id",
			Required: true,
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
			Name:     "objective-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "tool-call-id",
			Required: true,
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
			Name:     "objective-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "tool-call-id",
			Required: true,
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

func handleObjectivesToolCallsList(ctx context.Context, cmd *cli.Command) error {
	client := gocadenyacomcadenyago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := gocadenyacomcadenyago.ObjectiveToolCallListParams{}

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

	format := cmd.Root().String("format")
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
		return ShowJSON(os.Stdout, "objectives:tool-calls list", obj, format, transform)
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
		return ShowJSONIterator(os.Stdout, "objectives:tool-calls list", iter, format, transform, maxItems)
	}
}

func handleObjectivesToolCallsApprove(ctx context.Context, cmd *cli.Command) error {
	client := gocadenyacomcadenyago.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := gocadenyacomcadenyago.ObjectiveToolCallApproveParams{}

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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "objectives:tool-calls approve", obj, format, transform)
}

func handleObjectivesToolCallsDeny(ctx context.Context, cmd *cli.Command) error {
	client := gocadenyacomcadenyago.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := gocadenyacomcadenyago.ObjectiveToolCallDenyParams{}

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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "objectives:tool-calls deny", obj, format, transform)
}
