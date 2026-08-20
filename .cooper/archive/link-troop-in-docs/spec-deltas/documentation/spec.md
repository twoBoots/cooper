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
