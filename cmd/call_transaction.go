package cmd

import (
	"hepic-cli/internal/api"
	"hepic-cli/internal/call"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var callTransactionCmd = &cobra.Command{
	Use:   "transaction",
	Short: "Get SIP call transaction details",
	Long: `Retrieve detailed transaction information for a SIP call by Call-ID.

Examples:
  hepic call transaction --call-id "abc123" --from 2025-01-01
  hepic call transaction --call-id "abc123" --from 2025-01-01 --to 2025-01-31`,
	RunE: runCallTransaction,
}

func init() {
	callCmd.AddCommand(callTransactionCmd)

	callTransactionCmd.Flags().String("call-id", "", "SIP Call-ID (required)")
	callTransactionCmd.Flags().String("from", "", "Start time (RFC3339 or YYYY-MM-DD, required)")
	callTransactionCmd.Flags().String("to", "", "End time (RFC3339 or YYYY-MM-DD, default: now)")

	callTransactionCmd.MarkFlagRequired("call-id")
	callTransactionCmd.MarkFlagRequired("from")
}

func runCallTransaction(cmd *cobra.Command, args []string) error {
	callID, _ := cmd.Flags().GetString("call-id")
	from, _ := cmd.Flags().GetString("from")
	to, _ := cmd.Flags().GetString("to")

	client, err := api.NewClient()
	if err != nil {
		return err
	}

	params, err := call.NewSearchParams(from, to, "", "", callID)
	if err != nil {
		return err
	}

	result, err := call.GetTransaction(cmd.Context(), client, params)
	if err != nil {
		return err
	}

	return output.Print(result)
}
