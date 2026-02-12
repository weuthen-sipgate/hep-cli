package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"go.yaml.in/yaml/v3"
)

// Config holds the application configuration.
type Config struct {
	Host   string `json:"host" yaml:"host"`
	Token  string `json:"token" yaml:"token"`
	Format string `json:"format,omitempty" yaml:"format,omitempty"`
}

// ConfigDir returns the path to ~/.hepic.
func ConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot determine home directory: %w", err)
	}
	return filepath.Join(home, ".hepic"), nil
}

// ConfigPath returns the full path to the config file.
func ConfigPath() (string, error) {
	dir, err := ConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.yaml"), nil
}

// Load reads the current effective configuration from viper.
func Load() *Config {
	return &Config{
		Host:   viper.GetString("host"),
		Token:  viper.GetString("token"),
		Format: viper.GetString("format"),
	}
}

// Save writes the config to ~/.hepic/config.yaml.
func Save(cfg *Config) error {
	dir, err := ConfigDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("cannot create config directory: %w", err)
	}

	path, err := ConfigPath()
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("cannot marshal config: %w", err)
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("cannot write config file: %w", err)
	}
	return nil
}

// Validate checks that required configuration values are set.
func Validate() error {
	if viper.GetString("host") == "" {
		return fmt.Errorf("host is not configured. Run 'hepic init' or set HEPIC_HOST")
	}
	if viper.GetString("token") == "" {
		return fmt.Errorf("token is not configured. Run 'hepic init' or set HEPIC_TOKEN")
	}
	return nil
}
