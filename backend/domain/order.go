package domain

// Order representa una compra confirmada
type Order struct {
	ID     uint        `json:"id"`
	UserID uint        `json:"user_id"`
	Total  float64     `json:"total"`
	Status string      `json:"status"`
	Items  []OrderItem `json:"items"`
}

// OrderItem representa un producto dentro de una orden
type OrderItem struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	Product   Product `json:"product"`
}
