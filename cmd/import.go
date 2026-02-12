package cmd

import (
	"context"
	"fmt"
	"os"

	"hepic-cli/internal/admin"
	"hepic-cli/internal/api"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import data into the HEPIC platform",
	Long:  "Import data files (e.g., PCAP captures) into the HEPIC platform.",
}

var importPcapCmd = &cobra.Command{
	Use:   "pcap",
	Short: "Import a PCAP file",
	Long: `Import a PCAP capture file into the HEPIC platform.

Use --now to import immediately instead of queuing.

Examples:
  hepic import pcap --file capture.pcap
  hepic import pcap --file capture.pcap --now`,
	RunE: func(cmd *cobra.Command, args []string) error {
		filePath, _ := cmd.Flags().GetString("file")
		now, _ := cmd.Flags().GetBool("now")

		// Verify file exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return fmt.Errorf("file not found: %s", filePath)
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		var result interface{}
		if now {
			result, err = admin.ImportPCAPNow(context.Background(), client, filePath)
		} else {
			result, err = admin.ImportPCAP(context.Background(), client, filePath)
		}
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(importCmd)

	importCmd.AddCommand(importPcapCmd)

	importPcapCmd.Flags().String("file", "", "Path to the PCAP file to import (required)")
	importPcapCmd.Flags().Bool("now", false, "Import immediately instead of queuing")
	importPcapCmd.MarkFlagRequired("file")
}
