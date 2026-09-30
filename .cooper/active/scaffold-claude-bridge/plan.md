# Implementation Plan: Scaffold Claude Bridge (`.claude/skills` and `CLAUDE.md`)

## Phase 1: Test Suite & Failure Demonstration (Red)
- [~] Task: Claude Bridge Installer Tests
  - [ ] Sub-task: Write unit test in `cmd/installer_test.go` verifying greenfield creation of `.claude/skills` symlink and `CLAUDE.md` bridge (Red)
  - [ ] Sub-task: Write unit test in `cmd/installer_test.go` verifying preservation of existing `.claude/skills` entity (Red)
  - [ ] Sub-task: Write unit test in `cmd/installer_test.go` verifying append and idempotent non-duplication of `@AGENTS.md` in existing `CLAUDE.md` (Red)
  - [ ] Sub-task: Execute `go test ./cmd/ -run TestInstallScript_ClaudeBridge` and confirm test failures (Red)
- [ ] Task: Phase 1 Verification & Checkpoint

## Phase 2: Installer Implementation (Green)
- [ ] Task: Implement Claude Bridge in `install.sh`
  - [ ] Sub-task: Implement `setup_claude_bridge` function in `install.sh` handling `.claude/skills` and `CLAUDE.md`
  - [ ] Sub-task: Invoke `setup_claude_bridge` during installer execution after AGENTS.md setup
  - [ ] Sub-task: Re-run `go test ./cmd/ -run TestInstallScript_ClaudeBridge` and verify all tests pass (Green)
  - [ ] Sub-task: Run all existing installer tests in `cmd/installer_test.go` to ensure no regressions
- [ ] Task: Phase 2 Verification & Checkpoint

## Phase 3: Setup Skill Documentation & Source Parity
- [ ] Task: Document Claude Bridge in Cooper Setup Skill
  - [ ] Sub-task: Update `skills/cooper-setup/SKILL.md` to document `.claude/skills` symlink and `CLAUDE.md` bridge setup
  - [ ] Sub-task: Synchronize `.agents/skills/cooper-setup/SKILL.md` byte-for-byte with `skills/cooper-setup/SKILL.md`
  - [ ] Sub-task: Run `cooper validate` to verify skill identity, spec structure, and links
- [ ] Task: Phase 3 Verification & Checkpoint

## Phase 4: End-to-End Verification & Remote Sync
- [ ] Task: Complete Test Suite & Manual Dry-Run
  - [ ] Sub-task: Run full test suite (`go test ./...`) across all packages
  - [ ] Sub-task: Perform end-to-end dry run of `install.sh` in an isolated temporary directory
  - [ ] Sub-task: Present manual verification results for user confirmation
- [ ] Task: Phase 4 Verification & Checkpoint
