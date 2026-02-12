package cmd

import (
	"github.com/spf13/cobra"
)

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users and authentication",
	Long:  "Create, list, update, and delete users. Manage user groups, import/export, and password changes.",
}

func init() {
	rootCmd.AddCommand(userCmd)
}
