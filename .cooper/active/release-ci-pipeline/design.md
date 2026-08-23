# Design Document: Automated Multi-Platform Release CI/CD Pipeline & v1.0.0 Baseline

## 1. Pipeline Architecture & Workflow Hierarchy

The Cooper CI/CD release system is structured into two complementary GitHub Actions workflows:

```
Push to main / PR / Tag
         │
         ▼
┌─────────────────────────────────────────────────────────────┐
│                 .github/workflows/ci.yml                    │
│   ├── gofmt formatting check                                │
│   ├── go vet static analysis                                │
│   └── go test -v -coverprofile=coverage.out ./...           │
└──────────────────────────────┬──────────────────────────────┘
                               │ Success
                               ▼
┌─────────────────────────────────────────────────────────────┐
│               .github/workflows/release.yml                 │
│                                                             │
│   1. [auto-tag] (Only on push to main)                      │
│      ├── Extract version from cmd/version.go                │
│      └── Create & push v<version> tag if missing on origin  │
│                                                             │
│   2. [build-and-release] (Matrix: 5 targets)                │
│      ├── linux/amd64   -> dist/cooper-linux-x86_64          │
│      ├── linux/arm64   -> dist/cooper-linux-aarch64         │
│      ├── darwin/amd64  -> dist/cooper-darwin-x86_64         │
│      ├── darwin/arm64  -> dist/cooper-darwin-aarch64        │
│      └── windows/amd64 -> dist/cooper-windows-x86_64.exe    │
│      └── Inject ldflags (Version, Commit, Date)             │
│                                                             │
│   3. [publish-release] (gh release CLI)                     │
│      ├── Upload matrix assets to semantic release (v1.0.0)  │
│      └── Upload matrix assets to 'latest' rolling release   │
└─────────────────────────────────────────────────────────────┘
```

---

## 2. Detailed Workflow Specifications

### 2.1 Continuous Integration Workflow (`ci.yml`)
- **Triggers**: `push` on `main`, `pull_request` on `main`, and `workflow_call`.
- **Environment**:
  - `FORCE_JAVASCRIPT_ACTIONS_TO_NODE20: true`
  - `NODE_NO_WARNINGS: "1"`
  - Go Toolchain: `1.27.0` via `actions/setup-go@v5`
- **Steps**:
  1. `actions/checkout@v4`
  2. `actions/setup-go@v5` (cache: true, go-version: `1.27.0`)
  3. `gofmt -l .` verification (fails if unformatted code is detected)
  4. `go vet ./...`
  5. `go test -v -coverprofile=coverage.out ./...` and `go tool cover -func=coverage.out`

### 2.2 Release Workflow (`release.yml`)
- **Triggers**: `push` on `main`, `push` on tags `v*`.
- **Permissions**: `contents: write`
- **Jobs**:
  1. **`ci`**: Reusable workflow call to `./.github/workflows/ci.yml`.
  2. **`auto-tag`**:
     - Runs on `ubuntu-latest` if `github.ref == 'refs/heads/main'`.
     - Uses `actions/checkout@v4` with `fetch-depth: 0`.
     - Extracts version:
       ```bash
       VERSION=$(grep -E '^\s*Version\s*=' cmd/version.go | head -n 1 | sed -E 's/.*"([^"]+)".*/\1/')
       ```
     - Checks if `v${VERSION}` tag exists on origin via `git ls-remote`.
     - If not, tags and pushes `v${VERSION}` using `github-actions[bot]`.
  3. **`build-and-release`**:
     - Matrix strategy:
       - `goos: linux`, `goarch: amd64`, `binary: cooper-linux-x86_64`
       - `goos: linux`, `goarch: arm64`, `binary: cooper-linux-aarch64`
       - `goos: darwin`, `goarch: amd64`, `binary: cooper-darwin-x86_64`
       - `goos: darwin`, `goarch: arm64`, `binary: cooper-darwin-aarch64`
       - `goos: windows`, `goarch: amd64`, `binary: cooper-windows-x86_64.exe`
     - Compilation with ldflags:
       ```bash
       go build -ldflags="-s -w -X github.com/twoBoots/cooper/cmd.Version=${VERSION} -X github.com/twoBoots/cooper/cmd.Commit=${COMMIT} -X github.com/twoBoots/cooper/cmd.Date=${BUILD_DATE}" -o dist/${{ matrix.binary }} .
       ```
     - Uploads artifacts using `actions/upload-artifact@v4`.
  4. **`publish-release`**:
     - Downloads all built assets using `actions/download-artifact@v4`.
     - Publishes to GitHub Releases via `gh release upload --clobber` or `gh release create`.
     - Uploads to both `v${VERSION}` and `latest`.

---

## 3. Binary & Package Metadata Alignment

### 3.1 Version Definition in `cmd/version.go`
- Update `Version = "1.0.0"` in `cmd/version.go`.
- Ensure `Commit` and `Date` remain inject-ready for ldflags.

---

## 4. Verification & Testing

1. **Local Test & Formatting Verification**:
   - `gofmt -l .` returns 0 modified files.
   - `go vet ./...` succeeds cleanly.
   - `go test -v -cover ./...` passes all tests with >80% coverage.
2. **Static Cross-Compilation Dry Run**:
   - Test building `GOOS=linux GOARCH=amd64 go build -o dist/cooper-linux-x86_64 .`
   - Test building `GOOS=darwin GOARCH=arm64 go build -o dist/cooper-darwin-aarch64 .`
   - Test building `GOOS=windows GOARCH=amd64 go build -o dist/cooper-windows-x86_64.exe .`
3. **Workflow Syntax Linting**:
   - Verify YAML structure, action versions, environment variables, and job dependencies.
