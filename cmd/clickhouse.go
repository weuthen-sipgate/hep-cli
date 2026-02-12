package cmd

import (
	"context"

	"hepic-cli/internal/admin"
	"hepic-cli/internal/api"
	"hepic-cli/internal/output"

	"github.com/spf13/cobra"
)

var clickhouseCmd = &cobra.Command{
	Use:   "clickhouse",
	Short: "ClickHouse database operations",
	Long:  "Execute queries against the ClickHouse database backend.",
}

var clickhouseQueryCmd = &cobra.Command{
	Use:   "query <sql>",
	Short: "Execute a raw ClickHouse query",
	Long: `Execute a raw SQL query against the ClickHouse database and return the results.

The query is provided as a positional argument.

Examples:
  hepic clickhouse query "SELECT count() FROM hep"
  hepic clickhouse query "SELECT * FROM hep LIMIT 10"`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := args[0]

		client, err := api.NewClient()
		if err != nil {
			return err
		}

		result, err := admin.ClickhouseQuery(context.Background(), client, query)
		if err != nil {
			output.PrintError(err)
			return err
		}

		return output.Print(result)
	},
}

func init() {
	rootCmd.AddCommand(clickhouseCmd)

	clickhouseCmd.AddCommand(clickhouseQueryCmd)
}
