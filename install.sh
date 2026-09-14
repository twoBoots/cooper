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
COOPER_RELEASE_BASE_URL="https://github.com/twoBoots/cooper/releases/latest/download"

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

# Resolve the first writable directory for the cooper binary.
# Echoes the chosen path, or nothing when no candidate is writable.
resolve_bin_dir() {
    local home_bin="${HOME:-}/.local/bin"

    if [ -d "/usr/local/bin" ] && [ -w "/usr/local/bin" ]; then
        echo "/usr/local/bin"
        return 0
    fi

    if [ -n "${HOME:-}" ]; then
        if mkdir -p "$home_bin" 2>/dev/null && [ -w "$home_bin" ]; then
            echo "$home_bin"
            return 0
        fi
    fi

    return 0
}

# Install the optional cooper binary.
#
# The binary provides 'cooper validate' for projects that want deterministic
# spec linting in CI. It is strictly optional: the .cooper/ and .agents/skills/
# workspace is fully usable without it, so every failure here is informational
# and this function always returns 0. install.sh runs under 'set -e', so
# scaffolding would otherwise abort on an air-gapped host, a rate-limited
# GitHub API, or an unwritable target directory.
install_cooper_binary() {
    local bin_dir="${1:-}"
    local os arch asset target

    if [ -z "$bin_dir" ]; then
        bin_dir="$(resolve_bin_dir)"
    fi

    if [ -z "$bin_dir" ]; then
        echo "  [i] No writable bin directory found. Continuing in zero-binary mode."
        return 0
    fi

    target="${bin_dir}/cooper"

    # Tier 1: compile from a local clone when a Go toolchain is present.
    if [ -n "$SCRIPT_DIR" ] && [ -f "$SCRIPT_DIR/main.go" ] && command -v go >/dev/null 2>&1; then
        if (cd "$SCRIPT_DIR" && go build -ldflags="-s -w" -o "$target" . 2>/dev/null); then
            _cooper_post_process_binary "$target"
            echo "  [✓] Compiled cooper binary -> $target"
            return 0
        fi
        echo "  [i] Local compilation failed. Trying pre-built release asset..."
    fi

    # Tier 2: download the pre-built release asset for this platform.
    os="$(uname -s 2>/dev/null | tr '[:upper:]' '[:lower:]' || true)"
    arch="$(uname -m 2>/dev/null || true)"

    case "$arch" in
        x86_64 | amd64) arch="x86_64" ;;
        arm64 | aarch64) arch="aarch64" ;;
        *) arch="" ;;
    esac

    if [ -n "$os" ] && [ -n "$arch" ]; then
        asset="cooper-${os}-${arch}"
        if command -v curl >/dev/null 2>&1; then
            if curl -fsSL "${COOPER_RELEASE_BASE_URL}/${asset}" -o "$target" 2>/dev/null; then
                chmod +x "$target" 2>/dev/null || true
                _cooper_post_process_binary "$target"
                echo "  [✓] Installed cooper binary -> $target"
                return 0
            fi
        elif command -v wget >/dev/null 2>&1; then
            if wget -qO "$target" "${COOPER_RELEASE_BASE_URL}/${asset}" 2>/dev/null; then
                chmod +x "$target" 2>/dev/null || true
                _cooper_post_process_binary "$target"
                echo "  [✓] Installed cooper binary -> $target"
                return 0
            fi
        fi
        rm -f "$target" 2>/dev/null || true
    fi

    # Tier 3: graceful zero-binary fallback.
    echo "  [i] Could not obtain the cooper binary (offline, unsupported platform, or no download tool)."
    echo "  [i] Continuing in zero-binary mode. The .cooper/ workspace and agent skills work without it."
    return 0
}

# Strip macOS quarantine and apply an ad-hoc signature so Gatekeeper does not
# block a freshly downloaded or compiled binary. Never fatal.
_cooper_post_process_binary() {
    local target="$1"
    if [ "$(uname -s 2>/dev/null || true)" = "Darwin" ]; then
        xattr -d com.apple.quarantine "$target" >/dev/null 2>&1 || true
        codesign -s - --force "$target" >/dev/null 2>&1 || true
    fi
    return 0
}

