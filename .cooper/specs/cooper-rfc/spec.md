# Capability Specification: Cooper RFC Planning & Review Protocol

## Purpose & Scope
Defines the architectural design, collaborative review, and approval workflow for the `cooper-rfc` skill within the Cooper Spec-Driven Development (SDD) framework.

## Requirements

### Requirement: Draft PR Reviewer Guidance
The `cooper-rfc` skill SHALL scaffold clear reviewer instructions in generated Draft Pull Request bodies to guide human peer review and approval actions.

#### Scenario: Include Reviewer Actions Block
- GIVEN an RFC Draft Pull Request is created by `cooper-rfc`
- WHEN the PR body (`prbody.md`) is scaffolded
- THEN it MUST include a dedicated `### 📝 Reviewer Actions` section explaining feedback submission, approval mechanics (GitHub Native `Approve` review or `/approve` comment trigger), and graduation triggers.

### Requirement: Automated PR Approval State Detection
The `cooper-rfc` skill SHALL inspect GitHub Pull Request status to detect native review approvals.

#### Scenario: Native GitHub Review Approval
- GIVEN an AI agent checking PR review status during `cooper-rfc`
- WHEN inspecting GitHub PR state with `gh pr view --json state,reviews,reviewDecision`
- THEN the agent MUST recognize `reviewDecision == "APPROVED"` or approved reviews as formal RFC approval and initiate track graduation.

### Requirement: PR Comment Approval Trigger
The `cooper-rfc` skill SHALL parse discussion comments for maintainer approval triggers.

#### Scenario: Parse /approve Comment Trigger
- GIVEN an AI agent checking PR comments with `gh pr view --comments`
- WHEN a comment contains `/approve` from a maintainer or reviewer
- THEN the agent MUST recognize the comment trigger as RFC approval and proceed with track graduation.

### Requirement: RFC Graduation & Merge Gate
The `cooper-rfc` skill SHALL update metadata, register child implementation tracks, mark the PR ready for merge, and enforce human merge approval before downstream execution begins.

#### Scenario: Transition to Approved and Register Tracks
- GIVEN an approved RFC in `cooper-rfc`
- WHEN graduating the RFC
- THEN the agent MUST update `metadata.json` status to `approved`, update `rfc.md` status to `Approved`, append decomposed child tracks to `.cooper/tracks.md`, mark the PR ready with `gh pr ready`, and await human maintainer merge to `main`.
