package controllers

import (
	"backend/dao"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 📦 Obtener carrito del usuario (GET /api/cart)
func GetCart(c *gin.Context) {
	email, _ := c.Get("email")

	// Buscar usuario por email (domain.User)
	user, err := dao.GetUserByEmail(email.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Obtener o crear el carrito (domain.Cart)
	cart, err := dao.GetOrCreateCartByUserID(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener el carrito"})
		return
	}

	c.JSON(http.StatusOK, cart)
}

// ➕ Agregar producto al carrito (POST /api/cart/add)
func AddToCart(c *gin.Context) {
	email, _ := c.Get("email")

	// Obtener usuario por email
	user, err := dao.GetUserByEmail(email.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Parsear el JSON recibido
	var body struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON inválido"})
		return
	}

	// Agregar al carrito
	err = dao.AddToCart(user.ID, body.ProductID, body.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo agregar el producto al carrito"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto agregado al carrito"})
}

// ❌ Eliminar un producto del carrito (DELETE /api/cart/remove)
func RemoveFromCart(c *gin.Context) {
	email, _ := c.Get("email")

	user, err := dao.GetUserByEmail(email.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	var body struct {
		ProductID uint `json:"product_id"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON inválido"})
		return
	}

	err = dao.RemoveFromCart(user.ID, body.ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el producto del carrito"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado del carrito"})
}

// 🧹 Vaciar carrito (DELETE /api/cart/clear)
func ClearCart(c *gin.Context) {
	email, _ := c.Get("email")

	user, err := dao.GetUserByEmail(email.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	err = dao.ClearCart(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo vaciar el carrito"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Carrito vaciado correctamente"})
}
