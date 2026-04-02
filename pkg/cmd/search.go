// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"go.cadenya.com/cadenya-sdk-go"
	"go.cadenya.com/cadenya-sdk-go/option"
	"os"

	"github.com/stainless-sdks/cadenya-cli/internal/apiquery"
	"github.com/stainless-sdks/cadenya-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var searchSearchToolsOrToolSets = cli.Command{
	Name:    "search-tools-or-tool-sets",
	Usage:   "Searches for tools or tool sets in the workspace",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "query",
			QueryPath: "query",
		},
	},
	Action:          handleSearchSearchToolsOrToolSets,
	HideHelpCommand: true,
}

func handleSearchSearchToolsOrToolSets(ctx context.Context, cmd *cli.Command) error {
	client := cadenya.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := cadenya.SearchSearchToolsOrToolSetsParams{}

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
	_, err = client.Search.SearchToolsOrToolSets(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "search search-tools-or-tool-sets", obj, format, transform)
}
