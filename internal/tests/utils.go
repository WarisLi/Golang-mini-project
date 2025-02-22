package tests

import (
	"os"
	"time"

	"github.com/WarisLi/Golang-mini-project/internal/adapters/http"
	"github.com/WarisLi/Golang-mini-project/internal/tests/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func setupAppTest() (*fiber.App, *mocks.MockProductRepository, *mocks.MockUserRepository) {
	app := fiber.New()
	mockProductRepo := new(mocks.MockProductRepository)
	mockUserRepo := new(mocks.MockUserRepository)
	http.SetupRoutes(app, mockProductRepo, mockUserRepo)

	return app, mockProductRepo, mockUserRepo
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
