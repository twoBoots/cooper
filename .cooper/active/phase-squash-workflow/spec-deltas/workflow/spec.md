# Capability Specification Delta: SDD Workflow & Track Lifecycle

- **Capability**: `workflow`
- **Track ID**: `phase-squash-workflow`

## Requirements

### Requirement: Pre-PR Phase Commit Squashing
The Cooper SDD framework SHALL require that before opening a Pull Request, all granular task and plan commits from each phase within an active track are squashed into an atomic, semantic milestone commit per phase.

#### Scenario: Squashing Phase Commits Prior to PR Creation
+ GIVEN an active track in `.worktrees/<track_id>` where all phases and tasks in `plan.md` are completed
+ WHEN the track enters the finalization step prior to opening a Pull Request (`gh pr create`)
+ THEN each phase MUST be consolidated into a single commit formatted as `feat(<scope>): Phase <N> - <Phase Title>` or `fix(<scope>): Phase <N> - <Phase Title>`
+ AND the squashed commit body MUST contain the summary of completed tasks within that phase.

### Requirement: Phase Git Note Verification Consolidation
When phase commits are squashed, the framework SHALL attach a consolidated Phase Verification report as a Git Note to the newly generated squashed commit SHA.

#### Scenario: Attaching Verification Git Note to Squashed Phase SHA
+ GIVEN a squashed Phase commit
+ WHEN the commit SHA is generated
+ THEN a Git Note MUST be attached containing the phase verification status, automated test results, user manual verification confirmation, and timestamp.

### Requirement: Plan Checkpoint SHA Tracking
The framework SHALL record the squashed commit SHA at the phase checkpoint level in `plan.md` rather than maintaining transient per-task commit SHAs.

#### Scenario: Recording Phase Checkpoint SHA
+ GIVEN completed tasks within a phase
+ WHEN the phase is squashed
+ THEN the phase header or checkpoint line in `plan.md` MUST record `[checkpoint: <squashed_sha>]`
+ AND individual task list items SHALL remain clean status indicators (`- [x] Task Title`).

### Requirement: Safe Remote Synchronization
When pushing a squashed branch history to remote, the framework SHALL use lease-protected force push.

#### Scenario: Pushing Rewritten Track History
+ GIVEN a track branch with squashed phase commits that were previously pushed as granular commits
+ WHEN synchronizing with remote `origin`
+ THEN the command executed MUST be `git push --force-with-lease origin <track_id>`.
