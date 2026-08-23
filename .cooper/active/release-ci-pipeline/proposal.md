# Proposal: Automated Multi-Platform Release CI/CD Pipeline & v1.0.0 Baseline

## Intent
Establish an automated Continuous Integration (CI) and multi-platform binary Release pipeline for Cooper using GitHub Actions, adopt the standard Bender release archetype, and establish `v1.0.0` as the baseline production release.

## Background & Context
The Cooper Go CLI provides core Spec-Driven Development (SDD) governance, living capability specs, and embedded MCP tooling. The CLI includes an in-place self-updater (`cooper update` / `cooper self-update`) powered by `github.com/twoBoots/bender/pkg/updater`.

However, the Cooper repository currently lacks GitHub Actions CI workflows (`.github/workflows/ci.yml`) and Release automation (`.github/workflows/release.yml`). Consequently:
- Pull requests and trunk commits are not validated automatically with `gofmt`, `go vet`, and unit test coverage.
- Binary assets matching the platform naming matrix (`cooper-linux-x86_64`, `cooper-linux-aarch64`, `cooper-darwin-x86_64`, `cooper-darwin-aarch64`, `cooper-windows-x86_64.exe`) are not published to GitHub Releases.
- `cooper update` and remote installation scripts cannot resolve released binary artifacts.

Porting and tailoring the hardened release pipeline from `bender` (the source of truth for CLI archetypes in twoBoots) will close this missing link, fulfill living capability spec requirements, and enable seamless binary updates for all users.

## Scope Boundaries

### In Scope
1. **GitHub Actions CI Workflow (`.github/workflows/ci.yml`)**:
   - Automated triggers on push to `main`, PRs to `main`, and workflow call.
   - Node 20 deprecation suppression flags (`FORCE_JAVASCRIPT_ACTIONS_TO_NODE20: true`, `NODE_NO_WARNINGS: "1"`).
   - Go formatting check (`gofmt -l .`), static analysis (`go vet ./...`), and test execution with coverage reporting.
2. **GitHub Actions Release Workflow (`.github/workflows/release.yml`)**:
   - Triggers on push to `main` and semantic tags (`v*`).
   - Automated semantic tagging job (`auto-tag`) extracting version from `cmd/version.go` and creating `v<version>` tag if not yet present on remote.
   - Multi-arch binary build matrix for Linux (x86_64, aarch64), macOS (x86_64, aarch64), and Windows (x86_64.exe).
   - Build metadata injection via `-ldflags` for `Version`, `Commit`, and `Date` into `github.com/twoBoots/cooper/cmd`.
   - Dual release publishing: publishes asset artifacts to both semantic tag releases (`v1.0.0`) and the `latest` rolling release via `gh release`.
3. **Set Baseline Production Version (`v1.0.0`)**:
   - Update `Version` in `cmd/version.go` from `"dev"` to `"1.0.0"`.
   - Update test suites to verify version output consistency.
4. **Living Spec Delta**:
   - Update `cli/spec.md` living spec to specify the automated release matrix and publishing lifecycle.

### Out of Scope
- Linux ARM 32-bit or non-standard architectures not supported by Bender updater platform conventions.
- Homebrew formula packaging (can be addressed in future tracks).
