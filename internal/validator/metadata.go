package validator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// TrackMetadata defines the metadata.json schema for a Cooper track.
type TrackMetadata struct {
	TrackID   string `json:"track_id"`
	Title     string `json:"title"`
	Type      string `json:"type"`   // "feature" | "bugfix" | "chore" | "rfc"
	Status    string `json:"status"` // "new" | "in_progress" | "completed" | "approved" | "planning"
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty"`
}

// ValidateTrackMetadataFile reads and validates a metadata.json file against schema and expected directory name.
func ValidateTrackMetadataFile(filePath string, expectedTrackID string) ([]ValidationError, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read metadata file %s: %w", filePath, err)
	}
	return ValidateTrackMetadataContent(filePath, expectedTrackID, data), nil
}

// ValidateTrackMetadataContent validates the raw bytes of a metadata.json file.
func ValidateTrackMetadataContent(filename string, expectedTrackID string, data []byte) []ValidationError {
	var errors []ValidationError

	var meta TrackMetadata
	if err := json.Unmarshal(data, &meta); err != nil {
		errors = append(errors, ValidationError{
			File:    filename,
			Line:    1,
			Message: fmt.Sprintf("Invalid JSON syntax: %v", err),
			Rule:    "metadata/json-syntax",
		})
		return errors
	}

	if meta.TrackID == "" {
		errors = append(errors, ValidationError{
			File:    filename,
			Line:    1,
			Message: "Field 'track_id' is required and cannot be empty",
			Rule:    "metadata/track-id-required",
		})
	} else if expectedTrackID != "" && meta.TrackID != expectedTrackID {
		errors = append(errors, ValidationError{
			File:    filename,
			Line:    1,
			Message: fmt.Sprintf("Track ID '%s' does not match directory name '%s'", meta.TrackID, expectedTrackID),
			Rule:    "metadata/track-id-match",
		})
	}

	if strings.TrimSpace(meta.Title) == "" {
		errors = append(errors, ValidationError{
			File:    filename,
			Line:    1,
			Message: "Field 'title' is required and cannot be empty",
			Rule:    "metadata/title-required",
		})
	}

	validTypes := map[string]bool{
		"feature": true,
		"bugfix":  true,
		"chore":   true,
		"rfc":     true,
	}
	if !validTypes[meta.Type] {
		errors = append(errors, ValidationError{
			File:    filename,
			Line:    1,
			Message: fmt.Sprintf("Invalid track type '%s'. Expected one of: feature, bugfix, chore, rfc", meta.Type),
			Rule:    "metadata/type-valid",
		})
	}

	validStatuses := map[string]bool{
		"new":         true,
		"in_progress": true,
		"in-progress": true,
		"completed":   true,
		"approved":    true,
		"planning":    true,
	}
	if !validStatuses[meta.Status] {
		errors = append(errors, ValidationError{
			File:    filename,
			Line:    1,
			Message: fmt.Sprintf("Invalid track status '%s'. Expected one of: new, in_progress, completed, approved, planning", meta.Status),
			Rule:    "metadata/status-valid",
		})
	}

	if strings.TrimSpace(meta.CreatedAt) == "" {
		errors = append(errors, ValidationError{
			File:    filename,
			Line:    1,
			Message: "Field 'created_at' is required and cannot be empty",
			Rule:    "metadata/created-at-required",
		})
	}

	return errors
}

// ValidateTracksRegistryParity checks that all active tracks in .cooper/active are referenced in .cooper/tracks.md.
func ValidateTracksRegistryParity(cooperDir string) []ValidationError {
	var errors []ValidationError

	tracksMdPath := filepath.Join(cooperDir, "tracks.md")
	tracksData, err := os.ReadFile(tracksMdPath)
	if err != nil {
		errors = append(errors, ValidationError{
			File:    tracksMdPath,
			Line:    1,
			Message: fmt.Sprintf("Missing or unreadable tracks registry: %v", err),
			Rule:    "registry/tracks-file",
		})
		return errors
	}
	tracksContent := string(tracksData)

	activeDir := filepath.Join(cooperDir, "active")
	entries, err := os.ReadDir(activeDir)
	if err != nil {
		// Active dir might not exist yet, not necessarily an error
		return errors
	}

	for _, entry := range entries {
		if entry.IsDir() {
			trackID := entry.Name()
			if !strings.Contains(tracksContent, trackID) {
				errors = append(errors, ValidationError{
					File:    tracksMdPath,
					Line:    1,
					Message: fmt.Sprintf("Active track '%s' is not registered in tracks.md", trackID),
					Rule:    "registry/track-parity",
				})
			}
		}
	}

	return errors
}
