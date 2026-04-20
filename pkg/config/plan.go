package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sort"

	"github.com/cadenya/cadenya-go"
	"github.com/cadenya/cadenya-go/option"
	"github.com/tidwall/gjson"
)

// Plan is the output of the read-only phase of `config apply`.
type Plan struct {
	Ops []Op
}

// Op is one intended action against one resource.
type Op struct {
	Kind   OpKind
	Target ResourceRef
	Change ChangeSet
}

type OpKind int

const (
	OpNoChange OpKind = iota
	OpCreate
	OpUpdate
	OpDelete
	OpReorder // memory-layer link position change
)

func (k OpKind) String() string {
	switch k {
	case OpNoChange:
		return "no change"
	case OpCreate:
		return "create"
	case OpUpdate:
		return "update"
	case OpDelete:
		return "delete"
	case OpReorder:
		return "reorder"
	}
	return "unknown"
}

// ResourceRef identifies a single resource within the plan.
//
// Some fields carry kind-specific meaning:
//   - For KindVariationAssignment (add): ExternalID is the target in kind:ext form
//     (e.g. "tool_set:crm_api"). Parent is "<agent_ext>/<variation_ext>".
//   - For KindVariationAssignment (delete): RowID is the server-side row id.
//     Parent is the same. ExternalID is kind:canonical_id of target for display.
//   - For KindVariationMemoryLayer: ExternalID is the memory_layer external_id.
//     Parent is "<agent_ext>/<variation_ext>". RowID is set on delete/reorder.
type ResourceRef struct {
	Kind       ResourceKind
	Parent     string
	ExternalID string
	RowID      string
}

func (r ResourceRef) Path() string {
	switch r.Kind {
	case KindToolSet:
		return "tool_sets." + r.ExternalID
	case KindTool:
		return "tool_sets." + r.Parent + ".tools." + r.ExternalID
	case KindMemoryLayer:
		return "memory_layers." + r.ExternalID
	case KindMemoryEntry:
		return "memory_layers." + r.Parent + ".entries." + r.ExternalID
	case KindAgent:
		return "agents." + r.ExternalID
	case KindVariation:
		return "agents." + r.Parent + ".variations." + r.ExternalID
	case KindVariationAssignment:
		return "agents." + r.Parent + ".assignments[" + r.ExternalID + "]"
	case KindVariationMemoryLayer:
		return "agents." + r.Parent + ".memory_layers[" + r.ExternalID + "]"
	}
	return "<unknown>"
}

type ResourceKind int

const (
	KindToolSet ResourceKind = iota
	KindTool
	KindMemoryLayer
	KindMemoryEntry
	KindAgent
	KindVariation
	KindVariationAssignment
	KindVariationMemoryLayer
)

// ChangeSet describes what will be written for an Op.
type ChangeSet struct {
	// For OpCreate / OpUpdate of {metadata, spec}-shaped resources:
	FieldPaths []string
	Body       map[string]any

	// For list reconciliation ops (assignments, memory-layer links, entries):
	TargetKind       string // "tool" | "tool_set" | "agent" | "memory_layer"
	TargetExternalID string // YAML-declared external_id of the target — sent as "external_id:<ext>" in the wire body
	Position         int    // memory-layer link position (add or reorder)
}

// planBuilder carries state across the multi-pass Build.
type planBuilder struct {
	ctx    context.Context
	client *cadenya.Client
	cfg    *Config
	plan   *Plan

	// Resolver cache: kind + ":" + external_id → lookup result (canonical id,
	// raw JSON, or "does not exist").
	resolved map[string]lookupResult

	// Reference errors collected across the whole Build pass. Unresolvable
	// YAML refs don't short-circuit; they accumulate and are joined into a
	// single error at the end so the user sees every broken reference in one
	// round-trip instead of fixing them one at a time.
	refErrors []error
}

type lookupResult struct {
	CanonicalID string
	RawJSON     string
	Exists      bool
}

