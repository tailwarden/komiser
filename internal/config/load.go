package config

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/civo/civogo"
	"github.com/digitalocean/godo"
	. "github.com/mlabouardy/komiser/models"
	"github.com/mlabouardy/komiser/providers"
	"github.com/oracle/oci-go-sdk/common"
)

func loadConfigFromFile(path string) (*Config, error) {
	filename, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("no such file %s", filename)
	}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return loadConfigFromBytes(yamlFile)
}

func loadConfigFromBytes(b []byte) (*Config, error) {
	var config Config

	err := toml.Unmarshal([]byte(b), &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func Load(configPath string) (*Config, []providers.ProviderClient, error) {
	config, err := loadConfigFromFile(configPath)
	if err != nil {
		return nil, nil, err
	}

	if len(config.SQLite.File) == 0 && config.Postgres.URI == "" {
		return nil, nil, errors.New("postgres URI or sqlite file is missing")
	}

	clients := make([]providers.ProviderClient, 0)

	if len(config.AWS) > 0 {
		for _, account := range config.AWS {
			if account.Source == "CREDENTIALS_FILE" {
				if len(account.Path) > 0 {
					cfg, err := awsConfig.LoadDefaultConfig(context.Background(), awsConfig.WithSharedConfigProfile(account.Profile), awsConfig.WithSharedCredentialsFiles(
						[]string{account.Path},
					))
					if err != nil {
						return nil, nil, err
					}
					clients = append(clients, providers.ProviderClient{
						AWSClient: &cfg,
						Name:      account.Name,
					})
				} else {
					cfg, err := awsConfig.LoadDefaultConfig(context.Background(), awsConfig.WithSharedConfigProfile(account.Profile))
					if err != nil {
						return nil, nil, err
					}
					clients = append(clients, providers.ProviderClient{
						AWSClient: &cfg,
						Name:      account.Name,
					})
				}
			} else if account.Source == "ENVIRONMENT_VARIABLES" {
				cfg, err := awsConfig.LoadDefaultConfig(context.Background())
				if err != nil {
					log.Fatal(err)
				}
				clients = append(clients, providers.ProviderClient{
					AWSClient: &cfg,
					Name:      account.Name,
				})
			}
		}
	}

	if len(config.DigitalOcean) > 0 {
		for _, account := range config.DigitalOcean {
			client := godo.NewFromToken(account.Token)
			clients = append(clients, providers.ProviderClient{
				DigitalOceanClient: client,
				Name:               account.Name,
			})
		}
	}

	if len(config.Oci) > 0 {
		for _, account := range config.Oci {
			if account.Source == "CREDENTIALS_FILE" {
				client := common.DefaultConfigProvider()
				clients = append(clients, providers.ProviderClient{
					OciClient: client,
					Name:      account.Name,
				})
			}
		}
	}

	if len(config.Civo) > 0 {
		for _, account := range config.Civo {
			client, err := civogo.NewClient(account.Token, "LON1")
			if err != nil {
				log.Fatal(err)
			}
			clients = append(clients, providers.ProviderClient{
				CivoClient: client,
				Name:       account.Name,
			})
		}
	}

	return config, clients, nil
}
