package routes

import (
	controllers "backend/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/ping", controllers.Ping)
	}

	return r
}
