package config

// Config is the root of a `cadenya config apply` YAML document.
//
// Presence semantics matter: fields that are absent from the YAML are NOT
// written on update. To preserve presence for map-typed spec blobs we keep
// them as map[string]any and derive the update_mask from the set of top-level
// keys the user wrote. For list-typed reconcilable fields (assignments,
// memory_layers, entries) we use *[]T so nil (absent) and empty ([]) are
// distinguishable — empty means "reconcile to no items," nil means "leave
// alone."
type Config struct {
	Version      int                         `yaml:"version"`
	ToolSets     map[string]*ToolSetNode     `yaml:"tool_sets,omitempty"`
	MemoryLayers map[string]*MemoryLayerNode `yaml:"memory_layers,omitempty"`
	Agents       map[string]*AgentNode       `yaml:"agents,omitempty"`
}

// ToolSetNode maps to a single entry under `tool_sets.<external_id>`.
type ToolSetNode struct {
	// ExternalID is filled from the map key during parse.
	ExternalID string `yaml:"-"`

	Name   string               `yaml:"name,omitempty"`
	Labels map[string]string    `yaml:"labels,omitempty"`
	Spec   map[string]any       `yaml:"spec,omitempty"`
	Tools  map[string]*ToolNode `yaml:"tools,omitempty"`
}

// ToolNode maps to a single entry under `tool_sets.<ts>.tools.<external_id>`.
type ToolNode struct {
	ExternalID string `yaml:"-"`

	Name   string            `yaml:"name,omitempty"`
	Labels map[string]string `yaml:"labels,omitempty"`
	Spec   map[string]any    `yaml:"spec,omitempty"`
}

// MemoryLayerNode maps to `memory_layers.<external_id>`.
//
// Entries is a tri-state pointer: nil means the user did not write `entries:`
// or `entries_from_files:`, so the layer's existing entry set is left alone.
// A non-nil (possibly empty) map means "reconcile current entries to exactly
// this set."
type MemoryLayerNode struct {
	ExternalID string `yaml:"-"`

	Name             string                   `yaml:"name,omitempty"`
	Labels           map[string]string        `yaml:"labels,omitempty"`
	Spec             map[string]any           `yaml:"spec,omitempty"`
	Entries          map[string]*EntryNode    `yaml:"entries,omitempty"`
	EntriesFromFiles []*EntriesFromFilesBlock `yaml:"entries_from_files,omitempty"`
}

// ReconcileEntries reports whether the user declared `entries:` or
// `entries_from_files:` on this layer. If false, the apply leaves the layer's
// current entry set untouched.
//
// Call after ExpandFromFiles — which ensures Entries is non-nil whenever a
// file-glob expansion ran, even if the glob produced no inline entries.
func (l *MemoryLayerNode) ReconcileEntries() bool {
	return l.Entries != nil
}

// EntryNode is a single memory entry, keyed by its map key (== MemoryEntry.key).
type EntryNode struct {
	// Key is filled from the map key during parse. For entries expanded from
	// `entries_from_files`, Key is derived from the key_from template.
	Key string `yaml:"-"`

	Description string `yaml:"description,omitempty"`
	Content     string `yaml:"content,omitempty"`

	// SourceFile, when non-empty, records the file this entry was loaded from
	// (used for error messages and for the >1MiB upload path).
	SourceFile string `yaml:"-"`
}

// EntriesFromFilesBlock is the CLI-only construct that expands a file glob
// into additional MemoryEntry ops before reconciliation.
type EntriesFromFilesBlock struct {
	Glob        string `yaml:"glob"`
	KeyFrom     string `yaml:"key_from"`
	StripPrefix string `yaml:"strip_prefix,omitempty"`
}

// AgentNode maps to `agents.<external_id>`.
type AgentNode struct {
	ExternalID string `yaml:"-"`

	Name       string                    `yaml:"name,omitempty"`
	Labels     map[string]string         `yaml:"labels,omitempty"`
	Spec       map[string]any            `yaml:"spec,omitempty"`
	Variations map[string]*VariationNode `yaml:"variations,omitempty"`
}

// VariationNode maps to `agents.<a>.variations.<external_id>`.
//
// Assignments and MemoryLayers use *[]T for the same tri-state reason as
// MemoryLayerNode.Entries.
type VariationNode struct {
	ExternalID string `yaml:"-"`

	Name         string            `yaml:"name,omitempty"`
	Labels       map[string]string `yaml:"labels,omitempty"`
	Spec         map[string]any    `yaml:"spec,omitempty"`
	Assignments  *[]*AssignmentRef `yaml:"assignments,omitempty"`
	MemoryLayers *[]string         `yaml:"memory_layers,omitempty"`
}

// AssignmentRef is a single-key map of kind → external_id. Exactly one of
// Tool, ToolSet, Agent must be non-empty after parsing; the parser enforces
// this and fails the plan otherwise.
type AssignmentRef struct {
	Tool    string `yaml:"tool,omitempty"`
	ToolSet string `yaml:"tool_set,omitempty"`
	Agent   string `yaml:"agent,omitempty"`
}
