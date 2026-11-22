# Technical Design: Bump Cooper CLI Patch Version to v1.2.1

## Architecture Overview
The Cooper release pipeline (`.github/workflows/release.yml`) inspects `Version` in `cmd/version.go` on pushes to `main`. When `v<Version>` does not yet exist on origin, the `auto-tag` job creates and pushes the annotated tag `v1.2.1` and publishes multi-platform release binaries for both `v1.2.1` and `latest`.

## Component Changes
1. **`cmd/version.go`**:
   - Update `Version = "1.2.0"` to `Version = "1.2.1"`.
2. **`cmd/version_test.go`**:
   - Add `TestDefaultVersionConstant` (and ensure tests restore `Version` after mutation) to assert the default compiled version is `"1.2.1"`.
