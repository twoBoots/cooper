# Konduktor

This repository documents and persists the customized **Conductor Workflow**. 

## Overview

The Conductor workflow is a structured approach to project management and development for AI agents and humans alike. See [gemini-cli-extensions/conductor](https://github.com/gemini-cli-extensions/conductor).

## Structure

- `conductor/workflow.md`: The customised workflow specification.
- `AGENT.md`: Agnostic instructions for AI agents to ensure compliance with this workflow. These can be copied into the relevant root agent file e.g. GEMINI.md or CLAUDE.md

## Usage

When setting up a new project with this workflow:
1. Copy the `conductor/workflow.md` file, overwriting the default workflow. If the current workflow file is customized, stop immediately and ask for clarification of next steps.
2. Copy the instructions from `AGENT.md` to your agent's primary instruction file (e.g., `GEMINI.md` or `CLAUDE.md`).
