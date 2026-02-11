package repository

import (
	"database/sql"
	"ecommerce/internal/model"
)

type CategoryRepository struct {
	DB *sql.DB
}

func (r *CategoryRepository) Create(c model.Category) error {
	_, err := r.DB.Exec("INSERT INTO categories(name) VALUES($1)", c.Name)
	return err
}

func (r *CategoryRepository) GetAll() ([]model.Category, error) {
	rows, err := r.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.Category
	for rows.Next() {
		var c model.Category
		rows.Scan(&c.ID, &c.Name)
		list = append(list, c)
	}
	return list, nil
}

func (r *CategoryRepository) Update(id int64, name string) error {
	_, err := r.DB.Exec("UPDATE categories SET name=$1 WHERE id=$2", name, id)
	return err
}

func (r *CategoryRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM categories WHERE id=$1", id)
	return err
}
