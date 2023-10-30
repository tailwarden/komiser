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

// CalculateMachineCost returns a calculated and normalized machine cost by hours in this month.
func CalculateMachineCost(ctx context.Context, client providers.ProviderClient, data CalculateMachineCostData) (float64, error) {
	machineTypeClient, err := compute.NewMachineTypesRESTClient(ctx, option.WithCredentials(client.GCPClient.Credentials))
	if err != nil {
		return 0, err
	}

	mtS := strings.Split(data.MachineType, "/")

	mt, err := machineTypeClient.Get(ctx, &computepb.GetMachineTypeRequest{
		MachineType: mtS[len(mtS)-1],
		Project:     data.Project,
		Zone:        data.Zone,
	})
	if err != nil {
		return 0, err
	}

	var opts = Opts{
		Commitment:  data.Commitment,
		Region:      utils.GcpGetRegionFromZone(data.Zone),
		NumOfCPU:    uint64(*mt.GuestCpus),
		NumOfMemory: uint64(*mt.MemoryMb / 1024),
	}
	if mt.Name != nil {
		switch {
		case nameToTypeMatch(*mt.Name, E2):
			opts.Type = E2
		case nameToTypeMatch(*mt.Name, C3):
			opts.Type = C3
		case nameToTypeMatch(*mt.Name, N2):
			opts.Type = N2
		case nameToTypeMatch(*mt.Name, N2D):
			opts.Type = N2D
		case nameToTypeMatch(*mt.Name, T2A):
			opts.Type = T2A
		case nameToTypeMatch(*mt.Name, T2D):
			opts.Type = T2D
		case nameToTypeMatch(*mt.Name, N1):
			opts.Type = N1
		case nameToTypeMatch(*mt.Name, C2):
			opts.Type = C2
		case nameToTypeMatch(*mt.Name, C2D):
			opts.Type = C2D
		case nameToTypeMatch(*mt.Name, M1):
			opts.Type = M1
		case nameToTypeMatch(*mt.Name, M2):
			opts.Type = M2
		case nameToTypeMatch(*mt.Name, M3):
			opts.Type = M3
		default:
			// Early return if machine type unsupported and do not need calculate cost
			return 0, errors.New("unsupported machine type")
		}
	}

	var cost float64
	hourlyRate, err := calculateMachineHourly(data.Pricing, opts)
	if err != nil {
		return 0, err
	}

	currentTime := time.Now()
	startOfMonth := utils.BeginningOfMonth(currentTime)

	created, err := time.Parse(time.RFC3339, data.CreationTimestamp)
	if err != nil {
		return 0, err
	}

	duration := currentTime.Sub(startOfMonth)
	if created.After(startOfMonth) {
		duration = currentTime.Sub(created)
	}

	normalizedHourlyRate := float64(hourlyRate) / 1000000000
	cost = normalizedHourlyRate * float64(duration.Hours())

	return cost, nil
}

func calculateMachineHourly(p *Pricing, opts Opts) (uint64, error) {
	switch opts.Type {
	case E2:
		return getHourly(p, opts, typeGetterE2)
	case C3:
		return getHourly(p, opts, typeGetterC3)
	case N2:
		return getHourly(p, opts, typeGetterN2)
	case N2D:
		return getHourly(p, opts, typeGetterN2D)
	case T2A:
		return getHourly(p, opts, typeGetterT2A)
	case T2D:
		return getHourly(p, opts, typeGetterT2D)
	case N1:
		return getHourly(p, opts, typeGetterN1)
	case C2D:
		return getHourly(p, opts, typeGetterC2D)
	case M3:
		return getHourly(p, opts, typeGetterM3)
	}
	return 0, errors.New("unknown type")
}

func typeGetterE2(p *Pricing, opts Opts) (Subtype, Subtype, error) {
	var core Subtype
	var memory Subtype
	switch opts.Commitment {
	case OnDemand:
		core = p.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.E2.Vmimagee2Core
		memory = p.Gcp.Compute.GCE.VmsOnDemand.MemoryPerGb.E2.Vmimagee2RAM
	case Spot:
		core = p.Gcp.Compute.GCE.VmsPreemptible.CoresPerCore.E2.Vmimagepreemptiblee2Core
		memory = p.Gcp.Compute.GCE.VmsPreemptible.MemoryPerGb.E2.Vmimagepreemptiblee2RAM
	case Commitment1YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit1Year.CoresPerCore.E2.Commitmente2CPU1Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit1Year.MemoryPerGb.E2.Commitmente2RAM1Yv1
	case Commitment3YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit3Year.CoresPerCore.E2.Commitmente2CPU3Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit3Year.MemoryPerGb.E2.Commitmente2RAM3Yv1
	default:
		return Subtype{}, Subtype{}, fmt.Errorf("commitment %q not supported", opts.Commitment)
	}
	return core, memory, nil
}

