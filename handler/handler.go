package handler

import (
	"log"

	"github.com/gin-gonic/gin"
)

// @Summary ping example
// @Description hello world
// @Security ApiKeyHeader
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/v1/ [get]
// @Tags api/v1
func HelloWorldHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}

func PaymentWebhookHandler(c *gin.Context) {
	var payment PaymentNotification
	err := c.BindJSON(&payment)
	if err != nil {
		log.Printf("Failed to bind JSON: %v", err)
		c.JSON(400, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	// Announce the payment
	go announcePayment(payment)

	c.JSON(200, gin.H{
		"message": "Payment received",
	})
}
