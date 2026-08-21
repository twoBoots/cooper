package cmd

import (
	"github.com/spf13/cobra"
)

var (
	verbose        bool
	nonInteractive bool
)

// NewRootCmd initializes the root cobra command and attaches subcommands.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "cooper",
		Short: "cooper - Spec-Driven Development (SDD) CLI",
		Long:  `cooper is the CLI for the Cooper Spec-Driven Development framework and Troop worktree management.`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output logging")
	rootCmd.PersistentFlags().BoolVar(&nonInteractive, "non-interactive", false, "Disable interactive prompts for automation/CI")

	// Register subcommands
	rootCmd.AddCommand(newVersionCmd())
	rootCmd.AddCommand(newValidateCmd())
	rootCmd.AddCommand(newInitCmd())
	rootCmd.AddCommand(newTrackCmd())
	rootCmd.AddCommand(newUpdateCmd())

	return rootCmd
}

// Execute runs the root command.
func Execute() error {
	rootCmd := NewRootCmd()
	return rootCmd.Execute()
}