// Build walks the parsed Config, issues read-only lookups via `client`, and
// returns a Plan.
func Build(ctx context.Context, client *cadenya.Client, cfg *Config) (*Plan, error) {
	b := &planBuilder{
		ctx:      ctx,
		client:   client,
		cfg:      cfg,
		plan:     &Plan{},
		resolved: map[string]lookupResult{},
	}
	if err := b.buildToolSetOps(); err != nil {
		return nil, err
	}
	if err := b.buildMemoryLayerOps(); err != nil {
		return nil, err
	}
	if err := b.buildAgentOps(); err != nil {
		return nil, err
	}
	if len(b.refErrors) > 0 {
		return nil, errors.Join(b.refErrors...)
	}
	return b.plan, nil
}

// -----------------------------------------------------------------------------
// tool sets
// -----------------------------------------------------------------------------

func (b *planBuilder) buildToolSetOps() error {
	for _, extID := range sortedStringKeys(b.cfg.ToolSets) {
		ts := b.cfg.ToolSets[extID]
		ref := ResourceRef{Kind: KindToolSet, ExternalID: extID}

		lr, err := b.lookupToolSet(extID)
		if err != nil {
			return err
		}
		op := planResourceSpec(ref, ts.Name, ts.Labels, ts.Spec, extID, lr)
		b.plan.Ops = append(b.plan.Ops, op)

		for _, tExtID := range sortedStringKeys(ts.Tools) {
			tool := ts.Tools[tExtID]
			tref := ResourceRef{Kind: KindTool, Parent: extID, ExternalID: tExtID}

			var tlr lookupResult
			if lr.Exists {
				tlr, err = b.lookupTool(extID, tExtID)
				if err != nil {
					return err
				}
			}
			b.plan.Ops = append(b.plan.Ops,
				planResourceSpec(tref, tool.Name, tool.Labels, tool.Spec, tExtID, tlr))
		}
	}
	return nil
}

// -----------------------------------------------------------------------------
// memory layers + entries
// -----------------------------------------------------------------------------

func (b *planBuilder) buildMemoryLayerOps() error {
	for _, extID := range sortedStringKeys(b.cfg.MemoryLayers) {
		ml := b.cfg.MemoryLayers[extID]
		ref := ResourceRef{Kind: KindMemoryLayer, ExternalID: extID}

		lr, err := b.lookupMemoryLayer(extID)
		if err != nil {
			return err
		}
		op := planResourceSpec(ref, ml.Name, ml.Labels, ml.Spec, extID, lr)
		b.plan.Ops = append(b.plan.Ops, op)

		// Entries reconciliation. Only run if the user declared entries: or
		// entries_from_files: on this layer.
		if !ml.ReconcileEntries() {
			continue
		}
		if err := b.buildEntryOps(extID, ml, lr); err != nil {
			return err
		}
	}
	return nil
}

