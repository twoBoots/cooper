# Track Proposal: Document Draft PR Rationale in Cooper RFC Lifecycle

- **Track ID**: `docs-rfc-draft-prs`
- **Type**: Documentation
- **Related Issue**: GitHub Issue #5 (`docs(rfc): Document rationale and benefits of Draft PRs for RFC lifecycle`)

## 1. Summary & Motivation

In Cooper Spec-Driven Development (SDD), architectural initiatives are initiated upstream via `cooper-rfc` and opened as **Draft Pull Requests** on GitHub/GitLab. To clarify why Draft PRs are chosen over standard mergeable Pull Requests and provide detailed architectural rationale without cluttering the core reference manual (`.cooper/COOPER.md`), this track introduces a dedicated documentation guide under `docs/rfc-draft-prs.md`.

## 2. Target Scope & Deliverables

1. **Dedicated Architecture Guide (`docs/rfc-draft-prs.md`)**:
   - Comprehensive breakdown of why Cooper architectural RFCs are opened as Draft PRs.
   - Comparative matrix: Draft RFC PR vs. Regular PR.
   - Detailed analysis of the 3 primary benefits (preventing premature track/spec commits, collaborative review surface, two-step milestone gate).
   - Sequence/flow diagram illustrating the RFC graduation and merge gate.

2. **Living Spec Delta (`spec-deltas/rfc-workflow/spec.md`)**:
   - Establish formal living specification requirements for the Cooper RFC lifecycle and Draft PR mechanics.

3. **Lightweight Reference Pointers**:
   - Ensure `.cooper/COOPER.md` and `README.md` link cleanly to `docs/rfc-draft-prs.md` where appropriate, without cluttering core cheatsheets.

## 3. Impact & Value
- **Zero Clutter**: Keeps `.cooper/COOPER.md` concise and focused as a developer cheatsheet.
- **Deep Reference**: Gives contributors and teams a complete, authoritative rationale for the Draft PR governance model.
