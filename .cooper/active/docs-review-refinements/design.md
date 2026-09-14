# Technical Design: Documentation Review Refinements

## Changes Overview
Update `docs/guide/workflow.md`:
1. Update Git Notes description:
   - From: `- **Git Notes**: Record task summary metadata (`git notes add -m`) against the task commit.`
   - To: `- **Git Notes**: Record task summary metadata (`git notes add -m`) against the task commit to preserve implementation context against drift.`
2. Update Quality Review intro:
   - From: `Before opening or merging a PR:`
   - To: `Before opening or merging a PR, act as a principal engineer to:`

## Quality Gates
- `npm test` passes 100%.
- `go run . validate` passes 100%.
- `npm run docs:build` compiles cleanly.
