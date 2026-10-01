# Implementation Plan: Scaffold Claude Bridge (`.claude/skills` and `CLAUDE.md`)

## Phase 1: Test Suite & Failure Demonstration (Red)
- [x] Task: Claude Bridge Installer Tests (d262e7d)
  - [x] Sub-task: Write unit test in `cmd/installer_test.go` verifying greenfield creation of `.claude/skills` symlink and `CLAUDE.md` bridge (Red)
  - [x] Sub-task: Write unit test in `cmd/installer_test.go` verifying preservation of existing `.claude/skills` entity (Red)
  - [x] Sub-task: Write unit test in `cmd/installer_test.go` verifying append and idempotent non-duplication of `@AGENTS.md` in existing `CLAUDE.md` (Red)
  - [x] Sub-task: Execute `go test ./cmd/ -run TestInstallScript_ClaudeBridge` and confirm test failures (Red)
- [x] Task: Phase 1 Verification & Checkpoint [checkpoint: 91a4a83]

## Phase 2: Installer Implementation (Green)
- [x] Task: Implement Claude Bridge in `install.sh` (6ffeb91)
  - [x] Sub-task: Implement `setup_claude_bridge` function in `install.sh` handling `.claude/skills` and `CLAUDE.md`
  - [x] Sub-task: Invoke `setup_claude_bridge` during installer execution after AGENTS.md setup
  - [x] Sub-task: Re-run `go test ./cmd/ -run TestInstallScript_ClaudeBridge` and verify all tests pass (Green)
  - [x] Sub-task: Run all existing installer tests in `cmd/installer_test.go` to ensure no regressions
- [x] Task: Phase 2 Verification & Checkpoint [checkpoint: 6fc7152]

## Phase 3: Setup Skill Documentation & Source Parity
- [x] Task: Document Claude Bridge in Cooper Setup Skill (be5b1a6)
  - [x] Sub-task: Update `skills/cooper-setup/SKILL.md` to document `.claude/skills` symlink and `CLAUDE.md` bridge setup
  - [x] Sub-task: Synchronize `.agents/skills/cooper-setup/SKILL.md` byte-for-byte with `skills/cooper-setup/SKILL.md`
  - [x] Sub-task: Run `cooper validate` to verify skill identity, spec structure, and links
- [x] Task: Phase 3 Verification & Checkpoint [checkpoint: 4adf093]

## Phase 4: End-to-End Verification & Remote Sync
- [ ] Task: Complete Test Suite & Manual Dry-Run
  - [ ] Sub-task: Run full test suite (`go test ./...`) across all packages
  - [ ] Sub-task: Perform end-to-end dry run of `install.sh` in an isolated temporary directory
  - [ ] Sub-task: Present manual verification results for user confirmation
- [ ] Task: Phase 4 Verification & Checkpoint
