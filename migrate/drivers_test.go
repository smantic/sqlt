package migrate

import "testing"

func TestOpen(t *testing.T) {

	var mysqlDSNs []string = []string{
		"username:password@protocol(address)/dbname?param=value",
		"username:password@protocol(address)/dbname?param=value&columnsWithAlias=true",
		"mysql://username:password@protocol(address)/dbname?param=value&columnsWithAlias=true",
		"/dbname",
		"@/",
		"tcp(127.0.0.1)/dbname",
		"tcp(de:ad:be:ef::ca:fe)/dbname",
		"",
	}

	for _, dsn := range mysqlDSNs {
		db, err := open("", dsn)
		if db != nil {
			t.Errorf("db was not nil")
			if err != nil {
				t.Logf("error: %s", err.Error())
			}
		}
	}

}
