# Cooper 🛢️🐒

[![CI](https://github.com/twoBoots/cooper/actions/workflows/ci.yml/badge.svg)](https://github.com/twoBoots/cooper/actions/workflows/ci.yml)
[![Release](https://github.com/twoBoots/cooper/actions/workflows/release.yml/badge.svg)](https://github.com/twoBoots/cooper/releases)

**Cooper** combines **OpenSpec's Living Spec Deltas** with **Conductor's** quality governance and **[Troop's](https://github.com/twoBoots/troop)** worktree isolation into a unified, agent-agnostic **Spec-Driven Development (SDD)** framework under `.cooper/`.

Cooper packages its own dedicated agent skills under `.agents/skills/`, making it 100% self-contained and free of external plugin prerequisites.

## Overview

Cooper is an agent-agnostic Spec-Driven Development (SDD) framework and CLI. It structures AI-assisted engineering into explicit specifications, isolated Git workspaces, and strict quality controls.

### Core Architecture

* **Living Capability Specs (`.cooper/specs/`)**: System behaviour is documented as living specifications. Changes are defined as human-reviewable Spec Deltas (`+` additions, `-` removals) before code is written.
* **Worktree Isolation (Troop)**: Work executes in dedicated worktrees (`.worktrees/<track_id>`), eliminating branch switching, stash conflicts, and dirty working trees across parallel agents.
* **Quality & Phase Gates**: Enforces test-first implementation, test coverage checks, and remote checkpoint syncs per phase.
* **Self-Contained Agent Skills**: Ships project-local skills (`.agents/skills/cooper-*`) plus an optional Go CLI for deterministic spec validation. Zero external plugins required, and no binary needed to use the framework.

### Value Delivered

| Operational Area | Without Cooper | With Cooper |
| :--- | :--- | :--- |
| **Architecture** | Code starts before consensus; architectural debate happens in PR code diffs. | Upstream Draft RFC PRs isolate design and spec deltas until team sign-off. |
| **Specifications** | Prompt drift across chat sessions; stale documentation. | Living specs (`.cooper/specs/`) updated automatically via Spec Deltas on PR merge. |
| **Workspaces** | Stash conflicts and branch collisions across concurrent agent runs. | Isolated Git worktrees per track via Troop (`.worktrees/<track_id>`). |
| **Quality** | Bypassed test suites or unverified generation. | Strict TDD (Red -> Green -> Refactor), >80% coverage check, and Git Notes audit trails. |

> 📖 **Deep Dive Documentation**:
> - [OpenSpec vs. Conductor Comparison & Cooper Architecture Blueprint](docs/openspec-vs-conductor-comparison.md)
> - [Why Cooper Architectural RFCs Mandate Draft Pull Requests](docs/rfc-draft-prs.md)

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
6. **Optional CLI Binary**: Compiles from a local clone when Go is available, otherwise downloads the release asset for your platform (with macOS quarantine stripping and ad-hoc signing). If neither succeeds — air-gapped, rate-limited, or no writable `bin` directory — the installer reports it and completes normally in zero-binary mode.

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

## 🚀 The `cooper` CLI

Cooper's framework is markdown, and your agent is the runtime. The optional Go binary exists for the one job an agent cannot reliably do on its own output: **deterministic specification checking**.

`install.sh` installs it when it can, and completes successfully when it cannot — the `.cooper/` workspace and agent skills work without it.

```bash
# Validate living capability specs, spec deltas, track metadata,
# registry parity, and documentation references
cooper validate
cooper validate --json          # machine-readable report for CI

# In-place binary self-updating from GitHub Releases
cooper update
cooper update --check
cooper update --force

cooper version
cooper --help
```

### Why validation is the whole CLI

`cooper validate` is fast, exhaustive, and costs no tokens. It catches malformed `GIVEN`/`WHEN`/`THEN` scenarios, missing normative keywords, track-ID and directory mismatches, invalid metadata, registry drift, and documentation references that point at files which do not exist. Crucially, it does this *without trusting the agent that wrote the spec*.

Wire it into your own CI when you want that enforcement:

```yaml
- name: Validate Cooper SDD Specs
  run: cooper validate
```

Cooper does not install this gate for you. Enforcement is yours to opt into.

### Scaffolding and track management

These are handled by `install.sh` and the agent skills respectively, not by the binary:

| Task | Use |
| :--- | :--- |
| Scaffold or migrate a repository | `install.sh` |
| Create a track and its worktree | `cooper-new-track` skill, `git agent-start` |
| Execute tasks, checkpoint a phase | `cooper-implement` skill |
| Inspect tracks and worktrees | `cooper-status` skill, `git troop` |

> **Removed in v1.2.0**: `cooper init`, `cooper track *`, `cooper mcp`, and `cooper mcp install`.
> Scaffolding moved to `install.sh` (a single source of truth — the two scaffolders had diverged), and track orchestration belongs to the skills that already performed it.
> If you previously ran `cooper mcp install`, remove the `cooper` entry from your editor's MCP config (`.cursor/mcp.json`, `~/.claude.json`, `~/.codeium/windsurf/mcp_config.json`, `.vscode/mcp.json`, or `~/.gemini/config/mcp_config.json`); it now points at a command that no longer exists.

## Workflow Summary

1. **Spawn Track Worktree**: `git agent-start <track_id>` (or invoke `cooper-new-track`)
   - Creates branch `<track_id>` and checks out isolated worktree at `.worktrees/<track_id>`.
2. **Develop in Worktree**: Navigate to `.worktrees/<track_id>` to work (or invoke `cooper-implement`).
   - Generate Spec Deltas (`.cooper/active/<track_id>/spec-deltas/`) and TDD plan (`plan.md`).
   - Follow TDD (Red -> Green -> Refactor), attach Git Notes commit summaries, execute phase synchronization (`git fetch origin main`), and create phase checkpoints following `.cooper/definition/workflow.md`.
3. **List Active Worktrees**: `git troop` (or invoke `cooper-status`).
4. **Review & PR**: Run `cooper-review`, push `<track_id>`, and submit PR (`gh pr create --body-file prbody.md`).
5. **Merge & Teardown**: Upon PR merge, Spec Deltas merge into `.cooper/specs/`, and teardown worktree via `git agent-stop <track_id>`.
