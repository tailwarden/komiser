package gcp

import (
	"context"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
)

func (gcp GCP) GetVpcNetworks() (int, error) {
	total := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err != nil {
		return total, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := compute.New(client)
	if err != nil {
		return total, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return total, err
	}
	for _, project := range projects {
		networks, err := svc.Networks.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return total, err
		}
		total += len(networks.Items)
	}
	return total, nil
}

func (gcp GCP) GetNetworkFirewalls() (int, error) {
	total := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err != nil {
		return total, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := compute.New(client)
	if err != nil {
		return total, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return total, err
	}

	for _, project := range projects {
		firewalls, err := svc.Firewalls.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return total, err
		}
		total += len(firewalls.Items)
	}
	return total, nil
}

func (gcp GCP) GetNetworkRouters() (int, error) {
	total := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err != nil {
		return total, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := compute.New(client)
	if err != nil {
		return total, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return total, err
	}

	for _, project := range projects {
		regions, err := svc.Regions.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return total, err
		}
		for _, region := range regions.Items {
			routers, err := svc.Routers.List(project.ID, region.Name).Do()
			if err != nil {
				log.Println(err)
				return total, err
			}
			total += len(routers.Items)
		}
	}
	return total, nil
}

func (gcp GCP) GetNatGateways() (int, error) {
	total := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err != nil {
		return total, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := compute.New(client)
	if err != nil {
		return total, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return total, err
	}

	for _, project := range projects {
		regions, err := svc.Regions.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return total, err
		}
		for _, region := range regions.Items {
			routers, err := svc.Routers.List(project.ID, region.Name).Do()
			if err != nil {
				log.Println(err)
				return total, err
			}

			for _, router := range routers.Items {
				total += len(router.Nats)
			}

		}
	}
	return total, nil
}

func (gcp GCP) GetSubnetsNumber() (int, error) {
	total := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err != nil {
		return total, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := compute.New(client)
	if err != nil {
		return total, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return total, err
	}

	for _, project := range projects {
		regions, err := svc.Regions.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return total, err
		}
		for _, region := range regions.Items {
			subnets, err := svc.Subnetworks.List(project.ID, region.Name).Do()
			if err != nil {
				log.Println(err)
				return total, err
			}
			total += len(subnets.Items)
		}
	}
	return total, nil
}

func (gcp GCP) GetExternalAddresses() (int, error) {
	total := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err != nil {
		return total, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := compute.New(client)
	if err != nil {
		return total, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return total, err
	}

	for _, project := range projects {
		regions, err := svc.Regions.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return total, err
		}
		for _, region := range regions.Items {
			addresses, err := svc.Addresses.List(project.ID, region.Name).Do()
			if err != nil {
				log.Println(err)
				return total, err
			}
			total += len(addresses.Items)
		}
	}
	return total, nil
}

func (gcp GCP) GetVpnTunnels() (int, error) {
	total := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err != nil {
		return total, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := compute.New(client)
	if err != nil {
		return total, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return total, err
	}

	for _, project := range projects {
		regions, err := svc.Regions.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return total, err
		}
		for _, region := range regions.Items {
			tunnels, err := svc.VpnTunnels.List(project.ID, region.Name).Do()
			if err != nil {
				log.Println(err)
				return total, err
			}
			total += len(tunnels.Items)
		}
	}
	return total, nil
}

func (gcp GCP) GetSSLCertificates() (int, error) {
	total := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err != nil {
		return total, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := compute.New(client)
	if err != nil {
		return total, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return total, err
	}

	for _, project := range projects {
		certificates, err := svc.SslCertificates.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return total, err
		}
		total += len(certificates.Items)
	}
	return total, nil
}

func (gcp GCP) GetSSLPolicies() (int, error) {
	total := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err != nil {
		return total, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := compute.New(client)
	if err != nil {
		return total, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return total, err
	}

	for _, project := range projects {
		policies, err := svc.SslPolicies.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return total, err
		}
		total += len(policies.Items)
	}
	return total, nil
}

func (gcp GCP) GetSecurityPolicies() (int, error) {
	total := 0

	src, err := google.DefaultTokenSource(oauth2.NoContext, compute.ComputeReadonlyScope)
	if err != nil {
		return total, err
	}
	client := oauth2.NewClient(context.Background(), src)

	svc, err := compute.New(client)
	if err != nil {
		return total, err
	}

	projects, err := gcp.GetProjects()
	if err != nil {
		return total, err
	}

	for _, project := range projects {
		policies, err := svc.SecurityPolicies.List(project.ID).Do()
		if err != nil {
			log.Println(err)
			return total, err
		}
		total += len(policies.Items)
	}
	return total, nil
}
