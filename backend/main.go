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

	// Ejecutamos las migraciones
	dao.AutoMigrateUser()
	dao.AutoMigrateProduct()
	dao.AutoMigrateCart()
	dao.AutoMigrateOrder()

	// Tomamos el puerto desde Azure (o 8080 local)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Configuramos el router principal
	r := routes.SetupRouter()

	// Endpoint de healthz para el pipeline
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Printf("Servidor escuchando en puerto %s...", port)
	err := r.Run(":" + port)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
