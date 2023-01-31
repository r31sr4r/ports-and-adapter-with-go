package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	db2 "github.com/r31sr4r/go-ports-and-adapters/adapters/db"
	"github.com/r31sr4r/go-ports-and-adapters/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product 1", 10)

	productService.Enable(product)
}