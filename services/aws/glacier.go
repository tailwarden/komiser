package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
)

func (aws AWS) ListVaults(cfg awsConfig.Config) (map[string]int, error) {
	output := make(map[string]int, 0)
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return output, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := glacier.NewFromConfig(cfg)
		res, err := svc.ListVaults(context.Background(), &glacier.ListVaultsInput{
			AccountId: awsConfig.String("-"),
		})
		if err != nil {
			return output, err
		}

		output["vaults"] += len(res.VaultList)

		for _, vault := range res.VaultList {
			output["total"] += int(vault.SizeInBytes)
		}

	}
	return output, nil
}
