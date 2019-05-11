package aws

import (
	"net/http"
)

func (handler *GCPHandler) RedisInstancesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_redis_instances")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetRedisInstances()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "redis:CloudPlatformScope is missing")
		} else {
			handler.cache.Set("gcp_redis_instances", response)
			respondWithJSON(w, 200, response)
		}
	}
}
