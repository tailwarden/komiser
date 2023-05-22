package gcpcomputepricing

type GCE struct {
	VmsOnDemand                                      VmsOnDemand        `json:"vms_on_demand"`
	VmsPreemptible                                   VmsPreemptible     `json:"vms_preemptible"`
	VmsCommit1Year                                   VmsCommit1Year     `json:"vms_commit_1_year"`
	VmsCommit3Year                                   VmsCommit3Year     `json:"vms_commit_3_year"`
	NetworkOther                                     NetworkOther       `json:"network_other"`
	FlexCommit1Year                                  FlexCommit1Year    `json:"flex_commit_1_year"`
	FlexCommit3Year                                  FlexCommit3Year    `json:"flex_commit_3_year"`
	PremiumImage                                     PremiumImage       `json:"premium_image"`
	Ingress                                          Ingress            `json:"ingress"`
	VmsReservation                                   VmsReservation     `json:"vms_reservation"`
	Storagemultiregionalarchivesnapshotdatastorage   Subtype            `json:"storagemultiregionalarchivesnapshotdatastorage"`
	Storagemultiregionalarchivesnapshotearlydeletion Subtype            `json:"storagemultiregionalarchivesnapshotearlydeletion"`
	Storagemultiregionalarchivesnapshotretrieval     Subtype            `json:"storagemultiregionalarchivesnapshotretrieval"`
	Storageregionalarchivesnapshotearlydeletion      Subtype            `json:"storageregionalarchivesnapshotearlydeletion"`
	Storageregionalarchivesnapshotretrieval          Subtype            `json:"storageregionalarchivesnapshotretrieval"`
	MemoryPerGb                                      GCEMemoryPerGb     `json:"memory:_per_gb"`
	Management                                       GCEManagement      `json:"management"`
	WorkloadManager                                  GCEWorkloadManager `json:"workload_manager"`
}

type VmsOnDemand struct {
	AdvancedNetworking                                       VmsOnDemandAdvancedNetworking `json:"advanced_networking"`
	MemoryPerGb                                              VmsOnDemandMemoryPerGb        `json:"memory:_per_gb"`
	CoresPerCore                                             VmsOnDemandCoresPerCore       `json:"cores:_per_core"`
	Vmimagecomputeoptimizedsoletenancycore                   Subtype                       `json:"vmimagecomputeoptimizedsoletenancycore"`
	Vmimagecomputeoptimizedsoletenancycoresoletenancypremium Subtype                       `json:"vmimagecomputeoptimizedsoletenancycoresoletenancypremium"`
	Vmimagecomputeoptimizedsoletenancyram                    Subtype                       `json:"vmimagecomputeoptimizedsoletenancyram"`
	Vmimagecomputeoptimizedsoletenancyramsoletenancypremium  Subtype                       `json:"vmimagecomputeoptimizedsoletenancyramsoletenancypremium"`
	Vmimagelargesoletenancycore                              Subtype                       `json:"vmimagelargesoletenancycore"`
	Vmimagelargesoletenancycoresoletenancypremium            Subtype                       `json:"vmimagelargesoletenancycoresoletenancypremium"`
	Vmimagelargesoletenancyram                               Subtype                       `json:"vmimagelargesoletenancyram"`
	Vmimagelargesoletenancyramsoletenancypremium             Subtype                       `json:"vmimagelargesoletenancyramsoletenancypremium"`
	Vmimagen2Soletenancycoresoletenancyovercommitpremium     Subtype                       `json:"vmimagen2soletenancycoresoletenancyovercommitpremium"`
	Vmimagen2Soletenancyramsoletenancyovercommitpremium      Subtype                       `json:"vmimagen2soletenancyramsoletenancyovercommitpremium"`
	Vmimagen2Soletenancyramsoletenancypremium                Subtype                       `json:"vmimagen2soletenancyramsoletenancypremium"`
	Vmimagesoletenancycoresoletenancyovercommitpremium       Subtype                       `json:"vmimagesoletenancycoresoletenancyovercommitpremium"`
	Vmimagesoletenancyramsoletenancyovercommitpremium        Subtype                       `json:"vmimagesoletenancyramsoletenancyovercommitpremium"`
}

type VmsOnDemandAdvancedNetworking struct {
	Advnet100Gbpstotalondemand Subtype `json:"advnet100gbpstotalondemand"`
	Advnet50Gbpstotalondemand  Subtype `json:"advnet50gbpstotalondemand"`
	Advnet75Gbpstotalondemand  Subtype `json:"advnet75gbpstotalondemand"`
}

