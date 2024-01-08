package main

import (
	"log"
	"sandbox-api/router"

	_ "sandbox-api/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Sandbox API
// @version         1.0
// @description     My first APIhub.
// @contact.name   KayDee
// @contact.email  phngkhuongduy@gmail.com
// @host      localhost:8080
// @securityDefinitions.apikey ApiKeyHeader
// @in header
// @name X-API-KEY
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}
	r := router.SetupRouter()

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger-ui/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8080")
}
