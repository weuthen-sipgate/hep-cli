package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Set via ldflags at build time
	Version   = "dev"
	BuildDate = "unknown"
	GitCommit = "unknown"
)

type versionInfo struct {
	Version   string `json:"version" yaml:"version"`
	BuildDate string `json:"build_date" yaml:"build_date"`
	GitCommit string `json:"git_commit" yaml:"git_commit"`
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show hepic-cli version information",
	RunE: func(cmd *cobra.Command, args []string) error {
		info := versionInfo{
			Version:   Version,
			BuildDate: BuildDate,
			GitCommit: GitCommit,
		}

		format := viper.GetString("format")
		switch format {
		case "table":
			fmt.Fprintf(os.Stdout, "hepic-cli %s (commit: %s, built: %s)\n", info.Version, info.GitCommit, info.BuildDate)
		default:
			enc := json.NewEncoder(os.Stdout)
			enc.SetIndent("", "  ")
			enc.Encode(info)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
