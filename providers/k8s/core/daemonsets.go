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

func DaemonSets(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	var config metav1.ListOptions

	opencostEnabled := true
	daemonsetsCost, err := oc.GetOpencostInfo(client.K8sClient.OpencostBaseUrl, "daemonset")
	if err != nil {
		log.Errorf("ERROR: Couldn't get daemonsets info from OpenCost: %v", err)
		log.Warn("Opencost disabled")
		opencostEnabled = false
	}

	for {
		res, err := client.K8sClient.Client.AppsV1().DaemonSets("").List(ctx, config)
		if err != nil {
			return nil, err
		}

		for _, daemonset := range res.Items {
			tags := make([]models.Tag, 0)

			for key, value := range daemonset.Labels {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: value,
				})
			}

			cost := 0.0
			if opencostEnabled {
				cost = daemonsetsCost[daemonset.Name].TotalCost
			}

			resources = append(resources, models.Resource{
				Provider:   "Kubernetes",
				Account:    client.Name,
				Service:    "DaemonSet",
				ResourceId: string(daemonset.UID),
				Name:       daemonset.Name,
				Region:     daemonset.Namespace,
				Cost:       cost,
				CreatedAt:  daemonset.CreationTimestamp.Time,
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
		"service":   "DaemonSet",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
