# Proposal: Living Files Stand Alone

## Problem Statement
Living definition files (`.cooper/definition/*.md`) and capability specs (`.cooper/specs/<capability>/spec.md`, at the repo root and in every barrel) accumulate planning provenance during RFC authoring, track creation, and spec-delta merges:

1. **Ticket, issue, PR, and track numbers** (`#28`, `GH-12`, `PR #41`).
2. **Decision and open-question markers** (`D14`, `DEC-1`, `OQ-4`).
3. **Bare section numbers** (`§4.2`, `§3`, `section 2.1`), including self-references.

These tokens force readers to chase closed issues or archived working docs to understand current behaviour, and bare section numbers break silently when headings reorder. Active tracks, RFCs, and archived tracks are point-in-time working records where those markers belong.

## Proposed Solution
Enforce **Living Files Stand Alone** via Cooper's workflow definitions, starter template, documentation spec, and agent skills - without adding Go validator regex rules or `.cooperrc` configuration.

### Scope & Rules
- **Covered Living Files:** All files in `.cooper/definition/` and `.cooper/specs/` (repo root and barrels).
- **Exempt Point-in-Time Records:** `.cooper/active/` and `.cooper/archive/`.
- **Forbidden in Living Files:**
  1. Ticket, issue, PR, and track numbers in prose and requirement text (including `GIVEN`/`WHEN`/`THEN` steps).
  2. Decision and open-question markers (`D14`, `DEC-1`, `OQ-4`).
  3. Bare section numbers (`§4.2`, `§3`, `section 2.1`), including self-references. Use exact heading names instead.
- **Non-Requirement Exceptions (Outside Requirement Text and Scenario Steps Only):**
  1. **Rationale links:** Link a frozen archived track, approved or archived RFC, or frozen macro contract by repository file path and exact heading name (`path/to/doc.md#Heading Name`).
  2. **Open known gaps:** Cite the tracking ticket while the gap is open; remove the entry once the ticket ships.

## Acceptance Criteria
1. `.cooper/definition/workflow.md` and `templates/workflow.md` define **Living Files Stand Alone** under Guiding Principles and require provenance stripping during Track Finalization Spec Delta merges.
2. `.cooper/specs/documentation/spec.md` defines `### Requirement: Living Files Stand Alone` covering scope, forbidden references, and the two non-requirement exceptions.
3. `cooper-new-track`, `cooper-rfc`, `cooper-implement`, `cooper-review`, and `cooper-setup` enforce the rule (including brownfield audits in `cooper-setup`), with byte-for-byte parity between `skills/` and `.agents/skills/`.
4. Existing living files in `.cooper/definition/` and `.cooper/specs/` are audited and clean.
