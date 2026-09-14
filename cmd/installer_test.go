package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const installScript = "../install.sh"

func requireBash(t *testing.T) string {
	t.Helper()
	bash, err := exec.LookPath("bash")
	if err != nil {
		t.Skip("bash not available on this platform")
	}
	return bash
}

func TestInstallScript_SyntaxIsValid(t *testing.T) {
	bash := requireBash(t)

	out, err := exec.Command(bash, "-n", installScript).CombinedOutput()
	if err != nil {
		t.Fatalf("install.sh failed syntax check: %v\n%s", err, out)
	}
}

// TestInstallScript_BinaryInstallIsNonFatal enforces the installer spec
// scenario "Graceful Zero-Binary Fallback". install.sh runs under `set -e`, so
// any unguarded failure in the binary step would abort scaffolding. The step is
// therefore a function that must return 0 even when no toolchain, no download
// tool, and no writable target directory are available.
func TestInstallScript_BinaryInstallIsNonFatal(t *testing.T) {
	bash := requireBash(t)

	scriptPath, err := filepath.Abs(installScript)
	if err != nil {
		t.Fatalf("failed resolving install.sh: %v", err)
	}

	// An empty PATH removes go, curl, wget, git and every other helper, which
	// is the worst case the fallback must survive. Note this exits early at
	// bin-dir resolution; the download-failure path is covered separately by
	// TestInstallScript_FailedDownloadPreservesExistingBinary.
	emptyBin := t.TempDir()
	target := t.TempDir()

	// Source the script in library mode, then invoke only the binary step.
	program := `
set -e
export COOPER_INSTALL_LIB_ONLY=1
. "` + scriptPath + `"
install_cooper_binary "` + target + `"
echo "RETURNED:$?"
`

	cmd := exec.Command(bash, "-c", program)
	cmd.Env = append(os.Environ(), "PATH="+emptyBin, "HOME="+target)
	out, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("install_cooper_binary aborted when the binary was unobtainable; "+
			"scaffolding must continue without it.\nerror: %v\noutput:\n%s", err, out)
	}
	if !strings.Contains(string(out), "RETURNED:0") {
		t.Errorf("expected install_cooper_binary to return 0 in zero-binary fallback, got:\n%s", out)
	}
}

// TestInstallScript_FailedDownloadPreservesExistingBinary exercises the Tier 2
// download-failure path with a writable bin directory, which the test above
// does not reach.
//
// Writing the download straight to the destination truncates it before the
// transfer completes, so a failed download on an offline or rate-limited host
// silently uninstalled a working cooper. Each tier now stages into a temp file
// and only replaces the destination on success.
func TestInstallScript_FailedDownloadPreservesExistingBinary(t *testing.T) {
	bash := requireBash(t)

	scriptPath, err := filepath.Abs(installScript)
	if err != nil {
		t.Fatalf("failed resolving install.sh: %v", err)
	}

	binDir := t.TempDir()
	existing := filepath.Join(binDir, "cooper")
	const sentinel = "existing working cooper"

	if err := os.WriteFile(existing, []byte("#!/bin/sh\necho \""+sentinel+"\"\n"), 0755); err != nil {
		t.Fatalf("failed seeding an existing binary: %v", err)
	}

	// SCRIPT_DIR is redirected away from the repository so Tier 1 cannot
	// compile, and the release URL is pointed at a path that cannot resolve so
	// Tier 2 must fail.
	program := `
set -e
export COOPER_INSTALL_LIB_ONLY=1
. "` + scriptPath + `"
SCRIPT_DIR="` + t.TempDir() + `"
COOPER_RELEASE_BASE_URL="https://127.0.0.1:1/nonexistent"
install_cooper_binary "` + binDir + `"
echo "RETURNED:$?"
`

	cmd := exec.Command(bash, "-c", program)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("install_cooper_binary aborted on download failure: %v\n%s", err, out)
	}
	if !strings.Contains(string(out), "RETURNED:0") {
		t.Errorf("expected return 0 on download failure, got:\n%s", out)
	}

	data, err := os.ReadFile(existing)
	if err != nil {
		t.Fatalf("a failed download removed the previously installed binary at %s: %v", existing, err)
	}
	if !strings.Contains(string(data), sentinel) {
		t.Errorf("a failed download corrupted the previously installed binary; content is now:\n%s", data)
	}
}