func typeGetterC3(p *Pricing, opts Opts) (Subtype, Subtype, error) {
	var core Subtype
	var memory Subtype
	switch opts.Commitment {
	// TODO there are multiple, sole tenancy, premium and standard
	case OnDemand:
		core = p.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.C3.Vmimagec3Standardcore
		memory = p.Gcp.Compute.GCE.VmsOnDemand.MemoryPerGb.C3.Vmimagec3Standardram
	case Spot:
		core = p.Gcp.Compute.GCE.VmsPreemptible.CoresPerCore.C3.Vmimagepreemptiblec3Standardcore
		memory = p.Gcp.Compute.GCE.VmsPreemptible.MemoryPerGb.C3.Vmimagepreemptiblec3Standardram
	case Commitment1YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit1Year.CoresPerCore.C3.Commitmentc3CPU1Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit1Year.MemoryPerGb.C3.Commitmentc3RAM1Yv1
	case Commitment3YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit3Year.CoresPerCore.C3.Commitmentc3CPU3Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit3Year.MemoryPerGb.C3.Commitmentc3RAM3Yv1
	default:
		return Subtype{}, Subtype{}, fmt.Errorf("commitment %q not supported", opts.Commitment)
	}
	return core, memory, nil
}

func typeGetterN2(p *Pricing, opts Opts) (Subtype, Subtype, error) {
	var core Subtype
	var memory Subtype
	switch opts.Commitment {
	case OnDemand:
		core = p.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.N2.Vmimagen2Standardcore
		memory = p.Gcp.Compute.GCE.VmsOnDemand.MemoryPerGb.N2.Vmimagen2Standardram
	case Spot:
		core = p.Gcp.Compute.GCE.VmsPreemptible.CoresPerCore.N2.Vmimagepreemptiblen2Standardcore
		memory = p.Gcp.Compute.GCE.VmsPreemptible.MemoryPerGb.N2.Vmimagepreemptiblen2Standardram
	case Commitment1YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit1Year.CoresPerCore.N2.Commitmentn2CPU1Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit1Year.MemoryPerGb.N2.Commitmentn2RAM1Yv1
	case Commitment3YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit3Year.CoresPerCore.N2.Commitmentn2CPU3Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit3Year.MemoryPerGb.N2.Commitmentn2RAM3Yv1
	default:
		return Subtype{}, Subtype{}, fmt.Errorf("commitment %q not supported", opts.Commitment)
	}
	return core, memory, nil
}

func typeGetterN2D(p *Pricing, opts Opts) (Subtype, Subtype, error) {
	var core Subtype
	var memory Subtype
	switch opts.Commitment {
	case OnDemand:
		core = p.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.N2D.Vmimagen2Dstandardcore
		memory = p.Gcp.Compute.GCE.VmsOnDemand.MemoryPerGb.N2D.Vmimagen2Dstandardram
	case Spot:
		core = p.Gcp.Compute.GCE.VmsPreemptible.CoresPerCore.N2D.Vmimagepreemptiblen2Dstandardcore
		memory = p.Gcp.Compute.GCE.VmsPreemptible.MemoryPerGb.N2D.Vmimagepreemptiblen2Dstandardram
	case Commitment1YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit1Year.CoresPerCore.N2D.Commitmentn2Dcpu1Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit1Year.MemoryPerGb.N2D.Commitmentn2Dram1Yv1
	case Commitment3YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit3Year.CoresPerCore.N2D.Commitmentn2Dcpu3Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit3Year.MemoryPerGb.N2D.Commitmentn2Dram3Yv1
	default:
		return Subtype{}, Subtype{}, fmt.Errorf("commitment %q not supported", opts.Commitment)
	}
	return core, memory, nil
}

func typeGetterT2A(p *Pricing, opts Opts) (Subtype, Subtype, error) {
	var core Subtype
	var memory Subtype
	switch opts.Commitment {
	case OnDemand:
		core = p.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.T2A.Vmimaget2Astandardcore
		memory = p.Gcp.Compute.GCE.VmsOnDemand.MemoryPerGb.T2A.Vmimaget2Astandardram
	case Spot:
		core = p.Gcp.Compute.GCE.VmsPreemptible.CoresPerCore.T2A.Vmimagepreemptiblet2Astandardcore
		memory = p.Gcp.Compute.GCE.VmsPreemptible.MemoryPerGb.T2A.Vmimagepreemptiblet2Astandardram
	default:
		return Subtype{}, Subtype{}, fmt.Errorf("commitment %q not supported", opts.Commitment)
	}
	return core, memory, nil
}

