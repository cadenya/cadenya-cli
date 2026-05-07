// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
)

func TestBulkWorkspaceResourcesResultsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"bulk-workspace-resources:results", "list",
			"--max-items", "10",
			"--workspace-id", "workspaceId",
			"--bulk-workspace-apply-id", "bulkWorkspaceApplyId",
			"--action", "ACTION_UNSPECIFIED",
			"--cursor", "cursor",
			"--limit", "0",
			"--sort-order", "sortOrder",
			"--type", "type",
		)
	})
}
