package cmd

import (
	"context"

	"hepic-cli/internal/agent"
	"hepic-cli/internal/api"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var agentSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search agents by GUID and type",
	Long: `Search for capture agents using a GUID and agent type.

Both --guid and --type are required.

Examples:
  hepic agent search --guid abc-123 --type home
  hepic agent search --guid abc-123 --type remote --format table`,
	RunE: runAgentSearch,
}

func init() {
	agentCmd.AddCommand(agentSearchCmd)

	agentSearchCmd.Flags().String("guid", "", "Agent GUID to search for (required)")
	agentSearchCmd.Flags().String("type", "", "Agent type to search for (required)")
	agentSearchCmd.MarkFlagRequired("guid")
	agentSearchCmd.MarkFlagRequired("type")
}

func runAgentSearch(cmd *cobra.Command, args []string) error {
	guid, _ := cmd.Flags().GetString("guid")
	agentType, _ := cmd.Flags().GetString("type")

	client, err := api.NewClient()
	if err != nil {
		return err
	}

	result, err := agent.Search(context.Background(), client, guid, agentType)
	if err != nil {
		return err
	}

	return output.Print(result)
}
