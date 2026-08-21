# Spec Delta: Cooper Bender CLI Integration

## Requirements

### Requirement: Binary Self-Update Command
+ The Cooper CLI SHALL provide an `update` command (and alias `self-update`) to perform in-place binary upgrades from GitHub Releases.
+
+ #### Scenario: Check Update Availability
+ - GIVEN a user running `cooper update --check`
+ - WHEN a newer semantic release is published on `twoBoots/cooper`
+ - THEN the CLI outputs the latest available version and notifies the user to run `cooper update`.
+
+ #### Scenario: Apply In-Place Update
+ - GIVEN a user running `cooper update`
+ - WHEN a newer version is available
+ - THEN the CLI downloads the platform asset matching current OS and architecture, replaces the binary in-place, applies macOS quarantine stripping and code signing if on Darwin, and reports success.
+
+ #### Scenario: Already Up To Date
+ - GIVEN a user running `cooper update`
+ - WHEN the local binary is equal to or newer than the latest release
+ - THEN the CLI notifies the user that the binary is already up to date unless `--force` is specified.

### Requirement: Embedded Stdio Model Context Protocol (MCP) Server
+ The Cooper CLI SHALL provide an `mcp` command that starts a JSON-RPC 2.0 MCP server over stdio.
+
+ #### Scenario: MCP Server Initialization
+ - GIVEN an AI assistant starting `cooper mcp`
+ - WHEN the client sends an `initialize` JSON-RPC request
+ - THEN the server responds with server name `cooper-mcp`, server version, and tool capabilities.
+
+ #### Scenario: Expose Cooper SDD Tools
+ - GIVEN an initialized Cooper MCP server
+ - WHEN the client sends a `tools/list` request
+ - THEN the server returns `cooper_get_version`, `cooper_init_project`, `cooper_track_create`, `cooper_track_status`, `cooper_validate`, and `cooper_self_update`.

### Requirement: AI Assistant MCP Client Configuration Installer
+ The Cooper CLI SHALL provide an `mcp install` command to automatically register Cooper into editor MCP configuration files.
+
+ #### Scenario: Auto-Detect and Configure Clients
+ - GIVEN supported AI assistant configuration files (Cursor, Antigravity, Claude Desktop, Claude Code, Windsurf, VS Code)
+ - WHEN running `cooper mcp install`
+ - THEN the CLI safely merges `cooper` (`command: "cooper"`, `args: ["mcp"]`) into the detected client config files.
