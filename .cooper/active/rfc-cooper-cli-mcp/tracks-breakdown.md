# Implementation Track Breakdown: Cooper Go CLI & MCP Server

This initiative is decomposed into four focused, sequential execution tracks:

---

### Track 1: Core Go CLI & SDD Validator (`track-cooper-cli-core`)
- **Scope**:
  - Initialize Go module (`go.mod`) with Cobra CLI framework and embedded templates (`go:embed`).
  - Implement `cooper init` (project scaffolding & Conductor/OpenSpec migration).
  - Implement `cooper validate` / `cooper lint` (GIVEN/WHEN/THEN syntax validation, metadata schema checks, and link integrity verification).
  - Implement `cooper track [new|status|checkpoint|close]` wrapping Troop worktree commands.
- **Dependencies**: None
- **Living Specs Touched**: `.cooper/specs/cooper-cli/spec.md`

---

### Track 2: Upstream Reconciliation & 3-Way Diff Engine (`track-cooper-updater-diff3`)
- **Scope**:
  - Implement `.cooper/manifest.json` version fingerprinting.
  - Implement `internal/updater` fetching latest upstream releases from GitHub.
  - Implement 3-way diff and merge engine (`base` vs `local_custom` vs `upstream_release`).
  - Implement `cooper update` command with interactive conflict resolution.
- **Dependencies**: Track 1
- **Living Specs Touched**: `.cooper/specs/cooper-cli/spec.md`

---

### Track 3: Embedded Agent MCP Server (`track-cooper-embedded-mcp`)
- **Scope**:
  - Implement `internal/mcp` JSON-RPC server over stdio (`cooper mcp`).
  - Implement track tools: `cooper_track_create`, `cooper_checkpoint_record`, `cooper_spec_delta_validate`.
  - Implement agent-guided update tools: `cooper_check_updates`, `cooper_diff_updates`, `cooper_apply_update`.
  - Provide MCP configuration generator for Cursor, Claude Code, and Antigravity.
- **Dependencies**: Track 1, Track 2
- **Living Specs Touched**: `.cooper/specs/cooper-cli/spec.md`

---

### Track 4: Installer Integration & Release Packaging (`track-cooper-installer-packaging`)
- **Scope**:
  - Update `install.sh` to prompt for optional CLI binary installation.
  - Implement OS/architecture detection (`darwin`/`linux`, `arm64`/`amd64`) and release fetching.
  - Configure GitHub Actions CI release workflow with cross-platform binaries attached to releases.
  - Document CLI commands, MCP setup, and update flows in `docs/` and `README.md`.
- **Dependencies**: Track 1, Track 2, Track 3
- **Living Specs Touched**: `.cooper/specs/installer/spec.md`, `.cooper/specs/documentation/spec.md`
