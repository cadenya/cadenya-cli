package commands

import (
	cli "github.com/urfave/cli/v3"

	sdk "go.cadenya.com/cadenya-go"
)

const bodySchemaAPIKeysRemoveWorkspace = "{\"$defs\":{},\"properties\":{\"workspaceId\":{\"type\":\"string\"}},\"required\":[\"workspaceId\"],\"type\":\"object\"}"

// APIKeysRemoveWorkspaceConversion is the typed request produced from one urfave command.
type APIKeysRemoveWorkspaceConversion struct {
	Params sdk.APIKeyRemoveWorkspaceParams
	Body   any
}

// ConvertAPIKeysRemoveWorkspace reads urfave values according to Redwood's operation IR.
func ConvertAPIKeysRemoveWorkspace(cmd *cli.Command, out *APIKeysRemoveWorkspaceConversion) error {
	values := map[string]any{}
	_schema := parseBodySchema(bodySchemaAPIKeysRemoveWorkspace)
	_body := newBodyBuilder()
	_strict := cmd.Bool("strict")
	var _rawBody any
	if cmd.IsSet("file") {
		if err := _body.applyFile("file", cmd.String("file"), _schema, _strict); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if cmd.IsSet("workspace-id") {
		_v, err := stringArg("workspace-id", cmd.String("workspace-id"))
		if err != nil {
			return cli.Exit(err.Error(), 2)
		}
		if err := _body.set("workspace-id", []string{"workspaceId"}, _v); err != nil {
			return cli.Exit(err.Error(), 2)
		}
	}
	if err := _body.finish(_schema, map[string]string{"workspaceId": "--workspace-id"}, _strict); err != nil {
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
