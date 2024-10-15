package user

import (
	"context"
	"github.com/Justdanru/bhs-test/internal/models"
	"github.com/Justdanru/bhs-test/internal/usecase/repository"
	"github.com/Justdanru/bhs-test/internal/usecase/service"
)

type Service struct {
	userRepository repository.UserRepository
}

func NewService(
	userRepository repository.UserRepository,
) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

func (s *Service) User(ctx context.Context, filter service.UserFilter) (*models.User, error) {
	return s.userRepository.Get(ctx, repository.GetFilter{
		Id:       filter.Id,
		Username: filter.Username,
		Limit:    filter.Limit,
		Offset:   filter.Offset,
	})
}

func (s *Service) Register(ctx context.Context, username string, password string) (*models.User, error) {
	user, err := models.NewUser(username, password)
	if err != nil {
		return nil, err
	}

	user, err = s.userRepository.Add(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
