package commands

import (
	cli "github.com/urfave/cli/v3"

	sdk "go.cadenya.com/cadenya-go"
)

// ToolSetRulesRetrieveConversion is the typed request produced from one urfave command.
type ToolSetRulesRetrieveConversion struct {
	Params sdk.ToolSetRuleRetrieveParams
	Body   any
}

// ConvertToolSetRulesRetrieve reads urfave values according to Redwood's operation IR.
func ConvertToolSetRulesRetrieve(cmd *cli.Command, out *ToolSetRulesRetrieveConversion) error {
	values := map[string]any{}
	if cmd.IsSet("workspace-id") {
		values["workspaceId"] = cmd.String("workspace-id")
	}
	if err := decodeParams(values, &out.Params); err != nil {
		return cli.Exit(err.Error(), 2)
	}
	return nil
}
