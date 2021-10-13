package azure

import (
	"context"
	"github.com/Azure/azure-go-for-sdk/services/web/mgmt/2020-12-01/web"
	"time"
)

func getCertificatesClient(subscriptionID string) web.CertificatesClient {
	certClient := web.NewCertificatesClient(subscriptionID)
	return certClient
}

func (azure Azure) ListCertificates(subscriptionID string) (int64, error) {
	certsClient := getCertificatesClient(subscriptionID)
	var filter string
	var sum int64
	ctx := context.Background()
	if certsClient != nil {
		certificateCollectionPage, err := certsClient.List(ctx, filter)

		for err != nil {
			certCollection := certificateCollectionPage.Response()
			sum += int64(len(certCollection.Count))
			err = certificateCollectionPage.NextWithContext(ctx)
		}

		if err != nil {
			return sum, err
		}

		return sum, nil
	}
}

func (azure Azure) ListExpiredCertificates(subscriptionID string) (int64, error) {
	certsClient := getCertificatesClient(subscriptionID)
	var filter string
	var sum int64
	ctx := context.Background()

	if certsClient != nil {
		certificateCollectionPage, err := certsClient.List(ctx, filter)

		for err != nil {
			certCollection := certificateCollectionPage.Response()
			certContracts := certCollection.Value

			for _, certContract := range certContracts {
				certContract.ExpirationDate.Sub(time.Now())

			}
		}
	}
}
