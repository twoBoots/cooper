# Capability Spec Delta: Cooper RFC Review & Approval Protocol

## Added Requirements

+ GIVEN an RFC Draft Pull Request is created by `cooper-rfc`
+ WHEN the PR body (`prbody.md`) is scaffolded
+ THEN it MUST include a dedicated `### Reviewer Actions` section explaining feedback submission, approval mechanics (GitHub Native `Approve` or `/approve` comment), and graduation triggers.

+ GIVEN an AI agent checking PR review status during `cooper-rfc`
+ WHEN inspecting GitHub PR state with `gh pr view --json state,reviews,reviewDecision`
+ THEN the agent MUST recognize `reviewDecision == "APPROVED"` or approved reviews as formal RFC approval and initiate track graduation.

+ GIVEN an AI agent checking PR comments with `gh pr view --comments`
+ WHEN a comment contains `/approve` from a maintainer or reviewer
+ THEN the agent MUST recognize the comment trigger as RFC approval and proceed with track graduation.

+ GIVEN an approved RFC in `cooper-rfc`
+ WHEN graduating the RFC
+ THEN the agent MUST update `metadata.json` status to `approved`, update `rfc.md` status to `Approved`, append decomposed child tracks to `.cooper/tracks.md`, mark the PR ready with `gh pr ready`, and await human maintainer merge to `main`.

## Removed / Modified Requirements

- GIVEN an RFC in `cooper-rfc`
- WHEN waiting for approval
- THEN relying solely on explicit user prompts in chat without automated PR status detection is replaced by automated PR review inspection and comment trigger detection.
