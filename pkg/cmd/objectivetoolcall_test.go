// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
)

func TestObjectivesToolCallsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives:tool-calls", "list",
			"--max-items", "10",
			"--workspace-id", "workspaceId",
			"--objective-id", "objectiveId",
			"--cursor", "cursor",
			"--include-info=true",
			"--limit", "0",
			"--status", "TOOL_CALL_STATUS_UNSPECIFIED",
		)
	})
}

func TestObjectivesToolCallsApprove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives:tool-calls", "approve",
			"--workspace-id", "workspaceId",
			"--objective-id", "objectiveId",
			"--tool-call-id", "toolCallId",
		)
	})
}

func TestObjectivesToolCallsDeny(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives:tool-calls", "deny",
			"--workspace-id", "workspaceId",
			"--objective-id", "objectiveId",
			"--tool-call-id", "toolCallId",
			"--memo", "memo",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("memo: memo")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"objectives:tool-calls", "deny",
			"--workspace-id", "workspaceId",
			"--objective-id", "objectiveId",
			"--tool-call-id", "toolCallId",
		)
	})
}
