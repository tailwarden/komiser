package azure

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-12-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-06-01/resources"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/pkg/errors"
)

func getVMClient(subscriptionID string) compute.VirtualMachinesClient {
	vmClient := compute.NewVirtualMachinesClient(subscriptionID)
	return vmClient
}

func getGroups(subscriptionID string) ([]string, error) {
	tab := make([]string, 0)
	var err error

	grClient := resources.NewGroupsClient(subscriptionID)
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
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
	Image  string `json:"image"`
	Region string `json:"region"`
	Status string `json:"status"`
	Disk   int    `json:"disk"`
}

func (azure Azure) DescribeVMs(subscriptionID string) ([]Vm, error) {
	a, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		panic(err)
	}
	vmClient := getVMClient(subscriptionID)
	vmClient.Authorizer = a
	var filter string
	groups, err := getGroups(subscriptionID)
	for _, group := range groups {
		log.Println(group)
	}
	//TODO Move below into new function and call using goroutine
	filter = "myResourceGroup"
	listOfVms := make([]Vm, 0)
	ctx := context.Background()
	for vm, _ := vmClient.ListComplete(ctx, filter); vm.NotDone(); vm.Next() {
		i := vm.Value()
		listOfVms = append(listOfVms, Vm{
			Disk:   50,
			Image:  *i.Name,
			Status: *i.Name,
			Region: *i.Name,
		})

	}
	log.Println(listOfVms)
	return listOfVms, nil
}
