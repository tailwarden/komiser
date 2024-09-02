package providers

import (
	"context"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/civo/civogo"
	"github.com/digitalocean/godo"
	"github.com/linode/linodego"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/ovh/go-ovh/ovh"
	"github.com/patrickmn/go-cache"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/tailwarden/komiser/models"
	tccvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"go.mongodb.org/atlas/mongodbatlas"
	"golang.org/x/oauth2/google"
	"k8s.io/client-go/kubernetes"
)

type FetchDataFunction func(ctx context.Context, client ProviderClient) ([]models.Resource, error)

type ProviderClient struct {
	AWSClient          *aws.Config
	DigitalOceanClient *godo.Client
	OciClient          common.ConfigurationProvider
	CivoClient         *civogo.Client
	K8sClient          *K8sClient
	LinodeClient       *linodego.Client
	TencentClient      *tccvm.Client
	AzureClient        *AzureClient
	ScalewayClient     *scw.Client
	MongoDBAtlasClient *mongodbatlas.Client
	GCPClient          *GCPClient
	OVHClient          *ovh.Client
	Cache              *cache.Cache
	Name               string
}

type AzureClient struct {
	Credentials    *azidentity.ClientSecretCredential
	SubscriptionId string
}

type GCPClient struct {
	Credentials *google.Credentials
}

type K8sClient struct {
	Client          *kubernetes.Clientset
	OpencostBaseUrl string
}

type WorkerPool struct {
	numWorkers int
	tasks      chan func()
	wg         sync.WaitGroup
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		numWorkers: numWorkers,
		tasks:      make(chan func()),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		go wp.worker()
	}
}

func (wp *WorkerPool) SubmitTask(task func()) {
	wp.wg.Add(1)
	wp.tasks <- task
}

func (wp *WorkerPool) Wait() {
	wp.wg.Wait()
	close(wp.tasks)
}

func (wp *WorkerPool) worker() {
	for task := range wp.tasks {
		task()
		wp.wg.Done()
	}
}
