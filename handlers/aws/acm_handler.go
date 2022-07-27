package aws

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"
)

func (handler *AWSHandler) APIGatewayListCertificatesHandler(w http.ResponseWriter, r *http.Request) {
	profile := r.Header.Get("profile")
	cfg, err := config.LoadDefaultConfig(context.Background())

	if handler.multiple {
		cfg, err = config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile(profile))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't read "+profile+" profile")
		}
	}

	key := fmt.Sprintf("aws.%s.acm.certificates", profile)

	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.ListCertificates(cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "acm:ListCertificates is missing")
		} else {
			handler.cache.Set(key, response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) APIGatewayExpiredCertificatesHandler(w http.ResponseWriter, r *http.Request) {
	profile := r.Header.Get("profile")
	cfg, err := config.LoadDefaultConfig(context.Background())

	if handler.multiple {
		cfg, err = config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile(profile))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't read "+profile+" profile")
		}
	}

	key := fmt.Sprintf("aws.%s.acm.expired", profile)

	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.ListExpiredCertificates(cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "acm:ListCertificates is missing")
		} else {
			handler.cache.Set(key, response)
			respondWithJSON(w, 200, response)
		}
	}
}
