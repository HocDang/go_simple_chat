package app

import (
	"chat-server/config"
	"chat-server/internal/cache"
	"chat-server/internal/container"
	"chat-server/internal/db"
	"chat-server/internal/delivery/http"
	"log"

	"github.com/gin-gonic/gin"
)

func StartServer(env *config.Env) {
	// Init Router
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies([]string{})

	// Init connect to database
	database := db.InitPostgres(env)
	defer db.Close(database)

	// Init connect to Redis
	redis := cache.InitRedis(env)
	defer redis.Close()

	// Init Elasticsearch
	db.InitElasticsearch()

	// Init Container
	container := container.NewContainer(database)

	// Register Routes
	http.RegisterRoutes(router, container)

	log.Println("🚀 Server started at", env.ServerAddress)
	router.Run(env.ServerAddress)
}
