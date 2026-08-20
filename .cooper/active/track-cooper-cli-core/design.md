# Technical Design: Core Go CLI & SDD Validator

- **Track ID**: `track-cooper-cli-core`
- **Architectural Reference**: [`rfc-cooper-cli-mcp`](.cooper/active/rfc-cooper-cli-mcp/rfc.md)

---

## 1. Architecture & Package Structure

```
cooper/
├── cmd/
│   ├── root.go                 # Root Cobra command, global flags (--non-interactive, --verbose)
│   ├── version.go              # Version command (dynamic version injection via ldflags)
│   ├── init.go                 # 'cooper init' command handler
│   ├── validate.go             # 'cooper validate' / 'cooper lint' command handler
│   └── track.go                # 'cooper track [new|status|checkpoint|close]' command handler
├── internal/
│   ├── validator/              # High-speed SDD validation engine
│   │   ├── spec_linter.go      # GIVEN/WHEN/THEN syntax & +/- delta validation
│   │   ├── metadata.go         # metadata.json and tracks.md schema & parity validation
│   │   ├── link_auditor.go     # Markdown link integrity & external tool attribution check
│   │   └── validator_test.go   # Unit tests (>80% coverage)
│   ├── scaffold/               # Scaffolding & migration engine
│   │   ├── init.go             # Project init & Conductor/OpenSpec migration
│   │   ├── embedded.go         # go:embed templates & skills filesystem
│   │   └── scaffold_test.go    # Unit tests (>80% coverage)
│   └── track/                  # Track orchestration & Troop worktree interface
│       ├── worktree.go         # git agent-start / git worktree wrapper
│       ├── checkpoint.go       # git notes & phase checkpoint generator
│       ├── registry.go         # tracks.md parser & updater
│       └── track_test.go       # Unit tests (>80% coverage)
├── main.go                     # Binary entrypoint executing cmd.Execute()
└── go.mod                      # Module definition: github.com/twoBoots/cooper
```

---

## 2. Command Surface Specification

### 2.1 `cooper version`
- Outputs current CLI version, Git commit hash, and build timestamp.
- JSON output support with `--json`.

### 2.2 `cooper init`
- Flags: `--force` (overwrite existing files), `--non-interactive` (skip prompts).
- Scaffolds `.cooper/`, `.agents/skills/`, and `AGENTS.md`.
- Automatically migrates existing `.conductor/` or `openspec/` directories if detected.

### 2.3 `cooper validate` (alias: `cooper lint`)
- Validates:
  1. Spec Delta syntax in all active tracks (`.cooper/active/*/spec-deltas/**/*.md`).
  2. Living spec formatting in `.cooper/specs/**/*.md`.
  3. JSON schemas of `.cooper/active/*/metadata.json`.
  4. Outbound repository link compliance for foundational dependencies.
- Exit code `0` on clean validation, exit code `1` on lint errors with file and line annotations.

### 2.4 `cooper track`
- `cooper track new <track_id> [--type feature|bugfix|chore] [--title "Title"]`
- `cooper track status [<track_id>]`
- `cooper track checkpoint <phase_number>`
- `cooper track close <track_id>`

---

## 3. Data Structures & Validation Contracts

```go
type TrackMetadata struct {
    TrackID     string `json:"track_id"`
    Title       string `json:"title"`
    Type        string `json:"type"`   // "feature" | "bugfix" | "chore" | "rfc"
    Status      string `json:"status"` // "new" | "in_progress" | "completed" | "approved"
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at,omitempty"`
}

type ValidationError struct {
    File    string `json:"file"`
    Line    int    `json:"line"`
    Message string `json:"message"`
    Rule    string `json:"rule"`
}
```

---

## 4. Test Strategy & TDD Workflow
- Unit testing with table-driven tests for every validator rule.
- Mock Git executor for testing worktree and checkpoint operations without external subprocess failures.
- Minimum 80% test coverage enforced via `go test -coverprofile=coverage.out ./...`.
