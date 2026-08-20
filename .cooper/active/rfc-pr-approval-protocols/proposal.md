# Track Proposal: RFC PR Approval Detection Protocols & Reviewer Instructions

## 1. Intent & Context
Currently, `cooper-rfc` specifies that an RFC graduates to approved once reviewers approve the RFC. However:
1. It does not define explicit CLI commands or automated logic for an AI agent to detect and verify human review approvals on the GitHub Pull Request (previously relying solely on manual chat prompts like "Approve the RFC").
2. The scaffolded Draft PR body does not provide explicit reviewer guidance for human team members on how to indicate approval (native GitHub Review "Approve" vs. comment triggers).
3. The RFC-to-Track lifecycle documentation needs clearer definition across the Two-Tier architecture in `.cooper/COOPER.md` and `.cooper/definition/workflow.md`.

## 2. Value & Benefit
- **Automated AI Governance**: Agents executing `cooper-rfc` can autonomously inspect PR review states (`gh pr view --json ...`) and discussion threads for `/approve` triggers.
- **Clear Team Onboarding**: Engineers and reviewers who receive Draft PRs immediately understand the review contract and how their actions drive track graduation.
- **Clean Architectural Separation**: Keeps downstream implementation rules (`workflow.md`) focused on TDD/quality governance while clearly documenting the upstream RFC-to-Track transition in `COOPER.md`.

## 3. Scope Boundaries
- **In Scope**:
  - Updating `skills/cooper-rfc/SKILL.md` and `.agents/skills/cooper-rfc/SKILL.md` with PR approval detection commands, `/approve` trigger recognition, and reviewer callout templates in Step 5.3.
  - Updating `.cooper/COOPER.md` to document the end-to-end RFC review, approval detection, registration, and merge handoff lifecycle.
  - Updating `.cooper/definition/workflow.md` (Principle 9) to clarify the upstream-to-downstream transition.
  - Creating Living Capability Spec for `cooper-rfc` in `.cooper/specs/cooper-rfc/spec.md` and Spec Delta in `.cooper/active/rfc-pr-approval-protocols/spec-deltas/cooper-rfc/spec.md`.
- **Out of Scope**:
  - Modifying other skills (`cooper-new-track`, `cooper-implement`, `cooper-review`, `cooper-status`, `cooper-setup`) beyond standard consistency.
