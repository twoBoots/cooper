# Spec Delta: Living Files Stand Alone

Target Spec: `.cooper/specs/documentation/spec.md`

## Added Requirements

### + Requirement: Living Files Stand Alone
+ Files in `.cooper/definition/` and `.cooper/specs/` (at the repo root and in every barrel) SHALL stand alone. They MUST NOT contain ticket, issue, PR, or track numbers; decision or open-question markers (numbered `D`, `DEC`, or `OQ` identifiers); or bare section numbers (section-symbol numbers or positional section references, including self-references), except under the two non-requirement exceptions below.

#### + Scenario: Living files reject ticket, decision-marker, and bare section-number references
+ - GIVEN a file in `.cooper/definition/` or `.cooper/specs/` at the repo root or in a barrel
+ - WHEN requirement text, GIVEN/WHEN/THEN scenario steps, or living prose is authored
+ - THEN it MUST NOT include ticket, issue, PR, or track numbers
+ - AND it MUST NOT include decision or open-question markers (numbered `D`, `DEC`, or `OQ` identifiers)
+ - AND it MUST NOT use bare section numbers (including self-references), using exact heading names instead
+ - AND `.cooper/active/` and `.cooper/archive/` remain exempt.

#### + Scenario: Rationale links outside requirement text cite frozen artifacts by file path and heading name
+ - GIVEN prose outside requirement text and `GIVEN`/`WHEN`/`THEN` scenario steps in `.cooper/definition/` or `.cooper/specs/`
+ - WHEN referencing design rationale in an archived track, an approved or archived RFC, or a frozen macro contract
+ - THEN the note MAY link that artifact by repository file path and exact heading name
+ - AND it MUST NOT use bare section numbers, ticket numbers, or decision markers.

#### + Scenario: Open known-gap entries outside requirement text cite tracking tickets until resolved
+ - GIVEN prose outside requirement text in `.cooper/definition/` or `.cooper/specs/` documenting an open known gap
+ - WHEN the tracking ticket has not shipped
+ - THEN the entry MAY cite that tracking ticket
+ - AND once the ticket ships, the entry MUST be removed.

#### + Scenario: Spec delta finalization strips planning provenance before merging into living specs
+ - GIVEN a completed track merging `.cooper/active/<track_id>/spec-deltas/` into `.cooper/specs/<capability>/spec.md`
+ - WHEN integrating requirement changes during Track Finalization
+ - THEN ticket numbers, decision or open-question markers, and bare section numbers MUST be stripped or replaced with exact heading names before commit.

#### + Scenario: Brownfield audit flags forbidden references in existing living files
+ - GIVEN a repository audited by `cooper-setup` in brownfield mode
+ - WHEN scanning `.cooper/definition/` and `.cooper/specs/` at the root and in every barrel
+ - THEN `cooper-setup` MUST flag ticket numbers, decision or open-question markers, and bare section numbers for cleanup.
