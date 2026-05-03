// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
)

func TestObjectivesTasksRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives:tasks", "retrieve",
			"--workspace-id", "workspaceId",
			"--objective-id", "objectiveId",
			"--id", "id",
		)
	})
}

func TestObjectivesTasksList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives:tasks", "list",
			"--max-items", "10",
			"--workspace-id", "workspaceId",
			"--objective-id", "objectiveId",
			"--cursor", "cursor",
			"--limit", "0",
			"--sort-order", "sortOrder",
		)
	})
}
