package validator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidateTrackMetadata_Valid(t *testing.T) {
	content := `{
  "track_id": "test-track",
  "title": "Test Track Title",
  "type": "feature",
  "status": "new",
  "created_at": "2026-08-20T20:55:10+10:00"
}`
	errs := ValidateTrackMetadataContent("active/test-track/metadata.json", "test-track", []byte(content))
	if len(errs) != 0 {
		t.Fatalf("expected 0 errors, got: %+v", errs)
	}
}

func TestValidateTrackMetadata_InvalidJSON(t *testing.T) {
	content := `{ invalid json `
	errs := ValidateTrackMetadataContent("active/test-track/metadata.json", "test-track", []byte(content))
	if len(errs) == 0 {
		t.Fatal("expected errors for invalid JSON, got 0")
	}
	if errs[0].Rule != "metadata/json-syntax" {
		t.Errorf("expected metadata/json-syntax rule, got: %s", errs[0].Rule)
	}
}

func TestValidateTrackMetadata_MissingFields(t *testing.T) {
	content := `{
  "track_id": "",
  "title": "",
  "type": "invalid-type",
  "status": "invalid-status"
}`
	errs := ValidateTrackMetadataContent("active/test-track/metadata.json", "test-track", []byte(content))
	if len(errs) < 4 {
		t.Fatalf("expected at least 4 errors, got: %+v", errs)
	}
}

func TestValidateTrackMetadata_TrackIDMismatch(t *testing.T) {
	content := `{
  "track_id": "different-id",
  "title": "Valid Title",
  "type": "feature",
  "status": "in_progress",
  "created_at": "2026-08-20T20:55:10Z"
}`
	errs := ValidateTrackMetadataContent("active/test-track/metadata.json", "test-track", []byte(content))
	if len(errs) == 0 {
		t.Fatal("expected track_id mismatch error, got 0")
	}
	if errs[0].Rule != "metadata/track-id-match" {
		t.Errorf("expected metadata/track-id-match, got: %s", errs[0].Rule)
	}
}

func TestValidateTrackMetadataFile(t *testing.T) {
	tmpDir := t.TempDir()
	metaPath := filepath.Join(tmpDir, "metadata.json")
	content := `{
  "track_id": "my-track",
  "title": "My Track",
  "type": "chore",
  "status": "completed",
  "created_at": "2026-08-20T20:55:10Z"
}`
	if err := os.WriteFile(metaPath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write metadata file: %v", err)
	}

	errs, err := ValidateTrackMetadataFile(metaPath, "my-track")
	if err != nil {
		t.Fatalf("unexpected read error: %v", err)
	}
	if len(errs) != 0 {
		t.Fatalf("expected 0 errors, got: %+v", errs)
	}

	_, err = ValidateTrackMetadataFile(filepath.Join(tmpDir, "nonexistent.json"), "my-track")
	if err == nil {
		t.Fatalf("expected error for nonexistent file, got nil")
	}
}

func TestValidateTracksRegistryParity(t *testing.T) {
	tmpDir := t.TempDir()
	cooperDir := filepath.Join(tmpDir, ".cooper")
	activeDir := filepath.Join(cooperDir, "active", "test-track-1")
	if err := os.MkdirAll(activeDir, 0755); err != nil {
		t.Fatalf("failed to create active track dir: %v", err)
	}

	// Case 1: Active track present in tracks.md
	tracksMdPath := filepath.Join(cooperDir, "tracks.md")
	tracksContent := `# Tracks Registry

- [ ] **Track: Test Track** (test-track-1)
  - Worktree: .worktrees/test-track-1
`
	if err := os.WriteFile(tracksMdPath, []byte(tracksContent), 0644); err != nil {
		t.Fatalf("failed to write tracks.md: %v", err)
	}

	errs := ValidateTracksRegistryParity(cooperDir)
	if len(errs) != 0 {
		t.Fatalf("expected 0 errors for matching track, got: %+v", errs)
	}

	// Case 2: Active track missing in tracks.md
	if err := os.WriteFile(tracksMdPath, []byte("# Empty Tracks Registry\n"), 0644); err != nil {
		t.Fatalf("failed to write tracks.md: %v", err)
	}

	errs = ValidateTracksRegistryParity(cooperDir)
	if len(errs) == 0 {
		t.Fatalf("expected parity error for missing track in tracks.md, got 0")
	}
}
