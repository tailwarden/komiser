package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
)

func (awsClient AWS) ListESDomains(cfg awsConfig.Config) (int64, error) {
	var sum int64
	regions, err := awsClient.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := elasticsearchservice.NewFromConfig(cfg)
		res, err := svc.ListDomainNames(context.Background(), &elasticsearchservice.ListDomainNamesInput{})
		if err != nil {
			return sum, err
		}

		sum += int64(len(res.DomainNames))
	}
	return sum, nil
}
