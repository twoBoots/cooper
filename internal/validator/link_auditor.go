package validator

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var linkRegex = regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)

// AuditMarkdownLinks scans markdown content for links and validates local targets.
func AuditMarkdownLinks(filePath string, content string, rootDir string) []ValidationError {
	var errors []ValidationError
	lines := strings.Split(content, "\n")
	fileDir := filepath.Dir(filePath)

	inCodeBlock := false

	for i, line := range lines {
		lineNum := i + 1
		trimmed := strings.TrimSpace(line)

		// Handle fenced code blocks
		if strings.HasPrefix(trimmed, "```") || strings.HasPrefix(trimmed, "~~~") {
			inCodeBlock = !inCodeBlock
			continue
		}

		if inCodeBlock {
			continue
		}

		matches := linkRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			linkText := match[1]
			linkTarget := strings.TrimSpace(match[2])

			// Skip template placeholders e.g. <track_id>
			if strings.Contains(linkTarget, "<") || strings.Contains(linkTarget, "{") {
				continue
			}

			// Handle troop attribution rule
			if strings.EqualFold(linkText, "Troop") && !strings.Contains(linkTarget, "github.com/twoBoots/troop") {
				errors = append(errors, ValidationError{
					File:    filePath,
					Line:    lineNum,
					Message: "Links with text 'Troop' must target 'https://github.com/twoBoots/troop'",
					Rule:    "link/troop-attribution",
				})
			}

			// Skip web and special schemes
			if strings.HasPrefix(linkTarget, "http://") ||
				strings.HasPrefix(linkTarget, "https://") ||
				strings.HasPrefix(linkTarget, "mailto:") ||
				strings.HasPrefix(linkTarget, "#") {
				continue
			}

			// Clean fragment and query
			targetPath := linkTarget
			if idx := strings.Index(targetPath, "#"); idx != -1 {
				targetPath = targetPath[:idx]
			}
			if idx := strings.Index(targetPath, "?"); idx != -1 {
				targetPath = targetPath[:idx]
			}
			if targetPath == "" {
				continue
			}

			// Resolve path
			var fullPath string
			if filepath.IsAbs(targetPath) {
				fullPath = targetPath
			} else {
				// If target starts with relative to repo root (e.g. .cooper/...)
				candidateRelToFile := filepath.Join(fileDir, targetPath)
				candidateRelToRoot := filepath.Join(rootDir, targetPath)

				if _, err := os.Stat(candidateRelToFile); err == nil {
					fullPath = candidateRelToFile
				} else if _, err := os.Stat(candidateRelToRoot); err == nil {
					fullPath = candidateRelToRoot
				} else {
					fullPath = candidateRelToFile
				}
			}

			if _, err := os.Stat(fullPath); err != nil {
				errors = append(errors, ValidationError{
					File:    filePath,
					Line:    lineNum,
					Message: "Linked target file does not exist: " + linkTarget,
					Rule:    "link/target-exists",
				})
			}
		}
	}

	return errors
}

// AuditRepositoryLinks scans all markdown files in the repository for broken local links.
func AuditRepositoryLinks(rootDir string) []ValidationError {
	var errors []ValidationError

	_ = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			name := info.Name()
			if name == ".git" || name == ".worktrees" || name == "node_modules" || name == "bin" || name == "templates" || name == "assets" {
				return filepath.SkipDir
			}
			return nil
		}

		if strings.HasSuffix(strings.ToLower(path), ".md") {
			data, err := os.ReadFile(path)
			if err == nil {
				fileErrs := AuditMarkdownLinks(path, string(data), rootDir)
				errors = append(errors, fileErrs...)
			}
		}
		return nil
	})

	return errors
}