// TestInstallScript_RelocatesTroopReferences enforces the installer spec
// requirement "Troop Reference Relocation" together with the scaffolding
// scenario that every path referenced by the installed AGENTS.md must exist.
//
// install.sh moves TROOP.md into .cooper/TROOP.md, but the AGENTS.md written by
// the Troop installer links to the original root path. Relocating the file
// without rewriting the reference leaves a dangling link that 'cooper validate'
// reports in every freshly scaffolded project.
func TestInstallScript_RelocatesTroopReferences(t *testing.T) {
	bash := requireBash(t)

	scriptPath, err := filepath.Abs(installScript)
	if err != nil {
		t.Fatalf("failed resolving install.sh: %v", err)
	}

	target := t.TempDir()

	agentsPath := filepath.Join(target, "AGENTS.md")
	original := "# Agent Guidelines\n\nSee [Troop Reference](TROOP.md) for worktree commands.\nAlso see `TROOP.md` for details.\n"
	if err := os.WriteFile(agentsPath, []byte(original), 0644); err != nil {
		t.Fatalf("failed seeding AGENTS.md: %v", err)
	}
	if err := os.WriteFile(filepath.Join(target, "TROOP.md"), []byte("# Troop\n"), 0644); err != nil {
		t.Fatalf("failed seeding TROOP.md: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(target, ".cooper"), 0755); err != nil {
		t.Fatalf("failed creating .cooper: %v", err)
	}

	program := `
set -e
export COOPER_INSTALL_LIB_ONLY=1
. "` + scriptPath + `"
cd "` + target + `"
relocate_troop_reference
`

	cmd := exec.Command(bash, "-c", program)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("relocate_troop_reference failed: %v\n%s", err, out)
	}

	if _, statErr := os.Stat(filepath.Join(target, ".cooper", "TROOP.md")); statErr != nil {
		t.Errorf("TROOP.md was not relocated into .cooper/: %v", statErr)
	}

	updated, err := os.ReadFile(agentsPath)
	if err != nil {
		t.Fatalf("failed reading AGENTS.md: %v", err)
	}

	got := string(updated)
	if strings.Contains(got, "](TROOP.md)") {
		t.Errorf("AGENTS.md still contains a dangling markdown link to the relocated TROOP.md:\n%s", got)
	}
	if strings.Contains(got, "`TROOP.md`") {
		t.Errorf("AGENTS.md still contains a dangling inline-code reference to the relocated TROOP.md:\n%s", got)
	}
	if !strings.Contains(got, ".cooper/TROOP.md") {
		t.Errorf("AGENTS.md does not reference the relocated .cooper/TROOP.md:\n%s", got)
	}
}

// TestInstallScript_LibraryModeDoesNotScaffold asserts the sourcing guard used
// by the test above does not itself perform an installation, so sourcing the
// script is side-effect free.
func TestInstallScript_LibraryModeDoesNotScaffold(t *testing.T) {
	bash := requireBash(t)

	scriptPath, err := filepath.Abs(installScript)
	if err != nil {
		t.Fatalf("failed resolving install.sh: %v", err)
	}

	target := t.TempDir()

	program := `
set -e
export COOPER_INSTALL_LIB_ONLY=1
cd "` + target + `"
. "` + scriptPath + `"
echo "SOURCED_OK"
`

	cmd := exec.Command(bash, "-c", program)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("sourcing install.sh in library mode failed: %v\n%s", err, out)
	}
	if !strings.Contains(string(out), "SOURCED_OK") {
		t.Fatalf("expected clean sourcing, got:\n%s", out)
	}

	if _, statErr := os.Stat(filepath.Join(target, ".cooper")); statErr == nil {
		t.Error("sourcing install.sh in library mode scaffolded .cooper/; it must be side-effect free")
	}
}
