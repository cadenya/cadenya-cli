package cmd

import (
	"github.com/urfave/cli/v3"
)

// configCommand is the parent group for configuration-as-code subcommands.
// It is registered from cmd.go's Commands slice.
var configCommand = &cli.Command{
	Name:     "config",
	Category: "CONFIG",
	Usage:    "Apply a workspace configuration from a YAML file",
	Suggest:  true,
	Commands: []*cli.Command{
		&configApply,
	},
	HideHelpCommand: true,
}
