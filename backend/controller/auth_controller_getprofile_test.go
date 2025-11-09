package controllers

import (
	"backend/dao"
	"backend/database"
	"backend/domain"
	"backend/testutil"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

// ---- helpers para armar requests/response ----

func newJSONReq(t *testing.T, method, path string, body any) *http.Request {
	t.Helper()
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			t.Fatalf("encode json: %v", err)
		}
	}
	req := httptest.NewRequest(method, path, &buf)
	req.Header.Set(HeaderContentType, MIMEApplicationJSON)
	return req
}

func parseJSON(t *testing.T, w *httptest.ResponseRecorder, out any) {
	t.Helper()
	if err := json.NewDecoder(w.Body).Decode(out); err != nil {
		t.Fatalf("decode json: %v. body=%s", err, w.Body.String())
	}
}

// -----------------------------------------------------------
// 1) GetProfile cuando DB == nil  (ruta feliz mínima)
// -----------------------------------------------------------
func TestGetProfile_DBIsNil_Returns200MinimalPayload(t *testing.T) {
	// forzamos DB nil
	database.DB = nil

	gin.SetMode(gin.TestMode)
	r := gin.New()

	// ruta protegida
	r.GET("/api/me", func(c *gin.Context) {
		// simulamos que el middleware setea "email"
		c.Set("email", "me@test.com")
		GetProfile(c)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/me", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status=%d; body=%s", w.Code, w.Body.String())
	}
	var got map[string]any
	parseJSON(t, w, &got)
	if got["email"] != "me@test.com" {
		t.Fatalf("email esperado %q, got %v", "me@test.com", got["email"])
	}
}

// -----------------------------------------------------------
// 2) GetProfile con DB real (en memoria) -> NOT FOUND
// -----------------------------------------------------------
func TestGetProfile_DBUserNotFound_Returns404(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/api/me", func(c *gin.Context) {
		c.Set("userEmail", "noexiste@site.com") // otra vía: "userEmail"
		GetProfile(c)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/me", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("esperaba 404, got %d; body=%s", w.Code, w.Body.String())
	}
}

// -----------------------------------------------------------
// 3) GetProfile con DB real -> FOUND (User en contexto)
//    Cubrimos también getEmailAny con User, *User y map[string]any
// -----------------------------------------------------------

func TestGetProfile_DBUserFound_FromEmailKey_Returns200(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	// seed usuario
	if err := dao.CreateUser(domain.User{
		Name: "Vic", Email: "ok@test.com", Password: "x",
	}); err != nil {
		t.Fatalf("seed: %v", err)
	}

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/api/me", func(c *gin.Context) {
		c.Set("email", "ok@test.com") // vía "email"
		GetProfile(c)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/me", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status=%d; body=%s", w.Code, w.Body.String())
	}
	var got map[string]any
	parseJSON(t, w, &got)
	if got["email"] != "ok@test.com" || strings.TrimSpace(got["name"].(string)) == "" {
		t.Fatalf("payload inesperado: %#v", got)
	}
}

func TestGetProfile_DBUserFound_FromUserStruct_Returns200(t *testing.T) {
	testutil.SetupInMemoryDB(t)
	_ = dao.CreateUser(domain.User{Name: "A", Email: "struct@test.com", Password: "x"})

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/api/me", func(c *gin.Context) {
		c.Set("user", domain.User{Email: "struct@test.com"}) // domain.User
		GetProfile(c)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/me", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status=%d; body=%s", w.Code, w.Body.String())
	}
}

func TestGetProfile_DBUserFound_FromUserPtr_Returns200(t *testing.T) {
	testutil.SetupInMemoryDB(t)
	_ = dao.CreateUser(domain.User{Name: "B", Email: "ptr@test.com", Password: "x"})

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/api/me", func(c *gin.Context) {
		u := &domain.User{Email: "ptr@test.com"} // *domain.User
		c.Set("user", u)
		GetProfile(c)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/me", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status=%d; body=%s", w.Code, w.Body.String())
	}
}

func TestGetProfile_DBUserFound_FromUserMap_Returns200(t *testing.T) {
	testutil.SetupInMemoryDB(t)
	_ = dao.CreateUser(domain.User{Name: "C", Email: "map@test.com", Password: "x"})

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/api/me", func(c *gin.Context) {
		c.Set("user", map[string]any{"email": "map@test.com"}) // map[string]any
		GetProfile(c)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/me", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status=%d; body=%s", w.Code, w.Body.String())
	}
}
