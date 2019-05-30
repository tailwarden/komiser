package aws

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	. "github.com/mlabouardy/komiser/services/ini"
)

func (handler *AWSHandler) ConfigProfilesHandler(w http.ResponseWriter, r *http.Request) {
	if handler.multiple {
		sections, err := OpenFile(external.DefaultSharedCredentialsFilename())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't parse credentials file")
		}
		respondWithJSON(w, 200, sections.List())
	} else {
		respondWithJSON(w, 200, []string{})
	}

}
