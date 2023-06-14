package gcpcomputepricing

import (
	"errors"
	"fmt"
)

const (
	Standard = "pd-standard"
	SSD      = "pd-ssd"
	Balanced = "pd-balanced"
	Extreme  = "pd-extreme"
)

func CalculateDisk(p *Pricing, opts Opts) (uint64, error) {
	return getDiskMonthly(p, opts, typeDiskGetter)
}

func typeDiskGetter(p *Pricing, opts Opts) (Subtype, Subtype, error) {
	var capacity Subtype
	var snapshot Subtype
	switch opts.DiskType {
	case Standard:
		capacity = p.Gcp.Compute.PersistentDisk.Standard.Capacity.Storagepdcapacity
		snapshot = p.Gcp.Compute.PersistentDisk.Standard.Snapshot.Storagepdsnapshot
	case SSD:
		capacity = p.Gcp.Compute.PersistentDisk.SSD.Capacity.Storagepdssd
	case Balanced:
		capacity = p.Gcp.Compute.PersistentDisk.AsyncReplication.BalancedProtection.Asyncreplicationprotectionpdbalanced
		snapshot = p.Gcp.Compute.PersistentDisk.Snapshots.Storagemultiregionalstandardsnapshotearlydeletion
	}

	return capacity, snapshot, nil
}

func getDiskMonthly(p *Pricing, opts Opts, tg typeGetter) (uint64, error) {
	capacity, snapshot, err := tg(p, opts)
	if err != nil {
		return 0, err
	}

	var capacityPricePerRegion uint64 = 0
	if region, ok := capacity.Regions[opts.Region]; ok {
		if len(region.Prices) > 0 {
			capacityPricePerRegion = region.Prices[0].Nanos
		}
	} else {
		return 0, errors.New(fmt.Sprintf("capacity price not found for %q region", opts.Region))
	}

	var snapshotPricePerRegion uint64 = 0
	// SSD disks do not support snapshots
	if opts.DiskType != SSD {
		if region, ok := snapshot.Regions[opts.Region]; ok {
			if len(region.Prices) > 0 {
				snapshotPricePerRegion = region.Prices[0].Nanos
			}
		} else {
			return 0, errors.New(fmt.Sprintf("snapshot price not found for %q region", opts.Region))
		}
	}

	var sum uint64 = 0
	sum += capacityPricePerRegion * opts.DiskSize
	sum += snapshotPricePerRegion * opts.DiskSize

	return sum, nil
}
