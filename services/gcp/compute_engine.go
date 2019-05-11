package gcp

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	. "github.com/mlabouardy/komiser/models/gcp"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
	monitoring "google.golang.org/api/monitoring/v3"
)

func (gcp GCP) GetComputeInstances() ([]Instance, error) {
	instancesList := make([]Instance, 0)

	src, err := google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err != nil {
		return instancesList, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := compute.New(client)
	if err != nil {
		return instancesList, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return instancesList, err
	}

	for _, project := range projects {
		zones, err := svc.Zones.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return instancesList, err
		}

		for _, zone := range zones.Items {
			instances, err := svc.Instances.List(project.ID, zone.Name).Do()
			if err != nil {
				log.Println(err)
				return instancesList, err
			}

			public := true

			for _, instance := range instances.Items {
				for _, network := range instance.NetworkInterfaces {
					for _, config := range network.AccessConfigs {
						if config.NatIP == "" {
							public = false
						}
					}
				}

				parts := strings.Split(instance.MachineType, "/")
				machineType := parts[len(parts)-1]

				instancesList = append(instancesList, Instance{
					Name:        instance.Name,
					MachineType: machineType,
					Status:      instance.Status,
					CPUPlatform: instance.CpuPlatform,
					Public:      public,
					Zone:        zone.Name,
				})
			}
		}
	}
	return instancesList, nil
}

func (gcp GCP) GetRegions() ([]string, error) {
	listOfRegions := make([]string, 0)

	src, err := google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err != nil {
		return listOfRegions, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := compute.New(client)
	if err != nil {
		return listOfRegions, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return listOfRegions, err
	}

	for _, project := range projects {
		regions, err := svc.Regions.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return listOfRegions, err
		}

		for _, region := range regions.Items {
			listOfRegions = append(listOfRegions, region.Name)
		}
	}
	return listOfRegions, nil
}

func (gcp GCP) GetDisks() ([]Disk, error) {
	listOfDisks := make([]Disk, 0)

	src, err := google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err != nil {
		return listOfDisks, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := compute.New(client)
	if err != nil {
		return listOfDisks, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return listOfDisks, err
	}

	for _, project := range projects {
		zones, err := svc.Zones.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return listOfDisks, err
		}

		for _, zone := range zones.Items {
			disks, err := svc.Disks.List(project.ID, zone.Name).Do()
			if err != nil {
				log.Println(err)
				return listOfDisks, err
			}

			for _, disk := range disks.Items {
				listOfDisks = append(listOfDisks, Disk{
					SizeGb: disk.SizeGb,
					Status: disk.Status,
				})
			}
		}
	}
	return listOfDisks, nil
}

func (gcp GCP) GetDiskSnapshots() ([]Snapshot, error) {
	listOfSnapshots := make([]Snapshot, 0)

	src, err := google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err != nil {
		return listOfSnapshots, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := compute.New(client)
	if err != nil {
		return listOfSnapshots, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return listOfSnapshots, err
	}

	for _, project := range projects {
		snapshots, err := svc.Snapshots.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return listOfSnapshots, err
		}

		for _, snapshot := range snapshots.Items {
			listOfSnapshots = append(listOfSnapshots, Snapshot{
				SizeGb: snapshot.DiskSizeGb,
			})
		}
	}
	return listOfSnapshots, nil
}

func (gcp GCP) GetComputeImages() ([]Image, error) {
	listImages := make([]Image, 0)

	src, err := google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err != nil {
		return listImages, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := compute.New(client)
	if err != nil {
		return listImages, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return listImages, err
	}

	for _, project := range projects {
		images, err := svc.Images.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return listImages, err
		}

		for _, image := range images.Items {
			listImages = append(listImages, Image{
				SizeGb: image.DiskSizeGb,
			})
		}
	}
	return listImages, nil
}

func (gcp GCP) GetComputeCPUUtilization() ([]*monitoring.TimeSeries, error) {
	src, err := google.DefaultTokenSource(oauth2.NoContext, monitoring.MonitoringReadScope)
	if err != nil {
		return []*monitoring.TimeSeries{}, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := monitoring.New(client)
	if err != nil {
		return []*monitoring.TimeSeries{}, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return []*monitoring.TimeSeries{}, err
	}

	uri := fmt.Sprintf("projects/%s", projects[0].ID)
	data, err := svc.Projects.TimeSeries.
		List(uri).
		AggregationAlignmentPeriod("86400s").
		AggregationGroupByFields("project", "resource.labels.instance_name").
		AggregationPerSeriesAligner("ALIGN_MEAN").
		Filter(`metric.type="compute.googleapis.com/instance/cpu/utilization"`).
		IntervalEndTime(time.Now().Format("2006-01-02T15:04:05.000Z")).
		IntervalStartTime(time.Now().AddDate(0, 0, -7).Format("2006-01-02T15:04:05.000Z")).
		Do()

	if err != nil {
		log.Println(err)
		return []*monitoring.TimeSeries{}, err
	}

	return data.TimeSeries, nil
}

func (gcp GCP) GetQuotas() ([]Quota, error) {
	limits := make([]Quota, 0)

	src, err := google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err != nil {
		return limits, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := compute.New(client)
	if err != nil {
		return limits, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return limits, err
	}

	for _, project := range projects {
		quotas, err := svc.Projects.Get(project.ID).Do()
		if err != nil {
			log.Println(err)
			return limits, err
		}

		for _, quota := range quotas.Quotas {
			limits = append(limits, Quota{
				Metric: quota.Metric,
				Usage:  quota.Usage,
				Limit:  quota.Limit,
			})
		}
	}
	return limits, nil
}
