package testutil

import (
	"backend/dao"
	"backend/database"
	"testing"

	"github.com/glebarez/sqlite" // pure Go, sin CGO
	"gorm.io/gorm"
)

func SetupInMemoryDB(t *testing.T) {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("no se pudo abrir sqlite en memoria: %v", err)
	}
	database.DB = db

	// 👉 IMPORTANTE: migrá TODO lo que tus DAOs usan
	if err := database.DB.AutoMigrate(
		&dao.User{},
		&dao.Product{},
		&dao.CartEntity{},
		&dao.CartItemEntity{},
		&dao.OrderEntity{},
		&dao.OrderItemEntity{},
	); err != nil {
		t.Fatalf("automigrate falló: %v", err)
	}
}
