package dao_test

import (
	"backend/dao"
	"backend/database"
	"backend/domain"
	"backend/testutil"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestUsersDAO_CreateUser_HashesPassword(t *testing.T) {
	testutil.SetupInMemoryDB(t)

	err := dao.CreateUser(domain.User{
		Name: "Ana", Email: "ana@test.com", Password: "secreto",
	})
	if err != nil {
		t.Fatalf("CreateUser: %v", err)
	}

	var row dao.User
	if err := database.DB.Where("email = ?", "ana@test.com").First(&row).Error; err != nil {
		t.Fatalf("query user: %v", err)
	}
	if row.Password == "secreto" {
		t.Fatalf("la password no debería guardarse en claro")
	}
	if bcrypt.CompareHashAndPassword([]byte(row.Password), []byte("secreto")) != nil {
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
