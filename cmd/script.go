package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"hepic-cli/internal/api"
	"hepic-cli/internal/output"
	"hepic-cli/internal/script"

	"github.com/spf13/cobra"
)

var scriptCmd = &cobra.Command{
	Use:   "script",
	Short: "Manage scripts",
	Long:  "Create, list, update, and delete scripts on the HEPIC platform.",
}

var scriptListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all scripts",
	Long:  "Retrieve and display all scripts from the HEPIC platform.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := script.List(context.Background(), client)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

var scriptCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new script",
	Long:  "Create a new script on the HEPIC platform. Provide the script data as JSON via --data.",
	RunE: func(cmd *cobra.Command, args []string) error {
		dataStr, _ := cmd.Flags().GetString("data")

		var data interface{}
		if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
			return fmt.Errorf("invalid JSON in --data: %w", err)
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := script.Create(context.Background(), client, data)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

var scriptUpdateCmd = &cobra.Command{
	Use:   "update <uuid>",
	Short: "Update an existing script",
	Long:  "Update a script by UUID. Provide the updated data as JSON via --data.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uuid := args[0]
		dataStr, _ := cmd.Flags().GetString("data")

		var data interface{}
		if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
			return fmt.Errorf("invalid JSON in --data: %w", err)
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := script.Update(context.Background(), client, uuid, data)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

var scriptDeleteCmd = &cobra.Command{
	Use:   "delete <uuid>",
	Short: "Delete a script",
	Long:  "Delete a script by UUID. Requires confirmation unless --force is specified.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uuid := args[0]
		force, _ := cmd.Flags().GetBool("force")

		if !force {
			fmt.Fprintf(os.Stderr, "Delete script %s? [y/N]: ", uuid)
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				return fmt.Errorf("failed to read input: %w", err)
			}
			if strings.TrimSpace(strings.ToLower(input)) != "y" {
				return fmt.Errorf("operation cancelled")
			}
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := script.Delete(context.Background(), client, uuid)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(scriptCmd)

	scriptCmd.AddCommand(scriptListCmd)
	scriptCmd.AddCommand(scriptCreateCmd)
	scriptCmd.AddCommand(scriptUpdateCmd)
	scriptCmd.AddCommand(scriptDeleteCmd)

	scriptCreateCmd.Flags().String("data", "", "Script data as JSON")
	scriptCreateCmd.MarkFlagRequired("data")

	scriptUpdateCmd.Flags().String("data", "", "Script data as JSON")
	scriptUpdateCmd.MarkFlagRequired("data")

	scriptDeleteCmd.Flags().Bool("force", false, "Skip confirmation prompt")
}
