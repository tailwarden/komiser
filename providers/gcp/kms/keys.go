package kms

import (
	"context"
	"fmt"
	"strings"
	"time"

	kms "cloud.google.com/go/kms/apiv1"
	kmspb "cloud.google.com/go/kms/apiv1/kmspb"
	"github.com/sirupsen/logrus"
	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func Keys(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	kmsclient, err := kms.NewKeyManagementClient(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		return resources, err
	}

	parent := fmt.Sprintf("projects/%s/locations/global", client.GCPClient.Credentials.ProjectID)

	keyRingsIterator := kmsclient.ListKeyRings(ctx, &kmspb.ListKeyRingsRequest{Parent: parent})
	if keyRingsIterator == nil {
		logrus.Errorf("ListKeyRings returned nil iterator")
		return resources, nil
	}

	for {
		keyRing, err := keyRingsIterator.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			if strings.Contains(err.Error(), "SERVICE_DISABLED") {
				logrus.Warn(err.Error())
				return resources, nil
			} else {
				logrus.WithError(err).Errorf("failed to list key rings")
				return resources, err
			}
		}
		if keyRing == nil {
			continue
		}

		keyRingName := keyRing.GetName()[len(parent)+len("/keyRings/"):]

		keysIterator := kmsclient.ListCryptoKeys(ctx, &kmspb.ListCryptoKeysRequest{
			Parent: keyRing.GetName(),
		})
		if keysIterator == nil {
			logrus.Errorf("ListCryptoKeys returned nil iterator")
			return resources, nil
		}
		for {
			key, err := keysIterator.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				logrus.WithError(err).Errorf("failed to list keys")
				return resources, err
			}
			if key == nil {
				continue
			}

			keyMetadata, err := kmsclient.GetCryptoKey(ctx, &kmspb.GetCryptoKeyRequest{
				Name: key.GetName(),
			})
			if err != nil {
				logrus.WithError(err).Errorf("failed to load key metadata")
				return resources, err
			}
			if keyMetadata == nil {
				continue
			}

			tags := make([]models.Tag, 0)

			for key, value := range keyMetadata.GetLabels() {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: value,
				})
			}

			keyName := key.GetName()
			lastSlashIndex := strings.LastIndex(keyName, "/")
			if lastSlashIndex != -1 && lastSlashIndex < len(keyName)-1 {
				keyName = keyName[lastSlashIndex+1:]
			}

			resources = append(resources, models.Resource{
				Provider:   "GCP",
				Account:    client.Name,
				Service:    "KMS Key",
				ResourceId: key.GetName(),
				Region:     "global",
				Name:       keyName,
				FetchedAt:  time.Now(),
				Tags:       tags,
				Link:       fmt.Sprintf("https://console.cloud.google.com/security/kms/key/manage/global/%s/%s?project=%s", keyRingName, keyName, client.GCPClient.Credentials.ProjectID),
			})
		}
	}

	logrus.WithFields(logrus.Fields{
		"provider":  "GCP",
		"account":   client.Name,
		"service":   "KMS",
		"resources": len(resources),
	}).Info("Fetched resources")

	return resources, nil
}
