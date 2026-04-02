// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cadenya-cli/internal/mocktest"
)

func TestModelsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"models", "retrieve",
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
			"--cursor", "cursor",
			"--limit", "0",
			"--prefix", "prefix",
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
			"--id", "id",
		)
	})
}
