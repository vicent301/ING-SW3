package controllers

import (
	"backend/dao"
	"backend/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 📦 GET /api/cart
func GetCart(c *gin.Context) {
	email, ok := getEmailFromCtx(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no autorizado"})
		return
	}
	if database.DB == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "base de datos no inicializada"})
		return
	}
	user, err := dao.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	cart, err := dao.GetOrCreateCartByUserID(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener el carrito"})
		return
	}

	c.JSON(http.StatusOK, cart)
}

// ➕ POST /api/cart/add
func AddToCart(c *gin.Context) {
	email, ok := getEmailFromCtx(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no autorizado"})
		return
	}
	if database.DB == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "base de datos no inicializada"})
		return
	}
	user, err := dao.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	var body struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato JSON inválido"})
		return
	}
	if body.ProductID == 0 || body.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product_id y quantity deben ser válidos"})
		return
	}

	if err := dao.AddToCart(user.ID, body.ProductID, body.Quantity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo agregar el producto al carrito"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto agregado al carrito"})
}

// ❌ DELETE /api/cart/remove
func RemoveFromCart(c *gin.Context) {
	email, ok := getEmailFromCtx(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no autorizado"})
		return
	}
	if database.DB == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "base de datos no inicializada"})
		return
	}
	user, err := dao.GetUserByEmail(email)
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
	if body.ProductID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product_id requerido"})
		return
	}

	if err := dao.RemoveFromCart(user.ID, body.ProductID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el producto del carrito"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado del carrito"})
}

// 🧹 DELETE /api/cart/clear
func ClearCart(c *gin.Context) {
	email, ok := getEmailFromCtx(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no autorizado"})
		return
	}
	if database.DB == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "base de datos no inicializada"})
		return
	}
	user, err := dao.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	if err := dao.ClearCart(user.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo vaciar el carrito"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Carrito vaciado correctamente"})
}
