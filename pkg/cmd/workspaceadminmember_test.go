// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
)

func TestWorkspaceAdminMembersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-admin:members", "list",
			"--max-items", "10",
			"--workspace-id", "workspaceId",
			"--cursor", "cursor",
			"--limit", "0",
		)
	})
}

func TestWorkspaceAdminMembersAdd(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-admin:members", "add",
			"--workspace-id", "workspaceId",
			"--email", "email",
			"--profile-id", "profile_01HXKD2E5NQM3T9AYWCFS0AP08",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"email: email\n" +
			"profileId: profile_01HXKD2E5NQM3T9AYWCFS0AP08\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"workspace-admin:members", "add",
			"--workspace-id", "workspaceId",
		)
	})
}

func TestWorkspaceAdminMembersRemove(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"workspace-admin:members", "remove",
			"--workspace-id", "workspaceId",
			"--profile-id", "profileId",
		)
	})
}
