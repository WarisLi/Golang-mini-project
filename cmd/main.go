package main

import (
	_ "github.com/WarisLi/Golang-mini-project/docs"
	"github.com/WarisLi/Golang-mini-project/internal/adapters/database"
	"github.com/WarisLi/Golang-mini-project/internal/adapters/http"
	"github.com/WarisLi/Golang-mini-project/internal/config"
	"github.com/gofiber/fiber/v2"

	_ "github.com/lib/pq"
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
	db := config.SetupDB()
	productRepo := database.NewGormProductRepository(db)
	userRepo := database.NewGormUserRepository(db)

	app := fiber.New()
	http.SetupRoutes(app, productRepo, userRepo)

	app.Listen(":8080")
}
