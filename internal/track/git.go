package track

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// GitExecutor executes Git subcommands in a given working directory.
type GitExecutor interface {
	Run(dir string, args ...string) (string, error)
}

// DefaultGitExecutor uses the system git binary.
type DefaultGitExecutor struct{}

// Run executes a git command in dir and returns trimmed stdout.
func (g *DefaultGitExecutor) Run(dir string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	if dir != "" {
		cmd.Dir = dir
	}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("git %s failed: %v: %s", strings.Join(args, " "), err, strings.TrimSpace(stderr.String()))
	}
	return strings.TrimSpace(stdout.String()), nil
}
