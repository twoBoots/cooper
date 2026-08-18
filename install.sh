#!/usr/bin/env bash
set -e

# Cooper Remote/Local Installer Script
# Scaffolds Cooper (Spec-Driven Development with Living Spec Deltas + Troop Worktrees)
# into any Git repository, including native project-local agent skills under .agents/skills/.
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/twoBoots/cooper/main/install.sh | bash
#   or: ./install.sh [target_directory]

RAW_BASE_URL="https://raw.githubusercontent.com/twoBoots/cooper/main"
TROOP_RAW_BASE_URL="https://raw.githubusercontent.com/twoBoots/troop/main"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}" 2>/dev/null)" && pwd || true)"
TARGET_DIR="${1:-$(pwd)}"

cd "$TARGET_DIR"

if [ ! -d ".git" ] && ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    echo "Error: Target directory '$TARGET_DIR' is not a Git repository."
    echo "Please initialize git first using 'git init'."
    exit 1
fi

echo "🛢️ Installing Cooper (Cooper SDD + Living Specs + Troop Worktrees) into $(pwd)..."

# Helper function to fetch or copy a file from Cooper repo
get_cooper_file() {
    local filename="$1"
    local dest="${2:-$filename}"
    local dest_dir="$(dirname "$dest")"
    if [ "$dest_dir" != "." ] && [ ! -d "$dest_dir" ]; then
        mkdir -p "$dest_dir"
    fi

    if [ -n "$SCRIPT_DIR" ] && [ -f "$SCRIPT_DIR/$filename" ]; then
        cp "$SCRIPT_DIR/$filename" "$dest"
    elif command -v curl >/dev/null 2>&1; then
        curl -fsSL "$RAW_BASE_URL/$filename" -o "$dest"
    elif command -v wget >/dev/null 2>&1; then
        wget -qO "$dest" "$RAW_BASE_URL/$filename"
    else
        echo "Error: Neither curl nor wget found, and local $filename is missing."
        exit 1
    fi
}

# 1. Run Troop installer first (worktree setup, .gitaliases, .gitignore, TROOP.md)
echo "  [1/5] Setting up Troop worktree foundation..."
if [ -n "$SCRIPT_DIR" ] && [ -f "$SCRIPT_DIR/../troop/install.sh" ]; then
    bash "$SCRIPT_DIR/../troop/install.sh" "$TARGET_DIR"
elif command -v curl >/dev/null 2>&1; then
    curl -fsSL "$TROOP_RAW_BASE_URL/install.sh" | bash
elif command -v wget >/dev/null 2>&1; then
    wget -qO- "$TROOP_RAW_BASE_URL/install.sh" | bash
else
    echo "Error: Unable to fetch Troop installer."
    exit 1
fi

# Ensure base .cooper directory tree exists
mkdir -p .cooper/definition .cooper/code_styleguides .cooper/specs .cooper/active .cooper/archive .agents/skills

# Relocate TROOP.md into .cooper/ to keep project root clean
if [ -f "TROOP.md" ]; then
    mv "TROOP.md" ".cooper/TROOP.md"
    echo "  [✓] Relocated TROOP.md to .cooper/TROOP.md"
fi

CONDUCTOR_EXISTS=false
OPENSPEC_EXISTS=false

if [ -d "conductor" ]; then
    CONDUCTOR_EXISTS=true
fi
if [ -d "openspec" ]; then
    OPENSPEC_EXISTS=true
fi

# 2. Check and Migrate Existing Setup or Fetch Baseline Scaffolding
echo "  [2/5] Scaffolding & Migration Analysis..."

