package certficate

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	certificatemanager "cloud.google.com/go/certificatemanager/apiv1"
	certificatemanagerpb "cloud.google.com/go/certificatemanager/apiv1/certificatemanagerpb"
)

func Certificates(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	certificateManagerClient, err := certificatemanager.NewClient(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		logrus.WithError(err).Errorf("failed to create certificate client")
		return resources, err
	}

	reg := &certificatemanagerpb.ListCertificatesRequest{
		Parent: "projects/" + client.GCPClient.Credentials.ProjectID + "/locations/global",
	}
	certificates := certificateManagerClient.ListCertificates(ctx, reg)

	for {
		certificate, err := certificates.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			if strings.Contains(err.Error(), "SERVICE_DISABLED") {
				logrus.Warn(err.Error())
				return resources, nil
			} else {
				logrus.WithError(err).Errorf("failed to list certificates")
				return resources, err
			}
		}

		certificateNameWithoutProjectAndLocation := extractCertificateName(certificate.Name)

		resources = append(resources, models.Resource{
			Provider:   "GCP",
			Account:    client.Name,
			Service:    "Certificate",
			ResourceId: certificate.Name,
			Name:       certificateNameWithoutProjectAndLocation,
			CreatedAt:  certificate.CreateTime.AsTime(),
			Cost:       0,
			Metadata:   certificate.Labels,
			FetchedAt:  time.Now(),
			Link:       fmt.Sprintf("https://console.cloud.google.com/security/ccm/certificates/details/global/name/%s?project=%s", certificateNameWithoutProjectAndLocation, client.GCPClient.Credentials.ProjectID),
		})

	}

	logrus.WithFields(logrus.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "Certificate Manager",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil

}

func extractCertificateName(s string) string {
	pattern := `projects\/[^\/]+\/locations\/[^\/]+\/certificates\/([^\/]+)`

	regex := regexp.MustCompile(pattern)
	match := regex.FindStringSubmatch(s)

	if len(match) > 1 {
		return match[1]
	}
	return s
}
