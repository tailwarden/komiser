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

func Pods(ctx context.Context, client providers.ProviderClient, namespace string) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	var config metav1.ListOptions

	opencostEnabled := true
	podsCost, err := oc.GetOpencostInfo(client.K8sClient.OpencostBaseUrl, "pod")
	if err != nil {
		log.Errorf("ERROR: Couldn't get pods info from OpenCost: %v", err)
		log.Warn("Opencost disabled")
		opencostEnabled = false
	}

	config.Namespace = namespace // Set the namespace to filter pods by

	for {
		res, err := client.K8sClient.Client.CoreV1().Pods(namespace).List(ctx, config) // Filter pods by namespace
		if err != nil {
			return nil, err
		}

		for _, pod := range res.Items {
			tags := make([]models.Tag, 0)

			for key, value := range pod.Labels {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: value,
				})
			}

			cost := 0.0
			if opencostEnabled {
				cost = podsCost[pod.Name].TotalCost
			}

			resources = append(resources, models.Resource{
				Provider:   "Kubernetes",
				Account:    client.Name,
				Service:    "Pod",
				ResourceId: string(pod.UID),
				Name:       pod.Name,
				Region:     pod.Namespace, // Use the pod's namespace as the region
				Cost:       cost,
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
