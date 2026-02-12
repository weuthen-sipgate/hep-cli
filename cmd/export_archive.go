package cmd

import (
	"context"
	"fmt"
	"os"

	"hepic-cli/internal/api"
	"hepic-cli/internal/export"

	"github.com/spf13/cobra"
)

var exportArchiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Export transaction archive",
	Long: `Export a transaction archive for the given call ID.

Requires -o/--output since archive data is binary.

Examples:
  hepic export archive --call-id abc123 -o archive.tar.gz
  hepic export archive --call-id abc123 --from 2025-01-01 -o archive.tar.gz`,
	RunE: runExportArchive,
}

func init() {
	exportCmd.AddCommand(exportArchiveCmd)

	exportArchiveCmd.Flags().String("call-id", "", "Call ID to export (required)")
	exportArchiveCmd.Flags().String("from", "", "Start time (RFC3339, date, or unix ms)")
	exportArchiveCmd.Flags().String("to", "", "End time (RFC3339, date, or unix ms)")
	exportArchiveCmd.Flags().StringP("output", "o", "", "Output file path (required for binary archive)")
	exportArchiveCmd.MarkFlagRequired("call-id")
}

func runExportArchive(cmd *cobra.Command, args []string) error {
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

	body, err := export.ExportTransactionArchive(context.Background(), client, params)
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
