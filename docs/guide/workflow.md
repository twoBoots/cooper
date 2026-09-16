# Cooper Workflow & Lifecycle

The Cooper framework establishes a repeatable lifecycle for engineering changes, moving from collaborative architecture proposals down to isolated, test-driven implementations.

<div class="lifecycle-flow">
  <div class="lifecycle-card">
    <span class="lifecycle-badge">Phase 1</span>
    <div class="lifecycle-step">/cooper-rfc</div>
    <div class="lifecycle-sub">Architectural Initiative</div>
  </div>
  <div class="lifecycle-arrow">→</div>
  <div class="lifecycle-card">
    <span class="lifecycle-badge">Phase 2</span>
    <div class="lifecycle-step">/cooper-new-track</div>
    <div class="lifecycle-sub">Worktree &amp; Spec Delta</div>
  </div>
  <div class="lifecycle-arrow">→</div>
  <div class="lifecycle-card">
    <span class="lifecycle-badge">Phase 3</span>
    <div class="lifecycle-step">/cooper-implement</div>
    <div class="lifecycle-sub">Strict TDD &amp; Git Notes</div>
  </div>
  <div class="lifecycle-arrow">→</div>
  <div class="lifecycle-card">
    <span class="lifecycle-badge">Phase 4</span>
    <div class="lifecycle-step">/cooper-review</div>
    <div class="lifecycle-sub">Spec Delta &amp; PR Gate</div>
  </div>
  <div class="lifecycle-arrow">→</div>
  <div class="lifecycle-card">
    <span class="lifecycle-badge">Ship</span>
    <div class="lifecycle-step">PR &amp; Merge</div>
    <div class="lifecycle-sub">Teardown Worktree</div>
  </div>
</div>

## 1. Architectural Initiative - `/cooper-rfc`
When a proposed change may introduce large architecture modifications or affects multiple codeowners, start with an RFC:
- Draft an architectural RFC in `.cooper/active/rfc-<name>/rfc.md`.
- Open a **Draft Pull Request** early to gather stakeholder and peer feedback.
- Once approved, run `/cooper-new-track` to break down work into focused tracks.

See [Draft PRs in the Cooper RFC Lifecycle](../rfc-draft-prs.md) for details.

## 2. Track Planning - `/cooper-new-track`
For single-capability features, bug fixes:
- Spawn an isolated [Troop](https://twoboots.github.io/troop) worktree via `git agent-start <track_id>`.
- Generate the **Proposal** (`proposal.md`), **Technical Design** (`design.md`), and **Spec Delta** (`spec-deltas/<capability>/spec.md`).
- Formulate a TDD-enforced implementation plan (`plan.md`) with explicit phases and checkpoints.
- Register the track in `.cooper/tracks.md`.

## 3. Implementation & TDD Loop - `/cooper-implement`
Work strictly inside the dedicated [Troop](https://twoboots.github.io/troop) worktree (`.worktrees/<track_id>`):
- **Task Selection**: Mark the next task in-progress (`[~]`) and commit the status change.
- **Red Phase**: Write targeted tests validating the Spec Delta requirements and ensure test failure.
- **Green Phase**: Write the minimum necessary production code to achieve test success.
- **Refactor Phase**: Optimize, clean code, and verify >80% test coverage.
- **Git Notes**: Record task summary metadata (`git notes add -m`) against the task commit to preserve implementation context against drift.
- **Phase Verification Checkpoint**:
  - Run `git fetch origin main` to synchronize workflow rules and living specs.
  - Run the full automated test suite.
  - Complete manual verification.
  - Create a checkpoint commit and push (`git push origin <track_id>`).

## 4. Quality Review - `/cooper-review`
Before opening or merging a PR, act as a principal engineer to:
- Audit all changed files against the track's Spec Deltas.
- Confirm 100% adherence to project code styleguides and test coverage criteria.
- Archive the track into `.cooper/archive/<track_id>/` upon completion.
- Tear down the worktree cleanly using `git agent-stop <track_id>` after human approval and pull request merge.

<style>
.lifecycle-flow {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  margin: 24px 0;
  flex-wrap: wrap;
}
.lifecycle-card {
  flex: 1 1 120px;
  min-width: 110px;
  background-color: var(--vp-c-bg-soft);
  border: 1px solid var(--vp-c-divider);
  border-radius: 8px;
  padding: 12px;
  text-align: center;
  box-sizing: border-box;
}
.lifecycle-badge {
  display: inline-block;
  font-size: 11px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: var(--vp-c-brand-1);
  margin-bottom: 4px;
}
.lifecycle-step {
  font-family: var(--vp-font-family-mono);
  font-size: 13px;
  font-weight: 700;
  color: var(--vp-c-text-1);
  margin-bottom: 4px;
}
.lifecycle-sub {
  font-size: 12px;
  color: var(--vp-c-text-2);
  line-height: 1.3;
}
.lifecycle-arrow {
  font-size: 18px;
  font-weight: bold;
  color: var(--vp-c-text-3);
  user-select: none;
}
@media (max-width: 640px) {
  .lifecycle-flow {
    flex-direction: column;
    align-items: stretch;
  }
  .lifecycle-arrow {
    text-align: center;
    transform: rotate(90deg);
  }
  .lifecycle-card {
    min-width: 100%;
  }
}
</style>