func (b *planBuilder) buildEntryOps(layerExtID string, ml *MemoryLayerNode, layerLookup lookupResult) error {
	desired := ml.Entries
	if desired == nil {
		desired = map[string]*EntryNode{}
	}

	currentByKey := map[string]currentEntry{}
	if layerLookup.Exists {
		cs, err := b.listEntries(layerExtID)
		if err != nil {
			return err
		}
		for _, c := range cs {
			currentByKey[c.Key] = c
		}
	}

	for _, key := range sortedStringKeys(desired) {
		entry := desired[key]
		ref := ResourceRef{Kind: KindMemoryEntry, Parent: layerExtID, ExternalID: key}
		cur, found := currentByKey[key]
		if !found {
			b.plan.Ops = append(b.plan.Ops, Op{
				Kind:   OpCreate,
				Target: ref,
				Change: ChangeSet{
					Body: buildEntryCreateBody(entry),
				},
			})
			continue
		}
		delete(currentByKey, key)

		// Decide update vs no-change: compare description + content.
		if cur.Description == entry.Description && cur.Content == entry.Content {
			b.plan.Ops = append(b.plan.Ops, Op{Kind: OpNoChange, Target: ref})
			continue
		}
		paths := []string{}
		if cur.Description != entry.Description {
			paths = append(paths, "spec.description")
		}
		if cur.Content != entry.Content {
			paths = append(paths, "spec.content")
		}
		b.plan.Ops = append(b.plan.Ops, Op{
			Kind:   OpUpdate,
			Target: ref,
			Change: ChangeSet{
				FieldPaths: paths,
				Body:       buildEntryUpdateBody(entry, paths),
			},
		})
	}

	// Anything in currentByKey now is being deleted.
	for _, key := range sortedStringKeys(currentByKey) {
		cur := currentByKey[key]
		b.plan.Ops = append(b.plan.Ops, Op{
			Kind:   OpDelete,
			Target: ResourceRef{Kind: KindMemoryEntry, Parent: layerExtID, ExternalID: key, RowID: cur.RowID},
		})
	}
	return nil
}

// -----------------------------------------------------------------------------
// agents + variations (+ assignments + memory-layer links)
// -----------------------------------------------------------------------------

func (b *planBuilder) buildAgentOps() error {
	for _, extID := range sortedStringKeys(b.cfg.Agents) {
		a := b.cfg.Agents[extID]
		ref := ResourceRef{Kind: KindAgent, ExternalID: extID}

		lr, err := b.lookupAgent(extID)
		if err != nil {
			return err
		}
		op := planResourceSpec(ref, a.Name, a.Labels, a.Spec, extID, lr)
		b.plan.Ops = append(b.plan.Ops, op)

		for _, vExtID := range sortedStringKeys(a.Variations) {
			v := a.Variations[vExtID]
			vref := ResourceRef{Kind: KindVariation, Parent: extID, ExternalID: vExtID}

			var vlr lookupResult
			if lr.Exists {
				vlr, err = b.lookupVariation(extID, vExtID)
				if err != nil {
					return err
				}
			}
			b.plan.Ops = append(b.plan.Ops,
				planResourceSpec(vref, v.Name, v.Labels, v.Spec, vExtID, vlr))

			parent := extID + "/" + vExtID
			if err := b.buildAssignmentOps(parent, v, vlr); err != nil {
				return err
			}
			if err := b.buildMemoryLayerLinkOps(parent, v, vlr); err != nil {
				return err
			}
		}
	}
	return nil
}

