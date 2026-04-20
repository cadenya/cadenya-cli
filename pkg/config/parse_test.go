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
	cfg, err := Parse(strings.NewReader(src), map[string]string{})
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
	cfg, err := Parse(strings.NewReader(src), map[string]string{
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
	_, err := Parse(strings.NewReader(src), map[string]string{})
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
	_, err := Parse(strings.NewReader(src), map[string]string{})
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
	_, err := Parse(strings.NewReader(src), map[string]string{})
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
	_, err = Parse(strings.NewReader(src2), map[string]string{})
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
