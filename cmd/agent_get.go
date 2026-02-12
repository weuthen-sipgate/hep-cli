package cmd

import (

	"hepic-cli/internal/agent"
	"hepic-cli/internal/api"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var agentGetCmd = &cobra.Command{
	Use:   "get <uuid>",
	Short: "Get agent details by UUID",
	Long: `Get detailed information about a specific capture agent.

Examples:
  hepic agent get abc-123-def
  hepic agent get abc-123-def --format table`,
	Args: cobra.ExactArgs(1),
	RunE: runAgentGet,
}

func init() {
	agentCmd.AddCommand(agentGetCmd)
}

func runAgentGet(cmd *cobra.Command, args []string) error {
	uuid := args[0]

	client, err := api.NewClient()
	if err != nil {
		return err
	}

	result, err := agent.Get(cmd.Context(), client, uuid)
	if err != nil {
		return err
	}

	return output.Print(result)
}
