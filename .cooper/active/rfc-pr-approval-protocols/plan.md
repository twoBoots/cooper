# Implementation Plan: RFC PR Approval Detection Protocols & Reviewer Instructions

## Phase 1: Skill Enhancements (`cooper-rfc/SKILL.md`)

- [x] Task 1.1: Add Reviewer Guidance Block to PR Scaffolding (Step 5.3) (af2ccb8)
  - [x] Sub-task: Update `skills/cooper-rfc/SKILL.md` with `### 📝 Reviewer Actions` template block in `prbody.md`
  - [x] Sub-task: Update `.agents/skills/cooper-rfc/SKILL.md` with identical template block
  - [x] Sub-task: Verify markdown formatting and structure
- [x] Task 1.2: Add PR Approval State Detection & Comment Triggers (Section 6) (6fd001c)
  - [x] Sub-task: Update Step 6.1 in `skills/cooper-rfc/SKILL.md` with `gh pr view --json state,reviews,reviewDecision` command
  - [x] Sub-task: Add `/approve` comment trigger parsing instructions in Step 6.1
  - [x] Sub-task: Update `.agents/skills/cooper-rfc/SKILL.md` to match
- [x] Task 1.3: Align RFC Graduation & User Merge Gate Protocol (Step 6.2) (6fd001c)
  - [x] Sub-task: Document `metadata.json` status change, track registration, `gh pr ready`, and merge handoff
  - [x] Sub-task: Sync `.agents/skills/cooper-rfc/SKILL.md`
- [~] Task 1.4: Phase 1 Verification & Checkpoint
  - [ ] Sub-task: Verify skill syntax and step-by-step clarity
  - [ ] Sub-task: Checkpoint commit and phase sync

## Phase 2: Framework Documentation & Living Specs

- [ ] Task 2.1: Update `.cooper/COOPER.md` Two-Tier Planning Architecture
  - [ ] Sub-task: Document RFC Draft PR -> Review Approval -> Track Registration -> User PR Merge -> Downstream Track execution
  - [ ] Sub-task: Ensure terminology and commands align with `cooper-rfc` skill
- [ ] Task 2.2: Refine Principle 9 in `.cooper/definition/workflow.md`
  - [ ] Sub-task: Clarify handoff boundary from merged RFCs to downstream track workflow
- [ ] Task 2.3: Establish Living Capability Spec for `cooper-rfc`
  - [ ] Sub-task: Create `.cooper/specs/cooper-rfc/spec.md` with baseline requirements
- [ ] Task 2.4: Phase 2 Verification & Checkpoint
  - [ ] Sub-task: Verify living specs and documentation consistency
  - [ ] Sub-task: Checkpoint commit and phase sync

## Phase 3: Installer & Packaging Verification

- [ ] Task 3.1: Verify Installer Sync
  - [ ] Sub-task: Run/dry-run `install.sh` or verify skill file copies from `skills/` to `.agents/skills/`
- [ ] Task 3.2: Phase 3 Checkpoint & Finalization
  - [ ] Sub-task: Run repository verification and stage for PR
