# Technical Design: Dedicated Documentation for RFC Draft PRs

- **Track ID**: `docs-rfc-draft-prs`
- **Capability**: `rfc-workflow`

## 1. Documentation Architecture

To keep `.cooper/COOPER.md` clean and lightweight as a developer manual and cheatsheet, all detailed architectural rationale, comparisons, and workflows for RFC Draft PRs are placed in a dedicated guide:

`docs/rfc-draft-prs.md`

### 1.1 Content Structure of `docs/rfc-draft-prs.md`
1. **Executive Summary & SDD Rationale**: The two-tier model separating upstream architectural alignment from downstream execution tracks.
2. **Comparison Matrix: Draft RFC PR vs Regular PR**:
   | Dimension | Draft RFC PR | Regular PR |
   | :--- | :--- | :--- |
   | **GitHub Merge State** | 🚫 **Merge blocked** by GitHub. Cannot be accidentally merged prematurely. | ✅ Mergeable at any time by maintainers or auto-merge rules. |
   | **Intended Signal** | *"This architecture is under discussion and actively seeking feedback."* | *"This work is finished and ready to be committed to `main`."* |
   | **Living Spec & Track Integrity** | Spec deltas and track breakdown iterate on the branch until consensus is solid. | If merged early, stale or unaligned tracks and spec deltas get committed to `main`. |
   | **CI / Notification Noise** | Bypasses heavy code CI/CD pipelines (e.g. Docker builds, staging deploys) while retaining Markdown review. | Triggers full CI runs and aggressive reviewer assignment policies. |
   | **Graduation Trigger** | Mark `Ready for review` + register child tracks **only after** team approves the design. | Track registration has to happen upfront before team consensus is reached. |
3. **The 3 Primary Benefits of a Draft RFC PR**:
   - Benefit 1: Prevents Premature Commit of Unapproved Specs & Tracks (preserves `.cooper/tracks.md` integrity).
   - Benefit 2: Safe Collaborative Review Surface (rendered Markdown, inline diff comments, Mermaid diagrams).
   - Benefit 3: Clean Two-Step Milestone Gate.
4. **Milestone Flow Diagram**:
   ```mermaid
   sequenceDiagram
       autonumber
       actor Architect as System Architect / Agent
       actor Team as Engineering Team
       participant PR as GitHub Draft PR
       participant Registry as .cooper/tracks.md (main)
       participant Track as Cooper Track (.worktrees/<track_id>)

       Architect->>PR: Open Draft PR (gh pr create --draft)
       Team->>PR: Review & Iterate on Spec Deltas / Architecture
       Team->>PR: Approve RFC
       Architect->>PR: Register Finalized Tracks in .cooper/tracks.md
       Architect->>PR: Mark Ready for Review (gh pr ready)
       Team->>Registry: Merge PR to main
       Architect->>Track: Pick up child track via cooper-new-track
   ```

### 1.2 Capability Specification Model
We introduce a living capability specification under `.cooper/specs/rfc-workflow/spec.md` (via Spec Delta `.cooper/active/docs-rfc-draft-prs/spec-deltas/rfc-workflow/spec.md`).

### Requirements:
- `Requirement: Upstream Draft PR Review Surface`
- `Requirement: Two-Step Milestone Track Registration Gate`
- `Requirement: Documentation of Draft PR Rationale`
