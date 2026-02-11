package repository

import (
	"database/sql"
	"ecommerce/internal/model"
)

type ProductRepository struct {
	DB *sql.DB
}

func (r *ProductRepository) Create(p model.Product) error {
	_, err := r.DB.Exec(`
	INSERT INTO products(name, price, stock, size, color, category_id)
	VALUES($1,$2,$3,$4,$5,$6)`,
		p.Name, p.Price, p.Stock, p.Size, p.Color, p.CategoryID)
	return err
}

func (r *ProductRepository) GetAll() ([]model.Product, error) {
	rows, err := r.DB.Query(`
	SELECT id, name, price, stock, size, color, category_id FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Product
	for rows.Next() {
		var p model.Product
		rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.Size, &p.Color, &p.CategoryID)
		list = append(list, p)
	}
	return list, nil
}

func (r *ProductRepository) Update(id int64, p model.Product) error {
	_, err := r.DB.Exec(`
	UPDATE products SET name=$1, price=$2, stock=$3, size=$4, color=$5, category_id=$6
	WHERE id=$7`,
		p.Name, p.Price, p.Stock, p.Size, p.Color, p.CategoryID, id)
	return err
}

func (r *ProductRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM products WHERE id=$1", id)
	return err
}
