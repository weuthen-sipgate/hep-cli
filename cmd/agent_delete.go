package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"hepic-cli/internal/agent"
	"hepic-cli/internal/api"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var agentDeleteCmd = &cobra.Command{
	Use:   "delete <uuid>",
	Short: "Delete an agent by UUID",
	Long: `Delete a capture agent from the system.

Requires confirmation unless --force is used.

Examples:
  hepic agent delete abc-123-def
  hepic agent delete abc-123-def --force`,
	Args: cobra.ExactArgs(1),
	RunE: runAgentDelete,
}

func init() {
	agentCmd.AddCommand(agentDeleteCmd)

	agentDeleteCmd.Flags().Bool("force", false, "Skip confirmation prompt")
}

func runAgentDelete(cmd *cobra.Command, args []string) error {
	uuid := args[0]
	force, _ := cmd.Flags().GetBool("force")

	if !force {
		fmt.Fprintf(os.Stderr, "Delete agent %s? [y/N]: ", uuid)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}
		input = strings.TrimSpace(strings.ToLower(input))
		if input != "y" && input != "yes" {
			return fmt.Errorf("aborted")
		}
	}

	client, err := api.NewClient()
	if err != nil {
		return err
	}

	result, err := agent.Delete(context.Background(), client, uuid)
	if err != nil {
		return err
	}

	return output.Print(result)
}
