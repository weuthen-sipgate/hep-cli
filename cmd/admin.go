package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"hepic-cli/internal/admin"
	"hepic-cli/internal/api"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var adminCmd = &cobra.Command{
	Use:     "admin",
	Short:   "Admin operations and system information",
	Long:    "Manage admin profiles, config database, and view system version information.",
	GroupID: "admin",
}

var adminProfilesCmd = &cobra.Command{
	Use:   "profiles",
	Short: "List admin profiles",
	Long:  "Retrieve and display admin profiles from the HEPIC platform.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := admin.Profiles(cmd.Context(), client)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

var adminConfigDBCmd = &cobra.Command{
	Use:   "configdb",
	Short: "List config database tables",
	Long:  "Retrieve and display the list of config database tables.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := admin.ConfigDBTables(cmd.Context(), client)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

var adminResyncCmd = &cobra.Command{
	Use:   "resync",
	Short: "Resync config database",
	Long:  "Trigger a resynchronization of the config database. Requires confirmation unless --force is specified.",
	RunE: func(cmd *cobra.Command, args []string) error {
		force, _ := cmd.Flags().GetBool("force")

		if !force {
			fmt.Fprintf(os.Stderr, "Resync config database? [y/N]: ")
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

		result, err := admin.ConfigDBResync(cmd.Context(), client)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

var adminVersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show API and UI version information",
	Long:  "Retrieve and display both API and UI version information from the HEPIC platform.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		apiVersion, err := admin.APIVersion(cmd.Context(), client)
		if err != nil {
			return err
		}

		uiVersion, err := admin.UIVersion(cmd.Context(), client)
		if err != nil {
			return err
		}

		combined := map[string]interface{}{
			"api": apiVersion,
			"ui":  uiVersion,
		}

		return output.Print(combined)
	},
}

func init() {
	rootCmd.AddCommand(adminCmd)

	adminCmd.AddCommand(adminProfilesCmd)
	adminCmd.AddCommand(adminConfigDBCmd)
	adminCmd.AddCommand(adminResyncCmd)
	adminCmd.AddCommand(adminVersionCmd)

	adminResyncCmd.Flags().Bool("force", false, "Skip confirmation prompt")
}
