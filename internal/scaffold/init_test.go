package scaffold

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestInitGreenfieldProject(t *testing.T) {
	tmpDir := t.TempDir()

	err := InitProject(tmpDir, false)
	if err != nil {
		t.Fatalf("InitProject failed: %v", err)
	}

	// Verify .cooper structure
	expectedPaths := []string{
		".cooper/index.md",
		".cooper/tracks.md",
		".cooper/definition/product.md",
		".cooper/definition/workflow.md",
		".cooper/code_styleguides/go.md",
		".agents/skills/cooper-implement/SKILL.md",
		".agents/skills/cooper-status/SKILL.md",
		"AGENTS.md",
	}

	for _, relPath := range expectedPaths {
		fullPath := filepath.Join(tmpDir, relPath)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			t.Errorf("expected scaffolded file does not exist: %s", relPath)
		}
	}

	// Verify AGENTS.md content
	agentsData, err := os.ReadFile(filepath.Join(tmpDir, "AGENTS.md"))
	if err != nil {
		t.Fatalf("failed to read AGENTS.md: %v", err)
	}
	if !strings.Contains(string(agentsData), "Cooper SDD Framework") {
		t.Errorf("AGENTS.md does not contain expected template text: %s", string(agentsData))
	}
}

func TestInitProject_AlreadyInitialized(t *testing.T) {
	tmpDir := t.TempDir()

	err := InitProject(tmpDir, false)
	if err != nil {
		t.Fatalf("first InitProject failed: %v", err)
	}

	// Running again without force should return an error or skip safely
	err = InitProject(tmpDir, false)
	if err == nil {
		t.Fatal("expected error when initializing existing project without --force, got nil")
	}

	// Running with force should succeed
	err = InitProject(tmpDir, true)
	if err != nil {
		t.Fatalf("InitProject with force failed: %v", err)
	}
}

func TestInitProject_BrownfieldMigration(t *testing.T) {
	tmpDir := t.TempDir()

	// Simulate legacy .conductor
	conductorDir := filepath.Join(tmpDir, ".conductor", "tracks", "legacy-track")
	_ = os.MkdirAll(conductorDir, 0755)
	_ = os.WriteFile(filepath.Join(conductorDir, "plan.md"), []byte("# Legacy Plan"), 0644)

	err := InitProject(tmpDir, false)
	if err != nil {
		t.Fatalf("InitProject with brownfield migration failed: %v", err)
	}

	// Verify legacy track was migrated to .cooper/active/
	migratedPlan := filepath.Join(tmpDir, ".cooper", "active", "legacy-track", "plan.md")
	if _, err := os.Stat(migratedPlan); os.IsNotExist(err) {
		t.Errorf("expected legacy track to be migrated to %s", migratedPlan)
	}
}
