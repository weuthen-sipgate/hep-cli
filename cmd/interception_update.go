package cmd

import (
	"fmt"

	"hepic-cli/internal/api"
	"hepic-cli/internal/models"
	"hepic-cli/internal/output"
	"hepic-cli/internal/recording"

	"github.com/spf13/cobra"
)

var interceptionUpdateCmd = &cobra.Command{
	Use:   "update <uuid>",
	Short: "Update an existing interception",
	Long: `Update an existing call interception by UUID using PUT /interceptions/{uuid}.

Examples:
  hepic interception update abc-123 --caller "+4912345"
  hepic interception update abc-123 --callee "+4967890" --description "Updated"`,
	Args: cobra.ExactArgs(1),
	RunE: runInterceptionUpdate,
}

func init() {
	interceptionCmd.AddCommand(interceptionUpdateCmd)
	interceptionUpdateCmd.Flags().String("caller", "", "Caller number or pattern to intercept")
	interceptionUpdateCmd.Flags().String("callee", "", "Callee number or pattern to intercept")
	interceptionUpdateCmd.Flags().String("description", "", "Description of the interception")
	interceptionUpdateCmd.Flags().String("ip", "", "IP address to filter")
	interceptionUpdateCmd.Flags().Bool("status", true, "Enable or disable the interception")
	interceptionUpdateCmd.Flags().String("start-date", "", "Start date for the interception")
	interceptionUpdateCmd.Flags().String("stop-date", "", "Stop date for the interception")
}

func runInterceptionUpdate(cmd *cobra.Command, args []string) error {
	client, err := api.NewClient()
	if err != nil {
		return err
	}

	uuid := args[0]
	if uuid == "" {
		return fmt.Errorf("uuid is required")
	}

	data := models.InterceptionsStruct{}

	if v, _ := cmd.Flags().GetString("caller"); v != "" {
		data.SearchCaller = v
	}
	if v, _ := cmd.Flags().GetString("callee"); v != "" {
		data.SearchCallee = v
	}
	if v, _ := cmd.Flags().GetString("description"); v != "" {
		data.Description = v
	}
	if v, _ := cmd.Flags().GetString("ip"); v != "" {
		data.SearchIP = v
	}
	if cmd.Flags().Changed("status") {
		status, _ := cmd.Flags().GetBool("status")
		data.Status = status
	}
	if v, _ := cmd.Flags().GetString("start-date"); v != "" {
		data.StartDate = v
	}
	if v, _ := cmd.Flags().GetString("stop-date"); v != "" {
		data.StopDate = v
	}

	result, err := recording.UpdateInterception(cmd.Context(), client, uuid, data)
	if err != nil {
		return err
	}

	return output.Print(result)
}
