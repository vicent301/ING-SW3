package main

import (
	"backend/database"
	"backend/routes"
)

func main() {
	// Conexión a MySQL
	database.Connect()

	// Rutas
	r := routes.SetupRouter()

	// Puerto 8080
	r.Run(":8080")
}
