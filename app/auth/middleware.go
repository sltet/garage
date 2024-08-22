package auth

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sltet/garage/app/core"
	"go.uber.org/dig"
	"net/http"
	"strings"
)

func AuthMiddleware(ctn *dig.Container) gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := getToken(c); err == nil {
			ctn.Invoke(func(oauthServer OauthServerInterface) {
				data, err := oauthServer.LoadToken(token)
				if err != nil {
					c.JSON(http.StatusForbidden, gin.H{
						"message": "not authenticated",
						"error":   err.Error(),
					})
					c.Abort()
				}
				if data.IsExpired() {
					c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("/auth/%s", data.Client.GetId()))
					return
				}
				c.Next()
			})
		} else {
			session := sessions.Default(c)
			if user := session.Get("user"); user != nil {
				c.Next()

			} else {
				c.JSON(http.StatusForbidden, gin.H{
					"message": "not authenticated",
				})
				c.Abort()
			}
		}

	}
}

func getToken(c *gin.Context) (string, core.DetailedError) {
	bearerToken := c.Request.Header.Get("Authorization")
	if bearerToken == "" {
		return bearerToken, core.NewUnauthorizedError("Missing authorization header")
	}
	reqToken := strings.Split(bearerToken, " ")[1]
	if reqToken == "" {
		return reqToken, core.NewUnauthorizedError("Missing authorization header")
	}
	return reqToken, nil
}
