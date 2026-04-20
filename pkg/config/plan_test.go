package config

import (
	"strings"
	"testing"
)

// TestResolveRefMissingFailsWithDiagnostic covers the "unresolvable reference"
// error path of resolveRef — invoked without a network because the lookup
// callback is injectable.
func TestResolveRefMissingFailsWithDiagnostic(t *testing.T) {
	b := &planBuilder{cfg: &Config{}, plan: &Plan{}, resolved: map[string]lookupResult{}}
	notInConfig := func(string) bool { return false }
	notFound := func(_ string) (lookupResult, error) { return lookupResult{Exists: false}, nil }
	_, err := b.resolveRef("nowhere_ts", notInConfig, notFound, "tool_sets")
	if err == nil {
		t.Fatal("expected error for unresolvable ref")
	}
	for _, want := range []string{"tool_sets", "nowhere_ts", "not in this config", "workspace"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("error should mention %q, got: %v", want, err)
		}
	}
}

// TestResolveRefUsesConfigPresence checks that a ref pointing at a resource
// declared in the same YAML resolves to "" (will-create) when that resource
// doesn't exist in the workspace yet, rather than erroring.
func TestResolveRefUsesConfigPresence(t *testing.T) {
	b := &planBuilder{
		cfg: &Config{
			ToolSets: map[string]*ToolSetNode{"in_config": {ExternalID: "in_config"}},
		},
		plan:     &Plan{},
		resolved: map[string]lookupResult{},
	}
	inConfig := b.configHasToolSet
	notFound := func(_ string) (lookupResult, error) { return lookupResult{Exists: false}, nil }
	got, err := b.resolveRef("in_config", inConfig, notFound, "tool_sets")
	if err != nil {
		t.Fatalf("expected nil error for in-config will-create ref, got: %v", err)
	}
	if got != "" {
		t.Errorf("expected empty canonical id (will-create sentinel), got %q", got)
	}
}

// TestBuildReportsAllRefErrorsAtOnce exercises the bulk-error aggregation:
// a config with multiple broken refs should surface all of them in a single
// error so the user can fix them in one pass.
func TestBuildReportsAllRefErrorsAtOnce(t *testing.T) {
	asns := []*AssignmentRef{
		{ToolSet: "missing_a"},
		{ToolSet: "missing_b"},
	}
	mls := []string{"missing_layer_c"}
	cfg := &Config{
		Agents: map[string]*AgentNode{
			"a": {
				ExternalID: "a",
				Variations: map[string]*VariationNode{
					"default": {
						ExternalID:   "default",
						Assignments:  &asns,
						MemoryLayers: &mls,
					},
				},
			},
		},
	}

	// We don't have a real client. Build the plan state manually to exercise
	// resolveForAssignment / resolveMemoryLayer via a stub lookup chain.
	// Easier path: hit buildAssignmentOps / buildMemoryLayerLinkOps directly
	// and verify refErrors accumulate.
	b := &planBuilder{
		cfg:      cfg,
		plan:     &Plan{},
		resolved: map[string]lookupResult{},
	}
	// Pre-populate resolved map as if lookups ran and returned Exists=false.
	// resolveRef's workspace-side lookup callback will also return Exists=false
	// for anything not pre-cached, so simulate that by returning a cached miss.
	b.resolved["tool_set:missing_a"] = lookupResult{Exists: false}
	b.resolved["tool_set:missing_b"] = lookupResult{Exists: false}
	b.resolved["memory_layer:missing_layer_c"] = lookupResult{Exists: false}

	v := cfg.Agents["a"].Variations["default"]
	_ = b.buildAssignmentOps("a/default", v, lookupResult{Exists: true})
	_ = b.buildMemoryLayerLinkOps("a/default", v, lookupResult{Exists: true})

	if got := len(b.refErrors); got != 3 {
		t.Fatalf("expected 3 ref errors, got %d: %v", got, b.refErrors)
	}
	combined := ""
	for _, e := range b.refErrors {
		combined += e.Error() + "\n"
	}
	for _, want := range []string{"missing_a", "missing_b", "missing_layer_c"} {
		if !strings.Contains(combined, want) {
			t.Errorf("combined errors should mention %q, got:\n%s", want, combined)
		}
	}
}
