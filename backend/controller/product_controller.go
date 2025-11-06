package controllers

import (
	"backend/domain"
	"backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductController struct{ svc services.ProductServicePort }

func NewProductController(s services.ProductServicePort) *ProductController {
	return &ProductController{svc: s}
}

func (pc *ProductController) GetProducts(c *gin.Context) {
	search := c.Query("search")
	products, err := pc.svc.SearchProducts(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}
	p, err := pc.svc.GetProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
		return
	}
	c.JSON(http.StatusOK, p)
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var p domain.Product
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := pc.svc.CreateProduct(p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Producto creado correctamente"})
}
