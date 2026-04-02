// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/cadenya/cadenya-cli/internal/apiquery"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
	"github.com/cadenya/cadenya-sdk-go"
	"github.com/cadenya/cadenya-sdk-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var objectivesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new objective in the workspace",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-id",
			Required: true,
			BodyPath: "agentId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "data",
			Required: true,
			BodyPath: "data",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "CreateOperationMetadata contains the user-provided fields for creating\n an operation. Read-only fields (id, account_id, workspace_id, created_at, profile_id)\n are excluded since they are set by the server.",
			Required: true,
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "variation-id",
			Usage:    "Optional explicit variation selection. Overrides the agent's variation_selection_mode.",
			BodyPath: "variationId",
		},
	},
	Action:          handleObjectivesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"data": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "data.agent",
			Usage:      "Agent resource",
			InnerField: "agent",
		},
		&requestflag.InnerFlag[any]{
			Name:       "data.data",
			Usage:      "Represents a dynamically typed value which can be either null, a number, a string, a boolean, a recursive struct value, or a list of values.",
			InnerField: "data",
		},
		&requestflag.InnerFlag[string]{
			Name:       "data.initial-message",
			Usage:      "The initial message sent to the agent. This becomes the first user message in the LLM chat history.",
			InnerField: "initialMessage",
		},
		&requestflag.InnerFlag[string]{
			Name:       "data.parent-objective-id",
			Usage:      "A parent objective means the objective was spawned off using a separate agent to complete an objective",
			InnerField: "parentObjectiveId",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "data.secrets",
			Usage:      "Secrets that can be used in the headers for tool calls using the secret interpolation format.",
			InnerField: "secrets",
		},
		&requestflag.InnerFlag[string]{
			Name:       "data.system-prompt",
			Usage:      "system_prompt is read-only, derived from the selected variation's prompt",
			InnerField: "systemPrompt",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "data.variation",
			Usage:      "AgentVariation resource",
			InnerField: "variation",
		},
	},
	"metadata": {
		&requestflag.InnerFlag[string]{
			Name:       "metadata.external-id",
			Usage:      "External ID for the operation (e.g., a workflow ID from an external system)",
			InnerField: "externalId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "metadata.labels",
			Usage:      "Arbitrary key-value pairs for categorization and filtering\n Examples: {\"priority\": \"high\", \"source\": \"api\", \"workflow\": \"onboarding\"}",
			InnerField: "labels",
		},
	},
})

var objectivesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves an objective by ID from the workspace",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleObjectivesRetrieve,
	HideHelpCommand: true,
}

var objectivesList = cli.Command{
	Name:    "list",
	Usage:   "Lists all objectives in the workspace",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Usage:     "Agent ID for filtering",
			QueryPath: "agentId",
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
			Name:      "parent-objective-id",
			Usage:     "Optional filters",
			QueryPath: "parentObjectiveId",
		},
		&requestflag.Flag[string]{
			Name:      "profile-id",
			QueryPath: "profileId",
		},
		&requestflag.Flag[string]{
			Name:      "sort-order",
			Usage:     "Sort order for results (asc or desc by creation time)",
			QueryPath: "sortOrder",
		},
		&requestflag.Flag[string]{
			Name:      "state",
			Usage:     "Filter by state",
			QueryPath: "state",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleObjectivesList,
	HideHelpCommand: true,
}

var objectivesCancel = cli.Command{
	Name:    "cancel",
	Usage:   "Cancels a running or pending objective. The objective's state will be set to\nSTATE_CANCELLED.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "objective-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Usage:    "Optional reason for cancellation",
			BodyPath: "reason",
		},
	},
	Action:          handleObjectivesCancel,
	HideHelpCommand: true,
}

var objectivesContinue = requestflag.WithInnerFlags(cli.Command{
	Name:    "continue",
	Usage:   "Continues an objective that has completed",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "objective-id",
			Required: true,
		},
		&requestflag.Flag[bool]{
			Name:     "enqueue",
			Usage:    "When set to true, the message will be enqueued for when the agent loop is available to process it.",
			BodyPath: "enqueue",
		},
		&requestflag.Flag[string]{
			Name:     "message",
			Usage:    "The message to continue an objective that has completed (or you are enqueing)",
			BodyPath: "message",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "secret",
			Usage:    "Secrets that should be included with the message. Helpful for when you need to update secrets on the objective (IE: A secret expires and needs to be refreshed)",
			BodyPath: "secrets",
		},
	},
	Action:          handleObjectivesContinue,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"secret": {
		&requestflag.InnerFlag[string]{
			Name:       "secret.name",
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "secret.value",
			InnerField: "value",
		},
	},
})

var objectivesListContextWindows = cli.Command{
	Name:    "list-context-windows",
	Usage:   "Read-only list of the last five windows of execution for this objective, ordered\nby most recent first",
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleObjectivesListContextWindows,
	HideHelpCommand: true,
}

var objectivesListEvents = cli.Command{
	Name:    "list-events",
	Usage:   "Lists all events for an objective",
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
			Name:      "sort-order",
			Usage:     "Sort order for results (asc or desc by creation time)",
			QueryPath: "sortOrder",
		},
		&requestflag.Flag[string]{
			Name:      "window-id",
			Usage:     "Optional context window ID to filter events by",
			QueryPath: "windowId",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleObjectivesListEvents,
	HideHelpCommand: true,
}

func handleObjectivesCreate(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cadenya.ObjectiveNewParams{}

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
	_, err = client.Objectives.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "objectives create", obj, format, transform)
}

func handleObjectivesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
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
	_, err = client.Objectives.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "objectives retrieve", obj, format, transform)
}

func handleObjectivesList(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cadenya.ObjectiveListParams{}

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
		_, err = client.Objectives.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "objectives list", obj, format, transform)
	} else {
		iter := client.Objectives.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "objectives list", iter, format, transform, maxItems)
	}
}

func handleObjectivesCancel(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cadenya.ObjectiveCancelParams{}

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
	_, err = client.Objectives.Cancel(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "objectives cancel", obj, format, transform)
}

func handleObjectivesContinue(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cadenya.ObjectiveContinueParams{}

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
	_, err = client.Objectives.Continue(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "objectives continue", obj, format, transform)
}

func handleObjectivesListContextWindows(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cadenya.ObjectiveListContextWindowsParams{}

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
		_, err = client.Objectives.ListContextWindows(
			ctx,
			cmd.Value("objective-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "objectives list-context-windows", obj, format, transform)
	} else {
		iter := client.Objectives.ListContextWindowsAutoPaging(
			ctx,
			cmd.Value("objective-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "objectives list-context-windows", iter, format, transform, maxItems)
	}
}

func handleObjectivesListEvents(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("objective-id") && len(unusedArgs) > 0 {
		cmd.Set("objective-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cadenya.ObjectiveListEventsParams{}

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
		_, err = client.Objectives.ListEvents(
			ctx,
			cmd.Value("objective-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "objectives list-events", obj, format, transform)
	} else {
		iter := client.Objectives.ListEventsAutoPaging(
			ctx,
			cmd.Value("objective-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "objectives list-events", iter, format, transform, maxItems)
	}
}
