package cmd

import (
	"hepic-cli/internal/api"
	"hepic-cli/internal/call"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var callMessageCmd = &cobra.Command{
	Use:   "message",
	Short: "Search for SIP call messages",
	Long: `Search for raw SIP call messages with filters for time range, caller, callee, and call ID.

Examples:
  hepic call message --from 2025-01-01 --to 2025-01-31
  hepic call message --from 2025-01-01 --call-id "abc123"`,
	RunE: runCallMessage,
}

func init() {
	callCmd.AddCommand(callMessageCmd)

	callMessageCmd.Flags().String("from", "", "Start time (RFC3339 or YYYY-MM-DD, required)")
	callMessageCmd.Flags().String("to", "", "End time (RFC3339 or YYYY-MM-DD, default: now)")
	callMessageCmd.Flags().String("caller", "", "Filter by caller (from_user)")
	callMessageCmd.Flags().String("callee", "", "Filter by callee (ruri_user)")
	callMessageCmd.Flags().String("call-id", "", "Filter by SIP Call-ID")

	callMessageCmd.MarkFlagRequired("from")
}

func runCallMessage(cmd *cobra.Command, args []string) error {
	from, _ := cmd.Flags().GetString("from")
	to, _ := cmd.Flags().GetString("to")
	caller, _ := cmd.Flags().GetString("caller")
	callee, _ := cmd.Flags().GetString("callee")
	callID, _ := cmd.Flags().GetString("call-id")

	client, err := api.NewClient()
	if err != nil {
		output.PrintError(err)
		return err
	}

	params, err := call.NewSearchParams(from, to, caller, callee, callID)
	if err != nil {
		output.PrintError(err)
		return err
	}

	result, err := call.SearchMessage(cmd.Context(), client, params)
	if err != nil {
		output.PrintError(err)
		return err
	}

	return output.Print(result)
}