if [ "$CONDUCTOR_EXISTS" = true ]; then
    echo "  [→] Existing Conductor setup detected. Migrating to .cooper/ structure..."
    
    # Migrate global definitions
    [ -f "conductor/product.md" ] && cp "conductor/product.md" .cooper/definition/product.md
    [ -f "conductor/product-guidelines.md" ] && cp "conductor/product-guidelines.md" .cooper/definition/product-guidelines.md
    [ -f "conductor/tech-stack.md" ] && cp "conductor/tech-stack.md" .cooper/definition/tech-stack.md
    
    # Migrate styleguides
    if [ -d "conductor/code_styleguides" ]; then
        cp -r conductor/code_styleguides/* .cooper/code_styleguides/ 2>/dev/null || true
    fi
    
    # Migrate tracks into archive
    if [ -d "conductor/tracks" ]; then
        cp -r conductor/tracks/* .cooper/archive/ 2>/dev/null || true
    fi
    echo "  [✓] Conductor configuration migrated to .cooper/"
fi

if [ "$OPENSPEC_EXISTS" = true ]; then
    echo "  [→] Existing OpenSpec setup detected. Migrating living specs to .cooper/specs/..."
    
    if [ -d "openspec/specs" ]; then
        cp -r openspec/specs/* .cooper/specs/ 2>/dev/null || true
    fi
    if [ -d "openspec/changes" ]; then
        cp -r openspec/changes/* .cooper/active/ 2>/dev/null || true
    fi
    echo "  [✓] OpenSpec living specs and changes migrated to .cooper/"
fi

if [ "$CONDUCTOR_EXISTS" = false ] && [ "$OPENSPEC_EXISTS" = false ]; then
    echo "  [→] Greenfield project detected (neither Conductor nor OpenSpec found)."
    echo "  [→] Installing baseline Cooper templates..."
    
    get_cooper_file "templates/product.md" ".cooper/definition/product.md"
    get_cooper_file "templates/tech-stack.md" ".cooper/definition/tech-stack.md"
    get_cooper_file "templates/product-guidelines.md" ".cooper/definition/product-guidelines.md"
    get_cooper_file "templates/code_styleguides/typescript.md" ".cooper/code_styleguides/typescript.md"
    get_cooper_file "templates/code_styleguides/python.md" ".cooper/code_styleguides/python.md"
    get_cooper_file "templates/code_styleguides/go.md" ".cooper/code_styleguides/go.md"
    get_cooper_file "templates/code_styleguides/rust.md" ".cooper/code_styleguides/rust.md"
    
    echo "  [✓] Baseline scaffolding created under .cooper/"
fi

# 3. Install Cooper Hybrid workflow specification & COOPER.md reference
echo "  [3/5] Installing Cooper workflow specification & COOPER.md..."
get_cooper_file ".cooper/definition/workflow.md" ".cooper/definition/workflow.md"
get_cooper_file ".cooper/COOPER.md" ".cooper/COOPER.md"

# Scaffold Handshake Index (.cooper/index.md)
cat << 'EOF' > .cooper/index.md
# Project Context (.cooper)

## Definition
- [Product Definition](./definition/product.md)
- [Product Guidelines](./definition/product-guidelines.md)
- [Tech Stack](./definition/tech-stack.md)
- [Workflow](./definition/workflow.md)
- [Code Style Guides](./code_styleguides/)

## Living Specifications
- [Capability Specs](./specs/)

## Tracks
- [Tracks Registry](./tracks.md)
- [Active Tracks](./active/)
- [Archive](./archive/)

## Capabilities
- [Agent Skills](../.agents/skills/)
EOF

# Initialize tracks registry if missing
if [ ! -f ".cooper/tracks.md" ]; then
    cat << 'EOF' > .cooper/tracks.md
# Tracks Registry

All active and completed Cooper tracks are registered below.

---
EOF
fi
echo "  [✓] Installed .cooper/definition/workflow.md, COOPER.md & index.md"

# 4. Install Project-Local Agent Skills (.agents/skills/cooper-*)
echo "  [4/5] Installing project-local Cooper skills into .agents/skills/..."
get_cooper_file "skills/cooper-setup/SKILL.md" ".agents/skills/cooper-setup/SKILL.md"
get_cooper_file "skills/cooper-rfc/SKILL.md" ".agents/skills/cooper-rfc/SKILL.md"
get_cooper_file "skills/cooper-new-track/SKILL.md" ".agents/skills/cooper-new-track/SKILL.md"
get_cooper_file "skills/cooper-implement/SKILL.md" ".agents/skills/cooper-implement/SKILL.md"
get_cooper_file "skills/cooper-review/SKILL.md" ".agents/skills/cooper-review/SKILL.md"
get_cooper_file "skills/cooper-status/SKILL.md" ".agents/skills/cooper-status/SKILL.md"
echo "  [✓] Installed Cooper skills (.agents/skills/cooper-{setup,rfc,new-track,implement,review,status})"

# 5. Setup AGENTS.md
echo "  [5/5] Setting up AGENTS.md rules..."
TMP_TEMPLATE="$(mktemp)"
if [ -n "$SCRIPT_DIR" ] && [ -f "$SCRIPT_DIR/AGENTS.template.md" ]; then
    cp "$SCRIPT_DIR/AGENTS.template.md" "$TMP_TEMPLATE"
elif command -v curl >/dev/null 2>&1; then
    curl -fsSL "$RAW_BASE_URL/AGENTS.template.md" -o "$TMP_TEMPLATE"
elif command -v wget >/dev/null 2>&1; then
    wget -qO "$TMP_TEMPLATE" "$RAW_BASE_URL/AGENTS.template.md"
fi

if [ -f "AGENTS.md" ]; then
    if ! grep -qs "Cooper" AGENTS.md; then
        echo "" >> AGENTS.md
        cat "$TMP_TEMPLATE" >> AGENTS.md
        echo "  [✓] Appended Cooper rules to existing AGENTS.md"
    else
        echo "  [✓] Cooper rules already present in AGENTS.md"
    fi
else
    cp "$TMP_TEMPLATE" AGENTS.md
    echo "  [✓] Created AGENTS.md from Cooper template"
fi
rm -f "$TMP_TEMPLATE"

echo ""
echo "🛢️ Cooper SDD Framework successfully installed!"
echo ""
echo "Available Project Skills (.agents/skills/):"
echo "  - cooper-setup       : Re-audit & configure definitions/specs"
echo "  - cooper-rfc         : Collaborative RFC planning, Draft PR & spec deltas"
echo "  - cooper-new-track   : Spawn worktree & draft spec deltas/plan"
echo "  - cooper-implement   : Execute TDD cycle, Git Notes & phase sync"
echo "  - cooper-review      : Audit changes against spec deltas & style"
echo "  - cooper-status      : Overview of worktrees, tracks & checkpoints"
echo ""
echo "Workflow summary:"
echo "  1. Start track in isolated worktree : git agent-start <track_id>"
echo "  2. List active tracks               : git troop"
echo "  3. Develop & checkpoint             : Follow .cooper/definition/workflow.md"
echo "  4. Teardown track after PR merge    : git agent-stop <track_id>"
