package controllers

import (
	"backend/domain"
	"backend/tests/mocks"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetProductByID_OK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := new(mocks.ProductServiceMock)
	svc.On("GetProduct", uint(1)).Return(&domain.Product{ID: 1, Name: "Zapa"}, nil)

	pc := NewProductController(svc)
	r := gin.New()
	r.GET("/products/:id", pc.GetProductByID)

	req, _ := http.NewRequest(http.MethodGet, "/products/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"Zapa"`)
	svc.AssertExpectations(t)
}

func TestCreateProduct_OK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := new(mocks.ProductServiceMock)
	svc.On("CreateProduct", domain.Product{Name: "Zapa", Price: 100}).Return(nil)

	pc := NewProductController(svc)
	r := gin.New()
	r.POST("/products", pc.CreateProduct)

	body := `{"name":"Zapa","price":100}`
	req, _ := http.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Producto creado correctamente")
	svc.AssertExpectations(t)
}

func TestGetProductByID_BadID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := new(mocks.ProductServiceMock)
	pc := NewProductController(svc)
	r := gin.New()
	r.GET("/products/:id", pc.GetProductByID)

	req, _ := http.NewRequest(http.MethodGet, "/products/abc", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetProductByID_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := new(mocks.ProductServiceMock)
	svc.On("GetProduct", uint(2)).Return(nil, errors.New("not found"))

	pc := NewProductController(svc)
	r := gin.New()
	r.GET("/products/:id", pc.GetProductByID)

	req, _ := http.NewRequest(http.MethodGet, "/products/2", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Producto no encontrado")
	svc.AssertExpectations(t)
}

func TestGetProducts_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := new(mocks.ProductServiceMock)
	svc.On("SearchProducts", "zapas").Return(nil, errors.New("db down"))

	pc := NewProductController(svc)
	r := gin.New()
	r.GET("/products", pc.GetProducts)

	req, _ := http.NewRequest(http.MethodGet, "/products?search=zapas", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}

func TestCreateProduct_BadJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := new(mocks.ProductServiceMock)
	pc := NewProductController(svc)
	r := gin.New()
	r.POST("/products", pc.CreateProduct)

	req, _ := http.NewRequest(http.MethodPost, "/products", strings.NewReader(`{mal json}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateProduct_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := new(mocks.ProductServiceMock)
	svc.On("CreateProduct", domain.Product{Name: "Zapa", Price: 100}).Return(errors.New("db"))

	pc := NewProductController(svc)
	r := gin.New()
	r.POST("/products", pc.CreateProduct)

	body := `{"name":"Zapa","price":100}`
	req, _ := http.NewRequest(http.MethodPost, "/products", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	svc.AssertExpectations(t)
}

// ==== NUEVOS TESTS PARA SUBIR COVERAGE DEL CONTROLLER ====
func TestGetProducts_OK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := new(mocks.ProductServiceMock)
	svc.On("SearchProducts", "zapa").
		Return([]domain.Product{{ID: 1, Name: "Zapa 1"}, {ID: 2, Name: "Zapa 2"}}, nil)

	pc := NewProductController(svc)
	r := gin.New()
	r.GET("/products", pc.GetProducts)

	req, _ := http.NewRequest(http.MethodGet, "/products?search=zapa", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"Zapa 1"`)
	assert.Contains(t, w.Body.String(), `"Zapa 2"`)
	svc.AssertExpectations(t)
}

func TestGetProducts_EmptyQuery_OK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := new(mocks.ProductServiceMock)
	// cuando no viene ?search, el controller manda "" al servicio
	svc.On("SearchProducts", "").
		Return([]domain.Product{}, nil)

	pc := NewProductController(svc)
	r := gin.New()
	r.GET("/products", pc.GetProducts)

	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "[]") // lista vacía OK
	svc.AssertExpectations(t)
}
