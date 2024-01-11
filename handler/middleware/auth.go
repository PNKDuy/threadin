package middleware

import (
	"sandbox-api/configs"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Get the API key from the request header
// 		apiKey := c.GetHeader("X-API-Key")

// 		// Validate the API key
// 		if apiKey != "secret-key" {
// 			// Return an error response if the API key is invalid
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
// 			c.Abort()
// 			return
// 		}

// 		// Call the next handler
// 		c.Next()
// 	}
// }

// func AuthorizationMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Get the Authorization header from the request
// 		authorization := c.GetHeader("Authorization")

// 		// Validate the Authorization header
// 		if !strings.HasPrefix(authorization, "Bearer ") {
// 			// Return an error response if the Authorization header is invalid
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
// 			c.Abort()
// 			return
// 		}

// 		// Split the Authorization header to get the token
// 		_ = strings.TrimPrefix(authorization, "Bearer ")

// 		// Validate the token

// 		// Call the next handler
// 		c.Next()
// 	}
// }

func NewAuth(cfg *configs.Config) {
	store := sessions.NewCookieStore([]byte("secret"))
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.MaxAge = 60 * 60 * 24 * 7 // 1 week
	store.Options.Secure = false

	gothic.Store = store
	goth.UseProviders(
		google.New(cfg.GoogleClientID, cfg.GoogleClientSecret, "http://localhost:3000/auth/google/callback"),
	)
}
