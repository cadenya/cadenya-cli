// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/cadenya/cadenya-cli/internal/apiquery"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
	"github.com/cadenya/cadenya-go"
	"github.com/cadenya/cadenya-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var objectivesFeedbackCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Submits feedback for an objective's execution. Feedback scores are used by the\nagent variation scoring system to evaluate and rank variation performance.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "objective-id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "data",
			Required: true,
			BodyPath: "data",
		},
	},
	Action:          handleObjectivesFeedbackCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"data": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "data.attributes",
			Usage:      "Arbitrary key-value pairs to identify the source of the feedback.\n Since the submitting profile is typically an API key, use this to pass through\n application-specific identifiers (e.g., {\"user_id\": \"usr_123\", \"session_id\": \"abc\"}).",
			InnerField: "attributes",
		},
		&requestflag.InnerFlag[string]{
			Name:       "data.comment",
			Usage:      "Optional human-readable comment explaining the feedback",
			InnerField: "comment",
		},
		&requestflag.InnerFlag[float64]{
			Name:       "data.score",
			Usage:      "A score between -1.0 and 1.0 representing the quality of the objective's execution.\n -1.0 is the worst possible score, 0.0 is neutral, and 1.0 is the best.",
			InnerField: "score",
		},
	},
})

var objectivesFeedbackList = cli.Command{
	Name:    "list",
	Usage:   "Lists all feedback submitted for an objective",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleObjectivesFeedbackList,
	HideHelpCommand: true,
}

func handleObjectivesFeedbackCreate(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cadenya.ObjectiveFeedbackNewParams{}

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
	_, err = client.Objectives.Feedback.New(
		ctx,
		cmd.Value("objective-id").(string),
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
	return ShowJSON(os.Stdout, os.Stderr, "objectives:feedback create", obj, format, explicitFormat, transform)
}

func handleObjectivesFeedbackList(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cadenya.ObjectiveFeedbackListParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Objectives.Feedback.List(
			ctx,
			cmd.Value("objective-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, os.Stderr, "objectives:feedback list", obj, format, explicitFormat, transform)
	} else {
		iter := client.Objectives.Feedback.ListAutoPaging(
			ctx,
			cmd.Value("objective-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, os.Stderr, "objectives:feedback list", iter, format, explicitFormat, transform, maxItems)
	}
}
