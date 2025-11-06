package main

import (
	"log"
	"os"

	"backend/dao"
	"backend/database"
	"backend/routes"
	"backend/services"
)

func main() {
	// 1) Conectar BD
	db := database.Connect()
	//    o, si deja en un global: database.Connect(); db := database.DB

	// 2) Migraciones
	dao.AutoMigrateUser()
	dao.AutoMigrateProduct()
	dao.AutoMigrateCart()
	dao.AutoMigrateOrder()

	// 3) Repositorio concreto (GORM) e inyección en el servicio
	repo := dao.NewProductGormRepository(db) // <-- tu implementación concreta
	svc := services.NewProductService(repo)

	// 4) Router con servicio inyectado
	r := routes.SetupRouter(svc)

	// 5) Puerto (Azure u 8080 local)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor escuchando en puerto %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
