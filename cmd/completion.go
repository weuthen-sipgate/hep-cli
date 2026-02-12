package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish]",
	Short: "Generate shell completion scripts",
	Long: `Generate shell completion scripts for hepic-cli.

Bash:
  $ source <(hepic completion bash)
  # Or permanently:
  $ hepic completion bash > /etc/bash_completion.d/hepic

Zsh:
  $ source <(hepic completion zsh)
  # Or permanently (ensure ~/.zsh/completion is in your fpath):
  $ hepic completion zsh > ~/.zsh/completion/_hepic

Fish:
  $ hepic completion fish | source
  # Or permanently:
  $ hepic completion fish > ~/.config/fish/completions/hepic.fish`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		switch args[0] {
		case "bash":
			return rootCmd.GenBashCompletion(os.Stdout)
		case "zsh":
			return rootCmd.GenZshCompletion(os.Stdout)
		case "fish":
			return rootCmd.GenFishCompletion(os.Stdout, true)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
