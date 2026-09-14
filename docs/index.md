---
layout: home

hero:
  name: Cooper
  text: Spec-Driven Development Framework
  tagline: Eliminate agentic drift with grounded capability specs, Troop worktree isolation, and strict TDD.
  actions:
    - theme: brand
      text: Get Started
      link: /guide/getting-started
    - theme: alt
      text: View on GitHub
      link: https://github.com/twoBoots/cooper

features:
  - icon: 📑
    title: Living Capability Specs
    details: Ground agent planning in living specs (.cooper/specs/). Changes produce explicit Spec Deltas (+ / -) before code is written.
  - icon: 🌲
    title: Worktree Isolation
    details: Isolate parallel tracks inside dedicated Git worktrees via [Troop](https://github.com/twoBoots/troop), keeping trunk clean.
  - icon: 🎯
    title: TDD Quality Gates
    details: Enforce the Red -> Green -> Refactor cycle, maintain >80% code coverage, and enforce phase verification checkpoints.
  - icon: 📝
    title: Git Notes Metadata
    details: Automatically record task execution context, SHAs, and verification reports as Git Notes metadata on commits.
---

## Quickstart

Install the Cooper CLI with a single command:

```bash
curl -sSL https://raw.githubusercontent.com/twoBoots/cooper/main/install.sh | bash
```

---

## The Cooper Lifecycle

Cooper structures agentic engineering into clear, verifiable phases:

```text
  ┌─────────────────┐       ┌─────────────────┐       ┌─────────────────┐       ┌─────────────────┐
  │   cooper-rfc    │ ────> │ cooper-new-track│ ────> │ cooper-implement│ ────> │  cooper-review  │
  │ Architectural   │       │ Worktree Spawn  │       │ Strict TDD      │       │ Spec Delta & PR │
  │ Initiative      │       │ & Spec Delta    │       │ & Git Notes     │       │ Verification    │
  └─────────────────┘       └─────────────────┘       └─────────────────┘       └─────────────────┘
```

1. **RFC & Initiative Planning (`cooper-rfc`)**: Align on architectural scope, living spec impact, and draft PR reviews.
2. **Track Scaffolding (`cooper-new-track`)**: Spawn an isolated [Troop](https://github.com/twoBoots/troop) worktree and author the Spec Delta and TDD implementation plan.
3. **Execution (`cooper-implement`)**: Run the strict Red-Green-Refactor loop, attach Git Notes execution metadata, and pass phase checkpoints.
4. **Principal Review (`cooper-review`)**: Verify code against spec delta requirements, style guides, and test coverage before opening a PR.

---

## Learn More

- [Getting Started Guide](./guide/getting-started.md)
- [Workflow & Lifecycle Details](./guide/workflow.md)
- [Installation Guide](./INSTALL.md)
- [Compare with Other Frameworks](./openspec-vs-conductor-comparison.md)
