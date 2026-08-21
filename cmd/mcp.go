package cmd

import (
	"os"

	"github.com/spf13/cobra"
	cooperMCP "github.com/twoBoots/cooper/internal/mcp"
)

var (
	mcpTransport string
)

func newMCPCmd() *cobra.Command {
	mcpCmd := &cobra.Command{
		Use:     "mcp",
		Aliases: []string{"serve"},
		Short:   "Start the Model Context Protocol (MCP) server over stdio",
		Long: `Start the Cooper Model Context Protocol (MCP) server over stdio.

Exposes SDD Living Spec tools, project initialization, track management,
repository linting, and self-updating to AI coding assistants (Antigravity,
Claude Code, Cursor, Windsurf).`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cwd, err := os.Getwd()
			if err != nil {
				cwd = "."
			}
			return cooperMCP.RunMCPServer(cmd.InOrStdin(), cmd.OutOrStdout(), Version, Commit, Date, cwd)
		},
	}

	mcpCmd.Flags().StringVarP(&mcpTransport, "transport", "t", "stdio", "Transport protocol to use (stdio)")

	return mcpCmd
}
