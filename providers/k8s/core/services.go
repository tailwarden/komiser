package core

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	. "github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	oc "github.com/tailwarden/komiser/providers/k8s/opencost"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Services(ctx context.Context, client providers.ProviderClient) ([]Resource, error) {
	resources := make([]Resource, 0)

	var config metav1.ListOptions

	opencostEnabled := true
	serviceCost, err := oc.GetOpencostInfo(client.K8sClient.OpencostBaseUrl, "service")
	if err != nil {
		log.Errorf("ERROR: Couldn't get service info from OpenCost: %v", err)
		log.Warn("Opencost disabled")
		opencostEnabled = false
	}

	for {
		res, err := client.K8sClient.Client.CoreV1().Services("").List(ctx, config)
		if err != nil {
			return nil, err
		}

		for _, service := range res.Items {
			tags := make([]Tag, 0)

			for key, value := range service.Labels {
				tags = append(tags, Tag{
					Key:   key,
					Value: value,
				})
			}

			cost := 0.0
			if opencostEnabled {
				cost = serviceCost[service.Name].TotalCost
			}

			resources = append(resources, Resource{
				Provider:   "Kubernetes",
				Account:    client.Name,
				Service:    "Service",
				ResourceId: string(service.UID),
				Name:       service.Name,
				Region:     service.Namespace,
				Cost:       cost,
				CreatedAt:  service.CreationTimestamp.Time,
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
		"service":   "Service",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
