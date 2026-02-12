package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:     "export",
	Short:   "Export call data in various formats",
	GroupID: "call",
	Long: `Export call data as PCAP, SIPp, text, reports, or archives.

Available subcommands:
  pcap      Export call data as PCAP file
  sipp      Export messages as SIPp format
  text      Export messages as plain text
  report    Export transaction report
  archive   Export transaction archive
  action    Export action data by type`,
}

func init() {
	rootCmd.AddCommand(exportCmd)
}

// writeToFile writes the contents of r to the file at path and returns the
// number of bytes written. The ReadCloser is closed when done.
func writeToFile(r io.ReadCloser, path string) (int64, error) {
	defer r.Close()
	f, err := os.Create(path)
	if err != nil {
		return 0, fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()
	n, err := io.Copy(f, r)
	if err != nil {
		return n, fmt.Errorf("failed to write output file: %w", err)
	}
	return n, nil
}
