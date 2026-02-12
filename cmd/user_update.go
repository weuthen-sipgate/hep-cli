package cmd

import (
	"context"
	"fmt"

	"hepic-cli/internal/api"
	"hepic-cli/internal/output"
	"hepic-cli/internal/user"

	"github.com/spf13/cobra"
)

var userUpdateCmd = &cobra.Command{
	Use:   "update <uuid>",
	Short: "Update an existing user",
	Long:  "Update user fields for the specified user UUID.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uuid := args[0]

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		data := make(map[string]interface{})

		if cmd.Flags().Changed("username") {
			v, _ := cmd.Flags().GetString("username")
			data["username"] = v
		}
		if cmd.Flags().Changed("email") {
			v, _ := cmd.Flags().GetString("email")
			data["email"] = v
		}
		if cmd.Flags().Changed("firstname") {
			v, _ := cmd.Flags().GetString("firstname")
			data["firstname"] = v
		}
		if cmd.Flags().Changed("lastname") {
			v, _ := cmd.Flags().GetString("lastname")
			data["lastname"] = v
		}
		if cmd.Flags().Changed("usergroup") {
			v, _ := cmd.Flags().GetString("usergroup")
			data["usergroup"] = v
		}
		if cmd.Flags().Changed("department") {
			v, _ := cmd.Flags().GetString("department")
			data["department"] = v
		}
		if cmd.Flags().Changed("partid") {
			v, _ := cmd.Flags().GetUint16("partid")
			data["partid"] = v
		}

		if len(data) == 0 {
			return fmt.Errorf("no fields specified to update")
		}

		result, err := user.Update(context.Background(), client, uuid, data)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

func init() {
	userCmd.AddCommand(userUpdateCmd)

	userUpdateCmd.Flags().String("username", "", "Username")
	userUpdateCmd.Flags().String("email", "", "Email address")
	userUpdateCmd.Flags().String("firstname", "", "First name")
	userUpdateCmd.Flags().String("lastname", "", "Last name")
	userUpdateCmd.Flags().String("usergroup", "", "User group")
	userUpdateCmd.Flags().String("department", "", "Department")
	userUpdateCmd.Flags().Uint16("partid", 0, "Partition ID")
}
