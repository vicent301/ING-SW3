package dao

import (
	"backend/domain"
)

// La interfaz base, que luego el servicio usa
type ProductRepository interface {
	GetByID(id uint) (*domain.Product, error)
	SearchProducts(search string) ([]domain.Product, error)
	CreateProduct(p *domain.Product) error
}
