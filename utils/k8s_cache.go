package utils

import (
	"context"

	"github.com/tailwarden/komiser/providers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	SERVICES = "SERVICES"
)

func K8s_Cache(client *providers.ProviderClient, callerType string) (interface{}, error) {
	var response interface{}
	var err error
	switch callerType {
	case SERVICES:
		response, err = client.K8sClient.Client.CoreV1().Services("").List(context.Background(), metav1.ListOptions{})
		if err != nil {
			return response, err
		}
	}

	client.Cache.Set(callerType, response, 0)

	return response, nil
}
