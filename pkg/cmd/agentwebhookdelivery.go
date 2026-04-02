// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"go.cadenya.com/cadenya-go"
	"go.cadenya.com/cadenya-go/option"
	"os"

	"github.com/cadenya/cli/internal/apiquery"
	"github.com/cadenya/cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var agentsWebhookDeliveriesList = cli.Command{
	Name:    "list",
	Usage:   "Lists all webhook deliveries for an agent",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "agent-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor from previous response",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "event-type",
			Usage:     "Optional filter by event type",
			QueryPath: "eventType",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "objective-id",
			Usage:     "Optional filter by objective ID",
			QueryPath: "objectiveId",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAgentsWebhookDeliveriesList,
	HideHelpCommand: true,
}

func handleAgentsWebhookDeliveriesList(ctx context.Context, cmd *cli.Command) error {
	client := gocadenyacomcadenyago.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("agent-id") && len(unusedArgs) > 0 {
		cmd.Set("agent-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := gocadenyacomcadenyago.AgentWebhookDeliveryListParams{}

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
		_, err = client.Agents.WebhookDeliveries.List(
			ctx,
			cmd.Value("agent-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "agents:webhook-deliveries list", obj, format, transform)
	} else {
		iter := client.Agents.WebhookDeliveries.ListAutoPaging(
			ctx,
			cmd.Value("agent-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(os.Stdout, "agents:webhook-deliveries list", iter, format, transform, maxItems)
	}
}
