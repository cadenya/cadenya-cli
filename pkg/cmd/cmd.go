// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/cadenya/cadenya-cli/internal/autocomplete"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command            *cli.Command
	CommandErrorBuffer bytes.Buffer
)

func init() {
	Command = &cli.Command{
		Name:      "cadenya",
		Usage:     "CLI for the cadenya API",
		Suggest:   true,
		Version:   Version,
		ErrWriter: &CommandErrorBuffer,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
				Validator: func(baseURL string) error {
					return ValidateBaseURL(baseURL, "--base-url")
				},
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
			&requestflag.Flag[string]{
				Name:    "api-key",
				Sources: cli.EnvVars("CADENYA_API_KEY"),
			},
			&requestflag.Flag[string]{
				Name:    "webhook-key",
				Sources: cli.EnvVars("CADENYA_WEBHOOK_KEY"),
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "account",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&accountRetrieve,
					&accountRotateWebhookSigningKey,
				},
			},
			{
				Name:     "agents",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&agentsCreate,
					&agentsRetrieve,
					&agentsUpdate,
					&agentsList,
					&agentsDelete,
				},
			},
			{
				Name:     "agents:webhook-deliveries",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&agentsWebhookDeliveriesList,
				},
			},
			{
				Name:     "agent-variations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&agentVariationsCreate,
					&agentVariationsRetrieve,
					&agentVariationsUpdate,
					&agentVariationsList,
					&agentVariationsDelete,
					&agentVariationsAddAssignment,
					&agentVariationsRemoveAssignment,
				},
			},
			{
				Name:     "objectives",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&objectivesCreate,
					&objectivesRetrieve,
					&objectivesList,
					&objectivesCancel,
					&objectivesCompact,
					&objectivesContinue,
					&objectivesListContextWindows,
					&objectivesListEvents,
				},
			},
			{
				Name:     "objectives:tools",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&objectivesToolsList,
				},
			},
			{
				Name:     "objectives:tool-calls",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&objectivesToolCallsList,
					&objectivesToolCallsApprove,
					&objectivesToolCallsDeny,
				},
			},
			{
				Name:     "objectives:tasks",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&objectivesTasksRetrieve,
					&objectivesTasksList,
				},
			},
			{
				Name:     "objectives:feedback",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&objectivesFeedbackCreate,
					&objectivesFeedbackList,
				},
			},
			{
				Name:     "memory-layers",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&memoryLayersCreate,
					&memoryLayersRetrieve,
					&memoryLayersUpdate,
					&memoryLayersList,
					&memoryLayersDelete,
				},
			},
			{
				Name:     "memory-layers:entries",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&memoryLayersEntriesCreate,
					&memoryLayersEntriesRetrieve,
					&memoryLayersEntriesUpdate,
					&memoryLayersEntriesList,
					&memoryLayersEntriesDelete,
				},
			},
			{
				Name:     "uploads",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&uploadsCreate,
					&uploadsRetrieve,
				},
			},
			{
				Name:     "models",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&modelsRetrieve,
					&modelsList,
					&modelsSetStatus,
				},
			},
			{
				Name:     "search",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&searchSearchToolsOrToolSets,
				},
			},
			{
				Name:     "tool-sets",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&toolSetsCreate,
					&toolSetsRetrieve,
					&toolSetsUpdate,
					&toolSetsList,
					&toolSetsDelete,
					&toolSetsListEvents,
				},
			},
			{
				Name:     "tool-sets:tools",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&toolSetsToolsCreate,
					&toolSetsToolsRetrieve,
					&toolSetsToolsUpdate,
					&toolSetsToolsList,
					&toolSetsToolsDelete,
				},
			},
			{
				Name:     "api-keys",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&apiKeysCreate,
					&apiKeysRetrieve,
					&apiKeysUpdate,
					&apiKeysList,
					&apiKeysDelete,
					&apiKeysRotate,
				},
			},
			{
				Name:     "workspace-secrets",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&workspaceSecretsCreate,
					&workspaceSecretsRetrieve,
					&workspaceSecretsUpdate,
					&workspaceSecretsList,
					&workspaceSecretsDelete,
				},
			},
			{
				Name:     "workspaces",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&workspacesList,
					&workspacesGet,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "cadenya @manpages [-o cadenya.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "cadenya.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "cadenya.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