type VmsOnDemandMemoryPerGb struct {
	Vmstatesuspendedram                          Subtype                   `json:"vmstatesuspendedram"`
	Vmimagea2Highgpuram                          Subtype                   `json:"vmimagea2highgpuram"`
	Vmimagec2Dcustomextendedram                  Subtype                   `json:"vmimagec2dcustomextendedram"`
	Vmimagec2Dcustomram                          Subtype                   `json:"vmimagec2dcustomram"`
	Vmimagec2Dstandardram                        Subtype                   `json:"vmimagec2dstandardram"`
	C3                                           VmsOnDemandMemoryPerGbC3  `json:"c3"`
	Vmimagecomputeoptimizedram                   Subtype                   `json:"vmimagecomputeoptimizedram"`
	Vmimagecustomextendedram                     Subtype                   `json:"vmimagecustomextendedram"`
	Vmimagecustomram                             Subtype                   `json:"vmimagecustomram"`
	Vmimagee2RAM                                 Subtype                   `json:"vmimagee2ram"`
	G2                                           VmsOnDemandMemoryPerGbG2  `json:"g2"`
	Vmimagelargeram                              Subtype                   `json:"vmimagelargeram"`
	Vmimagelargerammemoryoptimizedupgradepremium Subtype                   `json:"vmimagelargerammemoryoptimizedupgradepremium"`
	M3                                           VmsOnDemandMemoryPerGbM3  `json:"m3"`
	Vmimagen1Standardram                         Subtype                   `json:"vmimagen1standardram"`
	Vmimagen2Customextendedram                   Subtype                   `json:"vmimagen2customextendedram"`
	Vmimagen2Customram                           Subtype                   `json:"vmimagen2customram"`
	Vmimagen2Dcustomextendedram                  Subtype                   `json:"vmimagen2dcustomextendedram"`
	Vmimagen2Dcustomram                          Subtype                   `json:"vmimagen2dcustomram"`
	Vmimagen2Dsoletenancyram                     Subtype                   `json:"vmimagen2dsoletenancyram"`
	Vmimagen2Dsoletenancyramsoletenancypremium   Subtype                   `json:"vmimagen2dsoletenancyramsoletenancypremium"`
	Vmimagen2Dstandardram                        Subtype                   `json:"vmimagen2dstandardram"`
	Vmimagen2Soletenancyram                      Subtype                   `json:"vmimagen2soletenancyram"`
	Vmimagen2Standardram                         Subtype                   `json:"vmimagen2standardram"`
	Vmimagesoletenancyram                        Subtype                   `json:"vmimagesoletenancyram"`
	Vmimagesoletenancyramsoletenancypremium      Subtype                   `json:"vmimagesoletenancyramsoletenancypremium"`
	T2A                                          VmsOnDemandMemoryPerGbT2A `json:"t2a"`
	Vmimaget2Dstandardram                        Subtype                   `json:"vmimaget2dstandardram"`
}

type VmsOnDemandMemoryPerGbC3 struct {
	Vmimagec3Soletenancyram                   Subtype `json:"vmimagec3soletenancyram"`
	Vmimagec3Soletenancyramsoletenancypremium Subtype `json:"vmimagec3soletenancyramsoletenancypremium"`
	Vmimagec3Standardram                      Subtype `json:"vmimagec3standardram"`
}

type VmsOnDemandMemoryPerGbG2 struct {
	Vmimageg2Soletenancyram                   Subtype `json:"vmimageg2soletenancyram"`
	Vmimageg2Soletenancyramsoletenancypremium Subtype `json:"vmimageg2soletenancyramsoletenancypremium"`
	Vmimageg2Standardram                      Subtype `json:"vmimageg2standardram"`
}

type VmsOnDemandMemoryPerGbM3 struct {
	Vmimagem3Soletenancyram                   Subtype `json:"vmimagem3soletenancyram"`
	Vmimagem3Soletenancyramsoletenancypremium Subtype `json:"vmimagem3soletenancyramsoletenancypremium"`
	Vmimagem3Standardram                      Subtype `json:"vmimagem3standardram"`
}

type VmsOnDemandMemoryPerGbT2A struct {
	Vmimaget2Astandardram Subtype `json:"vmimaget2astandardram"`
}