func (b *planBuilder) buildAssignmentOps(parent string, v *VariationNode, vLookup lookupResult) error {
	if v.Assignments == nil {
		return nil
	}

	// Current state — pull assignments out of the GET response's info.assignments.
	type currentAssn struct {
		RowID      string
		Kind       string // "tool" | "tool_set" | "agent"
		TargetID   string // canonical id
	}
	var current []currentAssn
	if vLookup.Exists {
		for _, a := range gjson.Get(vLookup.RawJSON, "info.assignments").Array() {
			ca := currentAssn{RowID: a.Get("id").String()}
			if id := a.Get("tool.id").String(); id != "" {
				ca.Kind = "tool"
				ca.TargetID = id
			} else if id := a.Get("toolSet.id").String(); id != "" {
				ca.Kind = "tool_set"
				ca.TargetID = id
			} else if id := a.Get("agent.id").String(); id != "" {
				ca.Kind = "agent"
				ca.TargetID = id
			}
			if ca.Kind != "" {
				current = append(current, ca)
			}
		}
	}

	// Resolve each desired ref to a canonical id (or "will-create") so we can
	// diff against current by canonical id.
	type desiredAssn struct {
		Kind       string // "tool" | "tool_set" | "agent"
		ExternalID string
		CanonID    string // empty if will-create
	}
	var desired []desiredAssn
	for i, ref := range *v.Assignments {
		kind, ext := classifyAssignment(ref)
		if kind == "" {
			b.refErrors = append(b.refErrors, fmt.Errorf("config: agents.%s.assignments[%d]: empty ref", parent, i))
			continue
		}
		canon, err := b.resolveForAssignment(kind, ext)
		if err != nil {
			b.refErrors = append(b.refErrors, err)
			continue
		}
		desired = append(desired, desiredAssn{Kind: kind, ExternalID: ext, CanonID: canon})
	}

	covered := map[string]bool{}
	for _, d := range desired {
		matched := false
		if d.CanonID != "" {
			for _, c := range current {
				if c.Kind == d.Kind && c.TargetID == d.CanonID {
					covered[c.RowID] = true
					matched = true
					break
				}
			}
		}
		if matched {
			b.plan.Ops = append(b.plan.Ops, Op{
				Kind:   OpNoChange,
				Target: ResourceRef{Kind: KindVariationAssignment, Parent: parent, ExternalID: d.Kind + ":" + d.ExternalID},
			})
		} else {
			b.plan.Ops = append(b.plan.Ops, Op{
				Kind:   OpCreate,
				Target: ResourceRef{Kind: KindVariationAssignment, Parent: parent, ExternalID: d.Kind + ":" + d.ExternalID},
				Change: ChangeSet{
					TargetKind:       d.Kind,
					TargetExternalID: d.ExternalID,
				},
			})
		}
	}
	for _, c := range current {
		if covered[c.RowID] {
			continue
		}
		b.plan.Ops = append(b.plan.Ops, Op{
			Kind:   OpDelete,
			Target: ResourceRef{Kind: KindVariationAssignment, Parent: parent, ExternalID: c.Kind + ":" + c.TargetID, RowID: c.RowID},
		})
	}
	return nil
}

func (b *planBuilder) buildMemoryLayerLinkOps(parent string, v *VariationNode, vLookup lookupResult) error {
	if v.MemoryLayers == nil {
		return nil
	}

	type currentLink struct {
		RowID    string
		TargetID string
		Position int
	}
	var current []currentLink
	if vLookup.Exists {
		for _, a := range gjson.Get(vLookup.RawJSON, "info.memoryLayerAssignments").Array() {
			current = append(current, currentLink{
				RowID:    a.Get("id").String(),
				TargetID: a.Get("memoryLayer.id").String(),
				Position: int(a.Get("position").Int()),
			})
		}
	}

	// Resolve each desired external_id.
	type desiredLink struct {
		ExternalID string
		Position   int
		CanonID    string // empty if will-create
	}
	var desired []desiredLink
	for i, extID := range *v.MemoryLayers {
		if extID == "" {
			b.refErrors = append(b.refErrors, fmt.Errorf("config: agents.%s.memory_layers[%d]: empty", parent, i))
			continue
		}
		canon, err := b.resolveMemoryLayer(extID)
		if err != nil {
			b.refErrors = append(b.refErrors, err)
			continue
		}
		desired = append(desired, desiredLink{ExternalID: extID, Position: i, CanonID: canon})
	}

	// Emit add/reorder per desired, then delete per current not in desired.
	coveredRows := map[string]bool{}
	for _, d := range desired {
		var match *currentLink
		if d.CanonID != "" {
			for i := range current {
				if current[i].TargetID == d.CanonID {
					match = &current[i]
					coveredRows[current[i].RowID] = true
					break
				}
			}
		}
		ref := ResourceRef{Kind: KindVariationMemoryLayer, Parent: parent, ExternalID: d.ExternalID}
		if match == nil {
			b.plan.Ops = append(b.plan.Ops, Op{
				Kind:   OpCreate,
				Target: ref,
				Change: ChangeSet{
					TargetKind:       "memory_layer",
					TargetExternalID: d.ExternalID,
					Position:         d.Position,
				},
			})
		} else if match.Position != d.Position {
			ref.RowID = match.RowID
			b.plan.Ops = append(b.plan.Ops, Op{
				Kind:   OpReorder,
				Target: ref,
				Change: ChangeSet{Position: d.Position},
			})
		} else {
			b.plan.Ops = append(b.plan.Ops, Op{Kind: OpNoChange, Target: ref})
		}
	}
	for _, c := range current {
		if coveredRows[c.RowID] {
			continue
		}
		b.plan.Ops = append(b.plan.Ops, Op{
			Kind:   OpDelete,
			Target: ResourceRef{Kind: KindVariationMemoryLayer, Parent: parent, ExternalID: c.TargetID, RowID: c.RowID},
		})
	}
	return nil
}

