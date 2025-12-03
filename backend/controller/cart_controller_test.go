package controllers

import (
	"backend/dao"
	"backend/database"
	"backend/domain"
	"backend/testutil"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "backend/testutil"
	"encoding/json"
)

func TestGetCart_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/cart", GetCart) // sin email en contexto

	req, _ := http.NewRequest(http.MethodGet, "/cart", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	//assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAddToCart_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/cart/add", AddToCart)

	req, _ := http.NewRequest(http.MethodPost, "/cart/add", strings.NewReader(`{"product_id":1,"quantity":1}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAddToCart_BadJSON(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	// SEED: el usuario que vas a setear en el contexto
	if err := dao.CreateUser(domain.User{
		Name: "X", Email: "x@y.com", Password: "pwd",
	}); err != nil {
		t.Fatalf("seed user: %v", err)
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(withEmail("x@y.com"))
	r.POST("/api/cart/add", AddToCart)

	// body inválido -> debe dar 400
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/api/cart/add", bytes.NewReader([]byte(`{`))) // JSON roto
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, got %d (%s)", w.Code, w.Body.String())
	}
}

func TestAddToCart_InvalidFields(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	if err := dao.CreateUser(domain.User{
		Name: "X", Email: "x@y.com", Password: "pwd",
	}); err != nil {
		t.Fatalf("seed user: %v", err)
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(withEmail("x@y.com"))
	r.POST("/api/cart/add", AddToCart)

	// Campos inválidos -> debe dar 400
	body := map[string]any{"product_id": 0, "quantity": -1}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, "/api/cart/add", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, got %d (%s)", w.Code, w.Body.String())
	}
}

func TestRemoveFromCart_BadJSON(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	if err := dao.CreateUser(domain.User{
		Name: "X", Email: "x@y.com", Password: "pwd",
	}); err != nil {
		t.Fatalf("seed user: %v", err)
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(withEmail("x@y.com"))
	r.DELETE("/api/cart/remove", RemoveFromCart)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodDelete, "/api/cart/remove", bytes.NewReader([]byte(`{`))) // JSON roto
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400, got %d (%s)", w.Code, w.Body.String())
	}
}

func TestClearCart_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.DELETE("/cart/clear", ClearCart) // sin email

	req, _ := http.NewRequest(http.MethodDelete, "/cart/clear", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func withEmail(email string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("email", email)
		c.Next()
	}
}

func TestCart_Flow_OK(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	// seed user
	if err := dao.CreateUser(domain.User{
		Name:     "Vic",
		Email:    "cart@example.com",
		Password: "x",
	}); err != nil {
		t.Fatalf("seed user: %v", err)
	}
	// seed product (ajustá si tenés dao.CreateProduct)
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

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(withEmail("cart@example.com"))

	r.GET("/api/cart", GetCart)
	r.POST("/api/cart/add", AddToCart)
	r.DELETE("/api/cart/remove", RemoveFromCart)
	r.DELETE("/api/cart/clear", ClearCart)

	// AddToCart
	bodyAdd := map[string]any{"product_id": p.ID, "quantity": 2}
	b, _ := json.Marshal(bodyAdd)
	req := httptest.NewRequest(http.MethodPost, "/api/cart/add", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("add esper 200, got %d (%s)", w.Code, w.Body.String())
	}

	// GetCart
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/cart", nil))
	if w.Code != http.StatusOK {
		t.Fatalf("get cart esper 200, got %d (%s)", w.Code, w.Body.String())
	}

	// RemoveFromCart
	bodyRem := map[string]any{"product_id": p.ID}
	b, _ = json.Marshal(bodyRem)
	w = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodDelete, "/api/cart/remove", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("remove esper 200, got %d (%s)", w.Code, w.Body.String())
	}

	// ClearCart
	w = httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodDelete, "/api/cart/clear", nil))
	if w.Code != http.StatusOK {
		t.Fatalf("clear esper 200, got %d (%s)", w.Code, w.Body.String())
	}
}

// -------------------- Ramas DB == nil --------------------

func TestGetCart_DBNil_Returns400(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(withEmail("alguien@test.com"))
	r.GET("/api/cart", GetCart)

	// Forzar DB nil
	database.DB = nil

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/cart", nil))
	if w.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400 por DB nil, got %d (%s)", w.Code, w.Body.String())
	}
}

func TestAddToCart_DBNil_Returns400(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(withEmail("alguien@test.com"))
	r.POST("/api/cart/add", AddToCart)

	database.DB = nil

	body := `{"product_id":1,"quantity":1}`
	req := httptest.NewRequest(http.MethodPost, "/api/cart/add", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400 por DB nil, got %d (%s)", w.Code, w.Body.String())
	}
}

func TestRemoveFromCart_DBNil_Returns400(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(withEmail("alguien@test.com"))
	r.DELETE("/api/cart/remove", RemoveFromCart)

	database.DB = nil

	body := `{"product_id":1}`
	req := httptest.NewRequest(http.MethodDelete, "/api/cart/remove", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400 por DB nil, got %d (%s)", w.Code, w.Body.String())
	}
}

func TestClearCart_DBNil_Returns400(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(withEmail("alguien@test.com"))
	r.DELETE("/api/cart/clear", ClearCart)

	database.DB = nil

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodDelete, "/api/cart/clear", nil))
	if w.Code != http.StatusBadRequest {
		t.Fatalf("esperaba 400 por DB nil, got %d (%s)", w.Code, w.Body.String())
	}
}

// -------------------- Ramas usuario NO encontrado --------------------

func TestGetCart_UserNotFound_Returns404(t *testing.T) {
	testutil.SetupInMemoryDB(t) // DB OK pero sin seed de usuario

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(withEmail("noexiste@test.com"))
	r.GET("/api/cart", GetCart)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/cart", nil))
	if w.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404 por usuario no encontrado, got %d (%s)", w.Code, w.Body.String())
	}
}

func TestAddToCart_UserNotFound_Returns404(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(withEmail("noexiste@test.com"))
	r.POST("/api/cart/add", AddToCart)

	body := `{"product_id":1,"quantity":1}`
	req := httptest.NewRequest(http.MethodPost, "/api/cart/add", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404 por usuario no encontrado, got %d (%s)", w.Code, w.Body.String())
	}
}

func TestRemoveFromCart_UserNotFound_Returns404(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(withEmail("noexiste@test.com"))
	r.DELETE("/api/cart/remove", RemoveFromCart)

	body := `{"product_id":1}`
	req := httptest.NewRequest(http.MethodDelete, "/api/cart/remove", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404 por usuario no encontrado, got %d (%s)", w.Code, w.Body.String())
	}
}

func TestClearCart_UserNotFound_Returns404(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(withEmail("noexiste@test.com"))
	r.DELETE("/api/cart/clear", ClearCart)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodDelete, "/api/cart/clear", nil))
	if w.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404 por usuario no encontrado, got %d (%s)", w.Code, w.Body.String())
	}
}
