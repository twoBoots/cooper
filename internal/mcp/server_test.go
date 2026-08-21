package mcp

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/twoBoots/bender/pkg/mcp"
	"github.com/twoBoots/bender/pkg/updater"
	"github.com/twoBoots/cooper/internal/scaffold"
)

func TestNewCooperServer_InitializeAndListTools(t *testing.T) {
	tmpDir := t.TempDir()
	srv := NewCooperServer("1.0.0", "abc1234", "2026-08-21", tmpDir)

	if srv.Name() != "cooper-mcp" {
		t.Errorf("expected server name cooper-mcp, got %s", srv.Name())
	}

	// Test initialize
	initReq := mcp.Request{
		JSONRPC: "2.0",
		ID:      rawID(1),
		Method:  "initialize",
	}
	resp := srv.HandleRequest(context.Background(), initReq)
	if resp.Error != nil {
		t.Fatalf("unexpected error on initialize: %v", resp.Error)
	}

	// Test tools/list
	listReq := mcp.Request{
		JSONRPC: "2.0",
		ID:      rawID(2),
		Method:  "tools/list",
	}
	listResp := srv.HandleRequest(context.Background(), listReq)
	if listResp.Error != nil {
		t.Fatalf("unexpected error on tools/list: %v", listResp.Error)
	}

	toolsRes, ok := listResp.Result.(mcp.ListToolsResult)
	if !ok {
		t.Fatalf("expected ListToolsResult, got %T", listResp.Result)
	}

	expectedTools := map[string]bool{
		"cooper_get_version":  false,
		"cooper_init_project": false,
		"cooper_track_create": false,
		"cooper_track_status": false,
		"cooper_validate":     false,
		"cooper_self_update":  false,
	}

	for _, tool := range toolsRes.Tools {
		expectedTools[tool.Name] = true
	}

	for name, found := range expectedTools {
		if !found {
			t.Errorf("expected tool %s to be registered", name)
		}
	}
}

func TestNewCooperServer_CallTool_GetVersion(t *testing.T) {
	tmpDir := t.TempDir()
	srv := NewCooperServer("1.2.3", "commit456", "2026-08-21", tmpDir)

	params, _ := json.Marshal(mcp.CallToolParams{
		Name:      "cooper_get_version",
		Arguments: map[string]interface{}{},
	})
	req := mcp.Request{
		JSONRPC: "2.0",
		ID:      rawID(10),
		Method:  "tools/call",
		Params:  params,
	}

	resp := srv.HandleRequest(context.Background(), req)
	if resp.Error != nil {
		t.Fatalf("unexpected error: %v", resp.Error)
	}

	callRes, ok := resp.Result.(mcp.CallToolResult)
	if !ok || len(callRes.Content) == 0 {
		t.Fatalf("expected CallToolResult with content, got %v", resp.Result)
	}

	if !strings.Contains(callRes.Content[0].Text, "1.2.3") || !strings.Contains(callRes.Content[0].Text, "commit456") {
		t.Errorf("expected version and commit in output, got: %s", callRes.Content[0].Text)
	}
}

func TestNewCooperServer_CallTool_InitProject(t *testing.T) {
	tmpDir := t.TempDir()
	srv := NewCooperServer("1.0.0", "abc", "date", tmpDir)

	targetDir := filepath.Join(tmpDir, "my-repo")

	params, _ := json.Marshal(mcp.CallToolParams{
		Name: "cooper_init_project",
		Arguments: map[string]interface{}{
			"path":  targetDir,
			"force": false,
		},
	})
	req := mcp.Request{
		JSONRPC: "2.0",
		ID:      rawID(20),
		Method:  "tools/call",
		Params:  params,
	}

	resp := srv.HandleRequest(context.Background(), req)
	if resp.Error != nil {
		t.Fatalf("unexpected error: %v", resp.Error)
	}

	callRes := resp.Result.(mcp.CallToolResult)
	if callRes.IsError {
		t.Fatalf("tool returned error: %s", callRes.Content[0].Text)
	}

	if _, err := os.Stat(filepath.Join(targetDir, ".cooper", "index.md")); err != nil {
		t.Errorf("expected .cooper/index.md to be created: %v", err)
	}

	// Try init again without force -> should return error result
	resp2 := srv.HandleRequest(context.Background(), req)
	callRes2 := resp2.Result.(mcp.CallToolResult)
	if !callRes2.IsError {
		t.Errorf("expected error on re-init without force, got success")
	}
}

