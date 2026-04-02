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

var apiKeysCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Creates a new API key in the workspace.",
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
			Usage:    "APIKeySpec contains the API Key-specific fields",
			Required: true,
			BodyPath: "spec",
		},
	},
	Action:          handleAPIKeysCreate,
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
			Name:       "spec.token",
			Usage:      "The actual token value (only returned on creation and rotation, read-only)",
			InnerField: "token",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.description",
			Usage:      "Description of what this API Key is used for",
			InnerField: "description",
		},
	},
})

var apiKeysRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves an API key by ID from the workspace",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAPIKeysRetrieve,
	HideHelpCommand: true,
}

var apiKeysUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Updates an API key in the workspace",
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
			Usage:    "APIKeySpec contains the API Key-specific fields",
			BodyPath: "spec",
		},
		&requestflag.Flag[string]{
			Name:     "update-mask",
			Usage:    "Fields to update",
			BodyPath: "updateMask",
		},
	},
	Action:          handleAPIKeysUpdate,
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
			Name:       "spec.token",
			Usage:      "The actual token value (only returned on creation and rotation, read-only)",
			InnerField: "token",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.description",
			Usage:      "Description of what this API Key is used for",
			InnerField: "description",
		},
	},
})

var apiKeysList = cli.Command{
	Name:    "list",
	Usage:   "Lists all API keys in the workspace",
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
	Action:          handleAPIKeysList,
	HideHelpCommand: true,
}

var apiKeysDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes an API key from the workspace",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAPIKeysDelete,
	HideHelpCommand: true,
}

var apiKeysRotate = cli.Command{
	Name:    "rotate",
	Usage:   "Rotates an API Key and returns a new token. All previous API Key tokens in use\nwill be invalidated.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleAPIKeysRotate,
	HideHelpCommand: true,
}

func handleAPIKeysCreate(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cadenya.APIKeyNewParams{}

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
	_, err = client.APIKeys.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "api-keys create", obj, format, transform)
}

func handleAPIKeysRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.APIKeys.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "api-keys retrieve", obj, format, transform)
}

func handleAPIKeysUpdate(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cadenya.APIKeyUpdateParams{}

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
	_, err = client.APIKeys.Update(
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
	return ShowJSON(os.Stdout, "api-keys update", obj, format, transform)
}

func handleAPIKeysList(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cadenya.APIKeyListParams{}

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
		_, err = client.APIKeys.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "api-keys list", obj, format, transform)
	} else {
		iter := client.APIKeys.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "api-keys list", iter, format, transform, maxItems)
	}
}

func handleAPIKeysDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.APIKeys.Delete(ctx, cmd.Value("id").(string), options...)
}

func handleAPIKeysRotate(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.APIKeys.Rotate(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "api-keys rotate", obj, format, transform)
}
