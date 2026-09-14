package cmd

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// skillsDir is the single authoritative agent skill tree. The
// slim-cli-validate-first track collapsed the former duplicate under
// internal/scaffold/assets/skills into this one.
const skillsDir = "../skills"

// cooperSkills are the six skills Cooper ships.
var cooperSkills = []string{
	"cooper-setup",
	"cooper-rfc",
	"cooper-new-track",
	"cooper-implement",
	"cooper-review",
	"cooper-status",
}

// requiredMandates must survive the consolidation of the two skill trees. They
// were introduced by the mandate-interactive-question-tools track and existed
// only in the embedded copy, so a naive deletion of that copy would have
// silently dropped them.
var requiredMandates = []string{
	"Interactive Question Protocol",
	"Native File Tools Mandate",
}

func readSkill(t *testing.T, name string) string {
	t.Helper()
	path := filepath.Join(skillsDir, name, "SKILL.md")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed reading %s: %v", path, err)
	}
	return string(data)
}

func TestSkills_RetainInteractiveAndFileToolMandates(t *testing.T) {
	for _, skill := range cooperSkills {
		content := readSkill(t, skill)
		for _, mandate := range requiredMandates {
			if !strings.Contains(content, mandate) {
				t.Errorf("skills/%s/SKILL.md is missing required mandate %q", skill, mandate)
			}
		}
	}
}

// subNumberedH2 matches a sub-numbered heading rendered at H2 level, such as
// "## 6.2 Title". Cooper's skill convention is "## N." for a section and
// "### N.M" for a subsection, so an H2 sub-number is a hierarchy regression.
var subNumberedH2 = regexp.MustCompile(`(?m)^## \d+\.\d+\s`)

func TestSkills_HeadingHierarchyIsWellFormed(t *testing.T) {
	for _, skill := range cooperSkills {
		content := readSkill(t, skill)
		if matches := subNumberedH2.FindAllString(content, -1); len(matches) > 0 {
			for _, m := range matches {
				t.Errorf("skills/%s/SKILL.md has sub-numbered heading at H2 level: %q (expected H3)",
					skill, strings.TrimSpace(m))
			}
		}
	}
}

// TestInstructions_NoFabricatedAttestationTemplate enforces the cli spec
// scenario "No Fabricated Attestation Templates In Agent Instructions".
//
// The deleted track.RecordCheckpoint was faithfully implementing the Git Note
// template in cooper-implement's SKILL.md, which supplied a pre-filled passing
// result and user approval for the agent to copy verbatim. Removing the Go
// function while leaving the instruction intact would reintroduce the defect by
// hand at the next checkpoint.
func TestInstructions_NoFabricatedAttestationTemplate(t *testing.T) {
	forbidden := []string{
		"Automated Tests: PASSED",
		"Manual Verification: APPROVED by user",
	}

	roots := []string{"../skills", "../.agents/skills", "../.cooper/definition"}

	for _, root := range roots {
		if _, err := os.Stat(root); err != nil {
			continue
		}

		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() || !strings.HasSuffix(path, ".md") {
				return err
			}

			data, readErr := os.ReadFile(path)
			if readErr != nil {
				return readErr
			}
			for _, needle := range forbidden {
				if strings.Contains(string(data), needle) {
					t.Errorf("%s supplies a pre-filled verification attestation %q; "+
						"the template must instruct the agent to record actual outcomes",
						filepath.ToSlash(path), needle)
				}
			}
			return nil
		})
		if err != nil {
			t.Fatalf("failed walking %s: %v", root, err)
		}
	}
}

// TestSkills_InstalledCopyMatchesSource guards the one legitimate second copy.
// Cooper dogfoods itself, so .agents/skills holds the installed instance of its
// own skills, exactly as a consumer project would. That copy is expected to
// exist, but it must not drift from the skills/ source -- undetected drift
// between two copies is the defect this track exists to remove.
func TestSkills_InstalledCopyMatchesSource(t *testing.T) {
	for _, skill := range cooperSkills {
		source := readSkill(t, skill)

		installedPath := filepath.Join("..", ".agents", "skills", skill, "SKILL.md")
		installed, err := os.ReadFile(installedPath)
		if err != nil {
			t.Errorf("failed reading installed copy %s: %v", installedPath, err)
			continue
		}

		if string(installed) != source {
			t.Errorf(".agents/skills/%s/SKILL.md has drifted from skills/%s/SKILL.md; re-run the installer to resync",
				skill, skill)
		}
	}
}

// TestSkills_NoDuplicateTree asserts that no SKILL.md exists outside the
// authoritative skills/ tree and its installed .agents/skills counterpart. The
// embedded tree under internal/scaffold/assets previously diverged with no sync
// step or parity check, so the installer and 'cooper init' produced different
// frameworks.
func TestSkills_NoDuplicateTree(t *testing.T) {
	var found []string

	err := filepath.Walk("..", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if info.Name() == ".git" || info.Name() == ".worktrees" {
				return filepath.SkipDir
			}
			return nil
		}
		if info.Name() == "SKILL.md" {
			found = append(found, filepath.ToSlash(path))
		}
		return nil
	})
	if err != nil {
		t.Fatalf("failed walking repository: %v", err)
	}

	for _, path := range found {
		if !strings.Contains(path, "/skills/") {
			t.Errorf("unexpected SKILL.md outside the authoritative tree: %s", path)
			continue
		}
		if strings.Contains(path, "internal/") {
			t.Errorf("duplicate SKILL.md tree still present under internal/: %s", path)
		}
	}
}
