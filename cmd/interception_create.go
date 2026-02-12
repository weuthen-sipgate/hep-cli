package cmd

import (

	"hepic-cli/internal/api"
	"hepic-cli/internal/models"
	"hepic-cli/internal/output"
	"hepic-cli/internal/recording"

	"github.com/spf13/cobra"
)

var interceptionCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new interception",
	Long: `Create a new call interception using POST /interceptions.

Examples:
  hepic interception create --caller "+4912345" --callee "+4967890"
  hepic interception create --caller "+4912345" --callee "+4967890" --description "Test intercept"`,
	RunE: runInterceptionCreate,
}

func init() {
	interceptionCmd.AddCommand(interceptionCreateCmd)
	interceptionCreateCmd.Flags().String("caller", "", "Caller number or pattern to intercept")
	interceptionCreateCmd.Flags().String("callee", "", "Callee number or pattern to intercept")
	interceptionCreateCmd.Flags().String("description", "", "Description of the interception")
	interceptionCreateCmd.Flags().String("ip", "", "IP address to filter")
	interceptionCreateCmd.Flags().Bool("status", true, "Enable or disable the interception")
	interceptionCreateCmd.Flags().String("start-date", "", "Start date for the interception")
	interceptionCreateCmd.Flags().String("stop-date", "", "Stop date for the interception")
}

func runInterceptionCreate(cmd *cobra.Command, args []string) error {
	client, err := api.NewClient()
	if err != nil {
		return err
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

	result, err := recording.CreateInterception(cmd.Context(), client, data)
	if err != nil {
		return err
	}

	return output.Print(result)
}
