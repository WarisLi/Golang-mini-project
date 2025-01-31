package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/WarisLi/Golang-mini-project/adapters"
	"github.com/WarisLi/Golang-mini-project/core"
	"github.com/WarisLi/Golang-mini-project/database"
	_ "github.com/WarisLi/Golang-mini-project/docs"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserData represents the user data extracted from the JWT token
type UserData struct {
	Username string
	Role     string
}

// userContextKey is the key used to store user data in the Fiber context
const userContextKey = "user"

// extractUserFromJWT is a middleware that extracts user data from the JWT token
func extractUserFromJWT(c *fiber.Ctx) error {
	user := &UserData{}

	// Extract the token from the Fiber context (inserted by the JWT middleware)
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	user.Username = claims["username"].(string)
	user.Role = claims["role"].(string)

	// Store the user data in the Fiber context
	c.Locals(userContextKey, user)

	return c.Next()
}

func checkRole(c *fiber.Ctx) error {
	user := c.Locals(userContextKey).(*UserData)

	if user.Role != "admin" {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}

func appLogger(c *fiber.Ctx) error {
	start := time.Now()
	fmt.Printf("URL = %s, Method = %s, Time = %s\n", c.OriginalURL(), c.Method(), start)

	return c.Next()
}

func setupApp(productRepo core.ProductRepository, userRepo core.UserRepository) *fiber.App {
	app := fiber.New()

	if err := godotenv.Load(); err != nil {
		log.Fatal("load .env error")
	}

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Use(appLogger)

	productService := core.NewProductService(productRepo)
	productHandler := adapters.NewHttpProductHandler(productService)

	userService := core.NewUserService(userRepo)
	userHandler := adapters.NewHttpUserHandler(userService)

	userGroup := app.Group("/user")
	userGroup.Post("", userHandler.CreateUser)
	userGroup.Post("/login", userHandler.LoginUser)

	// Middleware to extract user data from JWT
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	}))
	app.Use(extractUserFromJWT)

	productGroup := app.Group("/product")
	productGroup.Use(checkRole)
	productGroup.Get("", productHandler.GetProducts)
	productGroup.Get("/:id", productHandler.GetProduct)
	productGroup.Post("", productHandler.CreateProduct)
	productGroup.Put("/:id", productHandler.UpdateProduct)
	productGroup.Delete("/:id", productHandler.DeleteProduct)

	return app
}

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
	db := database.SetupDB()
	productRepo := adapters.NewGormProductRepository(db)
	userRepo := adapters.NewGormUserRepository(db)

	app := setupApp(productRepo, userRepo)

	app.Listen(":8080")
}
