# Cooper Installation & Migration Guide

`install.sh` is a one-line installer that scaffolds the **Cooper Spec-Driven Development (SDD) Framework** (`.cooper/`), **[Troop](https://twoboots.github.io/troop) Worktree Isolation** (`.worktrees/`), and native **Project Agent Skills** (`.agents/skills/`) into any target Git repository.

## Quick Installation

Run the following command inside your target repository:

```bash
curl -fsSL https://raw.githubusercontent.com/twoBoots/cooper/main/install.sh | bash
```

After installing, ask your agent to run `/cooper-setup` inside your project.

Alternatively, if running from a local clone of the Cooper repository:

```bash
/path/to/cooper/install.sh /path/to/your-project
```

## Installer Execution Flow & Migration Logic

<div class="installer-flow">
  <div class="flow-card">
    <div class="flow-badge">Step 1</div>
    <div class="flow-title">Run install.sh</div>
    <div class="flow-desc">Execute installer in target repository</div>
  </div>
  <div class="flow-arrow">↓</div>
  <div class="flow-card flow-audit">
    <div class="flow-badge">Step 2: Repository Audit</div>
    <div class="branch-grid">
      <div class="branch-item">
        <span class="branch-tag">Conductor</span>
        <div class="branch-text">Auto-migrates <code>conductor/</code> to <code>.cooper/</code></div>
      </div>
      <div class="branch-item">
        <span class="branch-tag">OpenSpec</span>
        <div class="branch-text">Auto-migrates <code>openspec/</code> to <code>.cooper/specs/</code></div>
      </div>
      <div class="branch-item">
        <span class="branch-tag">Greenfield</span>
        <div class="branch-text">Scaffolds baseline templates &amp; specs</div>
      </div>
    </div>
  </div>
  <div class="flow-arrow">↓</div>
  <div class="flow-card">
    <div class="flow-badge">Step 3</div>
    <div class="flow-title">Scaffold .cooper/ Structure</div>
    <div class="flow-desc">Initializes definitions, specs, tracks, and <code>COOPER.md</code></div>
  </div>
  <div class="flow-arrow">↓</div>
  <div class="flow-card">
    <div class="flow-badge">Step 4</div>
    <div class="flow-title">Install Agent Skills</div>
    <div class="flow-desc">Populates project-local <code>.agents/skills/cooper-*</code></div>
  </div>
  <div class="flow-arrow">↓</div>
  <div class="flow-card">
    <div class="flow-badge">Step 5</div>
    <div class="flow-title">Troop &amp; Rules Handshake</div>
    <div class="flow-desc">Runs Troop setup, moves <code>TROOP.md</code> to <code>.cooper/</code>, and injects guidelines into <code>AGENTS.md</code></div>
  </div>
</div>

### 1. [Troop](https://twoboots.github.io/troop) Foundation Setup
The installer runs the Troop installer ([twoBoots/troop](https://github.com/twoBoots/troop)) to establish Git worktree isolation:
* Sets up Git command aliases (`git agent-start <track_id>`, `git troop`, `git agent-stop <track_id>`).
* Updates `.gitignore` to exclude `.worktrees/`.
* Relocates `TROOP.md` reference guide to `.cooper/TROOP.md` to keep the project root clean.

### 2. Auto-Migration & Scaffolding Scenarios

#### Scenario A: Existing Conductor Setup Detected
If an existing `conductor/` directory is present in the target repository:
* **Definitions**: Migrates `product.md`, `tech-stack.md`, `product-guidelines.md`, and `workflow.md` to `.cooper/definition/`.
* **Code Style Guides**: Migrates `conductor/code_styleguides/` to `.cooper/code_styleguides/`.
* **Tracks**: Migrates `conductor/tracks/` to `.cooper/archive/`.

#### Scenario B: Existing OpenSpec Setup Detected
If an existing `openspec/` directory is present in the target repository:
* **Living Specs**: Copies capability specs from `openspec/specs/` directly to `.cooper/specs/`.
* **Active Changes**: Copies change proposals from `openspec/changes/` to `.cooper/active/`.

#### Scenario C: Greenfield Project (Neither Exists)
If neither `conductor/` nor `openspec/` is found:
* **Scaffolds Baseline Templates**: Installs project definitions (`product.md`, `tech-stack.md`, `product-guidelines.md`) and code styleguides (`typescript.md`, `python.md`, `go.md`, `rust.md`) from Cooper's native templates.
* **Scaffolds `.cooper/` Directory**: Initializes the baseline directory structure:
  ```
  .cooper/
  ├── index.md
  ├── COOPER.md
  ├── TROOP.md
  ├── tracks.md
  ├── definition/
  ├── code_styleguides/
  ├── specs/
  ├── active/
  └── archive/
  ```

### 3. Native Agent Skills Installation (`.agents/skills/`)
The installer installs self-contained project skills into `.agents/skills/`:
* `/cooper-setup` (`cooper-setup/SKILL.md`): Audits and configures the environment.
* `/cooper-rfc` (`cooper-rfc/SKILL.md`): Plans collaborative RFCs, living spec deltas, and Draft PR reviews ([RFC Draft PR Guide](rfc-draft-prs.md)).
* `/cooper-new-track` (`cooper-new-track/SKILL.md`): Spawns worktree and plans spec deltas.
* `/cooper-implement` (`cooper-implement/SKILL.md`): Executes TDD and phase sync.
* `/cooper-review` (`cooper-review/SKILL.md`): Audits implementation quality and spec delta fidelity.
* `/cooper-status` (`cooper-status/SKILL.md`): Displays active worktree and track progress.

### 4. Agent Rules Injection (`AGENTS.md`)
The installer creates or appends to `AGENTS.md` with rules instructing AI agents to follow:
* `.cooper/COOPER.md` quick reference.
* Native `/cooper-*` skill workflows (`/cooper-setup`, `/cooper-rfc`, `/cooper-new-track`, `/cooper-implement`, `/cooper-review`, `/cooper-status`).
* `.cooper/specs/` living spec reading.
* `.cooper/active/<track_id>/spec-deltas/` requirement diff generation.
* TDD Red/Green/Refactor cycle with Git Notes summaries (`git notes add -m`).
* Phase completion synchronization (`git fetch origin main` & `git push origin <track_id>`).
* [Troop](https://twoboots.github.io/troop) worktree isolation (`.worktrees/<track_id>`).

<style>
.installer-flow {
  display: flex;
  flex-direction: column;
  gap: 10px;
  margin: 24px 0;
}
.flow-card {
  background-color: var(--vp-c-bg-soft);
  border: 1px solid var(--vp-c-divider);
  border-radius: 8px;
  padding: 14px 18px;
  box-sizing: border-box;
}
.flow-badge {
  font-size: 11px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: var(--vp-c-brand-1);
  margin-bottom: 4px;
}
.flow-title {
  font-size: 15px;
  font-weight: 700;
  color: var(--vp-c-text-1);
  margin-bottom: 4px;
}
.flow-desc {
  font-size: 13px;
  color: var(--vp-c-text-2);
}
.flow-arrow {
  text-align: center;
  font-size: 18px;
  font-weight: bold;
  color: var(--vp-c-text-3);
  user-select: none;
  line-height: 1;
}
.branch-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: 10px;
  margin-top: 10px;
}
.branch-item {
  background-color: var(--vp-c-bg);
  border: 1px solid var(--vp-c-divider);
  border-radius: 6px;
  padding: 10px;
}
.branch-tag {
  display: inline-block;
  font-size: 11px;
  font-weight: 700;
  color: var(--vp-c-brand-1);
  margin-bottom: 4px;
}
.branch-text {
  font-size: 12px;
  color: var(--vp-c-text-2);
  line-height: 1.4;
}
</style>

