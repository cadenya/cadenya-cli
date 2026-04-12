// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
)

func TestAgentVariationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-variations", "create",
			"--agent-id", "agentId",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{compactionConfig: {summarization: {instructions: instructions}, toolResultClearing: {preserveRecentResults: 0}, triggerThreshold: 0}, constraints: {maxSubObjectives: 0, maxToolCalls: 0}, description: description, enableEpisodicMemory: true, episodicMemoryTtl: 0, modelConfig: {modelId: modelId, temperature: 0}, prompt: prompt, toolSelection: {assignedTools: {allowDiscovery: true}, autoDiscovery: {hints: [string], maxTools: 0}}, weight: 0}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentVariationsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-variations", "create",
			"--agent-id", "agentId",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.compaction-config", "{summarization: {instructions: instructions}, toolResultClearing: {preserveRecentResults: 0}, triggerThreshold: 0}",
			"--spec.constraints", "{maxSubObjectives: 0, maxToolCalls: 0}",
			"--spec.description", "description",
			"--spec.enable-episodic-memory=true",
			"--spec.episodic-memory-ttl", "0",
			"--spec.model-config", "{modelId: modelId, temperature: 0}",
			"--spec.prompt", "prompt",
			"--spec.tool-selection", "{assignedTools: {allowDiscovery: true}, autoDiscovery: {hints: [string], maxTools: 0}}",
			"--spec.weight", "0",
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
			"  compactionConfig:\n" +
			"    summarization:\n" +
			"      instructions: instructions\n" +
			"    toolResultClearing:\n" +
			"      preserveRecentResults: 0\n" +
			"    triggerThreshold: 0\n" +
			"  constraints:\n" +
			"    maxSubObjectives: 0\n" +
			"    maxToolCalls: 0\n" +
			"  description: description\n" +
			"  enableEpisodicMemory: true\n" +
			"  episodicMemoryTtl: 0\n" +
			"  modelConfig:\n" +
			"    modelId: modelId\n" +
			"    temperature: 0\n" +
			"  prompt: prompt\n" +
			"  toolSelection:\n" +
			"    assignedTools:\n" +
			"      allowDiscovery: true\n" +
			"    autoDiscovery:\n" +
			"      hints:\n" +
			"        - string\n" +
			"      maxTools: 0\n" +
			"  weight: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent-variations", "create",
			"--agent-id", "agentId",
		)
	})
}

func TestAgentVariationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-variations", "retrieve",
			"--agent-id", "agentId",
			"--id", "id",
		)
	})
}

func TestAgentVariationsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-variations", "update",
			"--agent-id", "agentId",
			"--id", "id",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{compactionConfig: {summarization: {instructions: instructions}, toolResultClearing: {preserveRecentResults: 0}, triggerThreshold: 0}, constraints: {maxSubObjectives: 0, maxToolCalls: 0}, description: description, enableEpisodicMemory: true, episodicMemoryTtl: 0, modelConfig: {modelId: modelId, temperature: 0}, prompt: prompt, toolSelection: {assignedTools: {allowDiscovery: true}, autoDiscovery: {hints: [string], maxTools: 0}}, weight: 0}",
			"--update-mask", "updateMask",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentVariationsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-variations", "update",
			"--agent-id", "agentId",
			"--id", "id",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.compaction-config", "{summarization: {instructions: instructions}, toolResultClearing: {preserveRecentResults: 0}, triggerThreshold: 0}",
			"--spec.constraints", "{maxSubObjectives: 0, maxToolCalls: 0}",
			"--spec.description", "description",
			"--spec.enable-episodic-memory=true",
			"--spec.episodic-memory-ttl", "0",
			"--spec.model-config", "{modelId: modelId, temperature: 0}",
			"--spec.prompt", "prompt",
			"--spec.tool-selection", "{assignedTools: {allowDiscovery: true}, autoDiscovery: {hints: [string], maxTools: 0}}",
			"--spec.weight", "0",
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
			"  compactionConfig:\n" +
			"    summarization:\n" +
			"      instructions: instructions\n" +
			"    toolResultClearing:\n" +
			"      preserveRecentResults: 0\n" +
			"    triggerThreshold: 0\n" +
			"  constraints:\n" +
			"    maxSubObjectives: 0\n" +
			"    maxToolCalls: 0\n" +
			"  description: description\n" +
			"  enableEpisodicMemory: true\n" +
			"  episodicMemoryTtl: 0\n" +
			"  modelConfig:\n" +
			"    modelId: modelId\n" +
			"    temperature: 0\n" +
			"  prompt: prompt\n" +
			"  toolSelection:\n" +
			"    assignedTools:\n" +
			"      allowDiscovery: true\n" +
			"    autoDiscovery:\n" +
			"      hints:\n" +
			"        - string\n" +
			"      maxTools: 0\n" +
			"  weight: 0\n" +
			"updateMask: updateMask\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent-variations", "update",
			"--agent-id", "agentId",
			"--id", "id",
		)
	})
}

func TestAgentVariationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-variations", "list",
			"--max-items", "10",
			"--agent-id", "agentId",
			"--cursor", "cursor",
			"--include-info=true",
			"--limit", "0",
			"--sort-order", "sortOrder",
		)
	})
}

func TestAgentVariationsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-variations", "delete",
			"--agent-id", "agentId",
			"--id", "id",
		)
	})
}

func TestAgentVariationsAddAssignment(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-variations", "add-assignment",
			"--agent-variation-id", "agentVariationId",
			"--sub-agent-id", "subAgentId",
			"--tool-id", "toolId",
			"--tool-set-id", "toolSetId",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"subAgentId: subAgentId\n" +
			"toolId: toolId\n" +
			"toolSetId: toolSetId\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agent-variations", "add-assignment",
			"--agent-variation-id", "agentVariationId",
		)
	})
}

func TestAgentVariationsRemoveAssignment(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agent-variations", "remove-assignment",
			"--agent-variation-id", "agentVariationId",
			"--id", "id",
		)
	})
}
