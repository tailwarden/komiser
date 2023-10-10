package providers

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/civo/civogo"
	"github.com/digitalocean/godo"
	"github.com/linode/linodego"
	"github.com/oracle/oci-go-sdk/common"
	"github.com/ovh/go-ovh/ovh"
	"github.com/scaleway/scaleway-sdk-go/scw"
	. "github.com/tailwarden/komiser/models"
	tccvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	"go.mongodb.org/atlas/mongodbatlas"
	"golang.org/x/oauth2/google"
	"k8s.io/client-go/kubernetes"
)

type FetchDataFunction func(ctx context.Context, client ProviderClient) ([]Resource, error)

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
