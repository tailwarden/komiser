package civo

func (civo Civo) GetPrivateNetworks(apiKey, regionCode string) (int, error) {
	client, err := getCivoClient(apiKey, regionCode)
	if err != nil {
		return 0, err
	}
	privateNetworks, err := client.ListNetworks()
	if err != nil {
		return 0, err
	}
	return len(privateNetworks), nil
}