// -----------------------------------------------------------------------------
// shared planning helper for {metadata, spec} resources
// -----------------------------------------------------------------------------

func planResourceSpec(
	ref ResourceRef,
	name string,
	labels map[string]string,
	spec map[string]any,
	externalID string,
	lr lookupResult,
) Op {
	if !lr.Exists {
		body := buildCreateBody(externalID, name, labels, spec)
		fieldPaths := declaredFieldPaths(name, labels, spec)
		return Op{
			Kind:   OpCreate,
			Target: ref,
			Change: ChangeSet{FieldPaths: fieldPaths, Body: body},
		}
	}

	declared := declaredFieldPaths(name, labels, spec)
	current := gjson.Parse(lr.RawJSON)

	var changed []string
	for _, path := range declared {
		if !fieldEqual(path, desiredValue(path, name, labels, spec), current) {
			changed = append(changed, path)
		}
	}
	if len(changed) == 0 {
		return Op{Kind: OpNoChange, Target: ref}
	}
	return Op{
		Kind:   OpUpdate,
		Target: ref,
		Change: ChangeSet{FieldPaths: changed, Body: buildUpdateBody(name, labels, spec, changed)},
	}
}

func declaredFieldPaths(name string, labels map[string]string, spec map[string]any) []string {
	var paths []string
	if name != "" {
		paths = append(paths, "metadata.name")
	}
	if labels != nil {
		paths = append(paths, "metadata.labels")
	}
	for k := range spec {
		paths = append(paths, "spec."+k)
	}
	sort.Strings(paths)
	return paths
}

func desiredValue(path string, name string, labels map[string]string, spec map[string]any) any {
	switch path {
	case "metadata.name":
		return name
	case "metadata.labels":
		return labels
	}
	const specPrefix = "spec."
	if len(path) > len(specPrefix) && path[:len(specPrefix)] == specPrefix {
		return spec[path[len(specPrefix):]]
	}
	return nil
}

// fieldEqual reports whether the desired value at `path` matches the current
// server state. Proto3 servers elide zero-valued scalars and empty collections
// from their JSON responses (standard protojson behavior), so the comparison
// normalizes both sides by recursively stripping default-valued keys and
// treats a missing path as equivalent to a root-level default.
func fieldEqual(path string, desired any, current gjson.Result) bool {
	cr := current.Get(path)
	if !cr.Exists() {
		// Server didn't echo the field. Match iff desired is a default
		// scalar/empty container or an object whose fields are all defaults.
		return isDefault(stripDefaults(desired))
	}
	var currentVal any
	if err := json.Unmarshal([]byte(cr.Raw), &currentVal); err != nil {
		return false
	}
	dj, err := json.Marshal(stripDefaults(desired))
	if err != nil {
		return false
	}
	cj, err := json.Marshal(stripDefaults(currentVal))
	if err != nil {
		return false
	}
	return string(dj) == string(cj)
}

