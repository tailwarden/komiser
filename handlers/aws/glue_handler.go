package aws

import (
	"net/http"
)

func (handler *AWSHandler) GlueGetCrawlersHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("glue_crawlers")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetGlueCrawlers(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "glue:GetCrawlers is missing")
		} else {
			handler.cache.Set("glue_crawlers", response)
			respondWithJSON(w, 200, response)
		}
	}
}

func (handler *AWSHandler) GlueGetJobsHandler(w http.ResponseWriter, r *http.Request) {
	response, found := handler.cache.Get("glue_jobs")
	if found {
		respondWithJSON(w, 200, response)
	} else {
		response, err := handler.aws.GetGlueJobs(handler.cfg)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "glue:GetJobs is missing")
		} else {
			handler.cache.Set("glue_jobs", response)
			respondWithJSON(w, 200, response)
		}
	}
}
