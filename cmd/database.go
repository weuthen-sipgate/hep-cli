package cmd

import (

	"hepic-cli/internal/api"
	"hepic-cli/internal/output"
	"hepic-cli/internal/statistic"

	"github.com/spf13/cobra"
)

var databaseCmd = &cobra.Command{
	Use:     "database",
	Short:   "Manage database nodes",
	Long:    "Query database node information from the HEPIC platform.",
	GroupID: "admin",
}

var databaseNodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "List database nodes",
	Long:  "Retrieve and display the list of database nodes.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := statistic.NodeList(cmd.Context(), client)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(databaseCmd)
	databaseCmd.AddCommand(databaseNodesCmd)
}
