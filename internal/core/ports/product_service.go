package ports

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/WarisLi/Golang-mini-project/internal/core/models"
)

// primary port
type ProductService interface {
	GetProducts() ([]models.Product, error)
	GetProduct(id uint) (*models.Product, error)
	CreateProduct(productInput models.ProductInput) error
	UpdateProduct(id uint, productInput models.ProductInput) error
	DeleteProduct(id uint) error
}

type productServiceImpl struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) ProductService {
	return &productServiceImpl{repo: repo}
}

// business logic
func (s *productServiceImpl) GetProducts() ([]models.Product, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *productServiceImpl) GetProduct(id uint) (*models.Product, error) {
	product, err := s.repo.GetOne(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productServiceImpl) CreateProduct(productInput models.ProductInput) error {
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
	var product models.Product
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

func (s *productServiceImpl) UpdateProduct(id uint, productInput models.ProductInput) error {
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
	var product models.Product
	err = json.Unmarshal(data, &product)
	if err != nil {
		fmt.Println("Error unmarshalling to Product:", err)
		return err
	}

	product.ID = id

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
