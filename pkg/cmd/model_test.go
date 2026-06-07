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
			"--bundle-key", "bundleKey",
			"--cursor", "cursor",
			"--include-info=true",
			"--limit", "0",
			"--prefix", "prefix",
			"--query", "query",
			"--sort-order", "sortOrder",
			"--status", "MODEL_STATUS_UNSPECIFIED",
		)
	})
}

func TestModelsSetStatus(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"models", "set-status",
			"--workspace-id", "workspaceId",
			"--id", "id",
			"--status", "MODEL_STATUS_UNSPECIFIED",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("status: MODEL_STATUS_UNSPECIFIED")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"models", "set-status",
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
			"--model-swap", "{currentModelId: currentModelId, nextModelId: nextModelId}",
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
			"--model-swap.next-model-id", "nextModelId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"modelSwaps:\n" +
			"  - currentModelId: currentModelId\n" +
			"    nextModelId: nextModelId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"models", "swap",
			"--workspace-id", "workspaceId",
		)
	})
}
