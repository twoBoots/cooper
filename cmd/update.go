package cmd

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"
	"github.com/twoBoots/bender/pkg/updater"
)

const (
	// DefaultCooperRepo is the default GitHub repository for Cooper releases.
	DefaultCooperRepo = "twoBoots/cooper"
)

var (
	updaterClientOverride *updater.Client

	updateCheckFlag         bool
	updateForceFlag         bool
	updateTargetVersionFlag string
	updateRepoFlag          string
	updateExecPathFlag      string
)

// SetUpdaterClient overrides the default updater client (used primarily in tests).
func SetUpdaterClient(c *updater.Client) {
	updaterClientOverride = c
}

// newUpdateCmd creates the 'cooper update' command.
func newUpdateCmd() *cobra.Command {
	updateCmd := &cobra.Command{
		Use:     "update",
		Aliases: []string{"self-update"},
		Short:   "Update Cooper CLI to the latest (or specified) version",
		Long: `Checks GitHub Releases for new versions of Cooper, downloads the
appropriate platform binary, and updates the local installation in place.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			opts := updater.Options{
				Repo:           updateRepoFlag,
				BinaryName:     "cooper",
				TargetVersion:  updateTargetVersionFlag,
				CurrentVersion: Version,
				ExecutablePath: updateExecPathFlag,
				Force:          updateForceFlag,
				CheckOnly:      updateCheckFlag,
				Client:         updaterClientOverride,
			}
			return RunUpdate(cmd.OutOrStdout(), cmd.ErrOrStderr(), opts, updaterClientOverride)
		},
	}

	updateCmd.Flags().BoolVarP(&updateCheckFlag, "check", "c", false, "Check for available updates without applying")
	updateCmd.Flags().BoolVarP(&updateForceFlag, "force", "f", false, "Force re-download and reinstall even if up to date")
	updateCmd.Flags().StringVarP(&updateTargetVersionFlag, "target-version", "t", "", "Target a specific release version tag (e.g. v1.3.0)")
	updateCmd.Flags().StringVar(&updateRepoFlag, "repo", DefaultCooperRepo, "GitHub repository to check for releases")
	updateCmd.Flags().StringVar(&updateExecPathFlag, "exec-path", "", "Target executable path to overwrite (advanced)")
	_ = updateCmd.Flags().MarkHidden("exec-path")
	_ = updateCmd.Flags().MarkHidden("repo")

	return updateCmd
}

// RunUpdate performs the update operation and writes user-friendly logs to the provided output writers.
func RunUpdate(out, errOut io.Writer, opts updater.Options, client *updater.Client) error {
	if client == nil {
		client = updater.NewClient()
	}
	opts.Client = client
	if opts.CurrentVersion == "" {
		opts.CurrentVersion = Version
	}

	fmt.Fprintf(out, "🛢️ Checking for updates (current version: v%s)...\n", opts.CurrentVersion)

	res, err := updater.SelfUpdate(opts)
	if err != nil {
		return fmt.Errorf("update failed: %w", err)
	}

	if opts.CheckOnly {
		if res.UpdateAvailable {
			fmt.Fprintf(out, "✨ %s\nRun 'cooper update' to install the new version.\n", res.Message)
		} else {
			fmt.Fprintf(out, "✅ %s\n", res.Message)
		}
		return nil
	}

	if res.Updated {
		fmt.Fprintf(out, "✅ %s\n", res.Message)
	} else {
		fmt.Fprintf(out, "ℹ️ %s\n", res.Message)
	}

	return nil
}
