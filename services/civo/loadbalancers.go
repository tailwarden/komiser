package civo

func (civo Civo) GetLoadBalancersCount(apiKey, regionCode string) (int, error) {
	client, err := getCivoClient(apiKey, regionCode)
	if err != nil {
		return 0, err
	}
	lbs, err := client.ListLoadBalancers()
	if err != nil {
		return 0, err
	}
	return len(lbs), nil
}
