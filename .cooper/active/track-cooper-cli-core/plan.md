# Implementation Plan: Core Go CLI & SDD Validator

- **Track ID**: `track-cooper-cli-core`
- **TDD Enforcement**: Minimum 80% test coverage (`go test -v -coverprofile=coverage.out ./...`)

---

## Phase 1: Go Module, Embedded Assets & CLI Entrypoint

- [x] Task: Go Module & Root CLI Scaffold (5b04689)
  - [x] Sub-task: Initialize `go.mod` (`github.com/twoBoots/cooper`) and add `github.com/spf13/cobra`
  - [x] Sub-task: Implement `cmd/root.go` and `main.go` with global flags
  - [x] Sub-task: Write unit tests for root command and version injection (`cmd/version.go`)
- [x] Task: Embedded Assets & Templates Engine (dac9e32)
  - [x] Sub-task: Implement `internal/scaffold/embedded.go` with `go:embed` for templates and skills
  - [x] Sub-task: Write unit tests verifying embedded assets integrity
- [x] Task: Phase 1 Verification & Checkpoint (7cb6a16)
  - [x] Sub-task: Run unit tests and verify clean compilation
  - [x] Sub-task: Phase 1 Checkpoint & remote sync (`git push origin track-cooper-cli-core`) [checkpoint: 7cb6a16]

---

## Phase 2: SDD Validator Engine (`internal/validator` & `cooper validate`)

- [x] Task: Spec Delta & Living Spec Syntax Linter (dabef58)
  - [x] Sub-task: Write unit tests for GIVEN/WHEN/THEN and `+`/`-` delta parsing (Red)
  - [x] Sub-task: Implement `internal/validator/spec_linter.go` (Green)
  - [x] Sub-task: Refactor and verify coverage >80% (Refactor)
- [x] Task: Track Metadata & Tracks Registry Validator (b0382e1)
  - [x] Sub-task: Write unit tests for `metadata.json` schema and `.cooper/tracks.md` parity (Red)
  - [x] Sub-task: Implement `internal/validator/metadata.go` (Green)
  - [x] Sub-task: Refactor and verify coverage >80% (Refactor)
- [ ] Task: Outbound Link Auditor & CLI Validate Command
  - [ ] Sub-task: Implement `internal/validator/link_auditor.go` for external tool links
  - [ ] Sub-task: Implement `cmd/validate.go` exposing `cooper validate` and `cooper lint`
  - [ ] Sub-task: Write unit and integration tests for validation CLI
- [ ] Task: Phase 2 Verification & Checkpoint
  - [ ] Sub-task: Run full test suite with coverage report
  - [ ] Sub-task: Phase 2 Checkpoint & remote sync (`git push origin track-cooper-cli-core`)

---

## Phase 3: Project Scaffolding & Track Orchestrator

- [ ] Task: Project Scaffolding Engine (`cooper init`)
  - [ ] Sub-task: Write unit tests for greenfield and brownfield migration (Red)
  - [ ] Sub-task: Implement `internal/scaffold/init.go` and `cmd/init.go` (Green)
  - [ ] Sub-task: Refactor and verify coverage >80% (Refactor)
- [ ] Task: Track Lifecycle & Troop Worktree Orchestrator (`cooper track`)
  - [ ] Sub-task: Write unit tests for track creation, checkpoint recording, and close (Red)
  - [ ] Sub-task: Implement `internal/track/` and `cmd/track.go` (Green)
  - [ ] Sub-task: Refactor and verify coverage >80% (Refactor)
- [ ] Task: Full Test Suite Quality Gate & Final Checkpoint
  - [ ] Sub-task: Run `go test -v -race -coverprofile=coverage.out ./...` and assert >80% coverage
  - [ ] Sub-task: Phase 3 Checkpoint & remote sync (`git push origin track-cooper-cli-core`)
