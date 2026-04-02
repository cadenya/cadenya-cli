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

var agentsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new agent in the workspace",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "CreateResourceMetadata contains the user-provided fields for creating\n a workspace-scoped resource. Read-only fields (id, account_id, workspace_id, profile_id,\n created_at) are excluded since they are set by the server.",
			Required: true,
			BodyPath: "metadata",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "spec",
			Usage:    "Agent specification (user-provided configuration)",
			Required: true,
			BodyPath: "spec",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "default-variation",
			Usage:    "Create agent variation request",
			BodyPath: "defaultVariation",
		},
	},
	Action:          handleAgentsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"metadata": {
		&requestflag.InnerFlag[string]{
			Name:       "metadata.name",
			Usage:      `Human-readable name for the resource (e.g., "Customer Support Agent", "Email Tool")`,
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "metadata.external-id",
			Usage:      "External ID for the resource (e.g., a workflow ID from an external system)",
			InnerField: "externalId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "metadata.labels",
			Usage:      "Arbitrary key-value pairs for categorization and filtering\n Examples: {\"environment\": \"production\", \"team\": \"platform\", \"version\": \"v2\"}",
			InnerField: "labels",
		},
	},
	"spec": {
		&requestflag.InnerFlag[string]{
			Name:       "spec.status",
			Usage:      "Status of the agent",
			InnerField: "status",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.variation-selection-mode",
			Usage:      "Controls how variations are automatically selected when creating objectives\n Defaults to RANDOM when unspecified",
			InnerField: "variationSelectionMode",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.description",
			Usage:      "Description of the agent's purpose",
			InnerField: "description",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.webhook-events-hmac-secret",
			Usage:      "The generated secret that will sign all webhooks that are sent to your configured Webhook URL.\n Formatted as \"wh_asdf1234\" per the https://www.standardwebhooks.com/ format.",
			InnerField: "webhookEventsHmacSecret",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.webhook-events-url",
			Usage:      "The URL that Cadenya will send events for any objective assigned to the agent.",
			InnerField: "webhookEventsUrl",
		},
	},
	"default-variation": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "default-variation.metadata",
			Usage:      "CreateResourceMetadata contains the user-provided fields for creating\n a workspace-scoped resource. Read-only fields (id, account_id, workspace_id, profile_id,\n created_at) are excluded since they are set by the server.",
			InnerField: "metadata",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "default-variation.spec",
			Usage:      "AgentVariationSpec defines the operational configuration for a variation",
			InnerField: "spec",
		},
		&requestflag.InnerFlag[string]{
			Name:       "default-variation.agent-id",
			Usage:      "Agent ID (from path)",
			InnerField: "agentId",
		},
	},
})

var agentsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves an agent by ID from the workspace",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAgentsRetrieve,
	HideHelpCommand: true,
}

var agentsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates an agent in the workspace",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "UpdateResourceMetadata contains the user-provided fields for updating\n a workspace-scoped resource. Read-only fields (id, account_id, workspace_id, profile_id,\n created_at) are excluded since they are set by the server.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "spec",
			Usage:    "Agent specification (user-provided configuration)",
			BodyPath: "spec",
		},
		&requestflag.Flag[string]{
			Name:     "update-mask",
			Usage:    "Fields to update",
			BodyPath: "updateMask",
		},
	},
	Action:          handleAgentsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"metadata": {
		&requestflag.InnerFlag[string]{
			Name:       "metadata.name",
			Usage:      `Human-readable name for the resource (e.g., "Customer Support Agent", "Email Tool")`,
			InnerField: "name",
		},
		&requestflag.InnerFlag[string]{
			Name:       "metadata.external-id",
			Usage:      "External ID for the resource (e.g., a workflow ID from an external system)",
			InnerField: "externalId",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "metadata.labels",
			Usage:      "Arbitrary key-value pairs for categorization and filtering\n Examples: {\"environment\": \"production\", \"team\": \"platform\", \"version\": \"v2\"}",
			InnerField: "labels",
		},
	},
	"spec": {
		&requestflag.InnerFlag[string]{
			Name:       "spec.status",
			Usage:      "Status of the agent",
			InnerField: "status",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.variation-selection-mode",
			Usage:      "Controls how variations are automatically selected when creating objectives\n Defaults to RANDOM when unspecified",
			InnerField: "variationSelectionMode",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.description",
			Usage:      "Description of the agent's purpose",
			InnerField: "description",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.webhook-events-hmac-secret",
			Usage:      "The generated secret that will sign all webhooks that are sent to your configured Webhook URL.\n Formatted as \"wh_asdf1234\" per the https://www.standardwebhooks.com/ format.",
			InnerField: "webhookEventsHmacSecret",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.webhook-events-url",
			Usage:      "The URL that Cadenya will send events for any objective assigned to the agent.",
			InnerField: "webhookEventsUrl",
		},
	},
})

var agentsList = cli.Command{
	Name:    "list",
	Usage:   "Lists all agents in the workspace",
	Suggest: true,
	Flags: []cli.Flag{
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
			Name:      "prefix",
			Usage:     "Filter expression (query param: prefix)",
			QueryPath: "prefix",
		},
		&requestflag.Flag[string]{
			Name:      "sort-order",
			Usage:     "Sort order for results (asc or desc by creation time)",
			QueryPath: "sortOrder",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAgentsList,
	HideHelpCommand: true,
}

var agentsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes an agent from the workspace",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAgentsDelete,
	HideHelpCommand: true,
}

func handleAgentsCreate(ctx context.Context, cmd *cli.Command) error {
	client := gocadenyacomcadenyasdkgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := gocadenyacomcadenyasdkgo.AgentNewParams{}

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
	_, err = client.Agents.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agents create", obj, format, transform)
}

func handleAgentsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := gocadenyacomcadenyasdkgo.NewClient(getDefaultRequestOptions(cmd)...)
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
	_, err = client.Agents.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agents retrieve", obj, format, transform)
}

func handleAgentsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := gocadenyacomcadenyasdkgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := gocadenyacomcadenyasdkgo.AgentUpdateParams{}

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
	_, err = client.Agents.Update(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agents update", obj, format, transform)
}

func handleAgentsList(ctx context.Context, cmd *cli.Command) error {
	client := gocadenyacomcadenyasdkgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := gocadenyacomcadenyasdkgo.AgentListParams{}

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
		_, err = client.Agents.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "agents list", obj, format, transform)
	} else {
		iter := client.Agents.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "agents list", iter, format, transform, maxItems)
	}
}

func handleAgentsDelete(ctx context.Context, cmd *cli.Command) error {
	client := gocadenyacomcadenyasdkgo.NewClient(getDefaultRequestOptions(cmd)...)
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

	return client.Agents.Delete(ctx, cmd.Value("id").(string), options...)
}
