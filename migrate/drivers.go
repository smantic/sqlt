package migrate

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
)

const (
	DRIVER_POSTGRES = "postgres"
	DRIVER_MYSQL    = "mysql"
)

func open(driver, dsn string) (*sql.DB, error) {

	var (
		db  *sql.DB
		err error
	)

	//if dsn == "" {
	//	return nil, fmt.Errof("undefined behavior to provide an empty dsn")
	//}

	if driver != "" {
		return sql.Open(driver, dsn)
	}

	pq.ParseURL()
	conn, err := pq.NewConnector(dsn)
	if err == nil {
		return sql.OpenDB(conn), nil
	}

	mysqlc, err := mysql.ParseDSN(dsn)
	if err == nil {
		conn, err := mysql.NewConnector(mysqlc)
		if err != nil {
			return sql.OpenDB(conn), nil
		}
	}

	db, err = sql.Open("mysql", dsn)
	if db != nil {
		return db, err
	}

	return db, err
}
