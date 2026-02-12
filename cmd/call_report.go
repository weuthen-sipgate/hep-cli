package cmd

import (
	"hepic-cli/internal/api"
	"hepic-cli/internal/call"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var callReportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate call reports (DTMF, log, QoS)",
	Long:  "Commands for generating call reports including DTMF tones, call logs, and Quality of Service metrics.",
}

var callReportDTMFCmd = &cobra.Command{
	Use:   "dtmf",
	Short: "Get DTMF report for a call",
	Long: `Retrieve DTMF (Dual-Tone Multi-Frequency) report for a SIP call.

Examples:
  hepic call report dtmf --call-id "abc123" --from 2025-01-01
  hepic call report dtmf --call-id "abc123" --from 2025-01-01 --to 2025-01-31`,
	RunE: runCallReportDTMF,
}

var callReportLogCmd = &cobra.Command{
	Use:   "log",
	Short: "Get log report for a call",
	Long: `Retrieve log report for a SIP call.

Examples:
  hepic call report log --call-id "abc123" --from 2025-01-01
  hepic call report log --call-id "abc123" --from 2025-01-01 --to 2025-01-31`,
	RunE: runCallReportLog,
}

var callReportQOSCmd = &cobra.Command{
	Use:   "qos",
	Short: "Get QoS report for a call",
	Long: `Retrieve Quality of Service (QoS) report for a SIP call.

Examples:
  hepic call report qos --call-id "abc123" --from 2025-01-01
  hepic call report qos --call-id "abc123" --from 2025-01-01 --to 2025-01-31`,
	RunE: runCallReportQOS,
}

func init() {
	callCmd.AddCommand(callReportCmd)

	// Add subcommands to report
	callReportCmd.AddCommand(callReportDTMFCmd)
	callReportCmd.AddCommand(callReportLogCmd)
	callReportCmd.AddCommand(callReportQOSCmd)

	// DTMF flags
	callReportDTMFCmd.Flags().String("call-id", "", "SIP Call-ID (required)")
	callReportDTMFCmd.Flags().String("from", "", "Start time (RFC3339 or YYYY-MM-DD, required)")
	callReportDTMFCmd.Flags().String("to", "", "End time (RFC3339 or YYYY-MM-DD, default: now)")
	callReportDTMFCmd.MarkFlagRequired("call-id")
	callReportDTMFCmd.MarkFlagRequired("from")

	// Log flags
	callReportLogCmd.Flags().String("call-id", "", "SIP Call-ID (required)")
	callReportLogCmd.Flags().String("from", "", "Start time (RFC3339 or YYYY-MM-DD, required)")
	callReportLogCmd.Flags().String("to", "", "End time (RFC3339 or YYYY-MM-DD, default: now)")
	callReportLogCmd.MarkFlagRequired("call-id")
	callReportLogCmd.MarkFlagRequired("from")

	// QoS flags
	callReportQOSCmd.Flags().String("call-id", "", "SIP Call-ID (required)")
	callReportQOSCmd.Flags().String("from", "", "Start time (RFC3339 or YYYY-MM-DD, required)")
	callReportQOSCmd.Flags().String("to", "", "End time (RFC3339 or YYYY-MM-DD, default: now)")
	callReportQOSCmd.MarkFlagRequired("call-id")
	callReportQOSCmd.MarkFlagRequired("from")
}

func runCallReportDTMF(cmd *cobra.Command, args []string) error {
	callID, _ := cmd.Flags().GetString("call-id")
	from, _ := cmd.Flags().GetString("from")
	to, _ := cmd.Flags().GetString("to")

	client, err := api.NewClient()
	if err != nil {
		output.PrintError(err)
		return err
	}

	params, err := call.NewSearchParams(from, to, "", "", callID)
	if err != nil {
		output.PrintError(err)
		return err
	}

	result, err := call.ReportDTMF(cmd.Context(), client, params)
	if err != nil {
		output.PrintError(err)
		return err
	}

	return output.Print(result)
}

func runCallReportLog(cmd *cobra.Command, args []string) error {
	callID, _ := cmd.Flags().GetString("call-id")
	from, _ := cmd.Flags().GetString("from")
	to, _ := cmd.Flags().GetString("to")

	client, err := api.NewClient()
	if err != nil {
		output.PrintError(err)
		return err
	}

	params, err := call.NewSearchParams(from, to, "", "", callID)
	if err != nil {
		output.PrintError(err)
		return err
	}

	result, err := call.ReportLog(cmd.Context(), client, params)
	if err != nil {
		output.PrintError(err)
		return err
	}

	return output.Print(result)
}

func runCallReportQOS(cmd *cobra.Command, args []string) error {
	callID, _ := cmd.Flags().GetString("call-id")
	from, _ := cmd.Flags().GetString("from")
	to, _ := cmd.Flags().GetString("to")

	client, err := api.NewClient()
	if err != nil {
		output.PrintError(err)
		return err
	}

	params, err := call.NewSearchParams(from, to, "", "", callID)
	if err != nil {
		output.PrintError(err)
		return err
	}

	result, err := call.ReportQOS(cmd.Context(), client, params)
	if err != nil {
		output.PrintError(err)
		return err
	}

	return output.Print(result)
}
