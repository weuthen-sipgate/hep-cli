package cmd

import (
	"hepic-cli/internal/api"
	"hepic-cli/internal/call"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var callDecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Decode SIP call messages",
	Long: `Decode SIP call messages into a structured format.

Examples:
  hepic call decode --from 2025-01-01 --to 2025-01-31
  hepic call decode --from 2025-01-01 --call-id "abc123"`,
	RunE: runCallDecode,
}

func init() {
	callCmd.AddCommand(callDecodeCmd)

	callDecodeCmd.Flags().String("from", "", "Start time (RFC3339 or YYYY-MM-DD, required)")
	callDecodeCmd.Flags().String("to", "", "End time (RFC3339 or YYYY-MM-DD, default: now)")
	callDecodeCmd.Flags().String("caller", "", "Filter by caller (from_user)")
	callDecodeCmd.Flags().String("callee", "", "Filter by callee (ruri_user)")
	callDecodeCmd.Flags().String("call-id", "", "Filter by SIP Call-ID")

	callDecodeCmd.MarkFlagRequired("from")
}

func runCallDecode(cmd *cobra.Command, args []string) error {
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

	result, err := call.DecodeMessage(cmd.Context(), client, params)
	if err != nil {
		return err
	}

	return output.Print(result)
}
