package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
)

func TestWorkspaceAdminProfilesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-admin:profiles", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--labels", "labels",
			"--limit", "0",
			"--query", "query",
		)
	})
}
