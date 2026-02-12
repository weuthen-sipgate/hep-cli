package cmd

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"hepic-cli/internal/admin"
	"hepic-cli/internal/api"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var settingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Manage user settings",
	Long:  "Create, list, update, and delete user settings on the HEPIC platform.",
}

var settingsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all user settings",
	Long:  "Retrieve and display all user settings from the HEPIC platform.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := admin.ListSettings(context.Background(), client)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

var settingsGetCmd = &cobra.Command{
	Use:   "get <category>",
	Short: "Get settings by category",
	Long:  "Retrieve user settings for a specific category.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		category := args[0]

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := admin.GetSettingsByCategory(context.Background(), client, category)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

var settingsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user setting",
	Long:  "Create a new user setting on the HEPIC platform. Provide the setting data as JSON via --data.",
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

		result, err := admin.CreateSetting(context.Background(), client, data)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

var settingsUpdateCmd = &cobra.Command{
	Use:   "update <uuid>",
	Short: "Update a user setting",
	Long:  "Update an existing user setting by UUID. Provide the updated data as JSON via --data.",
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

		result, err := admin.UpdateSetting(context.Background(), client, uuid, data)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

var settingsDeleteCmd = &cobra.Command{
	Use:   "delete <uuid>",
	Short: "Delete a user setting",
	Long:  "Delete a user setting by UUID. Requires confirmation unless --force is specified.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uuid := args[0]
		force, _ := cmd.Flags().GetBool("force")

		if !force {
			fmt.Fprintf(os.Stderr, "Delete setting %s? [y/N]: ", uuid)
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

		result, err := admin.DeleteSetting(context.Background(), client, uuid)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(settingsCmd)

	settingsCmd.AddCommand(settingsListCmd)
	settingsCmd.AddCommand(settingsGetCmd)
	settingsCmd.AddCommand(settingsCreateCmd)
	settingsCmd.AddCommand(settingsUpdateCmd)
	settingsCmd.AddCommand(settingsDeleteCmd)

	settingsCreateCmd.Flags().String("data", "", "Setting data as JSON")
	settingsCreateCmd.MarkFlagRequired("data")

	settingsUpdateCmd.Flags().String("data", "", "Setting data as JSON")
	settingsUpdateCmd.MarkFlagRequired("data")

	settingsDeleteCmd.Flags().Bool("force", false, "Skip confirmation prompt")
}
