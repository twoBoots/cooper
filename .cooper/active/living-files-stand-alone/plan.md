# Implementation Plan: Living Files Stand Alone

## Phase 1: Workflow Policy, Template & Living Documentation Spec
- [x] Task: Add Living Files Stand Alone to Workflow Policy, Starter Template, and Documentation Spec (97a6b3f)
  - [x] Sub-task: Write verification assertions in `tests/docs-workflow.test.mjs` checking `.cooper/definition/workflow.md`, `templates/workflow.md`, and `.cooper/specs/documentation/spec.md` for `Living Files Stand Alone`, banned categories, non-requirement exceptions, and Track Finalization provenance cleanup (Red)
  - [x] Sub-task: Update `.cooper/definition/workflow.md`, `templates/workflow.md`, and `.cooper/specs/documentation/spec.md` with the `Living Files Stand Alone` Guiding Principle, Track Finalization cleanup rule, and formal capability requirement/scenarios (Green)
  - [x] Sub-task: Verify formatting, heading consistency, and placeholder paths across workflow and spec files (Refactor)
- [x] Task: Phase 1 Verification & Checkpoint (26f16de) [checkpoint: 26f16de]
  - [x] Sub-task: Synchronize rules (`git fetch origin main`)
  - [x] Sub-task: Verify phase 1 changes and attach Git Note (`git notes add -m`)
  - [x] Sub-task: Commit phase checkpoint and push (`git push origin living-files-stand-alone`)

## Phase 2: Agent Skill Instructions & Living Files Audit (`skills/` & `.agents/skills/`)
- [x] Task: Enforce Living Files Stand Alone in Planning, RFC, Implementation, Review, and Setup Skills (52f5319)
  - [x] Sub-task: Verify baseline skill files and `.cooper/definition/` + `.cooper/specs/` living files (Red)
  - [x] Sub-task: Update `skills/cooper-new-track/SKILL.md`, `skills/cooper-rfc/SKILL.md`, `skills/cooper-implement/SKILL.md`, `skills/cooper-review/SKILL.md`, and `skills/cooper-setup/SKILL.md` and mirror byte-for-byte to `.agents/skills/` (Green)
  - [x] Sub-task: Audit skill headings (`### N.M`), required sections (`Interactive Question Protocol`, `Native File Tools Mandate`), backticked paths (`<...>` placeholders), and living files under `.cooper/definition/` and `.cooper/specs/` (Refactor)
- [~] Task: Phase 2 Verification & Checkpoint
  - [ ] Sub-task: Synchronize rules (`git fetch origin main`)
  - [ ] Sub-task: Verify phase 2 changes and attach Git Note (`git notes add -m`)
  - [ ] Sub-task: Commit phase checkpoint and push (`git push origin living-files-stand-alone`)
