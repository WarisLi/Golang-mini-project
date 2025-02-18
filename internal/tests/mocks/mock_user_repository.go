package mocks

import (
	"github.com/WarisLi/Golang-mini-project/internal/core/models"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user models.User) error {
	args := m.Called(user)

	return args.Error(0)
}

func (m *MockUserRepository) ValidateUser(requestUser models.User) error {
	args := m.Called(requestUser)

	return args.Error(0)
}
