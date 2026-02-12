package cmd

import (
	"context"

	"hepic-cli/internal/api"
	"hepic-cli/internal/output"
	"hepic-cli/internal/user"

	"github.com/spf13/cobra"
)

var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users",
	Long:  "Retrieve and display all users from the HEPIC platform.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := user.List(context.Background(), client)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

func init() {
	userCmd.AddCommand(userListCmd)
}
