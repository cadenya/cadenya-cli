// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/cadenya-cli/internal/mocktest"
)

func TestObjectivesToolsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives:tools", "list",
			"--max-items", "10",
			"--objective-id", "objectiveId",
			"--cursor", "cursor",
			"--limit", "0",
		)
	})
}
