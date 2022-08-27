package drivers

import (
	"database/sql"
	"fmt"
	"strings"
)

const (
	DRIVER_POSTGRES = "postgres"
	DRIVER_MYSQL    = "mysql"
)

// Open will take a DSN
// we do a special thing and expect the user to provide the scheme on the DSN URL
//
//	scheme://u:pass@protocol(address)/dbname
//
// even though this is not a standard exception across drivers.
// scheme is bare minimun that must be provided.
func Open(dsn string) (*sql.DB, error) {

	if dsn == "" {
		// cant be completely empty. must at least have the scheme.
		return nil, fmt.Errorf("empty DSN")
	}

	uri := strings.Split(dsn, "://")
	if len(uri) == 1 || uri[0] == "" {
		return nil, fmt.Errorf("expected a scheme on the DSN")
	}

	return open(uri[0], dsn)
}

func open(driver, dsn string) (*sql.DB, error) {

	var (
		db  *sql.DB
		err error
	)

	switch driver {
	case DRIVER_POSTGRES:
		db, err = sql.Open(DRIVER_POSTGRES, dsn)
	case DRIVER_MYSQL:
		// mysql expects no scheme in the dsn.
		dsn = dsn[8:] // mysql://dsn
		db, err = sql.Open(DRIVER_MYSQL, dsn)
	default:
		err = fmt.Errorf("unsupported driver")
	}

	return db, err
}
