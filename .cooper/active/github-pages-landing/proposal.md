# Proposal: GitHub Pages Entry-Level Informative Landing Site

## Problem Statement
The Cooper repository currently presents technical documentation in raw Markdown files (`README.md`, `docs/`, `.cooper/`). While thorough for contributors inspecting files directly, developers evaluating Cooper from the web or looking for a fast introduction lack an entry-level, visually engaging, and accessible landing page. A polished GitHub Pages site lowers the barrier to entry, highlights Cooper's core pillars (Spec-Driven Development + [Troop](https://github.com/twoBoots/troop) worktree isolation), provides immediate installation copy-pastes, and guides users into the documentation ecosystem.

## Proposed Solution
Scaffold a lightweight Static Site Generator (VitePress) site powered by npm and automated via GitHub Actions (`.github/workflows/pages.yml`). The site will provide:
1. **Polished Landing Page (`docs/index.md`)**:
   - Hero banner with concise value proposition and primary action links ("Get Started", "GitHub").
   - Instant Quickstart installation command (`curl -sSL ...`).
   - Core Pillars & Feature highlights (Cooper SDD Framework, [Troop](https://github.com/twoBoots/troop) Worktree Isolation, Git Notes Metadata, TDD Discipline).
   - High-level visual workflow diagram explaining the Cooper lifecycle (`cooper-rfc` -> `cooper-new-track` -> `cooper-implement` -> `cooper-review`).
2. **Navigation & Docs Integration**:
   - Clean top navigation bar and responsive layout.
   - Integration with existing repository docs (`INSTALL.md`, comparison guides, architecture overview).
   - Upstream linking to [Troop](https://github.com/twoBoots/troop) across all pages.
3. **Automated CI/CD**:
   - GitHub Actions workflow deploying the built VitePress static bundle to GitHub Pages upon push to `main`.

## Scope & Boundaries
- **In Scope**:
  - `package.json` setup for documentation build dependencies (`vitepress`).
  - VitePress configuration (`docs/.vitepress/config.mts`) with base path, nav, sidebar, theme settings, and social links.
  - Landing page template (`docs/index.md`) using VitePress home layout.
  - Integration of foundational guides (`docs/INSTALL.md`, overview pages).
  - GitHub Actions deployment workflow (`.github/workflows/pages.yml`).
  - Validation test scripts / verification for doc build and link integrity.
- **Out of Scope**:
  - Migration of internal `.cooper/` active track files or living specs to public web pages (those remain within the repository structure).
  - Dynamic backend services or authentication.
