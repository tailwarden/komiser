package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		// Set default value for tags in resources table, because sometimes it was 'null'
		// The Resources model was updated to set the default value to []string{}, but all older instances
		// of Komiser didn't have that default value set, so we need to update the database
		_, _ = db.NewUpdate().
			Table("resources").
			Set("tags = ?", []string{}).
			Where("tags = 'null'").
			Exec(ctx)

		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		// No rollback needed
		return nil
	})
}
