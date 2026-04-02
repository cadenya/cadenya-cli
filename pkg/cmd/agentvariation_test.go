// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cli/internal/mocktest"
	"github.com/cadenya/cli/internal/requestflag"
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
			"--spec", "{agentDocuments: [{documentId: documentId, documentMetadata: {name: name, externalId: externalId, labels: {foo: string}}, documentNamespaceId: documentNamespaceId, documentNamespaceMetadata: {name: name, externalId: externalId, labels: {foo: string}}}], agentTools: [{agentId: agentId, agentMetadata: {name: name, externalId: externalId, labels: {foo: string}}, toolId: toolId, toolMetadata: {name: name, externalId: externalId, labels: {foo: string}}, toolSetId: toolSetId, toolSetMetadata: {name: name, externalId: externalId, labels: {foo: string}}}], constraints: {maxSubObjectives: 0, maxToolCalls: 0}, description: description, enableEpisodicMemory: true, episodicMemoryTtl: 0, modelConfig: {modelId: modelId, temperature: 0}, prompt: prompt, toolSelection: {assignedTools: {allowDiscovery: true}, autoDiscovery: {hints: [string], maxTools: 0}}, weight: 0}",
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
			"--spec.agent-documents", "[{documentId: documentId, documentMetadata: {name: name, externalId: externalId, labels: {foo: string}}, documentNamespaceId: documentNamespaceId, documentNamespaceMetadata: {name: name, externalId: externalId, labels: {foo: string}}}]",
			"--spec.agent-tools", "[{agentId: agentId, agentMetadata: {name: name, externalId: externalId, labels: {foo: string}}, toolId: toolId, toolMetadata: {name: name, externalId: externalId, labels: {foo: string}}, toolSetId: toolSetId, toolSetMetadata: {name: name, externalId: externalId, labels: {foo: string}}}]",
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
			"  agentDocuments:\n" +
			"    - documentId: documentId\n" +
			"      documentMetadata:\n" +
			"        name: name\n" +
			"        externalId: externalId\n" +
			"        labels:\n" +
			"          foo: string\n" +
			"      documentNamespaceId: documentNamespaceId\n" +
			"      documentNamespaceMetadata:\n" +
			"        name: name\n" +
			"        externalId: externalId\n" +
			"        labels:\n" +
			"          foo: string\n" +
			"  agentTools:\n" +
			"    - agentId: agentId\n" +
			"      agentMetadata:\n" +
			"        name: name\n" +
			"        externalId: externalId\n" +
			"        labels:\n" +
			"          foo: string\n" +
			"      toolId: toolId\n" +
			"      toolMetadata:\n" +
			"        name: name\n" +
			"        externalId: externalId\n" +
			"        labels:\n" +
			"          foo: string\n" +
			"      toolSetId: toolSetId\n" +
			"      toolSetMetadata:\n" +
			"        name: name\n" +
			"        externalId: externalId\n" +
			"        labels:\n" +
			"          foo: string\n" +
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
			"--spec", "{agentDocuments: [{documentId: documentId, documentMetadata: {name: name, externalId: externalId, labels: {foo: string}}, documentNamespaceId: documentNamespaceId, documentNamespaceMetadata: {name: name, externalId: externalId, labels: {foo: string}}}], agentTools: [{agentId: agentId, agentMetadata: {name: name, externalId: externalId, labels: {foo: string}}, toolId: toolId, toolMetadata: {name: name, externalId: externalId, labels: {foo: string}}, toolSetId: toolSetId, toolSetMetadata: {name: name, externalId: externalId, labels: {foo: string}}}], constraints: {maxSubObjectives: 0, maxToolCalls: 0}, description: description, enableEpisodicMemory: true, episodicMemoryTtl: 0, modelConfig: {modelId: modelId, temperature: 0}, prompt: prompt, toolSelection: {assignedTools: {allowDiscovery: true}, autoDiscovery: {hints: [string], maxTools: 0}}, weight: 0}",
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
			"--spec.agent-documents", "[{documentId: documentId, documentMetadata: {name: name, externalId: externalId, labels: {foo: string}}, documentNamespaceId: documentNamespaceId, documentNamespaceMetadata: {name: name, externalId: externalId, labels: {foo: string}}}]",
			"--spec.agent-tools", "[{agentId: agentId, agentMetadata: {name: name, externalId: externalId, labels: {foo: string}}, toolId: toolId, toolMetadata: {name: name, externalId: externalId, labels: {foo: string}}, toolSetId: toolSetId, toolSetMetadata: {name: name, externalId: externalId, labels: {foo: string}}}]",
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
			"  agentDocuments:\n" +
			"    - documentId: documentId\n" +
			"      documentMetadata:\n" +
			"        name: name\n" +
			"        externalId: externalId\n" +
			"        labels:\n" +
			"          foo: string\n" +
			"      documentNamespaceId: documentNamespaceId\n" +
			"      documentNamespaceMetadata:\n" +
			"        name: name\n" +
			"        externalId: externalId\n" +
			"        labels:\n" +
			"          foo: string\n" +
			"  agentTools:\n" +
			"    - agentId: agentId\n" +
			"      agentMetadata:\n" +
			"        name: name\n" +
			"        externalId: externalId\n" +
			"        labels:\n" +
			"          foo: string\n" +
			"      toolId: toolId\n" +
			"      toolMetadata:\n" +
			"        name: name\n" +
			"        externalId: externalId\n" +
			"        labels:\n" +
			"          foo: string\n" +
			"      toolSetId: toolSetId\n" +
			"      toolSetMetadata:\n" +
			"        name: name\n" +
			"        externalId: externalId\n" +
			"        labels:\n" +
			"          foo: string\n" +
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
