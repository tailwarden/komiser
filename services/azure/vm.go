package azure

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-12-01/compute"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func getVMClient(subscriptionID string) compute.VirtualMachinesClient {
	vmClient := compute.NewVirtualMachinesClient(subscriptionID)
	return vmClient
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
	filter = "myResourceGroup"
	listOfVms := make([]Vm, 0)
	ctx := context.Background()
	vm, _ := vmClient.ListComplete(ctx, filter)
	log.Println(vm)
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