type VmsOnDemandCoresPerCore struct {
	Vmimagea2Highgpucore                          Subtype                    `json:"vmimagea2highgpucore"`
	Vmimagec2Dcustomcore                          Subtype                    `json:"vmimagec2dcustomcore"`
	Vmimagec2Dstandardcore                        Subtype                    `json:"vmimagec2dstandardcore"`
	C3                                            VmsOnDemandCoresPerCoreC3  `json:"c3"`
	Vmimagecomputeoptimizedcore                   Subtype                    `json:"vmimagecomputeoptimizedcore"`
	Vmimagecustomcore                             Subtype                    `json:"vmimagecustomcore"`
	Vmimagee2Core                                 Subtype                    `json:"vmimagee2core"`
	Vmimagef1Micro                                Subtype                    `json:"vmimagef1micro"`
	Vmimageg1Small                                Subtype                    `json:"vmimageg1small"`
	G2                                            VmsOnDemandCoresPerCoreG2  `json:"g2"`
	Vmimagelargecore                              Subtype                    `json:"vmimagelargecore"`
	Vmimagelargecorememoryoptimizedupgradepremium Subtype                    `json:"vmimagelargecorememoryoptimizedupgradepremium"`
	M3                                            VmsOnDemandCoresPerCoreM3  `json:"m3"`
	Vmimagen1Standardcore                         Subtype                    `json:"vmimagen1standardcore"`
	Vmimagen2Customcore                           Subtype                    `json:"vmimagen2customcore"`
	Vmimagen2Customextendedcore                   Subtype                    `json:"vmimagen2customextendedcore"`
	Vmimagen2Dcustomcore                          Subtype                    `json:"vmimagen2dcustomcore"`
	Vmimagen2Dsoletenancycore                     Subtype                    `json:"vmimagen2dsoletenancycore"`
	Vmimagen2Dsoletenancycoresoletenancypremium   Subtype                    `json:"vmimagen2dsoletenancycoresoletenancypremium"`
	Vmimagen2Dstandardcore                        Subtype                    `json:"vmimagen2dstandardcore"`
	Vmimagen2Soletenancycore                      Subtype                    `json:"vmimagen2soletenancycore"`
	Vmimagen2Soletenancycoresoletenancypremium    Subtype                    `json:"vmimagen2soletenancycoresoletenancypremium"`
	Vmimagen2Standardcore                         Subtype                    `json:"vmimagen2standardcore"`
	Vmimagesoletenancycore                        Subtype                    `json:"vmimagesoletenancycore"`
	Vmimagesoletenancycoresoletenancypremium      Subtype                    `json:"vmimagesoletenancycoresoletenancypremium"`
	T2A                                           VmsOnDemandCoresPerCoreT2A `json:"t2a"`
	Vmimaget2Dstandardcore                        Subtype                    `json:"vmimaget2dstandardcore"`
}

type VmsOnDemandCoresPerCoreC3 struct {
	Vmimagec3Soletenancycore                   Subtype `json:"vmimagec3soletenancycore"`
	Vmimagec3Soletenancycoresoletenancypremium Subtype `json:"vmimagec3soletenancycoresoletenancypremium"`
	Vmimagec3Standardcore                      Subtype `json:"vmimagec3standardcore"`
}

type VmsOnDemandCoresPerCoreG2 struct {
	Vmimageg2Soletenancycore                   Subtype `json:"vmimageg2soletenancycore"`
	Vmimageg2Soletenancycoresoletenancypremium Subtype `json:"vmimageg2soletenancycoresoletenancypremium"`
	Vmimageg2Standardcore                      Subtype `json:"vmimageg2standardcore"`
}

type VmsOnDemandCoresPerCoreM3 struct {
	Vmimagem3Soletenancycore                   Subtype `json:"vmimagem3soletenancycore"`
	Vmimagem3Soletenancycoresoletenancypremium Subtype `json:"vmimagem3soletenancycoresoletenancypremium"`
	Vmimagem3Standardcore                      Subtype `json:"vmimagem3standardcore"`
}

type VmsOnDemandCoresPerCoreT2A struct {
	Vmimaget2Astandardcore Subtype `json:"vmimaget2astandardcore"`
}

