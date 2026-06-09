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

var bulkWorkspaceResourcesRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a bulk workspace apply operation by ID.",
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
	Action:          handleBulkWorkspaceResourcesRetrieve,
	HideHelpCommand: true,
}

var bulkWorkspaceResourcesList = cli.Command{
	Name:    "list",
	Usage:   "Lists past and in-flight bulk workspace apply operations in the workspace.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "bundle-key",
			Usage:     "Filter by bundle_key — list every apply for a given bundle.",
			QueryPath: "bundleKey",
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
			Usage:     "Sort order for results (asc or desc by creation time)",
			QueryPath: "sortOrder",
		},
		&requestflag.Flag[string]{
			Name:      "state",
			Usage:     "Filter by lifecycle state.",
			QueryPath: "state",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleBulkWorkspaceResourcesList,
	HideHelpCommand: true,
}

var bulkWorkspaceResourcesApply = requestflag.WithInnerFlags(cli.Command{
	Name:    "apply",
	Usage:   "Asynchronously applies a declarative bundle of workspace resources. Returns the\noperation immediately in PENDING; clients poll Get to track progress.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "data",
			Required: true,
			BodyPath: "data",
		},
	},
	Action:          handleBulkWorkspaceResourcesApply,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"data": {
		&requestflag.InnerFlag[string]{
			Name:       "data.bundle-key",
			Usage:      "Required. Bundle ownership key. Resources created or updated by an\n Apply have their `metadata.bundle_key` set to this value. On\n subsequent applies with the same bundle_key, resources currently\n bearing this bundle_key but absent from the spec are soft-deleted.",
			InnerField: "bundleKey",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "data.agents",
			Usage:      "Agents to upsert, keyed by external_id.",
			InnerField: "agents",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "data.automatically-publish-agents",
			Usage:      "When true, every agent created or updated by this Apply has its\n state forced to STATE_PUBLISHED, regardless of the state declared on\n the agent's entry. Useful when the bundle represents a production\n configuration and you want all of its agents live without setting\n state: STATE_PUBLISHED on each entry.\n\n Default false: each agent entry's `state` controls (which is\n STATE_DRAFT on create when unspecified).",
			InnerField: "automaticallyPublishAgents",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "data.memory-layers",
			Usage:      "Memory layers to upsert, keyed by external_id.",
			InnerField: "memoryLayers",
		},
		&requestflag.InnerFlag[string]{
			Name:       "data.source-url",
			Usage:      "Optional URL pointing to the source of this apply (GitHub PR,\n Jenkins build, GitLab pipeline, etc.). Surfaced in the dashboard so\n users can jump from an apply back to the change that produced it.\n Free-form HTTPS URI; not interpreted by the server.",
			InnerField: "sourceUrl",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "data.tool-sets",
			Usage:      "Tool sets to upsert, keyed by external_id.",
			InnerField: "toolSets",
		},
	},
})

func handleBulkWorkspaceResourcesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.BulkWorkspaceResources.Get(
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
		Title:          "bulk-workspace-resources retrieve",
		Transform:      transform,
	})
}

func handleBulkWorkspaceResourcesList(ctx context.Context, cmd *cli.Command) error {
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

	params := cadenya.BulkWorkspaceResourceListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.BulkWorkspaceResources.List(
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
			Title:          "bulk-workspace-resources list",
			Transform:      transform,
		})
	} else {
		iter := client.BulkWorkspaceResources.ListAutoPaging(
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
			Title:          "bulk-workspace-resources list",
			Transform:      transform,
		})
	}
}

func handleBulkWorkspaceResourcesApply(ctx context.Context, cmd *cli.Command) error {
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

	params := cadenya.BulkWorkspaceResourceApplyParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.BulkWorkspaceResources.Apply(
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
		Title:          "bulk-workspace-resources apply",
		Transform:      transform,
	})
}
