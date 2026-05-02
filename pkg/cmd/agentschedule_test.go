// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

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
			"--agent-id", "agentId",
			"--metadata", "{name: name, bundleKey: bundleKey, externalId: externalId, labels: {foo: string}}",
			"--spec", "{initialMessage: initialMessage, schedule: {calendars: [{comment: comment, dayOfMonth: [{end: 0, start: 0, step: 0}], dayOfWeek: [{end: 0, start: 0, step: 0}], hour: [{end: 0, start: 0, step: 0}], minute: [{end: 0, start: 0, step: 0}], month: [{end: 0, start: 0, step: 0}], second: [{end: 0, start: 0, step: 0}]}], intervals: [{every: '-160513s', offset: '-160513s'}], timezone: timezone}, data: {}, overlapPolicy: OVERLAP_POLICY_UNSPECIFIED, status: AGENT_SCHEDULE_STATUS_UNSPECIFIED, variationId: variationId}",
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
			"--agent-id", "agentId",
			"--metadata.name", "name",
			"--metadata.bundle-key", "bundleKey",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.initial-message", "initialMessage",
			"--spec.schedule", "{calendars: [{comment: comment, dayOfMonth: [{end: 0, start: 0, step: 0}], dayOfWeek: [{end: 0, start: 0, step: 0}], hour: [{end: 0, start: 0, step: 0}], minute: [{end: 0, start: 0, step: 0}], month: [{end: 0, start: 0, step: 0}], second: [{end: 0, start: 0, step: 0}]}], intervals: [{every: '-160513s', offset: '-160513s'}], timezone: timezone}",
			"--spec.data", "{}",
			"--spec.overlap-policy", "OVERLAP_POLICY_UNSPECIFIED",
			"--spec.status", "AGENT_SCHEDULE_STATUS_UNSPECIFIED",
			"--spec.variation-id", "variationId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"metadata:\n" +
			"  name: name\n" +
			"  bundleKey: bundleKey\n" +
			"  externalId: externalId\n" +
			"  labels:\n" +
			"    foo: string\n" +
			"spec:\n" +
			"  initialMessage: initialMessage\n" +
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
			"  data: {}\n" +
			"  overlapPolicy: OVERLAP_POLICY_UNSPECIFIED\n" +
			"  status: AGENT_SCHEDULE_STATUS_UNSPECIFIED\n" +
			"  variationId: variationId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents:schedules", "create",
			"--agent-id", "agentId",
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
			"--agent-id", "agentId",
			"--id", "id",
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
			"--agent-id", "agentId",
			"--id", "id",
			"--metadata", "{name: name, bundleKey: bundleKey, externalId: externalId, labels: {foo: string}}",
			"--spec", "{initialMessage: initialMessage, schedule: {calendars: [{comment: comment, dayOfMonth: [{end: 0, start: 0, step: 0}], dayOfWeek: [{end: 0, start: 0, step: 0}], hour: [{end: 0, start: 0, step: 0}], minute: [{end: 0, start: 0, step: 0}], month: [{end: 0, start: 0, step: 0}], second: [{end: 0, start: 0, step: 0}]}], intervals: [{every: '-160513s', offset: '-160513s'}], timezone: timezone}, data: {}, overlapPolicy: OVERLAP_POLICY_UNSPECIFIED, status: AGENT_SCHEDULE_STATUS_UNSPECIFIED, variationId: variationId}",
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
			"--agent-id", "agentId",
			"--id", "id",
			"--metadata.name", "name",
			"--metadata.bundle-key", "bundleKey",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.initial-message", "initialMessage",
			"--spec.schedule", "{calendars: [{comment: comment, dayOfMonth: [{end: 0, start: 0, step: 0}], dayOfWeek: [{end: 0, start: 0, step: 0}], hour: [{end: 0, start: 0, step: 0}], minute: [{end: 0, start: 0, step: 0}], month: [{end: 0, start: 0, step: 0}], second: [{end: 0, start: 0, step: 0}]}], intervals: [{every: '-160513s', offset: '-160513s'}], timezone: timezone}",
			"--spec.data", "{}",
			"--spec.overlap-policy", "OVERLAP_POLICY_UNSPECIFIED",
			"--spec.status", "AGENT_SCHEDULE_STATUS_UNSPECIFIED",
			"--spec.variation-id", "variationId",
			"--update-mask", "updateMask",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"metadata:\n" +
			"  name: name\n" +
			"  bundleKey: bundleKey\n" +
			"  externalId: externalId\n" +
			"  labels:\n" +
			"    foo: string\n" +
			"spec:\n" +
			"  initialMessage: initialMessage\n" +
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
			"  data: {}\n" +
			"  overlapPolicy: OVERLAP_POLICY_UNSPECIFIED\n" +
			"  status: AGENT_SCHEDULE_STATUS_UNSPECIFIED\n" +
			"  variationId: variationId\n" +
			"updateMask: updateMask\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents:schedules", "update",
			"--agent-id", "agentId",
			"--id", "id",
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
			"--agent-id", "agentId",
			"--bundle-key", "bundleKey",
			"--cursor", "cursor",
			"--include-info=true",
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
			"--agent-id", "agentId",
			"--id", "id",
		)
	})
}
