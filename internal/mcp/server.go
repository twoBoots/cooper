package mcp

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/twoBoots/bender/pkg/mcp"
	"github.com/twoBoots/bender/pkg/updater"
	"github.com/twoBoots/cooper/internal/scaffold"
	"github.com/twoBoots/cooper/internal/track"
	"github.com/twoBoots/cooper/internal/validator"
)

var (
	mcpUpdaterClientOverride *updater.Client
)

// SetMCPUpdaterClient overrides the updater client for MCP tool testing.
func SetMCPUpdaterClient(c *updater.Client) {
	mcpUpdaterClientOverride = c
}

// NewCooperServer creates an MCP Server with all Cooper SDD tools and resources registered.
func NewCooperServer(version, commit, date, cwd string) *mcp.Server {
	if cwd == "" {
		cwd = "."
	}

	srv := mcp.NewServer("cooper-mcp", version, cwd)

	// Tool 1: cooper_get_version
	srv.RegisterTool(mcp.Tool{
		Name:        "cooper_get_version",
		Description: "Returns the current Cooper CLI version, commit hash, and build timestamp",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
		},
	}, func(ctx context.Context, args map[string]interface{}) (mcp.CallToolResult, error) {
		text := fmt.Sprintf("Cooper SDD CLI v%s (commit: %s, built at: %s)", version, commit, date)
		return mcp.NewTextResult(text, false), nil
	})

	// Tool 2: cooper_init_project
	srv.RegisterTool(mcp.Tool{
		Name:        "cooper_init_project",
		Description: "Scaffold standard .cooper/ directory structure, living specs, and agent skills in target repository",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"path": map[string]interface{}{
					"type":        "string",
					"description": "Target project root directory (defaults to current working directory)",
				},
				"force": map[string]interface{}{
					"type":        "boolean",
					"description": "Overwrite existing files if already initialized",
				},
			},
		},
	}, func(ctx context.Context, args map[string]interface{}) (mcp.CallToolResult, error) {
		targetDir := cwd
		if p, ok := args["path"].(string); ok && strings.TrimSpace(p) != "" {
			targetDir = p
		}
		force, _ := args["force"].(bool)

		err := scaffold.InitProject(targetDir, force)
		if err != nil {
			return mcp.NewErrorResult(fmt.Sprintf("Failed to initialize Cooper project: %v", err)), nil
		}

		return mcp.NewTextResult(fmt.Sprintf("Successfully initialized Cooper SDD project at %s", targetDir), false), nil
	})

	// Tool 3: cooper_track_create
	srv.RegisterTool(mcp.Tool{
		Name:        "cooper_track_create",
		Description: "Create a new Cooper SDD track workspace, metadata.json, proposal.md, design.md, and spec-deltas",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"path": map[string]interface{}{
					"type":        "string",
					"description": "Target project root directory (defaults to current working directory)",
				},
				"track_id": map[string]interface{}{
					"type":        "string",
					"description": "Unique kebab-case track identifier (e.g. auth-flow)",
				},
				"name": map[string]interface{}{
					"type":        "string",
					"description": "Human-readable track title",
				},
				"type": map[string]interface{}{
					"type":        "string",
					"description": "Track type: feature, bugfix, rfc, chore (default: feature)",
				},
			},
			Required: []string{"track_id"},
		},
	}, func(ctx context.Context, args map[string]interface{}) (mcp.CallToolResult, error) {
		targetDir := cwd
		if p, ok := args["path"].(string); ok && strings.TrimSpace(p) != "" {
			targetDir = p
		}
		trackID, _ := args["track_id"].(string)
		if strings.TrimSpace(trackID) == "" {
			return mcp.NewErrorResult("track_id is required"), nil
		}
		title, _ := args["name"].(string)
		trackType, _ := args["type"].(string)

		err := track.CreateTrack(targetDir, trackID, title, trackType, nil)
		if err != nil {
			return mcp.NewErrorResult(fmt.Sprintf("Failed to create track '%s': %v", trackID, err)), nil
		}

		return mcp.NewTextResult(fmt.Sprintf("Successfully created track '%s' in .cooper/active/%s/", trackID, trackID), false), nil
	})

	// Tool 4: cooper_track_status
	srv.RegisterTool(mcp.Tool{
		Name:        "cooper_track_status",
		Description: "Inspect and query status and task progress of active Cooper tracks",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"path": map[string]interface{}{
					"type":        "string",
					"description": "Target project root directory (defaults to current working directory)",
				},
				"track_id": map[string]interface{}{
					"type":        "string",
					"description": "Optional specific track ID to inspect",
				},
			},
		},
	}, func(ctx context.Context, args map[string]interface{}) (mcp.CallToolResult, error) {
		targetDir := cwd
		if p, ok := args["path"].(string); ok && strings.TrimSpace(p) != "" {
			targetDir = p
		}

		if tid, ok := args["track_id"].(string); ok && strings.TrimSpace(tid) != "" {
			info, err := track.GetTrackStatus(targetDir, tid)
			if err != nil {
				return mcp.NewErrorResult(err.Error()), nil
			}
			text := fmt.Sprintf("Track `%s` (%s)\nStatus: %s\nTasks: %d/%d completed",
				info.TrackID, info.Title, info.Status, info.CompletedTasks, info.TotalTasks)
			return mcp.NewTextResult(text, false), nil
		}

		// List all active tracks
		activeDir := filepath.Join(targetDir, ".cooper", "active")
		entries, err := os.ReadDir(activeDir)
		if err != nil {
			return mcp.NewTextResult("No active tracks directory found at "+activeDir, false), nil
		}

		var sb strings.Builder
		sb.WriteString("# Active Tracks Status\n\n")
		count := 0
		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}
			info, err := track.GetTrackStatus(targetDir, entry.Name())
			if err != nil {
				sb.WriteString(fmt.Sprintf("- **%s**: (metadata unavailable)\n", entry.Name()))
				continue
			}
			count++
			sb.WriteString(fmt.Sprintf("- **%s** (`%s`): %s [%d/%d tasks completed]\n",
				info.Title, info.TrackID, info.Status, info.CompletedTasks, info.TotalTasks))
		}

		if count == 0 {
			sb.WriteString("No active tracks currently in progress.\n")
		}

		return mcp.NewTextResult(sb.String(), false), nil
	})

	// Tool 5: cooper_validate
	srv.RegisterTool(mcp.Tool{
		Name:        "cooper_validate",
		Description: "Validate SDD Living Specs, Spec Deltas, and tracks registry parity for GIVEN/WHEN/THEN syntax compliance",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"path": map[string]interface{}{
					"type":        "string",
					"description": "Target project root directory (defaults to current working directory)",
				},
			},
		},
	}, func(ctx context.Context, args map[string]interface{}) (mcp.CallToolResult, error) {
		targetDir := cwd
		if p, ok := args["path"].(string); ok && strings.TrimSpace(p) != "" {
			targetDir = p
		}

		errs, err := validator.ValidateRepository(targetDir)
		if err != nil {
			return mcp.NewErrorResult(fmt.Sprintf("Validation error: %v", err)), nil
		}

		if len(errs) == 0 {
			return mcp.NewTextResult(fmt.Sprintf("✅ Validation Passed: Repository at %s conforms to Cooper SDD specifications.", targetDir), false), nil
		}

		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("❌ Validation Failed: %d issue(s) detected at %s:\n\n", len(errs), targetDir))
		for _, e := range errs {
			sb.WriteString(fmt.Sprintf("- [%s] %s (Line %d): %s\n", e.Rule, e.File, e.Line, e.Message))
		}
		return mcp.NewTextResult(sb.String(), true), nil
	})

	// Tool 6: cooper_self_update
	srv.RegisterTool(mcp.Tool{
		Name:        "cooper_self_update",
		Description: "Check for or apply in-place binary upgrades for Cooper CLI from GitHub Releases",
		InputSchema: mcp.ToolInputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"check_only": map[string]interface{}{
					"type":        "boolean",
					"description": "Check if an update is available without applying it (default: true)",
				},
				"force": map[string]interface{}{
					"type":        "boolean",
					"description": "Force re-download even if already on latest version",
				},
				"target_version": map[string]interface{}{
					"type":        "string",
					"description": "Specific semantic version tag (e.g. v1.2.0)",
				},
			},
		},
	}, func(ctx context.Context, args map[string]interface{}) (mcp.CallToolResult, error) {
		checkOnly := true
		if co, ok := args["check_only"].(bool); ok {
			checkOnly = co
		}
		force, _ := args["force"].(bool)
		targetVersion, _ := args["target_version"].(string)

		opts := updater.Options{
			Repo:           "twoBoots/cooper",
			BinaryName:     "cooper",
			TargetVersion:  targetVersion,
			CurrentVersion: version,
			Force:          force,
			CheckOnly:      checkOnly,
			Client:         mcpUpdaterClientOverride,
		}

		res, err := updater.SelfUpdate(opts)
		if err != nil {
			return mcp.NewErrorResult(fmt.Sprintf("Update failed: %v", err)), nil
		}

		return mcp.NewTextResult(res.Message, false), nil
	})

	// Register Standard Cooper Resources
	srv.RegisterResource(mcp.Resource{
		URI:         "cooper://index",
		Name:        "Cooper Index",
		Description: "Single source of truth project handshake index (.cooper/index.md)",
		MIMEType:    "text/markdown",
	}, func(ctx context.Context, uri string) (mcp.ReadResourceResult, error) {
		indexPath := filepath.Join(cwd, ".cooper", "index.md")
		data, err := os.ReadFile(indexPath)
		if err != nil {
			return mcp.ReadResourceResult{}, fmt.Errorf("failed to read .cooper/index.md: %w", err)
		}
		return mcp.ReadResourceResult{
			Contents: []mcp.ResourceContent{
				{
					URI:      uri,
					MIMEType: "text/markdown",
					Text:     string(data),
				},
			},
		}, nil
	})

	srv.RegisterResource(mcp.Resource{
		URI:         "cooper://tracks",
		Name:        "Cooper Tracks Registry",
		Description: "Tracks registry index (.cooper/tracks.md)",
		MIMEType:    "text/markdown",
	}, func(ctx context.Context, uri string) (mcp.ReadResourceResult, error) {
		tracksPath := filepath.Join(cwd, ".cooper", "tracks.md")
		data, err := os.ReadFile(tracksPath)
		if err != nil {
			return mcp.ReadResourceResult{}, fmt.Errorf("failed to read .cooper/tracks.md: %w", err)
		}
		return mcp.ReadResourceResult{
			Contents: []mcp.ResourceContent{
				{
					URI:      uri,
					MIMEType: "text/markdown",
					Text:     string(data),
				},
			},
		}, nil
	})

	return srv
}

// RunMCPServer starts the Cooper MCP server over the provided reader and writer.
func RunMCPServer(in io.Reader, out io.Writer, version, commit, date, cwd string) error {
	srv := NewCooperServer(version, commit, date, cwd)
	return srv.Serve(context.Background(), in, out)
}
