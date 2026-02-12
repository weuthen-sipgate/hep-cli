package cmd

import "github.com/spf13/cobra"

var interceptionCmd = &cobra.Command{
	Use:   "interception",
	Short: "Manage call interceptions",
	Long:  "Create, list, update, and delete call interceptions on the HEPIC platform.",
}

func init() {
	rootCmd.AddCommand(interceptionCmd)
}
