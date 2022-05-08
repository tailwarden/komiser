package civo

import (
	"encoding/json"
	"net/http"

	. "github.com/mlabouardy/komiser/services/cache"
	. "github.com/mlabouardy/komiser/services/civo"
)

type CivoHandler struct {
	cache Cache
	civo  Civo
}

func (handler *CivoHandler) GetCivoHandler() Civo {
	return handler.civo
}

func NewCivoHandler(cache Cache) *CivoHandler {
	civoHandler := CivoHandler{
		cache: cache,
		civo:  Civo{},
	}
	return &civoHandler
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
