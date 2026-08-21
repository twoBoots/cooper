# RFC: Establish "Suggestive, Not Prescriptive" Framework Ethos in Documentation & Lifecycle

- **RFC ID**: `rfc-suggestive-framework-ethos`
- **Author**: AI Agent & Team
- **Status**: In Review (Draft)
- **Created**: 2026-08-21

---

## 1. Summary & Motivation

### Problem Statement
Initial documentation for Cooper risked communicating a rigid, dogmatic framework mandate. Users and AI agents reading the repository might perceive Cooper's Spec-Driven Development (SDD), TDD loops, and worktree isolation as immutable constraints rather than an adaptable, empowering foundation.

### Proposed Solution
Articulate and enshrine Cooper's core ethos across framework documentation and agent guidelines: **Cooper is suggestive, not prescriptive**. 

Cooper provides a battle-tested starting point (living spec deltas, TDD execution loops, worktree isolation via Troop, and native agent skills), but the primary goal is for teams to **install it and adapt it** to their project's unique tech stack, workflows, conventions, and culture.

---

## 2. Architecture & Detailed Design

### 2.1 Core Pillars of the Suggestive Ethos

1. **Suggestive Baseline, Not a Rigid Mandate**:
   - Out-of-the-box defaults (e.g., TDD Red/Green/Refactor, coverage targets >80%, commit notes) represent recommended best practices, but every rule and template can be modified or extended.
2. **The 3-Step Lifecycle: Install → Adapt → Build**:
   - **1. Install**: Scaffold `.cooper/`, configure [Troop](https://github.com/twoBoots/troop) worktree aliases, and copy native agent skills to `.agents/skills/`.
   - **2. Adapt**: Customize `.cooper/definition/` (`product.md`, `tech-stack.md`, `workflow.md`), add conventions to `.cooper/code_styleguides/`, and tune `.agents/skills/`.
   - **3. Build**: Execute features with isolated worktrees and living spec deltas using `cooper-new-track` and `cooper-rfc`.
3. **Repository Ownership**:
   - All rules, skills, and configuration reside in-repo (`.cooper/` and `.agents/skills/`). There are no external runtime dependencies or vendor lock-in.
4. **Agent-Agnostic & Self-Contained**:
   - Any AI coding assistant (Claude Code, Gemini CLI, Cursor, Roo, Copilot, Antigravity) works out of the box by reading repository markdown instructions.

### 2.2 Documentation & Specification Changes

- **`README.md`**: Add prominent "💡 Philosophy: Suggestive, Not Prescriptive" section, the 3-step lifecycle diagram, and updated structure / adaptation guide.
- **`.cooper/specs/documentation/spec.md`**: Update capability requirements to require suggestive ethos framing across all top-level documentation.
- **Templates & Agent Guidelines**: Ensure setup skills (`cooper-setup`) and templates encourage user adaptation during project onboarding.

---

## 3. Alternatives Considered

| Alternative | Description | Trade-offs / Reason for Rejection |
| :--- | :--- | :--- |
| **Strict / Prescriptive Enforcement** | Treat all Cooper rules and workflows as mandatory invariants with strict linting and hard CI checks. | Too rigid; alienates projects with existing conventions, non-TDD workflows, or specialized delivery pipelines. |
| **Minimal / Unopinionated Scaffolding** | Provide empty directories without opinionated defaults or starter skills. | Fails to provide immediate value or clear guidance on how to run spec-driven development with AI agents. |
| **Suggestive Baseline with Guided Adaptation (Selected)** | Provide a complete, production-ready starting suite with clear instructions and tools to adapt everything. | Delivers immediate high value while granting teams full sovereignty over their engineering rules. |

---

## 4. Security, Performance & Scalability

- **Security**: No runtime security impact.
- **Performance**: Zero overhead; purely documentation, templates, and agent guidance.
- **Portability**: Improves portability by making it easier to adapt Cooper to diverse languages, frameworks, and CI/CD tools.

---

## 5. Rollout & Migration Strategy

The changes will be implemented via a single downstream track:
1. `track-readme-suggestive-ethos`: Update `README.md` to feature the suggestive philosophy, 3-step lifecycle, and adaptation instructions.

---

## 6. Open Questions & Discussion Topics

1. **Interactive Setup Customization**: Should `cooper-setup` and `cooper init` include additional interactive prompts to help users configure non-standard workflows (e.g. relaxing coverage or custom git workflows)?
2. **Ecosystem Guides**: Should we add dedicated recipe guides in `docs/` for adapting Cooper to specific environments (e.g., monorepos, microservices, mobile apps)?
