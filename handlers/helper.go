package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/patrickmn/go-cache"
	tccommon "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tccvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/civo/civogo"
	"github.com/digitalocean/godo"
	"github.com/linode/linodego"
	"github.com/mongodb-forks/digest"
	"github.com/oracle/oci-go-sdk/common"
	ovhPkg "github.com/ovh/go-ovh/ovh"
	"github.com/scaleway/scaleway-sdk-go/scw"
	log "github.com/sirupsen/logrus"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	"github.com/uptrace/bun"
	mdb "go.mongodb.org/atlas/mongodbatlas"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/getsentry/sentry-go"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/providers/aws"
	"github.com/tailwarden/komiser/providers/azure"
	"github.com/tailwarden/komiser/providers/civo"
	do "github.com/tailwarden/komiser/providers/digitalocean"
	"github.com/tailwarden/komiser/providers/gcp"
	"github.com/tailwarden/komiser/providers/k8s"
	"github.com/tailwarden/komiser/providers/linode"
	"github.com/tailwarden/komiser/providers/mongodbatlas"
	"github.com/tailwarden/komiser/providers/oci"
	"github.com/tailwarden/komiser/providers/ovh"
	"github.com/tailwarden/komiser/providers/scaleway"
	"github.com/tailwarden/komiser/providers/tencent"
	"github.com/tailwarden/komiser/utils"
)

const (
	CACHE_DURATION   = 3
	CLEANUP_DURATION = 4
)

var mu sync.Mutex

func triggerFetchingWorkflow(ctx context.Context, client providers.ProviderClient, provider string, db *bun.DB, regions []string, wp *providers.WorkerPool) {
	localHub := sentry.CurrentHub().Clone()

	defer func() {
		err := recover()
		if err != nil {
			log.WithField("err", err).Error(fmt.Sprintf("error fetching %s resources", provider))
			localHub.CaptureException(err.(error))
			localHub.Flush(2 * time.Second)
		}
	}()

	localHub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTag("provider", provider)
	})

	var analytics utils.Analytics
	telemetry := false
	switch provider {
	case "AWS":
		aws.FetchResources(ctx, client, regions, db, telemetry, analytics, wp)
	case "DigitalOcean":
		do.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "OCI":
		oci.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "Civo":
		civo.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "Kubernetes":
		k8s.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "Linode":
		linode.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "Tencent":
		tencent.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "Azure":
		azure.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "Scaleway":
		scaleway.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "MongoDBAtlas":
		mongodbatlas.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "GCP":
		gcp.FetchResources(ctx, client, db, telemetry, analytics, wp)
	case "OVH":
		ovh.FetchResources(ctx, client, db, telemetry, analytics, wp)
	}
}

func fetchResourcesForAccount(ctx context.Context, account models.Account, db *bun.DB, regions []string) {
	numWorkers := 64
	account.Status = ""
	wp := providers.NewWorkerPool(numWorkers)
	wp.Start()

	var wwg sync.WaitGroup
	workflowTrigger := func(client providers.ProviderClient, provider string) {
		wwg.Add(1)
		go func() {
			defer wwg.Done()
			triggerFetchingWorkflow(ctx, client, provider, db, regions, wp)
		}()
	}

	client, err := makeClientFromAccount(account)
	if err != nil {
		log.Error(err, account.Id, account.Provider)
		_, err := db.NewUpdate().Model(&account).Set("status = ? ", "INTEGRATION ISSUE").Where("name = ?", account.Name).Exec(ctx)
		if err != nil {
			log.Error(err)
			return
		}
		return
	}
	_, err = db.NewUpdate().Model(&account).Set("status = ? ", "SCANNING").Where("name = ?", account.Name).Exec(ctx)
	if err != nil {
		log.Error("Couldn't set status")
		return
	}
	log.Info("Scanning status set")
	if client.AWSClient != nil {
		workflowTrigger(*client, "AWS")
	} else if client.DigitalOceanClient != nil {
		workflowTrigger(*client, "DigitalOcean")
	} else if client.OciClient != nil {
		workflowTrigger(*client, "OCI")
	} else if client.CivoClient != nil {
		workflowTrigger(*client, "Civo")
	} else if client.K8sClient != nil {
		workflowTrigger(*client, "Kubernetes")
	} else if client.LinodeClient != nil {
		workflowTrigger(*client, "Linode")
	} else if client.TencentClient != nil {
		workflowTrigger(*client, "Tencent")
	} else if client.AzureClient != nil {
		workflowTrigger(*client, "Azure")
	} else if client.ScalewayClient != nil {
		workflowTrigger(*client, "Scaleway")
	} else if client.MongoDBAtlasClient != nil {
		workflowTrigger(*client, "MongoDBAtlas")
	} else if client.GCPClient != nil {
		workflowTrigger(*client, "GCP")
	} else if client.OVHClient != nil {
		workflowTrigger(*client, "OVH")
	}

	wwg.Wait()
	wp.Wait()
	_, err = db.NewUpdate().Model(&account).Set("status = ? ", "CONNECTED").Where("name = ?", account.Name).Exec(ctx)
	if err != nil {
		log.Error("Couldn't set status")
		return
	}
	log.Info("Scanning done")
}

