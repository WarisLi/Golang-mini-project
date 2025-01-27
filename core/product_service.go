package core

import (
	"encoding/json"
	"errors"
	"fmt"
)

// primary port
type ProductService interface {
	GetProducts() ([]Product, error)
	GetProduct(id uint) (*Product, error)
	CreateProduct(productInput ProductInput) error
	UpdateProduct(product Product) error
	DeleteProduct(id uint) error
}

type productServiceImpl struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return &productServiceImpl{repo: repo}
}

// business logic
func (s *productServiceImpl) GetProducts() ([]Product, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *productServiceImpl) GetProduct(id uint) (*Product, error) {
	product, err := s.repo.GetOne(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productServiceImpl) CreateProduct(productInput ProductInput) error {
	if productInput.Quantity <= 0 {
		return errors.New("quantity must be positive")
	}

	// Convert ProductInput to JSON
	data, err := json.Marshal(productInput)
	if err != nil {
		fmt.Println("Error marshalling ProductInput:", err)
		return err
	}

	// Convert JSON to Product
	var product Product
	err = json.Unmarshal(data, &product)
	if err != nil {
		fmt.Println("Error unmarshalling to Product:", err)
		return err
	}

	if err := s.repo.Save(product); err != nil {
		return err
	}

	return nil
}

func (s *productServiceImpl) UpdateProduct(product Product) error {
	if product.Quantity <= 0 {
		return errors.New("quantity must be positive")
	}

	if err := s.repo.Update(product); err != nil {
		return err
	}

	return nil
}

func (s *productServiceImpl) DeleteProduct(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}
