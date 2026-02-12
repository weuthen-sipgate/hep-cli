package cmd

import (

	"hepic-cli/internal/api"
	"hepic-cli/internal/output"
	"hepic-cli/internal/user"

	"github.com/spf13/cobra"
)

var userPasswordCmd = &cobra.Command{
	Use:   "password <uuid>",
	Short: "Change a user's password",
	Long:  "Update the password for the specified user UUID.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uuid := args[0]

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		password, _ := cmd.Flags().GetString("password")

		data := map[string]interface{}{
			"password": password,
		}

		result, err := user.UpdatePassword(cmd.Context(), client, uuid, data)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	userCmd.AddCommand(userPasswordCmd)

	userPasswordCmd.Flags().String("password", "", "New password")
	userPasswordCmd.MarkFlagRequired("password")
}
