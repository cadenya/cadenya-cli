package commands

import (
	cli "github.com/urfave/cli/v3"

	sdk "go.cadenya.com/cadenya-go"
)

// TrustedEndpointsListConversion is the typed request produced from one urfave command.
type TrustedEndpointsListConversion struct {
	Params sdk.TrustedEndpointListParams
	Body   any
}

// ConvertTrustedEndpointsList reads urfave values according to Redwood's operation IR.
func ConvertTrustedEndpointsList(cmd *cli.Command, out *TrustedEndpointsListConversion) error {
	values := map[string]any{}
	if cmd.IsSet("limit") {
		values["limit"] = cmd.Int32("limit")
	}
	if cmd.IsSet("cursor") {
		values["cursor"] = cmd.String("cursor")
	}
	if cmd.IsSet("labels") {
		values["labels"] = cmd.String("labels")
	}
	if err := decodeParams(values, &out.Params); err != nil {
		return cli.Exit(err.Error(), 2)
	}
	return nil
}
