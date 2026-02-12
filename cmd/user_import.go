package cmd

import (
	"context"

	"hepic-cli/internal/api"
	"hepic-cli/internal/output"
	"hepic-cli/internal/user"

	"github.com/spf13/cobra"
)

var userImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import users from a CSV file",
	Long:  "Import users from a CSV file into the HEPIC platform.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		filePath, _ := cmd.Flags().GetString("file")

		result, err := user.Import(context.Background(), client, filePath)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

func init() {
	userCmd.AddCommand(userImportCmd)

	userImportCmd.Flags().String("file", "", "Path to CSV file to import")
	userImportCmd.MarkFlagRequired("file")
}
