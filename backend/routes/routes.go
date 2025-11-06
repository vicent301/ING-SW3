package routes

import (
	"backend/controller" // <-- corregido
	"backend/middleware"
	"backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetupRouter(prodSvc services.ProductServicePort) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"https://appweb-front-prod-h7htdzbchbhsf6g2.northcentralus-01.azurewebsites.net",
			"https://appweb-front-qa-ctg3cwawggeag6g4.northcentralus-01.azurewebsites.net",
			"http://localhost:5173",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	pc := controllers.NewProductController(prodSvc)

	api := r.Group("/api")
	{
		// Health
		api.GET("/healthz", func(c *gin.Context) {
			c.String(http.StatusOK, "ok")
		})

		// Auth
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		api.GET("/products", pc.GetProducts)        // listado + ?search=
		api.GET("/products/:id", pc.GetProductByID) // detalle

		// Zona protegida
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			// Usuario
			protected.GET("/me", controllers.GetProfile)

			// Carrito
			protected.GET("/cart", controllers.GetCart)
			protected.POST("/cart/add", controllers.AddToCart)
			protected.DELETE("/cart/remove", controllers.RemoveFromCart)
			protected.DELETE("/cart/clear", controllers.ClearCart)

			// Productos (solo creación protegida)
			protected.POST("/products", pc.CreateProduct)

			// Órdenes
			protected.POST("/orders/create", controllers.CreateOrder)
			protected.GET("/orders", controllers.GetOrders)
		}
	}

	return r
}
