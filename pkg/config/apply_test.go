package config

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/cadenya/cadenya-go"
	"github.com/cadenya/cadenya-go/option"
)

// recordingServer is a tiny httptest-based stub. Each handler entry matches
// (METHOD, URL path) exactly and responds with a fixed JSON body + status.
// Unmatched requests fail the test so unexpected dispatches are caught.
type recordingServer struct {
	t        *testing.T
	mu       sync.Mutex
	requests []recordedRequest
	handlers map[string]stubResp
}

type recordedRequest struct {
	Method string
	Path   string
	Body   []byte
}

type stubResp struct {
	Status int
	Body   string
}

func newRecordingServer(t *testing.T) *recordingServer {
	return &recordingServer{t: t, handlers: map[string]stubResp{}}
}

func (s *recordingServer) on(method, path string, status int, body string) {
	s.handlers[method+" "+path] = stubResp{Status: status, Body: body}
}

func (s *recordingServer) start() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		s.mu.Lock()
		s.requests = append(s.requests, recordedRequest{Method: r.Method, Path: r.URL.Path, Body: body})
		resp, ok := s.handlers[r.Method+" "+r.URL.Path]
		s.mu.Unlock()
		if !ok {
			s.t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			http.Error(w, `{"code":5,"message":"not found"}`, 404)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.Status)
		io.WriteString(w, resp.Body)
	}))
}

func (s *recordingServer) take() []recordedRequest {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := s.requests
	s.requests = nil
	return out
}

func newTestClient(baseURL string) *cadenya.Client {
	return cadenya.NewClient(option.WithBaseURL(baseURL+"/"), option.WithAPIKey("test-key"))
}

// -----------------------------------------------------------------------------
// Build: bulk ref-error aggregation discards the plan
// -----------------------------------------------------------------------------

func TestBuildReturnsNilPlanOnRefErrors(t *testing.T) {
	// YAML declares an agent with a variation whose assignments and
	// memory_layers reference resources that exist neither in this config
	// nor in the workspace. Build should return nil plan + joined error.
	asns := []*AssignmentRef{
		{ToolSet: "missing_ts_1"},
		{ToolSet: "missing_ts_2"},
	}
	mls := []string{"missing_layer"}
	cfg := &Config{
		Agents: map[string]*AgentNode{
			"ag": {
				ExternalID: "ag",
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

	s := newRecordingServer(t)
	// Agent lookup succeeds (so we reach the variation's ref resolution).
	s.on("GET", "/v1/agents/external_id:ag", 200,
		`{"metadata":{"id":"agent_01","externalId":"ag"}}`)
	// Variation lookup: return 404 (no current state, but agent exists so
	// child resolution proceeds).
	s.on("GET", "/v1/agents/external_id:ag/variations/external_id:default", 404,
		`{"code":5,"message":"not found"}`)
	// Every ref lookup: 404 — nothing in workspace.
	for _, p := range []string{
		"/v1/tool_sets/external_id:missing_ts_1",
		"/v1/tool_sets/external_id:missing_ts_2",
		"/v1/memory_layers/external_id:missing_layer",
	} {
		s.on("GET", p, 404, `{"code":5,"message":"not found"}`)
	}
	srv := s.start()
	defer srv.Close()

	plan, err := Build(context.Background(), newTestClient(srv.URL), cfg)
	if err == nil {
		t.Fatal("expected Build to return a joined ref error")
	}
	if plan != nil {
		t.Errorf("expected nil plan when ref errors occur, got %+v", plan)
	}
	// Joined error should name every broken ref.
	for _, want := range []string{"missing_ts_1", "missing_ts_2", "missing_layer"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("joined error should mention %q, got: %v", want, err)
		}
	}
}

// -----------------------------------------------------------------------------
// Apply dispatch
// -----------------------------------------------------------------------------

func TestApplyDispatchesCreateToolSet(t *testing.T) {
	s := newRecordingServer(t)
	s.on("POST", "/v1/tool_sets", 200,
		`{"metadata":{"id":"toolset_01","name":"X","externalId":"x"}}`)
	srv := s.start()
	defer srv.Close()

	plan := &Plan{
		Ops: []Op{
			{
				Kind:   OpCreate,
				Target: ResourceRef{Kind: KindToolSet, ExternalID: "x"},
				Change: ChangeSet{
					Body: map[string]any{
						"metadata": map[string]any{"externalId": "x", "name": "X"},
						"spec":     map[string]any{"description": "d"},
					},
				},
			},
		},
	}
	var out bytes.Buffer
	result, err := Apply(context.Background(), newTestClient(srv.URL), plan, &out)
	if err != nil {
		t.Fatalf("Apply: %v", err)
	}
	if result.Applied != 1 || result.Unchanged != 0 || result.Failed != 0 {
		t.Errorf("result = %+v, want Applied=1", result)
	}

	reqs := s.take()
	if len(reqs) != 1 {
		t.Fatalf("want 1 request, got %d", len(reqs))
	}
	if reqs[0].Method != "POST" || reqs[0].Path != "/v1/tool_sets" {
		t.Errorf("unexpected request: %+v", reqs[0])
	}
	// Body shape: metadata.externalId + metadata.name + spec.description.
	var body map[string]any
	if err := json.Unmarshal(reqs[0].Body, &body); err != nil {
		t.Fatalf("parse body: %v", err)
	}
	meta, _ := body["metadata"].(map[string]any)
	if meta["externalId"] != "x" || meta["name"] != "X" {
		t.Errorf("metadata wrong: %+v", meta)
	}
}

func TestApplyNoChangeOpsSkipNetwork(t *testing.T) {
	// Handler that fails the test on any request — NoChange ops should not
	// dispatch.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected request for NoChange op: %s %s", r.Method, r.URL.Path)
		w.WriteHeader(500)
	}))
	defer srv.Close()

	plan := &Plan{
		Ops: []Op{
			{Kind: OpNoChange, Target: ResourceRef{Kind: KindToolSet, ExternalID: "a"}},
			{Kind: OpNoChange, Target: ResourceRef{Kind: KindAgent, ExternalID: "b"}},
		},
	}
	var out bytes.Buffer
	result, err := Apply(context.Background(), newTestClient(srv.URL), plan, &out)
	if err != nil {
		t.Fatalf("Apply: %v", err)
	}
	if result.Applied != 0 || result.Unchanged != 2 || result.Failed != 0 {
		t.Errorf("result = %+v, want Unchanged=2", result)
	}
}

