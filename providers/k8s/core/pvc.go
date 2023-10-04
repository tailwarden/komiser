package core

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	. "github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func PersistentVolumeClaims(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)

	var config metav1.ListOptions

	for {
		res, err := client.K8sClient.Client.CoreV1().PersistentVolumeClaims("").List(ctx, config)
		if err != nil {
			return nil, err
		}

		for _, pvc := range res.Items {
			tags := make([]Tag, 0)

			for key, value := range pvc.Labels {
				tags = append(tags, Tag{
					Key:   key,
					Value: value,
				})
			}

			resources = append(resources, Resource{
				Provider:   "Kubernetes",
				Account:    client.Name,
				Service:    "PersistentVolumeClaim",
				ResourceId: string(pvc.UID),
				Name:       pvc.Name,
				Region:     pvc.Namespace,
				Cost:       0,
				CreatedAt:  pvc.CreationTimestamp.Time,
				FetchedAt:  time.Now(),
				Tags:       tags,
			})
		}

		if res.GetContinue() == "" {
			break
		}

		config.Continue = res.GetContinue()
	}

	log.WithFields(log.Fields{
		"provider":  "Kubernetes",
		"account":   client.Name,
		"service":   "PersistentVolumeClaims",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
