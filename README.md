# Wrangler 🐒

**Wrangler** combines the **Conductor** methodology with worktree isolation via **JungleJim** ([twoBoots/junglejim](https://github.com/twoBoots/junglejim)) for wrangling code monkeys (i.e. human and AI developers).

## Quick Start / One-Line Installation

To set up **Wrangler** in any Git project repository, navigate to your target project folder and run:

```bash
curl -fsSL https://raw.githubusercontent.com/twoBoots/wrangler/main/install.sh | bash
```

Alternatively, if you have this repository cloned locally:

```bash
/path/to/wrangler/install.sh /path/to/your-project
```

### What `install.sh` does:

1. **JungleJim Setup**: Runs the `junglejim` installer to set up shared Git aliases (`git agent-start`, `git jims`, `git agent-stop`), `.gitignore` (`.worktrees/`), and installs `JUNGLEJIM.md`.
2. **Conductor Specification**: Copies `conductor/workflow.md` into your project's `conductor/` directory.
3. **Agent Rules**: Injects combined Conductor + JungleJim rules from `AGENTS.template.md` into your project's `AGENTS.md`.

## Structure

- `install.sh`: One-line automated installer script.
- `conductor/workflow.md`: The combined Conductor + JungleJim workflow specification.
- `AGENTS.template.md`: Template rules injected into `AGENTS.md` for AI agent instructions.

## Workflow Summary

1. **Spawn Track Worktree**: `git agent-start <track_id>`
   - Creates branch `feature/<track_id>` and checks out isolated worktree at `.worktrees/<track_id>`.
2. **Develop in Worktree**: Navigate to `.worktrees/<track_id>` to work.
   - Follow TDD, maintain high coverage, record plan updates, and create phase checkpoints following `conductor/workflow.md`.
3. **List Active Worktrees**: `git jims`
4. **Push & PR**: Push `feature/<track_id>` and submit PR.
5. **Teardown Worktree**: `git agent-stop <track_id>` after PR approval and merge.
