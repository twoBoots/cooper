# Project Workflow

## Guiding Principles

1. **The Plan is the Source of Truth:** All work must be tracked in `plan.md`
2. **The Tech Stack is Deliberate:** Changes to the tech stack must be documented in `tech-stack.md` *before* implementation
3. **Isolated Workflow Updates:** If changes are required for the `conductor/workflow.md` itself, they MUST be performed in isolation on a separate branch, committed, and submitted via a Pull Request before returning to the active track.
4. **Test-Driven Development:** Write unit tests before implementing functionality
4. **High Code Coverage:** Aim for >80% code coverage for all modules
5. **User Experience First:** Every decision should prioritize user experience
6. **Non-Interactive & CI-Aware:** Prefer non-interactive commands. Use `CI=true` for watch-mode tools (tests, linters) to ensure single execution.

## Track Workflow

### Track Branching Protocol
Before starting any work on a new track:
1. **Identify Track Type and Short Name:** Determine the track type (e.g., `feat`, `fix`, `chore`) and a descriptive short name.
2. **Create & Checkout Branch:** Create a new git branch using the pattern `<type>/<short_name>` (e.g., `feat/extension_popup`).
   ```bash
   git checkout -b <type>/<short_name>
   ```
3. **Commit Exclusively:** All commits related to the track (code, plan updates, checkpoints) MUST be made on this branch.

### Track Finalization (Pull Request)
Once the track is fully completed and has passed the **Code Review Process**:
1. **Push Branch:** Ensure the local branch is pushed to the remote repository.
   ```bash
   git push -u origin <track_id>
   ```
2. **Update Version:** Update the project version using the `npm version` command.
   - **Protocol:** You MUST ask the user to select between `major`, `minor`, `patch`, or `none` using the `ask_user` tool.
   - **Recommendation:** You should provide a recommendation based on your understanding of the changes:
     - **none:** Recommended for tracks that are NOT end-user facing (e.g., internal tests, documentation updates, maintenance chores).
     - **patch:** Recommended for bug fixes or changes that only affect the web dashboard or internal documentation without requiring a Chrome extension update.
     - **minor:** Recommended for new features that are backward compatible.
     - **major:** Recommended for incompatible API changes or significant architectural shifts.
   - **Action (if bump chosen):**
     - **Command:** `npm version [major|minor|patch] --no-git-tag-version`
     - **Synchronization:** Immediately after updating `package.json`, manually update the `version` field in `./src/manifest.json` to match.
     - **Commit:** Stage and commit both `package.json` and `./src/manifest.json` with a message following the format: `chore(release): bump version to v<new_version>`.
   - **Action (if none chosen):** Skip version update and proceed to the next step.
3. **Open Pull Request:** Use the GitHub CLI to open a PR. To avoid shell escaping issues (especially with backticks), follow this protocol:
   - **Step 3.1: Create Body File:** Create a temporary file named `prbody.md` containing the detailed description.
   - **Step 3.2: Generate PR Content:** The PR body MUST include:
     - **Closes:** Reference the Track ID.
     - **Work Summary:** A high-level outline of the work implemented.
     - **Requirements Mapping:** References to specific product requirements (e.g., REQ-123 or section titles from `product.md`) satisfied by this track.
     - **Manual Verification:** Step-by-step instructions for manual verification.
     - **Formatting Protocol:**
       - **File Links:** If referencing a markdown file (e.g., `conductor/product.md`), create a markdown link to that file on the current branch (e.g., `[conductor/product.md](../../blob/<branch_name>/conductor/product.md)`).
       - **Fallback:** If linking is not possible, surround the filename with backticks (e.g., `src/manifest.json`).
       - **URLs:** Format all URLs (including `chrome://extensions`) as markdown links (e.g., `[chrome://extensions](chrome://extensions)`).
   - **Step 3.3: Execute PR Creation:** Use the `--body-file` flag with the GitHub CLI.
     ```bash
     gh pr create --title "feat(conductor): <track_description>" --body-file prbody.md
     ```
   - **Step 3.4: Cleanup:** Delete the temporary `prbody.md` file immediately after the PR is successfully created.

## Task Workflow

All tasks follow a strict lifecycle:

## Track Lifecycle

### Track Initiation

When suggesting new tracks to begin:
1. **Check Existing Branches:** First, check all existing local and remote branches. A track may be being implemented in parallel on a different branch.
2. **Identify In-Progress Work:** If a branch already exists that appears to correspond to a potential track, and it is not the currently checked-out branch, clearly indicate this in the suggestion list.
3. **Avoid Duplication:** Do not suggest creating a new track for work that is already active on another branch unless specifically asked to branch off or restart.

### Track Startup

