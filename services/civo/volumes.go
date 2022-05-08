package civo

func (civo Civo) GetVolumesCount(apiKey, regionCode string) (int, error) {
	client, err := getCivoClient(apiKey, regionCode)
	if err != nil {
		return 0, err
	}
	volumes, err := client.ListVolumes()
	if err != nil {
		return 0, err
	}
	return len(volumes), nil
}
