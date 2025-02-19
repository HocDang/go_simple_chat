package usecases

import "chat-server/internal/domain/entities"

type UserUseCase interface {
	GetAllUsers() ([]entities.User, error)
}
