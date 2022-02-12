package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2021-09-01/containerservice"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	. "github.com/mlabouardy/komiser/models/azure"
)

func getManagedClustersClient(subscriptionID string) containerservice.ManagedClustersClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	mgdClustersClient := containerservice.NewManagedClustersClient(subscriptionID)
	mgdClustersClient.Authorizer = a
	return mgdClustersClient
}

func (azure Azure) DescribeManagedClusters(subscriptionID string) ([]Cluster, error) {
	mgdClustersClient := getManagedClustersClient(subscriptionID)
	ctx := context.Background()
	clusters := make([]Cluster, 0)
	for clusterItr, err := mgdClustersClient.ListComplete(ctx); clusterItr.NotDone(); clusterItr.Next() {
		if err != nil {
			return clusters, err
		}
		cluster := clusterItr.Value()
		clusters = append(clusters, Cluster{
			Name:              *cluster.Name,
			NodeResourceGroup: *cluster.ManagedClusterProperties.NodeResourceGroup,
			KubernetesVersion: *cluster.ManagedClusterProperties.KubernetesVersion,
			Location:          *cluster.Location,
		})
	}
	return clusters, nil
}
