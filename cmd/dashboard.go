package cmd

import (
	"encoding/json"
	"fmt"

	"hepic-cli/internal/api"
	"hepic-cli/internal/dashboard"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var dashboardCmd = &cobra.Command{
	Use:     "dashboard",
	Short:   "Manage dashboards",
	Long:    "List, update, and delete dashboards on the HEPIC platform.",
	GroupID: "data",
}

var dashboardListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all dashboards",
	Long:  "Retrieve and display all dashboards from the HEPIC platform.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := dashboard.List(cmd.Context(), client)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

var dashboardUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Create or update a dashboard",
	Long: `Create or update a dashboard by ID.

Examples:
  hepic dashboard update my-dashboard --data '{"name":"My Dashboard","type":1}'`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dashboardID := args[0]
		dataStr, _ := cmd.Flags().GetString("data")
		if dataStr == "" {
			return fmt.Errorf("--data flag is required")
		}

		var data interface{}
		if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
			return fmt.Errorf("invalid JSON in --data: %w", err)
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := dashboard.Store(cmd.Context(), client, dashboardID, data)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

var dashboardDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a dashboard",
	Long: `Delete a dashboard by ID. Requires --force to confirm.

Examples:
  hepic dashboard delete my-dashboard --force`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		force, _ := cmd.Flags().GetBool("force")
		if !force {
			return fmt.Errorf("--force flag is required to confirm deletion")
		}

		dashboardID := args[0]

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := dashboard.Delete(cmd.Context(), client, dashboardID)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(dashboardCmd)

	dashboardCmd.AddCommand(dashboardListCmd)

	dashboardCmd.AddCommand(dashboardUpdateCmd)
	dashboardUpdateCmd.Flags().String("data", "", "Dashboard data as JSON string (required)")

	dashboardCmd.AddCommand(dashboardDeleteCmd)
	dashboardDeleteCmd.Flags().Bool("force", false, "Confirm deletion")
}
