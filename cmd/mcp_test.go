package cmd

import (
	"bytes"
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
