package aws

import "net/http"

func (handler *GCPHandler) KMSCryptoKeysHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_kms_crypto_keys")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetKMSCryptoKeys()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "cloudkms:CloudPlatformScope is missing")
		} else {
			handler.cache.Set("gcp_kms_crypto_keys", response)
			respondWithJSON(w, 200, response)
		}
	}
}
