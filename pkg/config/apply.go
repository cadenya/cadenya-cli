package config

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/cadenya/cadenya-go"
	"github.com/cadenya/cadenya-go/option"
	"github.com/tidwall/gjson"
)

// Apply executes a Plan against the Cadenya API. Best-effort: stops on the
// first runtime failure and returns the partial Result.
func Apply(ctx context.Context, client *cadenya.Client, plan *Plan, out io.Writer) (*Result, error) {
	ac := &applyContext{
		ctx:        ctx,
		client:     client,
		canonicals: map[string]string{},
	}
	for k, v := range plan.Canonicals {
		ac.canonicals[k] = v
	}

	result := &Result{}
	for _, op := range plan.Ops {
		if op.Kind == OpNoChange {
			result.Unchanged++
			result.Outcomes = append(result.Outcomes, Outcome{Op: op})
			continue
		}

		RenderOpStart(out, op)

		var summary string
		var err error
		switch op.Kind {
		case OpCreate:
			summary, err = ac.dispatchCreate(op)
		case OpUpdate:
			summary, err = ac.dispatchUpdate(op)
		case OpDelete:
			summary, err = ac.dispatchDelete(op)
		case OpReorder:
			summary, err = ac.dispatchReorder(op)
		default:
			err = fmt.Errorf("config.Apply: unknown op kind %v", op.Kind)
		}

		RenderOpEnd(out, op, summary, err)
		result.Outcomes = append(result.Outcomes, Outcome{Op: op, Error: err})
		if err != nil {
			result.Failed++
			return result, err
		}
		result.Applied++
	}
	return result, nil
}

type Result struct {
	Applied   int
	Unchanged int
	Failed    int
	Outcomes  []Outcome
}

type Outcome struct {
	Op    Op
	Error error
}

// applyContext carries per-run state: the SDK client and a canonical-id
// cache. The cache is seeded from Plan.Canonicals (parents that already
// existed at plan time) and extended as create ops complete.
type applyContext struct {
	ctx        context.Context
	client     *cadenya.Client
	canonicals map[string]string
}

// recordCanonical extracts metadata.id from a create response and stores it
// for later sub-resource dispatches.
func (ac *applyContext) recordCanonical(kindKey string, raw []byte) string {
	id := gjson.GetBytes(raw, "metadata.id").String()
	if id != "" {
		ac.canonicals[kindKey] = id
	}
	return id
}

// parentID resolves a sub-resource's parent to a canonical id. Used when the
// server's nested-collection routes don't accept the `external_id:` form.
func (ac *applyContext) parentID(kindKey string) (string, error) {
	if id, ok := ac.canonicals[kindKey]; ok && id != "" {
		return id, nil
	}
	return "", fmt.Errorf("config: internal: no canonical id cached for %q (sub-resource dispatched before its parent was created?)", kindKey)
}

// resolveTargetID returns a canonical id for the target of an assignment or
// memory-layer link. Prefer the plan-time resolution; fall back to the
// apply-time canonicals cache (populated from create responses).
//
// Same-body `external_id:` resolution isn't supported by the server today,
// so we always send canonical ids in these bodies.
func (ac *applyContext) resolveTargetID(kind, extID, planTime string) (string, error) {
	if planTime != "" {
		return planTime, nil
	}
	key := kind + ":" + extID
	if id, ok := ac.canonicals[key]; ok && id != "" {
		return id, nil
	}
	return "", fmt.Errorf("config: internal: no canonical id cached for target %s (target created in same plan should have been recorded)", key)
}

// -----------------------------------------------------------------------------
// dispatchers
// -----------------------------------------------------------------------------

