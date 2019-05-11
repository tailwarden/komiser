package aws

import (
	"net/http"
)

func (handler *GCPHandler) StorageBucketsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_storage_buckets")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetTotalBuckets()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "storage:CloudPlatformReadOnlyScope is missing")
		} else {
			handler.cache.Set("gcp_storage_buckets", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) StorageBucketsSizeHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_storage_size")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetBucketSize()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "monitoring:MonitoringReadScope is missing")
		} else {
			handler.cache.Set("gcp_storage_size", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) StorageBucketsObjectsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_storage_objects")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetBucketObjects()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "monitoring:MonitoringReadScope is missing")
		} else {
			handler.cache.Set("gcp_storage_objects", response)
			respondWithJSON(w, 200, response)
		}
	}
}
