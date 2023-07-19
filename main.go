package main

import (
	"sandbox-api/router"

	_ "sandbox-api/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Sandbox API
// @version         1.0
// @description     My first APIhub.

// @contact.name   KayDee
// @contact.email  phngkhuongduy@gmail.com

// @host      localhost:8080
func main() {
	r := router.SetupRouter()
	r.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
