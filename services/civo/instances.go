package civo

func (civo Civo) GetInstanceCount(apiKey, regionCode string) (int, error) {
	client, err := getCivoClient(apiKey, regionCode)
	if err != nil {
		return 0, err
	}
	instances, err := client.ListAllInstances()
	if err != nil {
		return 0, err
	}
	return len(instances), nil
}
