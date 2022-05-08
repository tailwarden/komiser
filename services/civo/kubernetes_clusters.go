package civo

func (civo Civo) GetK8sClustersCount(apiKey, regionCode string) (int, error) {
	client, err := getCivoClient(apiKey, regionCode)
	if err != nil {
		return 0, err
	}
	page, err := client.ListKubernetesClusters()
	if err != nil {
		return 0, err
	}
	return len(page.Items), nil
}
