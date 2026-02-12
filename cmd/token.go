package cmd

import (
	"github.com/spf13/cobra"
)

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Manage API authentication tokens",
	Long:  "Create, list, and delete API authentication tokens for the HEPIC platform.",
}

func init() {
	rootCmd.AddCommand(tokenCmd)
}
