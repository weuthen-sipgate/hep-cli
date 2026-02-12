package cmd

import (
	"fmt"
	"io"
	"os"

	"hepic-cli/internal/api"
	"hepic-cli/internal/output"
	"hepic-cli/internal/user"

	"github.com/spf13/cobra"
)

var userExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export users as CSV",
	Long:  "Export all users from the HEPIC platform as a CSV file.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		body, err := user.Export(cmd.Context(), client)
		if err != nil {
			return err
		}
		defer body.Close()

		outFile, _ := cmd.Flags().GetString("output")

		if outFile != "" {
			f, err := os.Create(outFile)
			if err != nil {
				return fmt.Errorf("failed to create output file: %w", err)
			}
			defer f.Close()

			n, err := io.Copy(f, body)
			if err != nil {
				return fmt.Errorf("failed to write output file: %w", err)
			}

			result := map[string]interface{}{
				"status": "ok",
				"file":   outFile,
				"bytes":  n,
			}
			return output.Print(result)
		}

		// Write to stdout if no output file specified
		_, err = io.Copy(os.Stdout, body)
		return err
	},
}

func init() {
	userCmd.AddCommand(userExportCmd)

	userExportCmd.Flags().StringP("output", "o", "", "Output file path (writes to stdout if not specified)")
}
