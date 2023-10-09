package gcpcomputepricing

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	compute "cloud.google.com/go/compute/apiv1"
	"cloud.google.com/go/compute/apiv1/computepb"
	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/utils"
	"google.golang.org/api/option"
)

// CalculateDiskCost returns a calculated and normalized disk cost by hours in this month.
func CalculateDiskCost(ctx context.Context, client providers.ProviderClient, data CalculateDiskCostData) (float64, error) {
	diskTypeClient, err := compute.NewDiskTypesRESTClient(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		return 0, err
	}

	dtS := strings.Split(data.DiskType, "/")
	dt, err := diskTypeClient.Get(ctx, &computepb.GetDiskTypeRequest{
		DiskType: dtS[len(dtS)-1],
		Project:  data.Project,
		Zone:     data.Zone,
	})
	if err != nil {
		return 0, err
	}

	var opts = Opts{
		Region:   utils.GcpGetRegionFromZone(data.Zone),
		DiskSize: uint64(data.Size),
	}
	// Populating the disk type field with supported types
	if dt.Name != nil {
		switch {
		case nameToTypeMatch(*dt.Name, Standard):
			opts.Type = Standard
		case nameToTypeMatch(*dt.Name, SSD):
			opts.Type = SSD
		case nameToTypeMatch(*dt.Name, Balanced):
			opts.Type = Balanced
		case nameToTypeMatch(*dt.Name, Extreme):
			opts.Type = Extreme
		default:
			// Early return if disk type unsupported and do not need calculate cost
			return 0, errors.New("unsupported disk type")
		}
	}

	var cost float64
	monthlyRate, err := getDiskMonthly(data.Pricing, opts, typeDiskGet)
	if err != nil {
		return 0, err
	}

	currentTime := time.Now()
	startOfMonth := utils.BeginningOfMonth(currentTime)
	endOfMonth := utils.EndingOfMonth(currentTime)

	created, err := time.Parse(time.RFC3339, data.CreationTimestamp)
	if err != nil {
		return 0, err
	}

	duration := currentTime.Sub(startOfMonth)
	if created.After(startOfMonth) {
		duration = currentTime.Sub(created)
	}

	hourlyRate := monthlyRate / uint64(endOfMonth.Sub(startOfMonth).Hours())
	normalizedHourlyRate := float64(hourlyRate) / 1000000000
	cost = normalizedHourlyRate * float64(duration.Hours())

	return cost, nil
}

func typeDiskGet(p *Pricing, opts Opts) (Subtype, error) {
	var capacity Subtype
	switch opts.Type {
	case Standard:
		capacity = p.Gcp.Compute.PersistentDisk.Standard.Capacity.Storagepdcapacity
	case SSD:
		capacity = p.Gcp.Compute.PersistentDisk.SSD.Capacity.Storagepdssd
	case Balanced:
		capacity = p.Gcp.Compute.PersistentDisk.SSD.Capacity.Lite.Storagepdssdlitecapacity
	case Extreme:
		capacity = p.Gcp.Compute.PersistentDisk.SSD.Capacity.Extreme.Storagepdssdextremecapacity
	}

	return capacity, nil
}

// getDiskMonthly returs a calculated price by region for month.
func getDiskMonthly(p *Pricing, opts Opts, tg typeDiskGetter) (uint64, error) {
	capacity, err := tg(p, opts)
	if err != nil {
		return 0, err
	}

	var capacityPricePerRegion uint64 = 0
	if region, ok := capacity.Regions[opts.Region]; ok {
		if len(region.Prices) > 0 {
			capacityPricePerRegion = region.Prices[0].Nanos
		}
	} else {
		return 0, fmt.Errorf("capacity price not found for %q region", opts.Region)
	}

	var sum uint64 = 0
	sum += capacityPricePerRegion * opts.DiskSize

	return sum, nil
}
