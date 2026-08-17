# Technical Design: Relocate TROOP.md into .cooper/ and align repository structure

## Architecture & File Layout
The `.cooper/` layout in target projects and this repository will be:
```
.cooper/
├── TROOP.md                       # Relocated Troop reference manual
├── definition/                    # Core definitions & workflow rules
│   ├── product.md
│   ├── tech-stack.md
│   ├── product-guidelines.md
│   └── workflow.md
├── code_styleguides/
├── specs/
├── active/
└── archive/
```

## `install.sh` Adjustments
1. **Pre-creation**: Create `.cooper/` structure prior to running Troop installer or immediately move `TROOP.md` after Troop setup finishes:
   ```bash
   if [ -f "TROOP.md" ]; then
       mv "TROOP.md" ".cooper/TROOP.md"
       echo "  [✓] Moved TROOP.md to .cooper/TROOP.md"
   fi
   ```
2. **File Fetching**:
   Update `get_cooper_file "cooper/workflow.md" ".cooper/definition/workflow.md"` to:
   `get_cooper_file ".cooper/definition/workflow.md" ".cooper/definition/workflow.md"` (with fallback handling if needed).
3. **AGENTS.template.md**:
   Add reference to `.cooper/TROOP.md` for worktree isolation commands.
