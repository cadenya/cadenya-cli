// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
)

func TestWidgetSessionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"widget-sessions", "create",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--spec", "{widgetId: wgt_01HXKD2E5NQM3T9AYWCFMZZZBD, expiresAt: '2019-12-27T18:11:19.117Z', pinnedParameters: {foo: string}, subject: {id: customer-user-42, name: Jane Doe}, tenant: {id: acme-corp, name: Acme Corp}}",
			"--metadata", "{externalId: externalId, labels: {foo: string}}",
			"--secret", "{name: name, value: value}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(widgetSessionsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"widget-sessions", "create",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--spec.widget-id", "wgt_01HXKD2E5NQM3T9AYWCFMZZZBD",
			"--spec.expires-at", "2019-12-27T18:11:19.117Z",
			"--spec.pinned-parameters", "{foo: string}",
			"--spec.subject", "{id: customer-user-42, name: Jane Doe}",
			"--spec.tenant", "{id: acme-corp, name: Acme Corp}",
			"--metadata.external-id", "externalId",
			"--metadata.labels", "{foo: string}",
			"--secret.name", "name",
			"--secret.value", "value",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"spec:\n" +
			"  widgetId: wgt_01HXKD2E5NQM3T9AYWCFMZZZBD\n" +
			"  expiresAt: '2019-12-27T18:11:19.117Z'\n" +
			"  pinnedParameters:\n" +
			"    foo: string\n" +
			"  subject:\n" +
			"    id: customer-user-42\n" +
			"    name: Jane Doe\n" +
			"  tenant:\n" +
			"    id: acme-corp\n" +
			"    name: Acme Corp\n" +
			"metadata:\n" +
			"  externalId: externalId\n" +
			"  labels:\n" +
			"    foo: string\n" +
			"secrets:\n" +
			"  - name: name\n" +
			"    value: value\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"widget-sessions", "create",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
		)
	})
}

func TestWidgetSessionsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"widget-sessions", "retrieve",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--id", "id",
		)
	})
}

func TestWidgetSessionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"widget-sessions", "list",
			"--max-items", "10",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--cursor", "cursor",
			"--include-info=true",
			"--labels", "labels",
			"--limit", "0",
			"--sort-order", "sortOrder",
			"--state", "STATE_UNSPECIFIED",
			"--subject-id", "subjectId",
			"--tenant-id", "tenantId",
			"--widget-id", "widgetId",
		)
	})
}

func TestWidgetSessionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"widget-sessions", "delete",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--id", "id",
		)
	})
}

func TestWidgetSessionsDeleteTenant(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"widget-sessions", "delete-tenant",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--tenant-id", "tenantId",
		)
	})
}

func TestWidgetSessionsRevoke(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"widget-sessions", "revoke",
			"--workspace-id", "workspace_01HXKD2E5NQM3T9AYWCF133E3Q",
			"--id", "id",
		)
	})
}
