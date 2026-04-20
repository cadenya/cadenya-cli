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
//
// All parent/target ids are sent in the `external_id:<ext>` form — the server
// resolves them on every route (nested CRUD and nested subresource routes per
// AIP-122), so no client-side canonical-id plumbing is needed.
func Apply(ctx context.Context, client *cadenya.Client, plan *Plan, out io.Writer) (*Result, error) {
	ac := &applyContext{ctx: ctx, client: client}

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

type applyContext struct {
	ctx    context.Context
	client *cadenya.Client
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
	case KindTool:
		_, err = ac.client.ToolSets.Tools.New(
			ac.ctx, "external_id:"+op.Target.Parent,
			cadenya.ToolSetToolNewParams{}, opts...)
	case KindMemoryLayer:
		_, err = ac.client.MemoryLayers.New(ac.ctx, cadenya.MemoryLayerNewParams{}, opts...)
	case KindMemoryEntry:
		_, err = ac.client.MemoryLayers.Entries.New(
			ac.ctx, "external_id:"+op.Target.Parent,
			cadenya.MemoryLayerEntryNewParams{}, opts...)
	case KindAgent:
		_, err = ac.client.Agents.New(ac.ctx, cadenya.AgentNewParams{}, opts...)
	case KindVariation:
		_, err = ac.client.Agents.Variations.New(
			ac.ctx, "external_id:"+op.Target.Parent,
			cadenya.AgentVariationNewParams{}, opts...)
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
		_, err = ac.client.ToolSets.Update(ac.ctx,
			"external_id:"+op.Target.ExternalID,
			cadenya.ToolSetUpdateParams{}, opts...)
	case KindTool:
		_, err = ac.client.ToolSets.Tools.Update(ac.ctx,
			"external_id:"+op.Target.Parent,
			"external_id:"+op.Target.ExternalID,
			cadenya.ToolSetToolUpdateParams{}, opts...)
	case KindMemoryLayer:
		_, err = ac.client.MemoryLayers.Update(ac.ctx,
			"external_id:"+op.Target.ExternalID,
			cadenya.MemoryLayerUpdateParams{}, opts...)
	case KindMemoryEntry:
		_, err = ac.client.MemoryLayers.Entries.Update(ac.ctx,
			"external_id:"+op.Target.Parent,
			op.Target.ExternalID,
			cadenya.MemoryLayerEntryUpdateParams{}, opts...)
	case KindAgent:
		_, err = ac.client.Agents.Update(ac.ctx,
			"external_id:"+op.Target.ExternalID,
			cadenya.AgentUpdateParams{}, opts...)
	case KindVariation:
		_, err = ac.client.Agents.Variations.Update(ac.ctx,
			"external_id:"+op.Target.Parent,
			"external_id:"+op.Target.ExternalID,
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
		err := ac.client.MemoryLayers.Entries.Delete(ac.ctx,
			"external_id:"+op.Target.Parent, op.Target.RowID)
		if err != nil {
			return "", err
		}
		return "deleted", nil
	case KindVariationAssignment:
		agent, variation := splitVariationParent(op.Target.Parent)
		err := ac.client.Agents.Variations.RemoveAssignment(ac.ctx,
			"external_id:"+agent,
			"external_id:"+variation,
			op.Target.RowID)
		if err != nil {
			return "", err
		}
		return "removed", nil
	case KindVariationMemoryLayer:
		agent, variation := splitVariationParent(op.Target.Parent)
		err := ac.client.Agents.Variations.RemoveMemoryLayer(ac.ctx,
			"external_id:"+agent,
			"external_id:"+variation,
			op.Target.RowID)
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
	agent, variation := splitVariationParent(op.Target.Parent)
	payload, err := json.Marshal(map[string]any{"position": op.Change.Position})
	if err != nil {
		return "", err
	}
	_, err = ac.client.Agents.Variations.UpdateMemoryLayer(ac.ctx,
		"external_id:"+agent,
		"external_id:"+variation,
		op.Target.RowID,
		cadenya.AgentVariationUpdateMemoryLayerParams{},
		option.WithRequestBody("application/json", payload),
	)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("reordered → position=%d", op.Change.Position), nil
}

// -----------------------------------------------------------------------------
// variation subresource add helpers
// -----------------------------------------------------------------------------

func (ac *applyContext) addAssignment(op Op) (string, error) {
	agent, variation := splitVariationParent(op.Target.Parent)
	body := map[string]any{}
	target := "external_id:" + op.Change.TargetExternalID
	switch op.Change.TargetKind {
	case "tool":
		body["toolId"] = target
	case "tool_set":
		body["toolSetId"] = target
	case "agent":
		body["subAgentId"] = target
	default:
		return "", fmt.Errorf("config.addAssignment: unknown target kind %q", op.Change.TargetKind)
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	var raw []byte
	_, err = ac.client.Agents.Variations.AddAssignment(ac.ctx,
		"external_id:"+agent,
		"external_id:"+variation,
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
	agent, variation := splitVariationParent(op.Target.Parent)
	body := map[string]any{
		"memoryLayerId": "external_id:" + op.Change.TargetExternalID,
		"position":      op.Change.Position,
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	var raw []byte
	_, err = ac.client.Agents.Variations.AddMemoryLayer(ac.ctx,
		"external_id:"+agent,
		"external_id:"+variation,
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

// splitVariationParent splits "<agent_ext>/<variation_ext>" into (agent, variation).
// A parent without a "/" is treated as a variation with no agent scope.
func splitVariationParent(parent string) (agent, variation string) {
	a, v, ok := strings.Cut(parent, "/")
	if !ok {
		return "", a
	}
	return a, v
}

func createdSummary(raw []byte) string {
	id := gjson.GetBytes(raw, "metadata.id").String()
	if id == "" {
		return "created"
	}
	return "created " + id
}
