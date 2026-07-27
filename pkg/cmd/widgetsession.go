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

var widgetSessionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Mints a session against a widget and returns the session bearer token\n(`spec.token`, returned only on creation) plus the authoritative widget hostname\n(`info.host`). Asserting a tenant upserts the tenant record; attached secrets\nflow to every conversation the session creates.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "spec",
			Usage:    "WidgetSessionSpec is the configuration of a session, fixed at mint.",
			Required: true,
			BodyPath: "spec",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "CreateOperationMetadata contains the user-provided fields for creating\n an operation. Read-only fields (id, account_id, workspace_id, created_at, profile_id)\n are excluded since they are set by the server.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "secret",
			Usage:    "Secrets to attach to the session.",
			BodyPath: "secrets",
		},
	},
	Action:          handleWidgetSessionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"spec": {
		&requestflag.InnerFlag[string]{
			Name:       "spec.widget-id",
			Usage:      "Widget this session is minted against. Accepts the canonical `wgt_…` form\n or the `external_id:<value>` form.",
			InnerField: "widgetId",
		},
		&requestflag.InnerFlag[string]{
			Name:       "spec.token",
			Usage:      "The session bearer token. Returned only on creation — subsequent reads\n omit it. The token is short-lived; the widget refreshes it at the widget\n host without involving the customer's backend.",
			InnerField: "token",
		},
		&requestflag.InnerFlag[any]{
			Name:       "spec.expires-at",
			Usage:      "Hard session expiry. Tokens never outlive it; after it passes the session\n transitions to STATE_EXPIRED. Defaults to a server-chosen horizon when\n unset.",
			InnerField: "expiresAt",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.pinned-parameters",
			Usage:      "Parameters forced onto tool calls made by this session's conversations.\n A pinned parameter is an overlay on a tool's JSON schema: the parameter\n is removed from what the LLM sees, and its value is always overwritten\n server-side with the pinned value — so the model cannot be tricked into\n calling a tool with a different id than the one the session was minted\n for (e.g. pin \"workspaceId\" for an OpenAPI tool with a\n /workspaces/{workspaceId} path). Flows to every objective the session\n creates.",
			InnerField: "pinnedParameters",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.subject",
			Usage:      "SubjectAssertion identifies a person within a tenant in the customer's own\n namespace — typically their user id. Asserting a subject upserts the\n subject record under the asserted tenant and associates the created\n resource with it. A subject assertion is only valid alongside a tenant\n assertion: subject identifiers are scoped to their tenant.",
			InnerField: "subject",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "spec.tenant",
			Usage:      "TenantAssertion identifies a tenant in the customer's own namespace — their\n org, company, or team identifier for an end user. Asserting a tenant\n upserts the tenant record in the workspace (keyed on `id` as the tenant's\n external_id) and associates the created resource with it.",
			InnerField: "tenant",
		},
		&requestflag.InnerFlag[any]{
			Name:       "spec.token-expires-at",
			Usage:      "Expiry of the token returned in `token`. Distinct from `expires_at`,\n which bounds the session itself.",
			InnerField: "tokenExpiresAt",
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
			Usage:      "Key-value pairs for categorization and filtering. Values are 0-63\n alphanumeric characters with \"-\", \"_\", or \".\" allowed between; keys\n follow the same shape and additionally accept an optional DNS-subdomain\n prefix (e.g. \"cadenya.com/\") of at most 253 characters.\n Examples: {\"priority\": \"high\", \"source\": \"api\", \"workflow\": \"onboarding\"}",
			InnerField: "labels",
		},
	},
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

var widgetSessionsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Retrieves a widget session. The bearer token is never returned on reads.",
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
	Action:          handleWidgetSessionsRetrieve,
	HideHelpCommand: true,
}

