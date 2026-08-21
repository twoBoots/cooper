package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/twoBoots/cooper/internal/validator"
)

type ValidationReport struct {
	TotalErrors int                         `json:"total_errors"`
	Errors      []validator.ValidationError `json:"errors"`
}

func newValidateCmd() *cobra.Command {
	var dir string
	var jsonOutput bool

	validateCmd := &cobra.Command{
		Use:     "validate",
		Aliases: []string{"lint"},
		Short:   "Validate SDD capability specs, spec deltas, track metadata, and repository links",
		Long:    `Validate performs deterministic linting of living capability specs, active track spec deltas, track metadata JSON schemas, tracks.md registry parity, and repository markdown link integrity.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if dir == "" {
				var err error
				dir, err = os.Getwd()
				if err != nil {
					return fmt.Errorf("failed to get current working directory: %w", err)
				}
			}

			errs, err := validator.ValidateRepository(dir)
			if err != nil {
				return fmt.Errorf("validation execution failed: %w", err)
			}

			report := ValidationReport{
				TotalErrors: len(errs),
				Errors:      errs,
			}

			if jsonOutput {
				enc := json.NewEncoder(cmd.OutOrStdout())
				enc.SetIndent("", "  ")
				if err := enc.Encode(report); err != nil {
					return err
				}
			} else {
				if len(errs) == 0 {
					fmt.Fprintln(cmd.OutOrStdout(), "✓ Validation passed successfully. All specs, metadata, and links are valid.")
				} else {
					fmt.Fprintf(cmd.ErrOrStderr(), "✗ Found %d validation issue(s):\n\n", len(errs))
					for _, e := range errs {
						fmt.Fprintf(cmd.ErrOrStderr(), "  [%s] %s:%d - %s\n", e.Rule, e.File, e.Line, e.Message)
					}
				}
			}

			if len(errs) > 0 {
				return fmt.Errorf("validation failed with %d error(s)", len(errs))
			}

			return nil
		},
	}

	validateCmd.Flags().StringVarP(&dir, "dir", "d", "", "Directory path of the Cooper project to validate (defaults to current directory)")
	validateCmd.Flags().BoolVar(&jsonOutput, "json", false, "Output validation report in JSON format")

	return validateCmd
}
