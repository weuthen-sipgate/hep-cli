package cmd

import (
	"context"
	"fmt"
	"os"

	"hepic-cli/internal/api"
	"hepic-cli/internal/export"

	"github.com/spf13/cobra"
)

var exportReportCmd = &cobra.Command{
	Use:   "report",
	Short: "Export transaction report",
	Long: `Export a transaction report for the given call ID.

Requires -o/--output since report data is binary.

Examples:
  hepic export report --call-id abc123 -o report.pdf
  hepic export report --call-id abc123 --from 2025-01-01 -o report.pdf`,
	RunE: runExportReport,
}

func init() {
	exportCmd.AddCommand(exportReportCmd)

	exportReportCmd.Flags().String("call-id", "", "Call ID to export (required)")
	exportReportCmd.Flags().String("from", "", "Start time (RFC3339, date, or unix ms)")
	exportReportCmd.Flags().String("to", "", "End time (RFC3339, date, or unix ms)")
	exportReportCmd.Flags().StringP("output", "o", "", "Output file path (required for binary report)")
	exportReportCmd.MarkFlagRequired("call-id")
}

func runExportReport(cmd *cobra.Command, args []string) error {
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

	body, err := export.ExportTransactionReport(context.Background(), client, params)
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
