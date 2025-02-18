package ports

import (
	"github.com/WarisLi/Golang-mini-project/internal/core/models"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetOne(id uint) (*models.Product, error)
	Save(product models.Product) error
	Update(product models.Product) error
	Delete(id uint) error
}
