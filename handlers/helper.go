package handlers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

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
		} else if account.Credentials["source"] == "environment-variables" {
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

		return &providers.ProviderClient{
			K8sClient: &client,
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
