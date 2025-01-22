package adapters

import (
	"github.com/WarisLi/Golang-mini-project/core"
	"gorm.io/gorm"
)

type GormProductRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) core.ProductRepository {
	return &GormProductRepository{db: db}
}

func (r *GormProductRepository) GetAll() ([]core.Product, error) {
	var products []core.Product

	if result := r.db.Find(&products); result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *GormProductRepository) GetOne(id uint) (*core.Product, error) {
	var product core.Product

	if result := r.db.First(&product, id); result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *GormProductRepository) Save(product core.Product) error {
	if result := r.db.Create(&product); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *GormProductRepository) Update(product core.Product) error {
	if result := r.db.Model(&product).Updates(product); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormProductRepository) Delete(id uint) error {
	var product core.Product
	if result := r.db.Delete(&product, id); result.Error != nil {
		return result.Error
	}
	return nil
}
