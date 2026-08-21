# Design Document: Cooper Bender CLI Integration

## 1. Overview & Architecture

This design details the integration of `github.com/twoBoots/bender` packages into the `cooper` Go CLI codebase:
1. **`pkg/updater`**: Leveraged for in-place binary self-updating from GitHub Releases (`twoBoots/cooper`).
2. **`pkg/mcp`**: Leveraged for running a stdio JSON-RPC 2.0 MCP server and automating MCP client configuration for AI coding editors.

```
┌─────────────────────────────────────────────────────────────┐
│                         cooper CLI                          │
│                                                             │
│   ├── cmd/                                                  │
│   │   ├── init.go         (cooper init)                     │
│   │   ├── track.go        (cooper track [new|status|...])   │
│   │   ├── validate.go     (cooper validate)                 │
│   │   ├── update.go       (cooper update / self-update)     │
│   │   └── mcp.go          (cooper mcp [install])            │
│   │                                                         │
│   ├── internal/                                             │
│   │   ├── mcp/            (Cooper MCP tool registrations)   │
│   │   ├── scaffold/       (Project initialization logic)    │
│   │   ├── track/          (Track creation & management)     │
│   │   └── validator/      (GIVEN/WHEN/THEN SDD linter)      │
│   │                                                         │
│   └── (Imported Packages from github.com/twoBoots/bender)   │
│       ├── pkg/updater     (Self-updater engine)             │
│       └── pkg/mcp         (MCP stdio server & installer)    │
└─────────────────────────────────────────────────────────────┘
```

---

## 2. Command Architecture & Interfaces

### 2.1 Self-Updater Command (`cmd/update.go`)
- **Command**: `cooper update` (Aliases: `self-update`)
- **Flags**:
  - `-c, --check`: Query latest GitHub release and report availability without modifying the local binary.
  - `-f, --force`: Re-download and reinstall current/target version even if already up to date.
  - `-t, --target-version`: Specific semantic version tag to install (e.g. `v1.2.0`).
  - `--repo`: Target GitHub repo (hidden flag, default: `twoBoots/cooper`).
  - `--exec-path`: Target executable path to overwrite (hidden flag, default: current binary).
- **Execution Flow**:
  1. Instantiate `updater.Options` with `Repo: "twoBoots/cooper"`, `BinaryName: "cooper"`, `CurrentVersion: Version`.
  2. Call `updater.SelfUpdate(opts)`.
  3. Output status message and emojis to stdout/stderr.

### 2.2 MCP Server Command (`cmd/mcp.go`)
- **Command**: `cooper mcp` (Aliases: `serve`)
- **Flags**:
  - `-t, --transport`: Protocol transport (default: `stdio`).
- **Subcommand**: `cooper mcp install` (Aliases: `setup`, `configure`)
  - Flags:
    - `-c, --client`: Comma-separated list of client IDs (`cursor,antigravity,claude-desktop,claude-code,windsurf,vscode`).
    - `-a, --all`: Install across all supported clients.
    - `-y, --non-interactive`: Install into detected client config paths without prompting.
  - Calls `mcp.InstallClients` with `ServerName: "cooper"`, `Command: "cooper"`, `Args: ["mcp"]`.

---

## 3. Embedded MCP Server & Tool Registry (`internal/mcp/`)

The Cooper MCP server (`internal/mcp/server.go`) initializes `mcp.NewServer("cooper-mcp", Version, cwd)` and registers the following MCP tools:

### Tool 1: `cooper_get_version`
- **Description**: Returns the compiled version, Git commit hash, and build timestamp of the Cooper binary.
- **Input Schema**: `{}` (no parameters)
- **Output**: Formatted string with version details.

### Tool 2: `cooper_init_project`
- **Description**: Scaffolds standard `.cooper/` directory structure and base template files in the target project.
- **Input Schema**:
  - `path` (string, optional): Target repository root directory (defaults to server working directory).
- **Output**: JSON or summary text of created/verified files.

### Tool 3: `cooper_track_create`
- **Description**: Creates a new Cooper SDD track with `metadata.json`, `proposal.md`, and `spec-deltas/`.
- **Input Schema**:
  - `track_id` (string, required): Kebab-case track identifier.
  - `name` (string, required): Human-readable track title.
  - `type` (string, optional): Track type (`feature`, `bugfix`, `rfc`, `chore`, default: `feature`).
- **Output**: Summary of scaffolded track files.

### Tool 4: `cooper_track_status`
- **Description**: Inspects `.cooper/active/` and `.cooper/tracks.md` to report status of all active tracks.
- **Input Schema**: `{}` (no parameters)
- **Output**: Markdown list of active tracks and their metadata.

### Tool 5: `cooper_validate`
- **Description**: Validates living capability specs and active track spec-deltas for GIVEN/WHEN/THEN format compliance.
- **Input Schema**:
  - `path` (string, optional): Specific spec file or directory to validate (defaults to whole repository).
- **Output**: Linting results with error/warning counts.

### Tool 6: `cooper_self_update`
- **Description**: Checks for or applies binary self-updates for Cooper CLI via MCP.
- **Input Schema**:
  - `check_only` (boolean, optional): If true, checks update availability without applying (default: true).
  - `target_version` (string, optional): Specific release tag to install.
- **Output**: Result message indicating current and latest versions.

---

## 4. Testing Strategy

1. **Unit Testing (`cmd/update_test.go`)**:
   - Mock GitHub API using `httptest.Server`.
   - Test `cooper update --check`, `cooper update --force`, and error handling for failed network/corrupted releases.
2. **Unit Testing (`cmd/mcp_test.go`)**:
   - Test MCP CLI flag parsing and `RunMCPInstall` execution with temporary configuration directories.
3. **Integration Testing (`internal/mcp/server_test.go`)**:
   - Send JSON-RPC stdio payloads (`initialize`, `tools/list`, `tools/call`) over memory buffers (`bytes.Buffer`).
   - Validate each registered tool execution (`cooper_get_version`, `cooper_validate`, `cooper_track_create`, etc.).
4. **Coverage Mandate**: Maintain >80% code coverage across all newly added and modified packages.
