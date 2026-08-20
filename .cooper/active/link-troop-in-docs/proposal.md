# Track Proposal: Link Troop Mentions in Markdown Documentation to twoBoots/troop

## Summary
Update all Cooper markdown documentation files to link prominent mentions of **Troop** directly to its official GitHub repository ([twoBoots/troop](https://github.com/twoBoots/troop)).

## Problem Statement
Cooper is built upon a hybrid SDD architecture combining OpenSpec's living spec deltas, Conductor's quality gates, and Troop's Git worktree isolation. While Cooper frequently references Troop for worktree isolation workflows (`git agent-start`, `git troop`, `git agent-stop`), many markdown documentation files reference "Troop" as plain text or lack direct outbound links to the upstream repository. This makes it harder for new users, developers, and agents to discover Troop's standalone documentation and repository.

## Proposed Solution
Audit and update all markdown documentation files in Cooper (`README.md`, `AGENTS.md`, `AGENTS.template.md`, `CLAUDE.md`, `docs/`, `.cooper/`, `skills/`, `.agents/skills/`, and `templates/`):
- Link the primary/first prominent mention of Troop in each markdown document to `https://github.com/twoBoots/troop` using `[Troop](https://github.com/twoBoots/troop)`.
- Link Troop in summary tables and architecture overviews where tools/capabilities are listed.
- Preserve clean readability without excessive link repetition in dense command snippets.

## Scope Boundaries
- **In Scope:**
  - `README.md`
  - `AGENTS.md` and `AGENTS.template.md`
  - `CLAUDE.md` (if applicable)
  - `docs/*.md` (`INSTALL.md`, `openspec-vs-conductor-comparison.md`, `rfc-draft-prs.md`)
  - `.cooper/COOPER.md` and `.cooper/definition/workflow.md`
  - Project skills under `skills/` and `.agents/skills/` (`cooper-setup`, `cooper-rfc`, `cooper-new-track`, `cooper-implement`, `cooper-review`, `cooper-status`)
  - Starter templates under `templates/`
- **Out of Scope:**
  - Historical archived tracks in `.cooper/archive/` (to preserve historical record integrity).
  - Changing functional script code or CLI command names.
