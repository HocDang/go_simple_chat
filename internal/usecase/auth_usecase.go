package usecase

import (
	"chat-server/internal/domain/entities"
	"chat-server/internal/domain/repositories"
	"chat-server/pkg/utils"
	"errors"
)

type AuthUseCase struct {
	userRepo repositories.UserRepository
}

func NewAuthUseCase(userRepo repositories.UserRepository) *AuthUseCase {
	return &AuthUseCase{userRepo: userRepo}
}

func (uc *AuthUseCase) Register(email, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := entities.User{
		Email:    email,
		Password: hashedPassword,
	}

	return uc.userRepo.Create(&user)
}

func (uc *AuthUseCase) Login(email, password string) (*entities.User, error) {
	user, err := uc.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
