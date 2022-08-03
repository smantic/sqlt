package migrate

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io"
)

// Up runs the up migrations.
// pass in the sql to io.Reader
func Up(ctx context.Context, dsn string, r io.Reader) error {

	db, err := sql.Open("", dsn)
	if err != nil {
		return fmt.Errorf("failed to open the database: %w", err)
	}
	defer db.Close()

	return UpWithDB(ctx, db, r)
}

func globalRollback(tx *sql.Tx) {
	fmt.Printf("rolling back!!")
	err := tx.Rollback()
	if err != nil {
		fmt.Printf("err rolling back: %s", err)
	}
}

// UpWithDB is useful for when you want to provide the DB connection.
func UpWithDB(ctx context.Context, db *sql.DB, r io.Reader) error {

	err := db.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	//globaltx use to rollback everything if we get sent an interrupt.
	globaltx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	done := make(chan struct{})
	go func() {

		err := errors.New("migrations")
		// do migrations.
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
		err := globaltx.Commit()
		if err != nil {
			return err
		}
	}

	return nil
}
