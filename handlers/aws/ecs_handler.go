package aws

import (
	"net/http"
)

func (handler *AWSHandler) ECSHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("ecs")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.DescribeECS(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "ecs:DescribeClusters or ecs:DescribeTasks or ecs:DescribeServices is missing")
		} else {
			handler.cache.Set("ecs", response)
			respondWithJSON(w, 200, response)
		}
	}
}
