# Proposal: Bump Cooper CLI Patch Version to v1.2.1

## Problem Statement
Since the `v1.2.0` CLI release, several repository updates have shipped on `main`, including:
1. Updating the markdown link auditor (`link/troop-attribution` in `internal/validator/link_auditor.go`) to accept `https://twoboots.github.io/troop` alongside `https://github.com/twoBoots/troop`.
2. Scaffolding the `.claude/skills` and `CLAUDE.md` bridge in `install.sh` and `cooper-setup`.
3. Enforcing the **Living Files Stand Alone** policy across Cooper workflow definitions, templates, specs, and skills.
4. Publishing the VitePress documentation site and updated installation guides.

Because `cmd/version.go` remained at `1.2.0`, existing CLI users running `cooper update` or `cooper update --check` are not notified of a new release tag and do not receive the updated `link/troop-attribution` validator rule without `--force`.

## Proposed Solution
Bump the default CLI `Version` constant in `cmd/version.go` from `1.2.0` to `1.2.1` and add a unit test verifying the default semantic version constant so that merging to `main` triggers the automated `v1.2.1` release tag and binary build in `.github/workflows/release.yml`.

## Acceptance Criteria
1. `cmd/version.go` defines `Version = "1.2.1"`.
2. `cmd/version_test.go` verifies the default `Version` constant is `"1.2.1"`.
3. `go test ./...` and `go run . validate` pass with >80% coverage.
