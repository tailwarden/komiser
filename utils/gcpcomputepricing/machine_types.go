package gcpcomputepricing

import "errors"

const (
	E2  = "e2"
	C3  = "c3"
	N2  = "n2"
	N2D = "n2d"
	T2A = "t2a"
	T2D = "t2d"
	N1  = "n1"
	C2  = "c2"
	C2D = "c2d"
	M1  = "m1"
	M2  = "m2"
	M3  = "m3"
)

const (
	OnDemand                = "on-demand"
	Spot                    = "spot"
	Commitment1YearResource = "commitment-1-year-resource"
	Commitment3YearResource = "commitment-3-year-resource"
)

type Opts struct {
	Type        string
	Commitment  string
	Region      string
	NumOfCPU    uint64
	NumOfMemory uint64
}

func CalculateMachine(p *Pricing, opts Opts) (uint64, error) {
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

type typeGetter func(p *Pricing, opts Opts) (Subtype, Subtype, error)

func typeGetterE2(p *Pricing, opts Opts) (Subtype, Subtype, error) {
	var core Subtype
	var memory Subtype
	switch opts.Commitment {
	case OnDemand:
		core = p.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.Vmimagee2Core
		memory = p.Gcp.Compute.GCE.VmsOnDemand.MemoryPerGb.Vmimagee2RAM
	case Spot:
		core = p.Gcp.Compute.GCE.VmsPreemptible.CoresPerCore.Vmimagepreemptiblee2Core
		memory = p.Gcp.Compute.GCE.VmsPreemptible.MemoryPerGb.Vmimagepreemptiblee2RAM
	case Commitment1YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit1Year.CoresPerCore.Commitmente2CPU1Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit1Year.MemoryPerGb.Commitmente2RAM1Yv1
	case Commitment3YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit3Year.CoresPerCore.Commitmente2CPU3Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit3Year.MemoryPerGb.Commitmente2RAM3Yv1
	default:
		return Subtype{}, Subtype{}, errors.New("commitment not supported") // TODO improve error
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
		return Subtype{}, Subtype{}, errors.New("commitment not supported") // TODO improve error
	}
	return core, memory, nil
}

func typeGetterN2(p *Pricing, opts Opts) (Subtype, Subtype, error) {
	var core Subtype
	var memory Subtype
	switch opts.Commitment {
	case OnDemand:
		core = p.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.Vmimagen2Standardcore
		memory = p.Gcp.Compute.GCE.VmsOnDemand.MemoryPerGb.Vmimagen2Standardram
	case Spot:
		core = p.Gcp.Compute.GCE.VmsPreemptible.CoresPerCore.Vmimagepreemptiblen2Standardcore
		memory = p.Gcp.Compute.GCE.VmsPreemptible.MemoryPerGb.Vmimagepreemptiblen2Standardram
	case Commitment1YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit1Year.CoresPerCore.Commitmentn2CPU1Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit1Year.MemoryPerGb.Commitmentn2RAM1Yv1
	case Commitment3YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit3Year.CoresPerCore.Commitmentn2CPU3Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit3Year.MemoryPerGb.Commitmentn2RAM3Yv1
	default:
		return Subtype{}, Subtype{}, errors.New("commitment not supported") // TODO improve error
	}
	return core, memory, nil
}

func typeGetterN2D(p *Pricing, opts Opts) (Subtype, Subtype, error) {
	var core Subtype
	var memory Subtype
	switch opts.Commitment {
	case OnDemand:
		core = p.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.Vmimagen2Dstandardcore
		memory = p.Gcp.Compute.GCE.VmsOnDemand.MemoryPerGb.Vmimagen2Dstandardram
	case Spot:
		core = p.Gcp.Compute.GCE.VmsPreemptible.CoresPerCore.Vmimagepreemptiblen2Dstandardcore
		memory = p.Gcp.Compute.GCE.VmsPreemptible.MemoryPerGb.Vmimagepreemptiblen2Dstandardram
	case Commitment1YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit1Year.CoresPerCore.Commitmentn2Dcpu1Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit1Year.MemoryPerGb.Commitmentn2Dram1Yv1
	case Commitment3YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit3Year.CoresPerCore.Commitmentn2Dcpu3Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit3Year.MemoryPerGb.Commitmentn2Dram3Yv1
	default:
		return Subtype{}, Subtype{}, errors.New("commitment not supported") // TODO improve error
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
		return Subtype{}, Subtype{}, errors.New("commitment not supported") // TODO improve error
	}
	return core, memory, nil
}

func typeGetterT2D(p *Pricing, opts Opts) (Subtype, Subtype, error) {
	var core Subtype
	var memory Subtype
	switch opts.Commitment {
	case OnDemand:
		core = p.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.Vmimaget2Dstandardcore
		memory = p.Gcp.Compute.GCE.VmsOnDemand.MemoryPerGb.Vmimaget2Dstandardram
	case Spot:
		core = p.Gcp.Compute.GCE.VmsPreemptible.CoresPerCore.Vmimagepreemptiblet2Dstandardcore
		memory = p.Gcp.Compute.GCE.VmsPreemptible.MemoryPerGb.Vmimagepreemptiblet2Dstandardram
	case Commitment1YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit1Year.CoresPerCore.Commitmentt2Dcpu1Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit1Year.MemoryPerGb.Commitmentt2Dram1Yv1
	case Commitment3YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit3Year.CoresPerCore.Commitmentt2Dcpu3Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit3Year.MemoryPerGb.Commitmentt2Dram3Yv1
	default:
		return Subtype{}, Subtype{}, errors.New("commitment not supported") // TODO improve error
	}
	return core, memory, nil
}

func typeGetterN1(p *Pricing, opts Opts) (Subtype, Subtype, error) {
	var core Subtype
	var memory Subtype
	switch opts.Commitment {
	case OnDemand:
		core = p.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.Vmimagen1Standardcore
		memory = p.Gcp.Compute.GCE.VmsOnDemand.MemoryPerGb.Vmimagen1Standardram
	case Spot:
		core = p.Gcp.Compute.GCE.VmsPreemptible.CoresPerCore.Vmimagepreemptiblen1Standardcore
		memory = p.Gcp.Compute.GCE.VmsPreemptible.MemoryPerGb.Vmimagepreemptiblen1Standardram
	default:
		return Subtype{}, Subtype{}, errors.New("commitment not supported") // TODO improve error
	}
	return core, memory, nil
}

func typeGetterC2D(p *Pricing, opts Opts) (Subtype, Subtype, error) {
	var core Subtype
	var memory Subtype
	switch opts.Commitment {
	case OnDemand:
		core = p.Gcp.Compute.GCE.VmsOnDemand.CoresPerCore.Vmimagec2Dstandardcore
		memory = p.Gcp.Compute.GCE.VmsOnDemand.MemoryPerGb.Vmimagec2Dstandardram
	case Spot:
		core = p.Gcp.Compute.GCE.VmsPreemptible.CoresPerCore.Vmimagepreemptiblec2Dstandardcore
		memory = p.Gcp.Compute.GCE.VmsPreemptible.MemoryPerGb.Vmimagepreemptiblec2Dstandardram
	case Commitment1YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit1Year.CoresPerCore.Commitmentc2Dcpu1Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit1Year.MemoryPerGb.Commitmentc2Dram1Yv1
	case Commitment3YearResource:
		core = p.Gcp.Compute.GCE.VmsCommit3Year.CoresPerCore.Commitmentc2Dcpu3Yv1
		memory = p.Gcp.Compute.GCE.VmsCommit3Year.MemoryPerGb.Commitmentc2Dram3Yv1
	default:
		return Subtype{}, Subtype{}, errors.New("commitment not supported") // TODO improve error
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
		return Subtype{}, Subtype{}, errors.New("commitment not supported") // TODO improve error
	}
	return core, memory, nil
}

func getHourly(p *Pricing, opts Opts, tg typeGetter) (uint64, error) {
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
		return 0, errors.New("core not found for that region") // TODO improve error
	}

	var memoryPricePerRegion uint64 = 0
	if region, ok := memory.Regions[opts.Region]; ok {
		if len(region.Prices) > 0 {
			memoryPricePerRegion = region.Prices[0].Nanos
		}
	} else {
		return 0, errors.New("memory not found for that region") // TODO improve error
	}

	var sum uint64 = 0
	sum += corePricePerRegion * opts.NumOfCPU
	sum += memoryPricePerRegion * opts.NumOfMemory
	return sum, nil
}
