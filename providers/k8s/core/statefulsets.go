package core

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	oc "github.com/tailwarden/komiser/providers/k8s/opencost"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func StatefulSets(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	var config metav1.ListOptions

	opencostEnabled := true
	statefulsetsCost, err := oc.GetOpencostInfo(client.K8sClient.OpencostBaseUrl, "statefulset")
	if err != nil {
		log.Errorf("ERROR: Couldn't get statefulsets info from OpenCost: %v", err)
		log.Warn("Opencost disabled")
		opencostEnabled = false
	}

	for {
		res, err := client.K8sClient.Client.AppsV1().StatefulSets("").List(ctx, config)
		if err != nil {
			return nil, err
		}

		for _, statefulset := range res.Items {
			tags := make([]models.Tag, 0)

			for key, value := range statefulset.Labels {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: value,
				})
			}

			cost := 0.0
			if opencostEnabled {
				cost = statefulsetsCost[statefulset.Name].TotalCost
			}

			resources = append(resources, models.Resource{
				Provider:   "Kubernetes",
				Account:    client.Name,
				Service:    "StatefulSet",
				ResourceId: string(statefulset.UID),
				Name:       statefulset.Name,
				Region:     statefulset.Namespace,
				Cost:       cost,
				CreatedAt:  statefulset.CreationTimestamp.Time,
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
		"service":   "StatefulSet",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
