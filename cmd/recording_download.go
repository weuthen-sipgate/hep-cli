package cmd

import (
	"context"
	"fmt"
	"io"
	"os"

	"hepic-cli/internal/api"
	"hepic-cli/internal/recording"

	"github.com/spf13/cobra"
)

var recordingDownloadCmd = &cobra.Command{
	Use:   "download <uuid>",
	Short: "Download a recording file",
	Long: `Download a recording file (audio or PCAP) by UUID.

Examples:
  hepic recording download abc-123 -o call.wav
  hepic recording download abc-123 -o capture.pcap --type pcap`,
	Args: cobra.ExactArgs(1),
	RunE: runRecordingDownload,
}

func init() {
	recordingCmd.AddCommand(recordingDownloadCmd)
	recordingDownloadCmd.Flags().StringP("output", "o", "", "Output file path (required)")
	recordingDownloadCmd.Flags().String("type", "audio", "Download type: audio, pcap")
	recordingDownloadCmd.MarkFlagRequired("output")
}

func runRecordingDownload(cmd *cobra.Command, args []string) error {
	client, err := api.NewClient()
	if err != nil {
		return err
	}

	uuid := args[0]
	outputPath, _ := cmd.Flags().GetString("output")
	dlType, _ := cmd.Flags().GetString("type")

	if dlType != "audio" && dlType != "pcap" {
		return fmt.Errorf("invalid download type %q: must be 'audio' or 'pcap'", dlType)
	}

	body, err := recording.Download(context.Background(), client, dlType, uuid)
	if err != nil {
		return err
	}
	defer body.Close()

	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	n, err := io.Copy(file, body)
	if err != nil {
		return fmt.Errorf("failed to write recording data: %w", err)
	}

	fmt.Fprintf(os.Stderr, "Downloaded %d bytes to %s\n", n, outputPath)
	return nil
}
