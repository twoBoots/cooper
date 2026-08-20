# Track Proposal: Core Go CLI & SDD Validator

- **Track ID**: `track-cooper-cli-core`
- **Type**: Feature
- **Status**: Planning

---

## 1. Summary & Motivation
This track implements the foundational **Go CLI (`cooper`)** as decomposed from approved RFC [`rfc-cooper-cli-mcp`](.cooper/active/rfc-cooper-cli-mcp/rfc.md).

It establishes the Go module structure, embedded template assets, Cobra command-line interface, project initialization (`cooper init`), deterministic SDD validator (`cooper validate` / `cooper lint`), and high-level track orchestration (`cooper track [new|status|checkpoint|close]`).

---

## 2. Problem Statement
Currently, Cooper relies entirely on shell scripts and AI prompt instructions to manage tracks, validate living specifications, and scaffold projects. This lacks deterministic validation (e.g. malformed GIVEN/WHEN/THEN syntax or invalid metadata JSON passes undetected until runtime), and requires manual Git worktree and notes commands for every task.

---

## 3. Proposed Solution
1. **Go Module & CLI Framework**: Establish `github.com/twoBoots/cooper` Go module with Cobra CLI and embedded filesystems (`go:embed` for templates and skills).
2. **Deterministic SDD Validator (`cooper validate` / `cooper lint`)**: Fast binary validation of:
   - Living capability specs & spec deltas (`GIVEN`/`WHEN`/`THEN` syntax, `+`/`-` delta diff rules).
   - Track metadata JSON schema compliance.
   - Parity between `.cooper/tracks.md` and active `.worktrees/`.
   - Outbound repository link integrity.
3. **Project Initialization (`cooper init`)**: Brownfield Conductor/OpenSpec migration and greenfield project scaffolding.
4. **Track Lifecycle Wrapper (`cooper track ...`)**: Single CLI commands to spawn Troop worktrees, create track metadata, record phase checkpoints with Git Notes, and close tracks.
5. **Strict TDD Coverage**: Unit test suite with >80% coverage.

---

## 4. Scope Boundaries
- **In Scope**:
  - `go.mod` (Go 1.23) and dependencies (`github.com/spf13/cobra`).
  - `cmd/root.go`, `cmd/version.go`, `cmd/init.go`, `cmd/validate.go`, `cmd/track.go`, and `main.go`.
  - `internal/validator/` (Spec linter, metadata validator, link auditor).
  - `internal/scaffold/` (Project initialization, embedded templates).
  - `internal/track/` (Worktree management, phase checkpoints, metadata updates).
  - Comprehensive unit test suite.
- **Out of Scope (Addressed in Subsequent Tracks)**:
  - CLI binary self-updating and 3-way diff updater (Track 2: `track-cooper-updater-diff3`).
  - Embedded stdio MCP server (Track 3: `track-cooper-embedded-mcp`).
  - GitHub Actions CI/release matrix and installer updates (Track 4: `track-cooper-installer-packaging`).
