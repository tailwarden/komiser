package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tailwarden/komiser/internal/config"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Create configuration file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := config.Create(nil)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