func TestNewCooperServer_CallTool_TrackCreateAndStatus(t *testing.T) {
	tmpDir := t.TempDir()
	srv := NewCooperServer("1.0.0", "abc", "date", tmpDir)

	// First init project
	err := scaffold.InitProject(tmpDir, false)
	if err != nil {
		t.Fatalf("failed to init cooper: %v", err)
	}

	// Error case: missing track_id
	emptyParam, _ := json.Marshal(mcp.CallToolParams{
		Name:      "cooper_track_create",
		Arguments: map[string]interface{}{"path": tmpDir},
	})
	errResp := srv.HandleRequest(context.Background(), mcp.Request{
		JSONRPC: "2.0",
		ID:      rawID(29),
		Method:  "tools/call",
		Params:  emptyParam,
	})
	if !errResp.Result.(mcp.CallToolResult).IsError {
		t.Errorf("expected error for empty track_id")
	}

	// Create track
	params, _ := json.Marshal(mcp.CallToolParams{
		Name: "cooper_track_create",
		Arguments: map[string]interface{}{
			"path":     tmpDir,
			"track_id": "auth-flow",
			"name":     "Authentication Flow",
			"type":     "feature",
		},
	})
	req := mcp.Request{
		JSONRPC: "2.0",
		ID:      rawID(30),
		Method:  "tools/call",
		Params:  params,
	}

	resp := srv.HandleRequest(context.Background(), req)
	if resp.Error != nil {
		t.Fatalf("unexpected error: %v", resp.Error)
	}

	callRes := resp.Result.(mcp.CallToolResult)
	if callRes.IsError {
		t.Fatalf("tool returned error: %s", callRes.Content[0].Text)
	}

	// Query track status (all tracks)
	statusParams, _ := json.Marshal(mcp.CallToolParams{
		Name: "cooper_track_status",
		Arguments: map[string]interface{}{
			"path": tmpDir,
		},
	})
	statusReq := mcp.Request{
		JSONRPC: "2.0",
		ID:      rawID(31),
		Method:  "tools/call",
		Params:  statusParams,
	}

	statusResp := srv.HandleRequest(context.Background(), statusReq)
	if statusResp.Error != nil {
		t.Fatalf("unexpected error: %v", statusResp.Error)
	}

	statusRes := statusResp.Result.(mcp.CallToolResult)
	if !strings.Contains(statusRes.Content[0].Text, "auth-flow") {
		t.Errorf("expected status output to include auth-flow, got: %s", statusRes.Content[0].Text)
	}

	// Query track status for specific track_id
	singleParams, _ := json.Marshal(mcp.CallToolParams{
		Name: "cooper_track_status",
		Arguments: map[string]interface{}{
			"path":     tmpDir,
			"track_id": "auth-flow",
		},
	})
	singleResp := srv.HandleRequest(context.Background(), mcp.Request{
		JSONRPC: "2.0",
		ID:      rawID(32),
		Method:  "tools/call",
		Params:  singleParams,
	})
	singleRes := singleResp.Result.(mcp.CallToolResult)
	if !strings.Contains(singleRes.Content[0].Text, "auth-flow") {
		t.Errorf("expected single track status, got: %s", singleRes.Content[0].Text)
	}

	// Query invalid track
	invParams, _ := json.Marshal(mcp.CallToolParams{
		Name: "cooper_track_status",
		Arguments: map[string]interface{}{
			"path":     tmpDir,
			"track_id": "non-existent",
		},
	})
	invResp := srv.HandleRequest(context.Background(), mcp.Request{
		JSONRPC: "2.0",
		ID:      rawID(33),
		Method:  "tools/call",
		Params:  invParams,
	})
	if !invResp.Result.(mcp.CallToolResult).IsError {
		t.Errorf("expected error for non-existent track")
	}
}

func TestNewCooperServer_CallTool_Validate(t *testing.T) {
	tmpDir := t.TempDir()
	srv := NewCooperServer("1.0.0", "abc", "date", tmpDir)

	err := scaffold.InitProject(tmpDir, false)
	if err != nil {
		t.Fatalf("failed to init cooper: %v", err)
	}

	params, _ := json.Marshal(mcp.CallToolParams{
		Name: "cooper_validate",
		Arguments: map[string]interface{}{
			"path": tmpDir,
		},
	})
	req := mcp.Request{
		JSONRPC: "2.0",
		ID:      rawID(40),
		Method:  "tools/call",
		Params:  params,
	}

	resp := srv.HandleRequest(context.Background(), req)
	if resp.Error != nil {
		t.Fatalf("unexpected error: %v", resp.Error)
	}

	callRes := resp.Result.(mcp.CallToolResult)
	if callRes.IsError {
		t.Fatalf("tool returned error: %s", callRes.Content[0].Text)
	}
	if !strings.Contains(callRes.Content[0].Text, "Validation Passed") {
		t.Errorf("expected validation passed, got: %s", callRes.Content[0].Text)
	}

	// Create invalid spec file to test validation failure branch
	invalidSpec := filepath.Join(tmpDir, ".cooper", "specs", "invalid-cap", "spec.md")
	_ = os.MkdirAll(filepath.Dir(invalidSpec), 0755)
	_ = os.WriteFile(invalidSpec, []byte("broken content without keywords"), 0644)

	respFail := srv.HandleRequest(context.Background(), req)
	callResFail := respFail.Result.(mcp.CallToolResult)
	if !callResFail.IsError {
		t.Errorf("expected validation failure with broken spec file")
	}
}

