# Implementation Plan: Link Troop Mentions in Markdown Documentation to twoBoots/troop

## Phase 1: Core Framework Docs & Guides

- [x] Task: Audit and update Root & Agent Guidelines (32ec1ad)
  - [x] Sub-task: Update `README.md` to link first prominent Troop mention to `https://github.com/twoBoots/troop`
  - [x] Sub-task: Update `AGENTS.md` and `AGENTS.template.md` to link `[Troop](https://github.com/twoBoots/troop)` in Section 2
  - [x] Sub-task: Verify `CLAUDE.md` and root markdown files for Troop links
- [~] Task: Audit and update `docs/` Markdown Files
  - [ ] Sub-task: Update `docs/INSTALL.md` to link `[Troop](https://github.com/twoBoots/troop)`
  - [ ] Sub-task: Update `docs/openspec-vs-conductor-comparison.md` to link `[Troop](https://github.com/twoBoots/troop)`
  - [ ] Sub-task: Update `docs/rfc-draft-prs.md` to link `[Troop](https://github.com/twoBoots/troop)`
- [ ] Task: Phase 1 Verification & Checkpoint
  - [ ] Sub-task: Run automated markdown link check across Phase 1 files
  - [ ] Sub-task: Phase 1 Checkpoint & remote sync (`git push origin link-troop-in-docs`)

## Phase 2: Framework Definitions & Starter Templates

- [ ] Task: Update Cooper Internal Reference & Workflow
  - [ ] Sub-task: Update `.cooper/COOPER.md` to ensure primary link and table links point to `https://github.com/twoBoots/troop`
  - [ ] Sub-task: Update `.cooper/definition/workflow.md` to link `[Troop](https://github.com/twoBoots/troop)` in structure & protocol sections
- [ ] Task: Update Starter Templates (`templates/`)
  - [ ] Sub-task: Update `templates/product.md`, `templates/tech-stack.md`, and any template files mentioning Troop
- [ ] Task: Phase 2 Verification & Checkpoint
  - [ ] Sub-task: Run automated markdown link check across Phase 2 files
  - [ ] Sub-task: Phase 2 Checkpoint & remote sync (`git push origin link-troop-in-docs`)

## Phase 3: Project Skills & Final Audit

- [ ] Task: Audit and update Agent Skills (`skills/` and `.agents/skills/`)
  - [ ] Sub-task: Update `cooper-setup`, `cooper-rfc`, and `cooper-new-track` SKILL.md files
  - [ ] Sub-task: Update `cooper-implement`, `cooper-review`, and `cooper-status` SKILL.md files
  - [ ] Sub-task: Ensure both `skills/` and `.agents/skills/` are kept in sync
- [ ] Task: Comprehensive Markdown Link Verification & Quality Gate
  - [ ] Sub-task: Execute repository-wide markdown link audit script
  - [ ] Sub-task: Verify no broken markdown syntax or link collisions
- [ ] Task: Phase 3 Verification & Final Checkpoint
  - [ ] Sub-task: Phase 3 Checkpoint & remote sync (`git push origin link-troop-in-docs`)
