package main

import (
	"backend/dao"
	"backend/database"
	"backend/routes"
	"net/http"
)

func main() {
	// Conectamos la BD
	database.Connect()

	// Ejecutamos las migraciones (ahora sí podemos importar dao)
	dao.AutoMigrateUser()
	dao.AutoMigrateProduct()
	dao.AutoMigrateCart()
	dao.AutoMigrateOrder()

		http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	// Levantamos el servidor
	r := routes.SetupRouter()
	r.Run(":8080")
}
