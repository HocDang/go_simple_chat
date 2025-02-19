package usecases

import "chat-server/internal/domain/entities"

type UserUseCase interface {
	GetUsers() ([]entities.User, error)
	GetUserByID(id int) (*entities.User, error)
}
