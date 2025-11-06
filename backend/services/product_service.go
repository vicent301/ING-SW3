package services

import (
	"errors"

	"backend/dao"
	"backend/domain"
)

type ProductService struct{ repo dao.ProductRepository }

func NewProductService(r dao.ProductRepository) *ProductService { return &ProductService{repo: r} }

func (s *ProductService) GetProduct(id uint) (*domain.Product, error) {
	if id == 0 {
		return nil, errors.New("ID inválido")
	}
	return s.repo.GetByID(id)
}

func (s *ProductService) SearchProducts(search string) ([]domain.Product, error) {
	return s.repo.SearchProducts(search)
}

func (s *ProductService) CreateProduct(p domain.Product) error {
	return s.repo.CreateProduct(&p)
}
