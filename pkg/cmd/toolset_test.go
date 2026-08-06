// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
)

func TestToolSetsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets", "create",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{adapter: {mcp: {excludeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {exact: exact, type: exact, caseSensitive: true}}]}, headers: {foo: string}, includeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {exact: exact, type: exact, caseSensitive: true}}]}, justInTime: {enabled: true, failObjectiveOnToolListError: true}, toolApprovals: {always: true, type: always}, url: url}, type: mcp}, description: description}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(toolSetsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets", "create",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.adapter", "{mcp: {excludeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {exact: exact, type: exact, caseSensitive: true}}]}, headers: {foo: string}, includeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {exact: exact, type: exact, caseSensitive: true}}]}, justInTime: {enabled: true, failObjectiveOnToolListError: true}, toolApprovals: {always: true, type: always}, url: url}, type: mcp}",
			"--spec.description", "description",
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
			"  adapter:\n" +
			"    mcp:\n" +
			"      excludeTools:\n" +
			"        operator: OPERATOR_UNSPECIFIED\n" +
			"        filters:\n" +
			"          - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"            matcher:\n" +
			"              exact: exact\n" +
			"              type: exact\n" +
			"              caseSensitive: true\n" +
			"      headers:\n" +
			"        foo: string\n" +
			"      includeTools:\n" +
			"        operator: OPERATOR_UNSPECIFIED\n" +
			"        filters:\n" +
			"          - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"            matcher:\n" +
			"              exact: exact\n" +
			"              type: exact\n" +
			"              caseSensitive: true\n" +
			"      justInTime:\n" +
			"        enabled: true\n" +
			"        failObjectiveOnToolListError: true\n" +
			"      toolApprovals:\n" +
			"        always: true\n" +
			"        type: always\n" +
			"      url: url\n" +
			"    type: mcp\n" +
			"  description: description\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tool-sets", "create",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
		)
	})
}

func TestToolSetsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets", "retrieve",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--id", "toolset_01HXKD2E5NQM3T9AYWCFNRMN74",
		)
	})
}

func TestToolSetsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets", "update",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--id", "toolset_01HXKD2E5NQM3T9AYWCFNRMN74",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{adapter: {mcp: {excludeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {exact: exact, type: exact, caseSensitive: true}}]}, headers: {foo: string}, includeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {exact: exact, type: exact, caseSensitive: true}}]}, justInTime: {enabled: true, failObjectiveOnToolListError: true}, toolApprovals: {always: true, type: always}, url: url}, type: mcp}, description: description}",
			"--update-mask", "updateMask",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(toolSetsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets", "update",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--id", "toolset_01HXKD2E5NQM3T9AYWCFNRMN74",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.adapter", "{mcp: {excludeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {exact: exact, type: exact, caseSensitive: true}}]}, headers: {foo: string}, includeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {exact: exact, type: exact, caseSensitive: true}}]}, justInTime: {enabled: true, failObjectiveOnToolListError: true}, toolApprovals: {always: true, type: always}, url: url}, type: mcp}",
			"--spec.description", "description",
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
			"  adapter:\n" +
			"    mcp:\n" +
			"      excludeTools:\n" +
			"        operator: OPERATOR_UNSPECIFIED\n" +
			"        filters:\n" +
			"          - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"            matcher:\n" +
			"              exact: exact\n" +
			"              type: exact\n" +
			"              caseSensitive: true\n" +
			"      headers:\n" +
			"        foo: string\n" +
			"      includeTools:\n" +
			"        operator: OPERATOR_UNSPECIFIED\n" +
			"        filters:\n" +
			"          - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"            matcher:\n" +
			"              exact: exact\n" +
			"              type: exact\n" +
			"              caseSensitive: true\n" +
			"      justInTime:\n" +
			"        enabled: true\n" +
			"        failObjectiveOnToolListError: true\n" +
			"      toolApprovals:\n" +
			"        always: true\n" +
			"        type: always\n" +
			"      url: url\n" +
			"    type: mcp\n" +
			"  description: description\n" +
			"updateMask: updateMask\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tool-sets", "update",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--id", "toolset_01HXKD2E5NQM3T9AYWCFNRMN74",
		)
	})
}

func TestToolSetsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets", "list",
			"--max-items", "10",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--cursor", "cursor",
			"--include-info=true",
			"--labels", "labels",
			"--limit", "0",
			"--prefix", "prefix",
			"--query", "query",
			"--sort-order", "sortOrder",
			"--state", "STATE_UNSPECIFIED",
		)
	})
}

func TestToolSetsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets", "delete",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--id", "toolset_01HXKD2E5NQM3T9AYWCFNRMN74",
		)
	})
}

func TestToolSetsArchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets", "archive",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--id", "toolset_01HXKD2E5NQM3T9AYWCFNRMN74",
		)
	})
}

func TestToolSetsGetOpenAPISpec(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets", "get-openapi-spec",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--tool-set-id", "toolset_01HXKD2E5NQM3T9AYWCFNRMN74",
		)
	})
}

func TestToolSetsListEvents(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets", "list-events",
			"--max-items", "10",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--tool-set-id", "toolset_01HXKD2E5NQM3T9AYWCFNRMN74",
			"--cursor", "cursor",
			"--include-info=true",
			"--labels", "labels",
			"--limit", "0",
			"--sort-order", "sortOrder",
		)
	})
}

func TestToolSetsListUsage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets", "list-usage",
			"--max-items", "10",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--tool-set-id", "toolset_01HXKD2E5NQM3T9AYWCFNRMN74",
			"--cursor", "cursor",
			"--limit", "0",
			"--sort-order", "sortOrder",
			"--tool-id", "tool_01HXKD2E5NQM3T9AYWCFWVYY9K",
		)
	})
}

func TestToolSetsUnarchive(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"tool-sets", "unarchive",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--id", "toolset_01HXKD2E5NQM3T9AYWCFNRMN74",
		)
	})
}