type VmsPreemptible struct {
	AdvancedNetworking VmsPreemptibleAdvancedNetworking `json:"advanced_networking"`
	Cores1To64         VmsPreemptibleCores1To64         `json:"cores:_1_to_64"`
	MemoryPerGb        VmsPreemptibleMemoryPerGb        `json:"memory:_per_gb"`
	CoresPerCore       VmsPreemptibleCoresPerCore       `json:"cores:_per_core"`
}

type VmsPreemptibleAdvancedNetworking struct {
	Advnet100Gbpstotalpreemptible Subtype `json:"advnet100gbpstotalpreemptible"`
	Advnet50Gbpstotalpreemptible  Subtype `json:"advnet50gbpstotalpreemptible"`
	Advnet75Gbpstotalpreemptible  Subtype `json:"advnet75gbpstotalpreemptible"`
}

type VmsPreemptibleCores1To64 struct {
	Highcpu VmsPreemptibleHighcpu `json:"highcpu"`
}

type VmsPreemptibleHighcpu struct {
	Vmimagepreemptiblea2Highgpucore Subtype `json:"vmimagepreemptiblea2highgpucore"`
}

type VmsPreemptibleMemoryPerGb struct {
	Vmimagepreemptiblea2Highgpuram         Subtype                      `json:"vmimagepreemptiblea2highgpuram"`
	Vmimagepreemptiblec2Dcustomextendedram Subtype                      `json:"vmimagepreemptiblec2dcustomextendedram"`
	Vmimagepreemptiblec2Dcustomram         Subtype                      `json:"vmimagepreemptiblec2dcustomram"`
	Vmimagepreemptiblec2Dstandardram       Subtype                      `json:"vmimagepreemptiblec2dstandardram"`
	C3                                     VmsPreemptibleMemoryPerGbC3  `json:"c3"`
	Vmimagepreemptiblecomputeoptimizedram  Subtype                      `json:"vmimagepreemptiblecomputeoptimizedram"`
	Vmimagepreemptiblecustomextendedram    Subtype                      `json:"vmimagepreemptiblecustomextendedram"`
	Vmimagepreemptiblecustomram            Subtype                      `json:"vmimagepreemptiblecustomram"`
	Vmimagepreemptiblee2RAM                Subtype                      `json:"vmimagepreemptiblee2ram"`
	G2                                     VmsPreemptibleMemoryPerGbG2  `json:"g2"`
	Vmimagepreemptiblelargeram             Subtype                      `json:"vmimagepreemptiblelargeram"`
	M3                                     VmsPreemptibleMemoryPerGbM3  `json:"m3"`
	Vmimagepreemptiblen1Standardram        Subtype                      `json:"vmimagepreemptiblen1standardram"`
	Vmimagepreemptiblen2Customextendedram  Subtype                      `json:"vmimagepreemptiblen2customextendedram"`
	Vmimagepreemptiblen2Customram          Subtype                      `json:"vmimagepreemptiblen2customram"`
	Vmimagepreemptiblen2Dcustomextendedram Subtype                      `json:"vmimagepreemptiblen2dcustomextendedram"`
	Vmimagepreemptiblen2Dcustomram         Subtype                      `json:"vmimagepreemptiblen2dcustomram"`
	Vmimagepreemptiblen2Dstandardram       Subtype                      `json:"vmimagepreemptiblen2dstandardram"`
	Vmimagepreemptiblen2Standardram        Subtype                      `json:"vmimagepreemptiblen2standardram"`
	T2A                                    VmsPreemptibleMemoryPerGbT2A `json:"t2a"`
	Vmimagepreemptiblet2Dstandardram       Subtype                      `json:"vmimagepreemptiblet2dstandardram"`
}

type VmsPreemptibleMemoryPerGbC3 struct {
	Vmimagepreemptiblec3Standardram Subtype `json:"vmimagepreemptiblec3standardram"`
}

type VmsPreemptibleMemoryPerGbG2 struct {
	Vmimagepreemptibleg2Standardram Subtype `json:"vmimagepreemptibleg2standardram"`
}

type VmsPreemptibleMemoryPerGbM3 struct {
	Vmimagepreemptiblem3Standardram Subtype `json:"vmimagepreemptiblem3standardram"`
}

type VmsPreemptibleMemoryPerGbT2A struct {
	Vmimagepreemptiblet2Astandardram Subtype `json:"vmimagepreemptiblet2astandardram"`
}

