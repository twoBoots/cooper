package track

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type mockGitExecutor struct {
	commands [][]string
	failOn   string
}

func (m *mockGitExecutor) Run(dir string, args ...string) (string, error) {
	m.commands = append(m.commands, args)
	if m.failOn != "" && len(args) > 0 && args[0] == m.failOn {
		return "", errors.New("simulated git error")
	}
	if len(args) > 0 && args[0] == "rev-parse" {
		return "abc12345", nil
	}
	return "", nil
}

func TestCreateTrack(t *testing.T) {
	tmpDir := t.TempDir()
	cooperDir := filepath.Join(tmpDir, ".cooper")
	_ = os.MkdirAll(filepath.Join(cooperDir, "active"), 0755)
	_ = os.WriteFile(filepath.Join(cooperDir, "tracks.md"), []byte("# Tracks Registry\n\n---\n"), 0644)

	mockGit := &mockGitExecutor{}
	err := CreateTrack(tmpDir, "test-feat", "", "", mockGit)
	if err != nil {
		t.Fatalf("CreateTrack failed: %v", err)
	}

	// Verify track scaffolded
	trackDir := filepath.Join(cooperDir, "active", "test-feat")
	if _, err := os.Stat(filepath.Join(trackDir, "metadata.json")); os.IsNotExist(err) {
		t.Errorf("expected metadata.json to exist in %s", trackDir)
	}
	if _, err := os.Stat(filepath.Join(trackDir, "proposal.md")); os.IsNotExist(err) {
		t.Errorf("expected proposal.md to exist in %s", trackDir)
	}
	if _, err := os.Stat(filepath.Join(trackDir, "design.md")); os.IsNotExist(err) {
		t.Errorf("expected design.md to exist in %s", trackDir)
	}
	if _, err := os.Stat(filepath.Join(trackDir, "plan.md")); os.IsNotExist(err) {
		t.Errorf("expected plan.md to exist in %s", trackDir)
	}
	if _, err := os.Stat(filepath.Join(trackDir, "index.md")); os.IsNotExist(err) {
		t.Errorf("expected index.md to exist in %s", trackDir)
	}

	// Verify tracks.md registration
	tracksData, _ := os.ReadFile(filepath.Join(cooperDir, "tracks.md"))
	if !strings.Contains(string(tracksData), "test-feat") {
		t.Errorf("expected tracks.md to contain registered track, got: %s", string(tracksData))
	}

	// Verify metadata contents
	metaData, _ := os.ReadFile(filepath.Join(trackDir, "metadata.json"))
	var meta map[string]interface{}
	_ = json.Unmarshal(metaData, &meta)
	if meta["track_id"] != "test-feat" || meta["title"] != "test-feat" || meta["type"] != "feature" {
		t.Errorf("unexpected metadata fields: %+v", meta)
	}
}

func TestRecordCheckpoint(t *testing.T) {
	tmpDir := t.TempDir()
	mockGit := &mockGitExecutor{}

	commitSHA, err := RecordCheckpoint(tmpDir, 1, "CLI Scaffold", mockGit)
	if err != nil {
		t.Fatalf("RecordCheckpoint failed: %v", err)
	}
	if commitSHA != "abc12345" {
		t.Errorf("expected commit SHA abc12345, got: %s", commitSHA)
	}

	if len(mockGit.commands) < 2 {
		t.Fatalf("expected git commit and git notes commands, got: %+v", mockGit.commands)
	}

	// Test git commit failure
	failGit := &mockGitExecutor{failOn: "commit"}
	_, err = RecordCheckpoint(tmpDir, 1, "Fail", failGit)
	if err == nil {
		t.Fatal("expected error on git commit failure, got nil")
	}

	// Test git rev-parse failure
	failGit = &mockGitExecutor{failOn: "rev-parse"}
	_, err = RecordCheckpoint(tmpDir, 1, "Fail", failGit)
	if err == nil {
		t.Fatal("expected error on git rev-parse failure, got nil")
	}

	// Test git notes failure
	failGit = &mockGitExecutor{failOn: "notes"}
	_, err = RecordCheckpoint(tmpDir, 1, "Fail", failGit)
	if err == nil {
		t.Fatal("expected error on git notes failure, got nil")
	}
}

