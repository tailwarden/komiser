package aws

import (
	"context"
	"log"

	. "github.com/mlabouardy/komiser/providers"
	. "github.com/mlabouardy/komiser/providers/aws/cloudfront"
	. "github.com/mlabouardy/komiser/providers/aws/dynamodb"
	. "github.com/mlabouardy/komiser/providers/aws/ec2"
	. "github.com/mlabouardy/komiser/providers/aws/ecs"
	. "github.com/mlabouardy/komiser/providers/aws/eks"
	. "github.com/mlabouardy/komiser/providers/aws/iam"
	. "github.com/mlabouardy/komiser/providers/aws/lambda"
	. "github.com/mlabouardy/komiser/providers/aws/s3"
	"github.com/uptrace/bun"
)

func listOfSupportedServices() []FetchDataFunction {
	return []FetchDataFunction{
		Instances,
		Functions,
		Buckets,
		SecurityGroups,
		Roles,
		KubernetesClusters,
		Distributions,
		Tables,
		Tasks,
		Services,
		EcsClusters,
	}
}

func FetchAwsData(ctx context.Context, client ProviderClient, db *bun.DB) {
	for _, region := range getRegions() {
		client.AWSClient.Region = region
		for _, function := range listOfSupportedServices() {
			resources, err := function(ctx, client)
			if err != nil {
				log.Printf("[%s][AWS] %s", client.Name, err)
			} else {
				for _, resource := range resources {
					db.NewInsert().Model(&resource).Exec(context.Background())
				}
			}
		}
	}

}

func getRegions() []string {
	return []string{"us-east-1", "us-east-2", "us-west-1", "us-west-2", "ca-central-1", "eu-north-1", "eu-west-1", "eu-west-2", "eu-west-3", "eu-central-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-southeast-1", "ap-southeast-2", "ap-south-1", "sa-east-1"}
}
