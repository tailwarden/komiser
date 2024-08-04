package firestore

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"google.golang.org/api/option"

	firestore "cloud.google.com/go/firestore"
)

func Documents(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	firestoreClient, err := firestore.NewClient(ctx, client.GCPClient.Credentials.ProjectID, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		logrus.WithError(err).Errorf("failed to create firestore client")
		return resources, err
	}

	itr := firestoreClient.Collections(ctx)

	list, err := itr.GetAll()
	if err != nil {
		if strings.Contains(err.Error(), "SERVICE_DISABLED") {
			logrus.Warn(err.Error())
			return resources, nil
		} else {
			logrus.WithError(err).Errorf("failed to list all firestore collections")
			return resources, err
		}
	}

	for _, collection := range list {
		docItr := collection.Documents(ctx)
		documents, err := docItr.GetAll()
		if err != nil {
			logrus.WithError(err).Errorf("failed to list all firestore documents from %s collection", collection.ID)
			continue
		}
		for _, document := range documents {
			resources = append(resources, models.Resource{
				Provider:   "GCP",
				Account:    client.Name,
				Service:    "Firestore Documents",
				ResourceId: document.Ref.ID,
				Name:       document.Ref.ID,
				Cost:       0,
				FetchedAt:  time.Now(),
				Link:       fmt.Sprintf("https://console.cloud.google.com/firestore/databases?project=%s", client.GCPClient.Credentials.ProjectID),
			})
		}
	}

	logrus.WithFields(logrus.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "Firestore Documents",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
