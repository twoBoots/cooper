# Proposal: Bump Go Toolchain and Runtime Baseline to 1.27.0

## Rationale
The Go language toolchain evolves with performance improvements, standard library updates, and language features. Upgrading the Go runtime baseline and module directive from Go `1.22.0` to `1.27.0` modernizes the Cooper CLI toolchain, ensures future compatibility, and aligns repository definition files and scaffolding templates with current platform standards.

## User Benefit
- Developers and contributors building the Cooper CLI benefit from up-to-date Go compiler optimizations and toolchain capabilities.
- New projects scaffolded via Cooper will declare the modernized Go runtime baseline in their `.cooper/definition/tech-stack.md`.

## Scope & Boundaries
- **In Scope**:
  - Update `go.mod` to `go 1.27.0`.
  - Update `.cooper/definition/tech-stack.md` to specify `Go 1.27+`.
  - Update template assets in `templates/tech-stack.md` and `internal/scaffold/assets/templates/tech-stack.md`.
  - Verify all existing Go packages, commands (`cmd/`), MCP server (`internal/mcp/`), and test suites compile and pass.
- **Out of Scope**:
  - Changing external dependency versions (`github.com/spf13/cobra`, `github.com/twoBoots/bender`).
  - Rewriting existing Go CLI command logic beyond Go 1.27 compatibility.
