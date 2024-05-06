package sql

import (
	"database/sql"
)

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
