package cmd

import (

	"hepic-cli/internal/api"
	"hepic-cli/internal/output"
	"hepic-cli/internal/user"

	"github.com/spf13/cobra"
)

var userCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new user",
	Long:  "Create a new user account on the HEPIC platform.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		username, _ := cmd.Flags().GetString("username")
		email, _ := cmd.Flags().GetString("email")
		password, _ := cmd.Flags().GetString("password")
		firstname, _ := cmd.Flags().GetString("firstname")
		lastname, _ := cmd.Flags().GetString("lastname")
		usergroup, _ := cmd.Flags().GetString("usergroup")
		department, _ := cmd.Flags().GetString("department")
		partid, _ := cmd.Flags().GetUint16("partid")

		data := map[string]interface{}{
			"username":   username,
			"email":      email,
			"password":   password,
			"firstname":  firstname,
			"lastname":   lastname,
			"usergroup":  usergroup,
			"department": department,
			"partid":     partid,
		}

		result, err := user.Create(cmd.Context(), client, data)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	userCmd.AddCommand(userCreateCmd)

	userCreateCmd.Flags().String("username", "", "Username for the new user")
	userCreateCmd.Flags().String("email", "", "Email address")
	userCreateCmd.Flags().String("password", "", "Password")
	userCreateCmd.Flags().String("firstname", "", "First name")
	userCreateCmd.Flags().String("lastname", "", "Last name")
	userCreateCmd.Flags().String("usergroup", "", "User group")
	userCreateCmd.Flags().String("department", "", "Department")
	userCreateCmd.Flags().Uint16("partid", 10, "Partition ID")

	userCreateCmd.MarkFlagRequired("username")
	userCreateCmd.MarkFlagRequired("email")
	userCreateCmd.MarkFlagRequired("password")
	userCreateCmd.MarkFlagRequired("firstname")
	userCreateCmd.MarkFlagRequired("lastname")
	userCreateCmd.MarkFlagRequired("usergroup")
	userCreateCmd.MarkFlagRequired("department")
}
