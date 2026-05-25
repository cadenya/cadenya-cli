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
			"--workspace-id", "workspaceId",
			"--metadata", "{name: name, bundleKey: bundleKey, externalId: externalId, labels: {foo: string}}",
			"--spec", "{adapter: {http: {baseUrl: baseUrl, headers: {foo: string}}, mcp: {excludeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, headers: {foo: string}, includeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, toolApprovals: {always: true, only: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}}, url: url}, openapi: {baseUrl: baseUrl, excludeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, headers: {foo: string}, includeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, serverName: serverName, toolApprovals: {always: true, only: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}}, uploadId: uploadId, url: url}}, description: description}",
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
			"--workspace-id", "workspaceId",
			"--metadata.name", "name",
			"--metadata.bundle-key", "bundleKey",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.adapter", "{http: {baseUrl: baseUrl, headers: {foo: string}}, mcp: {excludeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, headers: {foo: string}, includeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, toolApprovals: {always: true, only: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}}, url: url}, openapi: {baseUrl: baseUrl, excludeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, headers: {foo: string}, includeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, serverName: serverName, toolApprovals: {always: true, only: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}}, uploadId: uploadId, url: url}}",
			"--spec.description", "description",
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
			"  adapter:\n" +
			"    http:\n" +
			"      baseUrl: baseUrl\n" +
			"      headers:\n" +
			"        foo: string\n" +
			"    mcp:\n" +
			"      excludeTools:\n" +
			"        operator: OPERATOR_UNSPECIFIED\n" +
			"        filters:\n" +
			"          - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"            matcher:\n" +
			"              caseSensitive: true\n" +
			"              contains: contains\n" +
			"              endsWith: endsWith\n" +
			"              exact: exact\n" +
			"              regex: regex\n" +
			"              startsWith: startsWith\n" +
			"      headers:\n" +
			"        foo: string\n" +
			"      includeTools:\n" +
			"        operator: OPERATOR_UNSPECIFIED\n" +
			"        filters:\n" +
			"          - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"            matcher:\n" +
			"              caseSensitive: true\n" +
			"              contains: contains\n" +
			"              endsWith: endsWith\n" +
			"              exact: exact\n" +
			"              regex: regex\n" +
			"              startsWith: startsWith\n" +
			"      toolApprovals:\n" +
			"        always: true\n" +
			"        only:\n" +
			"          operator: OPERATOR_UNSPECIFIED\n" +
			"          filters:\n" +
			"            - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"              matcher:\n" +
			"                caseSensitive: true\n" +
			"                contains: contains\n" +
			"                endsWith: endsWith\n" +
			"                exact: exact\n" +
			"                regex: regex\n" +
			"                startsWith: startsWith\n" +
			"      url: url\n" +
			"    openapi:\n" +
			"      baseUrl: baseUrl\n" +
			"      excludeTools:\n" +
			"        operator: OPERATOR_UNSPECIFIED\n" +
			"        filters:\n" +
			"          - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"            matcher:\n" +
			"              caseSensitive: true\n" +
			"              contains: contains\n" +
			"              endsWith: endsWith\n" +
			"              exact: exact\n" +
			"              regex: regex\n" +
			"              startsWith: startsWith\n" +
			"      headers:\n" +
			"        foo: string\n" +
			"      includeTools:\n" +
			"        operator: OPERATOR_UNSPECIFIED\n" +
			"        filters:\n" +
			"          - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"            matcher:\n" +
			"              caseSensitive: true\n" +
			"              contains: contains\n" +
			"              endsWith: endsWith\n" +
			"              exact: exact\n" +
			"              regex: regex\n" +
			"              startsWith: startsWith\n" +
			"      serverName: serverName\n" +
			"      toolApprovals:\n" +
			"        always: true\n" +
			"        only:\n" +
			"          operator: OPERATOR_UNSPECIFIED\n" +
			"          filters:\n" +
			"            - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"              matcher:\n" +
			"                caseSensitive: true\n" +
			"                contains: contains\n" +
			"                endsWith: endsWith\n" +
			"                exact: exact\n" +
			"                regex: regex\n" +
			"                startsWith: startsWith\n" +
			"      uploadId: uploadId\n" +
			"      url: url\n" +
			"  description: description\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tool-sets", "create",
			"--workspace-id", "workspaceId",
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
			"--workspace-id", "workspaceId",
			"--id", "id",
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
			"--workspace-id", "workspaceId",
			"--id", "id",
			"--metadata", "{name: name, bundleKey: bundleKey, externalId: externalId, labels: {foo: string}}",
			"--spec", "{adapter: {http: {baseUrl: baseUrl, headers: {foo: string}}, mcp: {excludeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, headers: {foo: string}, includeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, toolApprovals: {always: true, only: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}}, url: url}, openapi: {baseUrl: baseUrl, excludeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, headers: {foo: string}, includeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, serverName: serverName, toolApprovals: {always: true, only: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}}, uploadId: uploadId, url: url}}, description: description}",
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
			"--workspace-id", "workspaceId",
			"--id", "id",
			"--metadata.name", "name",
			"--metadata.bundle-key", "bundleKey",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.adapter", "{http: {baseUrl: baseUrl, headers: {foo: string}}, mcp: {excludeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, headers: {foo: string}, includeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, toolApprovals: {always: true, only: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}}, url: url}, openapi: {baseUrl: baseUrl, excludeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, headers: {foo: string}, includeTools: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}, serverName: serverName, toolApprovals: {always: true, only: {operator: OPERATOR_UNSPECIFIED, filters: [{attribute: ATTRIBUTE_UNSPECIFIED, matcher: {caseSensitive: true, contains: contains, endsWith: endsWith, exact: exact, regex: regex, startsWith: startsWith}}]}}, uploadId: uploadId, url: url}}",
			"--spec.description", "description",
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
			"  adapter:\n" +
			"    http:\n" +
			"      baseUrl: baseUrl\n" +
			"      headers:\n" +
			"        foo: string\n" +
			"    mcp:\n" +
			"      excludeTools:\n" +
			"        operator: OPERATOR_UNSPECIFIED\n" +
			"        filters:\n" +
			"          - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"            matcher:\n" +
			"              caseSensitive: true\n" +
			"              contains: contains\n" +
			"              endsWith: endsWith\n" +
			"              exact: exact\n" +
			"              regex: regex\n" +
			"              startsWith: startsWith\n" +
			"      headers:\n" +
			"        foo: string\n" +
			"      includeTools:\n" +
			"        operator: OPERATOR_UNSPECIFIED\n" +
			"        filters:\n" +
			"          - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"            matcher:\n" +
			"              caseSensitive: true\n" +
			"              contains: contains\n" +
			"              endsWith: endsWith\n" +
			"              exact: exact\n" +
			"              regex: regex\n" +
			"              startsWith: startsWith\n" +
			"      toolApprovals:\n" +
			"        always: true\n" +
			"        only:\n" +
			"          operator: OPERATOR_UNSPECIFIED\n" +
			"          filters:\n" +
			"            - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"              matcher:\n" +
			"                caseSensitive: true\n" +
			"                contains: contains\n" +
			"                endsWith: endsWith\n" +
			"                exact: exact\n" +
			"                regex: regex\n" +
			"                startsWith: startsWith\n" +
			"      url: url\n" +
			"    openapi:\n" +
			"      baseUrl: baseUrl\n" +
			"      excludeTools:\n" +
			"        operator: OPERATOR_UNSPECIFIED\n" +
			"        filters:\n" +
			"          - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"            matcher:\n" +
			"              caseSensitive: true\n" +
			"              contains: contains\n" +
			"              endsWith: endsWith\n" +
			"              exact: exact\n" +
			"              regex: regex\n" +
			"              startsWith: startsWith\n" +
			"      headers:\n" +
			"        foo: string\n" +
			"      includeTools:\n" +
			"        operator: OPERATOR_UNSPECIFIED\n" +
			"        filters:\n" +
			"          - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"            matcher:\n" +
			"              caseSensitive: true\n" +
			"              contains: contains\n" +
			"              endsWith: endsWith\n" +
			"              exact: exact\n" +
			"              regex: regex\n" +
			"              startsWith: startsWith\n" +
			"      serverName: serverName\n" +
			"      toolApprovals:\n" +
			"        always: true\n" +
			"        only:\n" +
			"          operator: OPERATOR_UNSPECIFIED\n" +
			"          filters:\n" +
			"            - attribute: ATTRIBUTE_UNSPECIFIED\n" +
			"              matcher:\n" +
			"                caseSensitive: true\n" +
			"                contains: contains\n" +
			"                endsWith: endsWith\n" +
			"                exact: exact\n" +
			"                regex: regex\n" +
			"                startsWith: startsWith\n" +
			"      uploadId: uploadId\n" +
			"      url: url\n" +
			"  description: description\n" +
			"updateMask: updateMask\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"tool-sets", "update",
			"--workspace-id", "workspaceId",
			"--id", "id",
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
			"--workspace-id", "workspaceId",
			"--bundle-key", "bundleKey",
			"--cursor", "cursor",
			"--include-info=true",
			"--limit", "0",
			"--prefix", "prefix",
			"--query", "query",
			"--sort-order", "sortOrder",
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
			"--workspace-id", "workspaceId",
			"--id", "id",
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
			"--workspace-id", "workspaceId",
			"--tool-set-id", "toolSetId",
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
			"--workspace-id", "workspaceId",
			"--tool-set-id", "toolSetId",
			"--cursor", "cursor",
			"--include-info=true",
			"--limit", "0",
			"--sort-order", "sortOrder",
		)
	})
}
