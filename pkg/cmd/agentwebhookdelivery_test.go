// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
)

func TestAgentsWebhookDeliveriesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:webhook-deliveries", "list",
			"--max-items", "10",
			"--agent-id", "agentId",
			"--cursor", "cursor",
			"--event-type", "OBJECTIVE_EVENT_TYPE_UNSPECIFIED",
			"--limit", "0",
			"--objective-id", "objectiveId",
		)
	})
}