1. **Create Track Branch:** Create a new branch for the track using the formal pattern `<type>/<short_name>` (e.g., `feat/messaging`).
2. **Initialize Track Files:** Create the track directory and initial files (`index.md`, `metadata.json`, `spec.md`, `plan.md`) following the standard templates. Ensure the `plan.md` contains the initial strategy and tasks.
3. **Push to Remote:** Immediately push the new branch and the initialized track files to the remote repository (`git push -u origin <branch_name>`) so that the plan is visible to other agents and developers.

### Standard Task Workflow

1. **Select Task:** Choose the next available task from `plan.md` in sequential order

2. **Mark In Progress:** Before beginning work, edit `plan.md` and change the task from `[ ]` to `[~]`

3. **Write Failing Tests (Red Phase):**
   - Create a new test file for the feature or bug fix.
   - Write one or more unit tests that clearly define the expected behavior and acceptance criteria for the task.
   - **CRITICAL:** Run the tests and confirm that they fail as expected. This is the "Red" phase of TDD. Do not proceed until you have failing tests.

4. **Firestore Schema & Types (If applicable):**
   - If the task involves Firestore collections (read, write, or update):
     - Create or update the schema markdown in `./docs/schema/firestore/{collection_name}.md`.
     - For subcollections, use the path: `./docs/schema/firestore/{collection_name}/doc/{subcollection_name}.md`.
     - Documentation must include:
       - A brief description at the beginning outlining a general overview of what is being stored and its intended use.
       - Collection Name, Document Identifier, and Schema (Field Name, Type, Example Value, Description).
     - Implement the schema as TypeScript types in `./src/types/`.

5. **Implement to Pass Tests (Green Phase):**
   - Write the minimum amount of application code necessary to make the failing tests pass.
   - Run the test suite again and confirm that all tests now pass. This is the "Green" phase.

5. **Refactor (Optional but Recommended):**
   - With the safety of passing tests, refactor the implementation code and the test code to improve clarity, remove duplication, and enhance performance without changing the external behavior.
   - Rerun tests to ensure they still pass after refactoring.

6. **Verify Coverage:** Run coverage reports using the project's chosen tools. For example, in a Python project, this might look like:
   ```bash
   pytest --cov=app --cov-report=html
   ```
   Target: >80% coverage for new code. The specific tools and commands will vary by language and framework.

7. **Document Deviations:** If implementation differs from tech stack:
   - **STOP** implementation
   - Update `tech-stack.md` with new design
   - Add dated note explaining the change
   - Resume implementation

8. **Commit Code Changes:**
   - Stage all code changes related to the task.
   - Propose a clear, concise commit message e.g, `feat(ui): Create basic HTML structure for calculator`.
   - Perform the commit.

9. **Attach Task Summary with Git Notes:**
   - **Step 9.1: Get Commit Hash:** Obtain the hash of the *just-completed commit* (`git log -1 --format="%H"`).
   - **Step 9.2: Draft Note Content:** Create a detailed summary for the completed task. This should include the task name, a summary of changes, a list of all created/modified files, and the core "why" for the change.
   - **Step 9.3: Attach Note:** Use the `git notes` command to attach the summary to the commit.
     ```bash
     # The note content from the previous step is passed via the -m flag.
     git notes add -m "<note content>" <commit_hash>
     ```

10. **Get and Record Task Commit SHA:**
    - **Step 10.1: Update Plan:** Read `plan.md`, find the line for the completed task, update its status from `[~]` to `[x]`, and append the first 7 characters of the *just-completed commit's* commit hash.
    - **Step 10.2: Write Plan:** Write the updated content back to `plan.md`.

11. **Commit Plan Update:**
    - **Action:** Stage the modified `plan.md` file.
    - **Action:** Commit this change with a descriptive message (e.g., `conductor(plan): Mark task 'Create user model' as complete`).

### Phase Completion Verification and Checkpointing Protocol

**Trigger:** This protocol is executed immediately after a task is completed that also concludes a phase in `plan.md`.

1.  **Check for Workflow Updates:** Execute `git fetch origin main` and check if `conductor/workflow.md` on `origin/main` has changes that are not in the current branch.
    -   **Action:** If changes exist, notify the user and suggest a git strategy (merge or rebase) to bring the current track branch up to date with the latest workflow guidelines before proceeding.

2.  **Announce Protocol Start:** Inform the user that the phase is complete and the verification and checkpointing protocol has begun.

