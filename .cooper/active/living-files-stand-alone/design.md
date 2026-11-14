# Technical Design: Living Files Stand Alone

## Architecture Overview
This feature establishes self-contained living documentation through policy rules, specification requirements, agent skill instructions, and verification tests—without adding Go validator regex rules or `.cooperrc` configuration knobs.

```
.cooper/definition/workflow.md   <── Guiding Principle + Track Finalization cleanup
templates/workflow.md            <── Synchronized template for initialized repos
.cooper/specs/documentation/spec.md <── Canonical Requirement & Scenarios
        │
        ├── skills/cooper-new-track/SKILL.md   (keep delta prose self-contained)
        ├── skills/cooper-rfc/SKILL.md         (keep delta prose self-contained; RFCs frozen on approval/archive)
        ├── skills/cooper-implement/SKILL.md   (strip forbidden refs on delta merge / definition edits)
        ├── skills/cooper-review/SKILL.md      (flag violations in living files & delta merges)
        └── skills/cooper-setup/SKILL.md       (audit existing definition/ and specs/ in brownfield mode)
                │
                └── synced byte-for-byte to .agents/skills/<skill>/SKILL.md
```

---

## Component Design

### 1. Workflow Policy & Template (`.cooper/definition/workflow.md` & `templates/workflow.md`)
- **Guiding Principles:** Add `**Living Files Stand Alone:**` stating:
  - Applies to all files in `.cooper/definition/` and `.cooper/specs/` (at repository root and inside any barrel). `.cooper/active/` and `.cooper/archive/` are exempt point-in-time working records.
  - Living files SHALL NOT contain ticket/issue/PR/track numbers, decision or open-question markers (`D14`, `DEC-1`, `OQ-4`), or bare section numbers (`§4.2`, `section 3`), including self-references (use exact heading names instead).
  - Outside requirement text (requirements and GIVEN/WHEN/THEN scenario steps):
    - Rationale notes may link to a frozen archived track document, an approved or archived RFC, or a frozen macro contract by file path + heading name.
    - Open known-gap notes may cite their tracking ticket until resolved, after which the entry is removed.
- **Track Finalization (Merge & Spec Delta Integration):** Explicitly require agents/maintainers merging `spec-deltas/` into `.cooper/specs/<capability>/spec.md` to strip planning provenance (ticket numbers, decision markers, and bare section numbers, replacing section numbers with heading names) before committing the updated living spec.

### 2. Living Capability Specification (`.cooper/specs/documentation/spec.md`)
Add `### Requirement: Living Files Stand Alone` with five normative scenarios:
1. Rejecting ticket/issue/PR/track numbers, decision/open-question markers, and bare section numbers in `.cooper/definition/` and `.cooper/specs/` (root and barrels) while exempting `.cooper/active/` and `.cooper/archive/`.
2. Permitting rationale links outside requirement text that cite a frozen archived track, approved/archived RFC, or frozen macro contract using file path + heading name.
3. Permitting open known-gap entries outside requirement text to cite a tracking ticket until shipped, and removing the entry once resolved.
4. Requiring Spec Delta finalization merges to strip planning provenance before writing to `.cooper/specs/`.
5. Requiring `cooper-setup` brownfield audits to detect and report forbidden references in existing `definition/` and `specs/` files.

### 3. Skill Instructions (`skills/cooper-*/SKILL.md` & `.agents/skills/cooper-*/SKILL.md`)
All edits are made first in `skills/` and mirrored byte-for-byte to `.agents/skills/` so `TestSkills_InstalledCopyMatchesSource` passes.
- **`cooper-new-track/SKILL.md` (`### 3.3 Draft Track Artifacts`):** Instruct agents writing `spec-deltas/<capability>/spec.md` (and any edits to `definition/`) to write self-contained requirement and scenario prose free of ticket numbers, decision markers (`D14`, `DEC-1`, `OQ-4`), and bare section numbers, using heading names instead; keep provenance inside `proposal.md`/`design.md` or cite frozen RFCs/archived tracks by path + heading only outside requirement text.
- **`cooper-rfc/SKILL.md` (`### 4.2 spec-deltas/<capability>/spec.md`):** Instruct agents that RFC spec deltas targeting living specs follow **Living Files Stand Alone**, and note that approved/archived RFCs (like frozen macro contracts) may be cited from outside requirement text by file path + heading name.
- **`cooper-implement/SKILL.md` (`## 4. Track Finalization & Handshake` + Operational Standards):** Require stripping any ticket numbers, decision markers, or bare section numbers before merging Spec Deltas into living `.cooper/specs/` or modifying `.cooper/definition/`.
- **`cooper-review/SKILL.md` (`### 2.2 Verification Checks`):** Add a verification check ensuring touched `.cooper/definition/` or `.cooper/specs/` files (and outgoing Spec Deltas ready to merge) contain no banned ticket/decision/section-number references inside requirement text or prose outside the two bounded exceptions.
- **`cooper-setup/SKILL.md` (`### 1.2 Detect Project State & Maturity`):** In brownfield/audit discovery mode, inspect existing `.cooper/definition/` and `.cooper/specs/` (root and barrels) and flag any ticket numbers, decision/open-question markers, or bare section numbers for cleanup.

### 4. Validator & Test Alignment
- **`cmd/skills_test.go` Constraints:**
  - Preserve `"Interactive Question Protocol"` and `"Native File Tools Mandate"` in all skills.
  - Use `### N.M` (never `## N.M`) for subsection headings inside skills.
  - Keep `skills/<skill>/SKILL.md` and `.agents/skills/<skill>/SKILL.md` byte-for-byte identical.
  - Ensure no illustrative example path without `<...>` / `{...}` placeholders triggers `internal/validator/code_path_auditor.go` (`link/code-path-exists`).
- **`tests/docs-workflow.test.mjs`:** Add Node test assertions verifying that `.cooper/definition/workflow.md`, `templates/workflow.md`, `.cooper/specs/documentation/spec.md`, and the five updated skills contain the **Living Files Stand Alone** principle, exceptions, finalization cleanup, and brownfield audit guidance, and that living files in `.cooper/definition/` and `.cooper/specs/` are clean of banned patterns.

---

## Quality Gates
- `CI=true npm test` (`tests/docs-workflow.test.mjs`) passes.
- `go test ./...` (`cmd/skills_test.go`, `internal/validator/...`) and `go run . validate` pass.
- Manual audit confirms `.cooper/definition/*` and `.cooper/specs/*/spec.md` contain zero forbidden references.
