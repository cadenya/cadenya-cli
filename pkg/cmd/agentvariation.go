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

var agentsVariationsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new variation for an agent",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "CreateResourceMetadata contains the user-provided fields for creating\n a workspace-scoped resource. Read-only fields (id, account_id, workspace_id, profile_id,\n created_at) are excluded since they are set by the server.",
			Required: true,
			BodyPath: "metadata",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "spec",
			Usage:    "AgentVariationSpec defines the operational configuration for a variation",
			Required: true,
			BodyPath: "spec",
		},
	},
	Action:          handleAgentsVariationsCreate,
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
			Usage:      "Key-value pairs for categorization and filtering. Values are 0-63\n alphanumeric characters with \"-\", \"_\", or \".\" allowed between; keys\n follow the same shape and additionally accept an optional DNS-subdomain\n prefix (e.g. \"cadenya.com/\") of at most 253 characters.\n Examples: {\"environment\": \"production\", \"team\": \"platform\", \"version\": \"v2\"}",
			InnerField: "labels",
		},
	},
	"spec": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.compaction-config",
			Usage:      "CompactionConfig defines how context window compaction behaves for objectives using this variation.",
			InnerField: "compactionConfig",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.constraints",
			InnerField: "constraints",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.description",
			Usage:      "Human-readable description of what this variation does or when it should be used",
			InnerField: "description",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.first-user-message-template",
			Usage:      "Liquid template for the first user message of objectives using this variation.\n Rendered with CreateObjectiveRequest.first_user_message_data into\n Objective.first_user_message, the first user message in the LLM chat history.\n CreateObjectiveRequest.first_user_message, when set, overrides the rendered\n result. If neither this template nor first_user_message is present, objective\n creation is rejected with InvalidArgument.",
			InnerField: "firstUserMessageTemplate",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.model-config",
			Usage:      "ModelConfig defines the model configuration for a variation",
			InnerField: "modelConfig",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.progressive-discovery",
			Usage:      "ProgressiveDiscovery is used to indicate that the agent should automatically discover tools that are not explicitly assigned to it.\n Max tools is the maximum number of tools that can be discovered per search.\n Hints are optional hints for tool search. These are used in conjunction with the context-aware tool search and can help select the best tools for the task.",
			InnerField: "progressiveDiscovery",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.system-prompt-template",
			Usage:      "Liquid template for the system prompt of objectives using this variation.\n Rendered with CreateObjectiveRequest.system_prompt_data into Objective.system_prompt.",
			InnerField: "systemPromptTemplate",
		},
	},
})

var agentsVariationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a variation by ID from an agent",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleAgentsVariationsRetrieve,
	HideHelpCommand: true,
}

var agentsVariationsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates a variation for an agent",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
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
			Usage:    "AgentVariationSpec defines the operational configuration for a variation",
			BodyPath: "spec",
		},
		&requestflag.Flag[string]{
			Name:     "update-mask",
			Usage:    "Fields to update",
			BodyPath: "updateMask",
		},
	},
	Action:          handleAgentsVariationsUpdate,
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
			Usage:      "Key-value pairs for categorization and filtering. Values are 0-63\n alphanumeric characters with \"-\", \"_\", or \".\" allowed between; keys\n follow the same shape and additionally accept an optional DNS-subdomain\n prefix (e.g. \"cadenya.com/\") of at most 253 characters.\n Examples: {\"environment\": \"production\", \"team\": \"platform\", \"version\": \"v2\"}",
			InnerField: "labels",
		},
	},
	"spec": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.compaction-config",
			Usage:      "CompactionConfig defines how context window compaction behaves for objectives using this variation.",
			InnerField: "compactionConfig",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.constraints",
			InnerField: "constraints",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.description",
			Usage:      "Human-readable description of what this variation does or when it should be used",
			InnerField: "description",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.first-user-message-template",
			Usage:      "Liquid template for the first user message of objectives using this variation.\n Rendered with CreateObjectiveRequest.first_user_message_data into\n Objective.first_user_message, the first user message in the LLM chat history.\n CreateObjectiveRequest.first_user_message, when set, overrides the rendered\n result. If neither this template nor first_user_message is present, objective\n creation is rejected with InvalidArgument.",
			InnerField: "firstUserMessageTemplate",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.model-config",
			Usage:      "ModelConfig defines the model configuration for a variation",
			InnerField: "modelConfig",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.progressive-discovery",
			Usage:      "ProgressiveDiscovery is used to indicate that the agent should automatically discover tools that are not explicitly assigned to it.\n Max tools is the maximum number of tools that can be discovered per search.\n Hints are optional hints for tool search. These are used in conjunction with the context-aware tool search and can help select the best tools for the task.",
			InnerField: "progressiveDiscovery",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.system-prompt-template",
			Usage:      "Liquid template for the system prompt of objectives using this variation.\n Rendered with CreateObjectiveRequest.system_prompt_data into Objective.system_prompt.",
			InnerField: "systemPromptTemplate",
		},
	},
})

