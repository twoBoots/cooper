package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/twoBoots/bender/pkg/mcp"
	cooperMCP "github.com/twoBoots/cooper/internal/mcp"
)

var (
	mcpTransport             string
	mcpInstallClients        string
	mcpInstallAll            bool
	mcpInstallNonInteractive bool
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

	// Attach 'cooper mcp install' subcommand
	mcpCmd.AddCommand(newMCPInstallCmd())

	return mcpCmd
}

func newMCPInstallCmd() *cobra.Command {
	installCmd := &cobra.Command{
		Use:     "install",
		Aliases: []string{"setup", "configure"},
		Short:   "Configure Cooper MCP server in AI coding assistants (Cursor, Antigravity, Claude, Windsurf, VS Code)",
		Long: `Automatically detect and configure Cooper's MCP server in supported AI assistant configuration files.

Supported clients:
  * cursor          - Cursor IDE (.cursor/mcp.json)
  * antigravity     - Google Antigravity / agy (~/.gemini/config/mcp_config.json)
  * claude-desktop  - Anthropic Claude Desktop
  * claude-code     - Anthropic Claude Code (~/.claude.json)
  * windsurf        - Windsurf IDE (~/.codeium/windsurf/mcp_config.json)
  * vscode          - VS Code (.vscode/mcp.json)`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var clients []string
			if mcpInstallClients != "" {
				for _, c := range strings.Split(mcpInstallClients, ",") {
					if trimmed := strings.TrimSpace(c); trimmed != "" {
						clients = append(clients, trimmed)
					}
				}
			}
			cwd, _ := os.Getwd()
			homeDir, _ := os.UserHomeDir()
			return RunMCPInstall(cmd.OutOrStdout(), cwd, homeDir, clients, mcpInstallAll, mcpInstallNonInteractive)
		},
	}

	installCmd.Flags().StringVarP(&mcpInstallClients, "client", "c", "", "Comma-separated list of clients to configure (cursor, antigravity, claude-desktop, claude-code, windsurf, vscode)")
	installCmd.Flags().BoolVarP(&mcpInstallAll, "all", "a", false, "Configure all supported AI clients")
	installCmd.Flags().BoolVarP(&mcpInstallNonInteractive, "non-interactive", "y", false, "Run non-interactively configuring detected clients")

	return installCmd
}

// RunMCPInstall configures client target configuration files.
func RunMCPInstall(out io.Writer, cwd string, homeDir string, clientIDs []string, all bool, nonInteractive bool) error {
	supported := mcp.GetSupportedClients(cwd, homeDir)
	var selectedIDs []string

	if all {
		for _, s := range supported {
			selectedIDs = append(selectedIDs, s.ID)
		}
	} else if len(clientIDs) > 0 {
		selectedIDs = clientIDs
	} else {
		// Non-interactive or default: install into detected clients or all if none detected
		for _, s := range supported {
			if s.Detected {
				selectedIDs = append(selectedIDs, s.ID)
			}
		}
		if len(selectedIDs) == 0 {
			for _, s := range supported {
				selectedIDs = append(selectedIDs, s.ID)
			}
		}
	}

	if len(selectedIDs) == 0 {
		fmt.Fprintln(out, "ℹ️ No AI clients selected. MCP configuration unchanged.")
		return nil
	}

	fmt.Fprintln(out, "🛢️ Configuring Cooper MCP Server in AI coding assistants...")

	opts := mcp.InstallerOptions{
		ServerName: "cooper",
		Command:    "cooper",
		Args:       []string{"mcp"},
		Cwd:        cwd,
		HomeDir:    homeDir,
		ClientIDs:  selectedIDs,
	}

	results, err := mcp.InstallClients(opts)
	if err != nil {
		return err
	}

	for _, res := range results {
		if res.Error != nil {
			fmt.Fprintf(out, "  [✗] Failed %s (%s): %v\n", res.DisplayName, res.ConfigPath, res.Error)
		} else if res.Created {
			fmt.Fprintf(out, "  [✓] Configured %s -> Created %s\n", res.DisplayName, res.ConfigPath)
		} else if res.Updated {
			fmt.Fprintf(out, "  [✓] Configured %s -> Updated %s\n", res.DisplayName, res.ConfigPath)
		}
	}

	fmt.Fprintln(out, "\n✨ Cooper MCP configuration completed!")
	return nil
}
