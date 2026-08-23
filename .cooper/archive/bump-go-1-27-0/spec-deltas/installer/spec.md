# Capability Specification: Installer & Repo Structure

## Purpose & Scope
Defines the installation, scaffolding, and repository directory structure requirements for the Cooper Hybrid Framework.

## Requirements

### Requirement: Troop Reference Relocation
The installer SHALL place all [Troop](https://github.com/twoBoots/troop) reference material within the `.cooper/` directory to keep the root directory uncluttered.

#### Scenario: Relocate TROOP.md on Setup
- GIVEN Cooper installer executes in a project
- WHEN the Troop foundation setup completes
- THEN `TROOP.md` MUST be placed or moved into `.cooper/TROOP.md` instead of remaining at the root level.

### Requirement: Workflow Specification Fetching
The installer SHALL scaffold the Cooper workflow specification into `.cooper/definition/workflow.md`.

#### Scenario: Resolve Cooper Definition Workflow
- GIVEN Cooper installer fetches workflow specification
- WHEN resolving Cooper repository assets
- THEN `.cooper/definition/workflow.md` MUST be fetched and placed into `.cooper/definition/workflow.md`.

+ ### Requirement: Scaffold Go Runtime Baseline
+ The installer and scaffolding templates SHALL provide Go 1.27+ as the default Go runtime baseline in scaffolded tech stack definitions.
+ 
+ #### Scenario: Scaffolding Default Tech Stack Template
+ - GIVEN a user scaffolding a new Cooper project
+ - WHEN `.cooper/definition/tech-stack.md` is populated from scaffold templates
+ - THEN the Go runtime engine entry MUST specify `Go 1.27+`.