func TestApplyStopsOnFirstFailure(t *testing.T) {
	s := newRecordingServer(t)
	s.on("POST", "/v1/tool_sets", 200, `{"metadata":{"id":"ts_01","externalId":"ok"}}`)
	s.on("POST", "/v1/memory_layers", 500, `{"code":13,"message":"internal"}`)
	// Third op should not be dispatched — not registered on purpose.
	srv := s.start()
	defer srv.Close()

	plan := &Plan{
		Ops: []Op{
			{Kind: OpCreate, Target: ResourceRef{Kind: KindToolSet, ExternalID: "ok"}, Change: ChangeSet{Body: map[string]any{"metadata": map[string]any{"externalId": "ok"}}}},
			{Kind: OpCreate, Target: ResourceRef{Kind: KindMemoryLayer, ExternalID: "fail"}, Change: ChangeSet{Body: map[string]any{"metadata": map[string]any{"externalId": "fail"}}}},
			{Kind: OpCreate, Target: ResourceRef{Kind: KindAgent, ExternalID: "never"}, Change: ChangeSet{Body: map[string]any{"metadata": map[string]any{"externalId": "never"}}}},
		},
	}
	var out bytes.Buffer
	result, err := Apply(context.Background(), newTestClient(srv.URL), plan, &out)
	if err == nil {
		t.Fatal("expected error from 500 op")
	}
	if result.Applied != 1 {
		t.Errorf("Applied = %d, want 1", result.Applied)
	}
	if result.Failed != 1 {
		t.Errorf("Failed = %d, want 1", result.Failed)
	}
	// Third op (agent) must NOT have been dispatched. We tolerate SDK
	// retries on the 5xx, so just assert no agent request ever went out.
	reqs := s.take()
	for _, r := range reqs {
		if strings.Contains(r.Path, "/agents") {
			t.Errorf("agent op should not have been dispatched, got: %+v", r)
		}
	}
}

// -----------------------------------------------------------------------------
// Memory-layer link reorder (OpReorder)
// -----------------------------------------------------------------------------

func TestApplyReordersMemoryLayerLink(t *testing.T) {
	s := newRecordingServer(t)
	s.on("PATCH",
		"/v1/agents/external_id:ag/variations/external_id:default/memory_layer_assignments/avml_01",
		200, `{"id":"avml_01","position":2}`)
	srv := s.start()
	defer srv.Close()

	plan := &Plan{
		Ops: []Op{
			{
				Kind: OpReorder,
				Target: ResourceRef{
					Kind:       KindVariationMemoryLayer,
					Parent:     "ag/default",
					ExternalID: "layer_a",
					RowID:      "avml_01",
				},
				Change: ChangeSet{Position: 2},
			},
		},
	}
	var out bytes.Buffer
	result, err := Apply(context.Background(), newTestClient(srv.URL), plan, &out)
	if err != nil {
		t.Fatalf("Apply: %v", err)
	}
	if result.Applied != 1 {
		t.Errorf("Applied = %d, want 1", result.Applied)
	}
	reqs := s.take()
	if len(reqs) != 1 {
		t.Fatalf("want 1 request, got %d", len(reqs))
	}
	var body map[string]any
	if err := json.Unmarshal(reqs[0].Body, &body); err != nil {
		t.Fatalf("parse body: %v", err)
	}
	// Body should carry the new position.
	if pos, ok := body["position"].(float64); !ok || pos != 2 {
		t.Errorf("body position = %v, want 2", body["position"])
	}
}
