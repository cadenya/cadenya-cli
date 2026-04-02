// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
)

func TestWorkspaceSecretsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-secrets", "create",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{value: value}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(workspaceSecretsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-secrets", "create",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.value", "value",
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
			"  value: value\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workspace-secrets", "create",
		)
	})
}

func TestWorkspaceSecretsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-secrets", "retrieve",
			"--id", "id",
		)
	})
}

func TestWorkspaceSecretsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-secrets", "update",
			"--id", "id",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{value: value}",
			"--update-mask", "updateMask",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(workspaceSecretsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-secrets", "update",
			"--id", "id",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.value", "value",
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
			"  value: value\n" +
			"updateMask: updateMask\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workspace-secrets", "update",
			"--id", "id",
		)
	})
}

func TestWorkspaceSecretsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-secrets", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--include-info=true",
			"--limit", "0",
			"--prefix", "prefix",
			"--sort-order", "sortOrder",
		)
	})
}

func TestWorkspaceSecretsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-secrets", "delete",
			"--id", "id",
		)
	})
}
