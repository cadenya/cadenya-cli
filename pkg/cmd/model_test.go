// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
)

func TestModelsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"models", "retrieve",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}

func TestModelsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"models", "list",
			"--max-items", "10",
			"--workspace-id", "workspaceId",
			"--ai-provider-key-id", "aiProviderKeyId",
			"--cursor", "cursor",
			"--include-info=true",
			"--is-assigned=true",
			"--limit", "0",
			"--prefix", "prefix",
			"--query", "query",
			"--sort-order", "sortOrder",
			"--state", "STATE_UNSPECIFIED",
		)
	})
}

func TestModelsDisable(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"models", "disable",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}

func TestModelsEnable(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"models", "enable",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}

func TestModelsSwap(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"models", "swap",
			"--workspace-id", "workspaceId",
			"--model-swap", "{currentModelId: currentModelId, disableCurrentAfterSwap: true, nextModelId: nextModelId}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(modelsSwap)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"models", "swap",
			"--workspace-id", "workspaceId",
			"--model-swap.current-model-id", "currentModelId",
			"--model-swap.disable-current-after-swap=true",
			"--model-swap.next-model-id", "nextModelId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"modelSwaps:\n" +
			"  - currentModelId: currentModelId\n" +
			"    disableCurrentAfterSwap: true\n" +
			"    nextModelId: nextModelId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"models", "swap",
			"--workspace-id", "workspaceId",
		)
	})
}