// stripDefaults walks `v` recursively and removes any map keys whose values
// are proto3 defaults: false, 0, "", nil, empty array, empty object. The
// result is a canonicalized form that matches what a protojson-serialized
// server response (with default-elision) will look like.
func stripDefaults(v any) any {
	switch t := v.(type) {
	case map[string]any:
		out := make(map[string]any, len(t))
		for k, val := range t {
			if isDefault(val) {
				continue
			}
			out[k] = stripDefaults(val)
		}
		return out
	case []any:
		out := make([]any, 0, len(t))
		for _, e := range t {
			out = append(out, stripDefaults(e))
		}
		return out
	default:
		return v
	}
}

// isDefault reports whether v is a proto3 zero value that servers elide from
// JSON responses.
//
// Numeric values arrive as float64 because the inputs to this helper are
// always JSON-unmarshaled into `any` (see fieldEqual). If a caller ever
// passes a typed Go int directly, fold it in here — but today none do.
func isDefault(v any) bool {
	switch t := v.(type) {
	case nil:
		return true
	case bool:
		return !t
	case string:
		return t == ""
	case float64:
		return t == 0
	case map[string]any:
		return len(t) == 0
	case []any:
		return len(t) == 0
	}
	return false
}

func buildCreateBody(externalID, name string, labels map[string]string, spec map[string]any) map[string]any {
	metadata := map[string]any{"externalId": externalID}
	if name != "" {
		metadata["name"] = name
	}
	if labels != nil {
		metadata["labels"] = labels
	}
	body := map[string]any{"metadata": metadata}
	if len(spec) > 0 {
		body["spec"] = spec
	}
	return body
}

func buildUpdateBody(name string, labels map[string]string, spec map[string]any, changedPaths []string) map[string]any {
	body := map[string]any{}
	metadata := map[string]any{}
	for _, path := range changedPaths {
		switch path {
		case "metadata.name":
			metadata["name"] = name
		case "metadata.labels":
			metadata["labels"] = labels
		default:
			const specPrefix = "spec."
			if len(path) > len(specPrefix) && path[:len(specPrefix)] == specPrefix {
				key := path[len(specPrefix):]
				specMap, _ := body["spec"].(map[string]any)
				if specMap == nil {
					specMap = map[string]any{}
					body["spec"] = specMap
				}
				specMap[key] = spec[key]
			}
		}
	}
	if len(metadata) > 0 {
		body["metadata"] = metadata
	}
	return body
}

// -----------------------------------------------------------------------------
// entry body builders
// -----------------------------------------------------------------------------

func buildEntryCreateBody(entry *EntryNode) map[string]any {
	// metadata.name is api:required. Use the key as the human name by default.
	metadata := map[string]any{"name": entry.Key}
	spec := map[string]any{"key": entry.Key}
	if entry.Description != "" {
		spec["description"] = entry.Description
	}
	if entry.Content != "" {
		spec["content"] = entry.Content
	}
	return map[string]any{"metadata": metadata, "spec": spec}
}

func buildEntryUpdateBody(entry *EntryNode, paths []string) map[string]any {
	spec := map[string]any{}
	for _, p := range paths {
		switch p {
		case "spec.description":
			spec["description"] = entry.Description
		case "spec.content":
			spec["content"] = entry.Content
		}
	}
	return map[string]any{"spec": spec}
}

// -----------------------------------------------------------------------------
// lookup helpers
// -----------------------------------------------------------------------------

func (b *planBuilder) lookupToolSet(extID string) (lookupResult, error) {
	return b.lookup("tool_set", extID, func() ([]byte, error) {
		var raw []byte
		_, err := b.client.ToolSets.Get(b.ctx, "external_id:"+extID, option.WithResponseBodyInto(&raw))
		return raw, err
	})
}

func (b *planBuilder) lookupTool(toolSetExtID, toolExtID string) (lookupResult, error) {
	return b.lookup("tool", toolSetExtID+"/"+toolExtID, func() ([]byte, error) {
		var raw []byte
		_, err := b.client.ToolSets.Tools.Get(
			b.ctx,
			"external_id:"+toolSetExtID,
			"external_id:"+toolExtID,
			option.WithResponseBodyInto(&raw),
		)
		return raw, err
	})
}

