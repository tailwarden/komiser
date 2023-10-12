package config

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/BurntSushi/toml"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/civo/civogo"
	"github.com/digitalocean/godo"
	"github.com/linode/linodego"
	"github.com/mongodb-forks/digest"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/tailwarden/komiser/models"
	. "github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/utils"
	tccommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	tccvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"github.com/uptrace/bun"
	"go.mongodb.org/atlas/mongodbatlas"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func loadConfigFromFile(path string) (*Config, error) {
	filename, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("no such file %s", filename)
	}

	yamlFile, err := os.ReadFile(filename)
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

func Load(configPath string, telemetry bool, analytics utils.Analytics, db *bun.DB) (*Config, []providers.ProviderClient, []models.Account, error) {
	config, err := loadConfigFromFile(configPath)
	if err != nil {
		return nil, nil, nil, err
	}

	/*if len(config.SQLite.File) == 0 && config.Postgres.URI == "" {
		return nil, nil, nil, errors.New("postgres URI or sqlite file is missing")
	}*/

	clients := make([]providers.ProviderClient, 0)
	accounts := make([]models.Account, 0)

	if len(config.AWS) > 0 {
		for _, account := range config.AWS {
			cloudAccount := models.Account{
				Provider: "AWS",
				Name:     account.Name,
				Credentials: map[string]string{
					"profile": account.Profile,
					"path":    account.Path,
					"source":  account.Source,
				},
			}
			accounts = append(accounts, cloudAccount)

			if account.Source == "CREDENTIALS_FILE" {
				if len(account.Path) > 0 {
					cfg, err := awsConfig.LoadDefaultConfig(context.Background(), awsConfig.WithSharedConfigProfile(account.Profile), awsConfig.WithSharedCredentialsFiles(
						[]string{account.Path},
					))
					if err != nil {
						return nil, nil, nil, err
					}
					clients = append(clients, providers.ProviderClient{
						AWSClient: &cfg,
						Name:      account.Name,
					})
				} else {
					cfg, err := awsConfig.LoadDefaultConfig(context.Background(), awsConfig.WithSharedConfigProfile(account.Profile))
					if err != nil {
						return nil, nil, nil, err
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
		if telemetry {
			analytics.TrackEvent("connected_account", map[string]interface{}{
				"type":     len(config.AWS),
				"provider": "AWS",
			})
		}
	}

	if len(config.DigitalOcean) > 0 {
		for _, account := range config.DigitalOcean {
			cloudAccount := models.Account{
				Provider: "DigitalOcean",
				Name:     account.Name,
				Credentials: map[string]string{
					"token": account.Token,
				},
			}
			accounts = append(accounts, cloudAccount)

			client := godo.NewFromToken(account.Token)
			clients = append(clients, providers.ProviderClient{
				DigitalOceanClient: client,
				Name:               account.Name,
			})
		}
		if telemetry {
			analytics.TrackEvent("connected_account", map[string]interface{}{
				"type":     len(config.DigitalOcean),
				"provider": "DigitalOcean",
			})
		}
	}

	if len(config.Oci) > 0 {
		for _, account := range config.Oci {
			cloudAccount := models.Account{
				Provider: "OCI",
				Name:     account.Name,
				Credentials: map[string]string{
					"profile": account.Profile,
					"source":  account.Source,
				},
			}
			accounts = append(accounts, cloudAccount)

			if account.Source == "CREDENTIALS_FILE" {
				client := common.DefaultConfigProvider()
				clients = append(clients, providers.ProviderClient{
					OciClient: client,
					Name:      account.Name,
				})
			}
		}
		if telemetry {
			analytics.TrackEvent("connected_account", map[string]interface{}{
				"type":     len(config.Oci),
				"provider": "OCI",
			})
		}
	}

	if len(config.Civo) > 0 {
		for _, account := range config.Civo {
			cloudAccount := models.Account{
				Provider: "Civo",
				Name:     account.Name,
				Credentials: map[string]string{
					"token": account.Token,
				},
			}
			accounts = append(accounts, cloudAccount)

			client, err := civogo.NewClient(account.Token, "LON1")
			if err != nil {
				log.Fatal(err)
			}
			clients = append(clients, providers.ProviderClient{
				CivoClient: client,
				Name:       account.Name,
			})
		}
		if telemetry {
			analytics.TrackEvent("connected_account", map[string]interface{}{
				"type":     len(config.Civo),
				"provider": "Civo",
			})
		}
	}

	if len(config.Kubernetes) > 0 {
		for _, account := range config.Kubernetes {
			cloudAccount := models.Account{
				Provider: "Kubernetes",
				Name:     account.Name,
				Credentials: map[string]string{
					"path":     account.Path,
					"contexts": strings.Join(account.Contexts, ";"),
				},
			}

			accounts = append(accounts, cloudAccount)

			kubeConfig, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
				&clientcmd.ClientConfigLoadingRules{ExplicitPath: account.Path},
				&clientcmd.ConfigOverrides{}).ClientConfig()
			if err != nil {
				log.Fatal(err)
			}

			k8sClient, err := kubernetes.NewForConfig(kubeConfig)
			if err != nil {
				log.Fatal(err)
			}

			client := providers.K8sClient{
				Client:          k8sClient,
				OpencostBaseUrl: account.OpencostBaseUrl,
			}

			clients = append(clients, providers.ProviderClient{
				K8sClient: &client,
				Name:      account.Name,
			})
		}
		if telemetry {
			analytics.TrackEvent("connected_account", map[string]interface{}{
				"type":     len(config.Kubernetes),
				"provider": "Kubernetes",
			})
		}
	}

	if len(config.Linode) > 0 {
		for _, account := range config.Linode {
			cloudAccount := models.Account{
				Provider: "Linode",
				Name:     account.Name,
				Credentials: map[string]string{
					"token": account.Token,
				},
			}

			accounts = append(accounts, cloudAccount)

			tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: account.Token})
			oauth2Client := &http.Client{
				Transport: &oauth2.Transport{
					Source: tokenSource,
				},
			}
			client := linodego.NewClient(oauth2Client)

			clients = append(clients, providers.ProviderClient{
				LinodeClient: &client,
				Name:         account.Name,
			})
		}
		if telemetry {
			analytics.TrackEvent("connected_account", map[string]interface{}{
				"type":     len(config.Linode),
				"provider": "Linode",
			})
		}
	}

	if len(config.Tencent) > 0 {
		for _, account := range config.Tencent {
			cloudAccount := models.Account{
				Provider: "Tencent",
				Name:     account.Name,
				Credentials: map[string]string{
					"secretId":  account.SecretID,
					"secretKey": account.SecretKey,
				},
			}

			accounts = append(accounts, cloudAccount)

			credential := tccommon.NewCredential(account.SecretID, account.SecretKey)
			cpf := profile.NewClientProfile()
			cpf.Language = "en-US"
			client, err := tccvm.NewClient(credential, regions.Frankfurt, cpf)
			if err != nil {
				log.Fatal(err)
			}

			clients = append(clients, providers.ProviderClient{
				TencentClient: client,
				Name:          account.Name,
			})
		}
		if telemetry {
			analytics.TrackEvent("connected_account", map[string]interface{}{
				"type":     len(config.Tencent),
				"provider": "Tencent",
			})
		}
	}

	if len(config.Azure) > 0 {
		for _, account := range config.Azure {
			cloudAccount := models.Account{
				Provider: "Azure",
				Name:     account.Name,
				Credentials: map[string]string{
					"clientId":       account.ClientId,
					"clientSecret":   account.ClientSecret,
					"tenantId":       account.TenantId,
					"subscriptionId": account.SubscriptionId,
				},
			}

			accounts = append(accounts, cloudAccount)

			creds, err := azidentity.NewClientSecretCredential(account.TenantId, account.ClientId, account.ClientSecret, &azidentity.ClientSecretCredentialOptions{})
			if err != nil {
				log.Fatal(err)
			}

			client := providers.AzureClient{
				Credentials:    creds,
				SubscriptionId: account.SubscriptionId,
			}

			clients = append(clients, providers.ProviderClient{
				AzureClient: &client,
				Name:        account.Name,
			})
		}
		if telemetry {
			analytics.TrackEvent("connected_account", map[string]interface{}{
				"type":     len(config.Azure),
				"provider": "Azure",
			})
		}
	}

	if len(config.Scaleway) > 0 {
		for _, account := range config.Scaleway {
			cloudAccount := models.Account{
				Provider: "Scaleway",
				Name:     account.Name,
				Credentials: map[string]string{
					"accessKey":      account.AccessKey,
					"secretKey":      account.SecretKey,
					"organizationId": account.OrganizationId,
				},
			}

			accounts = append(accounts, cloudAccount)

			client, err := scw.NewClient(
				scw.WithDefaultOrganizationID(account.OrganizationId),
				scw.WithAuth(account.AccessKey, account.SecretKey),
			)
			if err != nil {
				log.Fatal(err)
			}

			clients = append(clients, providers.ProviderClient{
				ScalewayClient: client,
				Name:           account.Name,
			})
		}
		if telemetry {
			analytics.TrackEvent("connected_account", map[string]interface{}{
				"type":     len(config.Scaleway),
				"provider": "Scaleway",
			})
		}
	}

	if len(config.MongoDBAtlas) > 0 {
		for _, account := range config.MongoDBAtlas {
			cloudAccount := models.Account{
				Provider: "MongoDB",
				Name:     account.Name,
				Credentials: map[string]string{
					"publicKey":      account.PublicApiKey,
					"privateKey":     account.PrivateApiKey,
					"organizationId": account.OrganizationID,
				},
			}

			accounts = append(accounts, cloudAccount)

			t := digest.NewTransport(account.PublicApiKey, account.PrivateApiKey)
			tc, err := t.Client()
			if err != nil {
				log.Fatal(err.Error())
			}

			client := mongodbatlas.NewClient(tc)
			clients = append(clients, providers.ProviderClient{
				MongoDBAtlasClient: client,
				Name:               account.Name,
			})
		}
	}

	if len(config.GCP) > 0 {
		for _, account := range config.GCP {
			cloudAccount := models.Account{
				Provider: "GCP",
				Name:     account.Name,
				Credentials: map[string]string{
					"accountKey": account.ServiceAccountKeyPath,
				},
			}

			accounts = append(accounts, cloudAccount)

			data, err := os.ReadFile(account.ServiceAccountKeyPath)
			if err != nil {
				log.Fatal(err)
			}

			creds, err := google.CredentialsFromJSON(context.Background(), data, "https://www.googleapis.com/auth/cloud-platform")
			if err != nil {
				log.Fatal(err)
			}

			clients = append(clients, providers.ProviderClient{
				GCPClient: &providers.GCPClient{
					Credentials: creds,
				},
				Name: account.Name,
			})
		}
	}

	return config, clients, accounts, nil
}
