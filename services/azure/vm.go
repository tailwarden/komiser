package azure

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-12-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-06-01/resources"
	"github.com/Azure/go-autorest/autorest/azure/auth"
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

type Vm struct {
	Name   string `json:"name"`
	Image  string `json:"image"`
	Region string `json:"region"`
	Disk   int32  `json:"disk"`
}

func (azure Azure) DescribeVMs(subscriptionID string) ([]Vm, error) {
	vmClient := getVMClient(subscriptionID)
	groups, _ := getGroups(subscriptionID)
	listOfVms := make([]Vm, 0)
	ctx := context.Background()
	for _, group := range groups {
		for vm, _ := vmClient.ListComplete(ctx, group); vm.NotDone(); vm.Next() {
			i := vm.Value()
			listOfVms = append(listOfVms, Vm{
				Name:   *i.Name,
				Disk:   *i.StorageProfile.OsDisk.DiskSizeGB,
				Image:  *i.StorageProfile.ImageReference.Offer,
				Region: *i.Location,
			})

		}
	}
	return listOfVms, nil
}
