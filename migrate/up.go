package migrate

import (
	"context"
	"database/sql"
	"fmt"
	"io"
)

// Up runs the up migrations.
// pass in the sql to io.Reader
func Up(ctx context.Context, dsn string, r io.Reader) error {
	db, err := open("", dsn)
	if err != nil {
		return err
	}

	return UpWithDB(ctx, db, r)
}

// UpWithDB is useful for when you want to provide the DB connection.
func UpWithDB(ctx context.Context, db *sql.DB, r io.Reader) error {

	err := db.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	//use to rollback everything if we get sent an interrupt.
	globaltx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	done := make(chan struct{})
	go func() {

		apply(ctx, db, r)
		if err != nil {
			globaltx.Rollback()
		}

		close(done)
	}()

	select {
	case <-ctx.Done():
		globalRollback(globaltx)
		return ctx.Err()
	case <-done:
		// migrations are done.

		err := globaltx.Commit()
		if err != nil {
			globalRollback(globaltx)
			return fmt.Errorf("failed to commit transaction! %w", err)
		}
	}

	return nil
}

func globalRollback(tx *sql.Tx) {
	fmt.Printf("rolling back all changes!!")
	err := tx.Rollback()
	if err != nil {
		fmt.Printf("err rolling back: %s", err)
	}
}
