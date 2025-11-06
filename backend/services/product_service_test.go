package services

import (
	"backend/domain"
	"backend/tests/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetProduct_OK(t *testing.T) {
	repo := new(mocks.ProductRepositoryMock)
	svc := NewProductService(repo)
	expected := &domain.Product{ID: 1, Name: "Zapa"}

	repo.On("GetByID", uint(1)).Return(expected, nil)

	got, err := svc.GetProduct(1)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)
	repo.AssertExpectations(t)
}

func TestGetProduct_InvalidID(t *testing.T) {
	repo := new(mocks.ProductRepositoryMock)
	svc := NewProductService(repo)

	got, err := svc.GetProduct(0)
	assert.Nil(t, got)
	assert.Error(t, err)
}
func TestCreateProduct_RepoError(t *testing.T) {
	repo := new(mocks.ProductRepositoryMock)
	svc := NewProductService(repo)
	p := domain.Product{Name: "X"}

	repo.On("CreateProduct", &p).Return(errors.New("db"))
	err := svc.CreateProduct(p)

	assert.Error(t, err)
	repo.AssertExpectations(t)
}

func TestSearchProducts_RepoError(t *testing.T) {
	repo := new(mocks.ProductRepositoryMock)
	svc := NewProductService(repo)

	repo.On("SearchProducts", "a").Return(nil, errors.New("db"))
	_, err := svc.SearchProducts("a")

	assert.Error(t, err)
	repo.AssertExpectations(t)
}
