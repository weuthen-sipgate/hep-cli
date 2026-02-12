package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"hepic-cli/internal/api"
	"hepic-cli/internal/output"
	"hepic-cli/internal/user"

	"github.com/spf13/cobra"
)

var tokenDeleteCmd = &cobra.Command{
	Use:   "delete <uuid>",
	Short: "Delete an API token",
	Long:  "Delete an API authentication token by UUID. Requires confirmation unless --force is specified.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		uuid := args[0]
		force, _ := cmd.Flags().GetBool("force")

		if !force {
			fmt.Fprintf(os.Stderr, "Delete token %s? [y/N]: ", uuid)
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			if err != nil {
				return fmt.Errorf("failed to read input: %w", err)
			}
			if strings.TrimSpace(strings.ToLower(input)) != "y" {
				return fmt.Errorf("operation cancelled")
			}
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := user.DeleteToken(context.Background(), client, uuid)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

func init() {
	tokenCmd.AddCommand(tokenDeleteCmd)

	tokenDeleteCmd.Flags().Bool("force", false, "Skip confirmation prompt")
}
