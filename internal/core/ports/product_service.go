package ports

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	events "github.com/WarisLi/Golang-shared-events"

	"github.com/WarisLi/Golang-mini-project/internal/adapters/producer"
	"github.com/WarisLi/Golang-mini-project/internal/core/models"
)

type ProductService interface {
	GetProducts() ([]models.Product, error)
	GetProduct(id uint) (*models.Product, error)
	CreateProduct(productInput models.ProductInput) error
	UpdateProduct(id uint, productInput models.ProductInput) error
	DeleteProduct(id uint) error
}

type productServiceImpl struct {
	repo          ProductRepository
	eventProducer producer.EventProducer
}

func NewProductService(repo ProductRepository, eventProducer producer.EventProducer) ProductService {
	return &productServiceImpl{
		repo:          repo,
		eventProducer: eventProducer,
	}
}

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

	if product.Quantity < 100 {
		event := events.LowProductQuantityNotificationEvent{
			Name:     product.Name,
			Quantity: product.Quantity,
		}
		err := s.eventProducer.Produce(event)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

func (s *productServiceImpl) DeleteProduct(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}
