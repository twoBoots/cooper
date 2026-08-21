# Spec Delta: Framework Documentation & Guidelines (Suggestive Ethos)

## Capability: documentation

### ADD: Requirement: Suggestive Framework Framing & Project Adaptation
The framework's primary documentation (`README.md`, `COOPER.md`, `AGENTS.md`) SHALL explicitly communicate that Cooper is suggestive rather than prescriptive, and highlight the intent for projects to install and adapt the framework to their specific needs.

#### ADD: Scenario: Core Philosophy Presentation in README
- GIVEN the primary `README.md` file
- WHEN a reader visits or views the documentation
- THEN it MUST include a dedicated Philosophy section defining Cooper as suggestive not prescriptive
- AND it MUST explain that repository owners own all rules, definitions, styleguides, and agent skills in-repo.

#### ADD: Scenario: Three-Step Lifecycle Presentation
- GIVEN the top-level framework documentation
- WHEN presenting the developer workflow or onboarding guide
- THEN it MUST present the 3-step lifecycle: **Install → Adapt → Build**
- AND it MUST specify how users can customize `.cooper/definition/`, `.cooper/code_styleguides/`, and `.agents/skills/`.
