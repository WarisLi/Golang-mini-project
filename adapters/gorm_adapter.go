package adapters

import (
	"github.com/WarisLi/Golang-mini-project/core"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

// connect secondary port
func NewGormProductRepository(db *gorm.DB) core.ProductRepository {
	return &GormRepository{db: db}
}

func NewGormUserRepository(db *gorm.DB) core.UserRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) GetAll() ([]core.Product, error) {
	var products []core.Product

	if result := r.db.Find(&products); result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *GormRepository) GetOne(id uint) (*core.Product, error) {
	var product core.Product

	if result := r.db.First(&product, id); result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (r *GormRepository) Save(product core.Product) error {
	if result := r.db.Create(&product); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *GormRepository) Update(product core.Product) error {
	if result := r.db.Model(&product).Updates(product); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormRepository) Delete(id uint) error {
	var product core.Product
	if result := r.db.Delete(&product, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *GormRepository) Create(user core.User) error {
	if result := r.db.Create(&user); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *GormRepository) ValidateUser(requestUser core.User) error {
	var user core.User
	result := r.db.Where("username = ?", requestUser.Username).First(&user)
	if result.Error != nil {
		return result.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestUser.Password))
	if err != nil {
		return err
	}
	return nil
}
