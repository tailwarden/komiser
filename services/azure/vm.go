package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-12-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-06-01/resources"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	. "github.com/mlabouardy/komiser/models/azure"
	"github.com/pkg/errors"
)

func getVMClient(subscriptionID string) compute.VirtualMachinesClient {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	vmClient := compute.NewVirtualMachinesClient(subscriptionID)
	vmClient.Authorizer = a
	return vmClient
}

func getGroups(subscriptionID string) ([]string, error) {
	tab := make([]string, 0)
	var err error
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	grClient := resources.NewGroupsClient(subscriptionID)
	grClient.Authorizer = a
	for list, err := grClient.ListComplete(context.Background(), "", nil); list.NotDone(); err = list.Next() {
		if err != nil {
			return nil, errors.Wrap(err, "error traverising RG list")
		}
		rgName := *list.Value().Name
		tab = append(tab, rgName)
	}
	return tab, err
}

func (azure Azure) DescribeVMs(subscriptionID string) (map[string]interface{}, error) {
	outputVMsPerRegion := make(map[string]int, 0)
	outputVMsPerStatus := make(map[string]int, 0)
	outputVMsPerFamily := make(map[string]int, 0)
	vmClient := getVMClient(subscriptionID)
	groups, err := getGroups(subscriptionID)
	if err != nil {
		return map[string]interface{}{}, err
	}

	ctx := context.Background()
	for _, group := range groups {
		for vmItr, err := vmClient.ListComplete(ctx, group); vmItr.NotDone(); vmItr.Next() {
			if err != nil {
				return map[string]interface{}{}, err
			}
			vm := vmItr.Value()
			outputVMsPerStatus[vm.Status]++
			outputVMsPerRegion[*vm.Location]++
			outputVMsPerFamily[*vm.Type]++

		}
	}
	return map[string]interface{}{
		"region":  outputVMsPerRegion,
		"status":  outputVMsPerStatus,
		"familty": outputVMsPerFamily,
	}, nil
}

func (azure Azure) GetDisks(subscriptionID string) ([]Disk, error) {
	listOfDisks := make([]Disk, 0)
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	disksClient := compute.NewDisksClient(subscriptionID)
	disksClient.Authorizer = a
	ctx := context.Background()
	for disk, _ := disksClient.ListComplete(ctx); disk.NotDone(); disk.Next() {
		i := disk.Value()
		listOfDisks = append(listOfDisks, Disk{
			SizeGb: int64(*i.DiskSizeGB),
			Status: "Active",
		})
	}
	return listOfDisks, nil
}
