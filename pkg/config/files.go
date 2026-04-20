package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/adrg/frontmatter"
)

// maxInlineContentBytes mirrors MemoryEntryCreateSpec.content's 1 MiB cap.
// Files larger than this need to be routed through UploadService (not yet
// implemented — ExpandFromFiles fails the plan with a clear message).
const maxInlineContentBytes = 1 << 20 // 1 MiB

// ExpandFromFiles walks every memory_layer's entries_from_files blocks and
// merges the resulting entries into each layer's Entries map.
func ExpandFromFiles(cfg *Config, baseDir string) error {
	for extID, ml := range cfg.MemoryLayers {
		if len(ml.EntriesFromFiles) == 0 {
			continue
		}

		// Start from inline entries (if any) so collisions can be detected.
		if ml.Entries == nil {
			ml.Entries = map[string]*EntryNode{}
		}

		for i, block := range ml.EntriesFromFiles {
			expanded, err := expandGlob(block, baseDir)
			if err != nil {
				return fmt.Errorf("config: memory_layers.%s.entries_from_files[%d]: %w", extID, i, err)
			}
			for key, entry := range expanded {
				if existing, collide := ml.Entries[key]; collide {
					which := "another entries_from_files block"
					if existing.SourceFile == "" {
						which = "an inline entries: entry"
					}
					return fmt.Errorf("config: memory_layers.%s: entry key %q from %s collides with %s",
						extID, key, entry.SourceFile, which)
				}
				ml.Entries[key] = entry
			}
		}
	}
	return nil
}

// expandGlob evaluates a single block's glob against baseDir and returns the
// resulting (key → *EntryNode) map. An error is returned if the glob matches
// zero files or if any file can't be read.
func expandGlob(block *EntriesFromFilesBlock, baseDir string) (map[string]*EntryNode, error) {
	if block.Glob == "" {
		return nil, fmt.Errorf("glob: missing")
	}
	if block.KeyFrom == "" {
		return nil, fmt.Errorf("key_from: missing")
	}

	full := filepath.Join(baseDir, block.Glob)
	matches, err := doublestarGlob(full)
	if err != nil {
		return nil, fmt.Errorf("glob %q: %w", block.Glob, err)
	}
	if len(matches) == 0 {
		return nil, fmt.Errorf("glob %q matched no files", block.Glob)
	}

	out := map[string]*EntryNode{}
	for _, abs := range matches {
		// Relative path for template substitution.
		rel, err := filepath.Rel(baseDir, abs)
		if err != nil {
			rel = abs
		}
		rel = filepath.ToSlash(rel)

		key, err := renderKey(block.KeyFrom, block.StripPrefix, rel)
		if err != nil {
			return nil, fmt.Errorf("file %q: %w", rel, err)
		}
		if _, dup := out[key]; dup {
			return nil, fmt.Errorf("file %q produced duplicate key %q within this block", rel, key)
		}

		data, err := os.ReadFile(abs)
		if err != nil {
			return nil, fmt.Errorf("read %q: %w", rel, err)
		}

		desc, body, err := splitFrontmatter(data, abs)
		if err != nil {
			return nil, fmt.Errorf("file %q: %w", rel, err)
		}

		if len(body) > maxInlineContentBytes {
			return nil, fmt.Errorf("file %q: content %d bytes exceeds 1 MiB inline cap; upload path not yet implemented",
				rel, len(body))
		}

		out[key] = &EntryNode{
			Key:         key,
			Description: desc,
			Content:     string(body),
			SourceFile:  rel,
		}
	}
	return out, nil
}

// renderKey applies strip_prefix, then template tokens, to a matched relative
// path. The path should use forward-slash separators.
func renderKey(template, stripPrefix, matchedPath string) (string, error) {
	p := matchedPath
	if stripPrefix != "" {
		if !strings.HasPrefix(p, stripPrefix) {
			return "", fmt.Errorf("strip_prefix %q does not match path %q", stripPrefix, matchedPath)
		}
		p = p[len(stripPrefix):]
	}

	pWithoutExt := p
	if ext := filepath.Ext(p); ext != "" {
		pWithoutExt = strings.TrimSuffix(p, ext)
	}
	base := filepath.Base(p)
	baseWithoutExt := base
	if ext := filepath.Ext(base); ext != "" {
		baseWithoutExt = strings.TrimSuffix(base, ext)
	}
	dir := filepath.ToSlash(filepath.Dir(p))
	if dir == "." {
		dir = ""
	}

	replacer := strings.NewReplacer(
		"{path_without_ext}", pWithoutExt,
		"{path}", p,
		"{basename_without_ext}", baseWithoutExt,
		"{basename}", base,
		"{dirname}", dir,
	)
	return replacer.Replace(template), nil
}

// splitFrontmatter returns (description, body) for files that begin with
// YAML/TOML/JSON frontmatter delimited by ---, +++, or { } respectively.
// Only markdown files participate in the convention; other extensions are
// returned as ("", raw).
//
// Uses github.com/adrg/frontmatter, which handles CRLF, BOM, and all three
// delimiter styles.
func splitFrontmatter(raw []byte, sourcePath string) (desc string, body []byte, err error) {
	if !strings.HasSuffix(strings.ToLower(sourcePath), ".md") {
		return "", raw, nil
	}
	var meta struct {
		Description string `yaml:"description" toml:"description" json:"description"`
	}
	rest, err := frontmatter.Parse(bytes.NewReader(raw), &meta)
	if err != nil {
		return "", nil, fmt.Errorf("parse frontmatter: %w", err)
	}
	return meta.Description, rest, nil
}

// doublestarGlob returns matches for a glob pattern that may contain `**`.
// For now, we only support the single-`*` glob surface that Go's filepath.Glob
// supports; patterns with `**` fall back to a file-walk from the first
// non-wildcard prefix, filtered by a suffix check.
//
// This is adequate for the plan doc's canonical example (`skills/**/*.md`) and
// simple single-star patterns. More sophisticated glob needs can be added
// later by pulling in github.com/bmatcuk/doublestar.
func doublestarGlob(pattern string) ([]string, error) {
	if !strings.Contains(pattern, "**") {
		return filepath.Glob(pattern)
	}
	// Split at the first `**` occurrence.
	idx := strings.Index(pattern, "**")
	head := pattern[:idx]           // e.g. "<base>/skills/"
	tail := pattern[idx+len("**"):] // e.g. "/*.md"

	// Walk head, collecting files that match tail via filepath.Match after
	// stripping the head prefix.
	root := head
	// Trim trailing separator so Walk uses a directory.
	root = strings.TrimRight(root, string(filepath.Separator))
	root = strings.TrimRight(root, "/")
	if root == "" {
		root = "."
	}

	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasPrefix(path, root) {
			return nil
		}
		suffixPart := strings.TrimPrefix(path, root)
		suffixPart = strings.TrimLeft(suffixPart, string(filepath.Separator))
		// The tail pattern likely starts with "/" — trim it.
		trimmedTail := strings.TrimLeft(tail, "/")
		trimmedTail = strings.TrimLeft(trimmedTail, string(filepath.Separator))
		// Match the full suffix against a pattern that allows zero-or-more
		// intermediate directories before trimmedTail. Use filepath.Match only
		// on the final segment.
		base := filepath.Base(path)
		ok, err := filepath.Match(trimmedTail, base)
		if err != nil {
			return err
		}
		if ok {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(matches)
	return matches, nil
}
