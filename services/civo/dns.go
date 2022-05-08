package civo

func (civo Civo) GetDNSDomainsCount(apiKey, regionCode string) (int, error) {
	client, err := getCivoClient(apiKey, regionCode)
	if err != nil {
		return 0, err
	}
	dnsDomains, err := client.ListDNSDomains()
	if err != nil {
		return 0, err
	}
	return len(dnsDomains), nil
}