func typeGetterT2D(p *Pricing, opts Opts) (Subtype, Subtype, error) {
	var core Subtype
	var memory Subtype
	switch opts.Commitment {
	case OnDemand:
		core = p.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.T2D.Vmimaget2Dstandardcore
		memory = p.Gcp.Compute.GCE.VmsOnDemand.MemoryPerGb.T2D.Vmimaget2Dstandardram
	case Spot:
		core = p.Gcp.Compute.GCE.VmsPreemptible.CoresPerCore.T2D.Vmimagepreemptiblet2Dstandardcore
		memory = p.Gcp.Compute.GCE.VmsPreemptible.MemoryPerGb.T2D.Vmimagepreemptiblet2Dstandardram
	case Commitment1YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit1Year.CoresPerCore.T2D.Commitmentt2Dcpu1Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit1Year.MemoryPerGb.T2D.Commitmentt2Dram1Yv1
	case Commitment3YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit3Year.CoresPerCore.T2D.Commitmentt2Dcpu3Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit3Year.MemoryPerGb.T2D.Commitmentt2Dram3Yv1
	default:
		return Subtype{}, Subtype{}, fmt.Errorf("commitment %q not supported", opts.Commitment)
	}
	return core, memory, nil
}

func typeGetterN1(p *Pricing, opts Opts) (Subtype, Subtype, error) {
	var core Subtype
	var memory Subtype
	switch opts.Commitment {
	case OnDemand:
		core = p.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.N1.Vmimagen1Standardcore
		memory = p.Gcp.Compute.GCE.VmsOnDemand.MemoryPerGb.N1.Vmimagen1Standardram
	case Spot:
		core = p.Gcp.Compute.GCE.VmsPreemptible.CoresPerCore.N1.Vmimagepreemptiblen1Standardcore
		memory = p.Gcp.Compute.GCE.VmsPreemptible.MemoryPerGb.N1.Vmimagepreemptiblen1Standardram
	default:
		return Subtype{}, Subtype{}, fmt.Errorf("commitment %q not supported", opts.Commitment)
	}
	return core, memory, nil
}

func typeGetterC2D(p *Pricing, opts Opts) (Subtype, Subtype, error) {
	var core Subtype
	var memory Subtype
	switch opts.Commitment {
	case OnDemand:
		core = p.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.C2D.Vmimagec2Dstandardcore
		memory = p.Gcp.Compute.GCE.VmsOnDemand.MemoryPerGb.C2D.Vmimagec2Dstandardram
	case Spot:
		core = p.Gcp.Compute.GCE.VmsPreemptible.CoresPerCore.C2D.Vmimagepreemptiblec2Dstandardcore
		memory = p.Gcp.Compute.GCE.VmsPreemptible.MemoryPerGb.C2D.Vmimagepreemptiblec2Dstandardram
	case Commitment1YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit1Year.CoresPerCore.C2D.Commitmentc2Dcpu1Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit1Year.MemoryPerGb.C2D.Commitmentc2Dram1Yv1
	case Commitment3YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit3Year.CoresPerCore.C2D.Commitmentc2Dcpu3Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit3Year.MemoryPerGb.C2D.Commitmentc2Dram3Yv1
	default:
		return Subtype{}, Subtype{}, fmt.Errorf("commitment %q not supported", opts.Commitment)
	}
	return core, memory, nil
}

func typeGetterM3(p *Pricing, opts Opts) (Subtype, Subtype, error) {
	var core Subtype
	var memory Subtype
	switch opts.Commitment {
	case OnDemand:
		core = p.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.M3.Vmimagem3Standardcore
		memory = p.Gcp.Compute.GCE.VmsOnDemand.MemoryPerGb.M3.Vmimagem3Standardram
	case Spot:
		core = p.Gcp.Compute.GCE.VmsPreemptible.CoresPerCore.M3.Vmimagepreemptiblem3Standardcore
		memory = p.Gcp.Compute.GCE.VmsPreemptible.MemoryPerGb.M3.Vmimagepreemptiblem3Standardram
	case Commitment1YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit1Year.CoresPerCore.M3.Commitmentm3CPU1Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit1Year.MemoryPerGb.M3.Commitmentm3RAM1Yv1
	case Commitment3YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit3Year.CoresPerCore.M3.Commitmentm3CPU3Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit3Year.MemoryPerGb.M3.Commitmentm3RAM3Yv1
	default:
		return Subtype{}, Subtype{}, fmt.Errorf("commitment %q not supported", opts.Commitment)
	}
	return core, memory, nil
}

func getHourly(p *Pricing, opts Opts, tg typeMachineGetter) (uint64, error) {
	core, memory, err := tg(p, opts)
	if err != nil {
		return 0, err
	}

	var corePricePerRegion uint64 = 0
	if region, ok := core.Regions[opts.Region]; ok {
		if len(region.Prices) > 0 {
			corePricePerRegion = region.Prices[0].Nanos
		}
	} else {
		return 0, fmt.Errorf("core price not found for %q region", opts.Region)
	}

	var memoryPricePerRegion uint64 = 0
	if region, ok := memory.Regions[opts.Region]; ok {
		if len(region.Prices) > 0 {
			memoryPricePerRegion = region.Prices[0].Nanos
		}
	} else {
		return 0, fmt.Errorf("memory not found for %q region", opts.Region)
	}

	var sum uint64 = 0
	sum += corePricePerRegion * opts.NumOfCPU
	sum += memoryPricePerRegion * opts.NumOfMemory
	return sum, nil
}
