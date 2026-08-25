# Track Proposal: Pre-PR Phase Commit Squashing and Verification Protocol

- **Track ID**: `phase-squash-workflow`
- **Type**: feature
- **Status**: Planning

## 1. Summary

This track updates the Cooper SDD framework workflow definitions, agent guidelines, and local skills to implement a **Pre-PR Phase Commit Squashing Protocol**. 

During active task execution, implementers retain granular TDD commits and execution notes for agent safety and rollback. Prior to opening a Pull Request (`gh pr create`), all commits within each phase are consolidated into a single atomic, semantic milestone commit per phase.

## 2. Motivation & Problem Statement

Currently, executing a multi-phase Cooper track produces numerous granular commits:
- Task start/complete commits (`chore(cooper): Start task '<name>'`, `cooper(plan): Complete task '<name>'`)
- Intermediate Red/Green/Refactor code commits
- Frequent `plan.md` checkbox edits

While this granularity provides critical safety and rollback capabilities during active coding:
1. It pollutes the repository Git history with internal agent scaffolding churn.
2. Standard PR squash-on-merge destroys the phase narrative entirely by collapsing the entire feature into one monolithic commit.
3. Reviewers lose the ability to review architectural phases incrementally.

## 3. Proposed Solution

1. **Granular Execution Lifecycle (Preserved):** Implementers continue writing granular TDD commits and attaching task-level notes during execution within the isolated [Troop](https://github.com/twoBoots/troop) worktree.
2. **Pre-PR Phase Squashing:** During Track Finalization (immediately before `gh pr create`):
   - All commits for each phase are squashed into an atomic commit: `feat(<scope>): Phase <N> - <Phase Title>`.
   - Task execution summaries are consolidated into the phase commit message body.
   - The Phase Checkpoint verification report is attached as a Git Note on the squashed phase commit SHA.
   - `plan.md` retains clean checkbox statuses (`- [x] Task`) and records the squashed commit SHA at the phase checkpoint header (`[checkpoint: <squashed_sha>]`).
3. **Remote Force-Push Alignment:** Update remote synchronization protocols to use `git push --force-with-lease origin <track_id>` following final squashing.

## 4. Scope & Impacted Boundaries

- `.cooper/definition/workflow.md` & `.cooper/COOPER.md`
- `.cooper/specs/workflow/spec.md` (Living Capability Spec & Spec Delta)
- `AGENTS.md`, `AGENTS.template.md`, `internal/scaffold/assets/AGENTS.template.md`
- `.agents/skills/cooper-implement/SKILL.md`, `.agents/skills/cooper-review/SKILL.md`, `.agents/skills/cooper-new-track/SKILL.md`
- Go CLI track and checkpoint helpers (`internal/track/`)
