package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-07-01/compute"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	. "github.com/mlabouardy/komiser/models/azure"
)

func getSnapshotClient(subscriptionID string) compute.SnapshotsClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	snapshotClient := compute.NewSnapshotsClient(subscriptionID)
	snapshotClient.Authorizer = a
	return snapshotClient
}

func (azure Azure) DescribeSnapshots(subscriptionID string) ([]Snapshot, error) {
	snapshotClient := getSnapshotClient(subscriptionID)
	snapshots := make([]Snapshot, 0)
	ctx := context.Background()
	for sItr, err := snapshotClient.ListComplete(ctx); sItr.NotDone(); sItr.Next() {
		if err != nil {
			return snapshots, err
		}
		snapshot := sItr.Value()
		snapshots = append(snapshots, Snapshot{
			Size: *snapshot.DiskSizeGB,
		})
	}
	return snapshots, nil
}
