// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"go.cadenya.com/cadenya-sdk-go"
	"go.cadenya.com/cadenya-sdk-go/option"
	"os"

	"github.com/stainless-sdks/cadenya-cli/internal/apiquery"
	"github.com/stainless-sdks/cadenya-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var agentsVariationsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new variation for an agent",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-id",
			Required: true,
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
			Usage:      "Arbitrary key-value pairs for categorization and filtering\n Examples: {\"environment\": \"production\", \"team\": \"platform\", \"version\": \"v2\"}",
			InnerField: "labels",
		},
	},
	"spec": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "spec.agent-documents",
			Usage:      "Documents assigned to this variation.\n Can include individual documents or entire document namespaces (which include all documents in the namespace).",
			InnerField: "agentDocuments",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "spec.agent-tools",
			Usage:      "Tools assigned to this variation",
			InnerField: "agentTools",
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
		&requestflag.InnerFlag[bool]{
			Name:       "spec.enable-episodic-memory",
			Usage:      "Enable episodic memory for objectives using this variation.\n When true, the system automatically creates a document namespace for each objective\n using the objective's episodic_key as the external_id, allowing the agent to\n store and retrieve documents specific to that episode.",
			InnerField: "enableEpisodicMemory",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "spec.episodic-memory-ttl",
			Usage:      "How long episodic memories should be retained.\n After this duration, episodic document namespaces can be automatically cleaned up.\n If not set, episodic memories are retained indefinitely.",
			InnerField: "episodicMemoryTtl",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.model-config",
			Usage:      "ModelConfig defines the model configuration for a variation",
			InnerField: "modelConfig",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.prompt",
			Usage:      "The system prompt for this variation",
			InnerField: "prompt",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.tool-selection",
			InnerField: "toolSelection",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "spec.weight",
			Usage:      "Weight for weighted random selection (>= 0). P(v) = v.weight / sum(all_weights).\n Only used when the agent's variation_selection_mode is WEIGHTED. A weight of 0 means never auto-selected, but can still be chosen explicitly via variation_id on CreateObjectiveRequest.",
			InnerField: "weight",
		},
	},
})

var agentsVariationsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a variation by ID from an agent",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
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
			Name:     "agent-id",
			Required: true,
		},
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
			Usage:      "Arbitrary key-value pairs for categorization and filtering\n Examples: {\"environment\": \"production\", \"team\": \"platform\", \"version\": \"v2\"}",
			InnerField: "labels",
		},
	},
	"spec": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "spec.agent-documents",
			Usage:      "Documents assigned to this variation.\n Can include individual documents or entire document namespaces (which include all documents in the namespace).",
			InnerField: "agentDocuments",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "spec.agent-tools",
			Usage:      "Tools assigned to this variation",
			InnerField: "agentTools",
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
		&requestflag.InnerFlag[bool]{
			Name:       "spec.enable-episodic-memory",
			Usage:      "Enable episodic memory for objectives using this variation.\n When true, the system automatically creates a document namespace for each objective\n using the objective's episodic_key as the external_id, allowing the agent to\n store and retrieve documents specific to that episode.",
			InnerField: "enableEpisodicMemory",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "spec.episodic-memory-ttl",
			Usage:      "How long episodic memories should be retained.\n After this duration, episodic document namespaces can be automatically cleaned up.\n If not set, episodic memories are retained indefinitely.",
			InnerField: "episodicMemoryTtl",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.model-config",
			Usage:      "ModelConfig defines the model configuration for a variation",
			InnerField: "modelConfig",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.prompt",
			Usage:      "The system prompt for this variation",
			InnerField: "prompt",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.tool-selection",
			InnerField: "toolSelection",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "spec.weight",
			Usage:      "Weight for weighted random selection (>= 0). P(v) = v.weight / sum(all_weights).\n Only used when the agent's variation_selection_mode is WEIGHTED. A weight of 0 means never auto-selected, but can still be chosen explicitly via variation_id on CreateObjectiveRequest.",
			InnerField: "weight",
		},
	},
})

var agentsVariationsList = cli.Command{
	Name:    "list",
	Usage:   "Lists all variations for an agent",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-id",
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
			Name:     "agent-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAgentsVariationsDelete,
	HideHelpCommand: true,
}

func handleAgentsVariationsCreate(ctx context.Context, cmd *cli.Command) error {
	client := gocadenyacomcadenyasdkgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := gocadenyacomcadenyasdkgo.AgentVariationNewParams{}

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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agents:variations create", obj, format, transform)
}

func handleAgentsVariationsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := gocadenyacomcadenyasdkgo.NewClient(getDefaultRequestOptions(cmd)...)
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Agents.Variations.Get(
		ctx,
		cmd.Value("agent-id").(string),
		cmd.Value("id").(string),
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agents:variations retrieve", obj, format, transform)
}

func handleAgentsVariationsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := gocadenyacomcadenyasdkgo.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := gocadenyacomcadenyasdkgo.AgentVariationUpdateParams{}

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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "agents:variations update", obj, format, transform)
}

func handleAgentsVariationsList(ctx context.Context, cmd *cli.Command) error {
	client := gocadenyacomcadenyasdkgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := gocadenyacomcadenyasdkgo.AgentVariationListParams{}

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
		return ShowJSON(os.Stdout, "agents:variations list", obj, format, transform)
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
		return ShowJSONIterator(os.Stdout, "agents:variations list", iter, format, transform, maxItems)
	}
}

func handleAgentsVariationsDelete(ctx context.Context, cmd *cli.Command) error {
	client := gocadenyacomcadenyasdkgo.NewClient(getDefaultRequestOptions(cmd)...)
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

	return client.Agents.Variations.Delete(
		ctx,
		cmd.Value("agent-id").(string),
		cmd.Value("id").(string),
		options...,
	)
}
