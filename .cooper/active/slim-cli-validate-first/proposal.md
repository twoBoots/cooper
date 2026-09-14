# Track Proposal: Slim the Cooper CLI to a Validate-First Surface

- **Track ID**: `slim-cli-validate-first`
- **Type**: chore
- **Status**: Planning
- **Supersedes**: `rfc-cooper-cli-mcp` (approved 2026-08-20) — partial reversal, see §5

## 1. Summary

Cooper's product is a prompt framework: six `SKILL.md` files, a directory convention, and three git aliases. The agent is the runtime. Beside that product sits a 3,949-line Go binary that the framework never calls.

This track reduces the binary to the one capability that genuinely cannot be delegated to a language model — deterministic specification validation — and removes the surface that duplicates what the agent already does well.

## 2. Motivation

An audit of the CLI and MCP server against actual usage found:

- **Zero invocations.** No reference to `cooper init`, `cooper validate`, `cooper track`, or any `cooper_*` MCP tool exists in any of the six skills, either CI workflow, or any of the five archived tracks. Cooper does not use its own CLI.
- **The one valuable command is not wired in.** `cooper validate` caught 9 of 9 planted defects in a scratch repository — malformed requirement headers, missing normative keywords, a mistyped GIVEN, track-ID/directory mismatch, invalid type and status enums, a missing `created_at`, registry parity drift, and a dangling markdown link. CI does not run it.
- **The coverage gate is decorative.** `ci.yml` prints `go tool cover -func` and exits 0 regardless, despite `workflow.md` §6 mandating >80%.
- **Two scaffolders have already diverged.** `install.sh` copies skills from `skills/`; `cooper init` writes them from `internal/scaffold/assets/skills/`. Both are checked in with no sync step, and all six files now differ — the embedded copies carry the `mandate-interactive-question-tools` changes, the root copies do not. Two install paths produce two different frameworks.
- **`cooper init` emits a broken project.** It writes an `AGENTS.md` mandating `.cooper/COOPER.md`, `git agent-start`, and `git troop`, then ships none of them: no `COOPER.md`, no `TROOP.md`, no git aliases, no `.gitignore` entry for `.worktrees/`. `cooper validate` reports success on that project because the dangling reference is backticked rather than a markdown link.
- **`cooper track checkpoint` fabricates its audit trail.** `internal/track/track.go:94` hardcodes `Automated Tests: PASSED` and `Manual Verification: APPROVED by user` into the git note without running a test, checking an exit code, or prompting anyone. Cooper's README sells Git Notes audit trails as a quality guarantee; this command makes that guarantee false on demand.
- **The MCP server is 25% of the codebase wrapping local shell operations.** Six tools, every one a local file operation inside a repo the agent already has shell access to. They occupy context in every session of every project that installs them.

## 3. Scope

**In scope:**

1. Remove the MCP server and client installer (`internal/mcp/`, `cmd/mcp.go`, and tests) — 1,019 lines.
2. Remove `cooper track checkpoint` (fabricated audit records) and the `track` subcommand tree that duplicates `cooper-new-track` and `cooper-status`.
3. Collapse the two scaffolders into one source of truth and delete the duplicated skill tree.
4. Wire `cooper validate` and a real 80% coverage gate into `.github/workflows/ci.yml`.
5. Teach the link auditor to resolve backticked repository paths, so `cooper init` output stops passing validation while broken.
6. Update `README.md` and the `cli` / `installer` living specs to match the reduced surface.

**Out of scope:**

- The unbuilt upstream-sync engine from `rfc-cooper-cli-mcp` (`.cooper/manifest.json`, 3-way diff, `cooper_check_updates` / `cooper_diff_updates` / `cooper_apply_update`). See §5.
- The "suggestive, not prescriptive" documentation work owned by the in-review `rfc-suggestive-framework-ethos`.
- `cooper update` / `cooper version`, which are cheap, standard, and retained unchanged.

## 4. User Benefit

- A single install path that produces a working project instead of two paths producing two divergent ones.
- An audit trail that reflects reality, or no audit trail command at all.
- Specification drift caught by CI on every pull request rather than by whoever notices.
- ~1,800 fewer lines to maintain, and six fewer tool definitions in every consuming agent's context window.

## 5. Relationship to `rfc-cooper-cli-mcp` (Explicit Reversal)

This track partially reverses an approved RFC. That is recorded here deliberately rather than left implicit.

The approved RFC justified the MCP server with three tools — `cooper_check_updates`, `cooper_diff_updates`, and `cooper_apply_update` — backed by a `.cooper/manifest.json` of upstream file hashes and a 3-way merge engine. That design addressed RFC problem statement #1, *"No Safe Upstream Upgrades."* **None of those three tools was ever implemented.** There is no `internal/updater/`, no `manifest.go`, no `diff3.go`, and no `.cooper/manifest.json`. What shipped instead were six thin wrappers around local file operations.

Two consequences are acknowledged:

1. **The MCP is being removed in its as-built form, not its as-designed form.** The as-designed upstream-sync tooling remains an open, unaddressed idea; this track does not argue against it and does not foreclose it.
2. **RFC problem statement #1 remains unsolved after this track.** It is the demonstrated root cause of the skill-tree divergence fixed in Phase 2 — collapsing to one copy removes today's drift but installs no mechanism preventing recurrence for downstream consumers. A future RFC should address it.

The same RFC's tiered binary installation for `install.sh` (Tier 1 local compile / Tier 2 release download / Tier 3 zero-binary fallback) was likewise never built, which is why the documented Quick Start installs no binary. Phase 2 resolves this by making the binary optional and explicit rather than advertised and absent.

## 6. Risks

| Risk | Mitigation |
| :--- | :--- |
| Users who ran `cooper mcp install` have configs pointing at a removed command | Phase 1 ships a removal note in `README.md`; the stale entry fails visibly at startup rather than silently misbehaving |
| Mandatory CI validation collides with the in-review "suggestive, not prescriptive" ethos RFC | Gate is applied to **this** repository's CI only, not injected into scaffolded consumer projects |
| Removing `track` subcommands strands any external caller | No callers exist in-repo; commands are absent from all skills, CI, and archived tracks |
| Version is a breaking change to a released public surface | Bump minor with a documented removal list in the release notes |
