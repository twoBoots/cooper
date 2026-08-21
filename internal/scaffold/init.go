package scaffold

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// InitProject scaffolds a new Cooper SDD workspace or migrates legacy Conductor/OpenSpec projects.
func InitProject(targetDir string, force bool) error {
	cooperDir := filepath.Join(targetDir, ".cooper")

	// Check if already initialized
	if !force {
		if _, err := os.Stat(cooperDir); err == nil {
			return fmt.Errorf("cooper project already initialized at %s (use --force to overwrite)", targetDir)
		}
	}

	// Create foundational directories
	dirsToCreate := []string{
		filepath.Join(cooperDir, "definition"),
		filepath.Join(cooperDir, "code_styleguides"),
		filepath.Join(cooperDir, "specs"),
		filepath.Join(cooperDir, "active"),
		filepath.Join(cooperDir, "archive"),
		filepath.Join(targetDir, ".agents", "skills"),
	}

	for _, dir := range dirsToCreate {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	// Write AGENTS.md
	agentsTemplate, err := GetAgentsTemplate()
	if err == nil {
		agentsFile := filepath.Join(targetDir, "AGENTS.md")
		if force || !fileExists(agentsFile) {
			if err := os.WriteFile(agentsFile, agentsTemplate, 0644); err != nil {
				return fmt.Errorf("failed to write AGENTS.md: %w", err)
			}
		}
	}

	// Write .cooper/index.md
	indexContent := `# Project Context (.cooper)

## Definition
- [Product Definition](./definition/product.md)
- [Product Guidelines](./definition/product-guidelines.md)
- [Tech Stack](./definition/tech-stack.md)
- [Workflow](./definition/workflow.md)
- [Code Style Guides](./code_styleguides/)

## Living Specifications
- [Capability Specs](./specs/)

## Tracks
- [Tracks Registry](./tracks.md)
- [Active Tracks](./active/)
- [Archive](./archive/)

## Capabilities
- [Agent Skills](../.agents/skills/)
`
	indexPath := filepath.Join(cooperDir, "index.md")
	if force || !fileExists(indexPath) {
		if err := os.WriteFile(indexPath, []byte(indexContent), 0644); err != nil {
			return fmt.Errorf("failed to write index.md: %w", err)
		}
	}

	// Write .cooper/tracks.md
	tracksContent := `# Tracks Registry

All active and completed Cooper tracks are registered below.

---
`
	tracksPath := filepath.Join(cooperDir, "tracks.md")
	if force || !fileExists(tracksPath) {
		if err := os.WriteFile(tracksPath, []byte(tracksContent), 0644); err != nil {
			return fmt.Errorf("failed to write tracks.md: %w", err)
		}
	}

	// Extract templates into .cooper/definition and .cooper/code_styleguides
	_ = WalkAssets(func(assetPath string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}

		if strings.HasPrefix(assetPath, "assets/templates/") {
			rel := strings.TrimPrefix(assetPath, "assets/templates/")
			var destPath string
			if strings.HasPrefix(rel, "code_styleguides/") {
				destPath = filepath.Join(cooperDir, rel)
			} else if strings.HasSuffix(rel, ".md") && !strings.Contains(rel, "spec-") {
				destPath = filepath.Join(cooperDir, "definition", rel)
			}

			if destPath != "" {
				data, readErr := embeddedAssets.ReadFile(assetPath)
				if readErr == nil {
					_ = os.MkdirAll(filepath.Dir(destPath), 0755)
					if force || !fileExists(destPath) {
						_ = os.WriteFile(destPath, data, 0644)
					}
				}
			}
		}

		// Extract skills into .agents/skills/
		if strings.HasPrefix(assetPath, "assets/skills/") {
			rel := strings.TrimPrefix(assetPath, "assets/skills/")
			destPath := filepath.Join(targetDir, ".agents", "skills", rel)
			data, readErr := embeddedAssets.ReadFile(assetPath)
			if readErr == nil {
				_ = os.MkdirAll(filepath.Dir(destPath), 0755)
				if force || !fileExists(destPath) {
					_ = os.WriteFile(destPath, data, 0644)
				}
			}
		}

		return nil
	})

	// Brownfield migration: check for .conductor/tracks or openspec
	migrateConductor(targetDir, cooperDir)

	return nil
}

func migrateConductor(targetDir string, cooperDir string) {
	conductorTracks := filepath.Join(targetDir, ".conductor", "tracks")
	if entries, err := os.ReadDir(conductorTracks); err == nil {
		for _, entry := range entries {
			if entry.IsDir() {
				src := filepath.Join(conductorTracks, entry.Name())
				dst := filepath.Join(cooperDir, "active", entry.Name())
				_ = copyDir(src, dst)
			}
		}
	}
}

func copyDir(src string, dst string) error {
	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(targetPath, 0755)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(targetPath, data, 0644)
	})
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