func makeClientFromAccount(account models.Account) (*providers.ProviderClient, error) {
	if account.Provider == "aws" {
		if account.Credentials["source"] == "credentials-file" {
			if len(account.Credentials["path"]) > 0 {
				cfg, err := awsConfig.LoadDefaultConfig(context.Background(), awsConfig.WithSharedConfigProfile(account.Credentials["profile"]), awsConfig.WithSharedCredentialsFiles(
					[]string{account.Credentials["path"]},
				))
				if err != nil {
					return nil, err
				}
				return &providers.ProviderClient{
					AWSClient: &cfg,
					Name:      account.Name,
				}, nil
			} else {
				cfg, err := awsConfig.LoadDefaultConfig(context.Background(), awsConfig.WithSharedConfigProfile(account.Credentials["profile"]))
				if err != nil {
					return nil, err
				}
				return &providers.ProviderClient{
					AWSClient: &cfg,
					Name:      account.Name,
				}, err
			}
		} else {
			cfg, err := awsConfig.LoadDefaultConfig(context.Background())
			if err != nil {
				return nil, err
			}
			return &providers.ProviderClient{
				AWSClient: &cfg,
				Name:      account.Name,
			}, nil
		}
	}

	if account.Provider == "digitalocean" {
		client := godo.NewFromToken(account.Credentials["token"])
		return &providers.ProviderClient{
			DigitalOceanClient: client,
			Name:               account.Name,
		}, nil
	}

	if account.Provider == "oci" {
		if account.Credentials["source"] == "CREDENTIALS_FILE" {
			return &providers.ProviderClient{
				OciClient: common.DefaultConfigProvider(),
				Name:      account.Name,
			}, nil
		}
	}

	if account.Provider == "civo" {
		client, err := civogo.NewClient(account.Credentials["token"], "LON1")
		if err != nil {
			return nil, err
		}
		return &providers.ProviderClient{
			CivoClient: client,
			Name:       account.Name,
		}, nil
	}

	if account.Provider == "kubernetes" {
		kubeConfig, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			&clientcmd.ClientConfigLoadingRules{ExplicitPath: account.Credentials["path"]},
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
			OpencostBaseUrl: account.Credentials["opencostBaseUrl"],
		}

		cache := cache.New(CACHE_DURATION, CLEANUP_DURATION)
		return &providers.ProviderClient{
			K8sClient: &client,
			Cache:     cache, // Alpha feature for dependency
			Name:      account.Name,
		}, nil
	}

	if account.Provider == "linode" {
		tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: account.Credentials["token"]})
		oauth2Client := &http.Client{
			Transport: &oauth2.Transport{
				Source: tokenSource,
			},
		}
		client := linodego.NewClient(oauth2Client)
		return &providers.ProviderClient{
			LinodeClient: &client,
			Name:         account.Name,
		}, nil
	}

	if account.Provider == "tencent" {
		credential := tccommon.NewCredential(account.Credentials["secretId"], account.Credentials["secretKey"])
		cpf := profile.NewClientProfile()
		cpf.Language = "en-US"
		client, err := tccvm.NewClient(credential, regions.Frankfurt, cpf)
		if err != nil {
			return nil, err
		}

		return &providers.ProviderClient{
			TencentClient: client,
			Name:          account.Name,
		}, nil
	}

	if account.Provider == "azure" {
		creds, err := azidentity.NewClientSecretCredential(account.Credentials["tenantId"], account.Credentials["clientId"], account.Credentials["clientSecret"], &azidentity.ClientSecretCredentialOptions{})
		if err != nil {
			log.Fatal(err)
		}

		client := providers.AzureClient{
			Credentials:    creds,
			SubscriptionId: account.Credentials["subscriptionId"],
		}

		return &providers.ProviderClient{
			AzureClient: &client,
			Name:        account.Name,
		}, nil
	}

	if account.Provider == "scaleway" {
		client, err := scw.NewClient(
			scw.WithDefaultOrganizationID(account.Credentials["organizationId"]),
			scw.WithAuth(account.Credentials["accessKey"], account.Credentials["secretKey"]),
		)
		if err != nil {
			return nil, err
		}

		return &providers.ProviderClient{
			ScalewayClient: client,
			Name:           account.Name,
		}, nil
	}

	if account.Provider == "mongodb" {
		t := digest.NewTransport(account.Credentials["publicApiKey"], account.Credentials["privateApiKey"])
		tc, err := t.Client()
		if err != nil {
			log.Fatal(err.Error())
		}

		client := mdb.NewClient(tc)
		return &providers.ProviderClient{
			MongoDBAtlasClient: client,
			Name:               account.Name,
		}, nil
	}

	if account.Provider == "gcp" {
		data, err := os.ReadFile(account.Credentials["accountKey"])
		if err != nil {
			log.Fatal(err)
		}

		creds, err := google.CredentialsFromJSON(context.Background(), data, "https://www.googleapis.com/auth/cloud-platform")
		if err != nil {
			log.Fatal(err)
		}

		return &providers.ProviderClient{
			GCPClient: &providers.GCPClient{
				Credentials: creds,
			},
			Name: account.Name,
		}, nil
	}

	if account.Provider == "ovh" {
		client, err := ovhPkg.NewClient(
			account.Credentials["endpoint"],
			account.Credentials["applicationKey"],
			account.Credentials["applicationSecret"],
			account.Credentials["consumerKey"],
		)
		if err != nil {
			return nil, err
		}

		return &providers.ProviderClient{
			OVHClient: client,
			Name:      account.Name,
		}, nil
	}
	return nil, fmt.Errorf("provider not supported")
}

