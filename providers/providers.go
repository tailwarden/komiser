package providers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/digitalocean/godo"
	. "github.com/mlabouardy/komiser/models"
)

type FetchDataFunction func(ctx context.Context, client ProviderClient) ([]Resource, error)

type ProviderClient struct {
	AWSClient          *aws.Config
	DigitalOceanClient *godo.Client
	Name               string
}
