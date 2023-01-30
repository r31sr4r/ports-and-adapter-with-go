package db_test

import (
	"database/sql"
	"log"
	"testing"
	"github.com/stretchr/testify/require"
	_ "github.com/mattn/go-sqlite3"
	"github.com/r31sr4r/go-ports-and-adapters/adapters/db"
	"github.com/r31sr4r/go-ports-and-adapters/application"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		id string,
		name string,
		status string,
		price float
	);`

	stmt, err := Db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	product := `INSERT INTO products (id, name, status, price) VALUES ("123", "test", "disabled", 10.00);`

	stmt, err := Db.Prepare(product)
	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("123")
	require.Nil(t, err)
	require.Equal(t, "123", product.GetID())
	require.Equal(t, "test", product.GetName())
	require.Equal(t, "disabled", product.GetStatus())
	require.Equal(t, 10.00, product.GetPrice())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product 1"
	product.Price = 25

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = "enabled"
	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())
	
}