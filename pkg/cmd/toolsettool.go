// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"go.cadenya.com/cadenya-go"
	"go.cadenya.com/cadenya-go/option"
	"os"

	"github.com/cadenya/cadenya-cli/internal/apiquery"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var toolSetsToolsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new tool in the tool set",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "tool-set-id",
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
			Required: true,
			BodyPath: "spec",
		},
	},
	Action:          handleToolSetsToolsCreate,
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
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.config",
			Usage:      "Config defines the adapter to use for the tool.\n This is used to determine how the tool is called.\n For example, if the tool is an HTTP tool, the adapter will be Http.\n If the tool is an inline tool, the adapter will be Inline.",
			InnerField: "config",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.parameters",
			InnerField: "parameters",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.status",
			Usage:      `Allowed values: "TOOL_STATUS_UNSPECIFIED", "TOOL_STATUS_AVAILABLE", "TOOL_STATUS_FILTERED", "TOOL_STATUS_ARCHIVED".`,
			InnerField: "status",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "spec.requires-approval",
			InnerField: "requiresApproval",
		},
	},
})

var toolSetsToolsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a tool by ID from the workspace",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "tool-set-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleToolSetsToolsRetrieve,
	HideHelpCommand: true,
}

var toolSetsToolsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates a tool in the tool set",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "tool-set-id",
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
			BodyPath: "spec",
		},
		&requestflag.Flag[string]{
			Name:     "update-mask",
			BodyPath: "updateMask",
		},
	},
	Action:          handleToolSetsToolsUpdate,
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
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.config",
			Usage:      "Config defines the adapter to use for the tool.\n This is used to determine how the tool is called.\n For example, if the tool is an HTTP tool, the adapter will be Http.\n If the tool is an inline tool, the adapter will be Inline.",
			InnerField: "config",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.parameters",
			InnerField: "parameters",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.status",
			Usage:      `Allowed values: "TOOL_STATUS_UNSPECIFIED", "TOOL_STATUS_AVAILABLE", "TOOL_STATUS_FILTERED", "TOOL_STATUS_ARCHIVED".`,
			InnerField: "status",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "spec.requires-approval",
			InnerField: "requiresApproval",
		},
	},
})

var toolSetsToolsList = cli.Command{
	Name:    "list",
	Usage:   "Lists all tools in the tool set",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "tool-set-id",
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
	Action:          handleToolSetsToolsList,
	HideHelpCommand: true,
}

var toolSetsToolsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a tool in the tool set",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "tool-set-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleToolSetsToolsDelete,
	HideHelpCommand: true,
}

func handleToolSetsToolsCreate(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tool-set-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-set-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cadenya.ToolSetToolNewParams{}

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
	_, err = client.ToolSets.Tools.New(
		ctx,
		cmd.Value("tool-set-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "tool-sets:tools create", obj, format, transform)
}

func handleToolSetsToolsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tool-set-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-set-id", unusedArgs[0])
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
	_, err = client.ToolSets.Tools.Get(
		ctx,
		cmd.Value("tool-set-id").(string),
		cmd.Value("id").(string),
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "tool-sets:tools retrieve", obj, format, transform)
}

func handleToolSetsToolsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tool-set-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-set-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cadenya.ToolSetToolUpdateParams{}

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
	_, err = client.ToolSets.Tools.Update(
		ctx,
		cmd.Value("tool-set-id").(string),
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
	return ShowJSON(os.Stdout, "tool-sets:tools update", obj, format, transform)
}

func handleToolSetsToolsList(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tool-set-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-set-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cadenya.ToolSetToolListParams{}

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
		_, err = client.ToolSets.Tools.List(
			ctx,
			cmd.Value("tool-set-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "tool-sets:tools list", obj, format, transform)
	} else {
		iter := client.ToolSets.Tools.ListAutoPaging(
			ctx,
			cmd.Value("tool-set-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "tool-sets:tools list", iter, format, transform, maxItems)
	}
}

func handleToolSetsToolsDelete(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("tool-set-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-set-id", unusedArgs[0])
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

	return client.ToolSets.Tools.Delete(
		ctx,
		cmd.Value("tool-set-id").(string),
		cmd.Value("id").(string),
		options...,
	)
}