func (ac *applyContext) dispatchCreate(op Op) (string, error) {
	body, err := json.Marshal(op.Change.Body)
	if err != nil {
		return "", fmt.Errorf("config: marshal create body for %s: %w", op.Target.Path(), err)
	}
	opts := []option.RequestOption{option.WithRequestBody("application/json", body)}
	var raw []byte
	opts = append(opts, option.WithResponseBodyInto(&raw))

	switch op.Target.Kind {
	case KindToolSet:
		_, err = ac.client.ToolSets.New(ac.ctx, cadenya.ToolSetNewParams{}, opts...)
		if err == nil {
			ac.recordCanonical("tool_set:"+op.Target.ExternalID, raw)
		}
	case KindTool:
		parent, perr := ac.parentID("tool_set:" + op.Target.Parent)
		if perr != nil {
			return "", perr
		}
		_, err = ac.client.ToolSets.Tools.New(ac.ctx, parent, cadenya.ToolSetToolNewParams{}, opts...)
		if err == nil {
			ac.recordCanonical("tool:"+op.Target.Parent+"/"+op.Target.ExternalID, raw)
		}
	case KindMemoryLayer:
		_, err = ac.client.MemoryLayers.New(ac.ctx, cadenya.MemoryLayerNewParams{}, opts...)
		if err == nil {
			ac.recordCanonical("memory_layer:"+op.Target.ExternalID, raw)
		}
	case KindMemoryEntry:
		parent, perr := ac.parentID("memory_layer:" + op.Target.Parent)
		if perr != nil {
			return "", perr
		}
		_, err = ac.client.MemoryLayers.Entries.New(ac.ctx, parent, cadenya.MemoryLayerEntryNewParams{}, opts...)
	case KindAgent:
		_, err = ac.client.Agents.New(ac.ctx, cadenya.AgentNewParams{}, opts...)
		if err == nil {
			ac.recordCanonical("agent:"+op.Target.ExternalID, raw)
		}
	case KindVariation:
		parent, perr := ac.parentID("agent:" + op.Target.Parent)
		if perr != nil {
			return "", perr
		}
		_, err = ac.client.AgentVariations.New(ac.ctx, parent, cadenya.AgentVariationNewParams{}, opts...)
		if err == nil {
			ac.recordCanonical("variation:"+op.Target.Parent+"/"+op.Target.ExternalID, raw)
		}
	case KindVariationAssignment:
		return ac.addAssignment(op)
	case KindVariationMemoryLayer:
		return ac.addMemoryLayerLink(op)
	default:
		return "", fmt.Errorf("config.dispatchCreate: unsupported kind %v", op.Target.Kind)
	}
	if err != nil {
		return "", err
	}
	return createdSummary(raw), nil
}

func (ac *applyContext) dispatchUpdate(op Op) (string, error) {
	body := map[string]any{}
	for k, v := range op.Change.Body {
		body[k] = v
	}
	body["updateMask"] = strings.Join(op.Change.FieldPaths, ",")
	payload, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("config: marshal update body for %s: %w", op.Target.Path(), err)
	}
	opts := []option.RequestOption{option.WithRequestBody("application/json", payload)}

	switch op.Target.Kind {
	case KindToolSet:
		_, err = ac.client.ToolSets.Update(ac.ctx, "external_id:"+op.Target.ExternalID, cadenya.ToolSetUpdateParams{}, opts...)
	case KindTool:
		parent, perr := ac.parentID("tool_set:" + op.Target.Parent)
		if perr != nil {
			return "", perr
		}
		_, err = ac.client.ToolSets.Tools.Update(
			ac.ctx, parent, "external_id:"+op.Target.ExternalID,
			cadenya.ToolSetToolUpdateParams{}, opts...)
	case KindMemoryLayer:
		_, err = ac.client.MemoryLayers.Update(ac.ctx, "external_id:"+op.Target.ExternalID, cadenya.MemoryLayerUpdateParams{}, opts...)
	case KindMemoryEntry:
		parent, perr := ac.parentID("memory_layer:" + op.Target.Parent)
		if perr != nil {
			return "", perr
		}
		_, err = ac.client.MemoryLayers.Entries.Update(
			ac.ctx, parent, op.Target.ExternalID,
			cadenya.MemoryLayerEntryUpdateParams{}, opts...)
	case KindAgent:
		_, err = ac.client.Agents.Update(ac.ctx, "external_id:"+op.Target.ExternalID, cadenya.AgentUpdateParams{}, opts...)
	case KindVariation:
		parent, perr := ac.parentID("agent:" + op.Target.Parent)
		if perr != nil {
			return "", perr
		}
		_, err = ac.client.AgentVariations.Update(
			ac.ctx, parent, "external_id:"+op.Target.ExternalID,
			cadenya.AgentVariationUpdateParams{}, opts...)
	default:
		return "", fmt.Errorf("config.dispatchUpdate: unsupported kind %v", op.Target.Kind)
	}
	if err != nil {
		return "", err
	}
	return "updated", nil
}

