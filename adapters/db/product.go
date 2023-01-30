package db


import (
	"database/sql"
	"github.com/r31sr4r/go-ports-and-adapters/application"	
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("SELECT id, name, status, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Status, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	if product.GetID() == "" {
		return p.create(product)
	}
	return p.update(product)
}

func (p *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare("INSERT INTO products (id, name, status, price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetStatus(), product.GetPrice())
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare("UPDATE products SET name = ?, status = ?, price = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(product.GetName(), product.GetStatus(), product.GetPrice(), product.GetID())
	if err != nil {
		return nil, err
	}
	err = stmt.Close()
	if err != nil {
		return nil, err
	}
	return product, nil
}

