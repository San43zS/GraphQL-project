package storage

import (
	"GraphQL-project/internal/storage/repsInterfaces"
)

type Storage interface {
	User() repsInterfaces.User
}
