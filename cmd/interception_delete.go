package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"hepic-cli/internal/api"
	"hepic-cli/internal/output"
	"hepic-cli/internal/recording"

	"github.com/spf13/cobra"
)

var interceptionDeleteCmd = &cobra.Command{
	Use:   "delete <uuid>",
	Short: "Delete an interception",
	Long: `Delete a call interception by UUID using DELETE /interceptions/{uuid}.

Without --force, prompts for confirmation before deleting.

Examples:
  hepic interception delete abc-123
  hepic interception delete abc-123 --force`,
	Args: cobra.ExactArgs(1),
	RunE: runInterceptionDelete,
}

func init() {
	interceptionCmd.AddCommand(interceptionDeleteCmd)
	interceptionDeleteCmd.Flags().Bool("force", false, "Skip deletion confirmation prompt")
}

func runInterceptionDelete(cmd *cobra.Command, args []string) error {
	client, err := api.NewClient()
	if err != nil {
		return err
	}

	uuid := args[0]
	force, _ := cmd.Flags().GetBool("force")

	if !force {
		fmt.Fprintf(os.Stderr, "Are you sure you want to delete interception %s? [y/N] ", uuid)
		reader := bufio.NewReader(os.Stdin)
		answer, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read confirmation: %w", err)
		}
		answer = strings.TrimSpace(answer)
		if answer != "y" && answer != "Y" {
			fmt.Fprintln(os.Stderr, "Deletion cancelled.")
			return nil
		}
	}

	result, err := recording.DeleteInterception(cmd.Context(), client, uuid)
	if err != nil {
		return err
	}

	return output.Print(result)
}