func populateConfigFromAccount(account models.Account, config *models.Config) error {
	switch account.Provider {
	case "aws":
		awsConfig := models.AWSConfig{
			Name:   account.Name,
			Source: account.Credentials["source"],
		}
		if account.Credentials["source"] == "credentials-file" {
			awsConfig.Profile = account.Credentials["profile"]
			if path, ok := account.Credentials["path"]; ok && len(path) > 0 {
				awsConfig.Path = account.Credentials["path"]
			}
		}
		config.AWS = append(config.AWS, awsConfig)

	case "digitalocean":
		digitalOceanConfig := models.DigitalOceanConfig{
			Name:  account.Name,
			Token: account.Credentials["token"],
		}
		config.DigitalOcean = append(config.DigitalOcean, digitalOceanConfig)

	case "oci":
		ociConfig := models.OciConfig{
			Name:    account.Name,
			Profile: account.Credentials["profile"],
			Source:  account.Credentials["source"],
		}
		config.Oci = append(config.Oci, ociConfig)

	case "civo":
		civoConfig := models.CivoConfig{
			Name:  account.Name,
			Token: account.Credentials["token"],
		}
		config.Civo = append(config.Civo, civoConfig)

	case "kubernetes":
		k8sConfig := models.KubernetesConfig{
			Name:     account.Name,
			Path:     account.Credentials["path"],
			Contexts: strings.Split(account.Credentials["contexts"], ";"),
		}
		config.Kubernetes = append(config.Kubernetes, k8sConfig)

	case "linode":
		linodeConfig := models.LinodeConfig{
			Name:  account.Name,
			Token: account.Credentials["token"],
		}
		config.Linode = append(config.Linode, linodeConfig)

	case "tencent":
		tencentConfig := models.TencentConfig{
			Name:      account.Name,
			SecretID:  account.Credentials["secretId"],
			SecretKey: account.Credentials["secretKey"],
		}
		config.Tencent = append(config.Tencent, tencentConfig)

	case "azure":
		azureConfig := models.AzureConfig{
			Name:           account.Name,
			ClientId:       account.Credentials["clientId"],
			ClientSecret:   account.Credentials["clientSecret"],
			TenantId:       account.Credentials["tenantId"],
			SubscriptionId: account.Credentials["subscriptionId"],
		}
		config.Azure = append(config.Azure, azureConfig)

	case "scaleway":
		scalewayConfig := models.ScalewayConfig{
			Name:           account.Name,
			AccessKey:      account.Credentials["accessKey"],
			SecretKey:      account.Credentials["secretKey"],
			OrganizationId: account.Credentials["organizationId"],
		}
		config.Scaleway = append(config.Scaleway, scalewayConfig)

	case "mongodb":
		mongoDBAtlasConfig := models.MongoDBAtlasConfig{
			Name:           account.Name,
			PublicApiKey:   account.Credentials["publicKey"],
			PrivateApiKey:  account.Credentials["privateKey"],
			OrganizationID: account.Credentials["organizationId"],
		}
		config.MongoDBAtlas = append(config.MongoDBAtlas, mongoDBAtlasConfig)

	case "gcp":
		gcpConfig := models.GCPConfig{
			Name:                  account.Name,
			ServiceAccountKeyPath: account.Credentials["accountKey"],
		}
		config.GCP = append(config.GCP, gcpConfig)

	case "ovh":
		ovhConfig := models.OVHConfig{
			Name:              account.Name,
			Endpoint:          account.Credentials["endpoint"],
			ApplicationKey:    account.Credentials["applicationKey"],
			ApplicationSecret: account.Credentials["applicationSecret"],
			ConsumerKey:       account.Credentials["consumerKey"],
		}
		config.OVH = append(config.OVH, ovhConfig)

	default:
		return fmt.Errorf("illegle provider")
	}

	return nil
}

