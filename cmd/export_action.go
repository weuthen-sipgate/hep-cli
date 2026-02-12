package cmd

import (
	"fmt"

	"hepic-cli/internal/api"
	"hepic-cli/internal/export"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var validActionTypes = []string{"active", "hepicapp", "logs", "picserver", "rtpagent"}

var exportActionCmd = &cobra.Command{
	Use:   "action <type>",
	Short: "Export action data by type",
	Long: `Export action data for the specified type.

Valid types: active, hepicapp, logs, picserver, rtpagent

Examples:
  hepic export action logs
  hepic export action active
  hepic export action rtpagent`,
	Args: cobra.ExactArgs(1),
	RunE: runExportAction,
	ValidArgs: validActionTypes,
}

func init() {
	exportCmd.AddCommand(exportActionCmd)
}

func runExportAction(cmd *cobra.Command, args []string) error {
	actionType := args[0]

	valid := false
	for _, t := range validActionTypes {
		if t == actionType {
			valid = true
			break
		}
	}
	if !valid {
		return fmt.Errorf("invalid action type %q; valid types: active, hepicapp, logs, picserver, rtpagent", actionType)
	}

	client, err := api.NewClient()
	if err != nil {
		return err
	}

	result, err := export.ExportAction(cmd.Context(), client, actionType)
	if err != nil {
		return err
	}

	return output.Print(result)
}
