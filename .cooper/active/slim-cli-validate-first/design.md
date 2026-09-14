# Technical Design: Slim the Cooper CLI to a Validate-First Surface

- **Track ID**: `slim-cli-validate-first`

## 1. Architecture

### 1.1 Guiding Principle

The binary earns its place only where a language model cannot be trusted to check its own work. Everything else in Cooper is markdown the agent reads and acts on.

```
BEFORE (3,949 lines)                  AFTER (~1,640 lines)
─────────────────────                 ─────────────────────
cooper init      ─┐                   cooper validate   ← the deterministic gate
cooper track *   ─┤ duplicate the     cooper update     ← binary self-update
cooper mcp       ─┤ skills or wrap    cooper version
cooper mcp install┤ the shell
cooper validate  ─┘ ← the only        install.sh        ← the single scaffolder
cooper update                         skills/           ← the single skill tree
cooper version
install.sh       ─┐ two scaffolders
cooper init      ─┘ already diverged
```

### 1.2 Decision: `install.sh` becomes the single scaffolder

`cooper init` is deleted rather than repaired. This is the key structural choice, and it resolves three findings at once:

1. **It fixes the broken-init defect by construction.** `cooper init` emitted an `AGENTS.md` referencing `.cooper/COOPER.md`, `git agent-start`, and `git troop` while shipping none of them. `install.sh` already installs all of these correctly by delegating to the Troop installer. Removing the broken path is cheaper and safer than teaching it to duplicate the working one.
2. **It permanently ends the skill-tree drift.** With `internal/scaffold/assets/` gone, exactly one copy of each `SKILL.md` exists. There is no second copy to diverge, so no sync step or CI parity check is needed.
3. **It preserves the zero-binary property.** `install.sh` scaffolds a complete, working Cooper project with no compiled artifact. Making the scaffolder a shell script rather than a Go command keeps the framework installable in environments that never obtain the binary — consistent with the ethos work in `rfc-suggestive-framework-ethos`.

The binary is thereby demoted from "how you install Cooper" to "an optional checker for projects that want CI enforcement."

### 1.3 Content migration direction (critical)

The two skill trees are **not** equal, and the direction of the merge matters. `internal/scaffold/assets/skills/` is the *newer* tree — it carries the `mandate-interactive-question-tools` changes (Interactive Question Protocol, Native File Tools Mandate) that root `skills/` lacks.

Therefore: **port embedded → root, then delete embedded.** Deleting the embedded tree first would silently discard the newer content. Phase 2 Task 2.1 does the port before Task 2.2 does the deletion, and a test asserts the mandate text survives.

`templates/` and `AGENTS.template.md` are byte-identical across both trees (verified), so those require no merge — only deletion of the embedded copy.

## 2. Component Breakdown

### 2.1 Deletions

| Path | Lines | Rationale |
| :--- | ---: | :--- |
| `internal/mcp/server.go`, `server_test.go` | 749 | Six tools wrapping local file operations |
| `cmd/mcp.go`, `cmd/mcp_test.go` | 270 | `mcp` + `mcp install` commands |
| `internal/track/` (`track.go`, `git.go`, tests) | 414 | Duplicates `cooper-new-track` / `cooper-status`; `RecordCheckpoint` fabricates audit records |
| `cmd/track.go`, `cmd/track_test.go` | 343 | Command surface for the above |
| `internal/scaffold/` (`init.go`, `embedded.go`, tests) | 434 | Second scaffolder; source of the drift |
| `cmd/init.go`, `cmd/init_test.go` | 104 | Command surface for the above |
| `internal/scaffold/assets/` | — | Duplicate skill/template tree (content ported first) |

Approximate total removed: **2,314 lines**, plus the duplicated asset tree.

### 2.2 Retained

| Path | Change |
| :--- | :--- |
| `internal/validator/` | Extended — backticked-path link auditing (§2.3) |
| `cmd/validate.go` | Unchanged |
| `cmd/update.go`, `cmd/version.go` | Unchanged |
| `cmd/root.go` | Subcommand registrations for `init`, `track`, `mcp` removed |
| `go.mod` | `github.com/twoBoots/bender` retained — `pkg/updater` still backs `cooper update`; only `pkg/mcp` usage is dropped |

### 2.3 Link auditor extension

The auditor currently resolves only true markdown link syntax — bracketed text followed by a parenthesised target. `AGENTS.md` references `.cooper/COOPER.md` in backticks, so a project missing that file validates clean.

Extend `internal/validator/link_auditor.go` to additionally resolve inline-code spans that are unambiguously repository paths, under conservative rules to avoid false positives on shell snippets and glob patterns:

- Contains a `/` **and** ends in a known documentation extension (`.md`, `.json`, `.yml`, `.yaml`), **or** ends in `/` and names an existing-or-missing directory.
- Does not begin with `-`, `$`, `~`, or a URL scheme.
- Contains no shell metacharacters (`*`, `|`, `<`, `>`, `` ` ``, spaces).
- Resolves relative to the containing file, then to the repository root.

New rule ID: `link/code-path-exists`. Emitted at the same severity as `link/target-exists`.

### 2.4 CI gates

`.github/workflows/ci.yml` gains two enforcement steps:

```yaml
- name: Validate Cooper SDD Specs
  run: go run . validate

- name: Enforce Coverage Threshold
  run: |
    COVERAGE=$(go tool cover -func=coverage.out | tail -1 | awk '{print substr($3, 1, length($3)-1)}')
    echo "Total coverage: ${COVERAGE}%"
    awk -v c="$COVERAGE" 'BEGIN { exit (c < 80) }' || { echo "Coverage ${COVERAGE}% below 80% threshold"; exit 1; }
```

Scoped to this repository's own CI. Neither gate is injected into scaffolded consumer projects, keeping the change compatible with the in-review "suggestive, not prescriptive" ethos.

### 2.5 `install.sh` binary step (optional, non-fatal)

A new step offers the binary without requiring it, matching the tiering the approved RFC specified but never shipped:

1. If run from a local clone with `go` on `PATH` → `go build -ldflags="-s -w" -o <bin_dir>/cooper .`
2. Else attempt download of `cooper-${OS}-${ARCH}` from GitHub Releases `latest`.
3. On Darwin, strip quarantine (`xattr -d com.apple.quarantine`) and ad-hoc sign (`codesign -s - --force`).
4. **On any failure, print a notice and continue successfully.** Scaffolding must never fail because the binary is unavailable.

Target directory: `/usr/local/bin` if writable, else `${HOME}/.local/bin`.

## 3. Coverage Consequence

Removing `internal/track` (96.6%), `internal/mcp` (90.6%), and `internal/scaffold` (91.3%) removes well-covered packages. The remaining surface is `internal/validator` (93.5%), `cmd` (88.0%, minus the deleted command tests), and uncovered `main.go`.

Projected total stays above the 80% gate, but this must be measured rather than assumed — Phase 3 Task 3.1 verifies the real number before the gate is armed, and raises validator coverage if the projection is wrong.

## 4. Public Surface Change

`cooper mcp`, `cooper mcp install`, `cooper init`, and `cooper track *` are removed from a released v1.1.0 binary. This is breaking.

- Version bumps to **1.2.0** with an explicit removal list in the release notes.
- `README.md`'s CLI section is rewritten to the three surviving commands.
- Users with stale `cooper mcp` entries in editor MCP configs will see the server fail to start — a visible failure, preferable to silent misbehaviour. Removal instructions are documented.
