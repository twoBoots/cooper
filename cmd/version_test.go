package cmd

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestRootCmd(t *testing.T) {
	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"--help"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error executing root help: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "cooper is the CLI for the Cooper Spec-Driven Development framework") {
		t.Errorf("expected output to contain description, got: %s", output)
	}
	if !strings.Contains(output, "--verbose") {
		t.Errorf("expected output to contain --verbose flag, got: %s", output)
	}
	if !strings.Contains(output, "--non-interactive") {
		t.Errorf("expected output to contain --non-interactive flag, got: %s", output)
	}
}

func TestExecute(t *testing.T) {
	// Set args to version so Execute doesn't run interactive/default root
	osArgs := make([]string, len(os.Args))
	copy(osArgs, os.Args)
	defer func() { os.Args = osArgs }()

	os.Args = []string{"cooper", "version"}
	if err := Execute(); err != nil {
		t.Errorf("unexpected error executing root: %v", err)
	}
}

func TestVersionCmd_Text(t *testing.T) {
	Version = "1.0.0"
	Commit = "abcdef"
	Date = "2026-08-20"

	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"version"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "cooper version 1.0.0 (commit: abcdef, built at: 2026-08-20)") {
		t.Errorf("unexpected version output: %s", out)
	}
}

func TestVersionCmd_JSON(t *testing.T) {
	Version = "1.0.0"
	Commit = "abcdef"
	Date = "2026-08-20"

	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs([]string{"version", "--json"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var data map[string]string
	if err := json.Unmarshal(buf.Bytes(), &data); err != nil {
		t.Fatalf("failed to parse json output: %v, raw: %s", err, buf.String())
	}

	if data["version"] != "1.0.0" || data["commit"] != "abcdef" || data["date"] != "2026-08-20" {
		t.Errorf("mismatched json fields: %+v", data)
	}
}
