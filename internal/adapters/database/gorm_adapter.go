package database

import (
	"errors"

	"github.com/WarisLi/Golang-mini-project/internal/core/models"
	"github.com/WarisLi/Golang-mini-project/internal/core/ports"
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormProductRepository(db *gorm.DB) ports.ProductRepository {
	return &GormRepository{db: db}
}

func NewGormUserRepository(db *gorm.DB) ports.UserRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) GetAll() ([]models.Product, error) {
	var products []models.Product

	if result := r.db.Find(&products); result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *GormRepository) GetOne(id uint) (*models.Product, error) {
	var product models.Product

	if result := r.db.First(&product, id); result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *GormRepository) Save(product models.Product) error {
	if result := r.db.Create(&product); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *GormRepository) Update(product models.Product) error {
	if result := r.db.Model(&product).Updates(product); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormRepository) Delete(id uint) error {
	var product models.Product
	result := r.db.Delete(&product, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected <= 0 {
		return errors.New("Delete failed")
	}
	return nil
}

func (r *GormRepository) GetUser(username string) (*models.User, error) {
	var user models.User
	result := r.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *GormRepository) Create(user models.User) error {
	if result := r.db.Create(&user); result.Error != nil {
		return result.Error
	}

	return nil
}
