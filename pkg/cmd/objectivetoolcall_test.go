// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
)

func TestObjectivesToolCallsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives:tool-calls", "retrieve",
			"--workspace-id", "workspaceId",
			"--objective-id", "objectiveId",
			"--tool-call-id", "toolCallId",
		)
	})
}

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
			"--execution-status", "TOOL_CALL_EXECUTION_STATUS_UNSPECIFIED",
			"--include-info=true",
			"--labels", "labels",
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

func TestObjectivesToolCallsSetContent(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives:tool-calls", "set-content",
			"--workspace-id", "workspaceId",
			"--objective-id", "objectiveId",
			"--tool-call-id", "toolCallId",
			"--content", "{audio: {data: data, mimeType: mimeType}, image: {data: data, mimeType: mimeType}, text: {text: text}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(objectivesToolCallsSetContent)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives:tool-calls", "set-content",
			"--workspace-id", "workspaceId",
			"--objective-id", "objectiveId",
			"--tool-call-id", "toolCallId",
			"--content.audio", "{data: data, mimeType: mimeType}",
			"--content.image", "{data: data, mimeType: mimeType}",
			"--content.text", "{text: text}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"content:\n" +
			"  - audio:\n" +
			"      data: data\n" +
			"      mimeType: mimeType\n" +
			"    image:\n" +
			"      data: data\n" +
			"      mimeType: mimeType\n" +
			"    text:\n" +
			"      text: text\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"objectives:tool-calls", "set-content",
			"--workspace-id", "workspaceId",
			"--objective-id", "objectiveId",
			"--tool-call-id", "toolCallId",
		)
	})
}
