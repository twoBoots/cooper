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
	// is the worst case the fallback must survive.
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
