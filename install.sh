#!/usr/bin/env bash
set -e

# Cooper Remote/Local Installer Script
# Scaffolds Cooper (Conductor Workflow + JungleJim Worktrees) into any Git repository.
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/twoBoots/cooper/main/install.sh | bash
#   or: ./install.sh [target_directory]

RAW_BASE_URL="https://raw.githubusercontent.com/twoBoots/cooper/main"
JJ_RAW_BASE_URL="https://raw.githubusercontent.com/twoBoots/junglejim/main"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}" 2>/dev/null)" && pwd || true)"
TARGET_DIR="${1:-$(pwd)}"

cd "$TARGET_DIR"

if [ ! -d ".git" ] && ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    echo "Error: Target directory '$TARGET_DIR' is not a Git repository."
    echo "Please initialize git first using 'git init'."
    exit 1
fi

echo "🛢️ Installing Cooper (Conductor + JungleJim) into $(pwd)..."

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

# 1. Run JungleJim installer first (worktree setup, .gitaliases, .gitignore, JUNGLEJIM.md)
echo "  [1/3] Setting up JungleJim worktree foundation..."
if [ -n "$SCRIPT_DIR" ] && [ -f "$SCRIPT_DIR/../junglejim/install.sh" ]; then
    "$SCRIPT_DIR/../junglejim/install.sh" "$TARGET_DIR"
elif command -v curl >/dev/null 2>&1; then
    curl -fsSL "$JJ_RAW_BASE_URL/install.sh" | bash
elif command -v wget >/dev/null 2>&1; then
    wget -qO- "$JJ_RAW_BASE_URL/install.sh" | bash
else
    echo "Error: Unable to fetch JungleJim installer."
    exit 1
fi

# 2. Install Conductor workflow specification
echo "  [2/3] Installing Conductor workflow specification..."
get_cooper_file "conductor/workflow.md" "conductor/workflow.md"
echo "  [✓] Installed conductor/workflow.md"

# 3. Setup AGENTS.md
echo "  [3/3] Setting up AGENTS.md rules..."
TMP_TEMPLATE="$(mktemp)"
if [ -n "$SCRIPT_DIR" ] && [ -f "$SCRIPT_DIR/AGENTS.template.md" ]; then
    cp "$SCRIPT_DIR/AGENTS.template.md" "$TMP_TEMPLATE"
elif command -v curl >/dev/null 2>&1; then
    curl -fsSL "$RAW_BASE_URL/AGENTS.template.md" -o "$TMP_TEMPLATE"
elif command -v wget >/dev/null 2>&1; then
    wget -qO "$TMP_TEMPLATE" "$RAW_BASE_URL/AGENTS.template.md"
fi

if [ -f "AGENTS.md" ]; then
    if ! grep -qs "Conductor" AGENTS.md; then
        echo "" >> AGENTS.md
        cat "$TMP_TEMPLATE" >> AGENTS.md
        echo "  [✓] Appended Cooper Conductor rules to existing AGENTS.md"
    else
        echo "  [✓] Cooper Conductor rules already present in AGENTS.md"
    fi
else
    cp "$TMP_TEMPLATE" AGENTS.md
    echo "  [✓] Created AGENTS.md from Cooper template"
fi
rm -f "$TMP_TEMPLATE"

echo ""
echo "🛢️ Cooper successfully installed!"
echo "Workflow summary:"
echo "  1. Start track in isolated worktree : git agent-start <track_id>"
echo "  2. List active tracks               : git jims"
echo "  3. Develop & checkpoint             : Follow conductor/workflow.md"
echo "  4. Teardown track after PR merge    : git agent-stop <track_id>"
