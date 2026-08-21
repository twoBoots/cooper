package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/twoBoots/cooper/internal/scaffold"
)

func newInitCmd() *cobra.Command {
	var dir string
	var force bool

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Scaffold a new Cooper SDD workspace or migrate legacy projects",
		Long:  `Init scaffolds the foundational .cooper directory, living specifications, workflow definitions, code styleguides, agent skills (.agents/skills), and AGENTS.md. Automatically detects and migrates legacy .conductor or openspec tracks.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if dir == "" {
				var err error
				dir, err = os.Getwd()
				if err != nil {
					return fmt.Errorf("failed to get current working directory: %w", err)
				}
			}

			if err := scaffold.InitProject(dir, force); err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "✓ Initialized Cooper SDD project in %s\n", dir)
			return nil
		},
	}

	initCmd.Flags().StringVarP(&dir, "dir", "d", "", "Target directory path to initialize (defaults to current directory)")
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "Force overwrite existing files")

	return initCmd
}
