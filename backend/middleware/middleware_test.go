package middleware_test

import (
	"backend/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthMiddleware_NoToken_Returns401(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(middleware.AuthMiddleware())
	r.GET("/private", func(c *gin.Context) {
		c.String(200, "ok")
	})

	req, _ := http.NewRequest(http.MethodGet, "/private", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
