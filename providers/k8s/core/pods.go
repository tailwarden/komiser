package core

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/models"
	"github.com/tailwarden/komiser/providers"
	oc "github.com/tailwarden/komiser/providers/k8s/opencost"
	"github.com/tailwarden/komiser/utils"
	v1 "k8s.io/api/core/v1"
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
				Relations:  getPodRelation(pod, client),
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

func getPodRelation(pod v1.Pod, client providers.ProviderClient) []models.Link {

	var rel []models.Link
	var err error
	owners := pod.GetOwnerReferences()
	for _, owner := range owners {
		rel = append(rel, models.Link{
			ResourceID: string(owner.UID),
			Type:       owner.Kind,
			Name:       owner.Name,
			Relation:   "USES",
		})
	}

	// check if service was cached before
	// if not call the service cache function, get list of services and cahce it for further use
	services, ok := client.Cache.Get(utils.SERVICES)
	if !ok {
		services, err = utils.K8s_Cache(&client, utils.SERVICES)
	}

	if err == nil {
		serviceList, _ := services.(*v1.ServiceList)

		for _, service := range serviceList.Items {
			selector := service.Spec.Selector
			if selectorMatchesPodLabels(selector, pod.Labels) {
				rel = append(rel, models.Link{
					ResourceID: string(service.UID),
					Type:       "Service",
					Name:       service.Name,
					Relation:   "USES",
				})
			}
		}
	}

	// for nodes
	node, err := client.K8sClient.Client.CoreV1().Nodes().Get(context.TODO(), pod.Spec.NodeName, metav1.GetOptions{})
	if err == nil {
		rel = append(rel, models.Link{
			ResourceID: string(node.UID),
			Type:       "Node",
			Name:       node.Name,
			Relation:   "USES",
		})
	}

	// for namespace
	namespace, err := client.K8sClient.Client.CoreV1().Namespaces().Get(context.TODO(), pod.Namespace, metav1.GetOptions{})
	if err == nil {
		rel = append(rel, models.Link{
			ResourceID: string(namespace.UID),
			Type:       "Namespace",
			Name:       namespace.Name,
			Relation:   "USES",
		})
	}

	// for serviceAccount
	sa, err := client.K8sClient.Client.CoreV1().ServiceAccounts(namespace.Name).Get(context.TODO(), pod.Spec.ServiceAccountName, metav1.GetOptions{})
	if err == nil {
		rel = append(rel, models.Link{
			ResourceID: string(sa.UID),
			Type:       "ServiceAccount",
			Name:       sa.Name,
			Relation:   "USES",
		})
	}

	return rel
}

func selectorMatchesPodLabels(selector, podLabels map[string]string) bool {
	if len(selector) == 0 {
		return false
	}

	for key, value := range selector {
		if podValue, exists := podLabels[key]; !exists || podValue != value {
			return false
		}
	}
	return true
}
