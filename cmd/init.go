package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"hepic-cli/internal/config"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize hepic-cli configuration",
	Long: `Set up the HEPIC API connection by providing host URL and API token.

Interactive mode (default):
  hepic init

Non-interactive mode:
  hepic init --host https://hepic.example.com --token your-api-key`,
	RunE: runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) error {
	host, _ := cmd.Flags().GetString("host")
	token, _ := cmd.Flags().GetString("token")

	// Interactive mode if flags not provided
	if host == "" || token == "" {
		reader := bufio.NewReader(os.Stdin)

		if host == "" {
			fmt.Fprint(os.Stderr, "HEPIC API Host URL: ")
			input, err := reader.ReadString('\n')
			if err != nil {
				return fmt.Errorf("failed to read input: %w", err)
			}
			host = strings.TrimSpace(input)
		}

		if token == "" {
			fmt.Fprint(os.Stderr, "API Token: ")
			input, err := reader.ReadString('\n')
			if err != nil {
				return fmt.Errorf("failed to read input: %w", err)
			}
			token = strings.TrimSpace(input)
		}
	}

	host = strings.TrimRight(host, "/")

	if host == "" {
		return fmt.Errorf("host is required")
	}
	if token == "" {
		return fmt.Errorf("token is required")
	}

	// Validate connection
	fmt.Fprintf(os.Stderr, "Validating connection to %s...\n", host)
	if err := validateConnection(host, token); err != nil {
		return fmt.Errorf("connection validation failed: %w", err)
	}

	cfg := &config.Config{
		Host:  host,
		Token: token,
	}

	if err := config.Save(cfg); err != nil {
		return fmt.Errorf("failed to save configuration: %w", err)
	}

	configPath, _ := config.ConfigPath()

	result := map[string]string{
		"status":      "ok",
		"config_path": configPath,
		"host":        host,
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.Encode(result)

	return nil
}

func validateConnection(host, token string) error {
	client := &http.Client{Timeout: 10 * time.Second}

	url := host + "/api/v3/version/api/info"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("invalid host URL: %w", err)
	}
	req.Header.Set("Auth-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("cannot reach host: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized || resp.StatusCode == http.StatusForbidden {
		return fmt.Errorf("authentication failed (HTTP %d) — check your API token", resp.StatusCode)
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("server returned HTTP %d", resp.StatusCode)
	}
	return nil
}
