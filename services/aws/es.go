package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
)

func (aws AWS) ListESDomains(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := elasticsearchservice.New(cfg)
		req := svc.ListDomainNamesRequest(&elasticsearchservice.ListDomainNamesInput{})
		res, err := req.Send(context.Background())
		if err != nil {
			return sum, err
		}

		sum += int64(len(res.DomainNames))
	}
	return sum, nil
}
