// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
)

func TestMemoryLayersCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"memory-layers", "create",
			"--workspace-id", "workspaceId",
			"--metadata", "{name: name, bundleKey: bundleKey, externalId: externalId, labels: {foo: string}}",
			"--spec", "{type: MEMORY_LAYER_TYPE_UNSPECIFIED, description: description}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(memoryLayersCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"memory-layers", "create",
			"--workspace-id", "workspaceId",
			"--metadata.name", "name",
			"--metadata.bundle-key", "bundleKey",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.type", "MEMORY_LAYER_TYPE_UNSPECIFIED",
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
			"  type: MEMORY_LAYER_TYPE_UNSPECIFIED\n" +
			"  description: description\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"memory-layers", "create",
			"--workspace-id", "workspaceId",
		)
	})
}

func TestMemoryLayersRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"memory-layers", "retrieve",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}

func TestMemoryLayersUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"memory-layers", "update",
			"--workspace-id", "workspaceId",
			"--id", "id",
			"--metadata", "{name: name, bundleKey: bundleKey, externalId: externalId, labels: {foo: string}}",
			"--spec", "{type: MEMORY_LAYER_TYPE_UNSPECIFIED, description: description}",
			"--update-mask", "updateMask",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(memoryLayersUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"memory-layers", "update",
			"--workspace-id", "workspaceId",
			"--id", "id",
			"--metadata.name", "name",
			"--metadata.bundle-key", "bundleKey",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.type", "MEMORY_LAYER_TYPE_UNSPECIFIED",
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
			"  type: MEMORY_LAYER_TYPE_UNSPECIFIED\n" +
			"  description: description\n" +
			"updateMask: updateMask\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"memory-layers", "update",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}

func TestMemoryLayersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"memory-layers", "list",
			"--max-items", "10",
			"--workspace-id", "workspaceId",
			"--bundle-key", "bundleKey",
			"--cursor", "cursor",
			"--include-info=true",
			"--limit", "0",
			"--prefix", "prefix",
			"--query", "query",
			"--sort-order", "sortOrder",
			"--type", "MEMORY_LAYER_TYPE_UNSPECIFIED",
		)
	})
}

func TestMemoryLayersDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"memory-layers", "delete",
			"--workspace-id", "workspaceId",
			"--id", "id",
		)
	})
}