var widgetSessionsList = cli.Command{
	Name:    "list",
	Usage:   "Lists widget sessions in a workspace, filterable by widget, tenant, subject, and\nstate",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor from previous response.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[bool]{
			Name:      "include-info",
			Usage:     "When true, the `info` field on each returned session is populated.\n Requests with this flag count more against your rate limit.",
			QueryPath: "includeInfo",
		},
		&requestflag.Flag[string]{
			Name:      "labels",
			Usage:     "Filters by metadata labels. Comma-separated key=value pairs,\n e.g. \"env=prod,team=ai\". A resource matches only if every pair\n matches exactly (AND semantics).",
			QueryPath: "labels",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return.",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "sort-order",
			Usage:     "Sort order for results (asc or desc by creation time).",
			QueryPath: "sortOrder",
		},
		&requestflag.Flag[string]{
			Name:      "state",
			Usage:     "Filter by state.",
			QueryPath: "state",
		},
		&requestflag.Flag[string]{
			Name:      "subject-id",
			Usage:     "Filter to sessions asserted for a subject. Accepts the canonical\n `subj_…` form or the `external_id:<value>` form; the external_id form is\n scoped within a tenant and requires `tenant_id` to also be set.",
			QueryPath: "subjectId",
		},
		&requestflag.Flag[string]{
			Name:      "tenant-id",
			Usage:     "Filter to sessions belonging to a tenant. Accepts the canonical\n `tenant_…` form or the `external_id:<value>` form.",
			QueryPath: "tenantId",
		},
		&requestflag.Flag[string]{
			Name:      "widget-id",
			Usage:     "Filter to sessions on a specific widget. Accepts the canonical `wgt_…`\n form or the `external_id:<value>` form.",
			QueryPath: "widgetId",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleWidgetSessionsList,
	HideHelpCommand: true,
}

var widgetSessionsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Deletes a session and its secrets. The session's conversations are\ndisassociated, not deleted; use the tenant-level delete for full erasure.",
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
	Action:          handleWidgetSessionsDelete,
	HideHelpCommand: true,
}

var widgetSessionsDeleteTenant = cli.Command{
	Name:    "delete-tenant",
	Usage:   "Deletes every session belonging to a tenant across all widgets in the workspace,\nalong with the conversations those sessions created — built for GDPR erasure\nrequests. The tenant is required; an empty value is rejected rather than\nmatching everything.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Required:  true,
			PathParam: "workspaceId",
		},
		&requestflag.Flag[string]{
			Name:      "tenant-id",
			Usage:     "Tenant whose sessions to delete. Required — an empty value is rejected\n rather than matching everything. Accepts the canonical `tenant_…` form or\n the `external_id:<value>` form.",
			QueryPath: "tenantId",
		},
	},
	Action:          handleWidgetSessionsDeleteTenant,
	HideHelpCommand: true,
}

var widgetSessionsRevoke = cli.Command{
	Name:    "revoke",
	Usage:   "Transitions a session to STATE_REVOKED. Outstanding tokens stop working\nimmediately, open event streams close within seconds, and the session's secrets\nare deleted. Terminal.",
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
	Action:          handleWidgetSessionsRevoke,
	HideHelpCommand: true,
}

func handleWidgetSessionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := cadenya.WidgetSessionNewParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WidgetSessions.New(ctx, params, options...)
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
		Title:          "widget-sessions create",
		Transform:      transform,
	})
}

func handleWidgetSessionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	params := cadenya.WidgetSessionGetParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WidgetSessions.Get(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "widget-sessions retrieve",
		Transform:      transform,
	})
}

func handleWidgetSessionsList(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := cadenya.WidgetSessionListParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.WidgetSessions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "widget-sessions list",
			Transform:      transform,
		})
	} else {
		iter := client.WidgetSessions.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "widget-sessions list",
			Transform:      transform,
		})
	}
}

func handleWidgetSessionsDelete(ctx context.Context, cmd *cli.Command) error {
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

	params := cadenya.WidgetSessionDeleteParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	return client.WidgetSessions.Delete(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleWidgetSessionsDeleteTenant(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := cadenya.WidgetSessionDeleteTenantParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WidgetSessions.DeleteTenant(ctx, params, options...)
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
		Title:          "widget-sessions delete-tenant",
		Transform:      transform,
	})
}

func handleWidgetSessionsRevoke(ctx context.Context, cmd *cli.Command) error {
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

	params := cadenya.WidgetSessionRevokeParams{
		WorkspaceID: cadenya.String(cmd.Value("workspace-id").(string)),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.WidgetSessions.Revoke(
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "widget-sessions revoke",
		Transform:      transform,
	})
}
