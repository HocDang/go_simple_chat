package app

import (
	"chat-server/config"
	"chat-server/internal/cache"
	"chat-server/internal/db"
	"chat-server/internal/delivery/http"
	"log"

	"github.com/gin-gonic/gin"
)

func StartServer(env *config.Env) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies([]string{})

	// Init DB
	database := db.InitPostgres(env)
	defer db.Close(database)

	// Init Cache redis with Interface
	var redis cache.CacheInterface
	redis = cache.InitRedis(env)
	defer redis.Close()

	// Init Elasticsearch
	db.InitElasticsearch()

	// Init Routes
	http.InitRoutes(router)

	log.Println("🚀 Server started at", env.ServerAddress)
	router.Run(env.ServerAddress)
}
