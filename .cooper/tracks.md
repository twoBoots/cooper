# Tracks Registry

All active and completed Cooper tracks are registered below.

---

- [x] **Track: Relocate TROOP.md into .cooper/ and align repo structure to .cooper/**
  - Archive: [.cooper/archive/relocate-troop-to-cooper/](.cooper/archive/relocate-troop-to-cooper/)
- [x] **Track: Document Draft PR Rationale in Cooper RFC Lifecycle**
  - Worktree: `.worktrees/docs-rfc-draft-prs`
  - Link: [.cooper/active/docs-rfc-draft-prs/index.md](.cooper/active/docs-rfc-draft-prs/index.md)

- [x] **Track: RFC PR Approval Detection Protocols & Reviewer Instructions**
  - Worktree: `.worktrees/rfc-pr-approval-protocols`
  - Link: [.cooper/active/rfc-pr-approval-protocols/index.md](.cooper/active/rfc-pr-approval-protocols/index.md)

- [x] **Track: Link Troop Mentions in Markdown Documentation to [twoBoots/troop](https://github.com/twoBoots/troop)**
  - Archive: [.cooper/archive/link-troop-in-docs/](.cooper/archive/link-troop-in-docs/)

- [ ] **RFC: Cooper Go CLI & MCP Server Architecture** (`rfc-cooper-cli-mcp`)
  - RFC Doc: [.cooper/active/rfc-cooper-cli-mcp/rfc.md](.cooper/active/rfc-cooper-cli-mcp/rfc.md)
  - Status: Approved
  - Decomposed Tracks:
    - [x] Track: `track-cooper-cli-core` (Scope: Go CLI project scaffolding, Cobra command tree, and SDD syntax validator)
        - Worktree: `.worktrees/track-cooper-cli-core`
        - Link: [.cooper/active/track-cooper-cli-core/index.md](.cooper/active/track-cooper-cli-core/index.md)
    - [ ] Track: `track-cooper-updater-diff3` (Scope: CLI binary self-update, 3-way diff reconciliation engine, and manifest fingerprinting)
        - Note: Binary self-update shipped via `cooper-bender-integration`. The 3-way diff and manifest fingerprinting remain unbuilt; RFC problem statement #1 ("No Safe Upstream Upgrades") is still open.
    - [ ] Track: `track-cooper-embedded-mcp` (Scope: Embedded stdio MCP server exposing SDD, self-update, and 3-way diff tools)
        - Note: Superseded by `slim-cli-validate-first`, which removes the as-built MCP server. The as-designed upstream-sync tooling is not foreclosed.
    - [ ] Track: `track-cooper-installer-packaging` (Scope: 3-tier install.sh script, CI validation, and GitHub Actions multi-arch release matrix)
- [x] **Track: Cooper Bender CLI Integration**
  - Archive: [.cooper/archive/cooper-bender-integration/](.cooper/archive/cooper-bender-integration/)
- [x] **Track: Bump Go Toolchain and Runtime Baseline to 1.27.0**
  - Archive: [.cooper/archive/bump-go-1-27-0/](.cooper/archive/bump-go-1-27-0/)
- [x] **Track: Automated Multi-Platform Release CI/CD Pipeline & v1.0.0 Baseline**
  - Worktree: `.worktrees/release-ci-pipeline`
  - Link: [.cooper/active/release-ci-pipeline/index.md](.cooper/active/release-ci-pipeline/index.md)


- [x] **Track: Mandate Interactive Question & Native File Tool Calls**
  - Archive: [.cooper/archive/mandate-interactive-question-tools/](.cooper/archive/mandate-interactive-question-tools/)

- [ ] **Track: Slim the Cooper CLI to a Validate-First Surface** (`slim-cli-validate-first`)
  - Worktree: `.worktrees/slim-cli-validate-first`
  - Link: [.cooper/active/slim-cli-validate-first/index.md](.cooper/active/slim-cli-validate-first/index.md)
  - Supersedes: `rfc-cooper-cli-mcp` (partial reversal — removes the as-built MCP server, `init`, and `track` commands)

