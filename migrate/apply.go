package migrate

import (
	"context"
	"database/sql"
	"io"
)

// apply migrations.
func apply(ctx context.Context, db *sql.DB, r io.Reader) error {

	statements, err := StatmentsFrom(r)
	if err != nil {
		return err
	}

	for _, s := range statements {
		_, err := db.ExecContext(ctx, s)
		if err != nil {
			return err
		}
	}

	return nil
}
