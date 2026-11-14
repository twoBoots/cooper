# Proposal: Living Files Stand Alone

## Problem Statement
Living definition files (`.cooper/definition/*.md`) and living capability specifications (`.cooper/specs/<capability>/spec.md`, both at the repository root and inside any barrel) frequently accumulate planning provenance tokens during RFC authoring, track creation, and spec-delta merging. Three families of references break self-containment:

1. **Ticket, issue, PR, and track number references** (for example `#28`, `GH-12`, `PR #41`).
2. **Decision and open-question markers** (for example `D14`, `DEC-1`, `OQ-4`).
3. **Bare section numbers** (for example `§4.2`, `§3`, `see section 2.1`), including self-references within the same file.

When these tokens land in living files, readers are forced to trace closed trackers or archived working documents to understand the current specification, and bare section numbers silently break whenever headings are added, removed, or reordered. Conversely, active tracks, RFCs, and archived tracks are point-in-time working records where those markers remain natural and useful.

## Proposed Solution
Establish **Living Files Stand Alone** as an explicit Guiding Principle and specification standard across Cooper's workflow definitions, starter template, living documentation specification, and agent skills—relying on workflow rules and agent discipline rather than rigid regex linters or custom configuration files.

### Scope & Rules
- **Covered Living Files:** All files under `.cooper/definition/` and `.cooper/specs/` (at the root and inside every barrel).
- **Exempt Point-in-Time Records:** Files under `.cooper/active/` and `.cooper/archive/`.
- **Banned Categories in Living Files:**
  1. Ticket, issue, PR, and track numbers in prose and requirement text (including `GIVEN`/`WHEN`/`THEN` scenario steps).
  2. Decision and open-question markers (`D14`, `DEC-1`, `OQ-4`) outright in all living files.
  3. Bare section numbers (`§4.2`, `§3`, `section 2.1`), including self-references, replaced by exact heading names.
- **Two Bounded Exceptions (Strictly Outside Requirement Text):**
  1. **Rationale links:** Non-requirement rationale sections may link to a frozen archived track document, an approved or archived RFC, or a frozen macro contract using its repository file path plus exact heading name (for example `path/to/doc.md#Heading Name`).
  2. **Open known gaps:** Non-requirement known-gap entries may cite their tracking ticket while the gap remains unresolved, and must be removed once the ticket ships.

### Enforcement Surface
- Update `.cooper/definition/workflow.md` and `templates/workflow.md` with the **Living Files Stand Alone** Guiding Principle and a Track Finalization requirement to strip or rewrite planning provenance when merging Spec Deltas into `.cooper/specs/`.
- Update `.cooper/specs/documentation/spec.md` with the formal `Living Files Stand Alone` requirement and scenarios.
- Update Cooper skills (`cooper-new-track`, `cooper-rfc`, `cooper-implement`, `cooper-review`, `cooper-setup`) in both `skills/` and `.agents/skills/` so planning, implementation cleanup, code review, and brownfield repository audits consistently uphold the standard.

## Acceptance Criteria
1. `.cooper/definition/workflow.md` and `templates/workflow.md` document **Living Files Stand Alone** under Guiding Principles and require provenance cleanup during Track Finalization Spec Delta merges.
2. `.cooper/specs/documentation/spec.md` includes `### Requirement: Living Files Stand Alone` covering scope, the three banned reference types, and the two non-requirement exceptions.
3. `cooper-new-track`, `cooper-rfc`, `cooper-implement`, `cooper-review`, and `cooper-setup` enforce the rule (including brownfield audit checks in `cooper-setup`), with byte-for-byte parity between `skills/` and `.agents/skills/`.
4. Documentation/workflow tests (`tests/docs-workflow.test.mjs`) verify the workflow rule, template, living spec, and skill directives, and existing living files under `.cooper/definition/` and `.cooper/specs/` are audited and clean.