type VmsPreemptibleCoresPerCore struct {
	Vmimagepreemptiblec2Dcustomcore        Subtype                       `json:"vmimagepreemptiblec2dcustomcore"`
	Vmimagepreemptiblec2Dstandardcore      Subtype                       `json:"vmimagepreemptiblec2dstandardcore"`
	C3                                     VmsPreemptibleCoresPerCoreC3  `json:"c3"`
	Vmimagepreemptiblecomputeoptimizedcore Subtype                       `json:"vmimagepreemptiblecomputeoptimizedcore"`
	Vmimagepreemptiblecustomcore           Subtype                       `json:"vmimagepreemptiblecustomcore"`
	Vmimagepreemptiblecustomextendedcore   Subtype                       `json:"vmimagepreemptiblecustomextendedcore"`
	Vmimagepreemptiblee2Core               Subtype                       `json:"vmimagepreemptiblee2core"`
	Vmimagepreemptiblef1Micro              Subtype                       `json:"vmimagepreemptiblef1micro"`
	Vmimagepreemptibleg1Small              Subtype                       `json:"vmimagepreemptibleg1small"`
	G2                                     VmsPreemptibleCoresPerCoreG2  `json:"g2"`
	Vmimagepreemptiblelargecore            Subtype                       `json:"vmimagepreemptiblelargecore"`
	M3                                     VmsPreemptibleCoresPerCoreM3  `json:"m3"`
	Vmimagepreemptiblen1Standardcore       Subtype                       `json:"vmimagepreemptiblen1standardcore"`
	Vmimagepreemptiblen2Customcore         Subtype                       `json:"vmimagepreemptiblen2customcore"`
	Vmimagepreemptiblen2Customextendedcore Subtype                       `json:"vmimagepreemptiblen2customextendedcore"`
	Vmimagepreemptiblen2Dcustomcore        Subtype                       `json:"vmimagepreemptiblen2dcustomcore"`
	Vmimagepreemptiblen2Dstandardcore      Subtype                       `json:"vmimagepreemptiblen2dstandardcore"`
	Vmimagepreemptiblen2Standardcore       Subtype                       `json:"vmimagepreemptiblen2standardcore"`
	T2A                                    VmsPreemptibleCoresPerCoreT2A `json:"t2a"`
	Vmimagepreemptiblet2Dstandardcore      Subtype                       `json:"vmimagepreemptiblet2dstandardcore"`
}

type VmsPreemptibleCoresPerCoreC3 struct {
	Vmimagepreemptiblec3Standardcore Subtype `json:"vmimagepreemptiblec3standardcore"`
}

type VmsPreemptibleCoresPerCoreG2 struct {
	Vmimagepreemptibleg2Standardcore Subtype `json:"vmimagepreemptibleg2standardcore"`
}

type VmsPreemptibleCoresPerCoreM3 struct {
	Vmimagepreemptiblem3Standardcore Subtype `json:"vmimagepreemptiblem3standardcore"`
}

type VmsPreemptibleCoresPerCoreT2A struct {
	Vmimagepreemptiblet2Astandardcore Subtype `json:"vmimagepreemptiblet2astandardcore"`
}

type VmsCommit1Year struct {
	CoresPerCore              VmsCommit1YearCoresPerCore `json:"cores:_per_core"`
	MemoryPerGb               VmsCommit1YearMemoryPerGb  `json:"memory:_per_gb"`
	Vmwareengineucs12Moprepay Subtype                    `json:"vmwareengineucs12moprepay"`
}

type VmsCommit1YearCoresPerCore struct {
	Commitmenta2Highgpucpu1Yv1        Subtype                      `json:"commitmenta2highgpucpu1yv1"`
	Commitmentc2Dcpu1Yv1              Subtype                      `json:"commitmentc2dcpu1yv1"`
	C3                                VmsCommit1YearCoresPerCoreC3 `json:"c3"`
	Commitmentcpucomputeoptimized1Yv1 Subtype                      `json:"commitmentcpucomputeoptimized1yv1"`
	Commitmentcpulargeinstance1Yv1    Subtype                      `json:"commitmentcpulargeinstance1yv1"`
	Commitmentcpu1Yv1                 Subtype                      `json:"commitmentcpu1yv1"`
	Commitmente2CPU1Yv1               Subtype                      `json:"commitmente2cpu1yv1"`
	G2                                VmsCommit1YearCoresPerCoreG2 `json:"g2"`
	M3                                VmsCommit1YearCoresPerCoreM3 `json:"m3"`
	Commitmentn2CPU1Yv1               Subtype                      `json:"commitmentn2cpu1yv1"`
	Commitmentn2Dcpu1Yv1              Subtype                      `json:"commitmentn2dcpu1yv1"`
	Commitmentt2Dcpu1Yv1              Subtype                      `json:"commitmentt2dcpu1yv1"`
}

