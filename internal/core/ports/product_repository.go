package ports

import (
	"github.com/WarisLi/Golang-mini-project/internal/core/models"
	"github.com/stretchr/testify/mock"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetOne(id uint) (*models.Product, error)
	Save(product models.Product) error
	Update(product models.Product) error
	Delete(id uint) error
}

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetAll() ([]models.Product, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductRepository) GetOne(id uint) (*models.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Product), args.Error(1)
}

func (m *MockProductRepository) Save(product models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Update(product models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}
