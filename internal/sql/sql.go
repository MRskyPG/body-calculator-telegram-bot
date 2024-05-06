package sql

import (
	"database/sql"
	"fmt"
)

type Product struct {
	ProductName   string
	Calories      float64
	Proteins      float64
	Fats          float64
	Carbohydrates float64
	ProductType   string
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

func GetProductsByType(db *sql.DB, pType string) ([]Product, error) {
	var products []Product
	str := fmt.Sprintf("SELECT * FROM get_products_by_type('%s');", pType)
	rows, err := db.Query(str)
	defer rows.Close()
	if err != nil {
		return products, err
	}

	for rows.Next() {
		var tmpProduct Product
		_ = rows.Scan(&tmpProduct.ProductName, &tmpProduct.Calories, &tmpProduct.Proteins, &tmpProduct.Fats, &tmpProduct.Carbohydrates, &tmpProduct.ProductType)
		products = append(products, tmpProduct)
	}
	return products, nil
}
