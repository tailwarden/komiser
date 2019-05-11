package aws

import (
	"net/http"
)

func (handler *GCPHandler) SqlInstancesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("sql_instances")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetSqlInstances()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "sqladmin:CloudPlatformScope is missing")
		} else {
			handler.cache.Set("sql_instances", response)
			respondWithJSON(w, 200, response)
		}
	}
}
