package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
)

func TestProfilesWhoami(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"profiles", "whoami",
		)
	})
}
