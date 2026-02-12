package cmd

import (
	"context"

	"hepic-cli/internal/api"
	"hepic-cli/internal/output"
	"hepic-cli/internal/user"

	"github.com/spf13/cobra"
)

var tokenListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all API tokens",
	Long:  "Retrieve and display all API authentication tokens.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := user.AuthTypes(context.Background(), client)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tokenCmd.AddCommand(tokenListCmd)
}
