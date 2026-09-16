package validator

import (
	"os"
	"path/filepath"
	"testing"
)

// hasRule reports whether any error carries the given rule ID.
func hasRule(errs []ValidationError, rule string) bool {
	for _, e := range errs {
		if e.Rule == rule {
			return true
		}
	}
	return false
}

// TestAuditMarkdownLinks_CodePathReferences covers the cli spec scenarios
// "Detect Dangling Backticked Path Reference" and "Ignore Non-Path Inline Code".
//
// AGENTS.md cites .cooper/COOPER.md in backticks rather than markdown link
// syntax, so a project missing that file previously validated clean.
func TestAuditMarkdownLinks_CodePathReferences(t *testing.T) {
	root := t.TempDir()

	// An existing target that must not be flagged.
	existingDir := filepath.Join(root, ".cooper")
	if err := os.MkdirAll(existingDir, 0755); err != nil {
		t.Fatalf("failed creating .cooper: %v", err)
	}
	if err := os.WriteFile(filepath.Join(existingDir, "COOPER.md"), []byte("# Cooper\n"), 0644); err != nil {
		t.Fatalf("failed writing COOPER.md: %v", err)
	}

	docPath := filepath.Join(root, "AGENTS.md")

	tests := []struct {
		name       string
		content    string
		wantFlag   bool
		wantReason string
	}{
		{
			name:       "dangling backticked markdown path is flagged",
			content:    "Refer to `.cooper/MISSING.md` for details.",
			wantFlag:   true,
			wantReason: "the file does not exist",
		},
		{
			name:     "existing backticked path is not flagged",
			content:  "Refer to `.cooper/COOPER.md` for details.",
			wantFlag: false,
		},
		{
			name:     "shell command with flags is not a path",
			content:  "Run `git agent-start <track_id>` to begin.",
			wantFlag: false,
		},
		{
			name:     "command flag is not a path",
			content:  "Pass `--force` to overwrite.",
			wantFlag: false,
		},
		{
			name:     "glob pattern is not a path",
			content:  "Styleguides live under `.cooper/code_styleguides/*.md` in the repo.",
			wantFlag: false,
		},
		{
			name:     "url in backticks is not a local path",
			content:  "See `https://github.com/twoBoots/troop` for details.",
			wantFlag: false,
		},
		{
			name:     "bare word without a slash is not a path",
			content:  "The `validate` command lints specs.",
			wantFlag: false,
		},
		{
			name:     "piped shell snippet is not a path",
			content:  "Run `curl -fsSL url | bash` to install.",
			wantFlag: false,
		},
		{
			name:     "home-relative path is skipped",
			content:  "Config lives at `~/.claude/settings.json` on disk.",
			wantFlag: false,
		},
		{
			name:     "template placeholder path is skipped",
			content:  "Track files live in `.cooper/active/<track_id>/plan.md` here.",
			wantFlag: false,
		},
		{
			name:     "non-documentation extension is skipped",
			content:  "The entrypoint is `cmd/root.go` in this repo.",
			wantFlag: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			errs := AuditMarkdownLinks(docPath, tc.content, root)
			got := hasRule(errs, "link/code-path-exists")

			if got != tc.wantFlag {
				t.Errorf("content %q: got flagged=%v, want %v (%s)\nerrors: %+v",
					tc.content, got, tc.wantFlag, tc.wantReason, errs)
			}
		})
	}
}

// TestAuditMarkdownLinks_CodePathInFencedBlockIgnored asserts that fenced code
// blocks, which routinely contain example paths that need not exist, are not
// audited.
func TestAuditMarkdownLinks_CodePathInFencedBlockIgnored(t *testing.T) {
	root := t.TempDir()
	docPath := filepath.Join(root, "AGENTS.md")

	content := "Example:\n\n```bash\ncat `.cooper/NOPE.md`\n```\n"

	errs := AuditMarkdownLinks(docPath, content, root)
	if hasRule(errs, "link/code-path-exists") {
		t.Errorf("paths inside fenced code blocks must not be audited, got: %+v", errs)
	}
}

// TestAuditMarkdownLinks_CodePathScopedToOperativeDocuments asserts the rule
// applies only to instruction documents an agent follows.
//
// Track documents are records and proposals: an archived track should cite the
// paths that existed when it ran, and a proposal should be free to describe a
// file that has not been built yet. Running this rule unscoped across the
// repository produced 40 false positives in .cooper/active and .cooper/archive
// against 0 real defects.
func TestAuditMarkdownLinks_CodePathScopedToOperativeDocuments(t *testing.T) {
	root := t.TempDir()
	content := "Refer to `.cooper/MISSING.md` for details."

	audited := []string{
		"AGENTS.md",
		"AGENTS.template.md",
		".cooper/index.md",
		".cooper/COOPER.md",
		"skills/cooper-setup/SKILL.md",
		".agents/skills/cooper-setup/SKILL.md",
	}

	exempt := []string{
		"README.md",
		"docs/install.md",
		".cooper/active/some-track/proposal.md",
		".cooper/active/some-track/plan.md",
		".cooper/archive/old-track/design.md",
		".cooper/specs/installer/spec.md",
		"templates/product.md",
	}

	for _, rel := range audited {
		docPath := filepath.Join(root, filepath.FromSlash(rel))
		errs := AuditMarkdownLinks(docPath, content, root)
		if !hasRule(errs, "link/code-path-exists") {
			t.Errorf("%s is an operative document and must be audited for dangling code paths", rel)
		}
	}

	for _, rel := range exempt {
		docPath := filepath.Join(root, filepath.FromSlash(rel))
		errs := AuditMarkdownLinks(docPath, content, root)
		if hasRule(errs, "link/code-path-exists") {
			t.Errorf("%s is a record, proposal or prose document and must not be audited; got: %+v", rel, errs)
		}
	}
}
