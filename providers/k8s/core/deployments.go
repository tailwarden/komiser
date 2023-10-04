package core

import (
	"context"
	"time"

	log "github.com/siruspen/logrus"
	. "github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Deployments(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)

	var config metav1.ListOptions

	for {
		res, err := client.K8sClient.Client.AppsV1().Deployments("").List(ctx, config)
		if err != nil {
			return nil, err
		}

		for _, deploy := range res.Items {
			tags := make([]Tag, 0)

			for key, value := range deploy.Labels {
				tags = append(tags, Tag{
					Key:   key,
					Value: value,
				})
			}

			resources = append(resources, Resource{
				Provider:   "Kubernetes",
				Account:    client.Name,
				Service:    "Deployment",
				ResourceId: string(deploy.UID),
				Name:       deploy.Name,
				Region:     deploy.Namespace,
				Cost:       0,
				CreatedAt:  deploy.CreationTimestamp.Time,
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
		"service":   "Deployment",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
