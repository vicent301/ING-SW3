package dao

import (
	"backend/database"
	"backend/domain"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Stock       int
	ImageURL    string
}

// AutoMigrar productos
func AutoMigrateProduct() {
	database.DB.AutoMigrate(&Product{})
}

// Crear producto
func CreateProduct(p domain.Product) error {
	product := Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		Stock:       p.Stock,
		ImageURL:    p.ImageURL,
	}
	return database.DB.Create(&product).Error
}

// Obtener todos los productos
func GetAllProducts() ([]domain.Product, error) {
	var entities []Product
	err := database.DB.Find(&entities).Error
	if err != nil {
		return nil, err
	}

	var products []domain.Product
	for _, e := range entities {
		products = append(products, domain.Product{
			ID:          e.ID,
			Name:        e.Name,
			Description: e.Description,
			Price:       e.Price,
			Stock:       e.Stock,
			ImageURL:    e.ImageURL,
		})
	}
	return products, nil
}

// Obtener producto por ID
func GetProductByID(id uint) (*domain.Product, error) {
	var e Product
	if err := database.DB.First(&e, id).Error; err != nil {
		return nil, err
	}
	return &domain.Product{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Price:       e.Price,
		Stock:       e.Stock,
		ImageURL:    e.ImageURL,
	}, nil
}

// 🔍 Buscar productos por palabra clave
func SearchProducts(keyword string) ([]domain.Product, error) {
	var entities []Product

	query := database.DB

	if keyword != "" {
		likePattern := "%" + keyword + "%"
		query = query.Where("name LIKE ? OR description LIKE ?", likePattern, likePattern)
	}

	if err := query.Find(&entities).Error; err != nil {
		return nil, err
	}

	var products []domain.Product
	for _, e := range entities {
		products = append(products, domain.Product{
			ID:          e.ID,
			Name:        e.Name,
			Description: e.Description,
			Price:       e.Price,
			Stock:       e.Stock,
			ImageURL:    e.ImageURL,
		})
	}

	return products, nil
}
