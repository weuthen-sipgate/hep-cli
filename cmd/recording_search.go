package cmd

import (

	"hepic-cli/internal/api"
	"hepic-cli/internal/output"
	"hepic-cli/internal/recording"

	"github.com/spf13/cobra"
)

var recordingSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for call recordings",
	Long:  "Search for call recordings within a time range using POST /call/recording/data.",
	RunE:  runRecordingSearch,
}

func init() {
	recordingCmd.AddCommand(recordingSearchCmd)
	recordingSearchCmd.Flags().String("from", "", "Start time (RFC3339, YYYY-MM-DD, or unix ms)")
	recordingSearchCmd.Flags().String("to", "", "End time (RFC3339, YYYY-MM-DD, or unix ms)")
}

func runRecordingSearch(cmd *cobra.Command, args []string) error {
	client, err := api.NewClient()
	if err != nil {
		return err
	}

	from, _ := cmd.Flags().GetString("from")
	to, _ := cmd.Flags().GetString("to")

	params, err := recording.NewSearchParams(from, to)
	if err != nil {
		return err
	}

	result, err := recording.SearchData(cmd.Context(), client, params)
	if err != nil {
		return err
	}

	return output.Print(result)
}
