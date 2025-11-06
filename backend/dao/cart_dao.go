package dao

import (
	"backend/database"
	"backend/domain"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type CartEntity struct {
	gorm.Model
	UserID uint
	Items  []CartItemEntity `gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE"`
}

func (CartEntity) TableName() string {
	return "carts"
}

type CartItemEntity struct {
	gorm.Model
	CartID    uint
	ProductID uint
	Quantity  int
	Product   Product `gorm:"foreignKey:ProductID"`
}

func (CartItemEntity) TableName() string {
	return "cart_items"
}

// 🔧 AutoMigrar tablas del carrito
func AutoMigrateCart() {
	database.DB.AutoMigrate(&CartEntity{}, &CartItemEntity{})
}

// 🔍 Obtener carrito de un usuario (crea uno si no existe)
func GetOrCreateCartByUserID(userID uint) (*domain.Cart, error) {
	var cartEntity CartEntity
	if err := database.DB.Preload("Items.Product").
		FirstOrCreate(&cartEntity, CartEntity{UserID: userID}).Error; err != nil {
		return nil, err
	}

	// Convertir a domain.Cart
	cart := domain.Cart{
		ID:     cartEntity.ID,
		UserID: cartEntity.UserID,
	}
	for _, item := range cartEntity.Items {
		cart.Items = append(cart.Items, domain.CartItem{
			ID:        item.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Product: domain.Product{
				ID:          item.Product.ID,
				Name:        item.Product.Name,
				Description: item.Product.Description,
				Price:       item.Product.Price,
				Stock:       item.Product.Stock,
				ImageURL:    item.Product.ImageURL,
			},
		})
	}
	return &cart, nil
}

// ➕ Agregar producto al carrito
func AddToCart(userID, productID uint, quantity int) error {
	var cart CartEntity
	if err := database.DB.FirstOrCreate(&cart, CartEntity{UserID: userID}).Error; err != nil {
		return err
	}

	// 🧩 VALIDAR que el producto exista antes de seguir
	var product Product
	if err := database.DB.First(&product, productID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("producto no encontrado (id=%d)", productID)
		}
		return err
	}

	// 🔎 Buscar si el producto ya está en el carrito
	var item CartItemEntity
	result := database.DB.Where("cart_id = ? AND product_id = ?", cart.ID, productID).First(&item)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		item = CartItemEntity{
			CartID:    cart.ID,
			ProductID: productID,
			Quantity:  quantity,
		}
		return database.DB.Create(&item).Error
	}

	item.Quantity += quantity
	return database.DB.Save(&item).Error
}

// ❌ Quitar producto del carrito
func RemoveFromCart(userID, productID uint) error {
	var cart CartEntity
	if err := database.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return err
	}
	return database.DB.Where("cart_id = ? AND product_id = ?", cart.ID, productID).Delete(&CartItemEntity{}).Error
}

// 🧹 Vaciar carrito
func ClearCart(userID uint) error {
	var cart CartEntity
	if err := database.DB.Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return err
	}
	return database.DB.Where("cart_id = ?", cart.ID).Delete(&CartItemEntity{}).Error
}
