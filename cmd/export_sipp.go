package cmd

import (
	"fmt"
	"os"

	"hepic-cli/internal/api"
	"hepic-cli/internal/export"

	"github.com/spf13/cobra"
)

var exportSippCmd = &cobra.Command{
	Use:   "sipp",
	Short: "Export messages as SIPp format",
	Long: `Export call messages as SIPp scenario file.

Requires --call-id and -o/--output since SIPp output is binary/XML data.

Examples:
  hepic export sipp --call-id abc123 -o scenario.xml
  hepic export sipp --call-id abc123 --from 2025-01-01 -o scenario.xml`,
	RunE: runExportSipp,
}

func init() {
	exportCmd.AddCommand(exportSippCmd)

	exportSippCmd.Flags().String("call-id", "", "Call ID to export (required)")
	exportSippCmd.Flags().String("from", "", "Start time (RFC3339, date, or unix ms)")
	exportSippCmd.Flags().String("to", "", "End time (RFC3339, date, or unix ms)")
	exportSippCmd.Flags().StringP("output", "o", "", "Output file path (required for binary SIPp)")
	exportSippCmd.MarkFlagRequired("call-id")
}

func runExportSipp(cmd *cobra.Command, args []string) error {
	outputPath, _ := cmd.Flags().GetString("output")
	if outputPath == "" {
		return fmt.Errorf("binary output requires -o flag; use -o to write to file")
	}

	callID, _ := cmd.Flags().GetString("call-id")
	from, _ := cmd.Flags().GetString("from")
	to, _ := cmd.Flags().GetString("to")

	params, err := export.NewExportParams(from, to, callID)
	if err != nil {
		return err
	}

	client, err := api.NewClient()
	if err != nil {
		return err
	}

	body, err := export.ExportSIPP(cmd.Context(), client, params)
	if err != nil {
		return err
	}

	n, err := writeToFile(body, outputPath)
	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Exported %d bytes to %s\n", n, outputPath)
	return nil
}
