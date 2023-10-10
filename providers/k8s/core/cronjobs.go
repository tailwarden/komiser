package core

import (
	"context"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	oc "github.com/tailwarden/komiser/providers/k8s/opencost"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CronJobs(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	var config metav1.ListOptions

	opencostEnabled := true
	cronjobsCost, err := oc.GetOpencostInfo(client.K8sClient.OpencostBaseUrl, "cronjob")
	if err != nil {
		log.Errorf("ERROR: Couldn't get cronjobs info from OpenCost: %v", err)
		log.Warn("Opencost disabled")
		opencostEnabled = false
	}

	for {
		res, err := client.K8sClient.Client.BatchV1().CronJobs("").List(ctx, config)
		if err != nil {
			return nil, err
		}

		for _, cronjob := range res.Items {
			tags := make([]models.Tag, 0)

			for key, value := range cronjob.Labels {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: value,
				})
			}

			cost := 0.0
			if opencostEnabled {
				cost = cronjobsCost[cronjob.Name].TotalCost
			}

			resources = append(resources, models.Resource{
				Provider:   "Kubernetes",
				Account:    client.Name,
				Service:    "CronJob",
				ResourceId: string(cronjob.UID),
				Name:       cronjob.Name,
				Region:     cronjob.Namespace,
				Cost:       cost,
				CreatedAt:  cronjob.CreationTimestamp.Time,
				FetchedAt:  time.Now(),
				Tags:       tags,
			})
			fmt.Printf("%+v", cronjob)
		}

		if res.GetContinue() == "" {
			break
		}

		config.Continue = res.GetContinue()

	}

	log.WithFields(log.Fields{
		"provider":  "Kubernetes",
		"account":   client.Name,
		"service":   "CronJob",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}
