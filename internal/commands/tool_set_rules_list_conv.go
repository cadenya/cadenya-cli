package commands

import (
	cli "github.com/urfave/cli/v3"

	sdk "go.cadenya.com/cadenya-go"
)

// ToolSetRulesListConversion is the typed request produced from one urfave command.
type ToolSetRulesListConversion struct {
	Params sdk.ToolSetRuleListParams
	Body   any
}

// ConvertToolSetRulesList reads urfave values according to Redwood's operation IR.
func ConvertToolSetRulesList(cmd *cli.Command, out *ToolSetRulesListConversion) error {
	values := map[string]any{}
	if cmd.IsSet("workspace-id") {
		values["workspaceId"] = cmd.String("workspace-id")
	}
	if err := decodeParams(values, &out.Params); err != nil {
		return cli.Exit(err.Error(), 2)
	}
	return nil
}
