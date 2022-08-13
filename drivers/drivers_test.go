package drivers

import "testing"

func TestOpen(t *testing.T) {

	dsns := []string{
		"mysql://username:password@protocol(address)/dbname?param=value",
		"mysql://username:password@protocol(address)/dbname?param=value&columnsWithAlias=true",
		"mysql://username:password@protocol(address)/dbname?param=value&columnsWithAlias=true",
		"mysql:///dbname",
		"mysql://@/",
		"mysql://tcp(127.0.0.1)/dbname",
		"mysql://tcp(127.0.0.1:1234)/dbname",
		//"mysql://tcp(de:ad:be:ef::ca:fe)/dbname",
		"mysql://",
		"postgres://[::1]:1234",
		"postgres://username:top%20secret@hostname.remote:1234/database",
		"postgres://",
	}

	for _, dsn := range dsns {
		dsn := dsn
		t.Run(dsn, func(t *testing.T) {
			db, err := Open(dsn)
			if err != nil {
				t.Error(err)
			}

			if db == nil {
				t.Error("nil DB")
			}

		})
	}
}
