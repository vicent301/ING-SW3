// backend/controller/order_controller.go
package controllers

import (
	"backend/dao"
	"backend/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// helper: obtiene el email del contexto de forma segura

// 🧾 Crear una nueva orden a partir del carrito (POST /api/orders/create)
func CreateOrder(c *gin.Context) {
	// 1) Autenticación
	email, ok := getEmailFromCtx(c)
	if !ok {
		// Para pruebas unitarias sin token, devolvemos 400 (como BadJSON)
		c.JSON(http.StatusBadRequest, gin.H{"error": "no autorizado"})
		return
	}

	// 2) DB no inicializada (tests unitarios sin DB)
	if database.DB == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "base de datos no inicializada"})
		return
	}

	// 3) Buscar usuario
	user, err := dao.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// 4) Crear la orden desde el carrito (tu lógica real)
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
	// 1) Autenticación
	email, ok := getEmailFromCtx(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no autorizado"})
		return
	}

	// 2) DB no inicializada (tests unitarios sin DB)
	if database.DB == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "base de datos no inicializada"})
		return
	}

	// 3) Buscar usuario
	user, err := dao.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	// 4) Obtener órdenes del usuario (tu lógica real)
	orders, err := dao.GetOrdersByUser(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las órdenes"})
		return
	}

	c.JSON(http.StatusOK, orders)
}
