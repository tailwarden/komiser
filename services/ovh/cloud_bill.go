package ovh

import (
	"fmt"

	ovhClient "github.com/ovh/go-ovh/ovh"
)

type Usage struct {
	MonthlyUsage struct {
		Instance []struct {
			TotalPrice float32 `json:"totalPrice"`
			Region     string  `json:"region"`
		} `json:"instance"`
	} `json:"monthlyUsage"`

	HourlyUsage struct {
		Instance []struct {
			TotalPrice float32 `json:"totalPrice"`
			Region     string  `json:"region"`
		} `json:"instance"`
		Snapshot []struct {
			TotalPrice float32 `json:"totalPrice"`
			Region     string  `json:"region"`
		} `json:"snapshot"`
		Storage []struct {
			TotalPrice float32 `json:"totalPrice"`
			Region     string  `json:"region"`
		} `json:"storage"`
		InstanceBandwidth []struct {
			TotalPrice float32 `json:"totalPrice"`
			Region     string  `json:"region"`
		} `json:"instanceBandwidth"`
		Volume []struct {
			TotalPrice float32 `json:"totalPrice"`
			Region     string  `json:"region"`
		} `json:"volume"`
		InstanceOptions []struct {
			TotalPrice float32 `json:"totalPrice"`
			Region     string  `json:"region"`
		} `json:"instanceOption"`
	} `json:"hourlyUsage"`
}

type Bill struct {
	Total    float32   `json:"total"`
	Services []Service `json:"services"`
}

type Service struct {
	Label string  `json:"label"`
	Total float32 `json:"total"`
}

func (ovh OVH) GetCurrentUsage() (Bill, error) {
	bill := Bill{}

	client, err := ovhClient.NewDefaultClient()
	if err != nil {
		return bill, err
	}

	projects, err := ovh.GetProjects()
	if err != nil {
		return bill, err
	}

	for _, project := range projects {
		usage := Usage{}
		err = client.Get(fmt.Sprintf("/cloud/project/%s/usage/current", project.ID), &usage)
		if err != nil {
			return bill, err
		}

		bill.Services = make([]Service, 6)

		bill.Services[0].Label = "Compute Instances"
		for _, instance := range usage.HourlyUsage.Instance {
			bill.Services[0].Total += instance.TotalPrice
			bill.Total += instance.TotalPrice
		}

		bill.Services[1].Label = "Snapshots"
		for _, snapshot := range usage.HourlyUsage.Snapshot {
			bill.Services[1].Total += snapshot.TotalPrice
			bill.Total += snapshot.TotalPrice
		}

		bill.Services[2].Label = "Storage"
		for _, storage := range usage.HourlyUsage.Storage {
			bill.Services[2].Total += storage.TotalPrice
			bill.Total += storage.TotalPrice
		}

		bill.Services[3].Label = "Consumed Bandwidth"
		for _, bandwidth := range usage.HourlyUsage.InstanceBandwidth {
			bill.Services[3].Total += bandwidth.TotalPrice
			bill.Total += bandwidth.TotalPrice
		}

		bill.Services[4].Label = "Volumes"
		for _, volume := range usage.HourlyUsage.Volume {
			bill.Services[4].Total += volume.TotalPrice
			bill.Total += volume.TotalPrice
		}

		bill.Services[5].Label = "Instance Options"
		for _, options := range usage.HourlyUsage.InstanceOptions {
			bill.Services[5].Total += options.TotalPrice
			bill.Total += options.TotalPrice
		}
	}

	return bill, nil
}
