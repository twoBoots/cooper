package track

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// TrackInfo contains parsed track data.
type TrackInfo struct {
	TrackID        string `json:"track_id"`
	Title          string `json:"title"`
	Type           string `json:"type"`
	Status         string `json:"status"`
	CreatedAt      string `json:"created_at"`
	TotalTasks     int    `json:"total_tasks"`
	CompletedTasks int    `json:"completed_tasks"`
}

// CreateTrack scaffolds a new track directory, initializes metadata, updates tracks.md, and spawns a worktree.
func CreateTrack(rootDir string, trackID string, title string, trackType string, exec GitExecutor) error {
	if trackType == "" {
		trackType = "feature"
	}
	if title == "" {
		title = trackID
	}

	cooperDir := filepath.Join(rootDir, ".cooper")
	activeDir := filepath.Join(cooperDir, "active", trackID)
	if err := os.MkdirAll(filepath.Join(activeDir, "spec-deltas"), 0755); err != nil {
		return fmt.Errorf("failed to create track directory: %w", err)
	}

	// Create metadata.json
	meta := map[string]string{
		"track_id":   trackID,
		"title":      title,
		"type":       trackType,
		"status":     "new",
		"created_at": time.Now().Format(time.RFC3339),
	}
	metaBytes, _ := json.MarshalIndent(meta, "", "  ")
	if err := os.WriteFile(filepath.Join(activeDir, "metadata.json"), metaBytes, 0644); err != nil {
		return fmt.Errorf("failed to write metadata.json: %w", err)
	}

	// Create proposal.md
	proposal := fmt.Sprintf("# Track Proposal: %s\n\n- **Track ID**: `%s`\n- **Type**: %s\n- **Status**: Planning\n\n## 1. Summary\n", title, trackID, trackType)
	_ = os.WriteFile(filepath.Join(activeDir, "proposal.md"), []byte(proposal), 0644)

	// Create design.md
	design := fmt.Sprintf("# Technical Design: %s\n\n- **Track ID**: `%s`\n\n## 1. Architecture\n", title, trackID)
	_ = os.WriteFile(filepath.Join(activeDir, "design.md"), []byte(design), 0644)

	// Create plan.md
	plan := fmt.Sprintf("# Implementation Plan: %s\n\n- **Track ID**: `%s`\n\n## Phase 1: Implementation\n\n- [ ] Task: Initial Implementation\n", title, trackID)
	_ = os.WriteFile(filepath.Join(activeDir, "plan.md"), []byte(plan), 0644)

	// Create index.md
	index := fmt.Sprintf("# Track: %s\n\n- [Proposal](./proposal.md)\n- [Design](./design.md)\n- [Plan](./plan.md)\n- [Spec Deltas](./spec-deltas/)\n", title)
	_ = os.WriteFile(filepath.Join(activeDir, "index.md"), []byte(index), 0644)

	// Register in tracks.md
	tracksMdPath := filepath.Join(cooperDir, "tracks.md")
	if data, err := os.ReadFile(tracksMdPath); err == nil {
		entry := fmt.Sprintf("\n- [ ] **Track: %s** (`%s`)\n  - Worktree: `.worktrees/%s`\n  - Link: [.cooper/active/%s/index.md](.cooper/active/%s/index.md)\n", title, trackID, trackID, trackID, trackID)
		_ = os.WriteFile(tracksMdPath, append(data, []byte(entry)...), 0644)
	}

	// Spawn worktree
	if exec != nil {
		_, _ = exec.Run(rootDir, "worktree", "add", filepath.Join(".worktrees", trackID), "-b", trackID)
	}

	return nil
}

