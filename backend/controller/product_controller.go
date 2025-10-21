package controllers

import (
	"backend/database"
	"backend/models"
	"net/http"
	_ "strconv"

	"github.com/gin-gonic/gin"
)

// Listar todos los productos
func GetProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

// Ver un producto por ID
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// Crear nuevo producto
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Producto creado correctamente", "product": product})
}

// Actualizar producto
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}

	var updated models.Product
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.Name = updated.Name
	product.Description = updated.Description
	product.Price = updated.Price
	product.Stock = updated.Stock
	product.ImageURL = updated.ImageURL

	database.DB.Save(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Producto actualizado", "product": product})
}

// Eliminar producto
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado"})
}
