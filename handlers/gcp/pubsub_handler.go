package aws

import (
	"net/http"
)

func (handler *GCPHandler) PubSubTopicsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("pubsub_topics")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetTopics()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "pubsub:PubsubScope is missing")
		} else {
			handler.cache.Set("pubsub_topics", response)
			respondWithJSON(w, 200, response)
		}
	}
}
