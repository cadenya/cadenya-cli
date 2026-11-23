package commands

import (
	cli "github.com/urfave/cli/v3"

	sdk "go.cadenya.com/cadenya-go"
)

// ToolSetRulesDeleteConversion is the typed request produced from one urfave command.
type ToolSetRulesDeleteConversion struct {
	Params sdk.ToolSetRuleDeleteParams
	Body   any
}

// ConvertToolSetRulesDelete reads urfave values according to Redwood's operation IR.
func ConvertToolSetRulesDelete(cmd *cli.Command, out *ToolSetRulesDeleteConversion) error {
	values := map[string]any{}
	if cmd.IsSet("workspace-id") {
		values["workspaceId"] = cmd.String("workspace-id")
	}
	if err := decodeParams(values, &out.Params); err != nil {
		return cli.Exit(err.Error(), 2)
	}
	return nil
}
