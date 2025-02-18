package route

import (
	"time"

	"w3s/go-backend/bootstrap"
	"w3s/go-backend/mongo"

	"github.com/gin-gonic/gin"
)

func NewHomeRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {

	group.GET("/", welcome)
}

func welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to API",
	})
}
