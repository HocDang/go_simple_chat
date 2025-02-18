package app

import (
	"chat-server/config"
	"chat-server/internal/db"
	"chat-server/internal/delivery/http"
	"log"

	"github.com/gin-gonic/gin"
)

func StartServer(env *config.Env) {
	router := gin.Default()

	// Init DB
	db.InitPostgres()
	db.InitRedis()
	db.InitElasticsearch()

	// Init Routes
	http.InitRoutes(router)

	log.Println("🚀 Server is running on :8080")
	router.Run(":8080")
}
