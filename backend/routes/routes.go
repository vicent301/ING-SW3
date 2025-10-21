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
		// Auth
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)
		api.GET("/ping", controllers.Ping)

		// Productos públicos
		api.GET("/products", controllers.GetProducts)
		api.GET("/products/:id", controllers.GetProductByID)

		// CRUD protegido (solo admin en futuro)
		protected := api.Group("/products")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("/", controllers.CreateProduct)
			protected.PUT("/:id", controllers.UpdateProduct)
			protected.DELETE("/:id", controllers.DeleteProduct)
		}

		// Perfil protegido
		user := api.Group("/user")
		user.Use(middleware.AuthMiddleware())
		{
			user.GET("/me", func(c *gin.Context) {
				email, _ := c.Get("email")
				c.JSON(200, gin.H{"message": "Bienvenido!", "email": email})
			})
		}
	}

	return r
}
