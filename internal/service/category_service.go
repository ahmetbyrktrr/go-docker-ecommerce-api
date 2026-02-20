package service

import (
	"ecommerce/internal/model"
	"ecommerce/internal/repository"
)

type CategoryService struct {
	Repo        *repository.CategoryRepository
	ProductRepo *repository.ProductRepository
}

func (s *CategoryService) Create(c model.Category) error {
	return s.Repo.Create(c)
}

func (s *CategoryService) GetAll() ([]model.Category, error) {
	return s.Repo.GetAll()
}

func (s *CategoryService) Update(id int64, name string, products []model.Product) error {
	// Kategori adını güncelle
	err := s.Repo.Update(id, name)
	if err != nil {
		return err
	}

	// Mevcut product ID'lerini topla
	var productIDs []int64
	for _, p := range products {
		if p.ID > 0 {
			productIDs = append(productIDs, p.ID)
		}
	}

	// Kategoriye ait olan ama listede olmayan productları sil
	err = s.ProductRepo.DeleteByCategoryIDAndNotInList(id, productIDs)
	if err != nil {
		return err
	}

	// Yeni productları ekle veya mevcut olanları güncelle
	for _, p := range products {
		p.CategoryID = id
		if p.ID > 0 {
			// Mevcut productı güncelle
			err = s.ProductRepo.Update(p.ID, p)
		} else {
			// Yeni product ekle
			err = s.ProductRepo.Create(p)
		}
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *CategoryService) Delete(id int64) error {
	return s.Repo.Delete(id)
}
