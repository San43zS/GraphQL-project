package repsInterfaces

import (
	"GraphQL-project/internal/models"
	"context"
)

//go:generate mockery --name User --output=./../mocksStorage/mocks
type User interface {
	Create(ctx context.Context, usr *models.User) (*models.User, error)
	Get(ctx context.Context, id string) (*models.User, error)
	Delete(ctx context.Context, id string) error
}
