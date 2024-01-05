package routers

import (
	"gin-boilerplate/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(route *gin.Engine) {
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Route Not Found"})
	})

	route.GET("/health", controllers.Health)

	// example
	example := route.Group("/example")
	example.GET("", controllers.GetAllExample)
	example.GET("/:id", controllers.GetOneExample)
	example.POST("", controllers.CreateExample)
	example.PUT("", controllers.UpdateExample)
	example.DELETE("/:id", controllers.DeleteExample)
}
