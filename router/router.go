package router

import (
	"sandbox-api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/", handler.HelloWorldHandler)
	}
	return router
}
