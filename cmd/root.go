package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "hepic",
	Short: "CLI for the HEPIC SIP capture and analysis platform",
	Long:  "hepic-cli provides command-line access to the HEPIC API for searching, analyzing, and managing SIP telecommunication data.",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, `{"error":"%s"}`+"\n", err.Error())
		return err
	}
	return nil
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().String("host", "", "HEPIC API host URL (overrides config/env)")
	rootCmd.PersistentFlags().String("token", "", "API key for authentication (overrides config/env)")
	rootCmd.PersistentFlags().String("format", "json", "Output format: json, table, yaml")
	rootCmd.PersistentFlags().Bool("verbose", false, "Enable verbose output (debug logging to stderr)")
	rootCmd.PersistentFlags().Bool("no-color", false, "Disable ANSI colors in output")

	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	viper.BindPFlag("format", rootCmd.PersistentFlags().Lookup("format"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("no-color", rootCmd.PersistentFlags().Lookup("no-color"))
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.hepic")

	viper.SetEnvPrefix("HEPIC")
	viper.AutomaticEnv()

	// Config file is optional — ignore if not found
	viper.ReadInConfig()
}
