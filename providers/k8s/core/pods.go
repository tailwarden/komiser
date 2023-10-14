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

func Pods(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	var config metav1.ListOptions

	opencostEnabled := true
	podsCost, err := oc.GetOpencostInfo(client.K8sClient.OpencostBaseUrl, "pod")
	if err != nil {
		log.Errorf("ERROR: Couldn't get pods info from OpenCost: %v", err)
		log.Warn("Opencost disabled")
		opencostEnabled = false
	}

	for {
		res, err := client.K8sClient.Client.CoreV1().Pods("").List(ctx, config)
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

			if len(pod.OwnerReferences) > 0 {
				// we use the owner kind of first owner only as the owner tag
				ownerTags := []models.Tag{
					{
						Key:   "owner_kind",
						Value: pod.OwnerReferences[0].Kind,
					},
					{
						Key:   "owner_name",
						Value: pod.OwnerReferences[0].Name,
					},
				}
				tags = append(tags, ownerTags...)
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
				Region:     pod.Namespace,
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
