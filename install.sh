#!/usr/bin/env bash
set -e

# Cooper Remote/Local Installer Script
# Scaffolds Cooper (Cooper Hybrid SDD Workflow + Troop Worktrees) into any Git repository.
# Handles auto-migration from existing Conductor or OpenSpec setups, or fetches baseline
# scaffolding from twoBoots/conductor if neither exists.
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/twoBoots/cooper/main/install.sh | bash
#   or: ./install.sh [target_directory]

RAW_BASE_URL="https://raw.githubusercontent.com/twoBoots/cooper/main"
CONDUCTOR_RAW_BASE_URL="https://raw.githubusercontent.com/twoBoots/conductor/main"
TROOP_RAW_BASE_URL="https://raw.githubusercontent.com/twoBoots/troop/main"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}" 2>/dev/null)" && pwd || true)"
TARGET_DIR="${1:-$(pwd)}"

cd "$TARGET_DIR"

if [ ! -d ".git" ] && ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    echo "Error: Target directory '$TARGET_DIR' is not a Git repository."
    echo "Please initialize git first using 'git init'."
    exit 1
fi

echo "🛢️ Installing Cooper (Cooper Hybrid SDD + Troop) into $(pwd)..."

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

# Helper function to fetch template from twoBoots/conductor repo
get_conductor_template() {
    local filename="$1"
    local dest="$2"
    local dest_dir="$(dirname "$dest")"
    if [ "$dest_dir" != "." ] && [ ! -d "$dest_dir" ]; then
        mkdir -p "$dest_dir"
    fi

    if command -v curl >/dev/null 2>&1; then
        curl -fsSL "$CONDUCTOR_RAW_BASE_URL/$filename" -o "$dest" 2>/dev/null || true
    elif command -v wget >/dev/null 2>&1; then
        wget -qO "$dest" "$CONDUCTOR_RAW_BASE_URL/$filename" 2>/dev/null || true
    fi
}

# 1. Run Troop installer first (worktree setup, .gitaliases, .gitignore, TROOP.md)
echo "  [1/4] Setting up Troop worktree foundation..."
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
mkdir -p .cooper/definition .cooper/code_styleguides .cooper/specs .cooper/active .cooper/archive

CONDUCTOR_EXISTS=false
OPENSPEC_EXISTS=false

if [ -d "conductor" ]; then
    CONDUCTOR_EXISTS=true
fi
if [ -d "openspec" ]; then
    OPENSPEC_EXISTS=true
fi

# 2. Check and Migrate Existing Setup or Fetch Baseline Scaffolding
echo "  [2/4] Scaffolding & Migration Analysis..."

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
    
    # Migrate tracks into archive / active
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
    echo "  [→] Fetching baseline scaffolding from twoBoots/conductor..."
    
    get_conductor_template "templates/product.md" ".cooper/definition/product.md"
    get_conductor_template "templates/tech-stack.md" ".cooper/definition/tech-stack.md"
    get_conductor_template "templates/product-guidelines.md" ".cooper/definition/product-guidelines.md"
    get_conductor_template "templates/code_styleguides/typescript.md" ".cooper/code_styleguides/typescript.md"
    get_conductor_template "templates/code_styleguides/python.md" ".cooper/code_styleguides/python.md"
    
    # Create initial product.md placeholder if template download failed
    if [ ! -s ".cooper/definition/product.md" ]; then
        cat << 'EOF' > .cooper/definition/product.md
# Product Definition

## Vision
Define the core product goals and target audience.

## Requirements
- Core functional goals
EOF
    fi

    # Create initial tech-stack.md placeholder if template download failed
    if [ ! -s ".cooper/definition/tech-stack.md" ]; then
        cat << 'EOF' > .cooper/definition/tech-stack.md
# Tech Stack Definition

- **Languages:** TypeScript / Python
- **Testing:** Jest / PyTest
- **CI/CD:** GitHub Actions
EOF
    fi

    echo "  [✓] Initial baseline scaffolding created under .cooper/"
fi

# 3. Install Cooper Hybrid workflow specification (.cooper/ & conductor/ compatibility link)
echo "  [3/4] Installing Cooper Hybrid workflow specification..."
get_cooper_file "cooper/workflow.md" ".cooper/definition/workflow.md"

# Backwards compatibility: maintain conductor/workflow.md
mkdir -p conductor
cp .cooper/definition/workflow.md conductor/workflow.md
echo "  [✓] Installed .cooper/definition/workflow.md (and backward-compatible conductor/workflow.md)"

# 4. Setup AGENTS.md
echo "  [4/4] Setting up AGENTS.md rules..."
TMP_TEMPLATE="$(mktemp)"
if [ -n "$SCRIPT_DIR" ] && [ -f "$SCRIPT_DIR/AGENTS.template.md" ]; then
    cp "$SCRIPT_DIR/AGENTS.template.md" "$TMP_TEMPLATE"
elif command -v curl >/dev/null 2>&1; then
    curl -fsSL "$RAW_BASE_URL/AGENTS.template.md" -o "$TMP_TEMPLATE"
elif command -v wget >/dev/null 2>&1; then
    wget -qO "$TMP_TEMPLATE" "$RAW_BASE_URL/AGENTS.template.md"
fi

if [ -f "AGENTS.md" ]; then
    if ! grep -qs "Cooper Hybrid" AGENTS.md; then
        echo "" >> AGENTS.md
        cat "$TMP_TEMPLATE" >> AGENTS.md
        echo "  [✓] Appended Cooper Hybrid rules to existing AGENTS.md"
    else
        echo "  [✓] Cooper Hybrid rules already present in AGENTS.md"
    fi
else
    cp "$TMP_TEMPLATE" AGENTS.md
    echo "  [✓] Created AGENTS.md from Cooper template"
fi
rm -f "$TMP_TEMPLATE"

echo ""
echo "🛢️ Cooper Hybrid SDD successfully installed!"
echo "Workflow summary:"
echo "  1. Start track in isolated worktree : git agent-start <track_id>"
echo "  2. List active tracks               : git troop"
echo "  3. Develop & checkpoint             : Follow .cooper/definition/workflow.md"
echo "  4. Teardown track after PR merge    : git agent-stop <track_id>"
