package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/tailwarden/komiser/models"
)

const DefaultFileName = "config.toml"

var demoConfig = models.Config{
	AWS: []models.AWSConfig{
		{
			Name:    "Demo",
			Source:  "CREDENTIALS_FILE",
			Profile: "default",
		},
	},
	SQLite: models.SQLiteConfig{
		File: "komiser.db",
	},
}

func Create(c *models.Config) error {
	if c == nil {
		c = &demoConfig
	}

	f, err := os.Create(DefaultFileName)
	if err != nil {
		return err
	}
	if err := toml.NewEncoder(f).Encode(*c); err != nil {
		return err
	}
	return f.Close()
}
