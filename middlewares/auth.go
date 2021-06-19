package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqKey := c.Request.Header.Get("X-Auth-Key")

		if reqKey != "pass_auth" {
			c.AbortWithStatus(http.StatusForbidden)
		}

		c.Next()
	}
}
