package http

import (
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/WarisLi/Golang-mini-project/internal/adapters/http/middleware"
	"github.com/WarisLi/Golang-mini-project/internal/core/ports"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App, productRepo ports.ProductRepository, userRepo ports.UserRepository) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// Middleware to log request information
	app.Use(middleware.AppLogger)

	productService := ports.NewProductService(productRepo)
	productHandler := NewHttpProductHandler(productService)

	userService := ports.NewUserService(userRepo)
	userHandler := NewHttpUserHandler(userService)

	userGroup := app.Group("/user")
	userGroup.Post("", userHandler.CreateUser)
	userGroup.Post("/login", userHandler.LoginUser)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	}))
	// Middleware to extract user data from JWT
	app.Use(middleware.JWTAuthMiddleware)

	productGroup := app.Group("/product")
	productGroup.Use(middleware.CheckRole)
	productGroup.Get("", productHandler.GetProducts)
	productGroup.Get("/:id", productHandler.GetProduct)
	productGroup.Post("", productHandler.CreateProduct)
	productGroup.Put("/:id", productHandler.UpdateProduct)
	productGroup.Delete("/:id", productHandler.DeleteProduct)
}
