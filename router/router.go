package router

import (
	"sandbox-api/handler"
	"sandbox-api/handler/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.Use(middleware.AuthMiddleware())
		v1.GET("/", handler.HelloWorldHandler)
	}
	return router
}
