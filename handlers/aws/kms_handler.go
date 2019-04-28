package aws

import (
	"net/http"
)

func (handler *AWSHandler) KMSKeysHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("kms_keys")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.ListKeys(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "kms:ListKeys is missing")
		} else {
			handler.cache.Set("kms_keys", response)
			respondWithJSON(w, 200, response)
		}
	}
}
