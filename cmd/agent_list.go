package cmd

import (

	"hepic-cli/internal/agent"
	"hepic-cli/internal/api"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var agentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all registered agents",
	Long: `List all capture agents registered in the system.

Examples:
  hepic agent list
  hepic agent list --format table`,
	RunE: runAgentList,
}

func init() {
	agentCmd.AddCommand(agentListCmd)
}

func runAgentList(cmd *cobra.Command, args []string) error {
	client, err := api.NewClient()
	if err != nil {
		return err
	}

	result, err := agent.List(cmd.Context(), client)
	if err != nil {
		return err
	}

	return output.Print(result)
}
