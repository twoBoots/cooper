# Proposal: Cooper Bender CLI Integration

## Intent
Integrate the generic Go CLI archetype, self-updater, and Model Context Protocol (MCP) server engine from [`github.com/twoBoots/bender`](https://github.com/twoBoots/bender) into the `cooper` CLI.

## Background & Context
Cooper provides Spec-Driven Development (SDD) governance, living capability specs, and Troop worktree isolation. In RFC [`rfc-cooper-cli-mcp`](.cooper/active/rfc-cooper-cli-mcp/rfc.md), Cooper specified a Go CLI binary with self-updating capabilities and an embedded stdio MCP server.

`twoBoots/bender` has extracted and hardened these generic Go CLI capabilities into clean, reusable packages (`pkg/updater` and `pkg/mcp`). Integrating `bender` allows `cooper` to eliminate duplicated logic and gain:
- Zero-dependency GitHub Release binary self-updating with platform asset resolution, macOS quarantine stripping, and ad-hoc codesigning (`pkg/updater`).
- Lightweight stdio JSON-RPC 2.0 MCP server with tool, resource, and prompt registration (`pkg/mcp`).
- Multi-client MCP configuration auto-installer supporting Cursor, Google Antigravity, Claude Desktop, Claude Code, Windsurf, and VS Code (`pkg/mcp.InstallClients`).

## Scope Boundaries

### In Scope
1. **Dependency Integration**: Add `github.com/twoBoots/bender` to `go.mod`.
2. **Binary Self-Update Command (`cooper update`)**:
   - Wire `cooper update` and `cooper self-update` to `updater.SelfUpdate` with flags `--check`, `--force`, `--target-version`, and `--repo` (defaulting to `twoBoots/cooper`).
3. **Embedded Stdio MCP Server (`cooper mcp`)**:
   - Implement `cooper mcp` command serving JSON-RPC 2.0 over `stdio`.
   - Expose core Cooper SDD tools:
     - `cooper_get_version`: Return CLI version, commit hash, build date.
     - `cooper_init_project`: Initialize `.cooper/` in target directory.
     - `cooper_track_create`: Scaffold new track proposal and metadata.
     - `cooper_track_status`: Query active track and worktree status.
     - `cooper_validate`: Lint SDD GIVEN/WHEN/THEN syntax and structure.
     - `cooper_self_update`: Check or trigger CLI self-update via MCP.
4. **Client MCP Auto-Installer (`cooper mcp install`)**:
   - Automatically configure `cooper` MCP server into detected AI coding assistant configurations (`.cursor/mcp.json`, Antigravity `mcp_config.json`, `claude_desktop_config.json`, `~/.claude.json`, Windsurf, VS Code).
5. **Quality & Test Coverage**:
   - Comprehensive unit and integration test suite with >80% test coverage following strict TDD.

### Out of Scope
- 3-way project template/skill content synchronization engine (`cooper update --templates` or manifest 3-way diffs; planned for subsequent track).
- Changes to git alias hooks or core Troop worktree scripts.
