package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "komiser",
	Short: "Cloud environment inspector",
	Long: `Komiser enables you to have a clear view into your cloud account,
gives helpful advice to reduce the cost and secure your environment.`,
}

func Execute() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