func (ac *applyContext) dispatchDelete(op Op) (string, error) {
	switch op.Target.Kind {
	case KindMemoryEntry:
		parent, perr := ac.parentID("memory_layer:" + op.Target.Parent)
		if perr != nil {
			return "", perr
		}
		err := ac.client.MemoryLayers.Entries.Delete(ac.ctx, parent, op.Target.RowID)
		if err != nil {
			return "", err
		}
		return "deleted", nil
	case KindVariationAssignment:
		variation, perr := ac.parentID("variation:" + op.Target.Parent)
		if perr != nil {
			return "", perr
		}
		err := ac.client.AgentVariations.RemoveAssignment(ac.ctx, variation, op.Target.RowID)
		if err != nil {
			return "", err
		}
		return "removed", nil
	case KindVariationMemoryLayer:
		variation, perr := ac.parentID("variation:" + op.Target.Parent)
		if perr != nil {
			return "", perr
		}
		err := ac.client.AgentVariations.RemoveMemoryLayer(ac.ctx, variation, op.Target.RowID)
		if err != nil {
			return "", err
		}
		return "removed", nil
	}
	return "", fmt.Errorf("config.dispatchDelete: unsupported kind %v", op.Target.Kind)
}

func (ac *applyContext) dispatchReorder(op Op) (string, error) {
	if op.Target.Kind != KindVariationMemoryLayer {
		return "", fmt.Errorf("config.dispatchReorder: unsupported kind %v", op.Target.Kind)
	}
	variation, err := ac.parentID("variation:" + op.Target.Parent)
	if err != nil {
		return "", err
	}
	payload, err := json.Marshal(map[string]any{"position": op.Change.Position})
	if err != nil {
		return "", err
	}
	_, err = ac.client.AgentVariations.UpdateMemoryLayer(
		ac.ctx, variation, op.Target.RowID,
		cadenya.AgentVariationUpdateMemoryLayerParams{},
		option.WithRequestBody("application/json", payload),
	)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("reordered → position=%d", op.Change.Position), nil
}

// -----------------------------------------------------------------------------
// variation-scoped add helpers
// -----------------------------------------------------------------------------

func (ac *applyContext) addAssignment(op Op) (string, error) {
	variation, err := ac.parentID("variation:" + op.Target.Parent)
	if err != nil {
		return "", err
	}
	targetID, err := ac.resolveTargetID(op.Change.TargetKind, op.Change.TargetExternalID, op.Change.TargetCanonicalID)
	if err != nil {
		return "", err
	}
	body := map[string]any{}
	switch op.Change.TargetKind {
	case "tool":
		body["toolId"] = targetID
	case "tool_set":
		body["toolSetId"] = targetID
	case "agent":
		body["subAgentId"] = targetID
	default:
		return "", fmt.Errorf("config.addAssignment: unknown target kind %q", op.Change.TargetKind)
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	var raw []byte
	_, err = ac.client.AgentVariations.AddAssignment(
		ac.ctx, variation,
		cadenya.AgentVariationAddAssignmentParams{},
		option.WithRequestBody("application/json", payload),
		option.WithResponseBodyInto(&raw),
	)
	if err != nil {
		return "", err
	}
	id := gjson.GetBytes(raw, "id").String()
	if id == "" {
		return "added", nil
	}
	return "added " + id, nil
}

func (ac *applyContext) addMemoryLayerLink(op Op) (string, error) {
	variation, err := ac.parentID("variation:" + op.Target.Parent)
	if err != nil {
		return "", err
	}
	targetID, err := ac.resolveTargetID("memory_layer", op.Change.TargetExternalID, op.Change.TargetCanonicalID)
	if err != nil {
		return "", err
	}
	body := map[string]any{
		"memoryLayerId": targetID,
		"position":      op.Change.Position,
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	var raw []byte
	_, err = ac.client.AgentVariations.AddMemoryLayer(
		ac.ctx, variation,
		cadenya.AgentVariationAddMemoryLayerParams{},
		option.WithRequestBody("application/json", payload),
		option.WithResponseBodyInto(&raw),
	)
	if err != nil {
		return "", err
	}
	id := gjson.GetBytes(raw, "id").String()
	if id == "" {
		return fmt.Sprintf("added (position=%d)", op.Change.Position), nil
	}
	return fmt.Sprintf("added %s (position=%d)", id, op.Change.Position), nil
}

func createdSummary(raw []byte) string {
	id := gjson.GetBytes(raw, "metadata.id").String()
	if id == "" {
		return "created"
	}
	return "created " + id
}
