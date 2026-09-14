package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// removedCommands are the subcommands deleted by the slim-cli-validate-first
// track. They duplicated the Cooper agent skills or wrapped local shell
// operations, and MUST NOT be reintroduced without a spec delta.
//
// Phase 1 removed "mcp" and "track". Phase 2 Task 2.2 removes "init" as part
// of collapsing the two scaffolders, and adds it to this list as its own Red
// step so that each phase leaves the suite green.
var removedCommands = []string{"track", "mcp"}

// retainedCommands are the three commands that constitute Cooper's
// validate-first surface.
var retainedCommands = []string{"validate", "update", "version"}

func TestRootCmd_RetainsValidateFirstSurface(t *testing.T) {
	rootCmd := NewRootCmd()

	registered := make(map[string]bool)
	for _, c := range rootCmd.Commands() {
		registered[c.Name()] = true
	}

	for _, name := range retainedCommands {
		if !registered[name] {
			t.Errorf("expected command %q to be registered on the root command", name)
		}
	}
}

func TestRootCmd_DoesNotRegisterRemovedCommands(t *testing.T) {
	rootCmd := NewRootCmd()

	registered := make(map[string]bool)
	for _, c := range rootCmd.Commands() {
		registered[c.Name()] = true
		for _, alias := range c.Aliases {
			registered[alias] = true
		}
	}

	for _, name := range removedCommands {
		if registered[name] {
			t.Errorf("command %q was removed by the slim-cli-validate-first track but is still registered", name)
		}
	}
}

// TestNoFabricatedVerificationAttestations enforces the spec requirement
// "Prohibition On Unverified Checkpoint Attestations". The deleted
// track.RecordCheckpoint hardcoded a passing test result and a user approval
// into a Git Note without running a test or prompting anyone. No source file
// may reintroduce that pattern.
func TestNoFabricatedVerificationAttestations(t *testing.T) {
	forbidden := []string{
		"Automated Tests: PASSED",
		"Manual Verification: APPROVED by user",
	}

	repoRoot := ".."

	err := filepath.Walk(repoRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// Track artifacts legitimately quote the forbidden strings when
			// documenting the defect being removed; only Go sources are scanned.
			if info.Name() == ".git" || info.Name() == ".worktrees" || info.Name() == ".cooper" {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		// This test file necessarily contains the literals it forbids.
		if strings.HasSuffix(path, "root_test.go") {
			return nil
		}

		data, readErr := os.ReadFile(path)
		if readErr != nil {
			return readErr
		}
		for _, needle := range forbidden {
			if strings.Contains(string(data), needle) {
				t.Errorf("%s contains fabricated verification attestation %q; checkpoint notes must report actual results", path, needle)
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("failed walking repository: %v", err)
	}
}
