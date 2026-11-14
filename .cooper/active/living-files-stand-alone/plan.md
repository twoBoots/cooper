# Implementation Plan: Living Files Stand Alone

## Phase 1: Workflow Policy, Template & Living Documentation Spec
- [ ] Task: Add Living Files Stand Alone to Workflow Policy, Starter Template, and Documentation Spec
  - [ ] Sub-task: Write verification assertions in `tests/docs-workflow.test.mjs` checking `.cooper/definition/workflow.md`, `templates/workflow.md`, and `.cooper/specs/documentation/spec.md` for `Living Files Stand Alone`, banned categories, non-requirement exceptions, and Track Finalization provenance cleanup (Red)
  - [ ] Sub-task: Update `.cooper/definition/workflow.md`, `templates/workflow.md`, and `.cooper/specs/documentation/spec.md` with the `Living Files Stand Alone` Guiding Principle, Track Finalization cleanup rule, and formal capability requirement/scenarios (Green)
  - [ ] Sub-task: Verify formatting, heading consistency, and placeholder paths across workflow and spec files (Refactor)
- [ ] Task: Phase 1 Verification & Checkpoint
  - [ ] Sub-task: Synchronize rules (`git fetch origin main`)
  - [ ] Sub-task: Verify phase 1 changes and attach Git Note (`git notes add -m`)
  - [ ] Sub-task: Commit phase checkpoint and push (`git push origin living-files-stand-alone`)

## Phase 2: Agent Skill Instructions (`skills/` & `.agents/skills/`)
- [ ] Task: Enforce Living Files Stand Alone in Planning, RFC, Implementation, Review, and Setup Skills
  - [ ] Sub-task: Add test assertions in `tests/docs-workflow.test.mjs` verifying `cooper-new-track`, `cooper-rfc`, `cooper-implement`, `cooper-review`, and `cooper-setup` enforce `Living Files Stand Alone` (including brownfield audit checks in `cooper-setup`) and maintain byte-for-byte parity between `skills/` and `.agents/skills/` (Red)
  - [ ] Sub-task: Update `skills/cooper-new-track/SKILL.md`, `skills/cooper-rfc/SKILL.md`, `skills/cooper-implement/SKILL.md`, `skills/cooper-review/SKILL.md`, and `skills/cooper-setup/SKILL.md` and mirror byte-for-byte to `.agents/skills/` (Green)
  - [ ] Sub-task: Audit skill headings (`### N.M`), required sections (`Interactive Question Protocol`, `Native File Tools Mandate`), and backticked paths (`<...>` placeholders) for `cmd/skills_test.go` and `code_path_auditor.go` compliance (Refactor)
- [ ] Task: Phase 2 Verification & Checkpoint
  - [ ] Sub-task: Synchronize rules (`git fetch origin main`)
  - [ ] Sub-task: Verify phase 2 changes and attach Git Note (`git notes add -m`)
  - [ ] Sub-task: Commit phase checkpoint and push (`git push origin living-files-stand-alone`)

## Phase 3: Living Files Audit & Regression Suite
- [ ] Task: Audit Living Definition and Specification Files for Forbidden Provenance Tokens
  - [ ] Sub-task: Add living-file cleanliness assertions in `tests/docs-workflow.test.mjs` scanning `.cooper/definition/*.md` and `.cooper/specs/*/spec.md` for forbidden ticket/issue/PR numbers, decision/open-question markers, and bare section numbers (Red)
  - [ ] Sub-task: Audit `.cooper/definition/` and `.cooper/specs/` and remediate any violations if found (Green)
  - [ ] Sub-task: Verify full track artifacts and spec delta formatting (Refactor)
- [ ] Task: Phase 3 Verification & Checkpoint
  - [ ] Sub-task: Synchronize rules (`git fetch origin main`)
  - [ ] Sub-task: Verify phase 3 changes and attach Git Note (`git notes add -m`)
  - [ ] Sub-task: Commit phase checkpoint and push (`git push origin living-files-stand-alone`)
