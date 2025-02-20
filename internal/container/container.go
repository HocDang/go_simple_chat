package container

import (
	"chat-server/internal/repository/postgres"
	"chat-server/internal/search"
	"chat-server/internal/usecase"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/go-pg/pg/v10"
)

type Container struct {
	AuthUseCase    *usecase.AuthUseCase
	UserUseCase    *usecase.UserUseCase
	MessageUseCase *usecase.MessageUseCase
}

func NewContainer(db *pg.DB, es *elasticsearch.Client) *Container {
	userRepo := postgres.NewUserPgRepository(db)
	messageRepo := postgres.NewMessagePgRepository(db)
	seachMessage := search.NewEsMessage(es, "messages")

	authUseCase := usecase.NewAuthUseCase(userRepo)
	userUseCase := usecase.NewUserUseCase(userRepo)

	messageUseCase := usecase.NewMessageUseCase(messageRepo, seachMessage)

	return &Container{
		AuthUseCase:    authUseCase,
		UserUseCase:    userUseCase,
		MessageUseCase: messageUseCase,
	}
}
