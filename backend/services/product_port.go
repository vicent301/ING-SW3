package services

import "backend/domain"

type ProductServicePort interface {
	SearchProducts(search string) ([]domain.Product, error)
	GetProduct(id uint) (*domain.Product, error)
	CreateProduct(p domain.Product) error
}
