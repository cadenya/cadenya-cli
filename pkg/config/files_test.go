package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestExpandFromFiles(t *testing.T) {
	dir := t.TempDir()
	skillsDir := filepath.Join(dir, "skills", "postmortem")
	if err := os.MkdirAll(skillsDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	// File with frontmatter.
	writeFile(t, filepath.Join(skillsDir, "write.md"),
		"---\ndescription: When to write a postmortem\n---\nThe body of the skill.\n")

	// File without frontmatter.
	writeFile(t, filepath.Join(dir, "skills", "general.md"), "Plain content.\n")

	cfg := &Config{
		MemoryLayers: map[string]*MemoryLayerNode{
			"company_skills": {
				ExternalID: "company_skills",
				EntriesFromFiles: []*EntriesFromFilesBlock{
					{
						Glob:        "skills/**/*.md",
						KeyFrom:     "{path_without_ext}",
						StripPrefix: "skills/",
					},
				},
			},
		},
	}
	if err := ExpandFromFiles(cfg, dir); err != nil {
		t.Fatalf("ExpandFromFiles: %v", err)
	}

	entries := cfg.MemoryLayers["company_skills"].Entries
	if got := len(entries); got != 2 {
		t.Fatalf("expected 2 entries, got %d: %v", got, entries)
	}

	pm, ok := entries["postmortem/write"]
	if !ok {
		t.Fatalf("missing key postmortem/write, keys: %v", mapKeys(entries))
	}
	if pm.Description != "When to write a postmortem" {
		t.Errorf("description = %q", pm.Description)
	}
	if !strings.Contains(pm.Content, "The body of the skill.") {
		t.Errorf("content = %q", pm.Content)
	}

	gen := entries["general"]
	if gen == nil {
		t.Fatalf("missing key general")
	}
	if gen.Description != "" {
		t.Errorf("expected empty description for non-frontmatter file, got %q", gen.Description)
	}
	if !strings.Contains(gen.Content, "Plain content.") {
		t.Errorf("content = %q", gen.Content)
	}
}

func TestExpandFromFilesZeroMatchFails(t *testing.T) {
	cfg := &Config{
		MemoryLayers: map[string]*MemoryLayerNode{
			"x": {
				ExternalID: "x",
				EntriesFromFiles: []*EntriesFromFilesBlock{
					{Glob: "nonexistent/*.md", KeyFrom: "{basename}"},
				},
			},
		},
	}
	err := ExpandFromFiles(cfg, t.TempDir())
	if err == nil || !strings.Contains(err.Error(), "matched no files") {
		t.Errorf("expected zero-match error, got: %v", err)
	}
}

func TestExpandFromFilesCollisionFails(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "a.md"), "one")
	writeFile(t, filepath.Join(dir, "a.txt"), "two")

	cfg := &Config{
		MemoryLayers: map[string]*MemoryLayerNode{
			"x": {
				ExternalID: "x",
				Entries: map[string]*EntryNode{
					"a": {Key: "a", Content: "inline"},
				},
				EntriesFromFiles: []*EntriesFromFilesBlock{
					{Glob: "a.md", KeyFrom: "{basename_without_ext}"},
				},
			},
		},
	}
	err := ExpandFromFiles(cfg, dir)
	if err == nil || !strings.Contains(err.Error(), "collides") {
		t.Errorf("expected collision error, got: %v", err)
	}
}

func TestRenderKeyTokens(t *testing.T) {
	cases := []struct {
		template, stripPrefix, path, want string
	}{
		{"{path_without_ext}", "skills/", "skills/postmortem/write.md", "postmortem/write"},
		{"{basename}", "", "skills/postmortem/write.md", "write.md"},
		{"{basename_without_ext}", "", "skills/postmortem/write.md", "write"},
		{"{dirname}", "skills/", "skills/postmortem/write.md", "postmortem"},
	}
	for _, c := range cases {
		got, err := renderKey(c.template, c.stripPrefix, c.path)
		if err != nil {
			t.Errorf("renderKey(%q, %q, %q): %v", c.template, c.stripPrefix, c.path, err)
			continue
		}
		if got != c.want {
			t.Errorf("renderKey(%q, %q, %q) = %q, want %q", c.template, c.stripPrefix, c.path, got, c.want)
		}
	}
}

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

func mapKeys[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
