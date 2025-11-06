package routes

import (
	"backend/domain"
	"backend/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoutes_GetProductByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := new(mocks.ProductServiceMock)
	svc.On("GetProduct", uint(1)).Return(&domain.Product{ID: 1, Name: "Zapa"}, nil)

	r := SetupRouter(svc)

	req, _ := http.NewRequest(http.MethodGet, "/api/products/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"Zapa"`)
	svc.AssertExpectations(t)
}
func TestRoutes_Healthz_OK(t *testing.T) {
	gin.SetMode(gin.TestMode)
	svc := new(mocks.ProductServiceMock)
	r := SetupRouter(svc)

	req, _ := http.NewRequest(http.MethodGet, "/api/healthz", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "ok", w.Body.String())
}
