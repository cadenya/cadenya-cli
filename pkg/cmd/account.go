// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"go.cadenya.com/cadenya-sdk-go"
	"go.cadenya.com/cadenya-sdk-go/option"
	"os"

	"github.com/stainless-sdks/cadenya-cli/internal/apiquery"
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "account retrieve", obj, format, transform)
}
