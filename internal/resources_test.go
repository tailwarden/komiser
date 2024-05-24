package internal

import (
	"context"
	"io"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/tailwarden/komiser/internal/config"
	"github.com/tailwarden/komiser/utils"
)

// BenchmarkFactorial benchmarks the Factorial function.
func BenchmarkFetchResources(b *testing.B) {
	// Setup
	ctx := context.TODO()
	log.SetOutput(io.Discard)
	analytics.Init()
	cfg, clients, accounts, err := config.Load("/workspaces/komiser/config.toml", false, analytics)
	if err != nil {
		b.Fatalf("Error during config setup: %v", err)
	}
	db, err = utils.SetupDBConnection(cfg)
	if err != nil {
		b.Fatalf("Error during DB setup: %v", err)
	}
	err = utils.SetupSchema(db, cfg, accounts)
	if err != nil {
		b.Fatalf("Error during schema setup: %v", err)
	}

	// The benchmark function will run b.N times
	for i := 0; i < b.N; i++ {
		fetchResources(ctx, clients, []string{}, false)
	}
}
