package usecases

import "chat-server/internal/domain/entities"

type AuthUseCase interface {
	GetAllUsers() ([]entities.User, error)
}
