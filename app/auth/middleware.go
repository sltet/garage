package auth

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
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