3.  **Ensure Test Coverage for Phase Changes:**
    -   **Step 3.1: Determine Phase Scope:** To identify the files changed in this phase, you must first find the starting point. Read `plan.md` to find the Git commit SHA of the *previous* phase's checkpoint. If no previous checkpoint exists, the scope is all changes since the first commit.
    -   **Step 3.2: List Changed Files:** Execute `git diff --name-only <previous_checkpoint_sha> HEAD` to get a precise list of all files modified during this phase.
    -   **Step 3.3: Verify and Create Tests:** For each file in the list:
        -   **CRITICAL:** First, check its extension. Exclude non-code files (e.g., `.json`, `.md`, `.yaml`).
        -   For each remaining code file, verify a corresponding test file exists.
        -   If a test file is missing, you **must** create one. Before writing the test, **first, analyze other test files in the repository to determine the correct naming convention and testing style.** The new tests **must** validate the functionality described in this phase's tasks (`plan.md`).

4.  **Execute Automated Tests with Proactive Debugging:**
    -   Before execution, you **must** announce the exact shell command you will use to run the tests.
    -   **Example Announcement:** "I will now run the automated test suite to verify the phase. **Command:** `CI=true npm test`"
    -   Execute the announced command.
    -   If tests fail, you **must** inform the user and begin debugging. You may attempt to propose a fix a **maximum of two times**. If the tests still fail after your second proposed fix, you **must stop**, report the persistent failure, and ask the user for guidance.

5.  **Perform Build if Required:**
    -   Determine if the changes in this phase require a build step (e.g., for a Chrome extension, a web app, or a compiled language) before they can be manually verified.
    -   If a build is required, execute the appropriate build command (e.g., `npm run build`).
    -   If the build fails, you **must** inform the user and attempt to debug (maximum of two attempts). If the build still fails after your second proposed fix, you **must stop**, report the persistent failure, and ask the user for guidance.

6.  **Propose a Detailed, Actionable Manual Verification Plan:**
    -   **CRITICAL:** To generate the plan, first analyze `product.md`, `product-guidelines.md`, and `plan.md` to determine the user-facing goals of the completed phase.
    -   You **must** generate a step-by-step plan that walks the user through the verification process, including any necessary commands and specific, expected outcomes.
    -   The plan you present to the user **must** follow this format:

        **For a Frontend Change:**
        ```
        The automated tests have passed. For manual verification, please follow these steps:

        **Manual Verification Steps:**
        1.  **Start the development server with the command:** `npm run dev`
        2.  **Open your browser to:** `http://localhost:3000`
        3.  **Confirm that you see:** The new user profile page, with the user's name and email displayed correctly.
        ```

        **For a Backend Change:**
        ```
        The automated tests have passed. For manual verification, please follow these steps:

        **Manual Verification Steps:**
        1.  **Ensure the server is running.**
        2.  **Execute the following command in your terminal:** `curl -X POST http://localhost:8080/api/v1/users -d '{"name": "test"}'`
        3.  **Confirm that you receive:** A JSON response with a status of `201 Created`.
        ```

7.  **Await Explicit User Feedback:**
    -   After presenting the detailed plan, ask the user for confirmation: "**Does this meet your expectations? Please confirm with yes or provide feedback on what needs to be changed.**"
    -   **PAUSE** and await the user's response. Do not proceed without an explicit yes or confirmation.

8.  **Create Checkpoint Commit:**
    -   Stage all changes. If no changes occurred in this step, proceed with an empty commit.
    -   Perform the commit with a clear and concise message (e.g., `conductor(checkpoint): Checkpoint end of Phase X`).

9.  **Attach Auditable Verification Report using Git Notes:**
    -   **Step 9.1: Draft Note Content:** Create a detailed verification report including the automated test command, the manual verification steps, and the user's confirmation.
    -   **Step 9.2: Attach Note:** Use the `git notes` command and the full commit hash from the previous step to attach the full report to the checkpoint commit.

10. **Get and Record Phase Checkpoint SHA:**
    -   **Step 10.1: Get Commit Hash:** Obtain the hash of the *just-created checkpoint commit* (`git log -1 --format="%H"`).
    -   **Step 10.2: Update Plan:** Read `plan.md`, find the heading for the completed phase, and append the first 7 characters of the commit hash in the format `[checkpoint: <sha>]`.
    -   **Step 10.3: Write Plan:** Write the updated content back to `plan.md`.

11. **Commit Plan Update:**
    - **Action:** Stage the modified `plan.md` file.
    - **Action:** Commit this change with a descriptive message following the format `conductor(plan): Mark phase '<PHASE NAME>' as complete`.

12. **Push to Remote:**
    - **Action:** Push the current branch to the remote repository (`git push origin <branch_name>`) to ensure the checkpoint and plan updates are backed up and visible.

13. **Announce Completion:** Inform the user that the phase is complete and the checkpoint has been created, with the detailed verification report attached as a git note.

### Quality Gates

Before marking any task complete, verify:

- [ ] All tests pass
- [ ] Code coverage meets requirements (>80%)
- [ ] Code follows project's code style guidelines (as defined in `code_styleguides/`)
- [ ] All public functions/methods are documented (e.g., docstrings, JSDoc, GoDoc)
- [ ] Type safety is enforced (e.g., type hints, TypeScript types, Go types)
- [ ] No linting or static analysis errors (using the project's configured tools)
- [ ] Works correctly on mobile (if applicable)
- [ ] Documentation updated if needed
- [ ] No security vulnerabilities introduced

## Development Commands

**AI AGENT INSTRUCTION: This section should be adapted to the project's specific language, framework, and build tools.**

### Setup
```bash
# Example: Commands to set up the development environment (e.g., install dependencies, configure database)
# e.g., for a Node.js project: npm install
# e.g., for a Go project: go mod tidy
```

### Daily Development
```bash
# Example: Commands for common daily tasks (e.g., start dev server, run tests, lint, format)
# e.g., for a Node.js project: npm run dev, npm test, npm run lint
# e.g., for a Go project: go run main.go, go test ./..., go fmt ./...
```

### Before Committing
```bash
# Example: Commands to run all pre-commit checks (e.g., format, lint, type check, run tests)
# e.g., for a Node.js project: npm run check
# e.g., for a Go project: make check (if a Makefile exists)
```

## Testing Requirements

### Unit Testing
- Every module must have corresponding tests.
- Use appropriate test setup/teardown mechanisms (e.g., fixtures, beforeEach/afterEach).
- Mock external dependencies.
- Test both success and failure cases.

### Integration Testing
- Test complete user flows
- Verify database transactions
- Test authentication and authorization
- Check form submissions

### Mobile Testing
- Test on actual iPhone when possible
- Use Safari developer tools
- Test touch interactions
- Verify responsive layouts
- Check performance on 3G/4G

## Code Review Process

### Self-Review Checklist
Before requesting review:

1. **Functionality**
   - Feature works as specified
   - Edge cases handled
   - Error messages are user-friendly

2. **Code Quality**
   - Follows style guide
   - DRY principle applied
   - Clear variable/function names
   - Appropriate comments

3. **Testing**
   - Unit tests comprehensive
   - Integration tests pass
   - Coverage adequate (>80%)

4. **Security**
   - No hardcoded secrets
   - Input validation present
   - SQL injection prevented
   - XSS protection in place

5. **Performance**
   - Database queries optimized
   - Images optimized
   - Caching implemented where needed

6. **Mobile Experience**
   - Touch targets adequate (44x44px)
   - Text readable without zooming
   - Performance acceptable on mobile
   - Interactions feel native

## Commit Guidelines

### Message Format
```
<type>(<scope>): <description>

