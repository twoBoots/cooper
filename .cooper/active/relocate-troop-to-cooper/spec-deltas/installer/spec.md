# Capability Spec Delta: Installer & Repo Structure

## Added Requirements
+ GIVEN Cooper installer executes in a project
+ WHEN the Troop foundation setup completes
+ THEN `TROOP.md` MUST be placed or moved into `.cooper/TROOP.md` instead of remaining at the root level.

+ GIVEN Cooper installer fetches workflow specification
+ WHEN resolving Cooper repository assets
+ THEN `.cooper/definition/workflow.md` MUST be fetched and placed into `.cooper/definition/workflow.md`.

## Removed / Modified Requirements
- GIVEN Cooper repository structure
- WHEN inspecting workflow specifications
- THEN `cooper/workflow.md` is removed and replaced by `.cooper/definition/workflow.md`.
