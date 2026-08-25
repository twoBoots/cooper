# Technical Design: Pre-PR Phase Commit Squashing and Verification Protocol

- **Track ID**: `phase-squash-workflow`
- **Status**: Planning

## 1. Architecture Overview

The Cooper SDD lifecycle is structured around discrete, verified **Phases**. This design formalizes the distinction between **Development History** (granular, safety-oriented, private to the worktree) and **Published History** (phase-atomic, bisectable, published to `main`).

```
[Phase 1 Active Tasks] ──> Granular TDD Commits (Private)
[Phase 2 Active Tasks] ──> Granular TDD Commits (Private)
             │
             ▼
   Track All Phases Done
             │
             ▼
[Pre-PR Phase Squash Step]
   ├─> Squash Phase 1 Commits ──> "feat(<scope>): Phase 1 - <Title>" + Git Note
   ├─> Squash Phase 2 Commits ──> "feat(<scope>): Phase 2 - <Title>" + Git Note
   ├─> Update plan.md Checkpoints ──> [checkpoint: <squashed_sha>]
   └─> Push Remote ──> git push --force-with-lease origin <track_id>
             │
             ▼
[Open Pull Request] ──> gh pr create (Atomic multi-commit PR)
```

## 2. Commit Message & Git Notes Specification

### 2.1 Squashed Phase Commit Message Format
```text
feat(<scope>): Phase <N> - <Phase Title>

- <Task 1 Title>: <Brief summary of changes>
- <Task 2 Title>: <Brief summary of changes>
- Verified against Spec Delta: .cooper/active/<track_id>/spec-deltas/<capability>/spec.md
- Test Coverage: >80%
```

### 2.2 Phase Git Note Attachment
On the squashed commit SHA, attach the consolidated Phase Verification note:
```bash
git notes add -m "Phase <N> Checkpoint Verification
Status: PASSED
Automated Tests: PASSED (>80% coverage)
Manual Verification: APPROVED by user
Spec Delta: .cooper/active/<track_id>/spec-deltas/<capability>/spec.md
Timestamp: <ISO-8601 Timestamp>" <squashed_phase_sha>
```

### 2.3 `plan.md` Checkpoint Structure
Individual tasks are checked off cleanly:
```markdown
## Phase 1: Core Storage Engine [checkpoint: 4a8b1c2]
- [x] Task: Key-Value Database Schema
  - [x] Sub-task: Write migration tests (Red)
  - [x] Sub-task: Implement schema migration (Green)
- [x] Task: Session Serialization
  - [x] Sub-task: Write serialization unit tests (Red)
  - [x] Sub-task: Implement JSON codec (Green)
- [x] Task: Phase 1 Verification & Checkpoint
```

## 3. Squashing Execution Mechanisms

Agents or developers can execute phase squashing using one of two supported methods:
1. **Interactive Rebase / Soft Reset Sequence:** Reset to base/phase boundary, stage all changes for that phase, commit with structured phase message, attach Git Note.
2. **Cooper CLI Helper Command (Future/Enhancement):** `cooper track finalize <track_id> --squash-phases`.

## 4. Skill & Rule Updates

1. **`cooper-implement`**:
   - Updates Section 4 (Track Finalization) with explicit Pre-PR Phase Squashing instructions before prompt to create PR.
   - Clarifies that per-task commits during execution remain granular.
2. **`cooper-review`**:
   - Updates review criteria to verify that PR branches present one clean, squashed commit per phase prior to merge approval.
3. **`cooper-new-track`**:
   - Formats generated `plan.md` templates with phase-level checkpoint placeholders `[checkpoint: ]`.
4. **`workflow.md` & `AGENTS.md`**:
   - Adds rule specifying the Pre-PR Phase Commit Squashing Protocol and `--force-with-lease` remote push rule.
