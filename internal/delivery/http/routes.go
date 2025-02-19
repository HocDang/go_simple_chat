package http

import (
	"chat-server/internal/container"
	"chat-server/internal/delivery/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, container *container.Container) {
	authHandler := NewAuthHandler(container.AuthUseCase)
	userHandler := NewUserHandler(container.UserUseCase)
	chatHandler := NewMessageHandler(container.MessageUseCase)

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
		}

		user := api.Group("/user")
		{
			user.Use(middleware.JwtAuthMiddleware())
			user.GET("/", userHandler.GetUsers)
		}

		messages := api.Group("/messages")
		{
			messages.Use(middleware.JwtAuthMiddleware())
			messages.POST("/", chatHandler.SendMessage)
			messages.GET("/:id", chatHandler.GetMessages)
		}
	}
}
