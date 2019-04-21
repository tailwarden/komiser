package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/swf"
)

func (aws AWS) GetSWFDomains(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := swf.New(cfg)
		req := svc.ListDomainsRequest(&swf.ListDomainsInput{})
		res, _ := req.Send(context.Background())
		if res != nil {
			sum += int64(len(res.DomainInfos))
		}
	}
	return sum, nil
}
