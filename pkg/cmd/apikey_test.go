// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
)

func TestAPIKeysCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"api-keys", "create",
			"--workspace-id", "workspaceId",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{description: description, permissions: [string]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(apiKeysCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"api-keys", "create",
			"--workspace-id", "workspaceId",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.description", "description",
			"--spec.permissions", "[string]",
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
			"  description: description\n" +
			"  permissions:\n" +
			"    - string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"api-keys", "create",
			"--workspace-id", "workspaceId",
		)
	})
}

func TestAPIKeysRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"api-keys", "retrieve",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}

func TestAPIKeysUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"api-keys", "update",
			"--workspace-id", "workspaceId",
			"--id", "id",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{description: description, permissions: [string]}",
			"--update-mask", "updateMask",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(apiKeysUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"api-keys", "update",
			"--workspace-id", "workspaceId",
			"--id", "id",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.description", "description",
			"--spec.permissions", "[string]",
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
			"  description: description\n" +
			"  permissions:\n" +
			"    - string\n" +
			"updateMask: updateMask\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"api-keys", "update",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}

func TestAPIKeysList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"api-keys", "list",
			"--max-items", "10",
			"--workspace-id", "workspaceId",
			"--cursor", "cursor",
			"--include-info=true",
			"--labels", "labels",
			"--limit", "0",
			"--prefix", "prefix",
			"--query", "query",
			"--sort-order", "sortOrder",
		)
	})
}

func TestAPIKeysDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"api-keys", "delete",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}

func TestAPIKeysDisable(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"api-keys", "disable",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}

func TestAPIKeysEnable(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"api-keys", "enable",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}

func TestAPIKeysRotate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"api-keys", "rotate",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}
