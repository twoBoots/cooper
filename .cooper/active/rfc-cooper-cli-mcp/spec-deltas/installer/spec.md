# Capability Specification Delta: Installer & Repo Structure

## Requirements

### Requirement: Optional CLI Binary Installation
The installer SHALL optionally prompt the user to install the pre-compiled `cooper` Go CLI binary to the user's path.

#### Scenario: User Accepts CLI Installation
- GIVEN the installer executes in an interactive terminal
- WHEN prompted to install the `cooper` CLI binary and the user accepts
- THEN the installer MUST detect host OS and architecture (e.g., `darwin/arm64`, `linux/amd64`), download the matching pre-compiled release binary, place it into `~/.local/bin/cooper` (or `/usr/local/bin`), and verify executable permissions.

#### Scenario: Zero-Binary Fallback on Decline or Offline
- GIVEN the installer executes in a headless, air-gapped, or offline environment, or the user declines the CLI prompt
- WHEN resolving installation dependencies
- THEN the installer MUST complete successfully in file-only mode, scaffolding `.cooper/`, `.agents/skills/`, and `AGENTS.md` without requiring the binary.

### Requirement: Upstream Version Manifest Generation
The installer SHALL generate a release manifest `.cooper/manifest.json` recording the upstream base version and file hashes to enable future 3-way updates.

#### Scenario: Create Version Manifest on Setup
- GIVEN the installer finishes scaffolding Cooper infrastructure
- WHEN writing configuration state
- THEN it MUST write `.cooper/manifest.json` recording the current version tag and template SHA256 hashes.
