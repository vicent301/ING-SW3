package controllers

import (
	"backend/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 🧾 Crear una nueva orden a partir del carrito (POST /api/orders/create)
func CreateOrder(c *gin.Context) {
	email, _ := c.Get("email")

	// Buscar usuario por email
	user, err := dao.GetUserByEmail(email.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Crear la orden desde el carrito
	order, err := dao.CreateOrderFromCart(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Orden creada correctamente",
		"order":   order,
	})
}

// 📦 Obtener todas las órdenes del usuario (GET /api/orders)
func GetOrders(c *gin.Context) {
	email, _ := c.Get("email")

	// Buscar usuario por email
	user, err := dao.GetUserByEmail(email.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// Obtener las órdenes del usuario
	orders, err := dao.GetOrdersByUser(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las órdenes"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
