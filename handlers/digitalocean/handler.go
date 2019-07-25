package digitalocean

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/digitalocean/godo"
	. "github.com/mlabouardy/komiser/services/cache"
	. "github.com/mlabouardy/komiser/services/digitalocean"
	"golang.org/x/oauth2"
)

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

type DigitalOceanHandler struct {
	cache        Cache
	client       *godo.Client
	digitalocean DigitalOcean
}

func NewDigitalOceanHandler(cache Cache) *DigitalOceanHandler {
	tokenSource := &TokenSource{
		AccessToken: os.Getenv("DIGITALOCEAN_ACCESS_TOKEN"),
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	digitaloceanHandler := DigitalOceanHandler{
		cache:        cache,
		client:       client,
		digitalocean: DigitalOcean{},
	}
	return &digitaloceanHandler
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
