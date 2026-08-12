package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
)

func TestAgentsSchedulesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:schedules", "create",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--agent-id", "agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{schedule: {calendars: [{comment: comment, dayOfMonth: [{end: 0, start: 0, step: 0}], dayOfWeek: [{end: 0, start: 0, step: 0}], hour: [{end: 0, start: 0, step: 0}], minute: [{end: 0, start: 0, step: 0}], month: [{end: 0, start: 0, step: 0}], second: [{end: 0, start: 0, step: 0}]}], intervals: [{every: '-160513s', offset: '-160513s'}], timezone: timezone}, firstUserMessage: firstUserMessage, firstUserMessageData: {}, overlapPolicy: OVERLAP_POLICY_UNSPECIFIED, systemPromptData: {}, variationId: agentvar_01HXKD2E5NQM3T9AYWCF32BSPP}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentsSchedulesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:schedules", "create",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--agent-id", "agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.schedule", "{calendars: [{comment: comment, dayOfMonth: [{end: 0, start: 0, step: 0}], dayOfWeek: [{end: 0, start: 0, step: 0}], hour: [{end: 0, start: 0, step: 0}], minute: [{end: 0, start: 0, step: 0}], month: [{end: 0, start: 0, step: 0}], second: [{end: 0, start: 0, step: 0}]}], intervals: [{every: '-160513s', offset: '-160513s'}], timezone: timezone}",
			"--spec.first-user-message", "firstUserMessage",
			"--spec.first-user-message-data", "{}",
			"--spec.overlap-policy", "OVERLAP_POLICY_UNSPECIFIED",
			"--spec.system-prompt-data", "{}",
			"--spec.variation-id", "agentvar_01HXKD2E5NQM3T9AYWCF32BSPP",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"metadata:\n" +
			"  name: name\n" +
			"  externalId: externalId\n" +
			"  labels:\n" +
			"    foo: string\n" +
			"spec:\n" +
			"  schedule:\n" +
			"    calendars:\n" +
			"      - comment: comment\n" +
			"        dayOfMonth:\n" +
			"          - end: 0\n" +
			"            start: 0\n" +
			"            step: 0\n" +
			"        dayOfWeek:\n" +
			"          - end: 0\n" +
			"            start: 0\n" +
			"            step: 0\n" +
			"        hour:\n" +
			"          - end: 0\n" +
			"            start: 0\n" +
			"            step: 0\n" +
			"        minute:\n" +
			"          - end: 0\n" +
			"            start: 0\n" +
			"            step: 0\n" +
			"        month:\n" +
			"          - end: 0\n" +
			"            start: 0\n" +
			"            step: 0\n" +
			"        second:\n" +
			"          - end: 0\n" +
			"            start: 0\n" +
			"            step: 0\n" +
			"    intervals:\n" +
			"      - every: '-160513s'\n" +
			"        offset: '-160513s'\n" +
			"    timezone: timezone\n" +
			"  firstUserMessage: firstUserMessage\n" +
			"  firstUserMessageData: {}\n" +
			"  overlapPolicy: OVERLAP_POLICY_UNSPECIFIED\n" +
			"  systemPromptData: {}\n" +
			"  variationId: agentvar_01HXKD2E5NQM3T9AYWCF32BSPP\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents:schedules", "create",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--agent-id", "agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
		)
	})
}

func TestAgentsSchedulesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:schedules", "retrieve",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--agent-id", "agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
			"--id", "as_01HXKD2E5NQM3T9AYWCFMZZZBD",
		)
	})
}

func TestAgentsSchedulesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:schedules", "update",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--agent-id", "agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
			"--id", "as_01HXKD2E5NQM3T9AYWCFMZZZBD",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{schedule: {calendars: [{comment: comment, dayOfMonth: [{end: 0, start: 0, step: 0}], dayOfWeek: [{end: 0, start: 0, step: 0}], hour: [{end: 0, start: 0, step: 0}], minute: [{end: 0, start: 0, step: 0}], month: [{end: 0, start: 0, step: 0}], second: [{end: 0, start: 0, step: 0}]}], intervals: [{every: '-160513s', offset: '-160513s'}], timezone: timezone}, firstUserMessage: firstUserMessage, firstUserMessageData: {}, overlapPolicy: OVERLAP_POLICY_UNSPECIFIED, systemPromptData: {}, variationId: agentvar_01HXKD2E5NQM3T9AYWCF32BSPP}",
			"--update-mask", "updateMask",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentsSchedulesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:schedules", "update",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--agent-id", "agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
			"--id", "as_01HXKD2E5NQM3T9AYWCFMZZZBD",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.schedule", "{calendars: [{comment: comment, dayOfMonth: [{end: 0, start: 0, step: 0}], dayOfWeek: [{end: 0, start: 0, step: 0}], hour: [{end: 0, start: 0, step: 0}], minute: [{end: 0, start: 0, step: 0}], month: [{end: 0, start: 0, step: 0}], second: [{end: 0, start: 0, step: 0}]}], intervals: [{every: '-160513s', offset: '-160513s'}], timezone: timezone}",
			"--spec.first-user-message", "firstUserMessage",
			"--spec.first-user-message-data", "{}",
			"--spec.overlap-policy", "OVERLAP_POLICY_UNSPECIFIED",
			"--spec.system-prompt-data", "{}",
			"--spec.variation-id", "agentvar_01HXKD2E5NQM3T9AYWCF32BSPP",
			"--update-mask", "updateMask",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"metadata:\n" +
			"  name: name\n" +
			"  externalId: externalId\n" +
			"  labels:\n" +
			"    foo: string\n" +
			"spec:\n" +
			"  schedule:\n" +
			"    calendars:\n" +
			"      - comment: comment\n" +
			"        dayOfMonth:\n" +
			"          - end: 0\n" +
			"            start: 0\n" +
			"            step: 0\n" +
			"        dayOfWeek:\n" +
			"          - end: 0\n" +
			"            start: 0\n" +
			"            step: 0\n" +
			"        hour:\n" +
			"          - end: 0\n" +
			"            start: 0\n" +
			"            step: 0\n" +
			"        minute:\n" +
			"          - end: 0\n" +
			"            start: 0\n" +
			"            step: 0\n" +
			"        month:\n" +
			"          - end: 0\n" +
			"            start: 0\n" +
			"            step: 0\n" +
			"        second:\n" +
			"          - end: 0\n" +
			"            start: 0\n" +
			"            step: 0\n" +
			"    intervals:\n" +
			"      - every: '-160513s'\n" +
			"        offset: '-160513s'\n" +
			"    timezone: timezone\n" +
			"  firstUserMessage: firstUserMessage\n" +
			"  firstUserMessageData: {}\n" +
			"  overlapPolicy: OVERLAP_POLICY_UNSPECIFIED\n" +
			"  systemPromptData: {}\n" +
			"  variationId: agentvar_01HXKD2E5NQM3T9AYWCF32BSPP\n" +
			"updateMask: updateMask\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents:schedules", "update",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--agent-id", "agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
			"--id", "as_01HXKD2E5NQM3T9AYWCFMZZZBD",
		)
	})
}

func TestAgentsSchedulesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:schedules", "list",
			"--max-items", "10",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--agent-id", "agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
			"--cursor", "cursor",
			"--include-info=true",
			"--labels", "labels",
			"--limit", "0",
			"--prefix", "prefix",
			"--query", "query",
			"--sort-order", "sortOrder",
		)
	})
}

func TestAgentsSchedulesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:schedules", "delete",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--agent-id", "agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
			"--id", "as_01HXKD2E5NQM3T9AYWCFMZZZBD",
		)
	})
}

func TestAgentsSchedulesArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:schedules", "archive",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--agent-id", "agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
			"--id", "as_01HXKD2E5NQM3T9AYWCFMZZZBD",
		)
	})
}

func TestAgentsSchedulesPause(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:schedules", "pause",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--agent-id", "agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
			"--id", "as_01HXKD2E5NQM3T9AYWCFMZZZBD",
		)
	})
}

func TestAgentsSchedulesResume(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:schedules", "resume",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--agent-id", "agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
			"--id", "as_01HXKD2E5NQM3T9AYWCFMZZZBD",
		)
	})
}
