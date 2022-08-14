package migrate

import (
	"context"
	"database/sql"
	"io"
)

// apply migrations
func apply(ctx context.Context, db *sql.DB, r io.Reader) error {

	statements, err := StatmentsFrom(r)
	if err != nil {
		return err
	}

	for _, s := range statements {
		result, err := db.ExecContext(ctx, s)
		_ = result
		if err != nil {
			return err
		}
	}

	return nil
}
