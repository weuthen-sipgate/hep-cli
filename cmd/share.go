package cmd

import (

	"hepic-cli/internal/admin"
	"hepic-cli/internal/api"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var shareCmd = &cobra.Command{
	Use:     "share",
	Short:   "Share call data, reports, and exports",
	GroupID: "call",
	Long: `Share various types of call data via the HEPIC platform.

Available subcommands:
  report       Share a call report (dtmf or log)
  transaction  Share a call transaction
  pcap         Share a PCAP export
  text         Share a text export
  ipalias      Share an IP alias
  mapping      Share a protocol mapping`,
}

var shareReportCmd = &cobra.Command{
	Use:   "report <type> <uuid>",
	Short: "Share a call report",
	Long:  "Share a call report by type (dtmf or log) and UUID.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		reportType := args[0]
		uuid := args[1]

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := admin.ShareReport(cmd.Context(), client, reportType, uuid)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

var shareTransactionCmd = &cobra.Command{
	Use:   "transaction <uuid>",
	Short: "Share a call transaction",
	Long:  "Share a call transaction by UUID.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uuid := args[0]

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := admin.ShareTransaction(cmd.Context(), client, uuid)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

var sharePcapCmd = &cobra.Command{
	Use:   "pcap <uuid>",
	Short: "Share a PCAP export",
	Long:  "Share a PCAP export by UUID.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uuid := args[0]

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := admin.SharePCAP(cmd.Context(), client, uuid)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

var shareTextCmd = &cobra.Command{
	Use:   "text <uuid>",
	Short: "Share a text export",
	Long:  "Share a text export by UUID.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uuid := args[0]

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := admin.ShareText(cmd.Context(), client, uuid)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

var shareIPAliasCmd = &cobra.Command{
	Use:   "ipalias <uuid>",
	Short: "Share an IP alias",
	Long:  "Share an IP alias by UUID.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uuid := args[0]

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := admin.ShareIPAlias(cmd.Context(), client, uuid)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

var shareMappingCmd = &cobra.Command{
	Use:   "mapping <uuid>",
	Short: "Share a protocol mapping",
	Long:  "Share a protocol mapping by UUID.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uuid := args[0]

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := admin.ShareMapping(cmd.Context(), client, uuid)
		if err != nil {
			return err
		}

		return output.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(shareCmd)

	shareCmd.AddCommand(shareReportCmd)
	shareCmd.AddCommand(shareTransactionCmd)
	shareCmd.AddCommand(sharePcapCmd)
	shareCmd.AddCommand(shareTextCmd)
	shareCmd.AddCommand(shareIPAliasCmd)
	shareCmd.AddCommand(shareMappingCmd)
}
