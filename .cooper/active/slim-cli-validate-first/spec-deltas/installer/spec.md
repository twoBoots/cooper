# Spec Delta: Installer & Repo Structure

## Capability: installer

## Requirements

### Requirement: Single Authoritative Scaffolding Mechanism
+ `install.sh` SHALL be the sole scaffolding mechanism for Cooper projects, and the repository MUST NOT retain a second scaffolder or a duplicate copy of the agent skill tree.

#### Scenario: Exactly One Authoritative Skill Tree Exists
+ - GIVEN the Cooper repository
+ - WHEN the agent skill sources are enumerated
+ - THEN exactly one authoritative source copy of each `cooper-*` SKILL.md MUST exist, under `skills/`
+ - AND no embedded duplicate MUST exist under `internal/`
+ - AND no `SKILL.md` MUST exist outside `skills/` and `.agents/skills/`.

#### Scenario: Installed Copy Matches Its Source
+ - GIVEN Cooper's own repository, which installs its skills into `.agents/skills/` exactly as a consumer project does
+ - WHEN the installed copy is compared against its `skills/` source
+ - THEN each `.agents/skills/cooper-*/SKILL.md` MUST be byte-identical to the corresponding file under `skills/`
+ - AND any divergence MUST fail validation, because two copies drifting apart unnoticed is the defect this capability exists to prevent.

#### Scenario: Newer Skill Content Is Preserved On Consolidation
+ - GIVEN two skill trees that have diverged, one carrying newer mandates than the other
+ - WHEN the trees are consolidated into a single source of truth
+ - THEN the surviving tree MUST retain the newer content, including the Interactive Question Protocol and Native File Tools Mandate sections
+ - AND no mandate present in either tree before consolidation MAY be lost.

#### Scenario: Scaffolded Project Satisfies Its Own Agent Guidelines
+ - GIVEN a user scaffolding a fresh repository via `install.sh`
+ - WHEN scaffolding completes
+ - THEN every path referenced by the installed `AGENTS.md` MUST exist, including `.cooper/COOPER.md` and `.cooper/TROOP.md`
+ - AND the Troop git aliases `git agent-start`, `git troop`, and `git agent-stop` MUST be registered
+ - AND `.gitignore` MUST exclude `.worktrees/`.

### Requirement: Optional Non-Fatal Binary Installation
+ `install.sh` SHALL offer the compiled `cooper` binary where obtainable, and MUST complete successfully when it cannot be obtained.

#### Scenario: Compile From Local Clone
+ - GIVEN `install.sh` executing from a local clone containing `main.go`
+ - AND `go` is present on `PATH`
+ - WHEN the binary installation step runs
+ - THEN the installer SHALL compile `cooper` and place it in the first writable directory of `/usr/local/bin` then `${HOME}/.local/bin`.

#### Scenario: Download Pre-Built Release Asset
+ - GIVEN `install.sh` executing without a local Go toolchain
+ - WHEN the binary installation step runs
+ - THEN the installer SHALL detect the host operating system and architecture
+ - AND SHALL attempt to download the matching `cooper-${OS}-${ARCH}` asset from the latest GitHub Release.

#### Scenario: macOS Gatekeeper Handling
+ - GIVEN a binary compiled or downloaded on Darwin
+ - WHEN the binary is placed on disk
+ - THEN the installer SHALL strip the quarantine attribute and apply an ad-hoc code signature.

#### Scenario: Graceful Zero-Binary Fallback
+ - GIVEN an air-gapped environment, a rate-limited GitHub API, or an unwritable target directory
+ - WHEN the binary installation step fails
+ - THEN the installer MUST print an informational notice
+ - AND MUST complete scaffolding successfully with a zero exit status
+ - AND the resulting `.cooper/` and `.agents/skills/` workspace MUST remain fully usable without the binary.
