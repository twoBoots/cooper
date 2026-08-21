package validator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAuditMarkdownLinks_Valid(t *testing.T) {
	tmpDir := t.TempDir()
	targetFile := filepath.Join(tmpDir, "target.md")
	if err := os.WriteFile(targetFile, []byte("# Target"), 0644); err != nil {
		t.Fatalf("failed to write target: %v", err)
	}

	content := `# Document
- Link to [Target](./target.md)
- External link to [Troop](https://github.com/twoBoots/troop)
`
	docPath := filepath.Join(tmpDir, "doc.md")
	errs := AuditMarkdownLinks(docPath, content, tmpDir)
	if len(errs) != 0 {
		t.Fatalf("expected 0 errors for valid links, got: %+v", errs)
	}
}

func TestAuditMarkdownLinks_BrokenRelativeLink(t *testing.T) {
	tmpDir := t.TempDir()
	content := `# Document
- Link to [Missing File](./missing-file.md)
`
	docPath := filepath.Join(tmpDir, "doc.md")
	errs := AuditMarkdownLinks(docPath, content, tmpDir)
	if len(errs) == 0 {
		t.Fatalf("expected error for broken relative link, got 0")
	}
	if errs[0].Rule != "link/target-exists" {
		t.Errorf("expected link/target-exists rule, got: %s", errs[0].Rule)
	}
}

func TestAuditRepositoryLinks(t *testing.T) {
	tmpDir := t.TempDir()
	docPath := filepath.Join(tmpDir, "README.md")
	if err := os.WriteFile(docPath, []byte("# Title\nValid file"), 0644); err != nil {
		t.Fatalf("failed to write file: %v", err)
	}

	errs := AuditRepositoryLinks(tmpDir)
	if len(errs) != 0 {
		t.Fatalf("expected 0 errors, got: %+v", errs)
	}
}

func TestValidateRepository(t *testing.T) {
	tmpDir := t.TempDir()
	cooperDir := filepath.Join(tmpDir, ".cooper")
	specsDir := filepath.Join(cooperDir, "specs", "auth")
	activeDir := filepath.Join(cooperDir, "active", "track-auth")
	deltasDir := filepath.Join(activeDir, "spec-deltas", "auth")

	if err := os.MkdirAll(specsDir, 0755); err != nil {
		t.Fatalf("failed to create specs dir: %v", err)
	}
	if err := os.MkdirAll(deltasDir, 0755); err != nil {
		t.Fatalf("failed to create deltas dir: %v", err)
	}

	// Living spec
	specContent := `# Capability Specification: Auth

## Requirements

### Requirement: Login
The system SHALL authenticate users.

#### Scenario: Login Success
- GIVEN credentials
- WHEN login
- THEN succeed
`
	if err := os.WriteFile(filepath.Join(specsDir, "spec.md"), []byte(specContent), 0644); err != nil {
		t.Fatalf("failed to write spec: %v", err)
	}

	// Delta spec
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
	if err := os.WriteFile(filepath.Join(deltasDir, "spec.md"), []byte(deltaContent), 0644); err != nil {
		t.Fatalf("failed to write delta: %v", err)
	}

	// Metadata
	metaContent := `{
  "track_id": "track-auth",
  "title": "Auth Track",
  "type": "feature",
  "status": "new",
  "created_at": "2026-08-20T20:00:00Z"
}`
	if err := os.WriteFile(filepath.Join(activeDir, "metadata.json"), []byte(metaContent), 0644); err != nil {
		t.Fatalf("failed to write metadata: %v", err)
	}

	// Tracks.md
	tracksContent := `# Tracks Registry
- [ ] **Track: Auth Track** (track-auth)
`
	if err := os.WriteFile(filepath.Join(cooperDir, "tracks.md"), []byte(tracksContent), 0644); err != nil {
		t.Fatalf("failed to write tracks.md: %v", err)
	}

	errs, err := ValidateRepository(tmpDir)
	if err != nil {
		t.Fatalf("unexpected validation error: %v", err)
	}
	if len(errs) != 0 {
		t.Fatalf("expected 0 errors, got: %+v", errs)
	}
}
