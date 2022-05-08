package civo

func (civo Civo) GetDisksCount(apiKey, regionCode string) (int, error) {
	client, err := getCivoClient(apiKey, regionCode)
	if err != nil {
		return 0, err
	}
	diskImages, err := client.ListDiskImages()
	if err != nil {
		return 0, err
	}
	return len(diskImages), nil
}
