package cmd

import (

	"hepic-cli/internal/api"
	"hepic-cli/internal/output"
	"hepic-cli/internal/statistic"

	"github.com/spf13/cobra"
)

var grafanaCmd = &cobra.Command{
	Use:     "grafana",
	Short:   "Interact with Grafana via proxy",
	Long:    "Retrieve Grafana dashboards, folders, organization info, and status via the HEPIC proxy.",
	GroupID: "monitoring",
}

var grafanaDashboardCmd = &cobra.Command{
	Use:   "dashboard <uid>",
	Short: "Get a Grafana dashboard by UID",
	Long: `Retrieve a Grafana dashboard by its UID.

Examples:
  hepic grafana dashboard abc123`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uid := args[0]

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := statistic.GetDashboard(cmd.Context(), client, uid)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

var grafanaFoldersCmd = &cobra.Command{
	Use:   "folders",
	Short: "List Grafana folders",
	Long:  "Retrieve all folders from Grafana.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := statistic.Folders(cmd.Context(), client)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

var grafanaOrgCmd = &cobra.Command{
	Use:   "org",
	Short: "Show Grafana organization info",
	Long:  "Retrieve organization information from Grafana.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := statistic.Org(cmd.Context(), client)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

var grafanaSearchCmd = &cobra.Command{
	Use:   "search <uid>",
	Short: "Search for a Grafana dashboard",
	Long: `Search for a Grafana dashboard by UID.

Examples:
  hepic grafana search abc123`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uid := args[0]

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := statistic.SearchDashboard(cmd.Context(), client, uid)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

var grafanaStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show Grafana connection status",
	Long:  "Retrieve the Grafana connection status.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := statistic.Status(cmd.Context(), client)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(grafanaCmd)

	grafanaCmd.AddCommand(grafanaDashboardCmd)
	grafanaCmd.AddCommand(grafanaFoldersCmd)
	grafanaCmd.AddCommand(grafanaOrgCmd)
	grafanaCmd.AddCommand(grafanaSearchCmd)
	grafanaCmd.AddCommand(grafanaStatusCmd)
}