func (b *planBuilder) lookupMemoryLayer(extID string) (lookupResult, error) {
	return b.lookup("memory_layer", extID, func() ([]byte, error) {
		var raw []byte
		_, err := b.client.MemoryLayers.Get(b.ctx, "external_id:"+extID, option.WithResponseBodyInto(&raw))
		return raw, err
	})
}

func (b *planBuilder) lookupAgent(extID string) (lookupResult, error) {
	return b.lookup("agent", extID, func() ([]byte, error) {
		var raw []byte
		_, err := b.client.Agents.Get(b.ctx, "external_id:"+extID, option.WithResponseBodyInto(&raw))
		return raw, err
	})
}

func (b *planBuilder) lookupVariation(agentExtID, varExtID string) (lookupResult, error) {
	return b.lookup("variation", agentExtID+"/"+varExtID, func() ([]byte, error) {
		var raw []byte
		_, err := b.client.Agents.Variations.Get(
			b.ctx,
			"external_id:"+agentExtID,
			"external_id:"+varExtID,
			option.WithResponseBodyInto(&raw),
		)
		return raw, err
	})
}

func (b *planBuilder) lookup(kind, cacheKey string, do func() ([]byte, error)) (lookupResult, error) {
	k := kind + ":" + cacheKey
	if lr, ok := b.resolved[k]; ok {
		return lr, nil
	}
	raw, err := do()
	if err != nil {
		if isNotFound(err) {
			lr := lookupResult{Exists: false}
			b.resolved[k] = lr
			return lr, nil
		}
		return lookupResult{}, fmt.Errorf("config: lookup %s %q: %w", kind, cacheKey, err)
	}
	lr := lookupResult{
		Exists:      true,
		RawJSON:     string(raw),
		CanonicalID: gjson.GetBytes(raw, "metadata.id").String(),
	}
	b.resolved[k] = lr
	return lr, nil
}

// listEntries paginates all memory entries of a layer and returns their
// (key, description, content, row_id). content is fetched via a per-entry GET
// since List returns only summaries.
func (b *planBuilder) listEntries(layerExtID string) ([]currentEntry, error) {
	parent := "external_id:" + layerExtID
	iter := b.client.MemoryLayers.Entries.ListAutoPaging(
		b.ctx,
		parent,
		cadenya.MemoryLayerEntryListParams{},
	)

	var out []currentEntry
	for iter.Next() {
		e := iter.Current()
		out = append(out, currentEntry{
			RowID:       e.Metadata.ID,
			Key:         e.Spec.Key,
			Description: e.Spec.Description,
		})
	}
	if err := iter.Err(); err != nil {
		return nil, fmt.Errorf("config: list entries for memory_layers.%s: %w", layerExtID, err)
	}
	for i := range out {
		var raw []byte
		_, err := b.client.MemoryLayers.Entries.Get(
			b.ctx,
			parent,
			out[i].RowID,
			option.WithResponseBodyInto(&raw),
		)
		if err != nil {
			return nil, fmt.Errorf("config: get entry %s/%s: %w", layerExtID, out[i].Key, err)
		}
		out[i].Content = gjson.GetBytes(raw, "content").String()
	}
	return out, nil
}

type currentEntry struct {
	RowID       string
	Key         string
	Description string
	Content     string
}

// -----------------------------------------------------------------------------
// reference resolution for variation assignments + memory-layer links
// -----------------------------------------------------------------------------

// resolveForAssignment returns the canonical id of the target of an assignment
// ref, or "" if the target is also being created in this plan (no canonical id
// yet — treat as an add). Errors if the ref cannot be resolved at all.
func (b *planBuilder) resolveForAssignment(kind, extID string) (string, error) {
	switch kind {
	case "tool_set":
		return b.resolveRef(extID, b.configHasToolSet, b.lookupToolSet, "tool_sets")
	case "tool":
		return b.resolveToolRef(extID)
	case "agent":
		return b.resolveRef(extID, b.configHasAgent, b.lookupAgent, "agents")
	}
	return "", fmt.Errorf("config: unknown assignment kind %q", kind)
}

