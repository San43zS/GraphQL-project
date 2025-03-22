package graph

import (
	"GraphQL-project/internal/models"
	"context"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserList
}

type UserList interface {
	GetUser(ctx context.Context, id string) (*models.User, error)
}
