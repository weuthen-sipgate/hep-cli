package cmd

import (
	"context"
	"fmt"
	"io"
	"os"

	"hepic-cli/internal/api"
	"hepic-cli/internal/export"

	"github.com/spf13/cobra"
)

var exportTextCmd = &cobra.Command{
	Use:   "text",
	Short: "Export messages as plain text",
	Long: `Export call messages as human-readable plain text to stdout.

Examples:
  hepic export text --call-id abc123
  hepic export text --call-id abc123 --from 2025-01-01 --to 2025-01-02`,
	RunE: runExportText,
}

func init() {
	exportCmd.AddCommand(exportTextCmd)

	exportTextCmd.Flags().String("call-id", "", "Call ID to export (required)")
	exportTextCmd.Flags().String("from", "", "Start time (RFC3339, date, or unix ms)")
	exportTextCmd.Flags().String("to", "", "End time (RFC3339, date, or unix ms)")
	exportTextCmd.MarkFlagRequired("call-id")
}

func runExportText(cmd *cobra.Command, args []string) error {
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

	body, err := export.ExportText(context.Background(), client, params)
	if err != nil {
		return err
	}
	defer body.Close()

	n, err := io.Copy(os.Stdout, body)
	if err != nil {
		return fmt.Errorf("failed to write text output: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Exported %d bytes\n", n)
	return nil
}
