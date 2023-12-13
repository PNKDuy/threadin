package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the API key from the request header
		apiKey := c.GetHeader("X-API-Key")

		// Validate the API key
		if apiKey != "secret-key" {
			// Return an error response if the API key is invalid
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			c.Abort()
			return
		}

		// Call the next handler
		c.Next()
	}
}
