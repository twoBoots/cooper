# Design: Mandate Interactive Question & Native File Tool Calls

## Architecture & Interaction Standards

### 1. Interactive Questioning Protocol
All Cooper skills (`cooper-setup`, `cooper-rfc`, `cooper-new-track`, `cooper-implement`, `cooper-review`, `cooper-status`) and `.cooper/definition/workflow.md` will standardize on the following operational instruction:

- **Interactive Question Tools (Mandatory)**: When presenting single-choice or multiple-choice options or seeking user confirmation, agents MUST invoke available interactive question tools (e.g. `ask_question`) rather than printing text choice lists in chat.
- **Text Chat Fallback**: Formatting questions and option lists as plain text in chat is strictly reserved as a fallback when no interactive question tool is available in the agent runtime environment.
- **Context-Aware Recommendations**: Prefix preferred choices with `(Recommended: <explanation>)`.

### 2. Native File Tools Protocol
- **Native File Operations (Mandatory)**: File viewing, creation, and modification MUST be performed using dedicated native file tools (`view_file`, `write_to_file`, `replace_file_content`) when available.
- **Prohibition on Shell Piping for File Edits**: Agents MUST NOT use shell commands with heredocs (`cat << 'EOF' > file`), pipes, or output redirections (`echo "..." > file`) to create or edit files when native file tools are present in the environment.

## Target Modifications
1. **Agent Skills**:
   - `.agents/skills/cooper-new-track/SKILL.md`
   - `.agents/skills/cooper-rfc/SKILL.md`
   - `.agents/skills/cooper-review/SKILL.md`
   - `.agents/skills/cooper-implement/SKILL.md`
   - `.agents/skills/cooper-setup/SKILL.md`
   - `.agents/skills/cooper-status/SKILL.md`
2. **Scaffold Skill Assets & Templates**:
   - `internal/scaffold/assets/skills/cooper-new-track/SKILL.md`
   - `internal/scaffold/assets/skills/cooper-rfc/SKILL.md`
   - `internal/scaffold/assets/skills/cooper-review/SKILL.md`
   - `internal/scaffold/assets/skills/cooper-implement/SKILL.md`
   - `internal/scaffold/assets/skills/cooper-setup/SKILL.md`
   - `internal/scaffold/assets/skills/cooper-status/SKILL.md`
   - `internal/scaffold/assets/AGENTS.template.md`
3. **Workflow & Guidelines**:
   - `AGENTS.md`
   - `.cooper/definition/workflow.md`
4. **Living Capability Spec Delta**:
   - `.cooper/active/mandate-interactive-question-tools/spec-deltas/documentation/spec.md`
