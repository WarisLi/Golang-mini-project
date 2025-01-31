package core

import (
	"github.com/stretchr/testify/mock"
)

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetOne(id uint) (*Product, error)
	Save(product Product) error
	Update(product Product) error
	Delete(id uint) error
}

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetAll() ([]Product, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]Product), args.Error(1)
}

func (m *MockProductRepository) GetOne(id uint) (*Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Product), args.Error(1)
}

func (m *MockProductRepository) Save(product Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Update(product Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
