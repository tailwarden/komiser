package bigquery

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func Tables(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	bqclient, err := bigquery.NewClient(ctx, client.GCPClient.Credentials.ProjectID, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		return nil, err
	}

	datasetsIterator := bqclient.Datasets(ctx)

	for {
		dataset, err := datasetsIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			logrus.WithError(err).Errorf("failed to list dataset")
			return resources, err
		}

		tablesIterator := dataset.Tables(ctx)
		for {
			table, err := tablesIterator.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				logrus.WithError(err).Errorf("failed to list tables")
				return resources, err
			}

			tableMetadata, err := table.Metadata(ctx)
			if err != nil {
				logrus.WithError(err).Errorf("failed to load table metadata")
				return resources, err
			}

			tags := make([]models.Tag, 0)

			for key, value := range tableMetadata.Labels {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: value,
				})
			}

			resources = append(resources, models.Resource{
				Provider:   "GCP",
				Account:    client.Name,
				Service:    "BigQuery Table",
				ResourceId: table.TableID,
				Region:     tableMetadata.Location,
				Name:       tableMetadata.Name,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https://console.cloud.google.com/bigquery?project=%s&page=dataset&p=%s&d=%s", client.GCPClient.Credentials.ProjectID, client.GCPClient.Credentials.ProjectID, dataset.DatasetID),
			})
		}

	}

	logrus.WithFields(logrus.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "BigQuery",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
