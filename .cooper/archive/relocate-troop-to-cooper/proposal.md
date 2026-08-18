# Track Proposal: Relocate TROOP.md into .cooper/ and align repository structure

## Rationale & Goal
When Cooper is installed in a target project, `TROOP.md` is currently placed in the project's root folder. To keep project root directories clean and consolidate all framework documentation and guidelines under `.cooper/`, `TROOP.md` should be relocated to `.cooper/TROOP.md`.

Additionally, the `twoBoots/cooper` repository currently maintains `cooper/workflow.md` in a `cooper/` directory rather than `.cooper/`. Migrating `cooper/` to `.cooper/` aligns this repository with the exact structure expected in projects utilizing Cooper.

## Scope of Changes
1. **Repository Structure**: Move `cooper/workflow.md` to `.cooper/definition/workflow.md` and remove `cooper/`.
2. **Installer (`install.sh`)**:
   - Ensure base `.cooper/` directory exists before Troop install.
   - Relocate `TROOP.md` to `.cooper/TROOP.md` upon completion of the Troop installer step.
   - Update remote / local file resolution paths from `cooper/workflow.md` to `.cooper/definition/workflow.md`.
3. **Templates & Rules**:
   - Update `AGENTS.template.md` to reference `.cooper/TROOP.md`.
4. **Documentation**:
   - Update `README.md`, `docs/INSTALL.md`, `docs/openspec-vs-conductor-comparison.md`, and workflow files.
