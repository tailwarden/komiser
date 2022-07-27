package aws

import (
	"context"

	awsConfig "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
	"github.com/aws/aws-sdk-go-v2/service/acm/types"
)

func (aws AWS) ListCertificates(cfg awsConfig.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := acm.NewFromConfig(cfg)
		res, err := svc.ListCertificates(context.Background(), &acm.ListCertificatesInput{})
		if err != nil {
			return sum, err
		}

		sum += int64(len(res.CertificateSummaryList))
	}
	return sum, nil
}

func (aws AWS) ListExpiredCertificates(cfg awsConfig.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := acm.NewFromConfig(cfg)
		res, err := svc.ListCertificates(context.Background(), &acm.ListCertificatesInput{
			CertificateStatuses: []types.CertificateStatus{
				types.CertificateStatusExpired,
			},
		})
		if err != nil {
			return sum, err
		}

		sum += int64(len(res.CertificateSummaryList))
	}
	return sum, nil
}
