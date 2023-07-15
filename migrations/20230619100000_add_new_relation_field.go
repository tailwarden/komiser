package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		// adding new column relation for migration 
		db.ExecContext(ctx, "ALTER TABLE resources ADD COLUMN relations JSONB DEFAULT '[]'::jsonb;")
		
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		// No rollback needed
		return nil
	})
}
