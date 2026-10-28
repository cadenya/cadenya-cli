package commands

import (
	cli "github.com/urfave/cli/v3"

	sdk "go.cadenya.com/cadenya-go"
)

const bodySchemaAPIKeysCreate = "{\"$defs\":{\"APIKeySpec\":{\"properties\":{\"allWorkspaces\":{\"description\":\"When true the key reaches every workspace in the account, including\\n workspaces created later, and workspace_assignments is ignored on\\n requests. Reads then list every workspace the key currently reaches.\",\"type\":\"boolean\"},\"description\":{\"description\":\"Free-form description of what this API key is used for.\",\"type\":\"string\"},\"permissions\":{\"description\":\"Scopes granted to this key. Each entry is a colon-separated\\n resource:verb string (e.g. \\\"objectives:manage\\\").\\n\\n Resources: agents, objectives, tools, memory, api_keys, workspaces,\\n widgets, widget_sessions, tenants, endpoints (trusted endpoints),\\n secrets, account.\\n Verbs: read and manage, where manage implies read — a stored scope set\\n is normalized to drop \\\"x:read\\\" when \\\"x:manage\\\" is present. The secrets\\n and account resources support only manage. \\\"*\\\" is an explicit\\n full-access grant.\\n\\n api_keys also has the verb admin (\\\"api_keys:admin\\\"), which implies\\n manage. It is required to create or change a key that reaches every\\n workspace, and to manage keys created by other members; without it a\\n caller manages only its own keys.\\n\\n Some scopes depend on another: \\\"tools:manage\\\" cannot be granted\\n without \\\"secrets:manage\\\", because tool set headers and base URLs\\n resolve workspace secrets ({{ secrets.NAME }}) before the source is\\n validated, synced, or called. A set that holds a scope without its\\n dependency is rejected as invalid.\\n\\n Scopes are deny-by-default: a key with an empty list can call only\\n scope-free endpoints. Full access is always an explicit \\\"*\\\" grant.\",\"items\":{\"type\":\"string\"},\"type\":\"array\"},\"workspaceAssignments\":{\"description\":\"Complete set of workspaces the key is assigned to. Order has no meaning\\n and duplicates are collapsed. On create, at least one assignment is\\n required unless all_workspaces is set. On update, selecting\\n spec.workspace_assignments in update_mask replaces the entire set. Use\\n the :addWorkspace and :removeWorkspace actions to change one workspace.\",\"items\":{\"$ref\":\"APIKeyWorkspaceAssignment\"},\"type\":\"array\"}},\"required\":[],\"type\":\"object\"},\"APIKeyWorkspaceAssignment\":{\"properties\":{\"workspaceId\":{\"description\":\"Canonical workspace ID, or the `external_id:<value>` form.\",\"type\":\"string\"}},\"required\":[\"workspaceId\"],\"type\":\"object\"},\"CreateAccountResourceMetadata\":{\"properties\":{\"externalId\":{\"description\":\"External ID for the resource (e.g., a workflow ID from an external system)\",\"type\":\"string\"},\"labels\":{\"additionalProperties\":{\"type\":\"string\"},\"description\":\"Key-value pairs for categorization and filtering. Values are 0-63\\n alphanumeric characters with \\\"-\\\", \\\"_\\\", or \\\".\\\" allowed between; keys\\n follow the same shape and additionally accept an optional DNS-subdomain\\n prefix (e.g. \\\"cadenya.com/\\\") of at most 253 characters.\\n Examples: {\\\"environment\\\": \\\"production\\\", \\\"team\\\": \\\"platform\\\", \\\"version\\\": \\\"v2\\\"}\",\"type\":\"object\"},\"name\":{\"description\":\"Human-readable name for the resource (e.g., \\\"Production API Key\\\", \\\"Staging Workspace\\\")\",\"type\":\"string\"}},\"required\":[\"name\"],\"type\":\"object\"}},\"properties\":{\"metadata\":{\"$ref\":\"CreateAccountResourceMetadata\"},\"spec\":{\"$ref\":\"APIKeySpec\"}},\"required\":[\"metadata\",\"spec\"],\"type\":\"object\"}"

// APIKeysCreateConversion is the typed request produced from one urfave command.
type APIKeysCreateConversion struct {
	Params sdk.APIKeyCreateParams
	Body   any
}

// ConvertAPIKeysCreate reads urfave values according to Redwood's operation IR.
func ConvertAPIKeysCreate(cmd *cli.Command, out *APIKeysCreateConversion) error {
	values := map[string]any{}
	_schema := parseBodySchema(bodySchemaAPIKeysCreate)
	_body := newBodyBuilder()
	_strict := cmd.Bool("strict")
	var _rawBody any
	if cmd.IsSet("file") {
		if err := _body.applyFile("file", cmd.String("file"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("metadata") {
		if err := _body.applyDoc("metadata", []string{"metadata"}, cmd.String("metadata"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("spec") {
		if err := _body.applyDoc("spec", []string{"spec"}, cmd.String("spec"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("name") {
		_v, err := stringArg("name", cmd.String("name"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("name", []string{"metadata", "name"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("external-id") {
		_v, err := stringArg("external-id", cmd.String("external-id"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("external-id", []string{"metadata", "externalId"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("label") {
		if err := _body.applyEntries("label", []string{"metadata", "labels"}, cmd.StringSlice("label"), scalarString, nil); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("description") {
		_v, err := stringArg("description", cmd.String("description"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("description", []string{"spec", "description"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("permission") {
		if err := _body.applyScalarItems("permission", []string{"spec", "permissions"}, cmd.StringSlice("permission"), scalarString, nil); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("all-workspaces") {
		if err := _body.set("all-workspaces", []string{"spec", "allWorkspaces"}, cmd.Bool("all-workspaces")); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("workspace-assignment") {
		if err := _body.applyShorthandItems("workspace-assignment", []string{"spec", "workspaceAssignments"}, cmd.StringSlice("workspace-assignment"), shorthandSpec{Fields: []shorthandField{{Wire: "workspaceId", Key: "workspace-id", Kind: scalarString, Enum: nil, Required: true}}}); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if err := _body.finish(_schema, map[string]string{"metadata": "--metadata", "metadata.name": "--name", "metadata.externalId": "--external-id", "metadata.labels": "--label", "spec": "--spec", "spec.description": "--description", "spec.permissions": "--permission", "spec.allWorkspaces": "--all-workspaces", "spec.workspaceAssignments": "--workspace-assignment"}, _strict); err != nil {
		return cli.Exit(err.Error(), 2)
	}
	_ = _rawBody
	for _k, _v := range _body.body {
		values[_k] = _v
	}
	if err := decodeParams(values, &out.Params); err != nil {
		return cli.Exit(err.Error(), 2)
	}
	out.Body = _body.body
	return nil
}
