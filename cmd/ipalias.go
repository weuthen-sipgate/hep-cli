package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"hepic-cli/internal/api"
	"hepic-cli/internal/config_resources"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var ipaliasCmd = &cobra.Command{
	Use:   "ipalias",
	Short: "Manage IP aliases",
	Long:  "List, create, update, delete, import, and export IP aliases for the HEPIC platform.",
}

var ipaliasListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all IP aliases",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}
		result, err := config_resources.ListAliases(context.Background(), client)
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

var ipaliasCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new IP alias",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		ip, _ := cmd.Flags().GetString("ip")
		alias, _ := cmd.Flags().GetString("alias")
		port, _ := cmd.Flags().GetInt("port")
		mask, _ := cmd.Flags().GetInt("mask")
		group, _ := cmd.Flags().GetString("group")
		servertype, _ := cmd.Flags().GetString("servertype")
		status, _ := cmd.Flags().GetBool("status")

		data := map[string]interface{}{
			"ip":         ip,
			"alias":      alias,
			"port":       port,
			"mask":       mask,
			"group":      group,
			"servertype": servertype,
			"status":     status,
		}

		result, err := config_resources.CreateAlias(context.Background(), client, data)
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

var ipaliasUpdateCmd = &cobra.Command{
	Use:   "update <uuid>",
	Short: "Update an existing IP alias",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		uuid := args[0]
		data := make(map[string]interface{})

		if cmd.Flags().Changed("ip") {
			v, _ := cmd.Flags().GetString("ip")
			data["ip"] = v
		}
		if cmd.Flags().Changed("alias") {
			v, _ := cmd.Flags().GetString("alias")
			data["alias"] = v
		}
		if cmd.Flags().Changed("port") {
			v, _ := cmd.Flags().GetInt("port")
			data["port"] = v
		}
		if cmd.Flags().Changed("mask") {
			v, _ := cmd.Flags().GetInt("mask")
			data["mask"] = v
		}
		if cmd.Flags().Changed("group") {
			v, _ := cmd.Flags().GetString("group")
			data["group"] = v
		}
		if cmd.Flags().Changed("servertype") {
			v, _ := cmd.Flags().GetString("servertype")
			data["servertype"] = v
		}
		if cmd.Flags().Changed("status") {
			v, _ := cmd.Flags().GetBool("status")
			data["status"] = v
		}

		result, err := config_resources.UpdateAlias(context.Background(), client, uuid, data)
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

var ipaliasDeleteCmd = &cobra.Command{
	Use:   "delete <uuid>",
	Short: "Delete an IP alias",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		force, _ := cmd.Flags().GetBool("force")
		if !force {
			if !confirmAction("Delete IP alias " + args[0] + "?") {
				return fmt.Errorf("operation cancelled")
			}
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}
		result, err := config_resources.DeleteAlias(context.Background(), client, args[0])
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

var ipaliasDeleteAllCmd = &cobra.Command{
	Use:   "delete-all",
	Short: "Delete all IP aliases",
	RunE: func(cmd *cobra.Command, args []string) error {
		force, _ := cmd.Flags().GetBool("force")
		if !force {
			if !confirmAction("Delete ALL IP aliases? This cannot be undone") {
				return fmt.Errorf("operation cancelled")
			}
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}
		result, err := config_resources.DeleteAllAliases(context.Background(), client)
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

var ipaliasExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export IP aliases as CSV",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		body, err := config_resources.ExportAliases(context.Background(), client)
		if err != nil {
			return err
		}
		defer body.Close()

		outFile, _ := cmd.Flags().GetString("output")
		if outFile != "" {
			f, err := os.Create(outFile)
			if err != nil {
				return fmt.Errorf("failed to create output file: %w", err)
			}
			defer f.Close()
			n, err := io.Copy(f, body)
			if err != nil {
				return fmt.Errorf("failed to write output file: %w", err)
			}
			result := map[string]interface{}{
				"status": "ok",
				"file":   outFile,
				"bytes":  n,
			}
			return output.Print(result)
		}

		// Write to stdout
		_, err = io.Copy(os.Stdout, body)
		return err
	},
}

var ipaliasImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import IP aliases from a CSV file",
	RunE: func(cmd *cobra.Command, args []string) error {
		filePath, _ := cmd.Flags().GetString("file")
		if filePath == "" {
			return fmt.Errorf("--file is required")
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := config_resources.ImportAliases(context.Background(), client, filePath)
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

// confirmAction prompts the user for confirmation on stderr.
// Returns true if the user answers "y" or "yes".
func confirmAction(prompt string) bool {
	fmt.Fprintf(os.Stderr, "%s [y/N]: ", prompt)
	reader := bufio.NewReader(os.Stdin)
	answer, _ := reader.ReadString('\n')
	answer = strings.TrimSpace(strings.ToLower(answer))
	return answer == "y" || answer == "yes"
}

// parseJSONFlag parses a --data flag value as JSON into a map.
func parseJSONFlag(data string) (map[string]interface{}, error) {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		return nil, fmt.Errorf("invalid JSON in --data flag: %w", err)
	}
	return result, nil
}

func init() {
	rootCmd.AddCommand(ipaliasCmd)

	ipaliasCmd.AddCommand(ipaliasListCmd)

	ipaliasCmd.AddCommand(ipaliasCreateCmd)
	ipaliasCreateCmd.Flags().String("ip", "", "IP address (required)")
	ipaliasCreateCmd.Flags().String("alias", "", "Alias name (required)")
	ipaliasCreateCmd.Flags().Int("port", 0, "Port number")
	ipaliasCreateCmd.Flags().Int("mask", 32, "Network mask")
	ipaliasCreateCmd.Flags().String("group", "", "Group name")
	ipaliasCreateCmd.Flags().String("servertype", "", "Server type")
	ipaliasCreateCmd.Flags().Bool("status", true, "Status (active/inactive)")
	ipaliasCreateCmd.MarkFlagRequired("ip")
	ipaliasCreateCmd.MarkFlagRequired("alias")

	ipaliasCmd.AddCommand(ipaliasUpdateCmd)
	ipaliasUpdateCmd.Flags().String("ip", "", "IP address")
	ipaliasUpdateCmd.Flags().String("alias", "", "Alias name")
	ipaliasUpdateCmd.Flags().Int("port", 0, "Port number")
	ipaliasUpdateCmd.Flags().Int("mask", 32, "Network mask")
	ipaliasUpdateCmd.Flags().String("group", "", "Group name")
	ipaliasUpdateCmd.Flags().String("servertype", "", "Server type")
	ipaliasUpdateCmd.Flags().Bool("status", true, "Status (active/inactive)")

	ipaliasCmd.AddCommand(ipaliasDeleteCmd)
	ipaliasDeleteCmd.Flags().Bool("force", false, "Skip confirmation prompt")

	ipaliasCmd.AddCommand(ipaliasDeleteAllCmd)
	ipaliasDeleteAllCmd.Flags().Bool("force", false, "Skip confirmation prompt")

	ipaliasCmd.AddCommand(ipaliasExportCmd)
	ipaliasExportCmd.Flags().StringP("output", "o", "", "Output file path (writes to stdout if not set)")

	ipaliasCmd.AddCommand(ipaliasImportCmd)
	ipaliasImportCmd.Flags().String("file", "", "CSV file to import (required)")
	ipaliasImportCmd.MarkFlagRequired("file")
}
