package repositories

import "chat-server/internal/domain/entities"

type UserRepository interface {
	Create(user *entities.User) error
	GetByEmail(email string) (*entities.User, error)
	GetAll() ([]entities.User, error)
}
