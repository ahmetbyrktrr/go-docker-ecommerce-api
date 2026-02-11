package service

import (
	"ecommerce/internal/model"
	"ecommerce/internal/repository"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

func (s *ProductService) Create(p model.Product) error {
	return s.Repo.Create(p)
}

func (s *ProductService) GetAll() ([]model.Product, error) {
	return s.Repo.GetAll()
}

func (s *ProductService) Update(id int64, p model.Product) error {
	return s.Repo.Update(id, p)
}

func (s *ProductService) Delete(id int64) error {
	return s.Repo.Delete(id)
}
