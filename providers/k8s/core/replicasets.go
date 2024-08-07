package core

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/api/apps/v1"
)

func Replicasets(ctx context.Context, client providers.ProviderClient) ([]models.Resource, error) {
	resources := make([]models.Resource, 0)

	var config metav1.ListOptions

	for {
		res, err := client.K8sClient.Client.AppsV1().ReplicaSets("").List(ctx, config)
		if err != nil {
			return nil, err
		}

		for _, rs := range res.Items {
			tags := make([]models.Tag, 0)

			for key, value := range rs.Labels {
				tags = append(tags, models.Tag{
					Key:   key,
					Value: value,
				})
			}

			if len(rs.OwnerReferences) > 0 {
				// we use the owner kind of first owner only as the owner tag
				ownerTags := []models.Tag{
					{
						Key:   "owner_kind",
						Value: rs.OwnerReferences[0].Kind,
					},
					{
						Key:   "owner_name",
						Value: rs.OwnerReferences[0].Name,
					},
				}
				tags = append(tags, ownerTags...)
			}


			resources = append(resources, models.Resource{
				Provider:   "Kubernetes",
				Account:    client.Name,
				Service:    "Replicaset",
				ResourceId: string(rs.UID),
				Name:       rs.Name,
				Region:     rs.Namespace,
				Relations: getReplicasetRelation(rs),
				Cost:       0,
				CreatedAt:  rs.CreationTimestamp.Time,
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
		"service":   "Replicaset",
		"resources": len(resources),
	}).Info("Fetched resources")
	return resources, nil
}


func getReplicasetRelation(rs v1.ReplicaSet) []models.Link {

	var rel []models.Link

	owners := rs.GetOwnerReferences()
	for _, owner := range owners {
		rel = append(rel, models.Link{
			ResourceID: string(owner.UID),
			Type:       owner.Kind,
			Name:       owner.Name,
			Relation:   "USES",
		})
	}

	return rel
}