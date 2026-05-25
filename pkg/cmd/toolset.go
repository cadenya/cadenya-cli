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

var toolSetsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new tool set in the workspace",
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
			Required: true,
			BodyPath: "spec",
		},
	},
	Action:          handleToolSetsCreate,
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
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.adapter",
			InnerField: "adapter",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.description",
			InnerField: "description",
		},
	},
})

var toolSetsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a tool set by ID from the workspace",
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
	Action:          handleToolSetsRetrieve,
	HideHelpCommand: true,
}

var toolSetsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates a tool set in the workspace",
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
			BodyPath: "spec",
		},
		&requestflag.Flag[string]{
			Name:     "update-mask",
			BodyPath: "updateMask",
		},
	},
	Action:          handleToolSetsUpdate,
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
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.adapter",
			InnerField: "adapter",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.description",
			InnerField: "description",
		},
	},
})

var toolSetsList = cli.Command{
	Name:    "list",
	Usage:   "Lists all tool sets in the workspace",
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
			Name:      "query",
			Usage:     "Free-form search query",
			QueryPath: "query",
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
	Action:          handleToolSetsList,
	HideHelpCommand: true,
}

var toolSetsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a tool set in the workspace",
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
	Action:          handleToolSetsDelete,
	HideHelpCommand: true,
}

var toolSetsGetOpenAPISpec = cli.Command{
	Name:    "get-openapi-spec",
	Usage:   "Retrieves the current OpenAPI specification JSON that has been consumed by the\ntool set. Only applicable to tool sets using the OpenAPI adapter.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "tool-set-id",
			Required:  true,
			PathParam: "toolSetId",
		},
	},
	Action:          handleToolSetsGetOpenAPISpec,
	HideHelpCommand: true,
}

var toolSetsListEvents = cli.Command{
	Name:    "list-events",
	Usage:   "Lists all events (including sync status) for a tool set",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "tool-set-id",
			Required:  true,
			PathParam: "toolSetId",
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
	Action:          handleToolSetsListEvents,
	HideHelpCommand: true,
}

func handleToolSetsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := cadenya.ToolSetNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ToolSets.New(
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
		Title:          "tool-sets create",
		Transform:      transform,
	})
}

func handleToolSetsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.ToolSets.Get(
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
		Title:          "tool-sets retrieve",
		Transform:      transform,
	})
}

func handleToolSetsUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := cadenya.ToolSetUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.ToolSets.Update(
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
		Title:          "tool-sets update",
		Transform:      transform,
	})
}

func handleToolSetsList(ctx context.Context, cmd *cli.Command) error {
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

	params := cadenya.ToolSetListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.ToolSets.List(
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
			Title:          "tool-sets list",
			Transform:      transform,
		})
	} else {
		iter := client.ToolSets.ListAutoPaging(
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
			Title:          "tool-sets list",
			Transform:      transform,
		})
	}
}

func handleToolSetsDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.ToolSets.Delete(
		ctx,
		cmd.Value("workspace-id").(string),
		cmd.Value("id").(string),
		options...,
	)
}

func handleToolSetsGetOpenAPISpec(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workspace-id") && len(unusedArgs) > 0 {
		cmd.Set("workspace-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("tool-set-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-set-id", unusedArgs[0])
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
	_, err = client.ToolSets.GetOpenAPISpec(
		ctx,
		cmd.Value("workspace-id").(string),
		cmd.Value("tool-set-id").(string),
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
		Title:          "tool-sets get-openapi-spec",
		Transform:      transform,
	})
}

func handleToolSetsListEvents(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workspace-id") && len(unusedArgs) > 0 {
		cmd.Set("workspace-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("tool-set-id") && len(unusedArgs) > 0 {
		cmd.Set("tool-set-id", unusedArgs[0])
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

	params := cadenya.ToolSetListEventsParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.ToolSets.ListEvents(
			ctx,
			cmd.Value("workspace-id").(string),
			cmd.Value("tool-set-id").(string),
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
			Title:          "tool-sets list-events",
			Transform:      transform,
		})
	} else {
		iter := client.ToolSets.ListEventsAutoPaging(
			ctx,
			cmd.Value("workspace-id").(string),
			cmd.Value("tool-set-id").(string),
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
			Title:          "tool-sets list-events",
			Transform:      transform,
		})
	}
}
