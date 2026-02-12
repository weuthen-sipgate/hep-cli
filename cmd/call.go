package cmd

import (
	"github.com/spf13/cobra"
)

var callCmd = &cobra.Command{
	Use:   "call",
	Short: "Search and analyze SIP calls",
	Long:  "Commands for searching SIP call data, viewing transactions, and generating call reports (DTMF, log, QoS).",
}

func init() {
	rootCmd.AddCommand(callCmd)
}
