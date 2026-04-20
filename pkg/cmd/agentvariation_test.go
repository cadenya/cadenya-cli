// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
)

func TestAgentsVariationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:variations", "create",
			"--agent-id", "agentId",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{compactionConfig: {summarization: {instructions: instructions}, toolResultClearing: {preserveRecentResults: 0}, triggerThreshold: 0}, constraints: {maxSubObjectives: 0, maxToolCalls: 0}, description: description, enableEpisodicMemory: true, episodicMemoryTtl: 0, modelConfig: {modelId: modelId, temperature: 0}, prompt: prompt, toolSelection: {assignedTools: {allowDiscovery: true}, autoDiscovery: {hints: [string], maxTools: 0}}, weight: 0}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentsVariationsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:variations", "create",
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
			"agents:variations", "create",
			"--agent-id", "agentId",
		)
	})
}

func TestAgentsVariationsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:variations", "retrieve",
			"--agent-id", "agentId",
			"--id", "id",
		)
	})
}

func TestAgentsVariationsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:variations", "update",
			"--agent-id", "agentId",
			"--id", "id",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{compactionConfig: {summarization: {instructions: instructions}, toolResultClearing: {preserveRecentResults: 0}, triggerThreshold: 0}, constraints: {maxSubObjectives: 0, maxToolCalls: 0}, description: description, enableEpisodicMemory: true, episodicMemoryTtl: 0, modelConfig: {modelId: modelId, temperature: 0}, prompt: prompt, toolSelection: {assignedTools: {allowDiscovery: true}, autoDiscovery: {hints: [string], maxTools: 0}}, weight: 0}",
			"--update-mask", "updateMask",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(agentsVariationsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:variations", "update",
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
			"agents:variations", "update",
			"--agent-id", "agentId",
			"--id", "id",
		)
	})
}

func TestAgentsVariationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:variations", "list",
			"--max-items", "10",
			"--agent-id", "agentId",
			"--cursor", "cursor",
			"--include-info=true",
			"--limit", "0",
			"--sort-order", "sortOrder",
		)
	})
}

func TestAgentsVariationsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:variations", "delete",
			"--agent-id", "agentId",
			"--id", "id",
		)
	})
}

func TestAgentsVariationsAddAssignment(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:variations", "add-assignment",
			"--agent-id", "agentId",
			"--variation-id", "variationId",
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
			"agents:variations", "add-assignment",
			"--agent-id", "agentId",
			"--variation-id", "variationId",
		)
	})
}

func TestAgentsVariationsAddMemoryLayer(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:variations", "add-memory-layer",
			"--agent-id", "agentId",
			"--variation-id", "variationId",
			"--memory-layer-id", "memoryLayerId",
			"--position", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"memoryLayerId: memoryLayerId\n" +
			"position: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents:variations", "add-memory-layer",
			"--agent-id", "agentId",
			"--variation-id", "variationId",
		)
	})
}

func TestAgentsVariationsRemoveAssignment(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:variations", "remove-assignment",
			"--agent-id", "agentId",
			"--variation-id", "variationId",
			"--id", "id",
		)
	})
}

func TestAgentsVariationsRemoveMemoryLayer(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:variations", "remove-memory-layer",
			"--agent-id", "agentId",
			"--variation-id", "variationId",
			"--id", "id",
		)
	})
}

func TestAgentsVariationsUpdateMemoryLayer(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"agents:variations", "update-memory-layer",
			"--agent-id", "agentId",
			"--variation-id", "variationId",
			"--id", "id",
			"--position", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("position: 0")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents:variations", "update-memory-layer",
			"--agent-id", "agentId",
			"--variation-id", "variationId",
			"--id", "id",
		)
	})
}
