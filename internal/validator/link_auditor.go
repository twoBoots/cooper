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

	for i, line := range lines {
		lineNum := i + 1
		matches := linkRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			linkText := match[1]
			linkTarget := match[2]

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

			// Resolve relative path
			var fullPath string
			if filepath.IsAbs(targetPath) {
				fullPath = targetPath
			} else {
				fullPath = filepath.Join(fileDir, targetPath)
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
			if name == ".git" || name == ".worktrees" || name == "node_modules" || name == "bin" {
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
