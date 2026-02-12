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

var statisticCmd = &cobra.Command{
	Use:   "statistic",
	Short: "Query statistics and metrics",
	Long:  "Query database statistics, measurements, metrics, retentions, and time-series data.",
}

var statisticDBCmd = &cobra.Command{
	Use:   "db",
	Short: "Show database statistics",
	Long:  "Retrieve and display database statistics from the HEPIC platform.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := statistic.DBStats(context.Background(), client)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

var statisticDataCmd = &cobra.Command{
	Use:   "data",
	Short: "Query statistical data",
	Long: `Query statistical time-series data with optional time range.

Examples:
  hepic statistic data --data '{"param":{"search":{"query":"sip"}}}'
  hepic statistic data --from 2025-01-01 --to 2025-01-02 --data '{"param":{}}'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dataStr, _ := cmd.Flags().GetString("data")
		from, _ := cmd.Flags().GetString("from")
		to, _ := cmd.Flags().GetString("to")

		var data map[string]interface{}
		if dataStr != "" {
			if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
				return fmt.Errorf("invalid JSON in --data: %w", err)
			}
		} else {
			data = make(map[string]interface{})
		}

		// Inject timestamp if from/to provided
		if from != "" || to != "" {
			ts := make(map[string]interface{})
			if from != "" {
				ts["from"] = from
			}
			if to != "" {
				ts["to"] = to
			}
			data["timestamp"] = ts
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := statistic.Data(context.Background(), client, data)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

var statisticMetricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Query available metrics",
	Long: `Query available metrics from the statistics backend.

Examples:
  hepic statistic metrics --data '{"param":{}}'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dataStr, _ := cmd.Flags().GetString("data")

		var data interface{}
		if dataStr != "" {
			if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
				return fmt.Errorf("invalid JSON in --data: %w", err)
			}
		} else {
			data = map[string]interface{}{}
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := statistic.Metrics(context.Background(), client, data)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

var statisticMeasurementsCmd = &cobra.Command{
	Use:   "measurements <dbid>",
	Short: "Query measurements for a database",
	Long: `Query measurements for a specific database ID.

Examples:
  hepic statistic measurements mydb --data '{"param":{}}'`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		dbid := args[0]
		dataStr, _ := cmd.Flags().GetString("data")

		var data interface{}
		if dataStr != "" {
			if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
				return fmt.Errorf("invalid JSON in --data: %w", err)
			}
		} else {
			data = map[string]interface{}{}
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := statistic.Measurements(context.Background(), client, dbid, data)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

var statisticRetentionsCmd = &cobra.Command{
	Use:   "retentions",
	Short: "Query retention policies",
	Long: `Query retention policies from the statistics backend.

Examples:
  hepic statistic retentions --data '{"param":{}}'`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dataStr, _ := cmd.Flags().GetString("data")

		var data interface{}
		if dataStr != "" {
			if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
				return fmt.Errorf("invalid JSON in --data: %w", err)
			}
		} else {
			data = map[string]interface{}{}
		}

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := statistic.Retentions(context.Background(), client, data)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(statisticCmd)

	statisticCmd.AddCommand(statisticDBCmd)

	statisticCmd.AddCommand(statisticDataCmd)
	statisticDataCmd.Flags().String("data", "", "Query data as JSON string")
	statisticDataCmd.Flags().String("from", "", "Start time (RFC3339, date, or unix ms)")
	statisticDataCmd.Flags().String("to", "", "End time (RFC3339, date, or unix ms)")

	statisticCmd.AddCommand(statisticMetricsCmd)
	statisticMetricsCmd.Flags().String("data", "", "Query data as JSON string")

	statisticCmd.AddCommand(statisticMeasurementsCmd)
	statisticMeasurementsCmd.Flags().String("data", "", "Query data as JSON string")

	statisticCmd.AddCommand(statisticRetentionsCmd)
	statisticRetentionsCmd.Flags().String("data", "", "Query data as JSON string")
}
