package core

type ProductRepository interface {
	GetAll() ([]Product, error)
	GetOne(id uint) (*Product, error)
	Save(product Product) error
	Update(product Product) error
	Delete(id uint) error
}
