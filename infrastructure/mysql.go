package infrastructure

import (
	"database/sql"
	"os"
)

func initMySQLDriver() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return db, nil
}
