package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/WarisLi/Golang-mini-project/core"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	app := setup()

	tests := []struct {
		description  string
		requestBody  core.UsernamePassword
		expectStatus int
	}{
		{
			description:  "Valid input",
			requestBody:  core.UsernamePassword{Username: "test_user_1", Password: "1234"},
			expectStatus: fiber.StatusCreated,
		},
		{
			description:  "Missing param",
			requestBody:  core.UsernamePassword{Username: "test_user_2"},
			expectStatus: fiber.StatusBadRequest,
		},
		{
			description:  "Duplicate username",
			requestBody:  core.UsernamePassword{Username: "test_user_1", Password: "12345"},
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
}
