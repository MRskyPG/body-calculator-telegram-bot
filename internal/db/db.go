package db

import (
	"database/sql"
	"fmt"
	"os"
)

type dbConfig struct {
	PgDB, PgUser, PgPassword string
}

func GetParameters() *dbConfig {
	return &dbConfig{
		PgDB:       os.Getenv("POSTGRES_DATABASE"),
		PgUser:     os.Getenv("POSTGRES_USER"),
		PgPassword: os.Getenv("POSTGRES_PASSWORD"),
	}
}

func GetDB() (*sql.DB, error) {
	var err error
	var db *sql.DB

	cfg := GetParameters()
	connStr := fmt.Sprintf("port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"5439", cfg.PgUser, cfg.PgPassword, cfg.PgDB)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
