# Spec Delta: Living Files Stand Alone

Target Spec: `.cooper/specs/documentation/spec.md`

## Added Requirements

### + Requirement: Living Files Stand Alone
+ All living framework files under `.cooper/definition/` and `.cooper/specs/` (at the repository root and inside every barrel) SHALL remain self-contained and understandable without consulting external issue trackers, pull requests, planning markers, or positional section numbers. Living files MUST NOT contain ticket/issue/PR/track numbers, decision or open-question markers (such as `D14`, `DEC-1`, `OQ-4`), or bare section numbers (such as `§4.2` or `section 3`, including self-references), except under strictly bounded non-requirement clauses.

#### + Scenario: Living files reject ticket, decision-marker, and bare section-number references
+ - GIVEN any file under `.cooper/definition/` or `.cooper/specs/` at the repository root or inside a barrel
+ - WHEN requirement text, GIVEN/WHEN/THEN scenario steps, or general living prose references related behavior, rationale, or structure
+ - THEN the file MUST NOT include ticket, issue, PR, or track numbers
+ - AND the file MUST NOT include decision or open-question markers (`D14`, `DEC-1`, `OQ-4`)
+ - AND the file MUST NOT use bare section numbers (including self-references), using exact heading names instead
+ - AND files under `.cooper/active/` and `.cooper/archive/` remain exempt as point-in-time working records.

#### + Scenario: Rationale links outside requirement text cite frozen archived tracks or approved RFCs by file path and heading name
+ - GIVEN prose outside requirement text (outside `### Requirement:` descriptions and `GIVEN`/`WHEN`/`THEN` scenario steps) in a living definition or capability spec
+ - WHEN explaining historical design rationale anchored in an archived track document, an approved or archived RFC, or a frozen macro contract
+ - THEN the rationale note MAY link that frozen artifact using its repository file path paired with the exact heading name
+ - AND it MUST NOT replace the heading name with a bare section number, ticket number, or decision marker.

#### + Scenario: Open known-gap entries outside requirement text cite tracking tickets until resolved
+ - GIVEN prose outside requirement text in a living definition or capability spec documenting an active known gap
+ - WHEN the referenced tracking ticket has not yet shipped
+ - THEN the known-gap entry MAY cite that tracking ticket so readers know the limitation is tracked
+ - AND once the tracking ticket ships, the known-gap entry and its ticket citation MUST be removed.

#### + Scenario: Spec delta finalization strips planning provenance before merging into living specs
+ - GIVEN a completed track whose Spec Deltas in `.cooper/active/<track_id>/spec-deltas/` are being integrated into `.cooper/specs/<capability>/spec.md`
+ - WHEN merging requirement changes and updating living files during Track Finalization
+ - THEN all ticket numbers, decision/open-question markers, and bare section numbers MUST be stripped or rewritten to exact heading names before committing the living spec.

#### + Scenario: Brownfield audit flags forbidden references in existing living files
+ - GIVEN an existing repository audited by `cooper-setup` in brownfield discovery mode
+ - WHEN scanning `.cooper/definition/` and `.cooper/specs/` at the root and in every barrel
+ - THEN `cooper-setup` MUST identify and report existing ticket numbers, decision/open-question markers, and bare section numbers for remediation.
