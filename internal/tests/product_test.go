package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/WarisLi/Golang-mini-project/internal/core/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	app, mockProductRepo, _ := setupAppTest()
	token := generateMockJWT()

	mockProduct := []models.Product{
		{Name: "Mock product 1", Quantity: 200},
		{Name: "Mock product 2", Quantity: 100},
	}
	mockProductRepo.On("GetAll").Return(mockProduct, nil)

	tests := []struct {
		description  string
		expectStatus int
	}{
		{
			description:  "Valid case",
			expectStatus: fiber.StatusOK,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/product", nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			resp, _ := app.Test(req)

			assert.Equal(t, test.expectStatus, resp.StatusCode)
		})
	}
	mockProductRepo.AssertExpectations(t)
}

func TestGetProduct(t *testing.T) {
	app, mockProductRepo, _ := setupAppTest()
	token := generateMockJWT()

	mockProduct := &models.Product{Name: "Mock product 1", Quantity: 200}
	mockProductRepo.On("GetOne", uint(1000)).Return(mockProduct, nil)
	mockProductRepo.On("GetOne", uint(999999)).Return(nil, errors.New("product not found"))

	tests := []struct {
		description  string
		pathParam    int
		expectStatus int
	}{
		{
			description:  "Valid input",
			pathParam:    1000,
			expectStatus: fiber.StatusOK,
		},
		{
			description:  "Not found",
			pathParam:    999999,
			expectStatus: fiber.StatusBadRequest,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req := httptest.NewRequest("GET", fmt.Sprintf("/product/%d", test.pathParam), nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			resp, _ := app.Test(req)

			assert.Equal(t, test.expectStatus, resp.StatusCode)
		})
	}
	mockProductRepo.AssertExpectations(t)
}

func TestCreateProduct(t *testing.T) {
	app, mockProductRepo, _ := setupAppTest()
	token := generateMockJWT()

	validInput := models.Product{Name: "Book A", Quantity: 1000}
	mockProductRepo.On("Save", validInput).Return(nil)

	tests := []struct {
		description  string
		requestBody  models.ProductInput
		expectStatus int
	}{
		{
			description:  "Valid input",
			requestBody:  models.ProductInput{Name: "Book A", Quantity: 1000},
			expectStatus: fiber.StatusCreated,
		},
		{
			description:  "Missing param",
			requestBody:  models.ProductInput{Name: "Book A"},
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid Quantity",
			requestBody:  models.ProductInput{Name: "Book A", Quantity: 0},
			expectStatus: fiber.StatusBadRequest,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			reqBody, _ := json.Marshal(test.requestBody)
			req := httptest.NewRequest("POST", "/product", bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			resp, _ := app.Test(req)

			assert.Equal(t, test.expectStatus, resp.StatusCode)
		})
	}
	mockProductRepo.AssertExpectations(t)
}

func TestUpdateProducts(t *testing.T) {
	app, mockProductRepo, _ := setupAppTest()
	token := generateMockJWT()

	validInput := models.Product{ID: 1000, Name: "Book A", Quantity: 200}
	mockProductRepo.On("Update", validInput).Return(nil)

	tests := []struct {
		description  string
		requestBody  models.ProductInput
		pathParam    int
		expectStatus int
	}{
		{
			description:  "Valid input",
			requestBody:  models.ProductInput{Name: "Book A", Quantity: 200},
			pathParam:    1000,
			expectStatus: fiber.StatusOK,
		},
		{
			description:  "Missing param",
			requestBody:  models.ProductInput{Name: "Book A"},
			pathParam:    1000,
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid Quantity",
			requestBody:  models.ProductInput{Name: "Book A", Quantity: -5},
			pathParam:    1000,
			expectStatus: fiber.StatusBadRequest,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			reqBody, _ := json.Marshal(test.requestBody)
			req := httptest.NewRequest("PUT", fmt.Sprintf("/product/%d", test.pathParam), bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			resp, _ := app.Test(req)

			assert.Equal(t, test.expectStatus, resp.StatusCode)
		})
	}
	mockProductRepo.AssertExpectations(t)
}

func TestDeleteProducts(t *testing.T) {
	app, mockProductRepo, _ := setupAppTest()
	token := generateMockJWT()

	mockProductRepo.On("Delete", uint(1000)).Return(nil)
	mockProductRepo.On("Delete", uint(9999)).Return(errors.New("Not found"))

	tests := []struct {
		description  string
		pathParam    int
		expectStatus int
	}{
		{
			description:  "Valid input",
			pathParam:    1000,
			expectStatus: fiber.StatusOK,
		},
		{
			description:  "Invalid input",
			pathParam:    9999,
			expectStatus: fiber.StatusBadRequest,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req := httptest.NewRequest("DELETE", fmt.Sprintf("/product/%d", test.pathParam), nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
			resp, _ := app.Test(req)

			assert.Equal(t, test.expectStatus, resp.StatusCode)
		})
	}
	mockProductRepo.AssertExpectations(t)
}
