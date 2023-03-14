package handlers

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"time"

	"github.com/tailwarden/komiser/models"
)

type Row struct {
	ID       string   `csv:"id"`
	Provider string   `csv:"provider"`
	Account  string   `csv:"account"`
	Name     string   `csv:"name"`
	Region   string   `csv:"region"`
	Tags     []string `csv:"tags"`
	Cost     float64  `csv:"cost"`
}

func (handler *ApiHandler) DownloadInventoryCSV(w http.ResponseWriter, r *http.Request) {
	file, err := os.Create("/tmp/export.csv")
	if err != nil {
		respondWithError(w, 500, "Could not create file at /tmp")
		return
	}

	defer file.Close()
	defer os.Remove("/tmp/export.csv")

	fw := bufio.NewWriter(file)
	csvWriter := csv.NewWriter(fw)

	var resources []models.Resource
	err = handler.db.NewSelect().Model((*models.Resource)(nil)).Scan(handler.ctx, &resources)
	if err != nil {
		respondWithError(w, 500, "Could not read from DB")
		return
	}

	header := []string{"id", "provider", "account", "name", "region", "tags", "cost"}
	if err := csvWriter.Write(header); err != nil {
		respondWithError(w, 500, "Could not write CSV")
		return
	}

	for _, record := range resources {
		tags, err := json.Marshal(record.Tags)
		if err != nil {
			log.Fatalf("Could not marshal tags")
		}

		row := []string{
			record.ResourceId, record.Provider, record.Account, record.Name, record.Region, string(tags), fmt.Sprintf("%2.f", record.Cost),
		}
		if err := csvWriter.Write(row); err != nil {
			respondWithError(w, 500, "Could not write CSV")
			return
		}
	}

	cd := mime.FormatMediaType("attachment", map[string]string{"filename": "export.csv"})
	w.Header().Set("Content-Disposition", cd)
	w.Header().Set("Content-Type", "application/octet-stream")
	http.ServeContent(w, r, "export.csv", time.Now(), file)
}
