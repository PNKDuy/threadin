package handler

import "github.com/gin-gonic/gin"

const ThreadsAPIURL = "https://www.threads.net/api/graphql"

type PublicAPI struct {
	APPToken string
}

func prepareHeaders(c *gin.Context, token string) {
	for k, v := range defaultHeaders(token) {
		c.Request.Header.Set(k, v)
	}
}

func defaultHeaders(token string) map[string]string {
	return map[string]string{
		"Accept":          "*/*",
		"Accept-Language": "en-US,en;q=0.9",
		"Cache-Control":   "no-cache",
		"Content-Type":    "application/x-www-form-urlencoded",
	}
}
