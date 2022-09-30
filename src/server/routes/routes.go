package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucca-rodrigues/myStock-API-GO/src/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.GET("/", controllers.Find)
			users.POST("/", controllers.Create)
		}
	}
	return router
}