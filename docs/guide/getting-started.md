# Getting Started with Cooper

Cooper is a suggestive, agent-agnostic Spec-Driven Development (SDD) framework—a hybrid uniting [OpenSpec](https://openspec.dev)'s living capability specifications with Conductor's quality governance, paired with [Troop](https://twoboots.github.io/troop) worktree isolation. Designed to be customized to any tech stack, Cooper eliminates context drift, hallucinated requirements, and trunk pollution in autonomous AI agent engineering workflows.

## Prerequisites

- **Git** 2.30+
- **Go** 1.27+ (if building the CLI from source)

## Installation

You can install the Cooper CLI using the official installer:

```bash
curl -sSL https://raw.githubusercontent.com/twoBoots/cooper/main/install.sh | bash
```

For detailed platform instructions and manual verification steps, see the [Installation Guide](../INSTALL.md).

## Initializing Cooper in a Repository

To adopt Cooper in an existing repository, run the initialization skill:

```text
/cooper-setup
```

Or validate SDD syntax directly with the Cooper CLI:

```bash
cooper validate
```

When Cooper is initialized, a `.cooper/` directory is created with:
- **`definition/`**: Product definition, tech stack, and workflow guidelines.
- **`specs/`**: Living capability specifications grounding all future work.
- **`active/`**: Active tracks currently in development.
- **`archive/`**: Historical record of completed tracks.
- **`tracks.md`**: Master registry of all tracks.

## Core Concepts

### 1. Living Capability Specs (`.cooper/specs/`)
Living specs define truth for each system capability. Before code is written, any change must be articulated as a Spec Delta adding (`+`) or removing (`-`) specific GIVEN / WHEN / THEN behavioral requirements (see [OpenSpec.dev](https://openspec.dev)).

### 2. Isolated Worktrees with [Troop](https://twoboots.github.io/troop)
Never develop features directly on the main repository trunk. Cooper uses [Troop](https://twoboots.github.io/troop) to spin up ephemeral, isolated worktrees under `.worktrees/<track_id>`, guaranteeing reproducible builds without colliding with parallel tracks.

### 3. Strict TDD Loop
Every task adheres to Red -> Green -> Refactor:
1. **Red**: Write a failing unit/integration test capturing the Spec Delta requirement.
2. **Green**: Implement the minimum code required to satisfy the test.
3. **Refactor**: Clean up and optimize while keeping tests passing and coverage >80%.

### 4. Git Notes Metadata Tracking
Every task execution and phase verification checkpoint is recorded directly into Git Notes on commits, preserving full machine-readable provenance.

::: tip Ethos: Suggestive, Not Prescriptive
Cooper is suggestive, not prescriptive. It provides guardrails to keep agents grounded, but leaves teams free to adapt skills, tech stacks, and workflows to their project needs.
:::

Next, explore the [Cooper Workflow & Lifecycle](./workflow.md).

