package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestValidateCmd_CleanRepo(t *testing.T) {
	tmpDir := t.TempDir()
	cooperDir := filepath.Join(tmpDir, ".cooper")
	specsDir := filepath.Join(cooperDir, "specs", "auth")
	activeDir := filepath.Join(cooperDir, "active", "track-auth")
	deltasDir := filepath.Join(activeDir, "spec-deltas", "auth")

	_ = os.MkdirAll(specsDir, 0755)
	_ = os.MkdirAll(deltasDir, 0755)

	specContent := `# Capability Specification: Auth

## Requirements

### Requirement: Login
The system SHALL authenticate users.

#### Scenario: Login Success
- GIVEN credentials
- WHEN login
- THEN succeed
`
	_ = os.WriteFile(filepath.Join(specsDir, "spec.md"), []byte(specContent), 0644)

	deltaContent := `# Capability Specification Delta: Auth

## Requirements

### + Requirement: 2FA
+ The system SHALL require 2FA.
+
+ #### + Scenario: 2FA Success
+ - GIVEN user logged in
+ - WHEN 2FA code entered
+ - THEN session authenticated
`
	_ = os.WriteFile(filepath.Join(deltasDir, "spec.md"), []byte(deltaContent), 0644)

	metaContent := `{
  "track_id": "track-auth",
  "title": "Auth Track",
  "type": "feature",
  "status": "new",
  "created_at": "2026-08-20T20:00:00Z"
}`
	_ = os.WriteFile(filepath.Join(activeDir, "metadata.json"), []byte(metaContent), 0644)

	tracksContent := `# Tracks Registry
- [ ] **Track: Auth Track** (track-auth)
`
	_ = os.WriteFile(filepath.Join(cooperDir, "tracks.md"), []byte(tracksContent), 0644)

	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"validate", "--dir", tmpDir})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected validate failure on clean repo: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "Validation passed successfully") {
		t.Errorf("expected success message, got: %s", out)
	}
}

func TestValidateCmd_JSONOutput(t *testing.T) {
	tmpDir := t.TempDir()
	cooperDir := filepath.Join(tmpDir, ".cooper")
	_ = os.MkdirAll(cooperDir, 0755)
	_ = os.WriteFile(filepath.Join(cooperDir, "tracks.md"), []byte("# Tracks\n"), 0644)

	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"lint", "--dir", tmpDir, "--json"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected validate error with json: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, `"errors"`) {
		t.Errorf("expected json output with errors field, got: %s", out)
	}
}

func TestValidateCmd_ErrorsFound(t *testing.T) {
	tmpDir := t.TempDir()
	cooperDir := filepath.Join(tmpDir, ".cooper")
	specsDir := filepath.Join(cooperDir, "specs", "auth")
	_ = os.MkdirAll(specsDir, 0755)
	_ = os.WriteFile(filepath.Join(cooperDir, "tracks.md"), []byte("# Tracks\n"), 0644)

	// Broken spec
	_ = os.WriteFile(filepath.Join(specsDir, "spec.md"), []byte("invalid spec"), 0644)

	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"validate", "--dir", tmpDir})

	err := rootCmd.Execute()
	if err == nil {
		t.Fatal("expected error on broken spec, got nil")
	}
}