var agentsVariationsList = cli.Command{
	Name:    "list",
	Usage:   "Lists all variations for an agent",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor from previous response",
			QueryPath: "cursor",
		},
		&requestflag.Flag[bool]{
			Name:      "include-info",
			Usage:     "When true, the `info` field on each returned variation is populated.\n Requests with this flag count more against your rate limit.",
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
			Name:      "sort-order",
			Usage:     "Sort order for results (asc or desc by creation time)",
			QueryPath: "sortOrder",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAgentsVariationsList,
	HideHelpCommand: true,
}

var agentsVariationsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a variation from an agent",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleAgentsVariationsDelete,
	HideHelpCommand: true,
}

var agentsVariationsAddAssignment = cli.Command{
	Name:    "add-assignment",
	Usage:   "Assigns a tool, tool set, or sub-agent to a variation. Exactly one target ID\nmust be set.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
		&requestflag.Flag[string]{
			Name:      "variation-id",
			Required:  true,
			PathParam: "variationId",
		},
		&requestflag.Flag[string]{
			Name:     "tool-id",
			BodyPath: "toolId",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Allowed values: "toolId".`,
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "tool-set-id",
			BodyPath: "toolSetId",
		},
		&requestflag.Flag[string]{
			Name:     "sub-agent-id",
			BodyPath: "subAgentId",
		},
	},
	Action:          handleAgentsVariationsAddAssignment,
	HideHelpCommand: true,
}

var agentsVariationsAddMemoryLayer = cli.Command{
	Name:    "add-memory-layer",
	Usage:   "Attaches a memory layer to a variation at a given position in the variation's\nbaseline memory cascade.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
		&requestflag.Flag[string]{
			Name:      "variation-id",
			Required:  true,
			PathParam: "variationId",
		},
		&requestflag.Flag[string]{
			Name:     "memory-layer-id",
			Usage:    "Layer to attach. Accepts the canonical `memlyr_…` form or the `external_id:<value>` form.",
			Required: true,
			BodyPath: "memoryLayerId",
		},
		&requestflag.Flag[int64]{
			Name:     "position",
			Usage:    "Position in the baseline cascade (lower = more specific). If\n omitted, the server appends at the most general end (max existing\n position + 1).",
			BodyPath: "position",
		},
	},
	Action:          handleAgentsVariationsAddMemoryLayer,
	HideHelpCommand: true,
}

var agentsVariationsRemoveAssignment = cli.Command{
	Name:    "remove-assignment",
	Usage:   "Detaches an assignment from a variation, identified by the assignment ID\nreturned when it was added.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
		&requestflag.Flag[string]{
			Name:      "variation-id",
			Required:  true,
			PathParam: "variationId",
		},
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleAgentsVariationsRemoveAssignment,
	HideHelpCommand: true,
}

var agentsVariationsRemoveMemoryLayer = cli.Command{
	Name:    "remove-memory-layer",
	Usage:   "Detaches a memory layer assignment from a variation, identified by the\nassignment id.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
		&requestflag.Flag[string]{
			Name:      "variation-id",
			Required:  true,
			PathParam: "variationId",
		},
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleAgentsVariationsRemoveMemoryLayer,
	HideHelpCommand: true,
}

var agentsVariationsUpdateMemoryLayer = cli.Command{
	Name:    "update-memory-layer",
	Usage:   "Updates the position of a memory layer assignment on a variation.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "agent-id",
			Required:  true,
			PathParam: "agentId",
		},
		&requestflag.Flag[string]{
			Name:      "variation-id",
			Required:  true,
			PathParam: "variationId",
		},
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
		&requestflag.Flag[int64]{
			Name:     "position",
			Usage:    "New position. Only field currently updatable on an assignment.",
			BodyPath: "position",
		},
	},
	Action:          handleAgentsVariationsUpdateMemoryLayer,
	HideHelpCommand: true,
}

func handleAgentsVariationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
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

	params := cadenya.AgentVariationNewParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.Variations.New(
		ctx,
		cmd.Value("agent-id").(string),
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
		Title:          "agents:variations create",
		Transform:      transform,
	})
}

func handleAgentsVariationsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
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

	params := cadenya.AgentVariationGetParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.Variations.Get(
		ctx,
		cmd.Value("agent-id").(string),
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
		Title:          "agents:variations retrieve",
		Transform:      transform,
	})
}

func handleAgentsVariationsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
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

	params := cadenya.AgentVariationUpdateParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.Variations.Update(
		ctx,
		cmd.Value("agent-id").(string),
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
		Title:          "agents:variations update",
		Transform:      transform,
	})
}

func handleAgentsVariationsList(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
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

	params := cadenya.AgentVariationListParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Agents.Variations.List(
			ctx,
			cmd.Value("agent-id").(string),
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
			Title:          "agents:variations list",
			Transform:      transform,
		})
	} else {
		iter := client.Agents.Variations.ListAutoPaging(
			ctx,
			cmd.Value("agent-id").(string),
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
			Title:          "agents:variations list",
			Transform:      transform,
		})
	}
}

func handleAgentsVariationsDelete(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
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

	params := cadenya.AgentVariationDeleteParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	return client.Agents.Variations.Delete(
		ctx,
		cmd.Value("agent-id").(string),
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleAgentsVariationsAddAssignment(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("variation-id") && len(unusedArgs) > 0 {
		cmd.Set("variation-id", unusedArgs[0])
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

	params := cadenya.AgentVariationAddAssignmentParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.Variations.AddAssignment(
		ctx,
		cmd.Value("agent-id").(string),
		cmd.Value("variation-id").(string),
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
		Title:          "agents:variations add-assignment",
		Transform:      transform,
	})
}

func handleAgentsVariationsAddMemoryLayer(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("variation-id") && len(unusedArgs) > 0 {
		cmd.Set("variation-id", unusedArgs[0])
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

	params := cadenya.AgentVariationAddMemoryLayerParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.Variations.AddMemoryLayer(
		ctx,
		cmd.Value("agent-id").(string),
		cmd.Value("variation-id").(string),
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
		Title:          "agents:variations add-memory-layer",
		Transform:      transform,
	})
}

func handleAgentsVariationsRemoveAssignment(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("variation-id") && len(unusedArgs) > 0 {
		cmd.Set("variation-id", unusedArgs[0])
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

	params := cadenya.AgentVariationRemoveAssignmentParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	return client.Agents.Variations.RemoveAssignment(
		ctx,
		cmd.Value("agent-id").(string),
		cmd.Value("variation-id").(string),
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleAgentsVariationsRemoveMemoryLayer(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("variation-id") && len(unusedArgs) > 0 {
		cmd.Set("variation-id", unusedArgs[0])
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

	params := cadenya.AgentVariationRemoveMemoryLayerParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	return client.Agents.Variations.RemoveMemoryLayer(
		ctx,
		cmd.Value("agent-id").(string),
		cmd.Value("variation-id").(string),
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleAgentsVariationsUpdateMemoryLayer(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("variation-id") && len(unusedArgs) > 0 {
		cmd.Set("variation-id", unusedArgs[0])
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

	params := cadenya.AgentVariationUpdateMemoryLayerParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.Variations.UpdateMemoryLayer(
		ctx,
		cmd.Value("agent-id").(string),
		cmd.Value("variation-id").(string),
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
		Title:          "agents:variations update-memory-layer",
		Transform:      transform,
	})
}
