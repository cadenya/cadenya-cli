// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/cadenya/cadenya-cli/internal/mocktest"
	"github.com/cadenya/cadenya-cli/internal/requestflag"
)

func TestObjectivesFeedbackCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives:feedback", "create",
			"--objective-id", "objectiveId",
			"--data", "{attributes: {foo: string}, comment: comment, score: 0}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(objectivesFeedbackCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives:feedback", "create",
			"--objective-id", "objectiveId",
			"--data.attributes", "{foo: string}",
			"--data.comment", "comment",
			"--data.score", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"data:\n" +
			"  attributes:\n" +
			"    foo: string\n" +
			"  comment: comment\n" +
			"  score: 0\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"objectives:feedback", "create",
			"--objective-id", "objectiveId",
		)
	})
}

func TestObjectivesFeedbackList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"objectives:feedback", "list",
			"--max-items", "10",
			"--objective-id", "objectiveId",
			"--cursor", "cursor",
			"--limit", "0",
		)
	})
}
