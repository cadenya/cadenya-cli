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
			"--workspace-id", "workspaceId",
			"--agent-id", "agentId",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{compactionConfig: {summarization: {instructions: instructions}, toolResultClearing: {preserveRecentResults: 0}, triggerThreshold: 0}, constraints: {inactivityTimeout: '-160513s', maxSubObjectives: 0, maxToolCalls: 0}, description: description, firstUserMessageTemplate: firstUserMessageTemplate, modelConfig: {modelId: claude/opus-4.6, temperature: 0}, progressiveDiscovery: {hints: [string], maxTools: 0}, systemPromptTemplate: systemPromptTemplate}",
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
			"--workspace-id", "workspaceId",
			"--agent-id", "agentId",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.compaction-config", "{summarization: {instructions: instructions}, toolResultClearing: {preserveRecentResults: 0}, triggerThreshold: 0}",
			"--spec.constraints", "{inactivityTimeout: '-160513s', maxSubObjectives: 0, maxToolCalls: 0}",
			"--spec.description", "description",
			"--spec.first-user-message-template", "firstUserMessageTemplate",
			"--spec.model-config", "{modelId: claude/opus-4.6, temperature: 0}",
			"--spec.progressive-discovery", "{hints: [string], maxTools: 0}",
			"--spec.system-prompt-template", "systemPromptTemplate",
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
			"    inactivityTimeout: '-160513s'\n" +
			"    maxSubObjectives: 0\n" +
			"    maxToolCalls: 0\n" +
			"  description: description\n" +
			"  firstUserMessageTemplate: firstUserMessageTemplate\n" +
			"  modelConfig:\n" +
			"    modelId: claude/opus-4.6\n" +
			"    temperature: 0\n" +
			"  progressiveDiscovery:\n" +
			"    hints:\n" +
			"      - string\n" +
			"    maxTools: 0\n" +
			"  systemPromptTemplate: systemPromptTemplate\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents:variations", "create",
			"--workspace-id", "workspaceId",
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
			"--workspace-id", "workspaceId",
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
			"--workspace-id", "workspaceId",
			"--agent-id", "agentId",
			"--id", "id",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{compactionConfig: {summarization: {instructions: instructions}, toolResultClearing: {preserveRecentResults: 0}, triggerThreshold: 0}, constraints: {inactivityTimeout: '-160513s', maxSubObjectives: 0, maxToolCalls: 0}, description: description, firstUserMessageTemplate: firstUserMessageTemplate, modelConfig: {modelId: claude/opus-4.6, temperature: 0}, progressiveDiscovery: {hints: [string], maxTools: 0}, systemPromptTemplate: systemPromptTemplate}",
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
			"--workspace-id", "workspaceId",
			"--agent-id", "agentId",
			"--id", "id",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.compaction-config", "{summarization: {instructions: instructions}, toolResultClearing: {preserveRecentResults: 0}, triggerThreshold: 0}",
			"--spec.constraints", "{inactivityTimeout: '-160513s', maxSubObjectives: 0, maxToolCalls: 0}",
			"--spec.description", "description",
			"--spec.first-user-message-template", "firstUserMessageTemplate",
			"--spec.model-config", "{modelId: claude/opus-4.6, temperature: 0}",
			"--spec.progressive-discovery", "{hints: [string], maxTools: 0}",
			"--spec.system-prompt-template", "systemPromptTemplate",
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
			"    inactivityTimeout: '-160513s'\n" +
			"    maxSubObjectives: 0\n" +
			"    maxToolCalls: 0\n" +
			"  description: description\n" +
			"  firstUserMessageTemplate: firstUserMessageTemplate\n" +
			"  modelConfig:\n" +
			"    modelId: claude/opus-4.6\n" +
			"    temperature: 0\n" +
			"  progressiveDiscovery:\n" +
			"    hints:\n" +
			"      - string\n" +
			"    maxTools: 0\n" +
			"  systemPromptTemplate: systemPromptTemplate\n" +
			"updateMask: updateMask\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents:variations", "update",
			"--workspace-id", "workspaceId",
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
			"--workspace-id", "workspaceId",
			"--agent-id", "agentId",
			"--cursor", "cursor",
			"--include-info=true",
			"--labels", "labels",
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
			"--workspace-id", "workspaceId",
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
			"--workspace-id", "workspaceId",
			"--agent-id", "agentId",
			"--variation-id", "variationId",
			"--sub-agent-id", "agent_01HXKD2E5NQM3T9AYWCFMGWT9Y",
			"--tool-id", "tool_01HXKD2E5NQM3T9AYWCFWVYY9K",
			"--tool-set-id", "toolset_01HXKD2E5NQM3T9AYWCFNRMN74",
			"--type", "type",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"subAgentId: agent_01HXKD2E5NQM3T9AYWCFMGWT9Y\n" +
			"toolId: tool_01HXKD2E5NQM3T9AYWCFWVYY9K\n" +
			"toolSetId: toolset_01HXKD2E5NQM3T9AYWCFNRMN74\n" +
			"type: type\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents:variations", "add-assignment",
			"--workspace-id", "workspaceId",
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
			"--workspace-id", "workspaceId",
			"--agent-id", "agentId",
			"--variation-id", "variationId",
			"--memory-layer-id", "memlyr_01HXKD2E5NQM3T9AYWCFFFBMJH",
			"--position", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"memoryLayerId: memlyr_01HXKD2E5NQM3T9AYWCFFFBMJH\n" +
			"position: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"agents:variations", "add-memory-layer",
			"--workspace-id", "workspaceId",
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
			"--workspace-id", "workspaceId",
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
			"--workspace-id", "workspaceId",
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
			"--workspace-id", "workspaceId",
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
			"--workspace-id", "workspaceId",
			"--agent-id", "agentId",
			"--variation-id", "variationId",
			"--id", "id",
		)
	})
}