type VmsCommit1YearCoresPerCoreC3 struct {
	Commitmentc3CPU1Yv1 Subtype `json:"commitmentc3cpu1yv1"`
}

type VmsCommit1YearCoresPerCoreG2 struct {
	Commitmentg2CPU1Yv1 Subtype `json:"commitmentg2cpu1yv1"`
}

type VmsCommit1YearCoresPerCoreM3 struct {
	Commitmentm3CPU1Yv1 Subtype `json:"commitmentm3cpu1yv1"`
}

type VmsCommit1YearMemoryPerGb struct {
	Commitmenta2Highgpuram1Yv1        Subtype                     `json:"commitmenta2highgpuram1yv1"`
	Commitmentc2Dram1Yv1              Subtype                     `json:"commitmentc2dram1yv1"`
	C3                                VmsCommit1YearMemoryPerGbC3 `json:"c3"`
	Commitmente2RAM1Yv1               Subtype                     `json:"commitmente2ram1yv1"`
	G2                                VmsCommit1YearMemoryPerGbG2 `json:"g2"`
	M3                                VmsCommit1YearMemoryPerGbM3 `json:"m3"`
	Commitmentn2Dram1Yv1              Subtype                     `json:"commitmentn2dram1yv1"`
	Commitmentn2RAM1Yv1               Subtype                     `json:"commitmentn2ram1yv1"`
	Commitmentramcomputeoptimized1Yv1 Subtype                     `json:"commitmentramcomputeoptimized1yv1"`
	Commitmentramlargeinstance1Yv1    Subtype                     `json:"commitmentramlargeinstance1yv1"`
	Commitmentram1Yv1                 Subtype                     `json:"commitmentram1yv1"`
	Commitmentt2Dram1Yv1              Subtype                     `json:"commitmentt2dram1yv1"`
}

type VmsCommit1YearMemoryPerGbC3 struct {
	Commitmentc3RAM1Yv1 Subtype `json:"commitmentc3ram1yv1"`
}

type VmsCommit1YearMemoryPerGbG2 struct {
	Commitmentg2RAM1Yv1 Subtype `json:"commitmentg2ram1yv1"`
}

type VmsCommit1YearMemoryPerGbM3 struct {
	Commitmentm3RAM1Yv1 Subtype `json:"commitmentm3ram1yv1"`
}

type VmsCommit3Year struct {
	CoresPerCore              VmsCommit3YearCoresPerCore `json:"cores:_per_core"`
	MemoryPerGb               VmsCommit3YearMemoryPerGb  `json:"memory:_per_gb"`
	Vmwareengineucs36Moprepay Subtype                    `json:"vmwareengineucs36moprepay"`
}

type VmsCommit3YearCoresPerCore struct {
	Commitmenta2Highgpucpu3Yv1        Subtype                      `json:"commitmenta2highgpucpu3yv1"`
	Commitmentc2Dcpu3Yv1              Subtype                      `json:"commitmentc2dcpu3yv1"`
	C3                                VmsCommit3YearCoresPerCoreC3 `json:"c3"`
	Commitmentcpucomputeoptimized3Yv1 Subtype                      `json:"commitmentcpucomputeoptimized3yv1"`
	Commitmentcpulargeinstance1Yv1    Subtype                      `json:"commitmentcpulargeinstance1yv1"`
	Commitmentcpulargeinstance3Yv1    Subtype                      `json:"commitmentcpulargeinstance3yv1"`
	Commitmentcpu3Yv1                 Subtype                      `json:"commitmentcpu3yv1"`
	Commitmente2CPU3Yv1               Subtype                      `json:"commitmente2cpu3yv1"`
	G2                                VmsCommit3YearCoresPerCoreG2 `json:"g2"`
	M3                                VmsCommit3YearCoresPerCoreM3 `json:"m3"`
	Commitmentn2CPU3Yv1               Subtype                      `json:"commitmentn2cpu3yv1"`
	Commitmentn2Dcpu3Yv1              Subtype                      `json:"commitmentn2dcpu3yv1"`
	Commitmentt2Dcpu3Yv1              Subtype                      `json:"commitmentt2dcpu3yv1"`
}

