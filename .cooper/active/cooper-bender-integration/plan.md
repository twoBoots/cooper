# Implementation Plan: Cooper Bender CLI Integration

- **Track ID:** `cooper-bender-integration`
- **Worktree:** `.worktrees/cooper-bender-integration`
- **Spec Delta:** [`.cooper/active/cooper-bender-integration/spec-deltas/cooper-bender-integration/spec.md`](spec-deltas/cooper-bender-integration/spec.md)
- **Status:** Planning

---

## Phase 1: Module Dependency & Self-Update CLI (`cooper update`)

- [x] **Task 1.1: Dependency Scaffolding & `go.mod` Configuration** (`66bf75a`)
  - [x] Sub-task: Add `github.com/twoBoots/bender` dependency to `go.mod`
  - [x] Sub-task: Run `go mod tidy` and verify package accessibility

- [x] **Task 1.2: Implement `cooper update` Command** (`d7bc788`)
  - [x] Sub-task: Write unit tests in `cmd/update_test.go` covering `--check`, `--force`, `--target-version`, and error states with mocked GitHub API (Red)
  - [x] Sub-task: Implement `cmd/update.go` wiring `updater.SelfUpdate` with `twoBoots/cooper` defaults (Green)
  - [x] Sub-task: Register `UpdateCmd` in `cmd/root.go` and refactor for >80% test coverage (Refactor)

- [x] **Task 1.3: Phase 1 Verification & Checkpoint** (`aa8c725`) [checkpoint: aa8c725]
  - [x] Sub-task: Run `go test -v -cover ./cmd...`
  - [x] Sub-task: Execute phase sync (`git fetch origin main`) and record checkpoint

---

## Phase 2: Embedded MCP Server Engine & Tool Suite (`internal/mcp`)

- [x] **Task 2.1: Cooper MCP Server & SDD Tool Suite** (`dae2c1b`)
  - [x] Sub-task: Write unit & protocol tests in `internal/mcp/server_test.go` verifying `initialize`, `tools/list`, and tool executions (`cooper_get_version`, `cooper_init_project`, `cooper_track_create`, `cooper_track_status`, `cooper_validate`, `cooper_self_update`) (Red)
  - [x] Sub-task: Implement `internal/mcp/server.go` registering all Cooper SDD tools using `bender/pkg/mcp` (Green)
  - [x] Sub-task: Refactor tool schemas, error handling, and verify test coverage >80% (Refactor)

- [x] **Task 2.2: Stdio MCP CLI Command (`cooper mcp`)** (`b470ae2`)
  - [x] Sub-task: Write CLI unit tests in `cmd/mcp_test.go` for stdio server initiation (Red)
  - [x] Sub-task: Implement `cmd/mcp.go` connecting stdio transport to `internal/mcp.RunMCPServer` (Green)
  - [x] Sub-task: Register `MCPCmd` in `cmd/root.go` and refactor (Refactor)

- [x] **Task 2.3: Phase 2 Verification & Checkpoint** (`0b3d882`) [checkpoint: 0b3d882]
  - [x] Sub-task: Run `go test -v -cover ./internal/mcp/... ./cmd/...`
  - [x] Sub-task: Execute phase sync (`git fetch origin main`) and record checkpoint

---

## Phase 3: Multi-Client MCP Auto-Installer (`cooper mcp install`)

- [x] **Task 3.1: Implement `cooper mcp install` Command** (`9ce9c2e`)
  - [x] Sub-task: Write tests in `cmd/mcp_test.go` for client auto-detection and custom `--client` / `--all` installations across mock editor configs (Red)
  - [x] Sub-task: Implement `MCPInstallCmd` in `cmd/mcp.go` delegating to `mcp.InstallClients` for server `cooper` (Green)
  - [x] Sub-task: Refactor output formatting and verify coverage >80% (Refactor)

- [x] **Task 3.2: Phase 3 Verification & Checkpoint** (`ea04f0f`) [checkpoint: ea04f0f]
  - [x] Sub-task: Run full test suite `go test -v -race -cover ./...`
  - [x] Sub-task: Execute phase sync (`git fetch origin main`) and record checkpoint

---

## Phase 4: Integration Verification, Documentation & PR Preparation

- [ ] **Task 4.1: End-to-End Validation & Living Spec Linter**
  - [ ] Sub-task: Run `cooper validate` across all living capability specs and spec deltas
  - [ ] Sub-task: Verify static binary compilation `go build -o bin/cooper .`

- [ ] **Task 4.2: Documentation & Track Registry Update**
  - [ ] Sub-task: Update `README.md` and CLI documentation with `update` and `mcp` commands
  - [ ] Sub-task: Register completed track in `.cooper/tracks.md`

- [ ] **Task 4.3: Final Phase Verification & Checkpoint**
  - [ ] Sub-task: Run all tests with coverage report (`go test -v -coverprofile=coverage.out ./...`)
  - [ ] Sub-task: Perform code review (`cooper-review`) and prepare PR
