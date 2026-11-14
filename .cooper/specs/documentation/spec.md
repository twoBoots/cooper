# Capability Specification: Framework Documentation & Guidelines

## Purpose & Scope
Defines the standards and guidelines for project-level documentation, framework reference manuals, agent prompts, and templates within the Cooper ecosystem.

## Requirements

### Requirement: External Dependency Attribution & Repository Linking
All framework documentation, manuals, skill instructions, and templates referencing foundational external tools SHALL link directly to their upstream source repositories on primary introduction.

#### Scenario: Primary Mention of Troop in Markdown Documents
- GIVEN a Cooper markdown documentation file (`README.md`, `AGENTS.md`, `docs/*.md`, `.cooper/*.md`, `skills/*/*.md`, or `templates/*.md`)
- WHEN the document mentions the Troop worktree isolation tool
- THEN the first prominent mention of Troop MUST be formatted as a markdown link pointing to `https://github.com/twoBoots/troop` (e.g. `[Troop](https://github.com/twoBoots/troop)`).

#### Scenario: Tool Summary Tables & Capability Listings
- GIVEN a table, overview matrix, or capability list referencing Cooper foundation components
- WHEN Troop is listed as a tool or pillar
- THEN the entry MUST include a direct markdown link to `https://github.com/twoBoots/troop`.

### Requirement: Interactive Question & File Tool Protocols
When interacting with users or performing file operations during skill execution, agents MUST prioritize runtime interactive tools and native file tools over plain text or shell stream redirections.

#### Scenario: Interactive Question Tool Invocations
- GIVEN an agent executing a Cooper skill or workflow step requiring user input, selection, or confirmation
- AND the agent runtime provides an interactive question tool (such as `ask_question`)
- WHEN the agent presents single-choice or multiple-choice options
- THEN the agent MUST invoke the interactive question tool
- AND the agent MUST NOT output plain text numbered/bulleted option lists in chat.

#### Scenario: Fallback to Text Chat for Questions
- GIVEN an agent executing a Cooper skill or workflow step
- AND no interactive question tool is provided in the agent runtime environment
- WHEN the agent needs to ask questions or present choices
- THEN the agent SHALL format the question clearly in text chat, asking questions strictly one at a time.

#### Scenario: Native File Tools Enforcement
- GIVEN an agent runtime providing native file tools (`view_file`, `write_to_file`, `replace_file_content`)
- WHEN the agent needs to view, create, or edit files in the workspace
- THEN the agent MUST use native file tools
- AND the agent MUST NOT use shell commands with stream editors (`sed`, `awk`), heredocs, pipes, or stream redirections (`cat << 'EOF'`, `echo >`, `cat >`) to create or edit files.

### Requirement: Living Files Stand Alone
Files in `.cooper/definition/` and `.cooper/specs/` (at the repo root and in every barrel) SHALL stand alone. They MUST NOT contain ticket, issue, PR, or track numbers; decision or open-question markers (numbered `D`, `DEC`, or `OQ` identifiers); or bare section numbers (section-symbol numbers or positional section references, including self-references), except under the two non-requirement exceptions below.

#### Scenario: Living files reject ticket, decision-marker, and bare section-number references
- GIVEN a file in `.cooper/definition/` or `.cooper/specs/` at the repo root or in a barrel
- WHEN requirement text, GIVEN/WHEN/THEN scenario steps, or living prose is authored
- THEN it MUST NOT include ticket, issue, PR, or track numbers
- AND it MUST NOT include decision or open-question markers (numbered `D`, `DEC`, or `OQ` identifiers)
- AND it MUST NOT use bare section numbers (including self-references), using exact heading names instead
- AND `.cooper/active/` and `.cooper/archive/` remain exempt.

#### Scenario: Rationale links outside requirement text cite frozen artifacts by file path and heading name
- GIVEN prose outside requirement text and `GIVEN`/`WHEN`/`THEN` scenario steps in `.cooper/definition/` or `.cooper/specs/`
- WHEN referencing design rationale in an archived track, an approved or archived RFC, or a frozen macro contract
- THEN the note MAY link that artifact by repository file path and exact heading name
- AND it MUST NOT use bare section numbers, ticket numbers, or decision markers.

#### Scenario: Open known-gap entries outside requirement text cite tracking tickets until resolved
- GIVEN prose outside requirement text in `.cooper/definition/` or `.cooper/specs/` documenting an open known gap
- WHEN the tracking ticket has not shipped
- THEN the entry MAY cite that tracking ticket
- AND once the ticket ships, the entry MUST be removed.

#### Scenario: Spec delta finalization strips planning provenance before merging into living specs
- GIVEN a completed track merging `.cooper/active/<track_id>/spec-deltas/` into `.cooper/specs/<capability>/spec.md`
- WHEN integrating requirement changes during Track Finalization
- THEN ticket numbers, decision or open-question markers, and bare section numbers MUST be stripped or replaced with exact heading names before commit.

#### Scenario: Brownfield audit flags forbidden references in existing living files
- GIVEN a repository audited by `cooper-setup` in brownfield mode
- WHEN scanning `.cooper/definition/` and `.cooper/specs/` at the root and in every barrel
- THEN `cooper-setup` MUST flag ticket numbers, decision or open-question markers, and bare section numbers for cleanup.
