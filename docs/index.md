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
    details: Ground agent planning in living specs under <code>.cooper/specs/</code> following <a href="https://openspec.dev" target="_blank" rel="noopener">OpenSpec</a>. Changes produce explicit Spec Deltas (+ / -) before code is written.
  - icon: 🌲
    title: Worktree Isolation
    details: Isolate parallel tracks inside dedicated Git worktrees via <a href="https://twoboots.github.io/troop" target="_blank" rel="noopener">Troop</a>, keeping trunk clean.
  - icon: 🎯
    title: TDD Quality Gates
    details: Enforce the Red -> Green -> Refactor cycle, >80% code coverage, and human in the loop phase verification checkpoints.
  - icon: 📝
    title: Git Notes Metadata
    details: Automatically record task execution context, SHAs, and verification reports as Git Notes metadata on commits to preserve against history rewrites.
---

## Quickstart

Install the Cooper CLI with a single command:

```bash
curl -sSL https://raw.githubusercontent.com/twoBoots/cooper/main/install.sh | bash
```

---

## The Cooper Lifecycle

Cooper structures agentic engineering into clear, verifiable phases:

<div class="lifecycle-flow">
  <div class="lifecycle-card">
    <span class="lifecycle-badge">Phase 1</span>
    <div class="lifecycle-step">/cooper-rfc</div>
    <div class="lifecycle-sub">Architectural Initiative</div>
  </div>
  <div class="lifecycle-arrow">→</div>
  <div class="lifecycle-card">
    <span class="lifecycle-badge">Phase 2</span>
    <div class="lifecycle-step">/cooper-new-track</div>
    <div class="lifecycle-sub">Worktree &amp; Spec Delta</div>
  </div>
  <div class="lifecycle-arrow">→</div>
  <div class="lifecycle-card">
    <span class="lifecycle-badge">Phase 3</span>
    <div class="lifecycle-step">/cooper-implement</div>
    <div class="lifecycle-sub">Strict TDD &amp; Git Notes</div>
  </div>
  <div class="lifecycle-arrow">→</div>
  <div class="lifecycle-card">
    <span class="lifecycle-badge">Phase 4</span>
    <div class="lifecycle-step">/cooper-review</div>
    <div class="lifecycle-sub">Spec Delta &amp; PR Gate</div>
  </div>
</div>

1. **RFC & Initiative Planning (`/cooper-rfc`)**: Align on architectural scope, living spec impact, and draft PR reviews.
2. **Track Scaffolding (`/cooper-new-track`)**: Spawn an isolated [Troop](https://twoboots.github.io/troop) worktree and author the Spec Delta and TDD implementation plan.
3. **Execution (`/cooper-implement`)**: Run the strict Red-Green-Refactor loop, attach Git Notes execution metadata, and pass phase checkpoints.
4. **Principal Review (`/cooper-review`)**: Verify code against spec delta requirements, style guides, and test coverage before opening a PR.

---

## Learn More

- [Getting Started Guide](./guide/getting-started.md)
- [Workflow & Lifecycle Details](./guide/workflow.md)
- [Installation Guide](./INSTALL.md)
- [Compare with Other Frameworks](./openspec-vs-conductor-comparison.md)

<style>
.lifecycle-flow {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  gap: 8px;
  margin: 24px 0;
  flex-wrap: wrap;
}
.lifecycle-card {
  flex: 1 1 140px;
  min-width: 130px;
  background-color: var(--vp-c-bg-soft);
  border: 1px solid var(--vp-c-divider);
  border-radius: 8px;
  padding: 12px;
  text-align: center;
  box-sizing: border-box;
}
.lifecycle-badge {
  display: inline-block;
  font-size: 11px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  color: var(--vp-c-brand-1);
  margin-bottom: 4px;
}
.lifecycle-step {
  font-family: var(--vp-font-family-mono);
  font-size: 13px;
  font-weight: 700;
  color: var(--vp-c-text-1);
  margin-bottom: 4px;
}
.lifecycle-sub {
  font-size: 12px;
  color: var(--vp-c-text-2);
  line-height: 1.3;
}
.lifecycle-arrow {
  font-size: 18px;
  font-weight: bold;
  color: var(--vp-c-text-3);
  user-select: none;
}
@media (max-width: 640px) {
  .lifecycle-flow {
    flex-direction: column;
    align-items: stretch;
  }
  .lifecycle-arrow {
    text-align: center;
    transform: rotate(90deg);
  }
  .lifecycle-card {
    min-width: 100%;
  }
}
</style>
