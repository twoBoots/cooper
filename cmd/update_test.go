package cmd

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/twoBoots/bender/pkg/updater"
)

func TestUpdateCmd_CheckAvailable(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assetName, _ := updater.GetBinaryNameForPlatform("cooper", runtime.GOOS, runtime.GOARCH)
		release := updater.Release{
			TagName: "v2.0.0",
			Name:    "Release v2.0.0",
			Assets: []updater.Asset{
				{
					Name:               assetName,
					BrowserDownloadURL: "http://example.com/asset",
					Size:               1024,
				},
			},
		}
		_ = json.NewEncoder(w).Encode(release)
	}))
	defer server.Close()

	client := updater.NewClientWithBaseURL(server.URL)
	SetUpdaterClient(client)
	defer SetUpdaterClient(nil)

	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(errBuf)
	rootCmd.SetArgs([]string{"update", "--check", "--repo", "twoBoots/cooper"})

	origVersion := Version
	Version = "1.0.0"
	defer func() { Version = origVersion }()

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error executing update --check: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "v2.0.0") {
		t.Errorf("expected output to mention v2.0.0, got: %s", out)
	}
	if !strings.Contains(out, "Update available") && !strings.Contains(out, "update available") && !strings.Contains(out, "Run 'cooper update'") {
		t.Errorf("expected output to prompt update, got: %s", out)
	}
}

func TestUpdateCmd_AlreadyUpToDate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		release := updater.Release{
			TagName: "v1.0.0",
			Name:    "Release v1.0.0",
			Assets:  []updater.Asset{},
		}
		_ = json.NewEncoder(w).Encode(release)
	}))
	defer server.Close()

	client := updater.NewClientWithBaseURL(server.URL)
	SetUpdaterClient(client)
	defer SetUpdaterClient(nil)

	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(errBuf)
	rootCmd.SetArgs([]string{"update", "--check"})

	origVersion := Version
	Version = "1.0.0"
	defer func() { Version = origVersion }()

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "already up to date") && !strings.Contains(out, "latest version") {
		t.Errorf("expected already up to date message, got: %s", out)
	}
}

func TestUpdateCmd_ApplyUpdate(t *testing.T) {
	tmpDir := t.TempDir()
	fakeBinary := filepath.Join(tmpDir, "cooper")
	if err := os.WriteFile(fakeBinary, []byte("#!/bin/sh\necho cooper-v1"), 0755); err != nil {
		t.Fatalf("failed to create fake binary: %v", err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "asset.bin") {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("#!/bin/sh\necho cooper-v2"))
			return
		}
		assetName, err := updater.GetBinaryNameForPlatform("cooper", runtime.GOOS, runtime.GOARCH)
		if err != nil {
			t.Fatalf("unexpected error getting binary name: %v", err)
		}
		release := updater.Release{
			TagName: "v2.0.0",
			Name:    "Release v2.0.0",
			Assets: []updater.Asset{
				{
					Name:               assetName,
					BrowserDownloadURL: "http://" + r.Host + "/asset.bin",
					Size:               100,
				},
			},
		}
		_ = json.NewEncoder(w).Encode(release)
	}))
	defer server.Close()

	client := updater.NewClientWithBaseURL(server.URL)
	SetUpdaterClient(client)
	defer SetUpdaterClient(nil)

	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(errBuf)
	rootCmd.SetArgs([]string{"update", "--exec-path", fakeBinary})

	origVersion := Version
	Version = "1.0.0"
	defer func() { Version = origVersion }()

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("unexpected error executing update: %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "Successfully updated") && !strings.Contains(out, "v2.0.0") {
		t.Errorf("expected success update message, got: %s", out)
	}

	content, err := os.ReadFile(fakeBinary)
	if err != nil {
		t.Fatalf("failed to read binary: %v", err)
	}
	if !strings.Contains(string(content), "cooper-v2") {
		t.Errorf("expected replaced binary content, got: %s", string(content))
	}
}

func TestUpdateCmd_NetworkError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := updater.NewClientWithBaseURL(server.URL)
	SetUpdaterClient(client)
	defer SetUpdaterClient(nil)

	rootCmd := NewRootCmd()
	buf := new(bytes.Buffer)
	errBuf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(errBuf)
	rootCmd.SetArgs([]string{"update"})

	origVersion := Version
	Version = "1.0.0"
	defer func() { Version = origVersion }()

	err := rootCmd.Execute()
	if err == nil {
		t.Fatal("expected error on server 500, got nil")
	}
}
