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
//
// Canonicals carries every parent resource's canonical id that's known at
// plan time — populated from the GETs we do during Build. Apply extends it
// with canonical ids learned from create responses. Sub-resource dispatchers
// read from Canonicals to resolve the parent segment of their URL, since the
// server's nested-collection POST routes accept only canonical ids.
//
// Keys are "<kind>:<ext_id>" for top-level resources and
// "variation:<agent_ext>/<var_ext>" for variations.
type Plan struct {
	Ops        []Op
	Canonicals map[string]string
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
	TargetKind        string // "tool" | "tool_set" | "agent" | "memory_layer"
	TargetExternalID  string // YAML-declared external_id of the target
	TargetCanonicalID string // resolved at plan time if target already exists; empty means "will be created — look up at apply time"
	Position          int    // memory-layer link position (add or reorder)
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
		plan:     &Plan{Canonicals: map[string]string{}},
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
				tlr, err = b.lookupTool(lr.CanonicalID, extID, tExtID)
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
				vlr, err = b.lookupVariation(lr.CanonicalID, extID, vExtID)
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
			return fmt.Errorf("config: agents.%s.assignments[%d]: empty ref", parent, i)
		}
		canon, err := b.resolveForAssignment(kind, ext)
		if err != nil {
			return err
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
					TargetKind:        d.Kind,
					TargetExternalID:  d.ExternalID,
					TargetCanonicalID: d.CanonID,
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
			return fmt.Errorf("config: agents.%s.memory_layers[%d]: empty", parent, i)
		}
		canon, err := b.resolveMemoryLayer(extID)
		if err != nil {
			return err
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
					TargetKind:        "memory_layer",
					TargetExternalID:  d.ExternalID,
					TargetCanonicalID: d.CanonID,
					Position:          d.Position,
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

func fieldEqual(path string, desired any, current gjson.Result) bool {
	desiredJSON, err := json.Marshal(desired)
	if err != nil {
		return false
	}
	cr := current.Get(path)
	if !cr.Exists() {
		return false
	}
	return jsonBytesEqual(desiredJSON, []byte(cr.Raw))
}

func jsonBytesEqual(a, b []byte) bool {
	var av, bv any
	if err := json.Unmarshal(a, &av); err != nil {
		return false
	}
	if err := json.Unmarshal(b, &bv); err != nil {
		return false
	}
	ar, err := json.Marshal(av)
	if err != nil {
		return false
	}
	br, err := json.Marshal(bv)
	if err != nil {
		return false
	}
	return string(ar) == string(br)
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

func (b *planBuilder) lookupTool(toolSetCanonicalID, toolSetExtID, toolExtID string) (lookupResult, error) {
	return b.lookup("tool", toolSetExtID+"/"+toolExtID, func() ([]byte, error) {
		var raw []byte
		_, err := b.client.ToolSets.Tools.Get(
			b.ctx,
			toolSetCanonicalID,
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

func (b *planBuilder) lookupVariation(agentCanonicalID, agentExtID, varExtID string) (lookupResult, error) {
	return b.lookup("variation", agentExtID+"/"+varExtID, func() ([]byte, error) {
		var raw []byte
		_, err := b.client.AgentVariations.Get(
			b.ctx,
			agentCanonicalID,
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
	if lr.CanonicalID != "" {
		b.plan.Canonicals[k] = lr.CanonicalID
	}
	return lr, nil
}

// listEntries paginates all memory entries of a layer and returns their
// (key, description, content, row_id). content is fetched via a per-entry GET
// since List returns only summaries.
//
// Uses the layer's canonical id (which is already cached — callers only reach
// this code path when lookupMemoryLayer returned Exists=true), since the
// server's nested-collection routes reject the external_id: form.
func (b *planBuilder) listEntries(layerExtID string) ([]currentEntry, error) {
	layerID := b.plan.Canonicals["memory_layer:"+layerExtID]
	if layerID == "" {
		return nil, fmt.Errorf("config: internal: listEntries called with unresolved layer %q", layerExtID)
	}
	iter := b.client.MemoryLayers.Entries.ListAutoPaging(
		b.ctx,
		layerID,
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
			layerID,
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
		return b.resolveRef(KindToolSet, extID, b.lookupToolSet, "tool_sets")
	case "tool":
		// The plan's YAML assignment form is { tool: <tool_external_id> }, which
		// doesn't include the parent tool set. We scan our YAML to find which
		// tool_set owns the tool, then look it up scoped to that tool set.
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
		parentCanon := b.plan.Canonicals["tool_set:"+parentExt]
		if parentCanon == "" {
			// Parent tool_set is being created in this plan; tool can't exist yet.
			return "", nil
		}
		lr, err := b.lookupTool(parentCanon, parentExt, extID)
		if err != nil {
			return "", err
		}
		if !lr.Exists {
			return "", nil
		}
		return lr.CanonicalID, nil
	case "agent":
		return b.resolveRef(KindAgent, extID, b.lookupAgent, "agents")
	}
	return "", fmt.Errorf("config: unknown assignment kind %q", kind)
}

// resolveMemoryLayer resolves a YAML memory_layer external_id reference.
func (b *planBuilder) resolveMemoryLayer(extID string) (string, error) {
	return b.resolveRef(KindMemoryLayer, extID, b.lookupMemoryLayer, "memory_layers")
}

// resolveRef is the shared resolver for any kind with a single-arg lookup
// helper (tool_set, memory_layer, agent). Returns ("", nil) if the target is
// being created in this plan (will-create); (canon, nil) if it already exists
// in the workspace; or ("", err) if it's not findable at all.
func (b *planBuilder) resolveRef(_ ResourceKind, extID string, lookup func(string) (lookupResult, error), mapKey string) (string, error) {
	// Is it in our YAML with plan op != Create (i.e. already exists)?
	switch mapKey {
	case "tool_sets":
		if _, ok := b.cfg.ToolSets[extID]; ok {
			lr, err := lookup(extID)
			if err != nil {
				return "", err
			}
			if lr.Exists {
				return lr.CanonicalID, nil
			}
			return "", nil // will-create
		}
	case "memory_layers":
		if _, ok := b.cfg.MemoryLayers[extID]; ok {
			lr, err := lookup(extID)
			if err != nil {
				return "", err
			}
			if lr.Exists {
				return lr.CanonicalID, nil
			}
			return "", nil // will-create
		}
	case "agents":
		if _, ok := b.cfg.Agents[extID]; ok {
			lr, err := lookup(extID)
			if err != nil {
				return "", err
			}
			if lr.Exists {
				return lr.CanonicalID, nil
			}
			return "", nil
		}
	}
	// Not in our YAML — must already exist in workspace.
	lr, err := lookup(extID)
	if err != nil {
		return "", err
	}
	if !lr.Exists {
		return "", fmt.Errorf("config: reference to %s.%s is not in this config and does not exist in workspace", mapKey, extID)
	}
	return lr.CanonicalID, nil
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
