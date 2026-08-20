# Capability Specification Delta: Installer & Repo Structure

## Requirements

### Requirement: Tiered CLI Binary Registration & Fallback
The installer SHALL execute a 3-tier binary installation strategy (Local Compilation -> GitHub Release Download -> Zero-Binary Fallback) without hard failures.

#### Scenario: Tier 1 Local Source Build
- GIVEN the installer executes within a local source clone containing `main.go`
- WHEN `go` is found in the user's `PATH`
- THEN the installer MUST compile the binary locally, apply macOS codesign/quarantine fixes if on Darwin, and register the binary to `/usr/local/bin/cooper` (or `~/.local/bin/cooper`).

#### Scenario: Tier 2 Release Binary Download
- GIVEN no local source build is available
- WHEN resolving installation dependencies
- THEN the installer MUST detect host OS/arch and attempt to download the pre-compiled binary from GitHub Releases into the target binary directory.

#### Scenario: Tier 3 Zero-Binary Graceful Fallback
- GIVEN the installer executes in an offline or air-gapped environment where compilation and download fail
- WHEN resolving installation dependencies
- THEN the installer MUST complete successfully in zero-binary file mode, scaffolding `.cooper/`, `.agents/skills/`, and `AGENTS.md` without throwing errors.

### Requirement: Upstream Version Manifest Generation
The installer SHALL generate a release manifest `.cooper/manifest.json` recording the upstream base version and file hashes to enable future 3-way updates.

#### Scenario: Create Version Manifest on Setup
- GIVEN the installer finishes scaffolding Cooper infrastructure
- WHEN writing configuration state
- THEN it MUST write `.cooper/manifest.json` recording the current version tag and template SHA256 hashes.
