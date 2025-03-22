package scenarios

import (
	"GraphQL-project/internal/models"
	"GraphQL-project/internal/storage"
	"context"
	"errors"
)

var (
	userNotFoundErr = errors.New("user not found")
	emailEmptyErr   = errors.New("email is empty")
)

type User interface {
	CreateUser(ctx context.Context, usr *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id string) error
	GetUser(ctx context.Context, id string) (*models.User, error)
}

type Scenarios struct {
	storage storage.Storage
}

func New(storage storage.Storage) User {
	return &Scenarios{
		storage: storage,
	}
}

func (s *Scenarios) CreateUser(ctx context.Context, usr *models.User) (*models.User, error) {
	if usr.Email == "" {
		return nil, emailEmptyErr
	}

	newUser, err := s.storage.User().Create(ctx, usr)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (s *Scenarios) GetUser(ctx context.Context, id string) (*models.User, error) {
	usr, err := s.storage.User().Get(ctx, id)
	if err != nil {
		return nil, userNotFoundErr
	}
	return usr, err
}

func (s *Scenarios) DeleteUser(ctx context.Context, id string) error {
	if err := s.storage.User().Delete(ctx, id); err != nil {
		return userNotFoundErr
	}

	return nil
}
