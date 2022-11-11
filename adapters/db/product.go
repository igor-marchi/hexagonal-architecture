package db

import (
	"database/sql"

	"github.com/igor-marchi/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (application.IProductInterface, error) {
	var product application.Product
	stmt, err := p.db.Prepare("select id, name, price, status from product where id = ?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDb) Save(product application.IProductInterface) (application.IProductInterface, error) {
	var rows int
	p.db.QueryRow("select id from product where id=?", product.GetId()).Scan(&rows)
	if rows == 0 {
		_, err := p.create(product)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := p.update(product)
		if err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (p *ProductDb) create(product application.IProductInterface) (application.IProductInterface, error) {
	stmt, err := p.db.Prepare(`INSERT INTO product(id, name, price, status) VALUES (?,?,?,?)`)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product application.IProductInterface) (application.IProductInterface, error) {
	_, err := p.db.Exec(
		"update products set name = ?, price = ?, status = ? where id = ?",
		product.GetName(), product.GetPrice(), product.GetStatus(), product.GetId())
	if err != nil {
		return nil, err
	}

	return product, err
}
