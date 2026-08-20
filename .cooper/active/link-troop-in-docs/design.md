# Technical Design: Link Troop Mentions in Markdown Documentation to twoBoots/troop

## Overview
This design details the auditing, formatting standards, and file-by-file update strategy for linking Troop mentions across Cooper's markdown documentation to the `twoBoots/troop` GitHub repository.

## Target Link Specification
- **Target URL**: `https://github.com/twoBoots/troop`
- **Standard Link Text**: `[Troop](https://github.com/twoBoots/troop)`
- **Header / Explicit Intro Style**: `[Troop (twoBoots/troop)](https://github.com/twoBoots/troop)` or `[Troop](https://github.com/twoBoots/troop)`
- **Table References**: `| [Troop](https://github.com/twoBoots/troop) | ... |`

## Audit & Categorization by File Group

### 1. Root & Agent Configuration Files
- **`README.md`**: Update introductory overview, architecture breakdown, and workflow steps to link the first prominent mention of Troop.
- **`AGENTS.md` & `AGENTS.template.md`**: Update Section 2 ("Troop Worktree Isolation Protocol") heading and first mention.
- **`CLAUDE.md`**: Ensure repository references or guidelines link properly if Troop is mentioned.

### 2. General Documentation (`docs/`)
- **`docs/INSTALL.md`**: Update Troop foundation setup step.
- **`docs/openspec-vs-conductor-comparison.md`**: Update section 3.1 Troop Worktree Isolation heading and architecture summary.
- **`docs/rfc-draft-prs.md`**: Update architecture diagram notes and child track execution references.

### 3. Cooper Framework Definitions (`.cooper/`)
- **`.cooper/COOPER.md`**: Update introduction and capability comparison table.
- **`.cooper/definition/workflow.md`**: Update Section 2 (Workspace Structure) and Section 3 (Track Worktree Protocol) initial mentions.

### 4. Agent Skills (`skills/` and `.agents/skills/`)
- Audit and update `cooper-setup`, `cooper-rfc`, `cooper-new-track`, `cooper-implement`, `cooper-review`, and `cooper-status` in both `skills/` and `.agents/skills/` to ensure the first prominent mention links to `https://github.com/twoBoots/troop`.

### 5. Starter Templates (`templates/`)
- Update `templates/product.md`, `templates/tech-stack.md`, and any template documentation referencing Troop.

## Automated Verification Strategy
1. **Link Verification Script / Regex Check**:
   - Run an automated grep / audit to verify every active markdown document containing `Troop` has at least one valid markdown link `[Troop](https://github.com/twoBoots/troop)` or `[twoBoots/troop](https://github.com/twoBoots/troop)`.
2. **Exclusion Check**:
   - Confirm `.cooper/archive/` remains untouched.
