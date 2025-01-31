package core

import "github.com/stretchr/testify/mock"

// secondary port
type UserRepository interface {
	Create(user User) error
	ValidateUser(requestUser User) error
}

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user User) error {
	args := m.Called(user)

	return args.Error(0)
}

func (m *MockUserRepository) ValidateUser(requestUser User) error {
	args := m.Called(requestUser)

	return args.Error(0)
}
