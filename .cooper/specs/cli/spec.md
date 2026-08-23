# Capability Specification: Cooper Go CLI & MCP Server Engine

## Purpose & Scope
Defines the architectural behavior, commands, living spec validation, self-updating, and stdio Model Context Protocol (MCP) server engine for the Cooper CLI (`cooper`).

## Requirements

### Requirement: Binary Self-Update Command
The Cooper CLI SHALL provide an `update` command (and alias `self-update`) to perform in-place binary upgrades from GitHub Releases.

#### Scenario: Check Update Availability
- GIVEN a user running `cooper update --check`
- WHEN a newer semantic release is published on `twoBoots/cooper`
- THEN the CLI outputs the latest available version and notifies the user to run `cooper update`.

#### Scenario: Apply In-Place Update
- GIVEN a user running `cooper update`
- WHEN a newer version is available
- THEN the CLI downloads the platform asset matching current OS and architecture, replaces the binary in-place, applies macOS quarantine stripping and code signing if on Darwin, and reports success.

#### Scenario: Already Up To Date
- GIVEN a user running `cooper update`
- WHEN the local binary is equal to or newer than the latest release
- THEN the CLI notifies the user that the binary is already up to date unless `--force` is specified.

### Requirement: Embedded Stdio Model Context Protocol (MCP) Server
The Cooper CLI SHALL provide an `mcp` command that starts a JSON-RPC 2.0 MCP server over stdio.

#### Scenario: MCP Server Initialization
- GIVEN an AI assistant starting `cooper mcp`
- WHEN the client sends an `initialize` JSON-RPC request
- THEN the server responds with server name `cooper-mcp`, server version, and tool capabilities.

#### Scenario: Expose Cooper SDD Tools
- GIVEN an initialized Cooper MCP server
- WHEN the client sends a `tools/list` request
- THEN the server returns `cooper_get_version`, `cooper_init_project`, `cooper_track_create`, `cooper_track_status`, `cooper_validate`, and `cooper_self_update`.

### Requirement: AI Assistant MCP Client Configuration Installer
The Cooper CLI SHALL provide an `mcp install` command to automatically register Cooper into editor MCP configuration files.

#### Scenario: Auto-Detect and Configure Clients
- GIVEN supported AI assistant configuration files (Cursor, Antigravity, Claude Desktop, Claude Code, Windsurf, VS Code)
- WHEN running `cooper mcp install`
- THEN the CLI safely merges `cooper` (`command: "cooper"`, `args: ["mcp"]`) into the detected client config files.

### Requirement: SDD Repository Syntax Validation
The Cooper CLI SHALL provide a `validate` command to lint living capability specifications, active track spec deltas, and markdown links across the workspace.

#### Scenario: Validate Compliant Workspace
- GIVEN a repository adhering to GIVEN/WHEN/THEN living spec formatting and metadata schemas
- WHEN running `cooper validate`
- THEN the validator exits cleanly with code 0 and reports success.

### Requirement: Automated Multi-Platform Release CI/CD Pipeline
The Cooper repository SHALL provide automated Continuous Integration (CI) and Release pipelines to compile, tag, and publish cross-platform binaries to GitHub Releases on push to `main` and on tag push.

#### Scenario: Continuous Integration Validation
- GIVEN a push or pull request targeting the `main` branch
- WHEN GitHub Actions CI executes
- THEN it verifies code formatting (`gofmt`), runs linters (`go vet`), and executes all unit tests with code coverage reporting.

#### Scenario: Automated Semantic Tagging on Main
- GIVEN a commit merged to the `main` branch
- WHEN the version defined in `cmd/version.go` is not yet tagged on origin
- THEN the release workflow creates an annotated Git tag `v<version>` and pushes it to origin.

#### Scenario: Cross-Platform Binary Compilation & Publishing
- GIVEN a release build triggered by a `main` push or semantic tag
- WHEN the release workflow compiles the binary matrix
- THEN it builds static binaries for Linux (`amd64`, `arm64`), macOS (`amd64`, `arm64`), and Windows (`amd64`) with embedded version, commit hash, and build date metadata via `-ldflags`, and publishes the assets to both the semantic release tag and `latest` release on GitHub.

