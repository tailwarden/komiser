package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsClient "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
)

func (aws AWS) ListVaults(cfg aws.Config) (map[string]int, error) {
	output := make(map[string]int, 0)
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return output, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := glacier.New(cfg)
		req := svc.ListVaultsRequest(&glacier.ListVaultsInput{
			AccountId: awsClient.String("-"),
		})
		res, err := req.Send(context.Background())
		if err != nil {
			return output, err
		}

		output["vaults"] += len(res.VaultList)

		for _, vault := range res.VaultList {
			output["total"] += int(*vault.SizeInBytes)
		}

	}
	return output, nil
}
