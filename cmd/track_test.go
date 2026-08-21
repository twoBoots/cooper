package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTrackCmd_Help(t *testing.T) {
	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"track", "--help"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected track help error: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "new") || !strings.Contains(out, "status") || !strings.Contains(out, "close") {
		t.Errorf("expected track subcommands in help output, got: %s", out)
	}
}

func TestTrackNewCmd(t *testing.T) {
	tmpDir := t.TempDir()
	cooperDir := filepath.Join(tmpDir, ".cooper")
	_ = os.MkdirAll(filepath.Join(cooperDir, "active"), 0755)
	_ = os.WriteFile(filepath.Join(cooperDir, "tracks.md"), []byte("# Tracks\n"), 0644)

	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"track", "new", "test-new", "--title", "New Track", "--type", "feature", "--dir", tmpDir})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error running track new: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "Created track 'test-new'") {
		t.Errorf("expected creation output, got: %s", out)
	}
}

func TestTrackStatusCmd(t *testing.T) {
	tmpDir := t.TempDir()
	cooperDir := filepath.Join(tmpDir, ".cooper")
	trackDir := filepath.Join(cooperDir, "active", "my-track")
	_ = os.MkdirAll(trackDir, 0755)

	metaContent := `{
  "track_id": "my-track",
  "title": "My Track",
  "type": "feature",
  "status": "in_progress",
  "created_at": "2026-08-20T20:00:00Z"
}`
	_ = os.WriteFile(filepath.Join(trackDir, "metadata.json"), []byte(metaContent), 0644)
	planContent := `# Implementation Plan
- [x] Task 1
- [ ] Task 2
`
	_ = os.WriteFile(filepath.Join(trackDir, "plan.md"), []byte(planContent), 0644)

	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"track", "status", "my-track", "--dir", tmpDir})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected track status error: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "my-track") || !strings.Contains(out, "in_progress") {
		t.Errorf("expected track status output, got: %s", out)
	}

	// Test JSON output
	buf.Reset()
	rootCmd.SetArgs([]string{"track", "status", "my-track", "--dir", tmpDir, "--json"})
	err = rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected track status --json error: %v", err)
	}
	if !strings.Contains(buf.String(), `"track_id": "my-track"`) {
		t.Errorf("expected json output with track_id, got: %s", buf.String())
	}
}

func TestTrackCloseCmd(t *testing.T) {
	tmpDir := t.TempDir()
	cooperDir := filepath.Join(tmpDir, ".cooper")
	trackDir := filepath.Join(cooperDir, "active", "my-track")
	_ = os.MkdirAll(trackDir, 0755)

	metaContent := `{
  "track_id": "my-track",
  "title": "My Track",
  "type": "feature",
  "status": "in_progress",
  "created_at": "2026-08-20T20:00:00Z"
}`
	_ = os.WriteFile(filepath.Join(trackDir, "metadata.json"), []byte(metaContent), 0644)
	_ = os.WriteFile(filepath.Join(cooperDir, "tracks.md"), []byte("- [ ] **Track: My Track** (my-track)\n"), 0644)

	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"track", "close", "my-track", "--dir", tmpDir})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error running track close: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "Closed track 'my-track'") {
		t.Errorf("expected close output, got: %s", out)
	}
}

func TestTrackCheckpointCmd(t *testing.T) {
	tmpDir := t.TempDir()

	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	// Invalid phase number error test
	rootCmd.SetArgs([]string{"track", "checkpoint", "not-a-number", "--dir", tmpDir})
	err := rootCmd.Execute()
	if err == nil {
		t.Fatal("expected error on non-integer phase number, got nil")
	}
}

func TestTrackCmd_Errors(t *testing.T) {
	tmpDir := t.TempDir()
	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)

	// Status on non-existent track
	rootCmd.SetArgs([]string{"track", "status", "non-existent", "--dir", tmpDir})
	err := rootCmd.Execute()
	if err == nil {
		t.Error("expected error getting status of non-existent track, got nil")
	}

	// New track invalid dir
	rootCmd.SetArgs([]string{"track", "new", "test", "--dir", "/dev/null/invalid"})
	_ = rootCmd.Execute()

	// Close track on empty dir
	rootCmd.SetArgs([]string{"track", "close", "non-existent", "--dir", tmpDir})
	_ = rootCmd.Execute()
}
