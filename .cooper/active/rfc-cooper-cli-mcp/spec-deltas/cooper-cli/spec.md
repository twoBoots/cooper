# Capability Specification: Cooper Go CLI & MCP Server

## Purpose & Scope
Defines the CLI command interface, 3-way upstream reconciliation engine, and embedded Model Context Protocol (MCP) server for the Cooper Spec-Driven Development framework.

## Requirements

### Requirement: Deterministic Spec & Markdown Validation
The `cooper` CLI SHALL provide high-speed validation for living capability specs, spec deltas, track metadata, and repository documentation links.

#### Scenario: Validate Spec Delta Format
- GIVEN a track directory with `spec-deltas/<capability>/spec.md`
- WHEN running `cooper validate` or `cooper lint`
- THEN the tool MUST verify that requirement additions are prefixed with `+`, deprecations with `-`, scenarios conform to `GIVEN` / `WHEN` / `THEN` keywords, and metadata JSON files match schema.

### Requirement: Upstream Synchronization & 3-Way Reconciliation
The `cooper` CLI SHALL synchronize project skills, rules, and templates with upstream releases while preserving project-specific customizations.

#### Scenario: Compute 3-Way Merge on Update
- GIVEN a project initialized with version `V1` having local modifications to `.cooper/definition/workflow.md`
- WHEN `cooper update` or MCP `cooper_apply_update` runs against upstream version `V2`
- THEN the engine MUST compute a 3-way diff between base `V1`, local modified `V1'`, and remote `V2`, merging non-conflicting enhancements and flagging conflicts.

### Requirement: Embedded Agent MCP Server
The `cooper` CLI SHALL provide an embedded Model Context Protocol (MCP) server running over standard input/output (`stdio`).

#### Scenario: Serve Structured SDD Tools over MCP
- GIVEN an AI coding agent connected to `cooper mcp`
- WHEN the agent requests available tools
- THEN the server MUST expose `cooper_track_create`, `cooper_checkpoint_record`, `cooper_spec_delta_validate`, `cooper_check_updates`, `cooper_diff_updates`, and `cooper_apply_update`.

#### Scenario: Agent-Guided Upstream Upgrade over MCP
- GIVEN an agent invoking `cooper_diff_updates`
- WHEN inspecting upstream changes
- THEN the MCP server MUST return structured diffs categorized by non-conflicting updates and customized conflicts, allowing the agent to guide the user interactively through selective adoption.

### Requirement: Binary Self-Updating
The `cooper` CLI SHALL support updating its own binary directly from GitHub Releases.

#### Scenario: Self-Update to Latest Release
- GIVEN a user running `cooper self-update` or `cooper update --self`
- WHEN a newer version exists in GitHub Releases
- THEN the CLI MUST download the platform binary for the current OS/architecture, replace the running executable in-place, and apply macOS codesigning/quarantine fixes on Darwin.
