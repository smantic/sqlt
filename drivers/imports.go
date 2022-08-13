package migrate

// just to ensure we are always importing the drivers.
import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)
