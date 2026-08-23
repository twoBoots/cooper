# Implementation Plan: Automated Multi-Platform Release CI/CD Pipeline & v1.0.0 Baseline

- **Track ID:** `release-ci-pipeline`
- **Worktree:** `.worktrees/release-ci-pipeline`
- **Spec Delta:** [`.cooper/active/release-ci-pipeline/spec-deltas/cli/spec.md`](spec-deltas/cli/spec.md)
- **Status:** Planning

---

## Phase 1: Continuous Integration Workflow (`.github/workflows/ci.yml`)

- [x] **Task 1.1: Implement Reusable CI Workflow** (`b49b343`)
  - [x] Sub-task: Create `.github/workflows/ci.yml` with `push`, `pull_request`, and `workflow_call` triggers
  - [x] Sub-task: Configure Node 20 deprecation suppression env vars and Go 1.27.0 setup
  - [x] Sub-task: Add steps for `gofmt` check, `go vet ./...`, and `go test -v -coverprofile=coverage.out ./...`

- [x] **Task 1.2: Phase 1 Verification & Checkpoint** (`4441ee8`) [checkpoint: 4441ee8]
  - [x] Sub-task: Run local formatting check and `go test -v -cover ./...`
  - [x] Sub-task: Execute phase sync (`git fetch origin main`) and record checkpoint

---

## Phase 2: Multi-Platform Release Workflow (`.github/workflows/release.yml`)

- [x] **Task 2.1: Implement Automated Tagging & Cross-Platform Release Matrix** (`d560eff`)
  - [x] Sub-task: Create `.github/workflows/release.yml` with CI dependency and `auto-tag` job extracting `Version` from `cmd/version.go`
  - [x] Sub-task: Configure `build-and-release` matrix for Linux (`amd64`, `arm64`), macOS (`amd64`, `arm64`), and Windows (`amd64`) with `-ldflags`
  - [x] Sub-task: Implement `publish-release` job uploading assets to semantic version tag and `latest` rolling release

- [x] **Task 2.2: Phase 2 Verification & Checkpoint** (`03c1650`) [checkpoint: 03c1650]
  - [x] Sub-task: Validate YAML syntax and job dependencies
  - [x] Sub-task: Execute phase sync (`git fetch origin main`) and record checkpoint

---

## Phase 3: Version Baseline (`v1.0.0`) & Cross-Compilation Dry Run

- [x] **Task 3.1: Bump Version to `1.0.0` in `cmd/version.go`** (`7760c14`)
  - [x] Sub-task: Update `Version` in `cmd/version.go` to `"1.0.0"`
  - [x] Sub-task: Update/verify unit tests in `cmd/version_test.go` and `cmd/update_test.go` (Red -> Green -> Refactor)
  - [x] Sub-task: Verify local multi-target cross-compilation with `CGO_ENABLED=0 GOOS=... GOARCH=... go build`

- [x] **Task 3.2: Phase 3 Verification & Checkpoint** (`f865371`) [checkpoint: f865371]
  - [x] Sub-task: Run `go test -v -race -cover ./...`
  - [x] Sub-task: Execute phase sync (`git fetch origin main`) and record checkpoint

---

## Phase 4: Integration Verification, Spec Promotion & Track Finalization

- [x] **Task 4.1: Living Spec Validation & Promotion** (`eef85e0`)
  - [x] Sub-task: Run `cooper validate` across all specs and spec deltas
  - [x] Sub-task: Promote spec delta into `.cooper/specs/cli/spec.md`

- [x] **Task 4.2: Documentation & Track Registry Update** (`cf2b869`)
  - [x] Sub-task: Update `README.md` with CI badge and release binary information
  - [x] Sub-task: Update `.cooper/tracks.md` to reflect completed track

- [ ] **Task 4.3: Final Phase Verification & Checkpoint**
  - [ ] Sub-task: Run full test suite with coverage report (`go test -v -coverprofile=coverage.out ./...`)
  - [ ] Sub-task: Perform Principal Engineer code review (`cooper-review`) and prepare PR
