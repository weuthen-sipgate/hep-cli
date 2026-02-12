package cmd

import (
	"context"

	"hepic-cli/internal/api"
	"hepic-cli/internal/output"
	"hepic-cli/internal/user"

	"github.com/spf13/cobra"
)

var tokenCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new API token",
	Long:  "Create a new API authentication token with the specified name.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		name, _ := cmd.Flags().GetString("name")

		data := map[string]interface{}{
			"name": name,
		}

		result, err := user.CreateToken(context.Background(), client, data)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tokenCmd.AddCommand(tokenCreateCmd)

	tokenCreateCmd.Flags().String("name", "", "Name for the new token")
	tokenCreateCmd.MarkFlagRequired("name")
}
