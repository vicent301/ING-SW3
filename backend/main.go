package main

import (
	"backend/dao"
	"backend/database"
	"backend/routes"
	"github.com/gin-gonic/gin"
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

	// Endpoint de health check
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Obtener el puerto asignado por Azure
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback local
	}

	log.Printf("Servidor escuchando en puerto %s...", port)
	r.Run(":" + port)
}
