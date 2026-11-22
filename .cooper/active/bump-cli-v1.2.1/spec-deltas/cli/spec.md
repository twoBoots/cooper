# Spec Delta: Cooper Go CLI & MCP Server Engine

## Capability: cli

## Requirements

### Requirement: Automated Multi-Platform Release CI/CD Pipeline
+ The Cooper CLI SHALL maintain its default semantic version string in `cmd/version.go` (`1.2.1`) to drive automated semantic Git tagging and release asset publishing when merged to `main`.

#### Scenario: Default CLI Version Reflects Patch Release v1.2.1
+ - GIVEN a newly built Cooper CLI binary without custom `-ldflags` overrides
+ - WHEN a user runs `cooper version`
+ - THEN the CLI MUST report version `1.2.1`.
