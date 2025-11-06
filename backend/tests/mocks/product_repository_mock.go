package mocks

import (
	"backend/domain"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct{ mock.Mock }

func (m *ProductRepositoryMock) GetByID(id uint) (*domain.Product, error) {
	args := m.Called(id)
	if v := args.Get(0); v != nil {
		return v.(*domain.Product), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *ProductRepositoryMock) SearchProducts(search string) ([]domain.Product, error) {
	args := m.Called(search)
	if v := args.Get(0); v != nil {
		return v.([]domain.Product), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *ProductRepositoryMock) CreateProduct(p *domain.Product) error {
	args := m.Called(p)
	return args.Error(0)
}

// --- mock del servicio para tests de controller ---
type ProductServiceMock struct{ mock.Mock }

func (m *ProductServiceMock) SearchProducts(s string) ([]domain.Product, error) {
	args := m.Called(s)
	if v := args.Get(0); v != nil {
		return v.([]domain.Product), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *ProductServiceMock) GetProduct(id uint) (*domain.Product, error) {
	args := m.Called(id)
	if v := args.Get(0); v != nil {
		return v.(*domain.Product), args.Error(1)
	}
	return nil, args.Error(1)
}
func (m *ProductServiceMock) CreateProduct(p domain.Product) error {
	args := m.Called(p)
	return args.Error(0)
}
