package commands

import (
	cli "github.com/urfave/cli/v3"

	sdk "go.cadenya.com/cadenya-go"
)

// ObjectivesRemoveQueuedActionConversion is the typed request produced from one urfave command.
type ObjectivesRemoveQueuedActionConversion struct {
	Params sdk.ObjectiveRemoveQueuedActionParams
	Body   any
}

// ConvertObjectivesRemoveQueuedAction reads urfave values according to Redwood's operation IR.
func ConvertObjectivesRemoveQueuedAction(cmd *cli.Command, out *ObjectivesRemoveQueuedActionConversion) error {
	values := map[string]any{}
	if cmd.IsSet("workspace-id") {
		values["workspaceId"] = cmd.String("workspace-id")
	}
	if cmd.IsSet("queued-action-id") {
		values["queuedActionId"] = cmd.String("queued-action-id")
	}
	if err := decodeParams(values, &out.Params); err != nil {
		return cli.Exit(err.Error(), 2)
	}
	return nil
}
