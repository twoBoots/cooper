package scaffold

import (
	"io/fs"
	"strings"
	"testing"
)

func TestEmbeddedAssets(t *testing.T) {
	skills, err := ListSkills()
	if err != nil {
		t.Fatalf("failed to list skills: %v", err)
	}

	expectedSkills := []string{
		"cooper-implement",
		"cooper-new-track",
		"cooper-review",
		"cooper-rfc",
		"cooper-setup",
		"cooper-status",
	}

	for _, expected := range expectedSkills {
		found := false
		for _, s := range skills {
			if s == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected skill %s not found in listed skills: %v", expected, skills)
		}
	}

	// Test reading a skill
	skillContent, err := GetSkill("cooper-implement")
	if err != nil {
		t.Fatalf("failed to get skill cooper-implement: %v", err)
	}
	if !strings.Contains(string(skillContent), "Cooper Implement Skill") {
		t.Errorf("unexpected skill content for cooper-implement: %s", string(skillContent))
	}

	// Test reading a non-existent skill
	_, err = GetSkill("non-existent-skill")
	if err == nil {
		t.Errorf("expected error for non-existent skill, got nil")
	}

	// Test reading a template
	templateContent, err := GetTemplate("product.md")
	if err != nil {
		t.Fatalf("failed to get product.md template: %v", err)
	}
	if !strings.Contains(string(templateContent), "Product Definition") {
		t.Errorf("unexpected product.md template content: %s", string(templateContent))
	}

	// Test reading non-existent template
	_, err = GetTemplate("non-existent.md")
	if err == nil {
		t.Errorf("expected error for non-existent template, got nil")
	}

	// Test reading AGENTS template
	agentsContent, err := GetAgentsTemplate()
	if err != nil {
		t.Fatalf("failed to get AGENTS template: %v", err)
	}
	if !strings.Contains(string(agentsContent), "Cooper SDD Framework") {
		t.Errorf("unexpected AGENTS template content: %s", string(agentsContent))
	}

	// Test GetEmbeddedFS
	fsys := GetEmbeddedFS()
	if _, err := fsys.ReadFile("assets/AGENTS.template.md"); err != nil {
		t.Errorf("failed to read AGENTS.template.md directly from fs: %v", err)
	}

	// Test WalkAssets
	walkCount := 0
	err = WalkAssets(func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		walkCount++
		return nil
	})
	if err != nil {
		t.Fatalf("WalkAssets failed: %v", err)
	}
	if walkCount == 0 {
		t.Errorf("expected walkCount > 0, got %d", walkCount)
	}
}
