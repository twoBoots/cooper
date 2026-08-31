# Proposal: Mandate Interactive Question & Native File Tool Calls

## Context & Motivation
Cooper skills and framework guidelines guide AI agents through Spec-Driven Development (SDD) lifecycles. However, two common prompt ambiguities cause AI agents to deviate from runtime capabilities:
1. **Interactive Questions**: Guidelines instructing agents to *"ask questions strictly one at a time in text chat"* cause agents to print markdown bullet lists instead of invoking runtime interactive question tools (such as `ask_question`).
2. **File Manipulation**: Agents frequently use shell commands with stream editors, heredocs, pipes, and stream redirections (`sed`, `awk`, `cat << 'EOF' > ...`, `echo "..." > ...`) instead of dedicated native file manipulation tools (`write_to_file`, `replace_file_content`, `view_file`), bypassing IDE safety checks, permissions, and diff integrations.

## Objectives & Scope
- **Mandate Interactive Question Tools**: Explicitly instruct agents that whenever presenting single-choice or multiple-choice questions or confirmations, they MUST invoke runtime interactive question tools (e.g. `ask_question`). Plain text lists in chat are strictly a fallback when no such tool exists in the environment.
- **Mandate Native File Tools**: Explicitly prohibit shell stream editors (`sed`, `awk`), pipes, heredocs, and shell stream redirection for creating or editing files. Mandate the use of native file tools (`write_to_file`, `replace_file_content`, `view_file`) whenever available.
- **Comprehensive Coverage**: Update active skills (`.agents/skills/`), embedded scaffold skill assets (`internal/scaffold/assets/skills/`), root instructions (`AGENTS.md`, `internal/scaffold/assets/AGENTS.template.md`), and workflow definition (`.cooper/definition/workflow.md`).

## Impact & Value
- Standardizes human-in-the-loop UX across modern AI agent environments.
- Enhances execution safety and file integrity by eliminating fragile shell piping.
- Directly resolves GitHub Issue #16.
