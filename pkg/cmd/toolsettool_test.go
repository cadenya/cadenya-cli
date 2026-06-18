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
			"--workspace-id", "workspaceId",
			"--tool-set-id", "toolSetId",
			"--metadata", "{name: name, bundleKey: bundleKey, externalId: externalId, labels: {foo: string}}",
			"--spec", "{config: {http: {requestMethod: HTTP_METHOD_UNSPECIFIED, headers: {foo: string}, path: path, query: query, requestBodyContentType: requestBodyContentType, requestBodyTemplate: requestBodyTemplate}, mcp: {}, openapi: {method: method, path: path}}, description: description, parameters: {foo: bar}, requiresApproval: true, llmToolName: llmToolName}",
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
			"--workspace-id", "workspaceId",
			"--tool-set-id", "toolSetId",
			"--metadata.name", "name",
			"--metadata.bundle-key", "bundleKey",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.config", "{http: {requestMethod: HTTP_METHOD_UNSPECIFIED, headers: {foo: string}, path: path, query: query, requestBodyContentType: requestBodyContentType, requestBodyTemplate: requestBodyTemplate}, mcp: {}, openapi: {method: method, path: path}}",
			"--spec.description", "description",
			"--spec.parameters", "{foo: bar}",
			"--spec.requires-approval=true",
			"--spec.llm-tool-name", "llmToolName",
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
			"  config:\n" +
			"    http:\n" +
			"      requestMethod: HTTP_METHOD_UNSPECIFIED\n" +
			"      headers:\n" +
			"        foo: string\n" +
			"      path: path\n" +
			"      query: query\n" +
			"      requestBodyContentType: requestBodyContentType\n" +
			"      requestBodyTemplate: requestBodyTemplate\n" +
			"    mcp: {}\n" +
			"    openapi:\n" +
			"      method: method\n" +
			"      path: path\n" +
			"  description: description\n" +
			"  parameters:\n" +
			"    foo: bar\n" +
			"  requiresApproval: true\n" +
			"  llmToolName: llmToolName\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tool-sets:tools", "create",
			"--workspace-id", "workspaceId",
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
			"--workspace-id", "workspaceId",
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
			"--workspace-id", "workspaceId",
			"--tool-set-id", "toolSetId",
			"--id", "id",
			"--metadata", "{name: name, bundleKey: bundleKey, externalId: externalId, labels: {foo: string}}",
			"--spec", "{config: {http: {requestMethod: HTTP_METHOD_UNSPECIFIED, headers: {foo: string}, path: path, query: query, requestBodyContentType: requestBodyContentType, requestBodyTemplate: requestBodyTemplate}, mcp: {}, openapi: {method: method, path: path}}, description: description, parameters: {foo: bar}, requiresApproval: true, llmToolName: llmToolName}",
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
			"--workspace-id", "workspaceId",
			"--tool-set-id", "toolSetId",
			"--id", "id",
			"--metadata.name", "name",
			"--metadata.bundle-key", "bundleKey",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.config", "{http: {requestMethod: HTTP_METHOD_UNSPECIFIED, headers: {foo: string}, path: path, query: query, requestBodyContentType: requestBodyContentType, requestBodyTemplate: requestBodyTemplate}, mcp: {}, openapi: {method: method, path: path}}",
			"--spec.description", "description",
			"--spec.parameters", "{foo: bar}",
			"--spec.requires-approval=true",
			"--spec.llm-tool-name", "llmToolName",
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
			"  config:\n" +
			"    http:\n" +
			"      requestMethod: HTTP_METHOD_UNSPECIFIED\n" +
			"      headers:\n" +
			"        foo: string\n" +
			"      path: path\n" +
			"      query: query\n" +
			"      requestBodyContentType: requestBodyContentType\n" +
			"      requestBodyTemplate: requestBodyTemplate\n" +
			"    mcp: {}\n" +
			"    openapi:\n" +
			"      method: method\n" +
			"      path: path\n" +
			"  description: description\n" +
			"  parameters:\n" +
			"    foo: bar\n" +
			"  requiresApproval: true\n" +
			"  llmToolName: llmToolName\n" +
			"updateMask: updateMask\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tool-sets:tools", "update",
			"--workspace-id", "workspaceId",
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
			"--workspace-id", "workspaceId",
			"--tool-set-id", "toolSetId",
			"--bundle-key", "bundleKey",
			"--cursor", "cursor",
			"--include-info=true",
			"--limit", "0",
			"--name", "string",
			"--prefix", "prefix",
			"--query", "query",
			"--requires-approval=true",
			"--sort-order", "sortOrder",
			"--state", "STATE_UNSPECIFIED",
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
			"--workspace-id", "workspaceId",
			"--tool-set-id", "toolSetId",
			"--id", "id",
		)
	})
}

func TestToolSetsToolsOmit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets:tools", "omit",
			"--workspace-id", "workspaceId",
			"--tool-set-id", "toolSetId",
			"--id", "id",
		)
	})
}

func TestToolSetsToolsRestore(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets:tools", "restore",
			"--workspace-id", "workspaceId",
			"--tool-set-id", "toolSetId",
			"--id", "id",
		)
	})
}
