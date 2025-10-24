package routes

import (
	"backend/controller"
	"backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://appweb-front-prod-h7htdzbchbhsf6g2.northcentralus-01.azurewebsites.net",
			"https://appweb-front-qa-ctg3cwawggeag6g4.northcentralus-01.azurewebsites.net",
			"http://localhost:5173", // para pruebas locales
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
	{
		api.GET("/healthz", func(c *gin.Context) {
			c.String(http.StatusOK, "ok")
		})

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

			//ordenes
			protected.POST("/orders/create", controllers.CreateOrder)
			protected.GET("/orders", controllers.GetOrders)
		}
	}

	return r
}
