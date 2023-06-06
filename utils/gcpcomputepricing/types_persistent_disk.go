package gcpcomputepricing

type PersistentDisk struct {
	AsyncReplication AsyncReplication        `json:"async_replication"`
	Snapshots        PersistentDiskSnapshots `json:"snapshots"`
	Diskops          Diskops                 `json:"diskops"`
	Standard         PersistentDiskStandard  `json:"standard"`
	SSD              PersistentDiskSSD       `json:"ssd"`
}

type AsyncReplication struct {
	BalancedProtection BalancedProtection         `json:"balanced_protection"`
	SsdProtection      SSDProtection              `json:"ssd_protection"`
	Networking         AsyncReplicationNetworking `json:"networking"`
}

type BalancedProtection struct {
	Asyncreplicationprotectionpdbalanced Subtype                    `json:"asyncreplicationprotectionpdbalanced"`
	Regional                             BalancedProtectionRegional `json:"regional"`
}

type BalancedProtectionRegional struct {
	Asyncreplicationprotectionregionalpdbalanced Subtype `json:"asyncreplicationprotectionregionalpdbalanced"`
}

type SSDProtection struct {
	Asyncreplicationprotectionpdssd Subtype               `json:"asyncreplicationprotectionpdssd"`
	Regional                        SSDProtectionRegional `json:"regional"`
}

type SSDProtectionRegional struct {
	Asyncreplicationprotectionregionalpdssd Subtype `json:"asyncreplicationprotectionregionalpdssd"`
}

type AsyncReplicationNetworking struct {
	Asia         AsyncReplicationNetworkingContinent `json:"asia"`
	Europe       AsyncReplicationNetworkingContinent `json:"europe"`
	NorthAmerica AsyncReplicationNetworkingContinent `json:"north_america"`
	Oceania      AsyncReplicationNetworkingContinent `json:"oceania"`
}

type AsyncReplicationNetworkingContinent struct {
	Asynchronousreplicationinterregionnetworkegress Subtype `json:"asynchronousreplicationinterregionnetworkegress"`
}

type PersistentDiskSnapshots struct {
	Multiregionalsnapshotdownload                     Subtype `json:"multiregionalsnapshotdownload"`
	Multiregionalsnapshotupload                       Subtype `json:"multiregionalsnapshotupload"`
	Storagemultiregionalstandardsnapshotearlydeletion Subtype `json:"storagemultiregionalstandardsnapshotearlydeletion"`
	Storageregionalarchivesnapshotdatastorage         Subtype `json:"storageregionalarchivesnapshotdatastorage"`
	Storageregionalarchivesnapshotearlydeletion       Subtype `json:"storageregionalarchivesnapshotearlydeletion"`
	Storageregionalarchivesnapshotretrieval           Subtype `json:"storageregionalarchivesnapshotretrieval"`
	Storageregionalstandardsnapshotearlydeletion      Subtype `json:"storageregionalstandardsnapshotearlydeletion"`
}

type Diskops struct {
	Pdiorequests Subtype `json:"pdiorequests"`
}

type PersistentDiskStandard struct {
	Capacity PersistentDiskStandardCapacity `json:"capacity"`
	Snapshot PersistentDiskStandardSnapshot `json:"snapshot"`
}

type PersistentDiskStandardCapacity struct {
	Regional          PersistentDiskStandardCapacityRegional `json:"regional"`
	Storagepdcapacity Subtype                                `json:"storagepdcapacity"`
}

type PersistentDiskStandardCapacityRegional struct {
	Regionalstoragepdcapacity Subtype `json:"regionalstoragepdcapacity"`
}

type PersistentDiskStandardSnapshot struct {
	Storagepdsnapshot Subtype `json:"storagepdsnapshot"`
}

type PersistentDiskSSD struct {
	Capacity PersistentDiskSSDCapacity `json:"capacity"`
}

type PersistentDiskSSDCapacity struct {
	Regional     PersistentDiskSSDCapacityRegional     `json:"regional"`
	Storagepdssd Subtype                               `json:"storagepdssd"`
	Extreme      PersistentDiskSSDCapacityExtreme      `json:"extreme"`
	Lite         PersistentDiskSSDCapacityLite         `json:"lite"`
	RegionalLite PersistentDiskSSDCapacityRegionalLite `json:"regional_lite"`
}

type PersistentDiskSSDCapacityRegional struct {
	Regionalstoragepdssd Subtype `json:"regionalstoragepdssd"`
}

type PersistentDiskSSDCapacityExtreme struct {
	Storagepdssdextremecapacity Subtype `json:"storagepdssdextremecapacity"`
	Storagepdssdextremeiops     Subtype `json:"storagepdssdextremeiops"`
}

type PersistentDiskSSDCapacityLite struct {
	Storagepdssdlitecapacity Subtype `json:"storagepdssdlitecapacity"`
}

type PersistentDiskSSDCapacityRegionalLite struct {
	Storageregionalpdssdlitecapacity Subtype `json:"storageregionalpdssdlitecapacity"`
}
