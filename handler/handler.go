package handler

import (
	"net/http"
	"sandbox-api/configs"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
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

func AuthCallbackHandler(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	t, _ := template.New("foo").Parse(configs.UserTemplate)
	t.Execute(c.Writer, user)
}

func LogoutHandler(c *gin.Context) {
	gothic.Logout(c.Writer, c.Request)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func AuthHandler(c *gin.Context) {
	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		t, _ := template.New("foo").Parse(configs.UserTemplate)
		t.Execute(c.Writer, gothUser)
	} else {
		gothic.BeginAuthHandler(c.Writer, c.Request)
	}
}