# Move TROOP.md into .cooper/ and repoint any references to it.
#
# The Troop installer writes an AGENTS.md that links to TROOP.md at the
# repository root. Relocating the file without rewriting those references
# leaves a dangling link that 'cooper validate' reports in every freshly
# scaffolded project, so both forms -- the markdown link and the inline-code
# mention -- are updated here.
relocate_troop_reference() {
    if [ -f "TROOP.md" ]; then
        mkdir -p .cooper
        mv "TROOP.md" ".cooper/TROOP.md"
        echo "  [✓] Relocated TROOP.md to .cooper/TROOP.md"
    fi

    [ -f "AGENTS.md" ] || return 0

    local tmp
    tmp="$(mktemp)" || return 0

    # Rewrite only root-relative references, leaving any already pointing at
    # .cooper/ untouched.
    sed -e 's|](\./TROOP\.md)|](./.cooper/TROOP.md)|g' \
        -e 's|](TROOP\.md)|](.cooper/TROOP.md)|g' \
        -e 's|`TROOP\.md`|`.cooper/TROOP.md`|g' \
        "AGENTS.md" > "$tmp" 2>/dev/null || { rm -f "$tmp"; return 0; }

    if ! cmp -s "$tmp" "AGENTS.md"; then
        mv "$tmp" "AGENTS.md"
        echo "  [✓] Repointed AGENTS.md references to .cooper/TROOP.md"
    else
        rm -f "$tmp"
    fi

    return 0
}

# Library mode: when sourced with COOPER_INSTALL_LIB_ONLY set, define the
# functions above and stop, performing no installation. Used by the test suite.
if [ -n "${COOPER_INSTALL_LIB_ONLY:-}" ]; then
    return 0 2>/dev/null || exit 0
fi

cd "$TARGET_DIR"

if [ ! -d ".git" ] && ! git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
    echo "Error: Target directory '$TARGET_DIR' is not a Git repository."
    echo "Please initialize git first using 'git init'."
    exit 1
fi

echo "🛢️ Installing Cooper (Cooper SDD + Living Specs + Troop Worktrees) into $(pwd)..."

# 1. Run Troop installer first (worktree setup, .gitaliases, .gitignore, TROOP.md)
echo "  [1/6] Setting up Troop worktree foundation..."
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

# Relocate TROOP.md into .cooper/ to keep project root clean, repointing
# any references the Troop installer left behind.
relocate_troop_reference

CONDUCTOR_EXISTS=false
OPENSPEC_EXISTS=false

if [ -d "conductor" ]; then
    CONDUCTOR_EXISTS=true
fi
if [ -d "openspec" ]; then
    OPENSPEC_EXISTS=true
fi

# 2. Check and Migrate Existing Setup or Fetch Baseline Scaffolding
echo "  [2/6] Scaffolding & Migration Analysis..."

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
echo "  [3/6] Installing Cooper workflow specification & COOPER.md..."
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
echo "  [4/6] Installing project-local Cooper skills into .agents/skills/..."
get_cooper_file "skills/cooper-setup/SKILL.md" ".agents/skills/cooper-setup/SKILL.md"
get_cooper_file "skills/cooper-rfc/SKILL.md" ".agents/skills/cooper-rfc/SKILL.md"
get_cooper_file "skills/cooper-new-track/SKILL.md" ".agents/skills/cooper-new-track/SKILL.md"
get_cooper_file "skills/cooper-implement/SKILL.md" ".agents/skills/cooper-implement/SKILL.md"
get_cooper_file "skills/cooper-review/SKILL.md" ".agents/skills/cooper-review/SKILL.md"
get_cooper_file "skills/cooper-status/SKILL.md" ".agents/skills/cooper-status/SKILL.md"
echo "  [✓] Installed Cooper skills (.agents/skills/cooper-{setup,rfc,new-track,implement,review,status})"

# 5. Setup AGENTS.md
echo "  [5/6] Setting up AGENTS.md rules..."
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

# 6. Install the optional cooper binary (never fatal)
echo "  [6/6] Installing optional cooper binary..."
install_cooper_binary

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
