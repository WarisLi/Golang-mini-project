package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/WarisLi/Golang-mini-project/internal/core/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestCreateUser(t *testing.T) {
	app, _, mockUserRepo := setupAppTest()

	mockUserRepo.On("Create", mock.AnythingOfType("models.User")).Return(nil)

	tests := []struct {
		description  string
		requestBody  models.UsernamePassword
		expectStatus int
	}{
		{
			description:  "Valid input",
			requestBody:  models.UsernamePassword{Username: "mock_user_1", Password: "1234"},
			expectStatus: fiber.StatusCreated,
		},
		{
			description:  "Missing param",
			requestBody:  models.UsernamePassword{Username: "mock_user_2"},
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
	app, _, mockUserRepo := setupAppTest()

	validUsername := "mock_user_1"
	validPassword, _ := bcrypt.GenerateFromPassword([]byte("Pass@12345"), bcrypt.DefaultCost)

	invalidInput := "mock_user_2"

	mockUserRepo.On("GetUser", validUsername).Return(&models.User{Username: validUsername, Password: string(validPassword)}, nil)
	mockUserRepo.On("GetUser", invalidInput).Return(&models.User{}, errors.New("Invalid input"))

	tests := []struct {
		description  string
		requestBody  models.UsernamePassword
		expectStatus int
	}{
		{
			description:  "Valid input",
			requestBody:  models.UsernamePassword{Username: "mock_user_1", Password: "Pass@12345"},
			expectStatus: fiber.StatusOK,
		},
		{
			description:  "Invalid password",
			requestBody:  models.UsernamePassword{Username: "mock_user_1", Password: "WrongPass"},
			expectStatus: fiber.StatusUnauthorized,
		},
		{
			description:  "Missing param",
			requestBody:  models.UsernamePassword{Username: "mock_user_1"},
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description:  "Invalid username",
			requestBody:  models.UsernamePassword{Username: "mock_user_2", Password: "Abc@7890"},
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
