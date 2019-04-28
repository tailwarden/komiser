package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/acm"
)

func (aws AWS) ListCertificates(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := acm.New(cfg)
		req := svc.ListCertificatesRequest(&acm.ListCertificatesInput{})
		res, err := req.Send(context.Background())
		if err != nil {
			return sum, err
		}

		sum += int64(len(res.CertificateSummaryList))
	}
	return sum, nil
}

func (aws AWS) ListExpiredCertificates(cfg aws.Config) (int64, error) {
	var sum int64
	regions, err := aws.getRegions(cfg)
	if err != nil {
		return 0, err
	}
	for _, region := range regions {
		cfg.Region = region.Name
		svc := acm.New(cfg)
		req := svc.ListCertificatesRequest(&acm.ListCertificatesInput{
			CertificateStatuses: []acm.CertificateStatus{
				acm.CertificateStatusExpired,
			},
		})
		res, err := req.Send(context.Background())
		if err != nil {
			return sum, err
		}

		sum += int64(len(res.CertificateSummaryList))
	}
	return sum, nil
}