type VmsCommit3YearCoresPerCoreC3 struct {
	Commitmentc3CPU3Yv1 Subtype `json:"commitmentc3cpu3yv1"`
}

type VmsCommit3YearCoresPerCoreG2 struct {
	Commitmentg2CPU3Yv1 Subtype `json:"commitmentg2cpu3yv1"`
}

type VmsCommit3YearCoresPerCoreM3 struct {
	Commitmentm3CPU3Yv1 Subtype `json:"commitmentm3cpu3yv1"`
}

type VmsCommit3YearMemoryPerGb struct {
	Commitmenta2Highgpuram3Yv1        Subtype                     `json:"commitmenta2highgpuram3yv1"`
	Commitmentc2Dram3Yv1              Subtype                     `json:"commitmentc2dram3yv1"`
	C3                                VmsCommit3YearMemoryPerGbC3 `json:"c3"`
	Commitmente2RAM3Yv1               Subtype                     `json:"commitmente2ram3yv1"`
	G2                                VmsCommit3YearMemoryPerGbG2 `json:"g2"`
	M3                                VmsCommit3YearMemoryPerGbM3 `json:"m3"`
	Commitmentn2Dram3Yv1              Subtype                     `json:"commitmentn2dram3yv1"`
	Commitmentn2RAM3Yv1               Subtype                     `json:"commitmentn2ram3yv1"`
	Commitmentramcomputeoptimized3Yv1 Subtype                     `json:"commitmentramcomputeoptimized3yv1"`
	Commitmentramlargeinstance3Yv1    Subtype                     `json:"commitmentramlargeinstance3yv1"`
	Commitmentram3Yv1                 Subtype                     `json:"commitmentram3yv1"`
	Commitmentt2Dram3Yv1              Subtype                     `json:"commitmentt2dram3yv1"`
}

type VmsCommit3YearMemoryPerGbC3 struct {
	Commitmentc3RAM3Yv1 Subtype `json:"commitmentc3ram3yv1"`
}

type VmsCommit3YearMemoryPerGbG2 struct {
	Commitmentg2RAM3Yv1 Subtype `json:"commitmentg2ram3yv1"`
}

type VmsCommit3YearMemoryPerGbM3 struct {
	Commitmentm3RAM3Yv1 Subtype `json:"commitmentm3ram3yv1"`
}

type NetworkOther struct {
	Externalip            Subtype `json:"externalip"`
	Externalippreemptible Subtype `json:"externalippreemptible"`
}

type FlexCommit1Year struct {
	Gcecommitmentsucs12Mo Subtype `json:"gcecommitmentsucs12mo"`
}

type FlexCommit3Year struct {
	Gcecommitmentsucs36Mo Subtype `json:"gcecommitmentsucs36mo"`
}

type PremiumImage struct {
	Microsoft PremiumImageMicrosoft `json:"microsoft"`
	Rhel      PremiumImageRhel      `json:"rhel"`
}

type PremiumImageMicrosoft struct {
	Windows   PremiumImageMicrosoftWindows   `json:"windows"`
	SQLServer PremiumImageMicrosoftSQLServer `json:"sql_server"`
}

