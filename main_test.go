package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/WarisLi/Golang-mini-project/core"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	mockProductRepo := new(core.MockProductRepository)
	mockUserRepo := new(core.MockUserRepository)
	app := setupApp(mockProductRepo, mockUserRepo)

	mockUserRepo.On("Create", mock.AnythingOfType("core.User")).Return(nil)

	tests := []struct {
		description  string
		requestBody  core.UsernamePassword
		expectStatus int
	}{
		{
			description:  "Valid input",
			requestBody:  core.UsernamePassword{Username: "mock_user_1", Password: "1234"},
			expectStatus: fiber.StatusCreated,
		},
		{
			description:  "Missing param",
			requestBody:  core.UsernamePassword{Username: "mock_user_2"},
			expectStatus: fiber.StatusBadRequest,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			reqBody, _ := json.Marshal(test.requestBody)
			req := httptest.NewRequest("POST", "/user", bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, test.expectStatus, resp.StatusCode)
		})
	}
	mockUserRepo.AssertExpectations(t)
}

func TestLoginUser(t *testing.T) {
	mockProductRepo := new(core.MockProductRepository)
	mockUserRepo := new(core.MockUserRepository)
	app := setupApp(mockProductRepo, mockUserRepo)

	validInput := core.User{Username: "mock_user_1", Password: "Pass@12345"}
	invalidInput := core.User{Username: "mock_user_1", Password: "Abc@7890"}
	mockUserRepo.On("ValidateUser", validInput).Return(nil)
	mockUserRepo.On("ValidateUser", invalidInput).Return(errors.New("Invalid input"))

	tests := []struct {
		description  string
		requestBody  core.UsernamePassword
		expectStatus int
	}{
		{
			description:  "Valid input",
			requestBody:  core.UsernamePassword{Username: "mock_user_1", Password: "Pass@12345"},
			expectStatus: fiber.StatusOK,
		},
		{
			description:  "Missing param",
			requestBody:  core.UsernamePassword{Username: "mock_user_2"},
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid username or password",
			requestBody:  core.UsernamePassword{Username: "mock_user_1", Password: "Abc@7890"},
			expectStatus: fiber.StatusUnauthorized,
		},
	}

	// Run tests
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			reqBody, _ := json.Marshal(test.requestBody)
			req := httptest.NewRequest("POST", "/user/login", bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			assert.Equal(t, test.expectStatus, resp.StatusCode)
		})
	}
	mockUserRepo.AssertExpectations(t)
}

func generateMockJWT() string {
	claims := jwt.MapClaims{
		"username": "mock_user",
		"role":     "admin",
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return t
}

func TestGetProducts(t *testing.T) {
	mockProductRepo := new(core.MockProductRepository)
	mockUserRepo := new(core.MockUserRepository)
	app := setupApp(mockProductRepo, mockUserRepo)
	token := generateMockJWT()

	mockProduct := []core.Product{
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
	mockProductRepo := new(core.MockProductRepository)
	mockUserRepo := new(core.MockUserRepository)
	app := setupApp(mockProductRepo, mockUserRepo)
	token := generateMockJWT()

	mockProduct := &core.Product{Name: "Mock product 1", Quantity: 200}
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
	mockProductRepo := new(core.MockProductRepository)
	mockUserRepo := new(core.MockUserRepository)
	app := setupApp(mockProductRepo, mockUserRepo)
	token := generateMockJWT()

	validInput := core.Product{Name: "Book A", Quantity: 1000}
	mockProductRepo.On("Save", validInput).Return(nil)

	tests := []struct {
		description  string
		requestBody  core.ProductInput
		expectStatus int
	}{
		{
			description:  "Valid input",
			requestBody:  core.ProductInput{Name: "Book A", Quantity: 1000},
			expectStatus: fiber.StatusCreated,
		},
		{
			description:  "Missing param",
			requestBody:  core.ProductInput{Name: "Book A"},
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid Quantity",
			requestBody:  core.ProductInput{Name: "Book A", Quantity: 0},
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
	mockProductRepo := new(core.MockProductRepository)
	mockUserRepo := new(core.MockUserRepository)
	app := setupApp(mockProductRepo, mockUserRepo)
	token := generateMockJWT()

	validInput := core.Product{ID: 1000, Name: "Book A", Quantity: 200}
	mockProductRepo.On("Update", validInput).Return(nil)

	tests := []struct {
		description  string
		requestBody  core.ProductInput
		pathParam    int
		expectStatus int
	}{
		{
			description:  "Valid input",
			requestBody:  core.ProductInput{Name: "Book A", Quantity: 200},
			pathParam:    1000,
			expectStatus: fiber.StatusOK,
		},
		{
			description:  "Missing param",
			requestBody:  core.ProductInput{Name: "Book A"},
			pathParam:    1000,
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid Quantity",
			requestBody:  core.ProductInput{Name: "Book A", Quantity: -5},
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
	mockProductRepo := new(core.MockProductRepository)
	mockUserRepo := new(core.MockUserRepository)
	app := setupApp(mockProductRepo, mockUserRepo)
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