func deleteConfigAccounts(account models.Account, config *models.Config) error {
	switch strings.ToLower(account.Provider) {
	case "aws":

		updatedConfig := make([]models.AWSConfig, 0)
		for _, acc := range config.AWS {
			if acc.Name != account.Name {
				updatedConfig = append(updatedConfig, acc)
			}
		}
		config.AWS = updatedConfig

	case "digitalocean":
		updatedConfig := make([]models.DigitalOceanConfig, 0)
		for _, acc := range config.DigitalOcean {
			if acc.Name != account.Name {
				updatedConfig = append(updatedConfig, acc)
			}
		}
		config.DigitalOcean = updatedConfig

	case "oci":
		updatedConfig := make([]models.OciConfig, 0)
		for _, acc := range config.Oci {
			if acc.Name != account.Name {
				updatedConfig = append(updatedConfig, acc)
			}
		}
		config.Oci = updatedConfig

	case "civo":
		updatedConfig := make([]models.CivoConfig, 0)
		for _, acc := range config.Civo {
			if acc.Name != account.Name {
				updatedConfig = append(updatedConfig, acc)
			}
		}
		config.Civo = updatedConfig

	case "kubernetes":
		updatedConfig := make([]models.KubernetesConfig, 0)
		for _, acc := range config.Kubernetes {
			if acc.Name != account.Name {
				updatedConfig = append(updatedConfig, acc)
			}
		}
		config.Kubernetes = updatedConfig

	case "linode":
		updatedConfig := make([]models.LinodeConfig, 0)
		for _, acc := range config.Linode {
			if acc.Name != account.Name {
				updatedConfig = append(updatedConfig, acc)
			}
		}
		config.Linode = updatedConfig

	case "tencent":
		updatedConfig := make([]models.TencentConfig, 0)
		for _, acc := range config.Tencent {
			if acc.Name != account.Name {
				updatedConfig = append(updatedConfig, acc)
			}
		}
		config.Tencent = updatedConfig

	case "azure":
		updatedConfig := make([]models.AzureConfig, 0)
		for _, acc := range config.Azure {
			if acc.Name != account.Name {
				updatedConfig = append(updatedConfig, acc)
			}
		}
		config.Azure = updatedConfig

	case "scaleway":
		updatedConfig := make([]models.ScalewayConfig, 0)
		for _, acc := range config.Scaleway {
			if acc.Name != account.Name {
				updatedConfig = append(updatedConfig, acc)
			}
		}
		config.Scaleway = updatedConfig

	case "mongodb":
		updatedConfig := make([]models.MongoDBAtlasConfig, 0)
		for _, acc := range config.MongoDBAtlas {
			if acc.Name != account.Name {
				updatedConfig = append(updatedConfig, acc)
			}
		}
		config.MongoDBAtlas = updatedConfig

	case "gcp":
		updatedConfig := make([]models.GCPConfig, 0)
		for _, acc := range config.GCP {
			if acc.Name != account.Name {
				updatedConfig = append(updatedConfig, acc)
			}
		}
		config.GCP = updatedConfig

	case "ovh":
		updatedConfig := make([]models.OVHConfig, 0)
		for _, acc := range config.OVH {
			if acc.Name != account.Name {
				updatedConfig = append(updatedConfig, acc)
			}
		}
		config.OVH = updatedConfig

	default:
		return fmt.Errorf("illegle provider")
	}

	return nil
}

func updateConfig(path string, cfg *models.Config) error {
	mu.Lock()
	defer mu.Unlock()

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := toml.NewEncoder(f).Encode(cfg); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
