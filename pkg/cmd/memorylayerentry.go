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

var memoryLayersEntriesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new entry in a memory layer. Returns the detail view, including the\nresolved content body.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "memory-layer-id",
			Required:  true,
			PathParam: "memoryLayerId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "CreateResourceMetadata contains the user-provided fields for creating\n a workspace-scoped resource. Read-only fields (id, account_id, workspace_id, profile_id,\n created_at) are excluded since they are set by the server.",
			Required: true,
			BodyPath: "metadata",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "spec",
			Usage:    "MemoryEntryCreateSpec is the input shape for CreateMemoryEntry. It accepts\n either inline content or a reference to a completed Upload; exactly one of\n the two must be set.",
			Required: true,
			BodyPath: "spec",
		},
	},
	Action:          handleMemoryLayersEntriesCreate,
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
		&requestflag.InnerFlag[string]{
			Name:       "spec.key",
			Usage:      "See MemoryEntrySpec.key for the full rule set. Same constraints apply\n here.",
			InnerField: "key",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.content",
			Usage:      "Inline content, written directly into the entry.",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.upload-id",
			Usage:      "ID of a COMPLETE Upload. The server reads the object from storage,\n copies its bytes into the entry, and marks the upload consumed.",
			InnerField: "uploadId",
		},
	},
})

var memoryLayersEntriesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a memory entry by ID from a memory layer. Returns the detail view,\nincluding the content body.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "memory-layer-id",
			Required:  true,
			PathParam: "memoryLayerId",
		},
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleMemoryLayersEntriesRetrieve,
	HideHelpCommand: true,
}

var memoryLayersEntriesUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates a memory entry in a memory layer. Returns the detail view, including the\nresolved content body.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "memory-layer-id",
			Required:  true,
			PathParam: "memoryLayerId",
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
			Usage:    "MemoryEntryUpdateSpec is the input shape for UpdateMemoryEntry. Fields\n present in the request's update_mask are applied; unset fields are left\n alone. The source oneof is optional for updates — omit it to leave the\n body untouched, or set exactly one branch to replace it.",
			BodyPath: "spec",
		},
		&requestflag.Flag[string]{
			Name:     "update-mask",
			BodyPath: "updateMask",
		},
	},
	Action:          handleMemoryLayersEntriesUpdate,
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
		&requestflag.InnerFlag[string]{
			Name:       "spec.content",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.description",
			InnerField: "description",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.key",
			InnerField: "key",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.upload-id",
			InnerField: "uploadId",
		},
	},
})

var memoryLayersEntriesList = cli.Command{
	Name:    "list",
	Usage:   "Lists all entries in a memory layer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "memory-layer-id",
			Required:  true,
			PathParam: "memoryLayerId",
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
			Name:      "prefix",
			Usage:     "Filter by key prefix (e.g., \"skills/postmortem/\" to list all entries\n under that hierarchy). Matches against the entry's key, not its name.",
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
	Action:          handleMemoryLayersEntriesList,
	HideHelpCommand: true,
}

var memoryLayersEntriesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a memory entry from a memory layer",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "memory-layer-id",
			Required:  true,
			PathParam: "memoryLayerId",
		},
		&requestflag.Flag[string]{
			Name:      "id",
			Required:  true,
			PathParam: "id",
		},
	},
	Action:          handleMemoryLayersEntriesDelete,
	HideHelpCommand: true,
}

func handleMemoryLayersEntriesCreate(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workspace-id") && len(unusedArgs) > 0 {
		cmd.Set("workspace-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("memory-layer-id") && len(unusedArgs) > 0 {
		cmd.Set("memory-layer-id", unusedArgs[0])
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

	params := cadenya.MemoryLayerEntryNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.MemoryLayers.Entries.New(
		ctx,
		cmd.Value("workspace-id").(string),
		cmd.Value("memory-layer-id").(string),
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
		Title:          "memory-layers:entries create",
		Transform:      transform,
	})
}

func handleMemoryLayersEntriesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workspace-id") && len(unusedArgs) > 0 {
		cmd.Set("workspace-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("memory-layer-id") && len(unusedArgs) > 0 {
		cmd.Set("memory-layer-id", unusedArgs[0])
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
	_, err = client.MemoryLayers.Entries.Get(
		ctx,
		cmd.Value("workspace-id").(string),
		cmd.Value("memory-layer-id").(string),
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
		Title:          "memory-layers:entries retrieve",
		Transform:      transform,
	})
}

func handleMemoryLayersEntriesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workspace-id") && len(unusedArgs) > 0 {
		cmd.Set("workspace-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("memory-layer-id") && len(unusedArgs) > 0 {
		cmd.Set("memory-layer-id", unusedArgs[0])
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

	params := cadenya.MemoryLayerEntryUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.MemoryLayers.Entries.Update(
		ctx,
		cmd.Value("workspace-id").(string),
		cmd.Value("memory-layer-id").(string),
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
		Title:          "memory-layers:entries update",
		Transform:      transform,
	})
}

func handleMemoryLayersEntriesList(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workspace-id") && len(unusedArgs) > 0 {
		cmd.Set("workspace-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("memory-layer-id") && len(unusedArgs) > 0 {
		cmd.Set("memory-layer-id", unusedArgs[0])
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

	params := cadenya.MemoryLayerEntryListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.MemoryLayers.Entries.List(
			ctx,
			cmd.Value("workspace-id").(string),
			cmd.Value("memory-layer-id").(string),
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
			Title:          "memory-layers:entries list",
			Transform:      transform,
		})
	} else {
		iter := client.MemoryLayers.Entries.ListAutoPaging(
			ctx,
			cmd.Value("workspace-id").(string),
			cmd.Value("memory-layer-id").(string),
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
			Title:          "memory-layers:entries list",
			Transform:      transform,
		})
	}
}

func handleMemoryLayersEntriesDelete(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("workspace-id") && len(unusedArgs) > 0 {
		cmd.Set("workspace-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("memory-layer-id") && len(unusedArgs) > 0 {
		cmd.Set("memory-layer-id", unusedArgs[0])
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

	return client.MemoryLayers.Entries.Delete(
		ctx,
		cmd.Value("workspace-id").(string),
		cmd.Value("memory-layer-id").(string),
		cmd.Value("id").(string),
		options...,
	)
}
