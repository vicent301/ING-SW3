package controllers

import (
	"backend/dao"
	"backend/domain"
	"backend/testutil"
	"encoding/json"
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

func TestGetProfile_OK_DB(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testutil.SetupInMemoryDB(t)

	// seed user
	if err := dao.CreateUser(domain.User{
		Name: "Ana", Email: "ana@test.com", Password: "x",
	}); err != nil {
		t.Fatalf("seed user: %v", err)
	}

	r := gin.New()
	// inyectamos el email en contexto simulando AuthMiddleware
	r.GET("/me", func(c *gin.Context) { c.Set("email", "ana@test.com"); GetProfile(c) })

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/me", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("esperaba 200, got %d (%s)", w.Code, w.Body.String())
	}

	var body map[string]any
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	if body["email"] != "ana@test.com" {
		t.Fatalf("email inesperado: %#v", body)
	}
}
