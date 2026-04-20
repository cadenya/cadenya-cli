package config

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRenderPlanFormatsCounts(t *testing.T) {
	plan := &Plan{
		Ops: []Op{
			{Kind: OpCreate, Target: ResourceRef{Kind: KindToolSet, ExternalID: "a"}},
			{Kind: OpUpdate, Target: ResourceRef{Kind: KindAgent, ExternalID: "b"},
				Change: ChangeSet{FieldPaths: []string{"metadata.name", "spec.description"}}},
			{Kind: OpNoChange, Target: ResourceRef{Kind: KindMemoryLayer, ExternalID: "c"}},
			{Kind: OpDelete, Target: ResourceRef{Kind: KindMemoryEntry, Parent: "c", ExternalID: "gone"}},
		},
	}
	var buf bytes.Buffer
	if err := RenderPlan(&buf, plan); err != nil {
		t.Fatalf("RenderPlan: %v", err)
	}
	out := buf.String()

	for _, want := range []string{
		"Planning changes…",
		"tool_sets.a",
		"will create",
		"agents.b",
		"will update (metadata.name, spec.description)",
		"memory_layers.c",
		"no change",
		"memory_layers.c.entries.gone",
		"will delete",
		"Plan: 1 to create, 1 to update, 1 unchanged, 1 to delete.",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q:\n%s", want, out)
		}
	}
}

func TestRenderPlanOmitsDeletesWhenZero(t *testing.T) {
	plan := &Plan{Ops: []Op{{Kind: OpCreate, Target: ResourceRef{Kind: KindAgent, ExternalID: "x"}}}}
	var buf bytes.Buffer
	_ = RenderPlan(&buf, plan)
	summary := buf.String()
	if !strings.Contains(summary, "Plan: 1 to create, 0 to update, 0 unchanged.") {
		t.Errorf("summary should omit delete count when zero, got:\n%s", summary)
	}
	if strings.Contains(summary, "to delete") {
		t.Errorf("summary should not mention 'to delete' when count is 0:\n%s", summary)
	}
}

func TestRenderOpStartAndEnd(t *testing.T) {
	op := Op{
		Kind:   OpCreate,
		Target: ResourceRef{Kind: KindToolSet, ExternalID: "my_ts"},
	}
	var start, end bytes.Buffer
	RenderOpStart(&start, op)
	RenderOpEnd(&end, op, "created toolset_01HXQ", nil)

	if !strings.Contains(start.String(), "tool_sets.my_ts") || !strings.Contains(start.String(), "creating…") {
		t.Errorf("RenderOpStart missing expected content:\n%s", start.String())
	}
	if !strings.Contains(end.String(), "created toolset_01HXQ") {
		t.Errorf("RenderOpEnd should include summary text:\n%s", end.String())
	}
}

func TestRenderOpEndOnFailure(t *testing.T) {
	op := Op{Kind: OpCreate, Target: ResourceRef{Kind: KindToolSet, ExternalID: "broken"}}
	var buf bytes.Buffer
	RenderOpEnd(&buf, op, "", fmt.Errorf("boom"))
	if !strings.Contains(buf.String(), "failed: boom") {
		t.Errorf("failure line should prefix with 'failed:', got:\n%s", buf.String())
	}
}

func TestRenderSummaryCounts(t *testing.T) {
	var buf bytes.Buffer
	RenderSummary(&buf, &Result{Applied: 3, Unchanged: 5, Failed: 1})
	if !strings.Contains(buf.String(), "Applied 3 changes. 5 unchanged. 1 failed.") {
		t.Errorf("summary counts wrong, got:\n%s", buf.String())
	}
}

func TestRenderPlanIncludesReorder(t *testing.T) {
	plan := &Plan{
		Ops: []Op{
			{
				Kind: OpReorder,
				Target: ResourceRef{
					Kind: KindVariationMemoryLayer, Parent: "ag/default",
					ExternalID: "l", RowID: "avml_01",
				},
				Change: ChangeSet{Position: 2},
			},
		},
	}
	var buf bytes.Buffer
	_ = RenderPlan(&buf, plan)
	out := buf.String()
	if !strings.Contains(out, "will reorder → position=2") {
		t.Errorf("reorder line should show target position, got:\n%s", out)
	}
	// Counted as an update for summary purposes.
	if !strings.Contains(out, "Plan: 0 to create, 1 to update, 0 unchanged.") {
		t.Errorf("reorder should count as 'to update' in summary, got:\n%s", out)
	}
}
