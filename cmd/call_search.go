package cmd

import (
	"hepic-cli/internal/api"
	"hepic-cli/internal/call"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var callSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for SIP call data",
	Long: `Search for SIP call data with filters for time range, caller, callee, and call ID.

Examples:
  hepic call search --from 2025-01-01 --to 2025-01-31
  hepic call search --from 2025-01-01 --caller "+49123"
  hepic call search --from 2025-01-01 --callee "+49456" --format table
  hepic call search --from 2025-01-01 --call-id "abc123"`,
	RunE: runCallSearch,
}

func init() {
	callCmd.AddCommand(callSearchCmd)

	callSearchCmd.Flags().String("from", "", "Start time (RFC3339 or YYYY-MM-DD, required)")
	callSearchCmd.Flags().String("to", "", "End time (RFC3339 or YYYY-MM-DD, default: now)")
	callSearchCmd.Flags().String("caller", "", "Filter by caller (from_user)")
	callSearchCmd.Flags().String("callee", "", "Filter by callee (ruri_user)")
	callSearchCmd.Flags().String("call-id", "", "Filter by SIP Call-ID")

	callSearchCmd.MarkFlagRequired("from")
}

func runCallSearch(cmd *cobra.Command, args []string) error {
	from, _ := cmd.Flags().GetString("from")
	to, _ := cmd.Flags().GetString("to")
	caller, _ := cmd.Flags().GetString("caller")
	callee, _ := cmd.Flags().GetString("callee")
	callID, _ := cmd.Flags().GetString("call-id")

	client, err := api.NewClient()
	if err != nil {
		return err
	}

	params, err := call.NewSearchParams(from, to, caller, callee, callID)
	if err != nil {
		return err
	}

	result, err := call.SearchData(cmd.Context(), client, params)
	if err != nil {
		return err
	}

	return output.Print(result)
}
