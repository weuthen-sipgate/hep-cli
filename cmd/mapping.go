package cmd

import (
	"context"
	"fmt"

	"hepic-cli/internal/api"
	"hepic-cli/internal/config_resources"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var mappingCmd = &cobra.Command{
	Use:   "mapping",
	Short: "Manage protocol mappings",
	Long:  "List, create, update, delete, and reset protocol mappings.",
}

var mappingListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all protocol mappings",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}
		result, err := config_resources.ListMappings(context.Background(), client)
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

var mappingCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new protocol mapping",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		hepid, _ := cmd.Flags().GetInt("hepid")
		profile, _ := cmd.Flags().GetString("profile")

		data := map[string]interface{}{
			"hepid":   hepid,
			"profile": profile,
		}

		dataFlag, _ := cmd.Flags().GetString("data")
		if dataFlag != "" {
			extra, err := parseJSONFlag(dataFlag)
			if err != nil {
				return err
			}
			for k, v := range extra {
				data[k] = v
			}
		}

		result, err := config_resources.CreateMapping(context.Background(), client, data)
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

var mappingUpdateCmd = &cobra.Command{
	Use:   "update <uuid>",
	Short: "Update an existing protocol mapping",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		uuid := args[0]
		data := make(map[string]interface{})

		if cmd.Flags().Changed("hepid") {
			v, _ := cmd.Flags().GetInt("hepid")
			data["hepid"] = v
		}
		if cmd.Flags().Changed("profile") {
			v, _ := cmd.Flags().GetString("profile")
			data["profile"] = v
		}

		dataFlag, _ := cmd.Flags().GetString("data")
		if dataFlag != "" {
			extra, err := parseJSONFlag(dataFlag)
			if err != nil {
				return err
			}
			for k, v := range extra {
				data[k] = v
			}
		}

		result, err := config_resources.UpdateMapping(context.Background(), client, uuid, data)
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

var mappingDeleteCmd = &cobra.Command{
	Use:   "delete <uuid>",
	Short: "Delete a protocol mapping",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		force, _ := cmd.Flags().GetBool("force")
		if !force {
			if !confirmAction("Delete mapping " + args[0] + "?") {
				return fmt.Errorf("operation cancelled")
			}
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}
		result, err := config_resources.DeleteMapping(context.Background(), client, args[0])
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

var mappingProtocolsCmd = &cobra.Command{
	Use:   "protocols",
	Short: "List all protocol definitions",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}
		result, err := config_resources.ListAllProtocols(context.Background(), client)
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

var mappingResetCmd = &cobra.Command{
	Use:   "reset [uuid]",
	Short: "Reset protocol mappings to defaults",
	Long:  "Reset all protocol mappings to defaults, or reset a single mapping by UUID.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		if len(args) == 1 {
			// Reset single mapping
			result, err := config_resources.ResetOne(context.Background(), client, args[0])
			if err != nil {
				return err
			}
			return output.Print(result)
		}

		// Reset all mappings
		force, _ := cmd.Flags().GetBool("force")
		if !force {
			if !confirmAction("Reset ALL protocol mappings to defaults? This cannot be undone") {
				return fmt.Errorf("operation cancelled")
			}
		}

		result, err := config_resources.ResetAll(context.Background(), client)
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(mappingCmd)

	mappingCmd.AddCommand(mappingListCmd)

	mappingCmd.AddCommand(mappingCreateCmd)
	mappingCreateCmd.Flags().Int("hepid", 0, "HEP ID")
	mappingCreateCmd.Flags().String("profile", "", "Profile name")
	mappingCreateCmd.Flags().String("data", "", "Additional mapping data as JSON")

	mappingCmd.AddCommand(mappingUpdateCmd)
	mappingUpdateCmd.Flags().Int("hepid", 0, "HEP ID")
	mappingUpdateCmd.Flags().String("profile", "", "Profile name")
	mappingUpdateCmd.Flags().String("data", "", "Additional mapping data as JSON")

	mappingCmd.AddCommand(mappingDeleteCmd)
	mappingDeleteCmd.Flags().Bool("force", false, "Skip confirmation prompt")

	mappingCmd.AddCommand(mappingProtocolsCmd)

	mappingCmd.AddCommand(mappingResetCmd)
	mappingResetCmd.Flags().Bool("force", false, "Skip confirmation prompt")
}
