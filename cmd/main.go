package main

import (
	"os"

	_ "github.com/WarisLi/Golang-mini-project/docs"
	"github.com/WarisLi/Golang-mini-project/internal/adapters/database"
	"github.com/WarisLi/Golang-mini-project/internal/adapters/http"
	"github.com/WarisLi/Golang-mini-project/internal/adapters/producer"
	"github.com/WarisLi/Golang-mini-project/internal/config"
	"github.com/WarisLi/Golang-mini-project/internal/core/ports"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
	"gopkg.in/Shopify/sarama.v1"
)

// @title           Swagger API
// @version         1.0
// @description     This is a sample server for a Product API.

// @host      localhost:8080
// @BasePath /
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db := config.SetupDB()

	servers := []string{os.Getenv("KAFKA_SERVERS")}
	saramaProducer, err := sarama.NewSyncProducer(servers, nil)
	if err != nil {
		panic(err)
	}
	defer saramaProducer.Close()

	productRepo := database.NewGormProductRepository(db)
	userRepo := database.NewGormUserRepository(db)

	eventProducer := producer.NewEventProducer(saramaProducer)

	productService := ports.NewProductService(productRepo, eventProducer)
	productHandler := http.NewHttpProductHandler(productService)

	userService := ports.NewUserService(userRepo)
	userHandler := http.NewHttpUserHandler(userService)

	app := fiber.New()
	http.SetupRoutes(app, productHandler, userHandler)

	app.Listen(":8080")
}