func TestCloseTrack(t *testing.T) {
	tmpDir := t.TempDir()
	cooperDir := filepath.Join(tmpDir, ".cooper")
	trackDir := filepath.Join(cooperDir, "active", "test-feat")
	_ = os.MkdirAll(trackDir, 0755)

	metaContent := `{
  "track_id": "test-feat",
  "title": "Test Feature",
  "type": "feature",
  "status": "in_progress",
  "created_at": "2026-08-20T20:00:00Z"
}`
	_ = os.WriteFile(filepath.Join(trackDir, "metadata.json"), []byte(metaContent), 0644)
	tracksContent := `# Tracks Registry

- [ ] **Track: Test Feature** (test-feat)
  - Worktree: .worktrees/test-feat
`
	_ = os.WriteFile(filepath.Join(cooperDir, "tracks.md"), []byte(tracksContent), 0644)

	mockGit := &mockGitExecutor{}
	err := CloseTrack(tmpDir, "test-feat", mockGit)
	if err != nil {
		t.Fatalf("CloseTrack failed: %v", err)
	}

	// Verify metadata status updated to completed
	metaData, _ := os.ReadFile(filepath.Join(trackDir, "metadata.json"))
	var meta map[string]interface{}
	_ = json.Unmarshal(metaData, &meta)
	if meta["status"] != "completed" {
		t.Errorf("expected status 'completed', got: %v", meta["status"])
	}

	// Verify tracks.md checked off
	tracksData, _ := os.ReadFile(filepath.Join(cooperDir, "tracks.md"))
	if !strings.Contains(string(tracksData), "- [x] **Track: Test Feature**") {
		t.Errorf("expected tracks.md checkbox checked off, got: %s", string(tracksData))
	}
}

func TestGetTrackStatus(t *testing.T) {
	tmpDir := t.TempDir()
	cooperDir := filepath.Join(tmpDir, ".cooper")
	trackDir := filepath.Join(cooperDir, "active", "test-status")
	_ = os.MkdirAll(trackDir, 0755)

	metaContent := `{
  "track_id": "test-status",
  "title": "Test Status Track",
  "type": "bugfix",
  "status": "in_progress",
  "created_at": "2026-08-20T20:00:00Z"
}`
	_ = os.WriteFile(filepath.Join(trackDir, "metadata.json"), []byte(metaContent), 0644)
	planContent := `# Implementation Plan
- [x] Task 1
- [~] Task 2
- [ ] Task 3
`
	_ = os.WriteFile(filepath.Join(trackDir, "plan.md"), []byte(planContent), 0644)

	info, err := GetTrackStatus(tmpDir, "test-status")
	if err != nil {
		t.Fatalf("GetTrackStatus failed: %v", err)
	}

	if info.TrackID != "test-status" || info.CompletedTasks != 1 || info.TotalTasks != 3 {
		t.Errorf("unexpected track status info: %+v", info)
	}

	// Test non-existent track
	_, err = GetTrackStatus(tmpDir, "non-existent")
	if err == nil {
		t.Fatal("expected error for non-existent track, got nil")
	}

	// Test invalid JSON metadata
	_ = os.WriteFile(filepath.Join(trackDir, "metadata.json"), []byte("invalid json"), 0644)
	_, err = GetTrackStatus(tmpDir, "test-status")
	if err == nil {
		t.Fatal("expected error for invalid json metadata, got nil")
	}
}

func TestDefaultGitExecutor(t *testing.T) {
	exec := &DefaultGitExecutor{}
	out, err := exec.Run("", "version")
	if err != nil {
		t.Fatalf("git version failed: %v", err)
	}
	if !strings.Contains(out, "git version") {
		t.Errorf("unexpected git output: %s", out)
	}
}
