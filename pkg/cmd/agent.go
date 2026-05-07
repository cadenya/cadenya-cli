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

var agentsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new agent in the workspace",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
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
			Name:       "metadata.bundle-key",
			Usage:      "Optional bundle ownership key. See ResourceMetadata.bundle_key.",
			InnerField: "bundleKey",
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
		&requestflag.InnerFlag[any]{
			Name:       "spec.input-data-schema",
			Usage:      "InputDataSchema is used for enforcing a data input when objectives are created. This is valuable when using liquid formatting in agent variation\n prompts. Input data schema is also valuable when using an agent as a sub-agent, as the schema is used as the tool's input parameter schema. If omitted,\n the sub-agent schema will be loaded with a simple \"prompt\" free text string as its schema.",
			InnerField: "inputDataSchema",
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
			Usage:      "Agent ID. Accepts the canonical `agent_…` form or the `external_id:<value>` form.",
			InnerField: "agentId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "default-variation.workspace-id",
			Usage:      "Workspace ID.",
			InnerField: "workspaceId",
		},
	},
})

var agentsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves an agent by ID from the workspace",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
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
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
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
			Name:       "metadata.bundle-key",
			Usage:      "Optional bundle ownership key. See ResourceMetadata.bundle_key.",
			InnerField: "bundleKey",
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
		&requestflag.InnerFlag[any]{
			Name:       "spec.input-data-schema",
			Usage:      "InputDataSchema is used for enforcing a data input when objectives are created. This is valuable when using liquid formatting in agent variation\n prompts. Input data schema is also valuable when using an agent as a sub-agent, as the schema is used as the tool's input parameter schema. If omitted,\n the sub-agent schema will be loaded with a simple \"prompt\" free text string as its schema.",
			InnerField: "inputDataSchema",
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
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "bundle-key",
			Usage:     "Filter by bundle_key — return only resources owned by this bundle.",
			QueryPath: "bundleKey",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor from previous response",
			QueryPath: "cursor",
		},
		&requestflag.Flag[bool]{
			Name:      "include-info",
			Usage:     "When true, the `info` field on each returned agent is populated. Requests\n with this flag count more against your rate limit.",
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
			Name:      "query",
			Usage:     "Free-form search query",
			QueryPath: "query",
		},
		&requestflag.Flag[string]{
			Name:      "sort-order",
			Usage:     "Sort order for results (asc or desc by creation time)",
			QueryPath: "sortOrder",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by agent publication status",
			QueryPath: "status",
		},
		&requestflag.Flag[string]{
			Name:      "variation-selection-mode",
			Usage:     "Filter by variation selection mode",
			QueryPath: "variationSelectionMode",
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
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleAgentsDelete,
	HideHelpCommand: true,
}

func handleAgentsCreate(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workspace-id") && len(unusedArgs) > 0 {
		cmd.Set("workspace-id", unusedArgs[0])
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

	params := cadenya.AgentNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.New(
		ctx,
		cmd.Value("workspace-id").(string),
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
		Title:          "agents create",
		Transform:      transform,
	})
}

func handleAgentsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workspace-id") && len(unusedArgs) > 0 {
		cmd.Set("workspace-id", unusedArgs[0])
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
	_, err = client.Agents.Get(
		ctx,
		cmd.Value("workspace-id").(string),
		cmd.Value("id").(string),
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
		Title:          "agents retrieve",
		Transform:      transform,
	})
}

func handleAgentsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workspace-id") && len(unusedArgs) > 0 {
		cmd.Set("workspace-id", unusedArgs[0])
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := cadenya.AgentUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.Update(
		ctx,
		cmd.Value("workspace-id").(string),
		cmd.Value("id").(string),
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
		Title:          "agents update",
		Transform:      transform,
	})
}

func handleAgentsList(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workspace-id") && len(unusedArgs) > 0 {
		cmd.Set("workspace-id", unusedArgs[0])
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

	params := cadenya.AgentListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Agents.List(
			ctx,
			cmd.Value("workspace-id").(string),
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
			Title:          "agents list",
			Transform:      transform,
		})
	} else {
		iter := client.Agents.ListAutoPaging(
			ctx,
			cmd.Value("workspace-id").(string),
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
			Title:          "agents list",
			Transform:      transform,
		})
	}
}

func handleAgentsDelete(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workspace-id") && len(unusedArgs) > 0 {
		cmd.Set("workspace-id", unusedArgs[0])
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

	return client.Agents.Delete(
		ctx,
		cmd.Value("workspace-id").(string),
		cmd.Value("id").(string),
		options...,
	)
}
