package http

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/WarisLi/Golang-mini-project/internal/adapters/http/middleware"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/swagger"
)

func SetupRoutes(
	app *fiber.App,
	productHandler *HttpProductHandler,
	userHandler *HttpUserHandler,
) {
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	// Middleware to log request information
	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Bangkok",
	}))

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
