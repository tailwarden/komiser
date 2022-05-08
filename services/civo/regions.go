package civo

func (civo Civo) GetRegionsCount(apiKey, regionCode string) (int, error) {
	client, err := getCivoClient(apiKey, regionCode)
	if err != nil {
		return 0, err
	}
	regions, err := client.ListRegions()
	if err != nil {
		return 0, err
	}
	return len(regions), nil
}
