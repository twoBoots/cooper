package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestInitCmd(t *testing.T) {
	tmpDir := t.TempDir()

	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"init", "--dir", tmpDir})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error running cooper init: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "Initialized Cooper SDD project") {
		t.Errorf("expected success message, got: %s", out)
	}

	// Check AGENTS.md exists
	if _, err := os.Stat(filepath.Join(tmpDir, "AGENTS.md")); os.IsNotExist(err) {
		t.Error("expected AGENTS.md to be created")
	}
}
