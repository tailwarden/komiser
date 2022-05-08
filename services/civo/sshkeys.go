package civo

func (civo Civo) GetSSHKeysCount(apiKey, regionCode string) (int, error) {
	client, err := getCivoClient(apiKey, regionCode)
	if err != nil {
		return 0, err
	}
	sshKeys, err := client.ListSSHKeys()
	if err != nil {
		return 0, err
	}
	return len(sshKeys), nil
}
