package routes

import (
	"backend/controller"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		api.GET("/products", controllers.GetProducts)
		api.GET("/products/:id", controllers.GetProductByID)

		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			//usuario
			protected.GET("/me", controllers.GetProfile)
			// Carrito
			protected.GET("/cart", controllers.GetCart)
			protected.POST("/cart/add", controllers.AddToCart)
			protected.DELETE("/cart/remove", controllers.RemoveFromCart)
			protected.DELETE("/cart/clear", controllers.ClearCart)
			//productos
			protected.POST("/products", controllers.CreateProduct)

		}
	}

	return r
}
