package config

import (
	"fmt"
	"io"
	"strings"
)

// RenderPlan prints a Terraform-style plan summary to w.
//
// Reorder ops (currently only memory-layer link position changes) are
// counted as updates for the summary, since they mutate existing state,
// but get their own line prefix ("will reorder → position=N") for clarity.
func RenderPlan(w io.Writer, plan *Plan) error {
	var creates, updates, unchanged, deletes int
	fmt.Fprintln(w, "Planning changes…")
	for _, op := range plan.Ops {
		line := fmt.Sprintf("  %-48s ", op.Target.Path())
		switch op.Kind {
		case OpCreate:
			creates++
			line += "will create"
		case OpUpdate:
			updates++
			line += "will update (" + strings.Join(op.Change.FieldPaths, ", ") + ")"
		case OpReorder:
			updates++
			line += fmt.Sprintf("will reorder → position=%d", op.Change.Position)
		case OpNoChange:
			unchanged++
			line += "no change"
		case OpDelete:
			deletes++
			line += "will delete"
		}
		fmt.Fprintln(w, line)
	}
	fmt.Fprintln(w)
	summary := fmt.Sprintf("Plan: %d to create, %d to update, %d unchanged", creates, updates, unchanged)
	if deletes > 0 {
		summary += fmt.Sprintf(", %d to delete", deletes)
	}
	fmt.Fprintln(w, summary+".")
	fmt.Fprintln(w)
	return nil
}

// RenderOpStart prints the per-op leading line during apply.
func RenderOpStart(w io.Writer, op Op) {
	fmt.Fprintf(w, "  %-48s %s…\n", op.Target.Path(), opGerund(op.Kind))
}

// RenderOpEnd prints the per-op trailing line after the op finishes.
// summary is the "created ts_01HXQ…" / "updated" string from the dispatcher.
func RenderOpEnd(w io.Writer, op Op, summary string, err error) {
	if err != nil {
		fmt.Fprintf(w, "  %-48s %s\n", op.Target.Path(), "failed: "+err.Error())
		return
	}
	if summary != "" {
		fmt.Fprintf(w, "  %-48s %s\n", op.Target.Path(), summary)
	}
}

// RenderSummary prints the final summary line.
func RenderSummary(w io.Writer, result *Result) {
	fmt.Fprintln(w)
	fmt.Fprintf(w, "Applied %d changes. %d unchanged. %d failed.\n",
		result.Applied, result.Unchanged, result.Failed)
}

// RenderDiff is not implemented for the tool_sets vertical.
func RenderDiff(w io.Writer, op Op) error {
	_ = w
	_ = op
	return fmt.Errorf("config.RenderDiff: not implemented")
}

func opGerund(k OpKind) string {
	switch k {
	case OpCreate:
		return "creating"
	case OpUpdate:
		return "updating"
	case OpDelete:
		return "deleting"
	case OpReorder:
		return "reordering"
	}
	return ""
}
