package aws

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws/external"
)

func (handler *AWSHandler) DynamoDBTableHandler(w http.ResponseWriter, r *http.Request) {
	profile := r.Header.Get("profile")
	cfg, err := external.LoadDefaultAWSConfig()

	if handler.multiple {
		cfg, err = external.LoadDefaultAWSConfig(external.WithSharedConfigProfile(profile))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't read "+profile+" profile")
		}
	}

	key := fmt.Sprintf("aws.%s.dynamodb.tables", profile)

	response, found := handler.cache.Get(key)
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeDynamoDBTables(cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "dynamodb:ListTables or dynamodb:DescribeTable is missing")
		} else {
			handler.cache.Set(key, response)
			respondWithJSON(w, 200, response)
		}
	}
}
