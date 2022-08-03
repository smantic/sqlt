package migrate

import (
	"context"
	"database/sql"
	"io"
)

// apply migrations
func apply(ctx context.Context, db *sql.DB, r io.Reader) error {

	queries, err := QueriesFrom(r)
	if err != nil {
		return err
	}

	for _, q := range queries {
		result, err := db.ExecContext(ctx, q)
		_ = result
		if err != nil {
			return err
		}
	}

	return nil
}
