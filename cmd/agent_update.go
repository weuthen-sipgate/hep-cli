package cmd

import (
	"encoding/json"
	"fmt"

	"hepic-cli/internal/agent"
	"hepic-cli/internal/api"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var agentUpdateCmd = &cobra.Command{
	Use:   "update <uuid>",
	Short: "Update an agent by UUID",
	Long: `Update an existing capture agent with new configuration.

The --data flag accepts a JSON string with the fields to update.

Examples:
  hepic agent update abc-123-def --data '{"host":"10.0.0.2","port":9061}'
  hepic agent update abc-123-def --data '{"active":false}'`,
	Args: cobra.ExactArgs(1),
	RunE: runAgentUpdate,
}

func init() {
	agentCmd.AddCommand(agentUpdateCmd)

	agentUpdateCmd.Flags().String("data", "", "JSON data for the update (required)")
	agentUpdateCmd.MarkFlagRequired("data")
}

func runAgentUpdate(cmd *cobra.Command, args []string) error {
	uuid := args[0]
	dataStr, _ := cmd.Flags().GetString("data")

	var data interface{}
	if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
		return fmt.Errorf("invalid JSON in --data flag: %w", err)
	}

	client, err := api.NewClient()
	if err != nil {
		return err
	}

	result, err := agent.Update(cmd.Context(), client, uuid, data)
	if err != nil {
		return err
	}

	return output.Print(result)
}
