# Implementation Plan: Slim the Cooper CLI to a Validate-First Surface

- **Track ID**: `slim-cli-validate-first`
- **Coverage Target**: >80% total statements
- **Ordering Constraint**: Phase 2 Task 2.1 (port newer skill content) MUST complete before Task 2.2 (delete embedded tree), or newer mandates are lost.

---

## Phase 1: Remove the MCP Server & Track Orchestration

- [ ] **Task 1.1: Remove MCP Server & Client Installer**
  - [ ] Sub-task: Write `cmd/root_test.go` asserting `cooper --help` lists no `mcp` command and that `NewRootCmd()` registers no subcommand named `mcp` (Red)
  - [ ] Sub-task: Delete `internal/mcp/server.go`, `internal/mcp/server_test.go`, `cmd/mcp.go`, `cmd/mcp_test.go`
  - [ ] Sub-task: Remove the `newMCPCmd()` registration from `cmd/root.go` (Green)
  - [ ] Sub-task: Run `go mod tidy`; confirm `github.com/twoBoots/bender` is retained for `pkg/updater` and only `pkg/mcp` usage is dropped
  - [ ] Sub-task: Verify `go build ./...` and `go vet ./...` pass (Refactor)

- [ ] **Task 1.2: Remove Track Orchestration & the Fabricating Checkpoint**
  - [ ] Sub-task: Extend `cmd/root_test.go` asserting no `track` command is registered (Red)
  - [ ] Sub-task: Write a guard test asserting no source file in the repository contains the literal string `Automated Tests: PASSED` (Red — enforces the new spec prohibition and prevents reintroduction)
  - [ ] Sub-task: Delete `internal/track/`, `cmd/track.go`, `cmd/track_test.go`
  - [ ] Sub-task: Remove the `newTrackCmd()` registration from `cmd/root.go` (Green)
  - [ ] Sub-task: Confirm both guard tests pass (Refactor)

- [ ] **Task 1.3: Phase 1 Verification & Checkpoint**
  - [ ] Sub-task: Run `git fetch origin main` and merge any workflow/living-spec updates
  - [ ] Sub-task: Run `CGO_ENABLED=0 go test ./...` and record the result
  - [ ] Sub-task: Present a manual verification guide and obtain explicit user approval via interactive question tool
  - [ ] Sub-task: Create checkpoint commit and attach a Git Note reporting the **actual** test outcome and the **actual** recorded approval
  - [ ] Sub-task: `git push origin slim-cli-validate-first`

---

## Phase 2: Consolidate to a Single Scaffolder

- [ ] **Task 2.1: Port Newer Skill Content Into the Surviving Tree** *(must precede 2.2)*
  - [ ] Sub-task: Write a test asserting each `skills/cooper-*/SKILL.md` contains both `Interactive Question Protocol` and `Native File Tools Mandate` (Red — currently fails for all six)
  - [ ] Sub-task: For each of the six skills, port the newer content from `internal/scaffold/assets/skills/` into `skills/`, reconciling any root-only edits rather than overwriting wholesale
  - [ ] Sub-task: Diff each pair to confirm the surviving file is a superset of both inputs (Green)
  - [ ] Sub-task: Confirm the mandate test passes for all six (Refactor)

- [ ] **Task 2.2: Delete the Second Scaffolder & Duplicate Asset Tree**
  - [ ] Sub-task: Extend `cmd/root_test.go` asserting no `init` command is registered (Red)
  - [ ] Sub-task: Write a test asserting no duplicate `SKILL.md` tree exists outside `skills/` (Red)
  - [ ] Sub-task: Delete `internal/scaffold/` (including `assets/`), `cmd/init.go`, `cmd/init_test.go`
  - [ ] Sub-task: Remove the `newInitCmd()` registration from `cmd/root.go` (Green)
  - [ ] Sub-task: Confirm `templates/` and `AGENTS.template.md` remain at repository root as installer sources (Refactor)

