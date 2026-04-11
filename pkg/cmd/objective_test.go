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
			"--agent-id", "agentId",
			"--data", "{data: {}, initialMessage: initialMessage, secrets: [{name: name, value: value}]}",
			"--metadata", "{externalId: externalId, labels: {foo: string}}",
			"--variation-id", "variationId",
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
			"--agent-id", "agentId",
			"--data.data", "{}",
			"--data.initial-message", "initialMessage",
			"--data.secrets", "[{name: name, value: value}]",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--variation-id", "variationId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"agentId: agentId\n" +
			"data:\n" +
			"  data: {}\n" +
			"  initialMessage: initialMessage\n" +
			"  secrets:\n" +
			"    - name: name\n" +
			"      value: value\n" +
			"metadata:\n" +
			"  externalId: externalId\n" +
			"  labels:\n" +
			"    foo: string\n" +
			"variationId: variationId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"objectives", "create",
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
			"--id", "id",
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
			"--agent-id", "agentId",
			"--cursor", "cursor",
			"--include-info=true",
			"--limit", "0",
			"--parent-objective-id", "parentObjectiveId",
			"--profile-id", "profileId",
			"--sort-order", "sortOrder",
			"--state", "STATE_UNSPECIFIED",
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
			"--objective-id", "objectiveId",
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
			"--objective-id", "objectiveId",
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
			"--objective-id", "objectiveId",
			"--compaction-config", "{summarization: {instructions: instructions, minPreserveTurns: 0}, toolResultClearing: {preserveRecentResults: 0}, triggerThreshold: 0}",
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
			"--objective-id", "objectiveId",
			"--compaction-config.summarization", "{instructions: instructions, minPreserveTurns: 0}",
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
			"    minPreserveTurns: 0\n" +
			"  toolResultClearing:\n" +
			"    preserveRecentResults: 0\n" +
			"  triggerThreshold: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"objectives", "compact",
			"--objective-id", "objectiveId",
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
			"--objective-id", "objectiveId",
			"--enqueue=true",
			"--message", "message",
			"--secret", "{name: name, value: value}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(objectivesContinue)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives", "continue",
			"--objective-id", "objectiveId",
			"--enqueue=true",
			"--message", "message",
			"--secret.name", "name",
			"--secret.value", "value",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"enqueue: true\n" +
			"message: message\n" +
			"secrets:\n" +
			"  - name: name\n" +
			"    value: value\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"objectives", "continue",
			"--objective-id", "objectiveId",
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
			"--objective-id", "objectiveId",
			"--cursor", "cursor",
			"--include-info=true",
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
			"--objective-id", "objectiveId",
			"--cursor", "cursor",
			"--include-info=true",
			"--limit", "0",
			"--sort-order", "sortOrder",
			"--window-id", "windowId",
		)
	})
}
