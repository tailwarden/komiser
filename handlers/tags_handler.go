package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	. "github.com/tailwarden/komiser/models"
)

func (handler *ApiHandler) BulkUpdateTagsHandler(w http.ResponseWriter, r *http.Request) {
	var input BulkUpdateTag

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resource := Resource{Tags: input.Tags}

	for _, resourceId := range input.Resources {
		_, err = handler.db.NewUpdate().Model(&resource).Column("tags").Where("id = ?", resourceId).Exec(handler.ctx)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Error while updating tags")
			return
		}
	}

	respondWithJSON(w, 200, "Tags has been successfuly updated")
}

func (handler *ApiHandler) UpdateTagsHandler(w http.ResponseWriter, r *http.Request) {
	tags := make([]Tag, 0)

	vars := mux.Vars(r)
	resourceId, ok := vars["id"]

	if !ok {
		respondWithError(w, http.StatusBadRequest, "Resource id is missing")
		return
	}

	id, err := strconv.Atoi(resourceId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Resource id should be an integer")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&tags)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resource := Resource{Tags: tags}

	_, err = handler.db.NewUpdate().Model(&resource).Column("tags").Where("id = ?", id).Exec(handler.ctx)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error while updating tags")
		return
	}

	respondWithJSON(w, 200, "Tags has been successfuly updated")
}
