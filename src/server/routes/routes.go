package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lucca-rodrigues/myStock-API-GO/src/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		// route := main.Group("users")
		// {
		// 	route.GET("/users", controllers.Find)
		// 	route.POST("/users", controllers.Create)
		// }
		main.GET("/", func(c *gin.Context) {
			c.JSON(200, "Ok api started")
		})

		main.GET("/users", controllers.Find)
		main.POST("/users", controllers.Create)
		main.GET("/products", controllers.FindAllProducts)
		main.POST("/products", controllers.CreateProduct)
	}
	return router
}