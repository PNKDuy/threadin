package router

import (
	"sandbox-api/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		// v1.Use(middleware.AuthMiddleware())

		v1.GET("/", handler.HelloWorldHandler)
		v1.GET("/auth/{provider}/callback", handler.AuthCallbackHandler)
		v1.GET("/logout/{provider}", handler.LogoutHandler)

		v1.GET("/auth/{provider}", handler.AuthHandler)
	}
	return router
}
