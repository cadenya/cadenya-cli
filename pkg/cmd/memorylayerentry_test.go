// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
)

func TestMemoryLayersEntriesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"memory-layers:entries", "create",
			"--workspace-id", "workspaceId",
			"--memory-layer-id", "memoryLayerId",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{key: key, content: content, description: description, type: type, uploadId: upload_01HXKD2E5NQM3T9AYWCFZ05DNK}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(memoryLayersEntriesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"memory-layers:entries", "create",
			"--workspace-id", "workspaceId",
			"--memory-layer-id", "memoryLayerId",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.key", "key",
			"--spec.content", "content",
			"--spec.description", "description",
			"--spec.type", "type",
			"--spec.upload-id", "upload_01HXKD2E5NQM3T9AYWCFZ05DNK",
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
			"  key: key\n" +
			"  content: content\n" +
			"  description: description\n" +
			"  type: type\n" +
			"  uploadId: upload_01HXKD2E5NQM3T9AYWCFZ05DNK\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"memory-layers:entries", "create",
			"--workspace-id", "workspaceId",
			"--memory-layer-id", "memoryLayerId",
		)
	})
}

func TestMemoryLayersEntriesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"memory-layers:entries", "retrieve",
			"--workspace-id", "workspaceId",
			"--memory-layer-id", "memoryLayerId",
			"--id", "id",
		)
	})
}

func TestMemoryLayersEntriesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"memory-layers:entries", "update",
			"--workspace-id", "workspaceId",
			"--memory-layer-id", "memoryLayerId",
			"--id", "id",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{content: content, description: description, key: key, uploadId: upload_01HXKD2E5NQM3T9AYWCFZ05DNK}",
			"--update-mask", "updateMask",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(memoryLayersEntriesUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"memory-layers:entries", "update",
			"--workspace-id", "workspaceId",
			"--memory-layer-id", "memoryLayerId",
			"--id", "id",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.content", "content",
			"--spec.description", "description",
			"--spec.key", "key",
			"--spec.upload-id", "upload_01HXKD2E5NQM3T9AYWCFZ05DNK",
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
			"  content: content\n" +
			"  description: description\n" +
			"  key: key\n" +
			"  uploadId: upload_01HXKD2E5NQM3T9AYWCFZ05DNK\n" +
			"updateMask: updateMask\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"memory-layers:entries", "update",
			"--workspace-id", "workspaceId",
			"--memory-layer-id", "memoryLayerId",
			"--id", "id",
		)
	})
}

func TestMemoryLayersEntriesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"memory-layers:entries", "list",
			"--max-items", "10",
			"--workspace-id", "workspaceId",
			"--memory-layer-id", "memoryLayerId",
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

func TestMemoryLayersEntriesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"memory-layers:entries", "delete",
			"--workspace-id", "workspaceId",
			"--memory-layer-id", "memoryLayerId",
			"--id", "id",
		)
	})
}
