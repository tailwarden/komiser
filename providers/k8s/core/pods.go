package core

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	. "github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Pods(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)

	var config metav1.ListOptions

	for {
		res, err := client.K8sClient.CoreV1().Pods("").List(ctx, config)
		if err != nil {
			return nil, err
		}

		for _, pod := range res.Items {
			tags := make([]Tag, 0)

			for key, value := range pod.Labels {
				tags = append(tags, Tag{
					Key:   key,
					Value: value,
				})
			}

			resources = append(resources, Resource{
				Provider:   "Kubernetes",
				Account:    client.Name,
				Service:    "Pod",
				ResourceId: string(pod.UID),
				Name:       pod.Name,
				Region:     pod.Namespace,
				Cost:       0,
				CreatedAt:  pod.CreationTimestamp.Time,
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
		"service":   "Pod",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
