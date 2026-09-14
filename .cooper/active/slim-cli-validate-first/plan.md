# Implementation Plan: Slim the Cooper CLI to a Validate-First Surface

- **Track ID**: `slim-cli-validate-first`
- **Coverage Target**: >80% total statements
- **Ordering Constraint**: Phase 2 Task 2.1 (port newer skill content) MUST complete before Task 2.2 (delete embedded tree), or newer mandates are lost.

---

## Phase 1: Remove the MCP Server & Track Orchestration

- [x] **Task 1.1: Remove MCP Server & Client Installer** (`901e029`)
  - [x] Sub-task: Write `cmd/root_test.go` asserting `cooper --help` lists no `mcp` command and that `NewRootCmd()` registers no subcommand named `mcp` (Red)
  - [x] Sub-task: Delete `internal/mcp/server.go`, `internal/mcp/server_test.go`, `cmd/mcp.go`, `cmd/mcp_test.go`
  - [x] Sub-task: Remove the `newMCPCmd()` registration from `cmd/root.go` (Green)
  - [x] Sub-task: Run `go mod tidy`; confirm `github.com/twoBoots/bender` is retained for `pkg/updater` and only `pkg/mcp` usage is dropped
  - [x] Sub-task: Verify `go build ./...` and `go vet ./...` pass (Refactor)

- [x] **Task 1.2: Remove Track Orchestration & the Fabricating Checkpoint** (`6e335ae`)
  - [x] Sub-task: Extend `cmd/root_test.go` asserting no `track` command is registered (Red)
  - [x] Sub-task: Write a guard test asserting no source file in the repository contains the literal string `Automated Tests: PASSED` (Red — enforces the new spec prohibition and prevents reintroduction)
  - [x] Sub-task: Delete `internal/track/`, `cmd/track.go`, `cmd/track_test.go`
  - [x] Sub-task: Remove the `newTrackCmd()` registration from `cmd/root.go` (Green)
  - [x] Sub-task: Confirm both guard tests pass (Refactor)
  - Note: the `init` assertion was deferred into Task 2.2's own Red step so that each phase leaves the suite green.

- [x] **Task 1.3: Phase 1 Verification & Checkpoint** [checkpoint: `7eed176`]
  - [x] Sub-task: Run `git fetch origin main` and merge any workflow/living-spec updates — no upstream drift in `workflow.md` or `specs/`
  - [x] Sub-task: Run `CGO_ENABLED=0 go test ./...` and record the result — all packages pass, 93.1% coverage
  - [x] Sub-task: Present a manual verification guide and obtain explicit user approval via interactive question tool
  - [x] Sub-task: Create checkpoint commit and attach a Git Note reporting the **actual** test outcome and the **actual** recorded approval
  - [x] Sub-task: `git push origin slim-cli-validate-first`

---

## Phase 2: Consolidate to a Single Scaffolder

