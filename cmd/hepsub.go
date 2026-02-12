package cmd

import (
	"fmt"

	"hepic-cli/internal/api"
	"hepic-cli/internal/config_resources"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var hepsubCmd = &cobra.Command{
	Use:     "hepsub",
	Short:   "Manage HEP subscriptions",
	Long:    "List, create, update, delete, and search HEP subscriptions.",
	GroupID: "config",
}

var hepsubListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all HEP subscriptions",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}
		result, err := config_resources.ListHepsub(cmd.Context(), client)
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

var hepsubCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new HEP subscription",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		hepid, _ := cmd.Flags().GetInt("hepid")
		profile, _ := cmd.Flags().GetString("profile")
		hepAlias, _ := cmd.Flags().GetString("hep-alias")

		data := map[string]interface{}{
			"hepid":     hepid,
			"profile":   profile,
			"hep_alias": hepAlias,
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

		result, err := config_resources.CreateHepsub(cmd.Context(), client, data)
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

var hepsubUpdateCmd = &cobra.Command{
	Use:   "update <uuid>",
	Short: "Update an existing HEP subscription",
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
		if cmd.Flags().Changed("hep-alias") {
			v, _ := cmd.Flags().GetString("hep-alias")
			data["hep_alias"] = v
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

		result, err := config_resources.UpdateHepsub(cmd.Context(), client, uuid, data)
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

var hepsubDeleteCmd = &cobra.Command{
	Use:   "delete <uuid>",
	Short: "Delete a HEP subscription",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		force, _ := cmd.Flags().GetBool("force")
		if !force {
			if !confirmAction("Delete HEP subscription " + args[0] + "?") {
				return fmt.Errorf("operation cancelled")
			}
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}
		result, err := config_resources.DeleteHepsub(cmd.Context(), client, args[0])
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

var hepsubSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search HEP subscription data",
	RunE: func(cmd *cobra.Command, args []string) error {
		dataFlag, _ := cmd.Flags().GetString("data")
		if dataFlag == "" {
			return fmt.Errorf("--data is required")
		}

		searchData, err := parseJSONFlag(dataFlag)
		if err != nil {
			return err
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}
		result, err := config_resources.SearchHepsub(cmd.Context(), client, searchData)
		if err != nil {
			return err
		}
		return output.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(hepsubCmd)

	hepsubCmd.AddCommand(hepsubListCmd)

	hepsubCmd.AddCommand(hepsubCreateCmd)
	hepsubCreateCmd.Flags().Int("hepid", 0, "HEP ID")
	hepsubCreateCmd.Flags().String("profile", "", "Profile name")
	hepsubCreateCmd.Flags().String("hep-alias", "", "HEP alias")
	hepsubCreateCmd.Flags().String("data", "", "Additional data as JSON")

	hepsubCmd.AddCommand(hepsubUpdateCmd)
	hepsubUpdateCmd.Flags().Int("hepid", 0, "HEP ID")
	hepsubUpdateCmd.Flags().String("profile", "", "Profile name")
	hepsubUpdateCmd.Flags().String("hep-alias", "", "HEP alias")
	hepsubUpdateCmd.Flags().String("data", "", "Additional data as JSON")

	hepsubCmd.AddCommand(hepsubDeleteCmd)
	hepsubDeleteCmd.Flags().Bool("force", false, "Skip confirmation prompt")

	hepsubCmd.AddCommand(hepsubSearchCmd)
	hepsubSearchCmd.Flags().String("data", "", "Search criteria as JSON (required)")
	hepsubSearchCmd.MarkFlagRequired("data")
}
