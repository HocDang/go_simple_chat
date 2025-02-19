package usecases

import "chat-server/internal/domain/entities"

type AuthUseCase interface {
	Register(email string, password string) error
	Login(email string, password string) (*entities.User, error)
}
