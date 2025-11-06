package dao

import (
	"backend/domain"
	"gorm.io/gorm"
)

type ProductGormRepository struct{ db *gorm.DB }

func NewProductGormRepository(db *gorm.DB) *ProductGormRepository {
	return &ProductGormRepository{db: db}
}

func (r *ProductGormRepository) GetByID(id uint) (*domain.Product, error) {
	var p domain.Product
	if err := r.db.First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductGormRepository) SearchProducts(search string) ([]domain.Product, error) {
	var list []domain.Product
	q := r.db.Model(&domain.Product{})
	if search != "" {
		q = q.Where("name LIKE ?", "%"+search+"%")
	}
	if err := q.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *ProductGormRepository) CreateProduct(p *domain.Product) error {
	return r.db.Create(p).Error
}
