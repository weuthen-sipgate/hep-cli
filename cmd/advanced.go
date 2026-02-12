package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"hepic-cli/internal/admin"
	"hepic-cli/internal/api"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var advancedCmd = &cobra.Command{
	Use:     "advanced",
	Short:   "Manage advanced settings",
	Long:    "Create, list, update, and delete advanced settings on the HEPIC platform.",
	GroupID: "config",
}

var advancedListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all advanced settings",
	Long:  "Retrieve and display all advanced settings from the HEPIC platform.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := admin.ListAdvanced(cmd.Context(), client)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

var advancedGetCmd = &cobra.Command{
	Use:   "get <uuid>",
	Short: "Get an advanced setting",
	Long:  "Retrieve an advanced setting by UUID.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uuid := args[0]

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := admin.GetAdvanced(cmd.Context(), client, uuid)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

var advancedCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new advanced setting",
	Long:  "Create a new advanced setting on the HEPIC platform. Provide the setting data as JSON via --data.",
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

		result, err := admin.CreateAdvanced(cmd.Context(), client, data)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

var advancedUpdateCmd = &cobra.Command{
	Use:   "update <uuid>",
	Short: "Update an advanced setting",
	Long:  "Update an existing advanced setting by UUID. Provide the updated data as JSON via --data.",
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

		result, err := admin.UpdateAdvanced(cmd.Context(), client, uuid, data)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

var advancedDeleteCmd = &cobra.Command{
	Use:   "delete <uuid>",
	Short: "Delete an advanced setting",
	Long:  "Delete an advanced setting by UUID. Requires confirmation unless --force is specified.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uuid := args[0]
		force, _ := cmd.Flags().GetBool("force")

		if !force {
			fmt.Fprintf(os.Stderr, "Delete advanced setting %s? [y/N]: ", uuid)
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

		result, err := admin.DeleteAdvanced(cmd.Context(), client, uuid)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(advancedCmd)

	advancedCmd.AddCommand(advancedListCmd)
	advancedCmd.AddCommand(advancedGetCmd)
	advancedCmd.AddCommand(advancedCreateCmd)
	advancedCmd.AddCommand(advancedUpdateCmd)
	advancedCmd.AddCommand(advancedDeleteCmd)

	advancedCreateCmd.Flags().String("data", "", "Advanced setting data as JSON")
	advancedCreateCmd.MarkFlagRequired("data")

	advancedUpdateCmd.Flags().String("data", "", "Advanced setting data as JSON")
	advancedUpdateCmd.MarkFlagRequired("data")

	advancedDeleteCmd.Flags().Bool("force", false, "Skip confirmation prompt")
}
