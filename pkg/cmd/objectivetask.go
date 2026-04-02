// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/cadenya/cadenya-sdk-go"
	"github.com/cadenya/cadenya-sdk-go/option"
	"github.com/stainless-sdks/cadenya-cli/internal/apiquery"
	"github.com/stainless-sdks/cadenya-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var objectivesTasksRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a task by ID from an objective",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "objective-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleObjectivesTasksRetrieve,
	HideHelpCommand: true,
}

var objectivesTasksList = cli.Command{
	Name:    "list",
	Usage:   "Lists all tasks for an objective",
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
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "sort-order",
			Usage:     "Sort order for results",
			QueryPath: "sortOrder",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleObjectivesTasksList,
	HideHelpCommand: true,
}

func handleObjectivesTasksRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := gocadenyacomcadenyasdkgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
	_, err = client.Objectives.Tasks.Get(
		ctx,
		cmd.Value("objective-id").(string),
		cmd.Value("id").(string),
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "objectives:tasks retrieve", obj, format, transform)
}

func handleObjectivesTasksList(ctx context.Context, cmd *cli.Command) error {
	client := gocadenyacomcadenyasdkgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := gocadenyacomcadenyasdkgo.ObjectiveTaskListParams{}

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
		_, err = client.Objectives.Tasks.List(
			ctx,
			cmd.Value("objective-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "objectives:tasks list", obj, format, transform)
	} else {
		iter := client.Objectives.Tasks.ListAutoPaging(
			ctx,
			cmd.Value("objective-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "objectives:tasks list", iter, format, transform, maxItems)
	}
}
