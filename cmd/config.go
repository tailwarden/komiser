package cmd

import (
	"os"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/tailwarden/komiser/models"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Create configuration file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		c := models.Config{
			AWS: []models.AWSConfig{
				models.AWSConfig{
					Name:    "Demo",
					Source:  "CREDENTIALS_FILE",
					Profile: "default",
				},
			},
			SQLite: models.SQLiteConfig{
				File: "komiser.db",
			},
		}

		f, err := os.Create("config.toml")
		if err != nil {
			log.Fatal(err)
		}
		if err := toml.NewEncoder(f).Encode(c); err != nil {
			log.Fatal(err)
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)

		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
