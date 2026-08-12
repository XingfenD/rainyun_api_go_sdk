package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadTOML(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.toml")
	content := `
api_key = "sk-test123"
output = "json"
`
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		t.Fatal(err)
	}
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if cfg.APIKey != "sk-test123" {
		t.Errorf("APIKey = %q, want %q", cfg.APIKey, "sk-test123")
	}
	if cfg.Output != "json" {
		t.Errorf("Output = %q, want %q", cfg.Output, "json")
	}
}

func TestLoadDefaults(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "empty.toml")
	if err := os.WriteFile(path, []byte(""), 0600); err != nil {
		t.Fatal(err)
	}
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if cfg.Output != "table" {
		t.Errorf("default Output = %q, want %q", cfg.Output, "table")
	}
}

func TestLoadMissingFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "missing.toml")
	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load() missing file error: %v", err)
	}
	if cfg.Output != "table" {
		t.Errorf("default Output = %q, want %q", cfg.Output, "table")
	}
}

func TestSave(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.toml")
	cfg := &Config{APIKey: "secret", Output: "json"}
	if err := Save(path, cfg); err != nil {
		t.Fatalf("Save() error: %v", err)
	}
	loaded, err := Load(path)
	if err != nil {
		t.Fatalf("re-Load() error: %v", err)
	}
	if loaded.APIKey != "secret" {
		t.Errorf("roundtrip APIKey = %q", loaded.APIKey)
	}
	if loaded.Output != "json" {
		t.Errorf("roundtrip Output = %q", loaded.Output)
	}
}

func TestDefaultPath(t *testing.T) {
	path := DefaultPath()
	if path == "" {
		t.Error("DefaultPath() returned empty")
	}
}

func TestEnsureDir(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "sub", "dir")
	path := filepath.Join(dir, "config.toml")
	cfg := &Config{Output: "table"}
	if err := Save(path, cfg); err != nil {
		t.Fatalf("Save() to deep dir error: %v", err)
	}
}
