package cmd

import (
	"context"

	"hepic-cli/internal/admin"
	"hepic-cli/internal/api"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var troubleshootingCmd = &cobra.Command{
	Use:   "troubleshooting",
	Short: "Troubleshooting and diagnostics",
	Long:  "Access troubleshooting logs and diagnostic information from the HEPIC platform.",
}

var troubleshootingLogCmd = &cobra.Command{
	Use:   "log <type> <action>",
	Short: "View troubleshooting logs",
	Long: `Retrieve troubleshooting logs by type and action.

Both type and action are provided as positional arguments.

Examples:
  hepic troubleshooting log system status
  hepic troubleshooting log hep capture`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		logType := args[0]
		action := args[1]

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := admin.TroubleshootingLog(context.Background(), client, logType, action)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(troubleshootingCmd)

	troubleshootingCmd.AddCommand(troubleshootingLogCmd)
}
