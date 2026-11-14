# Technical Design: Living Files Stand Alone

## Architecture Overview
Enforces self-contained living documentation through workflow rules, capability specs, and agent skill instructions - without Go validator regex rules or `.cooperrc` configuration.

```
.cooper/definition/workflow.md      <-- Guiding Principle + Track Finalization cleanup
templates/workflow.md               <-- Synchronised template for initialised repos
.cooper/specs/documentation/spec.md <-- Canonical Requirement & Scenarios
        |
        +-- skills/cooper-new-track/SKILL.md   (keep delta prose self-contained)
        +-- skills/cooper-rfc/SKILL.md         (keep delta prose self-contained; RFCs frozen on approval/archive)
        +-- skills/cooper-implement/SKILL.md   (strip forbidden refs on delta merge / definition edits)
        +-- skills/cooper-review/SKILL.md      (flag violations in living files & delta merges)
        +-- skills/cooper-setup/SKILL.md       (audit existing definition/ and specs/ in brownfield mode)
                |
                +-- synced byte-for-byte to .agents/skills/<skill>/SKILL.md
```

---

## Component Design

### 1. Workflow Policy & Template (`.cooper/definition/workflow.md` & `templates/workflow.md`)
- **Guiding Principles:** Add `**Living Files Stand Alone:**` covering `.cooper/definition/` and `.cooper/specs/` (repo root and barrels), exempting `.cooper/active/` and `.cooper/archive/`, forbidding ticket/issue/PR/track numbers, decision/open-question markers, and bare section numbers, and defining the two non-requirement exceptions (Rationale Links and Open Known Gaps).
- **Track Finalization (Merge & Spec Delta Integration):** Strip ticket numbers, decision markers, and bare section numbers (replacing section numbers with exact heading names) when merging `spec-deltas/` into `.cooper/specs/`.

### 2. Living Capability Specification (`.cooper/specs/documentation/spec.md`)
Add `### Requirement: Living Files Stand Alone` with five scenarios:
1. Reject ticket/issue/PR/track numbers, decision/open-question markers, and bare section numbers in `.cooper/definition/` and `.cooper/specs/` while exempting `.cooper/active/` and `.cooper/archive/`.
2. Allow rationale links outside requirement text citing a frozen archived track, approved/archived RFC, or frozen macro contract by file path and heading name.
3. Allow open known-gap entries outside requirement text to cite a tracking ticket until resolved, then remove the entry.
4. Strip planning provenance during Spec Delta finalization merges into `.cooper/specs/`.
5. Flag forbidden references in existing `definition/` and `specs/` files during `cooper-setup` brownfield audits.

### 3. Skill Instructions (`skills/cooper-*/SKILL.md` & `.agents/skills/cooper-*/SKILL.md`)
Edit `skills/` and mirror byte-for-byte to `.agents/skills/`:
- **`cooper-new-track/SKILL.md` (`### 3.3 Draft Track Artifacts`):** Keep `spec-deltas/<capability>/spec.md` free of ticket numbers, decision markers, and bare section numbers; keep provenance in `proposal.md` and `design.md`.
- **`cooper-rfc/SKILL.md` (`### 4.2 spec-deltas/<capability>/spec.md`):** Keep RFC spec deltas free of forbidden tokens; note that approved/archived RFCs and frozen macro contracts may be cited outside requirement text by file path and heading name.
- **`cooper-implement/SKILL.md` (`## 4. Track Finalization & Handshake`):** Strip forbidden tokens before merging Spec Deltas into `.cooper/specs/` or editing `.cooper/definition/`.
- **`cooper-review/SKILL.md` (`### 2.2 Verification Checks`):** Check touched `.cooper/definition/`, `.cooper/specs/`, and `spec-deltas/` files for forbidden tokens.
- **`cooper-setup/SKILL.md` (`### 1.2 Detect Project State & Maturity`):** Audit existing `.cooper/definition/` and `.cooper/specs/` in brownfield mode.

### 4. Validator Alignment
- Preserve `"Interactive Question Protocol"` and `"Native File Tools Mandate"` in all skills.
- Use `### N.M` (never `## N.M`) for skill subsection headings.
- Keep `skills/<skill>/SKILL.md` and `.agents/skills/<skill>/SKILL.md` byte-for-byte identical.
- Use `<...>` placeholders in illustrative backticked paths so `internal/validator/code_path_auditor.go` (`link/code-path-exists`) passes.

---

## Quality Gates
- `go test ./...` (`cmd/skills_test.go`, `internal/validator/...`) and `go run . validate` pass.
- `.cooper/definition/*` and `.cooper/specs/*/spec.md` contain zero forbidden references.
