package service

import (
	"ecommerce/internal/model"
	"ecommerce/internal/repository"
)

type CategoryService struct {
	Repo *repository.CategoryRepository
}

func (s *CategoryService) Create(c model.Category) error {
	return s.Repo.Create(c)
}

func (s *CategoryService) GetAll() ([]model.Category, error) {
	return s.Repo.GetAll()
}

func (s *CategoryService) Update(id int64, name string) error {
	return s.Repo.Update(id, name)
}

func (s *CategoryService) Delete(id int64) error {
	return s.Repo.Delete(id)
}