- [x] **Task 2.1: Port Newer Skill Content Into the Surviving Tree** (`db37b7b`)
  - [x] Sub-task: Write a test asserting each `skills/cooper-*/SKILL.md` contains both `Interactive Question Protocol` and `Native File Tools Mandate` (Red — failed for all six)
  - [x] Sub-task: For each of the six skills, port the newer content from `internal/scaffold/assets/skills/` into `skills/`, reconciling any root-only edits rather than overwriting wholesale
  - [x] Sub-task: Diff each pair to confirm the surviving file is a superset of both inputs (Green)
  - [x] Sub-task: Confirm the mandate test passes for all six (Refactor)
  - Finding: the embedded copy had regressed `cooper-rfc` heading `### 6.2` to `## 6.2`; root's correct heading was restored during the reconcile. A heading-hierarchy guard now prevents recurrence.
  - Finding: `.agents/skills/` is a legitimate third copy (Cooper's own installed instance) and was already in sync. Guarded by a new source/installed parity test.

- [x] **Task 2.2: Delete the Second Scaffolder & Duplicate Asset Tree** (`6e477ec`)
  - [x] Sub-task: Extend `cmd/root_test.go` asserting no `init` command is registered (Red)
  - [x] Sub-task: Write a test asserting no duplicate `SKILL.md` tree exists outside `skills/` (Red)
  - [x] Sub-task: Delete `internal/scaffold/` (including `assets/`), `cmd/init.go`, `cmd/init_test.go`
  - [x] Sub-task: Remove the `newInitCmd()` registration from `cmd/root.go` (Green)
  - [x] Sub-task: Confirm `templates/` and `AGENTS.template.md` remain at repository root as installer sources (Refactor)
  - Verified byte-identical before deletion: `templates/` and `AGENTS.template.md` vs their embedded copies.
  - Also removed: dead embedded `spec-template.md` / `spec-delta-template.md`, which the extraction filter excluded and never wrote anywhere.

- [x] **Task 2.3: Optional Non-Fatal Binary Installation in `install.sh`** (`53f7f2d`)
  - [x] Sub-task: Write a shell test asserting the binary step returns 0 under an empty `PATH` (no go, curl, wget) and that sourcing is side-effect free (Red)
  - [x] Sub-task: Implement bin-dir resolution (`/usr/local/bin` if writable, else `${HOME}/.local/bin`)
  - [x] Sub-task: Implement Tier 1 local `go build`, Tier 2 release-asset download, Darwin quarantine strip and ad-hoc codesign
  - [x] Sub-task: Wrap the whole step so every failure prints a notice and continues with exit 0 (Green)
  - [x] Sub-task: Verify a fresh `install.sh` run yields `.cooper/COOPER.md`, `.cooper/TROOP.md`, the three Troop aliases, and a `.worktrees/` entry in `.gitignore` (Refactor)
  - **Defect found and fixed:** `cooper validate` on a freshly scaffolded project reported `AGENTS.md:7` linking to `TROOP.md` after `install.sh` had relocated it to `.cooper/TROOP.md`. Every project scaffolded before this commit carried that dangling reference. Fixed by `relocate_troop_reference` and covered by a test.
  - **Deviation from plan:** the end-to-end scaffold assertion is verified manually, not automated. A full `install.sh` run fetches the Troop installer over the network, so a CI test would be flaky. The four hermetic installer tests cover the non-fatal contract; the e2e evidence is recorded in the task Git Note.

- [x] **Task 2.4: Remove the Fabricated Attestation Template From Instructions** (`5adc0b5`, guard fix `5269ed5`)
  - [x] Sub-task: Write a test asserting no `.md` file under `skills/`, `.agents/skills/`, or `.cooper/definition/` contains the literal `Automated Tests: PASSED` or `Manual Verification: APPROVED by user` (Red — failed for `skills/cooper-implement/SKILL.md` and its installed copy)
  - [x] Sub-task: Replace the hardcoded Git Note template in `skills/cooper-implement/SKILL.md` §3.4 with an instruction to record the actual test command, its real outcome, and the user's actual recorded response (Green)
  - [x] Sub-task: Resync `.agents/skills/` installed copy and confirm the parity guard passes (Refactor)
  - [x] Sub-task: Confirm the guard passes and `cooper validate` stays clean
  - Rationale: `RecordCheckpoint` was implementing this template faithfully. Deleting the Go function while leaving the instruction intact just reintroduces the defect by hand at the next checkpoint.
  - **Plan premise corrected:** `.cooper/definition/workflow.md` does **not** carry the hardcoded strings — line 128 already directs the agent to attach the actual verification report. No change was needed there, so `workflow.md` guiding principle 4 (isolated branch) did not apply and no split was required.

- [x] **Task 2.5: Phase 2 Verification & Checkpoint** [checkpoint: `6a510af`]
  - [x] Sub-task: Run `git fetch origin main` and reconcile — no upstream drift
  - [x] Sub-task: Run full test suite and the `install.sh` scaffolding tests; record results — all pass, 93.4% coverage
  - [x] Sub-task: Manually scaffold a throwaway repo and confirm every `AGENTS.md` path reference resolves — passed after the `relocate_troop_reference` fix
  - [x] Sub-task: Obtain explicit user verification approval via interactive question tool
  - [x] Sub-task: Checkpoint commit with a truthful verification Git Note; `git push origin slim-cli-validate-first`

---

## Phase 3: Arm the Validation Gates

- [x] **Task 3.1: Measure Post-Removal Coverage Before Arming the Gate**
  - [x] Sub-task: Run `go test -coverprofile` across the reduced surface and record the true total
  - [x] Sub-task: If total is below 80%, add `internal/validator` tests until it clears — not required
  - [x] Sub-task: Record the measured figure in this plan for the checkpoint note
  - **Measured: 93.4%** (cmd 95.2%, internal/validator 93.5%, main.go 0%). The design's projection held — removing the well-covered `track`/`mcp`/`scaffold` packages *raised* the total from 91.2%, because the deleted command surface carried more uncovered branches than the guard tests replacing it. The 80% gate arms with 13.4 points of headroom.

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
