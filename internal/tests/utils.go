package tests

import (
	"os"
	"time"

	"github.com/WarisLi/Golang-mini-project/internal/adapters/http"
	"github.com/WarisLi/Golang-mini-project/internal/adapters/producer"
	"github.com/WarisLi/Golang-mini-project/internal/core/ports"
	"github.com/WarisLi/Golang-mini-project/internal/tests/mocks"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"gopkg.in/Shopify/sarama.v1"
)

func setupAppTest() (*fiber.App, *mocks.MockProductRepository, *mocks.MockUserRepository) {
	err := godotenv.Load("../../.env")
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	mockProductRepo := new(mocks.MockProductRepository)
	mockUserRepo := new(mocks.MockUserRepository)

	servers := []string{os.Getenv("KAFKA_SERVERS")}
	saramaProducer, err := sarama.NewSyncProducer(servers, nil)
	if err != nil {
		panic(err)
	}
	defer saramaProducer.Close()

	eventProducer := producer.NewEventProducer(saramaProducer)

	productService := ports.NewProductService(mockProductRepo, eventProducer)
	productHandler := http.NewHttpProductHandler(productService)

	userService := ports.NewUserService(mockUserRepo)
	userHandler := http.NewHttpUserHandler(userService)

	http.SetupRoutes(app, productHandler, userHandler)

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
