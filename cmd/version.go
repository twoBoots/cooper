package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version is populated at build time via -ldflags
	Version = "dev"
	// Commit is populated at build time via -ldflags
	Commit = "none"
	// Date is populated at build time via -ldflags
	Date = "unknown"
)

type VersionInfo struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

func newVersionCmd() *cobra.Command {
	var jsonOutput bool

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version and build information for cooper",
		RunE: func(cmd *cobra.Command, args []string) error {
			info := VersionInfo{
				Version: Version,
				Commit:  Commit,
				Date:    Date,
			}

			if jsonOutput {
				enc := json.NewEncoder(cmd.OutOrStdout())
				enc.SetIndent("", "  ")
				return enc.Encode(info)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "cooper version %s (commit: %s, built at: %s)\n", info.Version, info.Commit, info.Date)
			return nil
		},
	}

	versionCmd.Flags().BoolVar(&jsonOutput, "json", false, "Output version information in JSON format")

	return versionCmd
}
