package container

import (
	"chat-server/internal/repository/postgres"
	"chat-server/internal/usecase"

	"github.com/go-pg/pg/v10"
)

type Container struct {
	AuthUseCase    *usecase.AuthUseCase
	UserUseCase    *usecase.UserUseCase
	MessageUseCase *usecase.MessageUseCase
}

func NewContainer(db *pg.DB) *Container {
	userRepo := postgres.NewUserPgRepository(db)
	messageRepo := postgres.NewMessagePgRepository(db)

	authUseCase := usecase.NewAuthUseCase(userRepo)
	userUseCase := usecase.NewUserUseCase(userRepo)
	messageUseCase := usecase.NewMessageUseCase(messageRepo)

	return &Container{
		AuthUseCase:    authUseCase,
		UserUseCase:    userUseCase,
		MessageUseCase: messageUseCase,
	}
}
