package dao

import (
	"backend/database"
	"backend/domain"
	"errors"
	"time"

	"gorm.io/gorm"
)

type OrderEntity struct {
	gorm.Model
	UserID    uint
	Total     float64
	Status    string
	CreatedAt time.Time
	Items     []OrderItemEntity `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
}

func (OrderEntity) TableName() string {
	return "orders"
}

type OrderItemEntity struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Quantity  int
	Price     float64
	Product   Product `gorm:"foreignKey:ProductID"`
}

func (OrderItemEntity) TableName() string {
	return "order_items"
}

// 🔧 AutoMigrar tablas de órdenes
func AutoMigrateOrder() {
	database.DB.AutoMigrate(&OrderEntity{}, &OrderItemEntity{})
}

// 💾 Crear una nueva orden a partir del carrito
func CreateOrderFromCart(userID uint) (*domain.Order, error) {
	var cart CartEntity
	if err := database.DB.Preload("Items.Product").Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return nil, errors.New("no se encontró el carrito del usuario")
	}

	if len(cart.Items) == 0 {
		return nil, errors.New("el carrito está vacío")
	}

	var total float64
	for _, item := range cart.Items {
		total += float64(item.Quantity) * item.Product.Price
	}

	// Crear la orden
	order := OrderEntity{
		UserID: userID,
		Total:  total,
		Status: "PENDIENTE",
	}

	if err := database.DB.Create(&order).Error; err != nil {
		return nil, err
	}

	// Crear los ítems de la orden y descontar stock
	for _, item := range cart.Items {
		orderItem := OrderItemEntity{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Product.Price,
		}
		database.DB.Create(&orderItem)

		// Descontar stock
		database.DB.Model(&Product{}).
			Where("id = ?", item.ProductID).
			UpdateColumn("stock", gorm.Expr("stock - ?", item.Quantity))
	}

	// Vaciar carrito después de crear orden
	database.DB.Where("cart_id = ?", cart.ID).Delete(&CartItemEntity{})

	// Retornar como dominio
	return &domain.Order{
		ID:     order.ID,
		UserID: order.UserID,
		Total:  total,
		Status: order.Status,
	}, nil
}

// 🔍 Obtener todas las órdenes de un usuario
func GetOrdersByUser(userID uint) ([]domain.Order, error) {
	var orders []OrderEntity
	if err := database.DB.Preload("Items.Product").
		Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}

	var result []domain.Order
	for _, o := range orders {
		order := domain.Order{
			ID:     o.ID,
			UserID: o.UserID,
			Total:  o.Total,
			Status: o.Status,
		}
		for _, i := range o.Items {
			order.Items = append(order.Items, domain.OrderItem{
				ID:        i.ID,
				ProductID: i.ProductID,
				Quantity:  i.Quantity,
				Price:     i.Price,
				Product: domain.Product{
					ID:       i.Product.ID,
					Name:     i.Product.Name,
					Price:    i.Product.Price,
					ImageURL: i.Product.ImageURL,
				},
			})
		}
		result = append(result, order)
	}
	return result, nil
}
