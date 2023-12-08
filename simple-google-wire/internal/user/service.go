package user

import (
	"context"

	"github.com/nebula-aac/public-projects/simple-google-wire/internal/domain"
)

type service struct {
	repo domain.UserRepository
}

func (s *service) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	userEntity, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	// Convert UserEntity to User, assuming the structures are similar
	user := &domain.User{
		ID:       userEntity.ID,
		Username: userEntity.Username,
		Password: userEntity.Password,
		Email:    userEntity.Email,
	}

	return user, nil
}
