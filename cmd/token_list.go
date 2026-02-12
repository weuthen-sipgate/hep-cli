package cmd

import (

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

		result, err := user.AuthTypes(cmd.Context(), client)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tokenCmd.AddCommand(tokenListCmd)
}
