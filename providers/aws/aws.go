package aws

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	. "github.com/mlabouardy/komiser/providers/aws/ec2"
	"github.com/uptrace/bun"
)

func FetchAwsData(ctx context.Context, cfg aws.Config, account string, db *bun.DB) {
	for _, region := range getRegions() {
		cfg.Region = region
		resources, err := Instances(ctx, cfg, account)
		if err != nil {
			log.Println(err)
		} else {
			for _, resource := range resources {
				db.NewInsert().Model(&resource).Exec(context.Background())
			}
		}
	}

}

func getRegions() []string {
	return []string{"us-east-1", "us-east-2", "us-west-1", "us-west-2", "ca-central-1", "eu-north-1", "eu-west-1", "eu-west-2", "eu-west-3", "eu-central-1", "ap-northeast-1", "ap-northeast-2", "ap-northeast-3", "ap-southeast-1", "ap-southeast-2", "ap-south-1", "sa-east-1"}
}
