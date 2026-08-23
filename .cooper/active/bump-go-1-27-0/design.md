# Technical Design: Bump Go Toolchain and Runtime Baseline to 1.27.0

## Overview
This change bumps the Go toolchain directive in `go.mod` to `1.27.0` and updates all project definition documents and embedded scaffolding templates to reflect `Go 1.27+` as the supported runtime baseline.

## Technical Architecture & Impact

### 1. Root Module Directive (`go.mod`)
- Update `go 1.22.0` to `go 1.27.0`.
- Verify `go.sum` consistency and tidy dependencies if necessary (`go mod tidy`).

### 2. Project Definition & Scaffolding Templates
- **Framework Tech Stack Definition**: Update `.cooper/definition/tech-stack.md` line 5 from `Go 1.22+` to `Go 1.27+`.
- **Root Template Asset**: Update `templates/tech-stack.md` line 5 from `Go 1.22+` to `Go 1.27+`.
- **Embedded Scaffold Template Asset**: Update `internal/scaffold/assets/templates/tech-stack.md` line 5 from `Go 1.22+` to `Go 1.27+`.

### 3. Verification Strategy
- Run `go test -v ./...` across all packages (`cmd`, `internal/mcp`, `internal/scaffold`).
- Ensure code coverage exceeds >80%.
- Run `cooper validate` or `go run main.go validate` to ensure spec and link compliance.
