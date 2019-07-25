package digitalocean

import (
	"context"
	"time"

	"github.com/digitalocean/godo"
)

type Action struct {
	At       time.Time `json:"at"`
	Type     string    `json:"type"`
	Status   string    `json:"status"`
	Resource string    `json:"resource"`
	Region   string    `json:"region"`
}

func (dg DigitalOcean) DescribeActions(client *godo.Client) ([]Action, error) {
	listOfActions := make([]Action, 0)

	actions, _, err := client.Actions.List(context.TODO(), &godo.ListOptions{})
	if err != nil {
		return listOfActions, err
	}

	for _, action := range actions {
		listOfActions = append(listOfActions, Action{
			At:       action.StartedAt.Time,
			Region:   action.RegionSlug,
			Resource: action.ResourceType,
			Type:     action.Type,
			Status:   action.Status,
		})
	}

	return listOfActions, nil
}