// resolveMemoryLayer resolves a YAML memory_layer external_id reference.
func (b *planBuilder) resolveMemoryLayer(extID string) (string, error) {
	return b.resolveRef(extID, b.configHasMemoryLayer, b.lookupMemoryLayer, "memory_layers")
}

// resolveRef is the shared resolver for any workspace-scoped, single-level
// resource (tool_set, memory_layer, agent). Returns:
//   - (canon, nil) when the target already exists in the workspace.
//   - ("", nil) when the target is in our YAML but doesn't yet exist in the
//     workspace — i.e. it will be created in this plan. Callers treat this
//     sentinel as "definitely an add" during diff.
//   - ("", err) when the target is neither in the YAML nor in the workspace.
//
// `inConfig` checks the YAML presence. `lookup` does the workspace GET.
// `kindLabel` feeds the error message ("tool_sets", "agents", …).
func (b *planBuilder) resolveRef(
	extID string,
	inConfig func(string) bool,
	lookup func(string) (lookupResult, error),
	kindLabel string,
) (string, error) {
	lr, err := lookup(extID)
	if err != nil {
		return "", err
	}
	if lr.Exists {
		return lr.CanonicalID, nil
	}
	if inConfig(extID) {
		return "", nil // will-create
	}
	return "", fmt.Errorf("config: reference to %s.%s is not in this config and does not exist in workspace", kindLabel, extID)
}

// resolveToolRef handles the two-level case: YAML writes { tool: <ext> }
// without naming the parent tool_set. We scan our YAML to find the owning
// tool_set, then apply the same will-create / exists / missing semantics as
// resolveRef — with an extra step that propagates parent-being-created down
// to the tool.
func (b *planBuilder) resolveToolRef(extID string) (string, error) {
	var parentExt string
	for tsExt, ts := range b.cfg.ToolSets {
		if _, ok := ts.Tools[extID]; ok {
			parentExt = tsExt
			break
		}
	}
	if parentExt == "" {
		return "", fmt.Errorf("config: assignment references tool %q but it is not declared under any tool_set in this config", extID)
	}
	// If the parent tool_set is being created in this plan, the tool can't
	// exist yet — propagate the will-create sentinel.
	if tsLR, ok := b.resolved["tool_set:"+parentExt]; ok && !tsLR.Exists {
		return "", nil
	}
	lr, err := b.lookupTool(parentExt, extID)
	if err != nil {
		return "", err
	}
	if !lr.Exists {
		return "", nil // parent exists; tool will be created alongside
	}
	return lr.CanonicalID, nil
}

func (b *planBuilder) configHasToolSet(extID string) bool {
	_, ok := b.cfg.ToolSets[extID]
	return ok
}

func (b *planBuilder) configHasMemoryLayer(extID string) bool {
	_, ok := b.cfg.MemoryLayers[extID]
	return ok
}

func (b *planBuilder) configHasAgent(extID string) bool {
	_, ok := b.cfg.Agents[extID]
	return ok
}

func classifyAssignment(ref *AssignmentRef) (kind, extID string) {
	switch {
	case ref.ToolSet != "":
		return "tool_set", ref.ToolSet
	case ref.Tool != "":
		return "tool", ref.Tool
	case ref.Agent != "":
		return "agent", ref.Agent
	}
	return "", ""
}

// -----------------------------------------------------------------------------
// misc
// -----------------------------------------------------------------------------

// isNotFound detects a 404 from the SDK's error type.
func isNotFound(err error) bool {
	var apiErr *cadenya.Error
	if errors.As(err, &apiErr) && apiErr.StatusCode == 404 {
		return true
	}
	return false
}

func sortedStringKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
