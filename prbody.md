## Summary & Intent

This PR implements the foundational **Cooper Go CLI & SDD Validator (`cooper`)**, fulfilling track **`track-cooper-cli-core`** decomposed from approved RFC [`rfc-cooper-cli-mcp`](.cooper/active/rfc-cooper-cli-mcp/rfc.md).

It establishes the standalone Go binary runtime with embedded assets, high-speed deterministic specification linting, greenfield/brownfield project scaffolding, and Troop worktree lifecycle management.

---

## Spec Delta Summary

### Added Requirements (`+`)
- **`+ Requirement: Core CLI Entrypoint & Versioning`**: Implemented `cooper` Cobra command tree with global flags (`--verbose`, `--non-interactive`) and `cooper version` with text/JSON outputs and ldflags injection points.
- **`+ Requirement: Project Initialization & Migration`**: Implemented `cooper init` providing greenfield `.cooper/` and `.agents/skills/` scaffolding and automatic migration of legacy `.conductor/` and `openspec/` tracks.
- **`+ Requirement: Deterministic Spec & Markdown Validation`**: Implemented `cooper validate` / `cooper lint` checking document headers, `## Requirements`, normative keywords (`SHALL`/`MUST`), scenario syntax (`GIVEN`/`WHEN`/`THEN`), `metadata.json` schemas, `tracks.md` registry parity, and repository markdown link integrity.
- **`+ Requirement: Track & Worktree Orchestration`**: Implemented `cooper track [new|status|checkpoint|close]` wrapping Troop worktree creation, checkpoint Git Notes recording, and track finalization.

### Removed Requirements (`-`)
- *None*

---

## Completed Phases & Checkpoints

- **Phase 1: Go Module, Embedded Assets & CLI Entrypoint**
  - Commit: `5b04689` (`feat(cli): Initialize Go module, Cobra root command, and version subcommand`)
  - Commit: `dac9e32` (`feat(scaffold): Embed template assets and agent skills via go:embed`)
  - Checkpoint: `7cb6a16` (`cooper(checkpoint): Checkpoint end of Phase 1 - Go Module, Embedded Assets & CLI Entrypoint`)
- **Phase 2: SDD Validator Engine (`cooper validate`)**
  - Commit: `dabef58` (`feat(validator): Implement living spec and spec-delta syntax linter`)
  - Commit: `b0382e1` (`feat(validator): Implement track metadata JSON and tracks.md registry parity validator`)
  - Commit: `5d952d1` (`feat(validator): Implement outbound link auditor and cooper validate/lint CLI command`)
  - Checkpoint: `57613e5` (`cooper(checkpoint): Checkpoint end of Phase 2 - SDD Validator Engine`)
- **Phase 3: Project Scaffolding & Track Orchestrator (`cooper init` & `cooper track`)**
  - Commit: `2ca6772` (`feat(scaffold): Implement project initialization and brownfield migration engine (cooper init)`)
  - Commit: `73a9433` (`feat(track): Implement track lifecycle and Troop worktree orchestrator (cooper track)`)
  - Checkpoint: `980d743` (`cooper(checkpoint): Checkpoint end of Phase 3 - Project Scaffolding & Track Orchestrator`)

---

## Verification & Quality Gates

- **`go vet`**: 0 errors / 0 warnings.
- **Test Suite**: 34 unit & integration tests passing (100% pass rate).
- **Test Coverage Breakdown**:
  - `cmd`: **81.6%**
  - `internal/scaffold`: **90.7%**
  - `internal/track`: **94.5%**
  - `internal/validator`: **90.8%**
  - **Overall Project Coverage**: **>89%** (enforcing >80% threshold).
- **Git Notes**: Verification notes attached to every checkpoint commit and synchronized upstream to `origin/refs/notes/commits`.
