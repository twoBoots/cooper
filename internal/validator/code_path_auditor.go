package validator

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// inlineCodeRegex captures the contents of single-backtick inline code spans.
var inlineCodeRegex = regexp.MustCompile("`([^`\n]+)`")

// auditableExtensions are the documentation file types whose absence is worth
// reporting. Source files are excluded: prose routinely cites illustrative
// module paths that need not exist in the current tree.
var auditableExtensions = []string{".md", ".json", ".yml", ".yaml"}

// shellMetacharacters disqualify a span from being treated as a path. Their
// presence means the span is a command, a glob, or a redirection.
const shellMetacharacters = "*|<>$\"'()[]{}!?;&= \t"

// operativeDocuments are the files whose inline-code path references are
// audited. These are the instruction documents an agent actually follows: a
// dangling reference here sends the agent to a file that is not there.
//
// Everything else is deliberately out of scope. Track documents under
// .cooper/active/ and .cooper/archive/ are records and proposals -- an archived
// track should cite the paths that existed when it ran, and a proposal should
// be free to describe a file that has not been built yet. Prose under docs/ and
// README.md is illustrative. Auditing those produced 40 false positives against
// 0 real defects when this rule was first run across the repository.
var operativeDocuments = []string{
	"AGENTS.md",
	"AGENTS.template.md",
	".cooper/index.md",
	".cooper/COOPER.md",
}

// isOperativeDocument reports whether a file's inline-code references should be
// audited.
func isOperativeDocument(filePath string, rootDir string) bool {
	rel, err := filepath.Rel(rootDir, filePath)
	if err != nil {
		return false
	}
	rel = filepath.ToSlash(rel)

	for _, doc := range operativeDocuments {
		if rel == doc {
			return true
		}
	}

	// Agent skills, in both the source tree and the installed copy.
	if strings.HasSuffix(rel, "/SKILL.md") &&
		(strings.HasPrefix(rel, "skills/") || strings.HasPrefix(rel, ".agents/skills/")) {
		return true
	}

	return false
}

// auditCodePathReferences reports inline-code spans that name a repository
// documentation file which does not exist.
//
// The markdown link auditor only resolves true link syntax, so a reference
// written as `.cooper/COOPER.md` escapes it entirely. That is not hypothetical:
// AGENTS.md cites COOPER.md exactly that way, and a project missing the file
// validated clean.
//
// The rules below are deliberately conservative. A false positive here blocks a
// build over prose, so a span is audited only when it is unambiguously a
// repository path.
func auditCodePathReferences(filePath string, line string, lineNum int, fileDir string, rootDir string) []ValidationError {
	var errors []ValidationError

	if !isOperativeDocument(filePath, rootDir) {
		return nil
	}

	for _, match := range inlineCodeRegex.FindAllStringSubmatch(line, -1) {
		candidate := strings.TrimSpace(match[1])

		if !isAuditableRepositoryPath(candidate) {
			continue
		}

		if resolveRepositoryPath(candidate, fileDir, rootDir) {
			continue
		}

		errors = append(errors, ValidationError{
			File:    filePath,
			Line:    lineNum,
			Message: "Referenced repository path does not exist: " + candidate,
			Rule:    "link/code-path-exists",
		})
	}

	return errors
}

// isAuditableRepositoryPath applies the conservative filter described above.
func isAuditableRepositoryPath(candidate string) bool {
	if candidate == "" {
		return false
	}

	// A repository path always contains a separator. This alone excludes bare
	// command names such as `validate`.
	if !strings.Contains(candidate, "/") {
		return false
	}

	// Commands, globs, redirections and multi-word snippets.
	if strings.ContainsAny(candidate, shellMetacharacters) {
		return false
	}

	// Flags, home-relative and absolute paths, and URLs are out of scope.
	if strings.HasPrefix(candidate, "-") ||
		strings.HasPrefix(candidate, "~") ||
		strings.HasPrefix(candidate, "/") ||
		strings.Contains(candidate, "://") {
		return false
	}

	// Template placeholders are resolved at scaffold time, not now. The
	// metacharacter check already covers <> and {}; this guards the bare form.
	if strings.Contains(candidate, "<") || strings.Contains(candidate, "{") {
		return false
	}

	// Only documentation extensions are audited.
	lower := strings.ToLower(candidate)
	for _, ext := range auditableExtensions {
		if strings.HasSuffix(lower, ext) {
			return true
		}
	}

	return false
}

// resolveRepositoryPath reports whether the candidate exists relative to the
// citing file or to the repository root.
func resolveRepositoryPath(candidate string, fileDir string, rootDir string) bool {
	if _, err := os.Stat(filepath.Join(fileDir, candidate)); err == nil {
		return true
	}
	if _, err := os.Stat(filepath.Join(rootDir, candidate)); err == nil {
		return true
	}
	return false
}