- [ ] **Task 2.3: Optional Non-Fatal Binary Installation in `install.sh`**
  - [ ] Sub-task: Write a shell test scaffolding a throwaway git repo with `PATH` stripped of `go` and network egress blocked, asserting `install.sh` exits 0 and produces a complete `.cooper/` workspace (Red)
  - [ ] Sub-task: Implement bin-dir resolution (`/usr/local/bin` if writable, else `${HOME}/.local/bin`)
  - [ ] Sub-task: Implement Tier 1 local `go build`, Tier 2 release-asset download, Darwin quarantine strip and ad-hoc codesign
  - [ ] Sub-task: Wrap the whole step so every failure prints a notice and continues with exit 0 (Green)
  - [ ] Sub-task: Add a test asserting a fresh `install.sh` run yields `.cooper/COOPER.md`, `.cooper/TROOP.md`, the three Troop aliases, and a `.worktrees/` entry in `.gitignore` (Refactor)

- [ ] **Task 2.4: Phase 2 Verification & Checkpoint**
  - [ ] Sub-task: Run `git fetch origin main` and reconcile
  - [ ] Sub-task: Run full test suite and the `install.sh` scaffolding tests; record results
  - [ ] Sub-task: Manually scaffold a throwaway repo and confirm every `AGENTS.md` path reference resolves
  - [ ] Sub-task: Obtain explicit user verification approval via interactive question tool
  - [ ] Sub-task: Checkpoint commit with a truthful verification Git Note; `git push origin slim-cli-validate-first`

---

## Phase 3: Arm the Validation Gates

- [ ] **Task 3.1: Measure Post-Removal Coverage Before Arming the Gate**
  - [ ] Sub-task: Run `go test -coverprofile` across the reduced surface and record the true total
  - [ ] Sub-task: If total is below 80%, add `internal/validator` tests until it clears — do not weaken the threshold
  - [ ] Sub-task: Record the measured figure in this plan for the checkpoint note

- [ ] **Task 3.2: Backticked Repository Path Auditing**
  - [ ] Sub-task: Write table-driven tests in `internal/validator/link_auditor_test.go` covering: a dangling backticked `.md` path (must flag), an existing backticked path (must pass), and must-not-flag cases — shell snippets, glob patterns, flags such as `--force`, URLs, and bare words with no `/` (Red)
  - [ ] Sub-task: Implement `link/code-path-exists` in `link_auditor.go` under the §2.3 conservative rules (Green)
  - [ ] Sub-task: Run `cooper validate` against the whole repository and resolve any genuine dangling references it now surfaces (Refactor)

- [ ] **Task 3.3: Wire CI Enforcement**
  - [ ] Sub-task: Add the `Validate Cooper SDD Specs` step (`go run . validate`) to `.github/workflows/ci.yml`
  - [ ] Sub-task: Replace the decorative coverage print with the 80% enforcing gate
  - [ ] Sub-task: Verify locally that both steps fail on a deliberately broken spec and on a synthetic sub-threshold profile, then pass on a clean tree
  - [ ] Sub-task: Confirm neither gate is written into scaffolded consumer projects

- [ ] **Task 3.4: Documentation & Version Bump**
  - [ ] Sub-task: Rewrite the `README.md` CLI section to `validate`, `update`, `version`; remove all `cooper mcp`, `cooper init`, and `cooper track` examples
  - [ ] Sub-task: Add a removal note covering stale `cooper mcp` entries in editor MCP configs and how to remove them
  - [ ] Sub-task: Update the README structure diagram and Value Matrix to reflect the single scaffolder
  - [ ] Sub-task: Bump `cmd/version.go` `Version` to `1.2.0`
  - [ ] Sub-task: Run `cooper validate` to confirm documentation links and specs are clean

- [ ] **Task 3.5: Phase 3 Verification & Checkpoint**
  - [ ] Sub-task: Run `git fetch origin main` and reconcile
  - [ ] Sub-task: Run full suite with coverage; confirm the gate passes on measured numbers
  - [ ] Sub-task: Obtain explicit user verification approval via interactive question tool
  - [ ] Sub-task: Checkpoint commit with a truthful verification Git Note; `git push origin slim-cli-validate-first`
  - [ ] Sub-task: Hand off to `cooper-review`
