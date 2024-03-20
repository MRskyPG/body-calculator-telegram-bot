package sql

import (
	"database/sql"
)

func GetDB() (*sql.DB, error) {
	var err error
	var db *sql.DB

	db, err = sql.Open("postgres", "user=postgres password=qwerty dbname=postgres port=5436 sslmode=disable")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
