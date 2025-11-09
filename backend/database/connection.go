package database

import (
	"fmt"
	"log"

	"backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Conexión normal (producción)
func Connect() *gorm.DB {
	dsn := config.GetDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Error al conectar a la base de datos:", err)
	}
	fmt.Println("✅ Conexión exitosa a MySQL")

	DB = db
	return DB
}

// ✅ Conexión inyectable para tests (no toca una DB real)
func ConnectWithDialector(dial gorm.Dialector) *gorm.DB {
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Error al abrir dialector (test):", err)
	}
	DB = db
	return DB
}
