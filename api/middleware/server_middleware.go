package middleware

import (
	"github.com/gin-gonic/gin"
)

func ServerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		serverName := "IO - Go/1.22"
		c.Writer.Header().Set("Server", serverName)

	}
}
