package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// helper: inyecta userID en el contexto (simula auth middleware)
func withUser(id uint) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("userID", id)
		c.Next()
	}
}

func TestGetProfile_Unauthorized_NoUserInContext(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/me", GetProfile) // sin userID

	req, _ := http.NewRequest(http.MethodGet, "/me", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
