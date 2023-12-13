package timeseriesdatabase

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"medsecurity/config"
	insqlx "medsecurity/pkg/db/sqlx"
)

func InitTimeSeriesDatabase(config config.DatabaseConfig) (db map[string]*sqlx.DB, err error) {
	db = make(map[string]*sqlx.DB)

	for idx, dsn := range config.TimeSeriesStorageDSN {
		key := config.TimeSeriesKey[idx]

		db[key], err = insqlx.InitSQLX(dsn)
		if err != nil {
			return db, fmt.Errorf("[ERROR] failed on creating connection to database with key: %s, error: %s", key, err.Error())
		}

		fmt.Println("[INFO] successfully created connection to database with key: ", key)
	}

	return db, nil
}
