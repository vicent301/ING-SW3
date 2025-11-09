package controllers

import (
	"backend/dao"
	"backend/domain"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"backend/testutil"
	"encoding/json"
)

func TestRegister_BadJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST(RouteRegister, Register)

	req, _ := http.NewRequest(http.MethodPost, RouteRegister, strings.NewReader(`{mal json}`))
	req.Header.Set(HeaderContentType, MIMEApplicationJSON)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestRegister_MissingFields(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST(RouteRegister, Register)

	body := `{"email":"","password":""}`
	req, _ := http.NewRequest(http.MethodPost, RouteRegister, strings.NewReader(body))
	req.Header.Set(HeaderContentType, MIMEApplicationJSON)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// la mayoría de implementaciones valida y devuelve 400
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLogin_BadJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST(RouteLogin, Login)

	req, _ := http.NewRequest(http.MethodPost, RouteLogin, strings.NewReader(`{`))
	req.Header.Set(HeaderContentType, MIMEApplicationJSON)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLogin_InvalidCredentials(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST(RouteLogin, Login)

	body := `{"email":"nada@x.com","password":"mala"}`
	req, _ := http.NewRequest(http.MethodPost, RouteLogin, strings.NewReader(body))
	req.Header.Set(HeaderContentType, MIMEApplicationJSON)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// muchas implementaciones devuelven 401 o 400; ajustá si tu handler usa 400
	assert.Contains(t, []int{http.StatusUnauthorized, http.StatusBadRequest}, w.Code)
}
func TestRegister_OK(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST(RouteRegister, Register)

	body := map[string]string{
		"name":     "Vicente",
		"email":    "test@example.com",
		"password": "123456",
	}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, RouteRegister, bytes.NewReader(b))
	req.Header.Set(HeaderContentType, MIMEApplicationJSON)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d. body=%s", w.Code, w.Body.String())
	}
	// chequeo simple: usuario creado
	u, err := dao.GetUserByEmail("test@example.com")
	if err != nil || u == nil {
		t.Fatalf("no se creó el usuario: %v", err)
	}
}

func TestLogin_OK(t *testing.T) {
	testutil.SetupInMemoryDB(t)
	// seed usuario
	if err := dao.CreateUser(domain.User{
		Name:     "Vic",
		Email:    "login@example.com",
		Password: "clave123",
	}); err != nil {
		t.Fatalf("seed user error: %v", err)
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST(RouteLogin, Login)

	body := map[string]string{
		"email":    "login@example.com",
		"password": "clave123",
	}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, RouteLogin, bytes.NewReader(b))
	req.Header.Set(HeaderContentType, MIMEApplicationJSON)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("esperaba 200, obtuve %d. body=%s", w.Code, w.Body.String())
	}
	if len(w.Body.String()) == 0 {
		t.Fatalf("esperaba token en respuesta")
	}
}

func TestRegister_DuplicateEmail_Returns500(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	// seed: usuario existente
	if err := dao.CreateUser(domain.User{
		Name: "Vic", Email: "dup@example.com", Password: "x",
	}); err != nil {
		t.Fatalf("seed user: %v", err)
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST(RouteRegister, Register)

	body := map[string]string{"name": "Otro", "email": "dup@example.com", "password": "y"}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, RouteRegister, bytes.NewReader(b))
	req.Header.Set(HeaderContentType, MIMEApplicationJSON)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("esperaba 500 por email duplicado, got %d (%s)", w.Code, w.Body.String())
	}
}
func TestLogin_WrongPassword_Returns401(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	if err := dao.CreateUser(domain.User{
		Name: "Vic", Email: "login@ex.com", Password: "correcta",
	}); err != nil {
		t.Fatalf("seed user: %v", err)
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST(RouteLogin, Login)

	body := map[string]string{"email": "login@ex.com", "password": "mala"}
	b, _ := json.Marshal(body)
	req := httptest.NewRequest(http.MethodPost, RouteLogin, bytes.NewReader(b))
	req.Header.Set(HeaderContentType, MIMEApplicationJSON)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("esperaba 401, got %d (%s)", w.Code, w.Body.String())
	}
}
