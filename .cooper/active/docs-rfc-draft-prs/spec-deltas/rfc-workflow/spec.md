# Capability Specification: RFC Workflow & Draft PR Lifecycle

## Purpose & Scope
Defines the upstream architectural alignment workflow, Draft Pull Request review requirements, track registration gates, and supporting documentation for the Cooper RFC lifecycle.

## Requirements

### Requirement: Upstream Draft PR Review Surface
The Cooper framework SHALL mandate that all architectural RFC initiatives are initialized and reviewed as Draft Pull Requests.

#### Scenario: Opening RFC Pull Request
+ - GIVEN an architect or agent completes an RFC proposal and living spec deltas
+ - WHEN opening a pull request to the upstream repository
+ - THEN the pull request MUST be created with the `--draft` flag
+ - AND the PR description MUST outline the initiative, open questions, and impacted specs.

### Requirement: Two-Step Milestone Track Registration Gate
The Cooper framework SHALL prevent the registration and implementation of decomposed child tracks until team consensus and approval are reached on the RFC.

#### Scenario: RFC Iteration Prior to Approval
+ - GIVEN an RFC is in Draft PR review
+ - WHEN architectural discussion and spec delta iterations occur
+ - THEN child execution tracks SHALL NOT be registered in `.cooper/tracks.md` on trunk
+ - AND no downstream implementation tracks SHALL be spawned.

#### Scenario: RFC Graduation and Track Registration
+ - GIVEN an RFC has received formal approval from reviewers
+ - WHEN the RFC reaches consensus
+ - THEN the agent SHALL register the finalized child tracks into `.cooper/tracks.md`
+ - AND mark the PR ready for review (`gh pr ready`) for trunk merge.

### Requirement: Documentation of Draft PR Rationale
The Cooper framework SHALL provide dedicated documentation explaining the comparative advantages and governance rationale of Draft PRs over standard PRs.

#### Scenario: Accessing RFC Governance Guide
+ - GIVEN a contributor or agent seeking guidance on Cooper architectural workflows
+ - WHEN reading framework documentation
+ - THEN `docs/rfc-draft-prs.md` MUST provide the comparison matrix, 3 core benefits, and milestone lifecycle flow.
