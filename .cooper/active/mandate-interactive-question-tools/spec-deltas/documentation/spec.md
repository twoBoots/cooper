# Spec Delta: Framework Documentation & Guidelines

## Capability: documentation

## Requirements

### Requirement: Interactive Question & File Tool Protocols
+ When interacting with users or performing file operations during skill execution, agents MUST prioritize runtime interactive tools and native file tools over plain text or shell stream redirections.

#### Scenario: Interactive Question Tool Invocations
+ - GIVEN an agent executing a Cooper skill or workflow step requiring user input, selection, or confirmation
+ - AND the agent runtime provides an interactive question tool (such as `ask_question`)
+ - WHEN the agent presents single-choice or multiple-choice options
+ - THEN the agent MUST invoke the interactive question tool
+ - AND the agent MUST NOT output plain text numbered/bulleted option lists in chat.

#### Scenario: Fallback to Text Chat for Questions
+ - GIVEN an agent executing a Cooper skill or workflow step
+ - AND no interactive question tool is provided in the agent runtime environment
+ - WHEN the agent needs to ask questions or present choices
+ - THEN the agent SHALL format the question clearly in text chat, asking questions strictly one at a time.

#### Scenario: Native File Tools Enforcement
+ - GIVEN an agent runtime providing native file tools (`view_file`, `write_to_file`, `replace_file_content`)
+ - WHEN the agent needs to view, create, or edit files in the workspace
+ - THEN the agent MUST use native file tools
+ - AND the agent MUST NOT use shell commands with heredocs, pipes, or stream redirections (`cat << 'EOF'`, `echo >`) to create or edit files.
