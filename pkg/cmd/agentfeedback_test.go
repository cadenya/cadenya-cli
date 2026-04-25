// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
)

func TestAgentsFeedbackList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:feedback", "list",
			"--max-items", "10",
			"--agent-id", "agentId",
			"--agent-variation-id", "agentVariationId",
			"--created-after", "'2019-12-27T18:11:19.117Z'",
			"--created-before", "'2019-12-27T18:11:19.117Z'",
			"--cursor", "cursor",
			"--include-info=true",
			"--limit", "0",
			"--query", "query",
			"--sentiment", "FEEDBACK_SENTIMENT_UNSPECIFIED",
		)
	})
}
