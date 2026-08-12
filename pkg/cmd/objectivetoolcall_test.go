package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
)

func TestObjectivesToolCallsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives:tool-calls", "retrieve",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--objective-id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
			"--tool-call-id", "toolcall_01HXKD2E5NQM3T9AYWCFTANFGV",
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
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--objective-id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
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
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--objective-id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
			"--tool-call-id", "toolcall_01HXKD2E5NQM3T9AYWCFTANFGV",
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
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--objective-id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
			"--tool-call-id", "toolcall_01HXKD2E5NQM3T9AYWCFTANFGV",
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
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--objective-id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
			"--tool-call-id", "toolcall_01HXKD2E5NQM3T9AYWCFTANFGV",
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
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--objective-id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
			"--tool-call-id", "toolcall_01HXKD2E5NQM3T9AYWCFTANFGV",
			"--content", "{text: {text: text}, type: text}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"content:\n" +
			"  - text:\n" +
			"      text: text\n" +
			"    type: text\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"objectives:tool-calls", "set-content",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--objective-id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
			"--tool-call-id", "toolcall_01HXKD2E5NQM3T9AYWCFTANFGV",
		)
	})
}
