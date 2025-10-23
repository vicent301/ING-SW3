package main

import (
	"backend/dao"
	"backend/database"
	"backend/routes"
	"log"
	"os"
)

func main() {
	// Conectamos la BD
	database.Connect()

	// Migraciones
	dao.AutoMigrateUser()
	dao.AutoMigrateProduct()
	dao.AutoMigrateCart()
	dao.AutoMigrateOrder()

	// Router principal
	r := routes.SetupRouter()

	// Obtener el puerto asignado por Azure
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback local
	}

	log.Printf("Servidor escuchando en puerto %s...", port)
	r.Run(":" + port)
}
