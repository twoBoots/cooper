# Track Proposal: Scaffold Claude Bridge (`.claude/skills` and `CLAUDE.md`)

## Rationale & Goal
Cooper installs and maintains agent skills under `.agents/skills/` and agent guidelines under `AGENTS.md` as its agent-agnostic foundation. However, several popular agent CLI tools (e.g. Claude Code) look specifically for project skills under `.claude/skills/<name>/SKILL.md` and auto-load project guidelines from root `CLAUDE.md`.

In the `twoBoots/cooper` repository itself, this bridge is established using:
- `.claude/skills` -> symbolic link pointing to `../.agents/skills`
- `CLAUDE.md` -> containing `@AGENTS.md`

Currently, `install.sh` and the `cooper-setup` skill do not scaffold this bridge into consumer projects during installation or setup. Consequently, agents discovering skills via `.claude/skills/` see no available `cooper-*` skills even after successful installation.

This track addresses GitHub issue [#26](https://github.com/twoBoots/cooper/issues/26) by adding automated, idempotent scaffolding of `.claude/skills` and `CLAUDE.md` to `install.sh` and updating the `cooper-setup` skill documentation.

## User Benefit
- **Seamless Skill Discovery**: Agents utilizing `.claude/skills/` instantly discover all installed Cooper skills (`cooper-setup`, `cooper-rfc`, `cooper-new-track`, `cooper-implement`, `cooper-review`, `cooper-status`).
- **Single Source of Truth**: Rules are not duplicated; `CLAUDE.md` cleanly imports `@AGENTS.md`.
- **Zero Friction & Safe Idempotency**: Existing custom `CLAUDE.md` files or custom `.claude/skills` directories/links are respected and preserved without accidental clobbering.

## Scope of Changes
1. **Installer (`install.sh`)**:
   - Add step to create `.claude` directory if missing.
   - Create symbolic link `.claude/skills -> ../.agents/skills` if `.claude/skills` does not already exist.
   - If `CLAUDE.md` is missing, create it containing `@AGENTS.md`.
   - If `CLAUDE.md` exists and does not already reference `@AGENTS.md`, append `@AGENTS.md`.
2. **Skill Documentation (`skills/cooper-setup/SKILL.md` & `.agents/skills/cooper-setup/SKILL.md`)**:
   - Document the `.claude/skills` symlink and `CLAUDE.md` bridge in the setup instructions.
3. **Living Capability Spec Delta (`spec-deltas/installer/spec.md`)**:
   - Formulate requirement and scenarios for the Claude compatibility bridge.
4. **Test Suite (`cmd/installer_test.go`)**:
   - Unit tests covering greenfield creation, idempotency, and non-clobbering behavior for existing `CLAUDE.md` and `.claude/skills`.
