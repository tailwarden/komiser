package core

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	. "github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Ingress(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)

	var config metav1.ListOptions

	for {
		res, err := client.K8sClient.Client.NetworkingV1().Ingresses("").List(ctx, config)
		if err != nil {
			return nil, err
		}

		for _, ingress := range res.Items {
			tags := make([]Tag, 0)

			for key, value := range ingress.Labels {
				tags = append(tags, Tag{
					Key:   key,
					Value: value,
				})
			}

			resources = append(resources, Resource{
				Provider:   "Kubernetes",
				Account:    client.Name,
				Service:    "Ingress",
				ResourceId: string(ingress.UID),
				Name:       ingress.Name,
				Region:     ingress.Namespace,
				Cost:       0,
				CreatedAt:  ingress.CreationTimestamp.Time,
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
		"service":   "Ingress",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
