package gcpcomputepricing

import (
	"context"
	"fmt"
	"time"

	"github.com/tailwarden/komiser/providers"
	"github.com/tailwarden/komiser/utils"
)

func CalculateSnapshotCost(ctx context.Context, client providers.ProviderClient, data CalculateSnapshotCostData) (float64, error) {
	var opts = Opts{
		Region: data.Region,
	}

	var cost float64
	monthlyRate, err := getSnapshotMonthly(
		data.Pricing,
		opts,
		typeSnapshotGetter,
		uint64(data.StorageBytes))
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

func typeSnapshotGetter(p *Pricing, opts Opts) (Subtype, error) {
	// TODO: switch by snapshot type
	capacity := p.Gcp.Compute.PersistentDisk.Snapshots.Storageregionalstandardsnapshotearlydeletion

	return capacity, nil
}

func getSnapshotMonthly(p *Pricing, opts Opts, tg func(*Pricing, Opts) (Subtype, error), storageBytes uint64) (uint64, error) {
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

	normalizedSize := uint64(storageBytes) / 1024 / 1024 / 1024
	return capacityPricePerRegion * normalizedSize, nil
}
