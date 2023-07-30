package handler

import (
	"github.com/gin-gonic/gin"
)

// @Summary ping example
// @Description hello world
// @Produce json
// @Success 200 {string} Helloworld
// @Router /api/v1/ [get]
// @Tags api/v1
func HelloWorldHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world",
	})
}

// @Summary thread post get
// @Description threads post get
// @Produce json
// @Success 200
// @Router /threads/get [get]
// @Tags threads
func GetThreadsPost(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Threads post data",
	})
}