type PremiumImageMicrosoftWindows struct {
	Licensed1656378918552316916Core    Subtype `json:"licensed1656378918552316916core"`
	Licensed1656378918552316916F1Micro Subtype `json:"licensed1656378918552316916f1micro"`
	Licensed1656378918552316916G1Small Subtype `json:"licensed1656378918552316916g1small"`
	Licensed3284763237085719542Core    Subtype `json:"licensed3284763237085719542core"`
	Licensed3284763237085719542F1Micro Subtype `json:"licensed3284763237085719542f1micro"`
	Licensed3284763237085719542G1Small Subtype `json:"licensed3284763237085719542g1small"`
	Licensed4819555115818134498Core    Subtype `json:"licensed4819555115818134498core"`
	Licensed4819555115818134498F1Micro Subtype `json:"licensed4819555115818134498f1micro"`
	Licensed4819555115818134498G1Small Subtype `json:"licensed4819555115818134498g1small"`
	Licensed4874454843789519845Core    Subtype `json:"licensed4874454843789519845core"`
	Licensed4874454843789519845F1Micro Subtype `json:"licensed4874454843789519845f1micro"`
	Licensed4874454843789519845G1Small Subtype `json:"licensed4874454843789519845g1small"`
	Licensed6107784707477449232Core    Subtype `json:"licensed6107784707477449232core"`
	Licensed6107784707477449232F1Micro Subtype `json:"licensed6107784707477449232f1micro"`
	Licensed6107784707477449232G1Small Subtype `json:"licensed6107784707477449232g1small"`
	Licensed7695108898142923768Core    Subtype `json:"licensed7695108898142923768core"`
	Licensed7695108898142923768F1Micro Subtype `json:"licensed7695108898142923768f1micro"`
	Licensed7695108898142923768G1Small Subtype `json:"licensed7695108898142923768g1small"`
	Licensed7798417859637521376Core    Subtype `json:"licensed7798417859637521376core"`
	Licensed7798417859637521376F1Micro Subtype `json:"licensed7798417859637521376f1micro"`
	Licensed7798417859637521376G1Small Subtype `json:"licensed7798417859637521376g1small"`
}

type PremiumImageMicrosoftSQLServer struct {
	Licensed1741222371620352982Core5Ormore Subtype `json:"licensed1741222371620352982core5ormore"`
	Licensed3039072951948447844Core5Ormore Subtype `json:"licensed3039072951948447844core5ormore"`
	Licensed3042936622923550835Core5Ormore Subtype `json:"licensed3042936622923550835core5ormore"`
	Licensed3398668354433905558Core5Ormore Subtype `json:"licensed3398668354433905558core5ormore"`
	Licensed6213885950785916969Core5Ormore Subtype `json:"licensed6213885950785916969core5ormore"`
	Licensed6795597790302237536Core5Ormore Subtype `json:"licensed6795597790302237536core5ormore"`
}

type PremiumImageRhel struct {
	Licensed7883559014960410759Corerange04      Subtype `json:"licensed7883559014960410759corerange04"`
	Licensed7883559014960410759Corerange5Ormore Subtype `json:"licensed7883559014960410759corerange5ormore"`
}

type Ingress struct {
	Premium     IngressPremium     `json:"premium"`
	InterRegion IngressInterRegion `json:"inter-region"`
	InterZone   IngressInterZone   `json:"inter-zone"`
	Standard    IngressStandard    `json:"standard"`
	IntraZone   IngressIntraZone   `json:"intra-zone"`
}

type IngressPremium struct {
	Networkgoogleingress   Subtype `json:"networkgoogleingress"`
	Networkinternetingress Subtype `json:"networkinternetingress"`
}

type IngressInterRegion struct {
	Networkinterregioningress Subtype `json:"networkinterregioningress"`
}

type IngressInterZone struct {
	Networkinterzoneingress Subtype `json:"networkinterzoneingress"`
}

type IngressStandard struct {
	Networkinternetstandardtieringress Subtype `json:"networkinternetstandardtieringress"`
}

type IngressIntraZone struct {
	Networkintrazoneingress Subtype `json:"networkintrazoneingress"`
}

type VmsReservation struct {
	CoresPerCore VmsReservationCoresPerCore `json:"cores:_per_core"`
	MemoryPerGb  VmsReservationMemoryPerGb  `json:"memory:_per_gb"`
}

type VmsReservationCoresPerCore struct {
	Reservationa2Highgpucore Subtype `json:"reservationa2highgpucore"`
}

type VmsReservationMemoryPerGb struct {
	Reservationa2Highgpuram Subtype `json:"reservationa2highgpuram"`
}

type GCEMemoryPerGb struct {
	Vmimagen2Soletenancyramsoletenancyovercommitpremium Subtype `json:"vmimagen2soletenancyramsoletenancyovercommitpremium"`
	Vmimagen2Soletenancyramsoletenancypremium           Subtype `json:"vmimagen2soletenancyramsoletenancypremium"`
	Vmimagesoletenancyramsoletenancyovercommitpremium   Subtype `json:"vmimagesoletenancyramsoletenancyovercommitpremium"`
}

type GCEManagement struct {
	Agentshourscount         Subtype `json:"agentshourscount"`
	Cloudopsagentshourscount Subtype `json:"cloudopsagentshourscount"`
}

type GCEWorkloadManager struct {
	BillingScannedResources Subtype `json:"billing/scanned_resources"`
}
