package sql

import (
	"database/sql"
)

func GetDB() (*sql.DB, error) {
	var err error
	var db *sql.DB

	//check Makefile
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

func GetProductTypes(db *sql.DB) ([]string, error) {
	var productTypes []string
	str := "select * from get_all_product_types();"
	rows, err := db.Query(str)
	if err != nil {
		return productTypes, err
	}

	for rows.Next() {
		var tmp string
		_ = rows.Scan(&tmp)
		productTypes = append(productTypes, tmp)
	}
	return productTypes, nil
}
