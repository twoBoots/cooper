# Cooper 🛢️🐒

**Cooper** combines **OpenSpec's Living Spec Deltas** with **Conductor's** quality governance and **Troop's** worktree isolation ([troop](https://github.com/twoBoots/troop)) into a unified **Spec-Driven Development (SDD)** framework under `.cooper/`.

## Quick Start / One-Line Installation

To set up **Cooper** in any Git project repository, navigate to your target project folder and run:

```bash
curl -fsSL https://raw.githubusercontent.com/twoBoots/cooper/main/install.sh | bash
```

Alternatively, if you have this repository cloned locally:

```bash
/path/to/cooper/install.sh /path/to/your-project
```

### What `install.sh` does:

1. **Troop Setup**: Runs the `troop` installer to set up shared Git aliases (`git agent-start`, `git troop`, `git agent-stop`), `.gitignore` (`.worktrees/`), and installs `TROOP.md`.
2. **Cooper Hybrid Specification**: Copies workflow specifications into `.cooper/definition/workflow.md` (and `conductor/workflow.md`).
3. **Agent Rules**: Injects Cooper Hybrid + Troop rules from `AGENTS.template.md` into your project's `AGENTS.md`.

## Structure

```
.cooper/
├── definition/                    # Global project definitions (product.md, tech-stack.md, workflow.md)
├── code_styleguides/              # Language-specific conventions
├── specs/                         # LIVING CAPABILITY SPECS (OpenSpec Living Spec model)
├── active/                        # ACTIVE TRACKS (Living inside .worktrees/<track_id>/)
│   └── <track_id>/
│       ├── proposal.md            # High-level rationale
│       ├── design.md              # Technical design decisions
│       ├── plan.md                # TDD-enforced, phase-checkpointed plan
│       ├── metadata.json          # Track status metadata
│       └── spec-deltas/           # Requirement diffs (+ added, - removed)
└── archive/                       # HISTORICAL COMPLETED TRACKS
```

- `install.sh`: One-line automated installer script.
- `cooper/workflow.md`: The Cooper Hybrid SDD + Troop workflow specification.
- `AGENTS.template.md`: Template rules injected into `AGENTS.md` for AI agent instructions.
- `docs/openspec-vs-conductor-comparison.md`: Detailed comparison and blueprint.

## Workflow Summary

1. **Spawn Track Worktree**: `git agent-start <track_id>`
   - Creates branch `<track_id>` and checks out isolated worktree at `.worktrees/<track_id>`.
2. **Develop in Worktree**: Navigate to `.worktrees/<track_id>` to work.
   - Generate Spec Deltas (`.cooper/active/<track_id>/spec-deltas/`) and TDD plan (`plan.md`).
   - Follow TDD (Red -> Green -> Refactor), attach Git Notes commit summaries, execute phase synchronization (`git fetch origin main`), and create phase checkpoints following `.cooper/definition/workflow.md`.
3. **List Active Worktrees**: `git troop`
4. **Push & PR**: Push `<track_id>` and submit PR (`gh pr create --body-file prbody.md`).
5. **Merge & Teardown**: Upon PR merge, Spec Deltas merge into `.cooper/specs/`, and teardown worktree via `git agent-stop <track_id>`.
