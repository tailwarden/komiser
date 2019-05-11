package aws

import (
	"net/http"
)

func (handler *GCPHandler) BigQueryScannedStatementsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_bigquery_statements")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetBigQueryScannedStatements()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "monitoring:MonitoringReadScope is missing")
		} else {
			handler.cache.Set("gcp_bigquery_statements", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) BigQueryStoredBytesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_bigquery_storage")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetBigQueryStoredBytes()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "monitoring:MonitoringReadScope is missing")
		} else {
			handler.cache.Set("gcp_bigquery_storage", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) BigQueryTablesHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_bigquery_tables")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetBigQueryTables()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "bigquery:CloudPlatformReadOnlyScope is missing")
		} else {
			handler.cache.Set("gcp_bigquery_tables", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *GCPHandler) BigQueryDatasetsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("gcp_bigquery_datasets")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.gcp.GetBigQueryDatasets()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "bigquery:CloudPlatformReadOnlyScope is missing")
		} else {
			handler.cache.Set("gcp_bigquery_datasets", response)
			respondWithJSON(w, 200, response)
		}
	}
}
