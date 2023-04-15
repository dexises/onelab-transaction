package service

import (
	"context"
	"errors"
	"onelab/internal/model"
	"onelab/internal/repository"
)

type IUserService interface {
	Create(ctx context.Context, user model.UserCreate) error
	Get(ctx context.Context, id int) (model.User, error)
}

type UserService struct {
	repo *repository.Manager
}

func NewUserService(repo *repository.Manager) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(ctx context.Context, user model.UserCreate) error {
	// TODO user validation
	err := s.repo.User.Create(ctx, user)
	return err
}

func (s *UserService) Get(ctx context.Context, id int) (model.User, error) {
	if id < 1 {
		return model.User{}, errors.New("invalid id parameter")
	}

	user, err := s.repo.User.Get(ctx, id)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
