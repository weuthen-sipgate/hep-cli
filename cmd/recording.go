package cmd

import "github.com/spf13/cobra"

var recordingCmd = &cobra.Command{
	Use:     "recording",
	Short:   "Manage call recordings",
	Long:    "Search, download, and inspect call recordings from the HEPIC platform.",
	GroupID: "call",
}

func init() {
	rootCmd.AddCommand(recordingCmd)
}
