package providers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/civo/civogo"
	"github.com/digitalocean/godo"
	"github.com/linode/linodego"
	"github.com/oracle/oci-go-sdk/common"
	. "github.com/tailwarden/komiser/models"
	"k8s.io/client-go/kubernetes"
)

type FetchDataFunction func(ctx context.Context, client ProviderClient) ([]Resource, error)

type ProviderClient struct {
	AWSClient          *aws.Config
	DigitalOceanClient *godo.Client
	OciClient          common.ConfigurationProvider
	CivoClient         *civogo.Client
	K8sClient          *kubernetes.Clientset
	LinodeClient       *linodego.Client
	Name               string
}
