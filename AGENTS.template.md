# Agent Guidelines (Cooper + Troop Workflow)

## Rules

1. **Conductor Mandate**:
   - All feature development, bug fixes, and significant changes MUST follow the **Conductor Workflow**.
   - Refer to `conductor/workflow.md` for the specific track lifecycle (TDD, plan updates, git notes, checkpoints, and quality gates).

2. **Troop Isolation Protocol**:
   - Work inside an isolated worktree under `.worktrees/<track_id>`. Do NOT modify code in the repository main trunk directly unless explicitly instructed.
   - Base track worktrees off `main` using `git agent-start <track_id>`.

3. **Execution & Cleanup**:
   - List active track worktrees with `git troop`.
   - Teardown completed worktrees with `git agent-stop <track_id>` after PR approval and merge.

4. **Source of Truth**:
   - Read `conductor/workflow.md` for full lifecycle and workflow guidelines.
