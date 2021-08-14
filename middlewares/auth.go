package middlewares

import (
	"errors"
	"net/http"

	"github.com/coala/corobo-ng/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ErrCookieTokenEmpty = errors.New("cookie token is empty")
)

// Get token value from cookies
func GetCookieToken(c *gin.Context) (string, error) {
	cookie, _ := c.Cookie("token")

	if cookie == "" {
		return "", ErrCookieTokenEmpty
	}

	return cookie, nil
}

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := GetCookieToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Token not found",
				"success": false,
			})
			return
		}

		_, err = services.GetUserByToken(db, token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "User not found",
				"success": false,
			})
			return
		}
		c.Next()
	}
}
