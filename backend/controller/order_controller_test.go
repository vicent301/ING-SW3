package controllers

import (
	"backend/dao"
	"backend/database"
	"backend/domain"
	"backend/testutil"
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateOrder_BadJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(withUser(1))
	r.POST("/orders/create", CreateOrder)

	req, _ := http.NewRequest(http.MethodPost, "/orders/create", strings.NewReader(`{mal json}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetOrders_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/orders", GetOrders) // sin userID

	req, _ := http.NewRequest(http.MethodGet, "/orders", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestOrders_CreateAndList_OK(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	// seed user
	if err := dao.CreateUser(domain.User{
		Name:     "Vic",
		Email:    "order@example.com",
		Password: "x",
	}); err != nil {
		t.Fatalf("seed user: %v", err)
	}
	// seed product
	type Product struct {
		ID          uint `gorm:"primaryKey"`
		Name        string
		Description string
		Price       float64
		Stock       int
		ImageURL    string
	}
	if err := database.DB.AutoMigrate(&Product{}); err != nil {
		t.Fatal(err)
	}
	p := Product{Name: "Zapa X", Price: 99}
	if err := database.DB.Create(&p).Error; err != nil {
		t.Fatal(err)
	}

	// Agregar al carrito antes de crear orden
	if err := dao.AddToCart(1, p.ID, 3); err != nil { // ajusta userID si no es 1
		t.Fatalf("add to cart: %v", err)
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(withEmail("order@example.com"))
	r.POST("/api/orders/create", CreateOrder)
	r.GET("/api/orders", GetOrders)

	// CreateOrder
	req := httptest.NewRequest(http.MethodPost, "/api/orders/create", bytes.NewReader([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("create order esper 200, got %d (%s)", w.Code, w.Body.String())
	}

	// List
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/orders", nil))
	if w.Code != http.StatusOK {
		t.Fatalf("get orders esper 200, got %d (%s)", w.Code, w.Body.String())
	}
}
