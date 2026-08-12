package cmd

import (
	"context"
	"fmt"
	"go.cadenya.com/cadenya-go"
	"go.cadenya.com/cadenya-go/option"

	"github.com/cadenya/cadenya-cli/internal/apiquery"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var accountRetrieve = cli.Command{
	Name:            "retrieve",
	Usage:           "Retrieves the current account for the token accessing the API. Useful to check\nif the credentials are valid.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleAccountRetrieve,
	HideHelpCommand: true,
}

var accountRotateChallengeToken = cli.Command{
	Name:            "rotate-challenge-token",
	Usage:           "Rotates the challenge token sent in the X-Cadenya-Challenge-Token header on MCP\ntools/list requests. Returns only the new token.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleAccountRotateChallengeToken,
	HideHelpCommand: true,
}

var accountRotateWebhookSigningKey = cli.Command{
	Name:            "rotate-webhook-signing-key",
	Usage:           "Rotates the webhook signing key for the account. Returns only the new key.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleAccountRotateWebhookSigningKey,
	HideHelpCommand: true,
}

func handleAccountRetrieve(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Account.Get(ctx, options...)
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
		Title:          "account retrieve",
		Transform:      transform,
	})
}

func handleAccountRotateChallengeToken(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Account.RotateChallengeToken(ctx, options...)
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
		Title:          "account rotate-challenge-token",
		Transform:      transform,
	})
}

func handleAccountRotateWebhookSigningKey(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Account.RotateWebhookSigningKey(ctx, options...)
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
		Title:          "account rotate-webhook-signing-key",
		Transform:      transform,
	})
}
