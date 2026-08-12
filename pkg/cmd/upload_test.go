package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
)

func TestUploadsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"uploads", "create",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{contentType: contentType, filename: filename, sizeBytes: sizeBytes}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(uploadsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"uploads", "create",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.content-type", "contentType",
			"--spec.filename", "filename",
			"--spec.size-bytes", "sizeBytes",
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
			"  contentType: contentType\n" +
			"  filename: filename\n" +
			"  sizeBytes: sizeBytes\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"uploads", "create",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
		)
	})
}

func TestUploadsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"uploads", "retrieve",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--id", "upload_01HXKD2E5NQM3T9AYWCFZ05DNK",
		)
	})
}
