# Implementation Plan: Relocate TROOP.md to .cooper/ and align repository structure

## Phase 1: Repository Structure Realignment
- [x] Task: Migrate `cooper/workflow.md` to `.cooper/definition/workflow.md`
- [x] Task: Remove obsolete `cooper/` directory from repository
- [x] Task: Remove obsolete `conductor/` directory from repository

## Phase 2: Installer Update (`install.sh`)
- [x] Task: Update Troop setup in `install.sh` to move `TROOP.md` into `.cooper/TROOP.md`
- [x] Task: Update Cooper workflow fetch in `install.sh` to target `.cooper/definition/workflow.md`
- [x] Task: Remove copying of `workflow.md` into `conductor/workflow.md`

## Phase 3: Templates and Documentation Updates
- [x] Task: Update `AGENTS.template.md` to reference `.cooper/TROOP.md`
- [x] Task: Update `README.md` tree structure and descriptions
- [x] Task: Update `docs/INSTALL.md` and `docs/openspec-vs-conductor-comparison.md`
- [x] Task: Update `.cooper/definition/workflow.md` diagrams and paths

## Phase 4: Verification & Checkpoint
- [x] Task: Verify installer locally with a clean git workspace test
- [x] Task: Phase Verification & Checkpoint
