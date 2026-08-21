package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMCPCmd_ServeStdio(t *testing.T) {
	rootCmd := NewRootCmd()
	inBuf := bytes.NewBufferString(`{"jsonrpc":"2.0","id":1,"method":"initialize"}` + "\n")
	outBuf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)

	rootCmd.SetIn(inBuf)
	rootCmd.SetOut(outBuf)
	rootCmd.SetErr(errBuf)
	rootCmd.SetArgs([]string{"mcp"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error executing cooper mcp: %v", err)
	}

	if !strings.Contains(outBuf.String(), "cooper-mcp") {
		t.Errorf("expected cooper-mcp in server initialize response, got: %s", outBuf.String())
	}
}

func TestMCPCmd_Help(t *testing.T) {
	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetArgs([]string{"mcp", "--help"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error executing mcp --help: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "Model Context Protocol") {
		t.Errorf("expected MCP help text, got: %s", out)
	}
}

func TestMCPInstallCmd_Run(t *testing.T) {
	tmpDir := t.TempDir()
	homeDir := filepath.Join(tmpDir, "home")
	_ = os.MkdirAll(homeDir, 0755)

	out := new(bytes.Buffer)
	err := RunMCPInstall(out, tmpDir, homeDir, []string{"cursor"}, false, true)
	if err != nil {
		t.Fatalf("unexpected error running RunMCPInstall: %v", err)
	}

	cursorConfig := filepath.Join(tmpDir, ".cursor", "mcp.json")
	data, err := os.ReadFile(cursorConfig)
	if err != nil {
		t.Fatalf("expected .cursor/mcp.json to exist: %v", err)
	}

	if !strings.Contains(string(data), `"cooper"`) {
		t.Errorf("expected cooper in cursor config, got: %s", string(data))
	}
}

func TestMCPInstallCmd_All(t *testing.T) {
	tmpDir := t.TempDir()
	homeDir := filepath.Join(tmpDir, "home")
	_ = os.MkdirAll(homeDir, 0755)

	out := new(bytes.Buffer)
	err := RunMCPInstall(out, tmpDir, homeDir, nil, true, false)
	if err != nil {
		t.Fatalf("unexpected error with all: %v", err)
	}

	if !strings.Contains(out.String(), "Configured") {
		t.Errorf("expected configuration output, got: %s", out.String())
	}
}

func TestMCPInstallCmd_UnknownClient(t *testing.T) {
	tmpDir := t.TempDir()
	out := new(bytes.Buffer)
	err := RunMCPInstall(out, tmpDir, tmpDir, []string{"unknown-client"}, false, false)
	if err == nil {
		t.Fatal("expected error for unknown client ID, got nil")
	}
}

func TestMCPInstallCmd_CLIExecution(t *testing.T) {
	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetArgs([]string{"mcp", "install", "--help"})

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error on mcp install --help: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "Configure Cooper MCP server") && !strings.Contains(out, "cursor") {
		t.Errorf("expected install help output, got: %s", out)
	}

	// Test executing mcp install via CLI flags
	cliCmd := NewRootCmd()
	cliBuf := new(bytes.Buffer)
	cliCmd.SetOut(cliBuf)
	cliCmd.SetArgs([]string{"mcp", "install", "--client", "cursor", "--non-interactive"})
	if err := cliCmd.Execute(); err != nil {
		t.Fatalf("unexpected error running mcp install via CLI: %v", err)
	}
}
