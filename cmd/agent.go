package cmd

import (
	"github.com/spf13/cobra"
)

var agentCmd = &cobra.Command{
	Use:     "agent",
	Short:   "Manage capture agents",
	GroupID: "admin",
	Long: `Manage capture agents and their subscriptions.

Available subcommands:
  list      List all registered agents
  get       Get agent details by UUID
  update    Update an agent by UUID
  delete    Delete an agent by UUID
  search    Search agents by GUID and type
  type      List agents by type`,
}

func init() {
	rootCmd.AddCommand(agentCmd)
}
