package database

import (
	"fmt"
	"log"

	"backend/config"
	"backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := config.GetDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Error al conectar a la base de datos:", err)
	}

	fmt.Println("✅ Conexión exitosa a MySQL")

	// Auto-migrar modelos
	db.AutoMigrate(&models.User{})

	DB = db
}
