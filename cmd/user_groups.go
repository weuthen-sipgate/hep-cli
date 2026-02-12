package cmd

import (

	"hepic-cli/internal/api"
	"hepic-cli/internal/output"
	"hepic-cli/internal/user"

	"github.com/spf13/cobra"
)

var userGroupsCmd = &cobra.Command{
	Use:   "groups",
	Short: "List user groups",
	Long:  "Retrieve and display all available user groups.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := user.Groups(cmd.Context(), client)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	userCmd.AddCommand(userGroupsCmd)
}