[optional body]

[optional footer]
```

### Types
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation only
- `style`: Formatting, missing semicolons, etc.
- `refactor`: Code change that neither fixes a bug nor adds a feature
- `test`: Adding missing tests
- `chore`: Maintenance tasks

### Examples
```bash
git commit -m "feat(auth): Add remember me functionality"
git commit -m "fix(posts): Correct excerpt generation for short posts"
git commit -m "test(comments): Add tests for emoji reaction limits"
git commit -m "style(mobile): Improve button touch targets"
```

## Definition of Done

A task is complete when:

1. All code implemented to specification
2. Unit tests written and passing
3. Code coverage meets project requirements
4. Documentation complete (if applicable)
5. Code passes all configured linting and static analysis checks
6. Works beautifully on mobile (if applicable)
7. Implementation notes added to `plan.md`
8. Changes committed with proper message
9. Git note with task summary attached to the commit

## Emergency Procedures

### Critical Bug in Production
1. Create hotfix branch from main
2. Write failing test for bug
3. Implement minimal fix
4. Test thoroughly including mobile
5. Deploy immediately
6. Document in plan.md

### Data Loss
1. Stop all write operations
2. Restore from latest backup
3. Verify data integrity
4. Document incident
5. Update backup procedures

### Security Breach
1. Rotate all secrets immediately
2. Review access logs
3. Patch vulnerability
4. Notify affected users (if any)
5. Document and update security procedures

## Deployment Workflow

### Pre-Deployment Checklist
- [ ] All tests passing
- [ ] Coverage >80%
- [ ] No linting errors
- [ ] Mobile testing complete
- [ ] Environment variables configured
- [ ] Database migrations ready
- [ ] Backup created

### Deployment Steps
1. Merge feature branch to main
2. Tag release with version
3. Push to deployment service
4. Run database migrations
5. Verify deployment
6. Test critical paths
7. Monitor for errors

### Post-Deployment
1. Monitor analytics
2. Check error logs
3. Gather user feedback
4. Plan next iteration

## Continuous Improvement

- Review workflow weekly
- Update based on pain points
- Document lessons learned
- Optimize for user happiness
- Keep things simple and maintainable
