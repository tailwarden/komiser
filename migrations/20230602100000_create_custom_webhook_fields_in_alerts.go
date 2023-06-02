package migrations

import (
	"context"

	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.ExecContext(ctx, `
			SELECT is_slack, endpoint, secret FROM alerts LIMIT 1;
		`)

		if err != nil {
			_, err := db.ExecContext(ctx, `
				ALTER TABLE alerts
				ADD COLUMN is_slack BOOLEAN DEFAULT 0;
			`)
			if err != nil {
				return err
			}

			_, err = db.ExecContext(ctx, `
				ALTER TABLE alerts
				ADD COLUMN endpoint TEXT;
			`)
			if err != nil {
				return err
			}

			_, err = db.ExecContext(ctx, `
				ALTER TABLE alerts
				ADD COLUMN secret TEXT;
			`)
			if err != nil {
				return err
			}

			_, err = db.ExecContext(ctx, `
				UPDATE alerts
				SET is_slack = true;
			`)
			if err != nil {
				return err
			}
		}

		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		return nil
	})
}