func TestNewCooperServer_CallTool_SelfUpdate(t *testing.T) {
	tmpDir := t.TempDir()
	srv := NewCooperServer("1.0.0", "abc", "date", tmpDir)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		release := updater.Release{
			TagName: "v2.0.0",
			Name:    "Release v2.0.0",
			Assets:  []updater.Asset{},
		}
		_ = json.NewEncoder(w).Encode(release)
	}))
	defer server.Close()

	client := updater.NewClientWithBaseURL(server.URL)
	SetMCPUpdaterClient(client)
	defer SetMCPUpdaterClient(nil)

	params, _ := json.Marshal(mcp.CallToolParams{
		Name: "cooper_self_update",
		Arguments: map[string]interface{}{
			"check_only": true,
		},
	})
	req := mcp.Request{
		JSONRPC: "2.0",
		ID:      rawID(50),
		Method:  "tools/call",
		Params:  params,
	}

	resp := srv.HandleRequest(context.Background(), req)
	if resp.Error != nil {
		t.Fatalf("unexpected error: %v", resp.Error)
	}

	callRes := resp.Result.(mcp.CallToolResult)
	if callRes.IsError {
		t.Errorf("unexpected error in self_update result: %s", callRes.Content[0].Text)
	}
	if !strings.Contains(callRes.Content[0].Text, "Update available") {
		t.Errorf("expected update available message, got: %s", callRes.Content[0].Text)
	}
}

func TestNewCooperServer_Resources(t *testing.T) {
	tmpDir := t.TempDir()
	err := scaffold.InitProject(tmpDir, false)
	if err != nil {
		t.Fatalf("failed to init cooper: %v", err)
	}

	srv := NewCooperServer("1.0.0", "abc", "date", tmpDir)

	// List resources
	req := mcp.Request{
		JSONRPC: "2.0",
		ID:      rawID(60),
		Method:  "resources/list",
	}
	resp := srv.HandleRequest(context.Background(), req)
	if resp.Error != nil {
		t.Fatalf("unexpected error: %v", resp.Error)
	}

	resList := resp.Result.(mcp.ListResourcesResult)
	if len(resList.Resources) < 2 {
		t.Errorf("expected at least 2 resources registered, got %d", len(resList.Resources))
	}

	// Read index resource
	readParams, _ := json.Marshal(mcp.ReadResourceParams{
		URI: "cooper://index",
	})
	readResp := srv.HandleRequest(context.Background(), mcp.Request{
		JSONRPC: "2.0",
		ID:      rawID(61),
		Method:  "resources/read",
		Params:  readParams,
	})
	if readResp.Error != nil {
		t.Fatalf("unexpected error reading index: %v", readResp.Error)
	}

	// Read tracks resource
	readTracksParams, _ := json.Marshal(mcp.ReadResourceParams{
		URI: "cooper://tracks",
	})
	readTracksResp := srv.HandleRequest(context.Background(), mcp.Request{
		JSONRPC: "2.0",
		ID:      rawID(62),
		Method:  "resources/read",
		Params:  readTracksParams,
	})
	if readTracksResp.Error != nil {
		t.Fatalf("unexpected error reading tracks: %v", readTracksResp.Error)
	}
}

func TestRunMCPServer(t *testing.T) {
	tmpDir := t.TempDir()
	in := bytes.NewBufferString(`{"jsonrpc":"2.0","id":1,"method":"initialize"}` + "\n")
	out := new(bytes.Buffer)

	err := RunMCPServer(in, out, "1.0.0", "commit", "date", tmpDir)
	if err != nil {
		t.Fatalf("unexpected error running MCP server: %v", err)
	}

	if !strings.Contains(out.String(), "cooper-mcp") {
		t.Errorf("expected output to contain cooper-mcp, got: %s", out.String())
	}
}

func rawID(id int) *json.RawMessage {
	data, _ := json.Marshal(id)
	raw := json.RawMessage(data)
	return &raw
}
