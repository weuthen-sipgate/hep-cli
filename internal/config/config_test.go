package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
)

func TestLoadFromViper(t *testing.T) {
	viper.Reset()
	viper.Set("host", "https://example.com")
	viper.Set("token", "test-token")
	viper.Set("format", "table")

	cfg := Load()
	if cfg.Host != "https://example.com" {
		t.Errorf("expected host https://example.com, got %s", cfg.Host)
	}
	if cfg.Token != "test-token" {
		t.Errorf("expected token test-token, got %s", cfg.Token)
	}
	if cfg.Format != "table" {
		t.Errorf("expected format table, got %s", cfg.Format)
	}
}

func TestSaveAndLoadFile(t *testing.T) {
	tmpDir := t.TempDir()
	origHome := os.Getenv("HOME")
	t.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", origHome)

	cfg := &Config{
		Host:  "https://hepic.example.com",
		Token: "my-api-key",
	}

	if err := Save(cfg); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	configPath := filepath.Join(tmpDir, ".hepic", "config.yaml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Fatal("config file was not created")
	}

	// Verify file permissions
	info, err := os.Stat(configPath)
	if err != nil {
		t.Fatalf("cannot stat config file: %v", err)
	}
	if info.Mode().Perm() != 0600 {
		t.Errorf("expected file permissions 0600, got %o", info.Mode().Perm())
	}

	// Load via viper
	viper.Reset()
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		t.Fatalf("viper cannot read config: %v", err)
	}

	loaded := Load()
	if loaded.Host != cfg.Host {
		t.Errorf("expected host %s, got %s", cfg.Host, loaded.Host)
	}
	if loaded.Token != cfg.Token {
		t.Errorf("expected token %s, got %s", cfg.Token, loaded.Token)
	}
}

func TestValidate(t *testing.T) {
	viper.Reset()
	if err := Validate(); err == nil {
		t.Error("expected error when host is empty")
	}

	viper.Set("host", "https://example.com")
	if err := Validate(); err == nil {
		t.Error("expected error when token is empty")
	}

	viper.Set("token", "my-token")
	if err := Validate(); err != nil {
		t.Errorf("unexpected validation error: %v", err)
	}
}

func TestEnvOverridesConfig(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	cfg := &Config{Host: "https://from-file.com", Token: "file-token"}
	if err := Save(cfg); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	viper.Reset()
	viper.SetConfigFile(filepath.Join(tmpDir, ".hepic", "config.yaml"))
	viper.SetEnvPrefix("HEPIC")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	t.Setenv("HEPIC_HOST", "https://from-env.com")

	loaded := Load()
	if loaded.Host != "https://from-env.com" {
		t.Errorf("expected env to override config, got host %s", loaded.Host)
	}
	if loaded.Token != "file-token" {
		t.Errorf("expected token from file, got %s", loaded.Token)
	}
}
