package aws

import (
	"net/http"
)

func (handler *AWSHandler) DataPipelineListPipelines(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("datapipeline_pipelines")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.ListDataPipelines(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "datapipeline:ListPipelines is missing")
		} else {
			handler.cache.Set("datapipeline_pipelines", response)
			respondWithJSON(w, 200, response)
		}
	}
}
