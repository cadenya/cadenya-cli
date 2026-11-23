package commands

import (
	cli "github.com/urfave/cli/v3"

	sdk "go.cadenya.com/cadenya-go"
)

// ToolSetsSyncConversion is the typed request produced from one urfave command.
type ToolSetsSyncConversion struct {
	Params sdk.ToolSetSyncParams
	Body   any
}

// ConvertToolSetsSync reads urfave values according to Redwood's operation IR.
func ConvertToolSetsSync(cmd *cli.Command, out *ToolSetsSyncConversion) error {
	values := map[string]any{}
	if cmd.IsSet("workspace-id") {
		values["workspaceId"] = cmd.String("workspace-id")
	}
	if err := decodeParams(values, &out.Params); err != nil {
		return cli.Exit(err.Error(), 2)
	}
	return nil
}
