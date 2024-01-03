package tests

import (
	"gin-boilerplate/routers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.TestMode)

	routers.RegisterRoutes(r)

	return r
}
