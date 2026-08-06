// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
)

func TestAIProviderKeysCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai-provider-keys", "create",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{config: {openrouter: {region: region}, type: openrouter}, credentials: {apiKey: {apiKey: apiKey}, type: apiKey}, provider: AI_PROVIDER_UNSPECIFIED}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiProviderKeysCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai-provider-keys", "create",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.config", "{openrouter: {region: region}, type: openrouter}",
			"--spec.credentials", "{apiKey: {apiKey: apiKey}, type: apiKey}",
			"--spec.provider", "AI_PROVIDER_UNSPECIFIED",
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
			"    openrouter:\n" +
			"      region: region\n" +
			"    type: openrouter\n" +
			"  credentials:\n" +
			"    apiKey:\n" +
			"      apiKey: apiKey\n" +
			"    type: apiKey\n" +
			"  provider: AI_PROVIDER_UNSPECIFIED\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai-provider-keys", "create",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
		)
	})
}

func TestAIProviderKeysRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai-provider-keys", "retrieve",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--id", "aipk_01HXKD2E5NQM3T9AYWCFQ41VW3",
		)
	})
}

func TestAIProviderKeysUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai-provider-keys", "update",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--id", "aipk_01HXKD2E5NQM3T9AYWCFQ41VW3",
			"--metadata", "{name: name, externalId: externalId, labels: {foo: string}}",
			"--spec", "{config: {openrouter: {region: region}, type: openrouter}, credentials: {apiKey: {apiKey: apiKey}, type: apiKey}, provider: AI_PROVIDER_UNSPECIFIED}",
			"--update-mask", "updateMask",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(aiProviderKeysUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai-provider-keys", "update",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--id", "aipk_01HXKD2E5NQM3T9AYWCFQ41VW3",
			"--metadata.name", "name",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--spec.config", "{openrouter: {region: region}, type: openrouter}",
			"--spec.credentials", "{apiKey: {apiKey: apiKey}, type: apiKey}",
			"--spec.provider", "AI_PROVIDER_UNSPECIFIED",
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
			"    openrouter:\n" +
			"      region: region\n" +
			"    type: openrouter\n" +
			"  credentials:\n" +
			"    apiKey:\n" +
			"      apiKey: apiKey\n" +
			"    type: apiKey\n" +
			"  provider: AI_PROVIDER_UNSPECIFIED\n" +
			"updateMask: updateMask\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ai-provider-keys", "update",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--id", "aipk_01HXKD2E5NQM3T9AYWCFQ41VW3",
		)
	})
}

func TestAIProviderKeysList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai-provider-keys", "list",
			"--max-items", "10",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--cursor", "cursor",
			"--include-info=true",
			"--labels", "labels",
			"--limit", "0",
			"--prefix", "prefix",
			"--promotional=true",
			"--query", "query",
			"--sort-order", "sortOrder",
		)
	})
}

func TestAIProviderKeysDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ai-provider-keys", "delete",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--id", "aipk_01HXKD2E5NQM3T9AYWCFQ41VW3",
		)
	})
}
