# Spec Delta: Installer & Repo Structure

## Capability: installer

## Requirements

### Requirement: Claude Compatibility Bridge Scaffolding
+ `install.sh` and the `cooper-setup` workflow SHALL scaffold a `.claude/skills` symbolic link pointing to `../.agents/skills` and ensure `CLAUDE.md` bridges to `@AGENTS.md`.

#### Scenario: Greenfield Scaffolding of Claude Bridge
+ - GIVEN a target repository without `.claude/skills` or `CLAUDE.md`
+ - WHEN the installer or setup workflow scaffolds agent configuration
+ - THEN directory `.claude` MUST be created
+ - AND `.claude/skills` MUST be a symbolic link pointing to `../.agents/skills`
+ - AND `CLAUDE.md` MUST exist and contain `@AGENTS.md`.

#### Scenario: Preserving Existing Claude Skills Target
+ - GIVEN a target repository where `.claude/skills` already exists as a symlink, directory, or file
+ - WHEN the installer runs
+ - THEN the existing `.claude/skills` MUST remain untouched.

#### Scenario: Appending Bridge to Existing CLAUDE.md
+ - GIVEN a target repository where `CLAUDE.md` exists and does not contain `@AGENTS.md`
+ - WHEN the installer runs
+ - THEN `@AGENTS.md` MUST be appended to `CLAUDE.md`.

#### Scenario: Idempotent Existing CLAUDE.md Bridge
+ - GIVEN a target repository where `CLAUDE.md` already contains `@AGENTS.md`
+ - WHEN the installer runs
+ - THEN `CLAUDE.md` MUST NOT be modified and MUST NOT contain duplicate `@AGENTS.md` entries.
