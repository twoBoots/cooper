# Track Design: Scaffold Claude Bridge (`.claude/skills` and `CLAUDE.md`)

## Architectural Overview
Cooper maintains an agent-agnostic core under `.agents/skills/` and `AGENTS.md`. To provide immediate compatibility with Claude-oriented agent interfaces without fragmenting skill definitions or guidelines, Cooper introduces a shim layer:

```
[Repository Root]
├── AGENTS.md                  <-- Single source of truth for agent guidelines
├── CLAUDE.md                  <-- Shim containing '@AGENTS.md'
├── .agents/
│   └── skills/                <-- Single source of truth for skill directories
│       ├── cooper-setup/
│       ├── cooper-new-track/
│       └── ...
└── .claude/
    └── skills                 <-- Symbolic link: '../.agents/skills'
```

## Detailed Mechanics

### 1. Relative Symlink Creation
- Target: `../.agents/skills`
- Destination: `.claude/skills`
- Using a relative path (`../.agents/skills` rather than an absolute path) ensures the symlink remains valid across worktrees, branch checkouts, and when cloning or moving the repository to different paths or filesystems.
- Guard: `[ -e .claude/skills ] || ln -s ../.agents/skills .claude/skills`
  - The `-e` test verifies whether any file, directory, or link already exists at `.claude/skills`. If so, it leaves the existing path intact.

### 2. CLAUDE.md Injection & Preservation
- If `CLAUDE.md` does not exist:
  Create `CLAUDE.md` containing `@AGENTS.md\n`.
- If `CLAUDE.md` already exists:
  Check whether `@AGENTS.md` is present via `grep -qs "@AGENTS.md" CLAUDE.md`.
  - If absent, append `\n@AGENTS.md\n`.
  - If present, leave file untouched to prevent duplicate lines on re-runs.

### 3. Function Encapsulation in `install.sh`
- A dedicated helper function `setup_claude_bridge` will be introduced in `install.sh`.
- Executed immediately following step 5 (`Setting up AGENTS.md rules...`).
- When `COOPER_INSTALL_LIB_ONLY=1` is exported during test execution, `setup_claude_bridge` is directly callable for targeted Go unit tests.

### 4. Integration with `cooper-setup` Skill
- Section 4 ("Agent Guidelines (AGENTS.md)") and Section 2.8 ("Project-Local Agent Skills Installation (.agents/skills/)") of `skills/cooper-setup/SKILL.md` will document creating the `.claude/skills` symlink and `@AGENTS.md` bridge.
- Both `skills/cooper-setup/SKILL.md` and `.agents/skills/cooper-setup/SKILL.md` will be updated synchronously to preserve byte-level identity.
