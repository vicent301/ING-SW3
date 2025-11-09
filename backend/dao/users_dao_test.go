package dao_test

import (
	"backend/dao"
	"backend/database"
	"backend/domain"
	"backend/testutil"
	"os"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// testPassword obtiene la contraseña desde el entorno para cumplir con Sonar.
// Si no está definida, usa un valor de prueba no sensible.
func testPassword() string {
	if v := os.Getenv("TEST_PASSWORD"); v != "" {
		return v
	}
	return "test-password" // fixture de test, no es un secreto real
}

func TestUsersDAO_CreateUser_HashesPassword(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	pwd := testPassword()

	err := dao.CreateUser(domain.User{
		Name: "Ana", Email: "ana@test.com", Password: pwd,
	})
	if err != nil {
		t.Fatalf("CreateUser: %v", err)
	}

	var row dao.User
	if err := database.DB.Where("email = ?", "ana@test.com").First(&row).Error; err != nil {
		t.Fatalf("query user: %v", err)
	}

	// No debe guardarse en claro
	if row.Password == pwd {
		t.Fatalf("la password no debería guardarse en claro")
	}

	// El hash debe corresponder a la contraseña usada en el seed
	if bcrypt.CompareHashAndPassword([]byte(row.Password), []byte(pwd)) != nil {
		t.Fatalf("hash no corresponde con la password original")
	}
}

func TestUsersDAO_GetUserByEmail_NotFound(t *testing.T) {
	testutil.SetupInMemoryDB(t)
	if _, err := dao.GetUserByEmail("no@existe.com"); err == nil {
		t.Fatalf("esperaba error por usuario inexistente")
	}
}

func TestUsersDAO_GetUserByID_NotFound(t *testing.T) {
	testutil.SetupInMemoryDB(t)
	if _, err := dao.GetUserByID(9999); err == nil {
		t.Fatalf("esperaba error por ID inexistente")
	}
}
