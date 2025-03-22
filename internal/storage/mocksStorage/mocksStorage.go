package mocksStorage

import (
	"GraphQL-project/internal/storage/mocksStorage/mocks"
	"GraphQL-project/internal/storage/repsInterfaces"
)

type TestStorage struct {
	UserMock *mocks.User
}

func New() *TestStorage {
	return &TestStorage{
		UserMock: new(mocks.User),
	}
}

func (s *TestStorage) User() repsInterfaces.User {
	return s.UserMock
}
