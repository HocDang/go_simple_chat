package usecase

import (
	"context"
	"time"

	"w3s/go-backend/domain"
)

type profileUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (pu *profileUsecase) GetProfileByID(c context.Context, userID string) (*domain.Profile, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	user, err := pu.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &domain.Profile{ID: user.ID, Name: user.Name, Email: user.Email, CreatedAt: user.CreatedAt}, nil
}
