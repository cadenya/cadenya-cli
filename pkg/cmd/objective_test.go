// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
)

func TestObjectivesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives", "create",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--agent-id", "agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
			"--system-prompt-data", "{foo: bar}",
			"--episodic-memory", "{key: key}",
			"--first-user-message", "firstUserMessage",
			"--first-user-message-data", "{foo: bar}",
			"--memory-cascade", "{memoryLayerId: memlyr_01HXKD2E5NQM3T9AYWCFFFBMJH, memoryEntryId: mementry_01HXKD2E5NQM3T9AYWCF5E52Z0}",
			"--metadata", "{externalId: externalId, labels: {foo: string}}",
			"--pinned-parameters", "{foo: string}",
			"--secret", "{name: name, value: value}",
			"--subject", "{id: customer-user-42, name: Jane Doe}",
			"--tenant", "{id: acme-corp, name: Acme Corp}",
			"--variation-id", "agentvar_01HXKD2E5NQM3T9AYWCF32BSPP",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(objectivesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives", "create",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--agent-id", "agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
			"--system-prompt-data", "{foo: bar}",
			"--episodic-memory.key", "key",
			"--first-user-message", "firstUserMessage",
			"--first-user-message-data", "{foo: bar}",
			"--memory-cascade.memory-layer-id", "memlyr_01HXKD2E5NQM3T9AYWCFFFBMJH",
			"--memory-cascade.memory-entry-id", "mementry_01HXKD2E5NQM3T9AYWCF5E52Z0",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--pinned-parameters", "{foo: string}",
			"--secret.name", "name",
			"--secret.value", "value",
			"--subject.id", "customer-user-42",
			"--subject.name", "Jane Doe",
			"--tenant.id", "acme-corp",
			"--tenant.name", "Acme Corp",
			"--variation-id", "agentvar_01HXKD2E5NQM3T9AYWCF32BSPP",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agentId: agent_01HXKD2E5NQM3T9AYWCFMGWT9Y\n" +
			"systemPromptData:\n" +
			"  foo: bar\n" +
			"episodicMemory:\n" +
			"  key: key\n" +
			"firstUserMessage: firstUserMessage\n" +
			"firstUserMessageData:\n" +
			"  foo: bar\n" +
			"memoryCascade:\n" +
			"  - memoryLayerId: memlyr_01HXKD2E5NQM3T9AYWCFFFBMJH\n" +
			"    memoryEntryId: mementry_01HXKD2E5NQM3T9AYWCF5E52Z0\n" +
			"metadata:\n" +
			"  externalId: externalId\n" +
			"  labels:\n" +
			"    foo: string\n" +
			"pinnedParameters:\n" +
			"  foo: string\n" +
			"secrets:\n" +
			"  - name: name\n" +
			"    value: value\n" +
			"subject:\n" +
			"  id: customer-user-42\n" +
			"  name: Jane Doe\n" +
			"tenant:\n" +
			"  id: acme-corp\n" +
			"  name: Acme Corp\n" +
			"variationId: agentvar_01HXKD2E5NQM3T9AYWCF32BSPP\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"objectives", "create",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
		)
	})
}

func TestObjectivesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives", "retrieve",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
		)
	})
}

func TestObjectivesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives", "list",
			"--max-items", "10",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--agent-id", "agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
			"--agent-schedule-id", "agentScheduleId",
			"--cursor", "cursor",
			"--include-info=true",
			"--labels", "labels",
			"--limit", "0",
			"--parent-objective-id", "parentObjectiveId",
			"--profile-id", "profile_01HXKD2E5NQM3T9AYWCFS0AP08",
			"--sort-order", "sortOrder",
			"--state", "STATE_UNSPECIFIED",
			"--subject-id", "subjectId",
			"--tenant-id", "tenantId",
		)
	})
}

func TestObjectivesCancel(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives", "cancel",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--objective-id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
			"--reason", "reason",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("reason: reason")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"objectives", "cancel",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--objective-id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
		)
	})
}

func TestObjectivesCompact(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives", "compact",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--objective-id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
			"--compaction-config", "{summarization: {instructions: instructions}, toolResultClearing: {preserveRecentResults: 0}, triggerThreshold: 0}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(objectivesCompact)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives", "compact",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--objective-id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
			"--compaction-config.summarization", "{instructions: instructions}",
			"--compaction-config.tool-result-clearing", "{preserveRecentResults: 0}",
			"--compaction-config.trigger-threshold", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"compactionConfig:\n" +
			"  summarization:\n" +
			"    instructions: instructions\n" +
			"  toolResultClearing:\n" +
			"    preserveRecentResults: 0\n" +
			"  triggerThreshold: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"objectives", "compact",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--objective-id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
		)
	})
}

func TestObjectivesContinue(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives", "continue",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--objective-id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
			"--message", "message",
			"--enqueue=true",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"message: message\n" +
			"enqueue: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"objectives", "continue",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--objective-id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
		)
	})
}

func TestObjectivesListContextWindows(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives", "list-context-windows",
			"--max-items", "10",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--objective-id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
			"--cursor", "cursor",
			"--include-info=true",
			"--labels", "labels",
			"--limit", "0",
		)
	})
}

func TestObjectivesListEvents(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives", "list-events",
			"--max-items", "10",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--objective-id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
			"--cursor", "cursor",
			"--include-info=true",
			"--labels", "labels",
			"--limit", "0",
			"--since-event-id", "sinceEventId",
			"--sort-order", "sortOrder",
			"--window-id", "windowId",
		)
	})
}

func TestObjectivesRetrieveDiagnostics(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives", "retrieve-diagnostics",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--objective-id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
		)
	})
}

func TestObjectivesStreamEvents(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives", "stream-events",
			"--max-items", "10",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--objective-id", "obj_01HXKD2E5NQM3T9AYWCFQAZGFV",
		)
	})
}
