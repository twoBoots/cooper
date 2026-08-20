# Cooper 🛢️🐒

**Cooper** combines **OpenSpec's Living Spec Deltas** with **Conductor's** quality governance and **[Troop's](https://github.com/twoBoots/troop)** worktree isolation into a unified, agent-agnostic **Spec-Driven Development (SDD)** framework under `.cooper/`.

Cooper packages its own dedicated agent skills under `.agents/skills/`, making it 100% self-contained and free of external plugin prerequisites.

## Quick Start / One-Line Installation

To set up **Cooper** in any Git project repository, navigate to your target project folder and run:

```bash
curl -fsSL https://raw.githubusercontent.com/twoBoots/cooper/main/install.sh | bash
```

Alternatively, if you have this repository cloned locally:

```bash
/path/to/cooper/install.sh /path/to/your-project
```

> 📖 **Detailed Installer & Migration Guide**: For details on auto-migrating existing Conductor/OpenSpec setups or greenfield scaffolding, see [`docs/INSTALL.md`](docs/INSTALL.md).

## What `install.sh` Does

1. **[Troop](https://github.com/twoBoots/troop) Foundation**: Sets up shared Git aliases (`git agent-start`, `git troop`, `git agent-stop`), `.gitignore` (`.worktrees/`), and installs `TROOP.md` into `.cooper/TROOP.md`.
2. **Auto-Migration & Scaffolding**: Auto-migrates existing `conductor/` or `openspec/` setups to `.cooper/`, or scaffolds baseline templates from Cooper's self-contained template suite for greenfield repositories.
3. **Cooper Specification & Handshake**: Copies workflow specifications into `.cooper/definition/workflow.md`, installs `.cooper/COOPER.md`, and creates `.cooper/index.md` as the single source of truth.
4. **Native Agent Skills**: Installs project-local Cooper skills into `.agents/skills/` (`cooper-setup`, `cooper-rfc`, `cooper-new-track`, `cooper-implement`, `cooper-review`, `cooper-status`).
5. **Agent Rules**: Injects Cooper SDD + [Troop](https://github.com/twoBoots/troop) rules from `AGENTS.template.md` into your project's `AGENTS.md`.

## Structure

```
your-project/
├── .agents/
│   └── skills/                        # Packaged Project-Local Cooper Skills
│       ├── cooper-setup/SKILL.md      # Project initialization & scaffolding
│       ├── cooper-rfc/SKILL.md        # Upstream collaborative RFC & architectural design
│       ├── cooper-new-track/SKILL.md  # Worktree spawning & spec delta planning
│       ├── cooper-implement/SKILL.md  # TDD execution & phase sync
│       ├── cooper-review/SKILL.md     # Code & spec delta review
│       └── cooper-status/SKILL.md     # Worktrees & track overview
├── .cooper/
│   ├── index.md                       # Handshake index (Single Source of Truth)
│   ├── COOPER.md                      # Cooper SDD reference manual & cheatsheet
│   ├── TROOP.md                       # Troop worktree reference manual
│   ├── tracks.md                      # Tracks Registry
│   ├── definition/                    # Global project definitions (product.md, tech-stack.md, workflow.md)
│   ├── code_styleguides/              # Language-specific conventions
│   ├── specs/                         # LIVING CAPABILITY SPECS (OpenSpec Living Spec model)
│   ├── active/                        # ACTIVE TRACKS (Living inside .worktrees/<track_id>/)
│   │   └── <track_id>/
│   │       ├── proposal.md            # High-level rationale
│   │       ├── design.md              # Technical design decisions
│   │       ├── plan.md                # TDD-enforced, phase-checkpointed plan
│   │       ├── metadata.json          # Track status metadata
│   │       └── spec-deltas/           # Requirement diffs (+ added, - removed)
│   └── archive/                       # HISTORICAL COMPLETED TRACKS
├── .worktrees/                        # Isolated Git worktrees for active tracks
└── AGENTS.md                          # Universal agent guidelines
```

## Available Agent Skills (`.agents/skills/`)

| Skill | Description |
| :--- | :--- |
| **`cooper-setup`** | Audits codebase, scaffolds `.cooper/` definitions, styleguides, and configures [Troop](https://github.com/twoBoots/troop) worktrees. |
| **`cooper-rfc`** | Plans collaborative architectural initiatives, drafts RFCs & spec deltas, opens Draft PRs, and decomposes into tracks ([RFC Draft PR Guide](docs/rfc-draft-prs.md)). |
| **`cooper-new-track`** | Spawns `.worktrees/<track_id>`, inspects living specs, and drafts proposal, design, spec deltas, and plan. |
| **`cooper-implement`** | Executes TDD loop inside worktree, records Git Notes metadata, and runs phase synchronization & checkpoints. |
| **`cooper-review`** | Conducts Principal Software Engineer code review against spec deltas, styleguides, and test suites. |
| **`cooper-status`** | Displays real-time overview of active worktrees, track progress, and phase checkpoints. |

## Workflow Summary

1. **Spawn Track Worktree**: `git agent-start <track_id>` (or invoke `cooper-new-track`)
   - Creates branch `<track_id>` and checks out isolated worktree at `.worktrees/<track_id>`.
2. **Develop in Worktree**: Navigate to `.worktrees/<track_id>` to work (or invoke `cooper-implement`).
   - Generate Spec Deltas (`.cooper/active/<track_id>/spec-deltas/`) and TDD plan (`plan.md`).
   - Follow TDD (Red -> Green -> Refactor), attach Git Notes commit summaries, execute phase synchronization (`git fetch origin main`), and create phase checkpoints following `.cooper/definition/workflow.md`.
3. **List Active Worktrees**: `git troop` (or invoke `cooper-status`).
4. **Review & PR**: Run `cooper-review`, push `<track_id>`, and submit PR (`gh pr create --body-file prbody.md`).
5. **Merge & Teardown**: Upon PR merge, Spec Deltas merge into `.cooper/specs/`, and teardown worktree via `git agent-stop <track_id>`.
