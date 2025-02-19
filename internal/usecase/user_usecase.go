package usecase

import (
	"chat-server/internal/domain/entities"
	"chat-server/internal/domain/repositories"
)

type UserUseCase struct {
	userRepo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (u *UserUseCase) GetAllUsers() ([]entities.User, error) {
	users, err := u.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var domainUsers []entities.User
	for _, user := range users {
		domainUsers = append(domainUsers, entities.User{
			ID: user.ID,
		})
	}
	return domainUsers, nil
}
