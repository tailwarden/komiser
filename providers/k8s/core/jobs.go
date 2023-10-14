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

func Jobs(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	var config metav1.ListOptions

	opencostEnabled := true
	jobsCost, err := oc.GetOpencostInfo(client.K8sClient.OpencostBaseUrl, "job")
	if err != nil {
		log.Errorf("ERROR: Couldn't get jobs info from OpenCost: %v", err)
		log.Warn("Opencost disabled")
		opencostEnabled = false
	}

	for {
		res, err := client.K8sClient.Client.BatchV1().Jobs("").List(ctx, config)
		if err != nil {
			return nil, err
		}

		for _, job := range res.Items {
			tags := make([]models.Tag, 0)

			for key, value := range job.Labels {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: value,
				})
			}

			if len(job.OwnerReferences) > 0 {
				// we use the owner kind of first owner only as the owner tag
				ownerTags := []models.Tag{
					{
						Key:   "owner_kind",
						Value: job.OwnerReferences[0].Kind,
					},
					{
						Key:   "owner_name",
						Value: job.OwnerReferences[0].Name,
					},
				}
				tags = append(tags, ownerTags...)
			}

			cost := 0.0
			if opencostEnabled {
				cost = jobsCost[job.Name].TotalCost
			}

			resources = append(resources, models.Resource{
				Provider:   "Kubernetes",
				Account:    client.Name,
				Service:    "Job",
				ResourceId: string(job.UID),
				Name:       job.Name,
				Region:     job.Namespace,
				Cost:       cost,
				CreatedAt:  job.CreationTimestamp.Time,
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
		"service":   "Job",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
