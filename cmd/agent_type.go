package cmd

import (
	"context"

	"hepic-cli/internal/agent"
	"hepic-cli/internal/api"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var agentTypeCmd = &cobra.Command{
	Use:   "type <type>",
	Short: "List agents by type",
	Long: `List all capture agents filtered by their type.

Examples:
  hepic agent type home
  hepic agent type remote --format table`,
	Args: cobra.ExactArgs(1),
	RunE: runAgentType,
}

func init() {
	agentCmd.AddCommand(agentTypeCmd)
}

func runAgentType(cmd *cobra.Command, args []string) error {
	agentType := args[0]

	client, err := api.NewClient()
	if err != nil {
		return err
	}

	result, err := agent.ListByType(context.Background(), client, agentType)
	if err != nil {
		return err
	}

	return output.Print(result)
}
