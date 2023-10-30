package migrations

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		// adding new column relation for migration
		if db.Dialect().Name() == dialect.SQLite {
			_, _ = db.ExecContext(ctx, "ALTER TABLE resources ADD COLUMN relations TEXT DEFAULT '[]';")
		} else {
			_, _ = db.ExecContext(ctx, "ALTER TABLE resources ADD COLUMN relations JSONB DEFAULT '[]'::jsonb;")
		}
		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		// No rollback needed
		return nil
	})
}
