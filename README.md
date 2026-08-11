# Cooper 🛢️🐒

**Cooper** combines the **Conductor** ([conductor](https://github.com/gemini-cli-extensions/conductor)) methodology with worktree isolation via **Troop** ([troop](https://github.com/twoBoots/troop)) for crafting a barrel of monkeys (i.e. human and AI developers working in trees).

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
2. **Conductor Specification**: Copies `conductor/workflow.md` into your project's `conductor/` directory.
3. **Agent Rules**: Injects combined Conductor + Troop rules from `AGENTS.template.md` into your project's `AGENTS.md`.

## Structure

- `install.sh`: One-line automated installer script.
- `conductor/workflow.md`: The combined Conductor + Troop workflow specification.
- `AGENTS.template.md`: Template rules injected into `AGENTS.md` for AI agent instructions.

## Workflow Summary

1. **Spawn Track Worktree**: `git agent-start <track_id>`
   - Creates branch `<track_id>` and checks out isolated worktree at `.worktrees/<track_id>`.
2. **Develop in Worktree**: Navigate to `.worktrees/<track_id>` to work.
   - Follow TDD, maintain high coverage, record plan updates, and create phase checkpoints following `conductor/workflow.md`.
3. **List Active Worktrees**: `git troop`
4. **Push & PR**: Push `<track_id>` and submit PR.
5. **Teardown Worktree**: `git agent-stop <track_id>` after PR approval and merge.
