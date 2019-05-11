package aws

import (
	"net/http"
)

func (handler *AWSHandler) DynamoDBTableHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("aws_dynamodb")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeDynamoDBTables(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "dynamodb:ListTables or dynamodb:DescribeTable is missing")
		} else {
			handler.cache.Set("aws_dynamodb", response)
			respondWithJSON(w, 200, response)
		}
	}
}
