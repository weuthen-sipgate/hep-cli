package cmd

import (
	"context"

	"hepic-cli/internal/api"
	"hepic-cli/internal/output"
	"hepic-cli/internal/recording"

	"github.com/spf13/cobra"
)

var interceptionListCmd = &cobra.Command{
	Use:   "list",
	Short: "List active interceptions",
	Long:  "List all active call interceptions using GET /interceptions.",
	RunE:  runInterceptionList,
}

func init() {
	interceptionCmd.AddCommand(interceptionListCmd)
}

func runInterceptionList(cmd *cobra.Command, args []string) error {
	client, err := api.NewClient()
	if err != nil {
		return err
	}

	result, err := recording.ListInterceptions(context.Background(), client)
	if err != nil {
		return err
	}

	return output.Print(result)
}
