# Spec Delta: Cooper Go CLI & MCP Server Engine

## Capability: cli

## Requirements

### Requirement: Embedded Stdio Model Context Protocol (MCP) Server
- The Cooper CLI SHALL provide an `mcp` command that starts a JSON-RPC 2.0 MCP server over stdio.

#### Scenario: MCP Server Initialization
- - GIVEN an AI assistant starting `cooper mcp`
- - WHEN the client sends an `initialize` JSON-RPC request
- - THEN the server responds with server name `cooper-mcp`, server version, and tool capabilities.

#### Scenario: Expose Cooper SDD Tools
- - GIVEN an initialized Cooper MCP server
- - WHEN the client sends a `tools/list` request
- - THEN the server returns `cooper_get_version`, `cooper_init_project`, `cooper_track_create`, `cooper_track_status`, `cooper_validate`, and `cooper_self_update`.

### Requirement: AI Assistant MCP Client Configuration Installer
- The Cooper CLI SHALL provide an `mcp install` command to automatically register Cooper into editor MCP configuration files.

#### Scenario: Auto-Detect and Configure Clients
- - GIVEN supported AI assistant configuration files (Cursor, Antigravity, Claude Desktop, Claude Code, Windsurf, VS Code)
- - WHEN running `cooper mcp install`
- - THEN the CLI safely merges `cooper` (`command: "cooper"`, `args: ["mcp"]`) into the detected client config files.

### Requirement: Validate-First Command Surface
+ The Cooper CLI SHALL expose only `validate`, `update`, and `version` as commands, and MUST NOT provide project scaffolding, track orchestration, or Model Context Protocol server functionality.

#### Scenario: Removed Commands Are Absent
+ - GIVEN a user running `cooper --help`
+ - WHEN the CLI renders its command list
+ - THEN the output MUST contain `validate`, `update`, and `version`
+ - AND the output MUST NOT contain `init`, `track`, or `mcp`.

#### Scenario: Scaffolding Delegated To Installer
+ - GIVEN a user wishing to scaffold or migrate a repository into Cooper
+ - WHEN the user consults the documented installation path
+ - THEN `install.sh` SHALL be the single supported scaffolding mechanism
+ - AND the CLI MUST NOT offer a competing `init` command.

#### Scenario: Track Orchestration Delegated To Skills
+ - GIVEN a user wishing to create, inspect, checkpoint, or close a track
+ - WHEN the user invokes Cooper's workflow
+ - THEN the `cooper-new-track`, `cooper-implement`, and `cooper-status` agent skills SHALL be the single supported mechanism
+ - AND the CLI MUST NOT provide `track` subcommands that duplicate them.

### Requirement: Prohibition On Unverified Checkpoint Attestations
+ Cooper MUST NOT emit a Git Note asserting that automated tests passed or that a user approved a verification, unless that test run and that approval actually occurred.

#### Scenario: No Programmatic Fabrication Of Verification Records
+ - GIVEN any Cooper component capable of attaching a checkpoint Git Note
+ - WHEN the note body would assert an automated test result or a user approval
+ - THEN the asserted outcome MUST be derived from an actual executed test invocation and an actual recorded user confirmation
+ - AND hardcoded verification text such as `Automated Tests: PASSED` MUST NOT be written without a corresponding verified result.

#### Scenario: No Fabricated Attestation Templates In Agent Instructions
+ - GIVEN a Cooper agent skill or workflow definition instructing an agent to attach a checkpoint Git Note
+ - WHEN that instruction supplies the note body as a template
+ - THEN the template MUST direct the agent to record the actual test command, its real outcome, and the user's actual recorded response
+ - AND the template MUST NOT supply a pre-filled passing result or approval for the agent to copy verbatim.

### Requirement: SDD Repository Syntax Validation
+ The Cooper CLI SHALL additionally audit inline-code repository path references in markdown, so that documentation citing a missing file fails validation.

#### Scenario: Detect Dangling Backticked Path Reference
+ - GIVEN a markdown file citing a repository path inside an inline-code span, such as `.cooper/COOPER.md`
+ - AND no file exists at that resolved path
+ - WHEN running `cooper validate`
+ - THEN the validator MUST report a `link/code-path-exists` violation naming the file and line
+ - AND the command MUST exit with a non-zero status.

#### Scenario: Ignore Non-Path Inline Code
+ - GIVEN a markdown file containing inline-code spans that are shell commands, glob patterns, flags, or URLs
+ - WHEN running `cooper validate`
+ - THEN the validator MUST NOT report `link/code-path-exists` violations for those spans.

### Requirement: Automated Multi-Platform Release CI/CD Pipeline
+ The Continuous Integration pipeline SHALL additionally enforce Cooper SDD specification validity and the project's documented minimum test coverage, failing the build when either is violated.

#### Scenario: CI Fails On Specification Violation
+ - GIVEN a pull request introducing a malformed living spec, spec delta, track metadata file, or dangling documentation link
+ - WHEN GitHub Actions CI executes
+ - THEN the pipeline MUST run `cooper validate`
+ - AND the job MUST fail with the reported violations.

#### Scenario: CI Fails Below Coverage Threshold
+ - GIVEN a pull request reducing total statement coverage below 80%
+ - WHEN GitHub Actions CI executes the coverage step
+ - THEN the pipeline MUST report the measured percentage
+ - AND the job MUST exit non-zero.

#### Scenario: Enforcement Is Not Injected Into Consumer Projects
+ - GIVEN a project scaffolded by Cooper's installer
+ - WHEN the scaffolding completes
+ - THEN Cooper MUST NOT install mandatory validation or coverage gates into that project's own CI configuration
+ - AND such gates SHALL remain opt-in, preserving Cooper's suggestive rather than prescriptive posture.
