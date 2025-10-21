package main

import (
	"backend/dao"
	"backend/database"
	"backend/routes"
)

func main() {
	// Conectamos la BD
	database.Connect()

	// Ejecutamos las migraciones (ahora sí podemos importar dao)
	dao.AutoMigrateUser()
	dao.AutoMigrateProduct()
	dao.AutoMigrateCart()
	dao.AutoMigrateOrder()

	// Levantamos el servidor
	r := routes.SetupRouter()
	r.Run(":8080")
}
