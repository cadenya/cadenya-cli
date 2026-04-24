// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
)

func TestToolSetsToolsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets:tools", "create",
			"--tool-set-id", "toolSetId",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{config: {http: {requestMethod: HTTP_METHOD_UNSPECIFIED, headers: {foo: string}, path: path, query: query, requestBodyContentType: requestBodyContentType, requestBodyTemplate: requestBodyTemplate, toolName: toolName}, mcp: {toolDescription: toolDescription, toolName: toolName, toolTitle: toolTitle}}, description: description, parameters: {foo: bar}, status: TOOL_STATUS_UNSPECIFIED, requiresApproval: true}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(toolSetsToolsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets:tools", "create",
			"--tool-set-id", "toolSetId",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.config", "{http: {requestMethod: HTTP_METHOD_UNSPECIFIED, headers: {foo: string}, path: path, query: query, requestBodyContentType: requestBodyContentType, requestBodyTemplate: requestBodyTemplate, toolName: toolName}, mcp: {toolDescription: toolDescription, toolName: toolName, toolTitle: toolTitle}}",
			"--spec.description", "description",
			"--spec.parameters", "{foo: bar}",
			"--spec.status", "TOOL_STATUS_UNSPECIFIED",
			"--spec.requires-approval=true",
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
			"  config:\n" +
			"    http:\n" +
			"      requestMethod: HTTP_METHOD_UNSPECIFIED\n" +
			"      headers:\n" +
			"        foo: string\n" +
			"      path: path\n" +
			"      query: query\n" +
			"      requestBodyContentType: requestBodyContentType\n" +
			"      requestBodyTemplate: requestBodyTemplate\n" +
			"      toolName: toolName\n" +
			"    mcp:\n" +
			"      toolDescription: toolDescription\n" +
			"      toolName: toolName\n" +
			"      toolTitle: toolTitle\n" +
			"  description: description\n" +
			"  parameters:\n" +
			"    foo: bar\n" +
			"  status: TOOL_STATUS_UNSPECIFIED\n" +
			"  requiresApproval: true\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tool-sets:tools", "create",
			"--tool-set-id", "toolSetId",
		)
	})
}

func TestToolSetsToolsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets:tools", "retrieve",
			"--tool-set-id", "toolSetId",
			"--id", "id",
		)
	})
}

func TestToolSetsToolsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets:tools", "update",
			"--tool-set-id", "toolSetId",
			"--id", "id",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{config: {http: {requestMethod: HTTP_METHOD_UNSPECIFIED, headers: {foo: string}, path: path, query: query, requestBodyContentType: requestBodyContentType, requestBodyTemplate: requestBodyTemplate, toolName: toolName}, mcp: {toolDescription: toolDescription, toolName: toolName, toolTitle: toolTitle}}, description: description, parameters: {foo: bar}, status: TOOL_STATUS_UNSPECIFIED, requiresApproval: true}",
			"--update-mask", "updateMask",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(toolSetsToolsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets:tools", "update",
			"--tool-set-id", "toolSetId",
			"--id", "id",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.config", "{http: {requestMethod: HTTP_METHOD_UNSPECIFIED, headers: {foo: string}, path: path, query: query, requestBodyContentType: requestBodyContentType, requestBodyTemplate: requestBodyTemplate, toolName: toolName}, mcp: {toolDescription: toolDescription, toolName: toolName, toolTitle: toolTitle}}",
			"--spec.description", "description",
			"--spec.parameters", "{foo: bar}",
			"--spec.status", "TOOL_STATUS_UNSPECIFIED",
			"--spec.requires-approval=true",
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
			"  config:\n" +
			"    http:\n" +
			"      requestMethod: HTTP_METHOD_UNSPECIFIED\n" +
			"      headers:\n" +
			"        foo: string\n" +
			"      path: path\n" +
			"      query: query\n" +
			"      requestBodyContentType: requestBodyContentType\n" +
			"      requestBodyTemplate: requestBodyTemplate\n" +
			"      toolName: toolName\n" +
			"    mcp:\n" +
			"      toolDescription: toolDescription\n" +
			"      toolName: toolName\n" +
			"      toolTitle: toolTitle\n" +
			"  description: description\n" +
			"  parameters:\n" +
			"    foo: bar\n" +
			"  status: TOOL_STATUS_UNSPECIFIED\n" +
			"  requiresApproval: true\n" +
			"updateMask: updateMask\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tool-sets:tools", "update",
			"--tool-set-id", "toolSetId",
			"--id", "id",
		)
	})
}

func TestToolSetsToolsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets:tools", "list",
			"--max-items", "10",
			"--tool-set-id", "toolSetId",
			"--cursor", "cursor",
			"--include-info=true",
			"--limit", "0",
			"--prefix", "prefix",
			"--query", "query",
			"--sort-order", "sortOrder",
		)
	})
}

func TestToolSetsToolsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets:tools", "delete",
			"--tool-set-id", "toolSetId",
			"--id", "id",
		)
	})
}
