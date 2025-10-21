package domain

// Cart representa el carrito del usuario (vista JSON)
type Cart struct {
	ID     uint       `json:"id"`
	UserID uint       `json:"user_id"`
	Items  []CartItem `json:"items"`
}

// CartItem representa un producto dentro del carrito
type CartItem struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Product   Product `json:"product"`
}
