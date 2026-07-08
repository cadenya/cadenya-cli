// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
)

func TestAgentsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "create",
			"--workspace-id", "workspaceId",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{variationSelectionMode: VARIATION_SELECTION_MODE_UNSPECIFIED, description: description, enableEpisodicMemory: true, episodicMemoryTtl: 0, outputDefinition: {foo: bar}, systemPromptDataSchema: {foo: bar}, webhookEventsUrl: webhookEventsUrl}",
			"--default-variation", "{metadata: {name: name, externalId: externalId, labels: {foo: string}}, spec: {compactionConfig: {summarization: {instructions: instructions}, toolResultClearing: {preserveRecentResults: 0}, triggerThreshold: 0}, constraints: {inactivityTimeout: '-160513s', maxSubObjectives: 0, maxToolCalls: 0}, description: description, firstUserMessageTemplate: firstUserMessageTemplate, modelConfig: {modelId: modelId, temperature: 0}, progressiveDiscovery: {hints: [string], maxTools: 0, rerankThreshold: 0}, systemPromptTemplate: systemPromptTemplate, weight: 0}}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "create",
			"--workspace-id", "workspaceId",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.variation-selection-mode", "VARIATION_SELECTION_MODE_UNSPECIFIED",
			"--spec.description", "description",
			"--spec.enable-episodic-memory=true",
			"--spec.episodic-memory-ttl", "0",
			"--spec.output-definition", "{foo: bar}",
			"--spec.system-prompt-data-schema", "{foo: bar}",
			"--spec.webhook-events-url", "webhookEventsUrl",
			"--default-variation.metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--default-variation.spec", "{compactionConfig: {summarization: {instructions: instructions}, toolResultClearing: {preserveRecentResults: 0}, triggerThreshold: 0}, constraints: {inactivityTimeout: '-160513s', maxSubObjectives: 0, maxToolCalls: 0}, description: description, firstUserMessageTemplate: firstUserMessageTemplate, modelConfig: {modelId: modelId, temperature: 0}, progressiveDiscovery: {hints: [string], maxTools: 0, rerankThreshold: 0}, systemPromptTemplate: systemPromptTemplate, weight: 0}",
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
			"  variationSelectionMode: VARIATION_SELECTION_MODE_UNSPECIFIED\n" +
			"  description: description\n" +
			"  enableEpisodicMemory: true\n" +
			"  episodicMemoryTtl: 0\n" +
			"  outputDefinition:\n" +
			"    foo: bar\n" +
			"  systemPromptDataSchema:\n" +
			"    foo: bar\n" +
			"  webhookEventsUrl: webhookEventsUrl\n" +
			"defaultVariation:\n" +
			"  metadata:\n" +
			"    name: name\n" +
			"    externalId: externalId\n" +
			"    labels:\n" +
			"      foo: string\n" +
			"  spec:\n" +
			"    compactionConfig:\n" +
			"      summarization:\n" +
			"        instructions: instructions\n" +
			"      toolResultClearing:\n" +
			"        preserveRecentResults: 0\n" +
			"      triggerThreshold: 0\n" +
			"    constraints:\n" +
			"      inactivityTimeout: '-160513s'\n" +
			"      maxSubObjectives: 0\n" +
			"      maxToolCalls: 0\n" +
			"    description: description\n" +
			"    firstUserMessageTemplate: firstUserMessageTemplate\n" +
			"    modelConfig:\n" +
			"      modelId: modelId\n" +
			"      temperature: 0\n" +
			"    progressiveDiscovery:\n" +
			"      hints:\n" +
			"        - string\n" +
			"      maxTools: 0\n" +
			"      rerankThreshold: 0\n" +
			"    systemPromptTemplate: systemPromptTemplate\n" +
			"    weight: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents", "create",
			"--workspace-id", "workspaceId",
		)
	})
}

func TestAgentsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "retrieve",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}

func TestAgentsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "update",
			"--workspace-id", "workspaceId",
			"--id", "id",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{variationSelectionMode: VARIATION_SELECTION_MODE_UNSPECIFIED, description: description, enableEpisodicMemory: true, episodicMemoryTtl: 0, outputDefinition: {foo: bar}, systemPromptDataSchema: {foo: bar}, webhookEventsUrl: webhookEventsUrl}",
			"--update-mask", "updateMask",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "update",
			"--workspace-id", "workspaceId",
			"--id", "id",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.variation-selection-mode", "VARIATION_SELECTION_MODE_UNSPECIFIED",
			"--spec.description", "description",
			"--spec.enable-episodic-memory=true",
			"--spec.episodic-memory-ttl", "0",
			"--spec.output-definition", "{foo: bar}",
			"--spec.system-prompt-data-schema", "{foo: bar}",
			"--spec.webhook-events-url", "webhookEventsUrl",
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
			"  variationSelectionMode: VARIATION_SELECTION_MODE_UNSPECIFIED\n" +
			"  description: description\n" +
			"  enableEpisodicMemory: true\n" +
			"  episodicMemoryTtl: 0\n" +
			"  outputDefinition:\n" +
			"    foo: bar\n" +
			"  systemPromptDataSchema:\n" +
			"    foo: bar\n" +
			"  webhookEventsUrl: webhookEventsUrl\n" +
			"updateMask: updateMask\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents", "update",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}

func TestAgentsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "list",
			"--max-items", "10",
			"--workspace-id", "workspaceId",
			"--cursor", "cursor",
			"--include-info=true",
			"--labels", "labels",
			"--limit", "0",
			"--prefix", "prefix",
			"--query", "query",
			"--sort-order", "sortOrder",
			"--state", "STATE_UNSPECIFIED",
			"--variation-selection-mode", "VARIATION_SELECTION_MODE_UNSPECIFIED",
		)
	})
}

func TestAgentsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "delete",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}

func TestAgentsArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "archive",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}

func TestAgentsPublish(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "publish",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}

func TestAgentsUnarchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "unarchive",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}

func TestAgentsUnpublish(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents", "unpublish",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}
