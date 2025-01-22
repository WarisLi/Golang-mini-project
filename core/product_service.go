package core

import "errors"

type ProductService interface {
	GetProducts() ([]Product, error)
	GetProduct(id uint) (*Product, error)
	CreateProduct(product Product) error
	UpdateProduct(product Product) error
	DeleteProduct(id uint) error
}

type productServiceImpl struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return &productServiceImpl{repo: repo}
}

func (s *productServiceImpl) GetProducts() ([]Product, error) {
	// Business logic
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *productServiceImpl) GetProduct(id uint) (*Product, error) {
	// Business logic
	product, err := s.repo.GetOne(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productServiceImpl) CreateProduct(product Product) error {
	// Business logic
	if product.Quantity <= 0 {
		return errors.New("quantity must be positive")
	}

	if err := s.repo.Save(product); err != nil {
		return err
	}

	return nil
}

func (s *productServiceImpl) UpdateProduct(product Product) error {
	// Business logic
	if product.Quantity <= 0 {
		return errors.New("quantity must be positive")
	}

	if err := s.repo.Update(product); err != nil {
		return err
	}

	return nil
}

func (s *productServiceImpl) DeleteProduct(id uint) error {
	// Business logic
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}
