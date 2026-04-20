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
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{status: AGENT_STATUS_UNSPECIFIED, variationSelectionMode: VARIATION_SELECTION_MODE_UNSPECIFIED, description: description, inputDataSchema: {}, webhookEventsUrl: webhookEventsUrl}",
			"--default-variation", "{metadata: {name: name, externalId: externalId, labels: {foo: string}}, spec: {compactionConfig: {summarization: {instructions: instructions}, toolResultClearing: {preserveRecentResults: 0}, triggerThreshold: 0}, constraints: {maxSubObjectives: 0, maxToolCalls: 0}, description: description, enableEpisodicMemory: true, episodicMemoryTtl: 0, modelConfig: {modelId: modelId, temperature: 0}, progressiveDiscovery: {hints: [string], maxTools: 0, rerankThreshold: 0}, prompt: prompt, weight: 0}}",
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
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.status", "AGENT_STATUS_UNSPECIFIED",
			"--spec.variation-selection-mode", "VARIATION_SELECTION_MODE_UNSPECIFIED",
			"--spec.description", "description",
			"--spec.input-data-schema", "{}",
			"--spec.webhook-events-url", "webhookEventsUrl",
			"--default-variation.metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--default-variation.spec", "{compactionConfig: {summarization: {instructions: instructions}, toolResultClearing: {preserveRecentResults: 0}, triggerThreshold: 0}, constraints: {maxSubObjectives: 0, maxToolCalls: 0}, description: description, enableEpisodicMemory: true, episodicMemoryTtl: 0, modelConfig: {modelId: modelId, temperature: 0}, progressiveDiscovery: {hints: [string], maxTools: 0, rerankThreshold: 0}, prompt: prompt, weight: 0}",
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
			"  status: AGENT_STATUS_UNSPECIFIED\n" +
			"  variationSelectionMode: VARIATION_SELECTION_MODE_UNSPECIFIED\n" +
			"  description: description\n" +
			"  inputDataSchema: {}\n" +
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
			"      maxSubObjectives: 0\n" +
			"      maxToolCalls: 0\n" +
			"    description: description\n" +
			"    enableEpisodicMemory: true\n" +
			"    episodicMemoryTtl: 0\n" +
			"    modelConfig:\n" +
			"      modelId: modelId\n" +
			"      temperature: 0\n" +
			"    progressiveDiscovery:\n" +
			"      hints:\n" +
			"        - string\n" +
			"      maxTools: 0\n" +
			"      rerankThreshold: 0\n" +
			"    prompt: prompt\n" +
			"    weight: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents", "create",
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
			"--id", "id",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{status: AGENT_STATUS_UNSPECIFIED, variationSelectionMode: VARIATION_SELECTION_MODE_UNSPECIFIED, description: description, inputDataSchema: {}, webhookEventsUrl: webhookEventsUrl}",
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
			"--id", "id",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.status", "AGENT_STATUS_UNSPECIFIED",
			"--spec.variation-selection-mode", "VARIATION_SELECTION_MODE_UNSPECIFIED",
			"--spec.description", "description",
			"--spec.input-data-schema", "{}",
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
			"  status: AGENT_STATUS_UNSPECIFIED\n" +
			"  variationSelectionMode: VARIATION_SELECTION_MODE_UNSPECIFIED\n" +
			"  description: description\n" +
			"  inputDataSchema: {}\n" +
			"  webhookEventsUrl: webhookEventsUrl\n" +
			"updateMask: updateMask\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents", "update",
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
			"--cursor", "cursor",
			"--include-info=true",
			"--limit", "0",
			"--prefix", "prefix",
			"--sort-order", "sortOrder",
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
			"--id", "id",
		)
	})
}
