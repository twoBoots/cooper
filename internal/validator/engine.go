package validator

import (
	"os"
	"path/filepath"
	"strings"
)

// ValidateRepository performs full SDD validation over the Cooper project.
func ValidateRepository(rootDir string) ([]ValidationError, error) {
	var allErrors []ValidationError
	cooperDir := filepath.Join(rootDir, ".cooper")

	// 1. Validate Living Specs
	specsDir := filepath.Join(cooperDir, "specs")
	if _, err := os.Stat(specsDir); err == nil {
		_ = filepath.Walk(specsDir, func(path string, info os.FileInfo, err error) error {
			if err != nil || info.IsDir() {
				return nil
			}
			if strings.HasSuffix(path, ".md") {
				errs, err := LintSpecFile(path, false)
				if err == nil {
					allErrors = append(allErrors, errs...)
				}
			}
			return nil
		})
	}

	// 2. Validate Active Tracks: Spec Deltas & Metadata
	activeDir := filepath.Join(cooperDir, "active")
	if entries, err := os.ReadDir(activeDir); err == nil {
		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}
			trackID := entry.Name()
			trackPath := filepath.Join(activeDir, trackID)

			// Metadata check
			metaPath := filepath.Join(trackPath, "metadata.json")
			if _, err := os.Stat(metaPath); err == nil {
				errs, err := ValidateTrackMetadataFile(metaPath, trackID)
				if err == nil {
					allErrors = append(allErrors, errs...)
				}
			}

			// Spec deltas check
			deltasDir := filepath.Join(trackPath, "spec-deltas")
			if _, err := os.Stat(deltasDir); err == nil {
				_ = filepath.Walk(deltasDir, func(path string, info os.FileInfo, err error) error {
					if err != nil || info.IsDir() {
						return nil
					}
					if strings.HasSuffix(path, ".md") {
						errs, err := LintSpecFile(path, true)
						if err == nil {
							allErrors = append(allErrors, errs...)
						}
					}
					return nil
				})
			}
		}
	}

	// 3. Validate Tracks Registry Parity
	parityErrs := ValidateTracksRegistryParity(cooperDir)
	allErrors = append(allErrors, parityErrs...)

	// 4. Audit Markdown Links
	linkErrs := AuditRepositoryLinks(rootDir)
	allErrors = append(allErrors, linkErrs...)

	return allErrors, nil
}
