package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"hepic-cli/internal/api"
	"hepic-cli/internal/output"
	"hepic-cli/internal/statistic"

	"github.com/spf13/cobra"
)

var prometheusCmd = &cobra.Command{
	Use:   "prometheus",
	Short: "Query Prometheus metrics",
	Long:  "Query Prometheus data, values, and labels via the HEPIC platform.",
}

var prometheusQueryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query Prometheus data",
	Long: `Query Prometheus time-series data.

Examples:
  hepic prometheus query --data '{"param":{"query":"up"},"timestamp":{"from":1700000000,"to":1700003600}}'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dataStr, _ := cmd.Flags().GetString("data")
		if dataStr == "" {
			return fmt.Errorf("--data flag is required")
		}

		var data interface{}
		if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
			return fmt.Errorf("invalid JSON in --data: %w", err)
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := statistic.QueryData(context.Background(), client, data)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

var prometheusValueCmd = &cobra.Command{
	Use:   "value",
	Short: "Query a Prometheus metric value",
	Long: `Query a specific Prometheus metric value.

Examples:
  hepic prometheus value --data '{"param":{"query":"up"}}'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dataStr, _ := cmd.Flags().GetString("data")
		if dataStr == "" {
			return fmt.Errorf("--data flag is required")
		}

		var data interface{}
		if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
			return fmt.Errorf("invalid JSON in --data: %w", err)
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := statistic.QueryValue(context.Background(), client, data)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

var prometheusLabelsCmd = &cobra.Command{
	Use:   "labels",
	Short: "List all Prometheus labels",
	Long:  "Retrieve all available Prometheus labels.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := statistic.Labels(context.Background(), client)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

var prometheusLabelCmd = &cobra.Command{
	Use:   "label <name>",
	Short: "Get details for a Prometheus label",
	Long: `Retrieve details for a specific Prometheus label.

Examples:
  hepic prometheus label instance`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		label := args[0]

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := statistic.LabelDetail(context.Background(), client, label)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(prometheusCmd)

	prometheusCmd.AddCommand(prometheusQueryCmd)
	prometheusQueryCmd.Flags().String("data", "", "Query data as JSON string (required)")

	prometheusCmd.AddCommand(prometheusValueCmd)
	prometheusValueCmd.Flags().String("data", "", "Query data as JSON string (required)")

	prometheusCmd.AddCommand(prometheusLabelsCmd)

	prometheusCmd.AddCommand(prometheusLabelCmd)
}
