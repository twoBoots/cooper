# Getting Started with Cooper

Cooper is a Spec-Driven Development (SDD) framework paired with [Troop](https://github.com/twoBoots/troop) worktree isolation, designed to eliminate context drift, hallucinated requirements, and trunk pollution in autonomous AI agent engineering workflows.

## Prerequisites

- **Git** 2.30+
- **Go** 1.27+ (if building the CLI from source)
- **Node.js** 20+ (for documentation site development)

## Installation

You can install the Cooper CLI using the official installer:

```bash
curl -sSL https://raw.githubusercontent.com/twoBoots/cooper/main/install.sh | bash
```

Or via Go:

```bash
go install github.com/twoBoots/cooper@latest
```

For detailed platform instructions and manual verification steps, see the [Installation Guide](/INSTALL).

## Initializing Cooper in a Repository

To adopt Cooper in an existing repository, run the initialization skill or scaffold manually:

```bash
# Verify or validate SDD syntax in your repo
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
Living specs define truth for each system capability. Before code is written, any change must be articulated as a Spec Delta adding (`+`) or removing (`-`) specific GIVEN / WHEN / THEN behavioral requirements.

### 2. Isolated Worktrees with [Troop](https://github.com/twoBoots/troop)
Never develop features directly on the main repository trunk. Cooper uses [Troop](https://github.com/twoBoots/troop) to spin up ephemeral, isolated worktrees under `.worktrees/<track_id>`, guaranteeing reproducible builds without colliding with parallel tracks.

### 3. Strict TDD Loop
Every task adheres to Red -> Green -> Refactor:
1. **Red**: Write a failing unit/integration test capturing the Spec Delta requirement.
2. **Green**: Implement the minimum code required to satisfy the test.
3. **Refactor**: Clean up and optimize while keeping tests passing and coverage >80%.

### 4. Git Notes Metadata Tracking
Every task execution and phase verification checkpoint is recorded directly into Git Notes on commits, preserving full machine-readable provenance.

Next, explore the [Cooper Workflow & Lifecycle](/guide/workflow).
