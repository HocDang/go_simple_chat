package app

import (
	"chat-server/config"
	"chat-server/internal/bootstrap"
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
	db.Migrate(database)
	defer db.Close(database)

	// Init connect to Redis
	redis := cache.InitRedis(env)
	defer redis.Close()

	// Init Elasticsearch
	es := bootstrap.InitElasticsearch(env.EsHost, env.EsPort)

	// Init Container
	container := container.NewContainer(database, es)

	// Register Routes
	http.RegisterRoutes(router, container)

	log.Println("🚀 Server started at", env.ServerAddress)
	router.Run(env.ServerAddress)
}