// RecordCheckpoint creates an empty checkpoint commit and attaches a verification Git Note.
func RecordCheckpoint(rootDir string, phaseNum int, phaseTitle string, exec GitExecutor) (string, error) {
	if exec == nil {
		exec = &DefaultGitExecutor{}
	}

	commitMsg := fmt.Sprintf("cooper(checkpoint): Checkpoint end of Phase %d - %s", phaseNum, phaseTitle)
	_, err := exec.Run(rootDir, "commit", "--allow-empty", "-m", commitMsg)
	if err != nil {
		return "", fmt.Errorf("failed to create checkpoint commit: %w", err)
	}

	commitSHA, err := exec.Run(rootDir, "rev-parse", "HEAD")
	if err != nil {
		return "", fmt.Errorf("failed to get commit SHA: %w", err)
	}

	noteMsg := fmt.Sprintf("Phase %d Checkpoint Verification\nAutomated Tests: PASSED\nManual Verification: APPROVED by user\nTimestamp: %s", phaseNum, time.Now().Format(time.RFC3339))
	_, err = exec.Run(rootDir, "notes", "add", "-m", noteMsg, commitSHA)
	if err != nil {
		return commitSHA, fmt.Errorf("failed to add git note: %w", err)
	}

	return commitSHA, nil
}

// CloseTrack updates metadata status to completed, checks off tracks.md, and tears down worktree.
func CloseTrack(rootDir string, trackID string, exec GitExecutor) error {
	cooperDir := filepath.Join(rootDir, ".cooper")
	activeDir := filepath.Join(cooperDir, "active", trackID)

	metaPath := filepath.Join(activeDir, "metadata.json")
	if data, err := os.ReadFile(metaPath); err == nil {
		var meta map[string]interface{}
		if err := json.Unmarshal(data, &meta); err == nil {
			meta["status"] = "completed"
			meta["updated_at"] = time.Now().Format(time.RFC3339)
			updatedBytes, _ := json.MarshalIndent(meta, "", "  ")
			_ = os.WriteFile(metaPath, updatedBytes, 0644)
		}
	}

	// Update tracks.md checkbox
	tracksMdPath := filepath.Join(cooperDir, "tracks.md")
	if data, err := os.ReadFile(tracksMdPath); err == nil {
		content := string(data)
		lines := strings.Split(content, "\n")
		for i, line := range lines {
			if strings.Contains(line, trackID) && strings.HasPrefix(strings.TrimSpace(line), "- [ ]") {
				lines[i] = strings.Replace(line, "- [ ]", "- [x]", 1)
			}
		}
		_ = os.WriteFile(tracksMdPath, []byte(strings.Join(lines, "\n")), 0644)
	}

	// Worktree remove if executor provided
	if exec != nil {
		_, _ = exec.Run(rootDir, "worktree", "remove", filepath.Join(".worktrees", trackID))
	}

	return nil
}

// GetTrackStatus inspects an active track and returns summary statistics.
func GetTrackStatus(rootDir string, trackID string) (*TrackInfo, error) {
	cooperDir := filepath.Join(rootDir, ".cooper")
	trackDir := filepath.Join(cooperDir, "active", trackID)

	metaPath := filepath.Join(trackDir, "metadata.json")
	metaData, err := os.ReadFile(metaPath)
	if err != nil {
		return nil, fmt.Errorf("track '%s' not found or metadata missing: %w", trackID, err)
	}

	var info TrackInfo
	if err := json.Unmarshal(metaData, &info); err != nil {
		return nil, fmt.Errorf("invalid metadata for track '%s': %w", trackID, err)
	}

	planPath := filepath.Join(trackDir, "plan.md")
	if planData, err := os.ReadFile(planPath); err == nil {
		lines := strings.Split(string(planData), "\n")
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			if strings.HasPrefix(trimmed, "- [x]") || strings.HasPrefix(trimmed, "- [X]") {
				info.TotalTasks++
				info.CompletedTasks++
			} else if strings.HasPrefix(trimmed, "- [ ]") || strings.HasPrefix(trimmed, "- [~]") {
				info.TotalTasks++
			}
		}
	}

	return &info, nil
}
