# Implementation Plan: Pre-PR Phase Commit Squashing and Verification Protocol

- **Track ID**: `phase-squash-workflow`
- **Status**: Planning

## Phase 1: Framework Workflow & Guidelines Documentation Update [checkpoint: ]

- [ ] Task: Update Core Workflow Definition (`.cooper/definition/workflow.md`)
  - [ ] Sub-task: Draft Pre-PR Phase Squashing protocol and step-by-step rebase rules
  - [ ] Sub-task: Document Git Notes consolidation and `--force-with-lease` remote push guidelines
- [ ] Task: Update Framework Reference Manual & Cheatsheet (`.cooper/COOPER.md`)
  - [ ] Sub-task: Document phase commit structure and review benefits in lifecycle tables
- [ ] Task: Update Agent Guidelines & Scaffolding Templates
  - [ ] Sub-task: Update `AGENTS.md` operational rules
  - [ ] Sub-task: Update `AGENTS.template.md` and `internal/scaffold/assets/AGENTS.template.md`
- [ ] Task: Phase 1 Verification & Checkpoint

## Phase 2: Agent Skills Alignment [checkpoint: ]

- [ ] Task: Update `cooper-implement` Skill
  - [ ] Sub-task: Update Track Finalization section with phase squash execution steps
  - [ ] Sub-task: Add consolidated Git Notes attachment command instructions
- [ ] Task: Update `cooper-review` Skill
  - [ ] Sub-task: Add verification checklist item for clean 1-commit-per-phase PR history
- [ ] Task: Update `cooper-new-track` Skill
  - [ ] Sub-task: Update plan generation template to use checkpoint-level SHAs
- [ ] Task: Phase 2 Verification & Checkpoint

## Phase 3: Capability Specification & Validation Suite [checkpoint: ]

- [ ] Task: Establish Living Capability Spec (`.cooper/specs/workflow/spec.md`)
  - [ ] Sub-task: Initialize `.cooper/specs/workflow/spec.md` with base workflow rules and phase squashing requirements
- [ ] Task: Run SDD Linting, Validation & Test Suite
  - [ ] Sub-task: Run Go test suite (`go test ./...`)
  - [ ] Sub-task: Run Cooper specification and link validator
- [ ] Task: Phase 3 Verification & Checkpoint
