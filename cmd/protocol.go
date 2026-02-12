package cmd

import (
	"context"
	"fmt"

	"hepic-cli/internal/api"
	"hepic-cli/internal/config_resources"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var protocolCmd = &cobra.Command{
	Use:   "protocol",
	Short: "Manage protocol definitions",
	Long:  "Search, create, update, and delete protocol definitions.",
}

var protocolSearchCmd = &cobra.Command{
	Use:   "search <id>",
	Short: "Search for a protocol by ID",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}
		result, err := config_resources.SearchProtocol(context.Background(), client, args[0])
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

var protocolCreateCmd = &cobra.Command{
	Use:   "create <id>",
	Short: "Create a new protocol definition",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dataFlag, _ := cmd.Flags().GetString("data")
		if dataFlag == "" {
			return fmt.Errorf("--data is required")
		}

		data, err := parseJSONFlag(dataFlag)
		if err != nil {
			return err
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}
		result, err := config_resources.CreateProtocol(context.Background(), client, args[0], data)
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

var protocolUpdateCmd = &cobra.Command{
	Use:   "update <uuid>",
	Short: "Update an existing protocol definition",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dataFlag, _ := cmd.Flags().GetString("data")
		if dataFlag == "" {
			return fmt.Errorf("--data is required")
		}

		data, err := parseJSONFlag(dataFlag)
		if err != nil {
			return err
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}
		result, err := config_resources.UpdateProtocol(context.Background(), client, args[0], data)
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

var protocolDeleteCmd = &cobra.Command{
	Use:   "delete <uuid>",
	Short: "Delete a protocol definition",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		force, _ := cmd.Flags().GetBool("force")
		if !force {
			if !confirmAction("Delete protocol " + args[0] + "?") {
				return fmt.Errorf("operation cancelled")
			}
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}
		result, err := config_resources.DeleteProtocol(context.Background(), client, args[0])
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(protocolCmd)

	protocolCmd.AddCommand(protocolSearchCmd)

	protocolCmd.AddCommand(protocolCreateCmd)
	protocolCreateCmd.Flags().String("data", "", "Protocol data as JSON (required)")
	protocolCreateCmd.MarkFlagRequired("data")

	protocolCmd.AddCommand(protocolUpdateCmd)
	protocolUpdateCmd.Flags().String("data", "", "Protocol data as JSON (required)")
	protocolUpdateCmd.MarkFlagRequired("data")

	protocolCmd.AddCommand(protocolDeleteCmd)
	protocolDeleteCmd.Flags().Bool("force", false, "Skip confirmation prompt")
}
