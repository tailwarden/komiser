package azure

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
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
	//certificateCollectionPage, err := certsClient.List(ctx, filter)
	for cert, err := certsClient.ListComplete(ctx, filter); cert.NotDone(); cert.Next() {
		if err != nil {
			log.Print("got error while traverising", err)
		}
		sum = sum + 1
	}

	/* 	for err != nil {
	   		certCollection := certificateCollectionPage.Response()
	   		sum += int64(len(certCollection.Count))
	   		err = certificateCollectionPage.NextWithContext(ctx)
	   	}

	   	if err != nil {
	   		return sum, err
	   	} */

	return sum, nil
}

func (azure Azure) ListExpiredCertificates(subscriptionID string) (int64, error) {
	/* 	certsClient := getCertificatesClient(subscriptionID)
	   	var filter string
	   	var sum int64
	   	ctx := context.Background()

	   	certificateCollectionPage, err := certsClient.List(ctx, filter)

	   	for err != nil {
	   		certCollection := certificateCollectionPage.Response()
	   		certContracts := certCollection.Value
	   		var expiredCertCount int64
	   		for _, certContract := range certContracts {
	   			year, month, day := time.Now().Date()
	   			currentDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	   			timeLapsed := currentDate.Sub(certContract.ExpirationDate)
	   			if timeLapsed.Hours() > 0 {
	   				//No action needed as Certificate has not expired
	   			} else {
	   				expiredCertCount++
	   			}
	   		}
	   		sum += int64(len(certCollection.Count) - expiredCertCount)
	   		err = certificateCollectionPage.NextWithContext(ctx)
	   	} */
	var sum int64
	/* 	if err != nil {
		return sum, err
	} */

	return sum, nil
}
