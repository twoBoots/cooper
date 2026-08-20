# Capability Specification Delta: Cooper Go CLI & MCP Server

## Requirements

### + Requirement: Core CLI Entrypoint & Versioning
+ The `cooper` CLI SHALL provide a standard command-line interface with subcommands, version reporting, and global execution flags.
+
+ #### + Scenario: Report Version
+ - GIVEN a user executes `cooper version`
+ - WHEN version flags are parsed
+ - THEN the CLI MUST output the semver version, git commit hash, and build timestamp.

### + Requirement: Project Initialization & Migration
+ The `cooper` CLI SHALL scaffold the complete Cooper directory structure and migrate legacy Conductor or OpenSpec repositories.
+
+ #### + Scenario: Scaffold Greenfield Project
+ - GIVEN an empty project directory
+ - WHEN running `cooper init`
+ - THEN the CLI MUST create `.cooper/`, `.agents/skills/`, and `AGENTS.md`.
+
+ #### + Scenario: Migrate Brownfield Conductor Project
+ - GIVEN a project containing `.conductor/`
+ - WHEN running `cooper init`
+ - THEN the CLI MUST migrate tracks and definitions from `.conductor/` into `.cooper/`.

### + Requirement: Deterministic Spec & Markdown Validation
+ The `cooper` CLI SHALL provide high-speed validation for living capability specs, spec deltas, track metadata, and repository documentation links.
+
+ #### + Scenario: Validate Spec Delta Format
+ - GIVEN a track directory with `spec-deltas/<capability>/spec.md`
+ - WHEN running `cooper validate` or `cooper lint`
+ - THEN the tool MUST verify that requirement additions are prefixed with `+`, deprecations with `-`, and scenarios conform to `GIVEN` / `WHEN` / `THEN` keywords.
+
+ #### + Scenario: Validate Track Metadata Schema
+ - GIVEN active tracks in `.cooper/active/`
+ - WHEN running `cooper validate`
+ - THEN the tool MUST verify that each `metadata.json` conforms to schema with valid track ID, status, and timestamps.

### + Requirement: Track & Worktree Orchestration
+ The `cooper` CLI SHALL orchestrate track lifecycle operations and Troop worktree management.
+
+ #### + Scenario: Spawn New Track Worktree
+ - GIVEN a valid track ID and title
+ - WHEN running `cooper track new <track_id>`
+ - THEN the CLI MUST spawn an isolated worktree via `git agent-start <track_id>`, scaffold `.cooper/active/<track_id>/`, and register the track in `.cooper/tracks.md`.
