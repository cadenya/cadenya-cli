package config

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/tidwall/gjson"
)

func TestParseToolSets(t *testing.T) {
	src := `version: 1

tool_sets:
  crm_api:
    name: "CRM API"
    spec:
      description: "Internal CRM read/write"
      adapter:
        http:
          base_url: "https://crm.internal/api"
          headers:
            x-api-version: "2024-01"
    tools:
      lookup_customer:
        name: "Lookup Customer"
        spec:
          description: "Look up a customer by email"
          requires_approval: false
`
	cfg, _, err := Parse(strings.NewReader(src), map[string]string{})
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}

	ts, ok := cfg.ToolSets["crm_api"]
	if !ok {
		t.Fatalf("expected tool_sets.crm_api")
	}
	if ts.ExternalID != "crm_api" {
		t.Errorf("ExternalID = %q, want crm_api", ts.ExternalID)
	}
	if ts.Name != "CRM API" {
		t.Errorf("Name = %q, want CRM API", ts.Name)
	}
	if _, ok := ts.Spec["description"]; !ok {
		t.Errorf("spec.description missing from parsed spec: %v", ts.Spec)
	}

	tool, ok := ts.Tools["lookup_customer"]
	if !ok {
		t.Fatalf("expected tool lookup_customer")
	}
	if tool.ExternalID != "lookup_customer" {
		t.Errorf("tool.ExternalID = %q, want lookup_customer", tool.ExternalID)
	}
}

func TestParseEnvSubstitution(t *testing.T) {
	src := `version: 1
tool_sets:
  mcp:
    name: "Linear"
    spec:
      adapter:
        mcp:
          url: "https://mcp.linear.app/sse"
          headers:
            authorization: "Bearer $LINEAR_TOKEN"
`
	cfg, _, err := Parse(strings.NewReader(src), map[string]string{
		"LINEAR_TOKEN": "secret123",
	})
	if err != nil {
		t.Fatalf("Parse: %v", err)
	}

	adapter := cfg.ToolSets["mcp"].Spec["adapter"].(map[string]any)
	mcp := adapter["mcp"].(map[string]any)
	headers := mcp["headers"].(map[string]any)
	got := headers["authorization"].(string)
	want := "Bearer secret123"
	if got != want {
		t.Errorf("authorization = %q, want %q", got, want)
	}
}

func TestParseUnsetEnvVarFails(t *testing.T) {
	src := `version: 1
tool_sets:
  x:
    name: "X"
    spec:
      description: "$UNSET_VARIABLE"
`
	_, _, err := Parse(strings.NewReader(src), map[string]string{})
	if err == nil {
		t.Fatal("expected error for unset env var")
	}
	if !strings.Contains(err.Error(), "UNSET_VARIABLE") {
		t.Errorf("error should name the variable, got: %v", err)
	}
}

func TestParseVersionValidation(t *testing.T) {
	src := `version: 2
tool_sets: {}
`
	_, _, err := Parse(strings.NewReader(src), map[string]string{})
	if err == nil || !strings.Contains(err.Error(), "version") {
		t.Errorf("expected version error, got: %v", err)
	}
}

func TestParseAssignmentValidation(t *testing.T) {
	src := `version: 1
agents:
  a:
    variations:
      v:
        assignments:
          - {}
`
	_, _, err := Parse(strings.NewReader(src), map[string]string{})
	if err == nil || !strings.Contains(err.Error(), "tool/tool_set/agent") {
		t.Errorf("expected empty-assignment error, got: %v", err)
	}

	src2 := `version: 1
agents:
  a:
    variations:
      v:
        assignments:
          - tool: foo
            tool_set: bar
`
	_, _, err = Parse(strings.NewReader(src2), map[string]string{})
	if err == nil || !strings.Contains(err.Error(), "exactly one") {
		t.Errorf("expected multi-target error, got: %v", err)
	}
}

func TestFieldEqualMatchesProto3DefaultElision(t *testing.T) {
	cases := []struct {
		name    string
		desired any
		current string
		path    string
		want    bool
	}{
		{"bool false matches missing", false, `{"spec":{}}`, "spec.requiresApproval", true},
		{"bool true missing does not match", true, `{"spec":{}}`, "spec.requiresApproval", false},
		{"empty string matches missing", "", `{"spec":{}}`, "spec.name", true},
		{"non-empty string missing does not match", "foo", `{"spec":{}}`, "spec.name", false},
		{"empty map matches missing", map[string]any{}, `{"metadata":{}}`, "metadata.labels", true},
		{"non-empty map missing does not match", map[string]any{"a": "b"}, `{"metadata":{}}`, "metadata.labels", false},
		{"nested default stripped",
			map[string]any{"enabled": false, "name": "foo"},
			`{"spec":{"cfg":{"name":"foo"}}}`, "spec.cfg", true},
		{"nested non-default differs",
			map[string]any{"enabled": true, "name": "foo"},
			`{"spec":{"cfg":{"name":"foo"}}}`, "spec.cfg", false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := fieldEqual(c.path, c.desired, gjson.Parse(c.current))
			if got != c.want {
				dj, _ := json.Marshal(c.desired)
				t.Errorf("fieldEqual(%s, %s, %s) = %v, want %v", c.path, dj, c.current, got, c.want)
			}
		})
	}
}

// -----------------------------------------------------------------------------
// Unhappy-path parse tests
// -----------------------------------------------------------------------------

func TestParseMalformedYAML(t *testing.T) {
	src := "version: 1\n  tool_sets: {\n"
	_, _, err := Parse(strings.NewReader(src), map[string]string{})
	if err == nil {
		t.Fatal("expected parse error on malformed YAML")
	}
	if !strings.Contains(err.Error(), "config: parse YAML") {
		t.Errorf("error should be wrapped with 'config: parse YAML', got: %v", err)
	}
}

func TestParseUnknownTopLevelKey(t *testing.T) {
	src := "version: 1\ntool_set: {}\n" // typo: missing `s`
	_, _, err := Parse(strings.NewReader(src), map[string]string{})
	if err == nil {
		t.Fatal("expected parse error on unknown top-level key")
	}
	// goccy/go-yaml with Strict() reports the unknown field by name.
	if !strings.Contains(err.Error(), "tool_set") {
		t.Errorf("error should name the unknown field, got: %v", err)
	}
}

func TestParseWrongScalarType(t *testing.T) {
	src := "version: one\n" // string where int expected
	_, _, err := Parse(strings.NewReader(src), map[string]string{})
	if err == nil {
		t.Fatal("expected parse error on non-int version")
	}
	if !strings.Contains(err.Error(), "int") {
		t.Errorf("error should mention int type mismatch, got: %v", err)
	}
}

func TestParseWrongShape(t *testing.T) {
	// tool_sets is declared as a map, YAML provides a sequence.
	src := "version: 1\ntool_sets:\n  - a\n"
	_, _, err := Parse(strings.NewReader(src), map[string]string{})
	if err == nil {
		t.Fatal("expected parse error on wrong shape")
	}
}

func TestParseTrivialConfig(t *testing.T) {
	// Only a version, no resources. Valid — apply is a no-op.
	cfg, _, err := Parse(strings.NewReader("version: 1\n"), map[string]string{})
	if err != nil {
		t.Fatalf("expected trivial config to parse, got: %v", err)
	}
	if cfg.Version != 1 {
		t.Errorf("version = %d, want 1", cfg.Version)
	}
	if len(cfg.ToolSets) != 0 || len(cfg.MemoryLayers) != 0 || len(cfg.Agents) != 0 {
		t.Errorf("expected no resources, got: %+v", cfg)
	}
}
