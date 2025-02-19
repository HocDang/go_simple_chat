package middleware

import (
	"net/http"
	"strings"

	"chat-server/pkg/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, "Bearer ")

		if len(t) == 2 {
			authToken := t[1]
			userID, err := utils.ValidateToken(authToken)

			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}

			c.Set("x-user-id", userID)
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
}
