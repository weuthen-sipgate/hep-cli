package cmd

import (
	"context"
	"fmt"

	"hepic-cli/internal/api"
	"hepic-cli/internal/output"
	"hepic-cli/internal/recording"

	"github.com/spf13/cobra"
)

var recordingInfoCmd = &cobra.Command{
	Use:   "info <uuid>",
	Short: "Show recording metadata",
	Long:  "Retrieve metadata for a specific recording by UUID using GET /call/recording/info/{uuid}.",
	Args:  cobra.ExactArgs(1),
	RunE:  runRecordingInfo,
}

func init() {
	recordingCmd.AddCommand(recordingInfoCmd)
}

func runRecordingInfo(cmd *cobra.Command, args []string) error {
	client, err := api.NewClient()
	if err != nil {
		return err
	}

	uuid := args[0]
	if uuid == "" {
		return fmt.Errorf("uuid is required")
	}

	result, err := recording.Info(context.Background(), client, uuid)
	if err != nil {
		return err
	}

	return output.Print(result)
}
