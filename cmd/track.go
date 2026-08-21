package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/twoBoots/cooper/internal/track"
)

func newTrackCmd() *cobra.Command {
	trackCmd := &cobra.Command{
		Use:   "track",
		Short: "Manage Cooper SDD tracks and Troop worktrees",
		Long:  `Create, inspect, record checkpoints for, and close Cooper tracks.`,
	}

	trackCmd.AddCommand(newTrackNewCmd())
	trackCmd.AddCommand(newTrackStatusCmd())
	trackCmd.AddCommand(newTrackCheckpointCmd())
	trackCmd.AddCommand(newTrackCloseCmd())

	return trackCmd
}

func newTrackNewCmd() *cobra.Command {
	var trackType string
	var title string
	var dir string

	cmd := &cobra.Command{
		Use:   "new <track_id>",
		Short: "Spawn a new track worktree and initialize track artifacts",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			trackID := args[0]
			if dir == "" {
				var err error
				dir, err = os.Getwd()
				if err != nil {
					return err
				}
			}

			exec := &track.DefaultGitExecutor{}
			if err := track.CreateTrack(dir, trackID, title, trackType, exec); err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "✓ Created track '%s' in .cooper/active/%s and spawned worktree .worktrees/%s\n", trackID, trackID, trackID)
			return nil
		},
	}

	cmd.Flags().StringVarP(&trackType, "type", "t", "feature", "Track type (feature, bugfix, chore, rfc)")
	cmd.Flags().StringVar(&title, "title", "", "Track title")
	cmd.Flags().StringVarP(&dir, "dir", "d", "", "Project root directory")

	return cmd
}

func newTrackStatusCmd() *cobra.Command {
	var dir string
	var jsonOutput bool

	cmd := &cobra.Command{
		Use:   "status <track_id>",
		Short: "Display progress and task counts for an active track",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			trackID := args[0]
			if dir == "" {
				var err error
				dir, err = os.Getwd()
				if err != nil {
					return err
				}
			}

			info, err := track.GetTrackStatus(dir, trackID)
			if err != nil {
				return err
			}

			if jsonOutput {
				enc := json.NewEncoder(cmd.OutOrStdout())
				enc.SetIndent("", "  ")
				return enc.Encode(info)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Track: %s\n", info.TrackID)
			fmt.Fprintf(cmd.OutOrStdout(), "Title: %s\n", info.Title)
			fmt.Fprintf(cmd.OutOrStdout(), "Type: %s | Status: %s\n", info.Type, info.Status)
			fmt.Fprintf(cmd.OutOrStdout(), "Tasks: %d/%d completed\n", info.CompletedTasks, info.TotalTasks)
			return nil
		},
	}

	cmd.Flags().StringVarP(&dir, "dir", "d", "", "Project root directory")
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "Output in JSON format")

	return cmd
}

func newTrackCheckpointCmd() *cobra.Command {
	var dir string
	var title string

	cmd := &cobra.Command{
		Use:   "checkpoint <phase_number>",
		Short: "Record a phase checkpoint commit with verification Git Note",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			phaseNum, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid phase number '%s': %w", args[0], err)
			}
			if dir == "" {
				dir, err = os.Getwd()
				if err != nil {
					return err
				}
			}

			exec := &track.DefaultGitExecutor{}
			sha, err := track.RecordCheckpoint(dir, phaseNum, title, exec)
			if err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "✓ Phase %d checkpoint recorded: %s\n", phaseNum, sha)
			return nil
		},
	}

	cmd.Flags().StringVar(&title, "title", "Verification", "Phase title")
	cmd.Flags().StringVarP(&dir, "dir", "d", "", "Project root directory")

	return cmd
}

func newTrackCloseCmd() *cobra.Command {
	var dir string

	cmd := &cobra.Command{
		Use:   "close <track_id>",
		Short: "Close a completed track, update tracks.md, and remove worktree",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			trackID := args[0]
			if dir == "" {
				var err error
				dir, err = os.Getwd()
				if err != nil {
					return err
				}
			}

			exec := &track.DefaultGitExecutor{}
			if err := track.CloseTrack(dir, trackID, exec); err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "✓ Closed track '%s'\n", trackID)
			return nil
		},
	}

	cmd.Flags().StringVarP(&dir, "dir", "d", "", "Project root directory")

	return cmd
}
