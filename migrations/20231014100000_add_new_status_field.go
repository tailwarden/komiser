package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		// adding new column relation for migration
		_, _ = db.ExecContext(ctx, "ALTER TABLE accounts ADD COLUMN status TEXT;")
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		// No rollback needed
		return nil
	})
}
