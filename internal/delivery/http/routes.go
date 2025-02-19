package http

import (
	"chat-server/internal/container"

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

		users := api.Group("/users")
		{
			users.GET("/", userHandler.GetUsers)
		}

		chat := api.Group("/chat")
		{
			chat.POST("/send", chatHandler.SendMessage)
			chat.GET("/messages", chatHandler.GetMessages)
		}
	}
}
