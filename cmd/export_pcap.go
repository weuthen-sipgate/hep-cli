package cmd

import (
	"fmt"
	"os"

	"hepic-cli/internal/api"
	"hepic-cli/internal/export"

	"github.com/spf13/cobra"
)

var exportPcapCmd = &cobra.Command{
	Use:   "pcap",
	Short: "Export call data as PCAP file",
	Long: `Export call data as a PCAP capture file.

Requires --call-id and -o/--output since PCAP is binary data.

Examples:
  hepic export pcap --call-id abc123 -o capture.pcap
  hepic export pcap --call-id abc123 --from 2025-01-01 --to 2025-01-02 -o capture.pcap`,
	RunE: runExportPcap,
}

func init() {
	exportCmd.AddCommand(exportPcapCmd)

	exportPcapCmd.Flags().String("call-id", "", "Call ID to export (required)")
	exportPcapCmd.Flags().String("from", "", "Start time (RFC3339, date, or unix ms)")
	exportPcapCmd.Flags().String("to", "", "End time (RFC3339, date, or unix ms)")
	exportPcapCmd.Flags().StringP("output", "o", "", "Output file path (required for binary PCAP)")
	exportPcapCmd.MarkFlagRequired("call-id")
}

func runExportPcap(cmd *cobra.Command, args []string) error {
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

	body, err := export.ExportPCAPData(cmd.Context(), client, params)
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
