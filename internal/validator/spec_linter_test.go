package validator

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLintSpecContent_ValidLivingSpec(t *testing.T) {
	content := `# Capability Specification: Authentication System

## Requirements

### Requirement: User Login
The authentication system SHALL authenticate users using email and password.

#### Scenario: Successful Login
- GIVEN a registered user with valid credentials
- WHEN the user submits their credentials
- THEN the system MUST issue a session token.
`
	errors := LintSpecContent("specs/auth/spec.md", content, false)
	if len(errors) != 0 {
		t.Fatalf("expected 0 errors for valid living spec, got: %+v", errors)
	}
}

func TestLintSpecContent_ValidSpecDelta(t *testing.T) {
	content := `# Capability Specification Delta: Authentication System

## Requirements

### + Requirement: Multi-Factor Authentication
+ The authentication system SHALL require a TOTP code during login.
+
+ #### + Scenario: Successful 2FA Verification
+ - GIVEN a user who has entered valid credentials
+ - WHEN the user submits a valid 6-digit TOTP code
+ - THEN the system MUST authenticate the user session.
`
	errors := LintSpecContent("active/my-track/spec-deltas/auth/spec.md", content, true)
	if len(errors) != 0 {
		t.Fatalf("expected 0 errors for valid spec delta, got: %+v", errors)
	}
}

func TestLintSpecContent_InvalidLivingSpec_MissingHeader(t *testing.T) {
	content := `## Requirements

### Requirement: User Login
The authentication system SHALL authenticate users.
`
	errors := LintSpecContent("specs/auth/spec.md", content, false)
	if len(errors) == 0 {
		t.Fatal("expected errors for missing capability title, got 0")
	}

	foundTitleError := false
	for _, err := range errors {
		if err.Rule == "spec/title-header" {
			foundTitleError = true
		}
	}
	if !foundTitleError {
		t.Errorf("expected spec/title-header error, got: %+v", errors)
	}
}

func TestLintSpecContent_InvalidLivingSpec_MissingKeywords(t *testing.T) {
	content := `# Capability Specification: Authentication System

## Requirements

### Requirement: User Login
The authentication system will do something.

#### Scenario: Successful Login
- Step 1: User logs in
- Step 2: Session created
`
	errors := LintSpecContent("specs/auth/spec.md", content, false)
	if len(errors) < 2 {
		t.Fatalf("expected multiple errors (missing SHALL/MUST and missing GIVEN/WHEN/THEN), got: %+v", errors)
	}
}

func TestLintSpecContent_InvalidSpecDelta_MissingDeltaPrefix(t *testing.T) {
	content := `# Capability Specification Delta: Authentication System

## Requirements

### Requirement: User Login
The authentication system SHALL authenticate users.

#### Scenario: Successful Login
- GIVEN a user
- WHEN they login
- THEN session created
`
	errors := LintSpecContent("spec-deltas/auth/spec.md", content, true)
	if len(errors) == 0 {
		t.Fatal("expected errors for delta missing + or - prefixes, got 0")
	}

	foundPrefixError := false
	for _, err := range errors {
		if err.Rule == "spec-delta/prefix" {
			foundPrefixError = true
		}
	}
	if !foundPrefixError {
		t.Errorf("expected spec-delta/prefix error, got: %+v", errors)
	}
}

func TestLintSpecFile(t *testing.T) {
	tmpDir := t.TempDir()
	specPath := filepath.Join(tmpDir, "spec.md")
	content := `# Capability Specification: Test

## Requirements

### Requirement: Test Feature
The system SHALL support test features.

#### Scenario: Test Scenario
- GIVEN test state
- WHEN test executed
- THEN test succeeds.
`
	if err := os.WriteFile(specPath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write test spec file: %v", err)
	}

	errs, err := LintSpecFile(specPath, false)
	if err != nil {
		t.Fatalf("unexpected read error: %v", err)
	}
	if len(errs) != 0 {
		t.Fatalf("expected 0 errors, got: %+v", errs)
	}
}

func TestLintSpecFile_NotFound(t *testing.T) {
	_, err := LintSpecFile("non-existent-path.md", false)
	if err == nil {
		t.Fatalf("expected error for non-existent file, got nil")
	}
}
